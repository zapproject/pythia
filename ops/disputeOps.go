package ops

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zapproject/pythia/apiOracle"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	zap "github.com/zapproject/pythia/contracts"
	zap1 "github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/tracker"
	"github.com/zapproject/pythia/util"
)

/**
 * This file handles all operations related to disputes
 */

func Dispute(requestId *big.Int, timestamp *big.Int, minerIndex *big.Int, ctx context.Context) error {

	if !minerIndex.IsUint64() || minerIndex.Uint64() > 4 {
		return fmt.Errorf("miner index should be between 0 and 4 (got %s)", minerIndex.Text(10))
	}

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	addr := ctx.Value(zapCommon.PublicAddress).(common.Address)

	balance, err := instance.BalanceOf(nil, addr)
	if err != nil {
		return fmt.Errorf("failed to fetch balance: %s", err.Error())
	}
	var asBytes32 [32]byte
	copy(asBytes32[:], "0x8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc49239") //keccak256(disputeFee)
	disputeCost, err := instance.GetUintVar(nil, asBytes32)
	if err != nil {
		return fmt.Errorf("failed to get dispute cost: %s", err)
	}

	if balance.Cmp(disputeCost) < 0 {
		return fmt.Errorf("insufficient balance (%s ZAP) disputes require (%s ZAP)",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(disputeCost))
	}

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}

	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	tx, err := instance2.BeginDispute(auth, requestId, timestamp, minerIndex)
	if err != nil {
		return fmt.Errorf("failed to send dispute txn: %s", err.Error())
	}
	fmt.Printf("dispute started with txn: %s\n", tx.Hash().Hex())
	return nil
}

func Vote(_disputeId *big.Int, _supportsDispute bool, ctx context.Context) error {

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	addr := ctx.Value(zapCommon.PublicAddress).(common.Address)
	voted, err := instance.DidVote(nil, _disputeId, addr)
	if err != nil {
		return fmt.Errorf("failed to check if you've already voted: %v", err)
	}
	if voted {
		fmt.Printf("You have already voted on this dispute\n")
		return nil
	}

	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}
	tx, err := instance2.Vote(auth, _disputeId, _supportsDispute)
	if err != nil {
		return fmt.Errorf("failed to submit vote transaction: %s", err.Error())
	}

	fmt.Printf("Vote submitted with transaction %s\n", tx.Hash().Hex())
	return nil
}

func getNonceSubmissions(ctx context.Context, valueBlock *big.Int, dispute *zap1.ZapDisputeNewDispute) ([]*apiOracle.PriceStamp, error) {
	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	tokenAbi, err := abi.JSON(strings.NewReader(zap1.ZapLibraryABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse abi: %v", err)
	}
	contractAddress := ctx.Value(zapCommon.ContractAddress).(common.Address)
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)

	//just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi
	bar := bind.NewBoundContract(contractAddress, tokenAbi, nil, nil, nil)

	allVals, err := instance.GetSubmissionsByTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get other submitted values for dispute: %v", err)
	}

	allAddrs, err := instance.GetMinersByRequestIdAndTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get miner addresses for dispute: %v", err)
	}

	const blockStep = 67
	high := int64(valueBlock.Uint64())
	low := high - blockStep
	nonceSubmitID := tokenAbi.Events["NonceSubmitted"].ID()
	var timedValues [5]*apiOracle.PriceStamp
	// timedValues := make([]*apiOracle.PriceStamp, 5)
	found := 0
	for found < 5 {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(low),
			ToBlock:   big.NewInt(high),
			Addresses: []common.Address{contractAddress},
			Topics:    [][]common.Hash{{nonceSubmitID}},
		}

		logs, err := client.FilterLogs(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce logs: %v", err)
		}

		for _, l := range logs {
			nonceSubmit := zap1.ZapLibraryNonceSubmitted{}
			err := bar.UnpackLog(&nonceSubmit, "NonceSubmitted", l)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack into object: %v", err)
			}
			header, err := client.HeaderByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				return nil, fmt.Errorf("failed to get nonce block header: %v", err)
			}
			for i := 0; i < 5; i++ {
				if nonceSubmit.Miner == allAddrs[i] {
					valTime := time.Unix(int64(header.Time), 0)

					bigF := new(big.Float)
					bigF.SetInt(allVals[i])
					f, _ := bigF.Float64()

					timedValues[i] = &apiOracle.PriceStamp{
						Created:   valTime,
						PriceInfo: apiOracle.PriceInfo{Price: f},
					}
					found++
					break
				}
			}
		}
		high -= blockStep
		low = high - blockStep
	}
	return timedValues[:], nil
}

func List(ctx context.Context) error {
	cfg := config.GetConfig()

	tokenAbi, err := abi.JSON(strings.NewReader(zap1.ZapDisputeABI))
	if err != nil {
		return fmt.Errorf("failed to parse abi: %v", err)
	}
	contractAddress := ctx.Value(zapCommon.ContractAddress).(common.Address)
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)

	////just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi
	bar := bind.NewBoundContract(contractAddress, tokenAbi, nil, nil, nil)

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get latest eth block header: %v", err)
	}

	startBlock := big.NewInt(54) //big.NewInt(10e3 * 14)
	startBlock.Sub(header.Number, startBlock)
	newDisputeID := tokenAbi.Events["NewDispute"].ID()
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{newDisputeID}},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to filter eth logs: %v", err)
	}

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)

	fmt.Printf("There are currently %d open disputes\n", len(logs))
	fmt.Printf("-------------------------------------\n")
	for _, rawDispute := range logs {
		dispute := zap1.ZapDisputeNewDispute{}
		err := bar.UnpackLog(&dispute, "NewDispute", rawDispute)
		if err != nil {
			return fmt.Errorf("failed to unpack dispute event from logs: %v", err)
		}
		_, executed, votePassed, _, reportedAddr, reportingMiner, _, uintVars, currTally, err := instance.GetAllDisputeVars(nil, dispute.DisputeId)
		if err != nil {
			return fmt.Errorf("failed to get dispute details: %v", err)
		}

		votingEnds := time.Unix(uintVars[3].Int64(), 0)
		createdTime := votingEnds.Add(-7 * 24 * time.Hour)

		var descString string
		if executed {
			descString = "complete, "
			if votePassed {
				descString += "successful"
			} else {
				descString += "rejected"
			}
		} else {
			descString = "in progress"
		}

		fmt.Printf("Dispute %s (%s):\n", dispute.DisputeId.String(), descString)
		fmt.Printf("    Accused Party: %s\n", reportedAddr.Hex())
		fmt.Printf("    Disputed by: %s\n", reportingMiner.Hex())
		fmt.Printf("    Created on:  %s\n", createdTime.Format("3:04 PM January 02, 2006 MST"))
		fmt.Printf("    Fee: %s ZAP\n", util.FormatERC20Balance(uintVars[8]))
		fmt.Printf("    \n")
		fmt.Printf("    Value disputed for requestID %d:\n", dispute.RequestId.Uint64())

		allSubmitted, err := getNonceSubmissions(ctx, uintVars[5], &dispute)
		if err != nil {
			return fmt.Errorf("failed to get the values submitted by other miners for the disputed block: %v", err)
		}
		disputedValTime := allSubmitted[uintVars[6].Uint64()].Created

		for i := len(allSubmitted) - 1; i >= 0; i-- {
			sub := allSubmitted[i]
			valStr := fmt.Sprintf("%f\n", sub.Price)
			var pointerStr string
			if i == int(uintVars[6].Uint64()) {
				pointerStr = " <--disputed"
			}

			fmt.Printf("      %s @ %s%s\n", valStr, sub.Created.Format("3:04:05 PM"), pointerStr)
		}
		fmt.Printf("    \n")

		tmp := new(big.Float)
		tmp.SetInt(currTally)
		currTallyFloat, _ := tmp.Float64()
		tmp.SetInt(uintVars[7])
		currQuorum, _ := tmp.Float64()
		currTallyFloat += currQuorum
		currTallyRatio := currTallyFloat / 2 * currQuorum
		percentage := currTallyRatio * 100
		if percentage > 100 {
			percentage = 100
		}
		fmt.Printf("    Currently %.0f%% of %s ZAP support this dispute (%s votes)\n", percentage, util.FormatERC20Balance(uintVars[7]), uintVars[4])

		result := tracker.CheckValueAtTime(dispute.RequestId.Uint64(), uintVars[2], disputedValTime)
		if result == nil || len(result.Datapoints) < 0 {
			fmt.Printf("      No data available for recommendation\n")
			continue
		}
		fmt.Printf("      Recommendation:\n")
		fmt.Printf("      Vote %t\n", !result.WithinRange)
		fmt.Printf("      Submitted value %s, expected range %.0f to %0.f\n", uintVars[2].String(), result.Low, result.High)
		numToShow := 3
		if numToShow > len(result.Datapoints) {
			numToShow = len(result.Datapoints)
		}

		disputeTimeDuration := cfg.DisputeTimeDelta

		fmt.Printf("Based on %d locally saved datapoints within %.0f minutes (showing closest %d)\n",

			len(result.Datapoints), disputeTimeDuration.Minutes(), numToShow)

		minTotalDelta := time.Duration(math.MaxInt64)
		index := 0
		for i := 0; i < len(result.Datapoints)-numToShow; i++ {
			totalDelta := time.Duration(0)
			for j := 0; j < numToShow; j++ {
				delta := result.Times[i+j].Sub(disputedValTime)
				if delta < 0 {
					delta = -delta
				}
				totalDelta += delta
			}
			if totalDelta < minTotalDelta {
				minTotalDelta = totalDelta
				index = i
			}
		}
		for i := 0; i < numToShow; i++ {
			dp := result.Datapoints[index+i]
			t := result.Times[index+i]
			fmt.Printf("        %f, ", dp)
			delta := disputedValTime.Sub(t)
			if delta > 0 {
				fmt.Printf("%.0fs before\n", delta.Seconds())
			} else {
				fmt.Printf("%.0fs after\n", (-delta).Seconds())
			}
		}
		fmt.Printf("\n")
	}

	return nil
}

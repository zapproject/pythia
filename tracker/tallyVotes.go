package tracker

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/rpc"
)

type TallyVotesTracker struct {
	lastCheckedBlock uint64
}

func (b *TallyVotesTracker) String() string {
	return "TallyVotesTracker"
}

func (b *TallyVotesTracker) Exec(ctx context.Context) error {

	zapAbi, err := abi.JSON(strings.NewReader(contracts1.ZapABI))
	if err != nil {
		return fmt.Errorf("failed to parse abi: %v", err)
	}

	disputeAbi, err := abi.JSON(strings.NewReader(contracts1.ZapDisputeABI))
	if err != nil {
		return fmt.Errorf("failed to parse abi: %v", err)
	}

	contractAddress := ctx.Value(zapCommon.ContractAddress).(common.Address)
	instanceM := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	instanceZ := ctx.Value(zapCommon.TransactorContractContextKey).(*contracts1.ZapTransactor)

	bar1 := bind.NewBoundContract(contractAddress, zapAbi, nil, nil, nil)
	bar2 := bind.NewBoundContract(contractAddress, disputeAbi, nil, nil, nil)

	newDisputeID := zapAbi.Events["NewDispute"].ID
	disputeVoteTalliedId := disputeAbi.Events["DisputeVoteTallied"].ID

	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get latest eth block header: %v", err)
	}
	if b.lastCheckedBlock == 0 {
		b.lastCheckedBlock = 1
	}

	toCheck := header.Number.Uint64()

	const blockDelay = 70
	if toCheck-b.lastCheckedBlock < blockDelay {
		return nil
	}

	checkUntil := toCheck

	query1 := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(b.lastCheckedBlock)),
		// FromBlock: nil,
		ToBlock: big.NewInt(int64(checkUntil)),
		// ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{newDisputeID}},
	}
	logs1, err := client.FilterLogs(ctx, query1)
	if err != nil {
		return fmt.Errorf("failed to filter eth logs: %v", err)
	}
	// fmt.Println("Logs1: ", logs1)

	query2 := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(b.lastCheckedBlock)),
		// FromBlock: nil,
		ToBlock: big.NewInt(int64(checkUntil)),
		// ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{disputeVoteTalliedId}},
	}
	logs2, err := client.FilterLogs(ctx, query2)
	if err != nil {
		return fmt.Errorf("failed to filter eth logs: %v", err)
	}

	lastDispute := contracts1.ZapDisputeDisputeVoteTallied{}
	if len(logs2) > 0 {
		err = bar2.UnpackLog(&lastDispute, "DisputeVoteTallied", logs2[len(logs2)-1])
		if err != nil {
			return fmt.Errorf("failed to unpack into object: %v", err)
		}
	}
	// fmt.Println("LastTalliedVote: ", lastDispute.DisputeID)

	lastDisputeID := lastDispute.DisputeID
	if lastDisputeID == nil {
		lastDisputeID = big.NewInt(0)
	}

	for _, l := range logs1 {
		newDispute := contracts1.ZapNewDispute{}
		err := bar1.UnpackLog(&newDispute, "NewDispute", l)
		if err != nil {
			return fmt.Errorf("failed to unpack into object: %v", err)
		}
		// skip dispute events that have been tallied
		if newDispute.DisputeId.Cmp(lastDisputeID) < 1 {
			continue
		}

		array := [32]byte{}
		data := []byte("minExecutionDate")
		data = crypto.Keccak256(data)
		for i := 0; i < 32; i++ {
			array[i] = data[i]
		}
		endTime, _ := instanceM.GetDisputeUintVars(nil, newDispute.DisputeId, array)

		now := time.Now()
		epoch := now.Unix()

		// fmt.Println("End time: ", endTime)
		// fmt.Println("Epoch: ", epoch)

		if endTime.Cmp(big.NewInt(epoch)) == -1 {
			auth, _ := PrepareEthTransaction(ctx)
			_, err = instanceZ.TallyVotes(auth, newDispute.DisputeId)
			if err != nil {
				return fmt.Errorf("failed to make TallyVotes call: %v", err)
			}
			// cut filter range from most recently checked block
			b.lastCheckedBlock = newDispute.Raw.BlockNumber
		} else {
			break
		}
	}

	return nil
}

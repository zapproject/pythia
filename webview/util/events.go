package util

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/setup"
	"github.com/zapproject/pythia/token"
)

type EventLog struct {
	Timestamp uint64
	Log       interface{}
}

func GetTransferLogs() []EventLog {
	tokenABI, err := abi.JSON(strings.NewReader(token.ERC20BasicABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.TokenAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, tokenABI, nil, nil, nil)

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	transferID := tokenABI.Events["Transfer"].ID
	// fmt.Println("TransferID: ", transferID)
	transferQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   header.Number,
		// ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.TokenAddress).(common.Address),
		},
		Topics: [][]common.Hash{{transferID}},
	}

	// fmt.Println("Block: ", header.Number)

	logs, err := client.FilterLogs(setup.CTX, transferQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}
	// fmt.Println("transfer logs: ", logs)

	transfers := []EventLog{}
	for _, l := range logs {
		transfer := token.ZapTokenBSCTransfer{}
		err := bar.UnpackLog(&transfer, "Transfer", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		if transfer.From.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() ||
			transfer.To.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, transfer}
			transfers = append(transfers, log)
		}
	}

	return transfers
}

func GetApprovalLogs() []EventLog {
	tokenABI, err := abi.JSON(strings.NewReader(token.ERC20ABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.TokenAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, tokenABI, nil, nil, nil)

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	approvalID := tokenABI.Events["Approval"].ID

	approvalQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   header.Number,
		ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.TokenAddress).(common.Address),
		},
		Topics: [][]common.Hash{{approvalID}},
	}

	logs, err := client.FilterLogs(setup.CTX, approvalQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	approvals := []EventLog{}
	for _, l := range logs {
		approval := token.ZapTokenBSCApproval{}
		err := bar.UnpackLog(&approval, "Approval", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		if approval.Owner.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() ||
			approval.Spender.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, approval}
			approvals = append(approvals, log)
		}
	}

	return approvals
}

func GetNewStakeLogs() []EventLog {
	stakeABI, err := abi.JSON(strings.NewReader(contracts1.ZapStakeABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.ContractAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, stakeABI, nil, nil, nil)

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	stakeID := stakeABI.Events["NewStake"].ID

	stakeQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   header.Number,
		ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.ContractAddress).(common.Address),
		},
		Topics: [][]common.Hash{{stakeID}},
	}

	logs, err := client.FilterLogs(setup.CTX, stakeQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	stakes := []EventLog{}
	for _, l := range logs {
		stake := contracts1.ZapStakeNewStake{}
		err := bar.UnpackLog(&stake, "NewStake", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		if stake.Sender.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, stakes}
			stakes = append(stakes, log)
		}
	}

	return stakes
}

func GetStakeRequestedLogs() []EventLog {
	requestedABI, err := abi.JSON(strings.NewReader(contracts1.ZapStakeABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.ContractAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, requestedABI, nil, nil, nil)

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	requestedID := requestedABI.Events["StakeWithdrawRequested"].ID

	requestedQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   header.Number,
		ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.ContractAddress).(common.Address),
		},
		Topics: [][]common.Hash{{requestedID}},
	}

	logs, err := client.FilterLogs(setup.CTX, requestedQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	requests := []EventLog{}
	for _, l := range logs {
		request := contracts1.ZapStakeStakeWithdrawRequested{}
		err := bar.UnpackLog(&request, "StakeWithdrawRequested", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		if request.Sender.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, requests}
			requests = append(requests, log)
		}
	}

	return requests
}

func GetStakeWithdawLogs() []EventLog {
	withdrawABI, err := abi.JSON(strings.NewReader(contracts1.ZapStakeABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.ContractAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, withdrawABI, nil, nil, nil)

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	withdrawID := withdrawABI.Events["StakeWithdrawn"].ID

	// fmt.Println("Before filter query")
	withdrawQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   header.Number,
		ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.ContractAddress).(common.Address),
		},
		Topics: [][]common.Hash{{withdrawID}},
	}

	logs, err := client.FilterLogs(setup.CTX, withdrawQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	withdraws := []EventLog{}
	for _, l := range logs {
		withdraw := contracts1.ZapStakeStakeWithdrawn{}
		err := bar.UnpackLog(&withdraw, "StakeWithdrawn", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		if withdraw.Sender.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, withdraw}
			withdraws = append(withdraws, log)
		}
	}

	return withdraws
}

func GetMinedLogs() []EventLog {
	minedABI, err := abi.JSON(strings.NewReader(contracts1.ZapLibraryABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.ContractAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, minedABI, nil, nil, nil)

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	minedID := minedABI.Events["NonceSubmitted"].ID

	minedQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   header.Number,
		ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.ContractAddress).(common.Address),
		},
		Topics: [][]common.Hash{{minedID}},
	}

	logs, err := client.FilterLogs(setup.CTX, minedQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	mineds := []EventLog{}
	for _, l := range logs {
		mined := contracts1.ZapLibraryNonceSubmitted{}
		err := bar.UnpackLog(&mined, "NonceSubmitted", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		if mined.Miner.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, mined}
			mineds = append(mineds, log)
		}
	}

	return mineds
}

func GetNewDisputeLogs() []EventLog {
	disputeABI, err := abi.JSON(strings.NewReader(contracts1.ZapABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.ContractAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, disputeABI, nil, nil, nil)

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	disputeID := disputeABI.Events["NewDispute"].ID

	disputeQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   header.Number,
		ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.ContractAddress).(common.Address),
		},
		Topics: [][]common.Hash{{disputeID}},
	}

	logs, err := client.FilterLogs(setup.CTX, disputeQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	disputes := []EventLog{}
	for _, l := range logs {

		dispute := contracts1.ZapNewDispute{}
		err := bar.UnpackLog(&dispute, "NewDispute", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		if dispute.Miner.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, dispute}
			disputes = append(disputes, log)
		}
	}

	return disputes
}

func GetVotedLogs() []EventLog {
	votedABI, err := abi.JSON(strings.NewReader(contracts1.ZapDisputeABI))
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to parse abi: %v", err)
	}
	contractAddress := setup.CTX.Value(ZapCommon.ContractAddress).(common.Address)
	client := setup.CTX.Value(ZapCommon.ClientContextKey).(rpc.ETHClient)

	bar := bind.NewBoundContract(contractAddress, votedABI, nil, nil, nil)

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }

	votedID := votedABI.Events["Voted"].ID

	votedQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   header.Number,
		ToBlock: nil,
		Addresses: []common.Address{
			setup.CTX.Value(ZapCommon.ContractAddress).(common.Address),
		},
		Topics: [][]common.Hash{{votedID}},
	}

	logs, err := client.FilterLogs(setup.CTX, votedQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	voteds := []EventLog{}
	for _, l := range logs {
		voted := contracts1.ZapDisputeVoted{}
		err := bar.UnpackLog(&voted, "Voted", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}

		if voted.Voter.Hex() == setup.CTX.Value(ZapCommon.PublicAddress).(common.Address).Hex() {
			time, _ := client.HeaderByNumber(setup.CTX, big.NewInt(int64(l.BlockNumber)))
			log := EventLog{time.Time, voted}
			voteds = append(voteds, log)
		}
	}

	return voteds
}

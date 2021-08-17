package webview

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/webview/webview"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/setup"
	"github.com/zapproject/pythia/token"
)

func showWallet(w webview.WebView) {
	setup.App()

	w.SetTitle("Wallet")
	w.SetSize(800, 600, webview.HintMin)

	addr := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	instance := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	zapBalance, _ := instance.BalanceOf(nil, addr)

	w.Bind("balance", func() string {
		return zapBalance.String()
	})

	w.Bind("showTxs", func() {
		showTxs()
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/wallet.html")
	p = "file://" + p
	w.Navigate(p)
	w.Run()
}

func showTxs() {
	transferLogs := getTransferLogs()
	fmt.Println("Transfer logs: ", transferLogs)

	approvalLogs := getApprovalLogs()
	fmt.Println("Approval logs: ", approvalLogs)
}

func getTransferLogs() []token.ZapTokenBSCTransfer {
	tokenABI, err := abi.JSON(strings.NewReader(token.ERC20BasicABI))
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

	transferID := tokenABI.Events["Transfer"].ID
	// fmt.Println("TransferID: ", transferID)
	transferQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		// ToBlock:   big.NewInt(header.Number.Int64()),
		ToBlock: nil,
		// setup.CTX.Value(ZapCommon.PublicAddress).(common.Address),
		Addresses: []common.Address{setup.CTX.Value(ZapCommon.TokenAddress).(common.Address)},
		Topics:    [][]common.Hash{{transferID}},
	}

	// fmt.Println("Block: ", header.Number)

	logs, err := client.FilterLogs(setup.CTX, transferQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}
	// fmt.Println("transfer logs: ", logs)

	transfers := []token.ZapTokenBSCTransfer{}
	for _, l := range logs {
		transfer := token.ZapTokenBSCTransfer{}
		err := bar.UnpackLog(&transfer, "Transfer", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		transfers = append(transfers, transfer)
	}

	return transfers
}

func getApprovalLogs() []token.ZapTokenBSCApproval {
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
		ToBlock:   nil,
		Addresses: []common.Address{setup.CTX.Value(ZapCommon.TokenAddress).(common.Address)},
		// Addresses: []common.Address{setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)},
		Topics: [][]common.Hash{{approvalID}},
	}

	logs, err := client.FilterLogs(setup.CTX, approvalQuery)
	if err != nil {
		fmt.Errorf("\U0001F6AB failed to get nonce logs: %v", err)
	}

	fmt.Println("approce logs: ", logs)
	approvals := []token.ZapTokenBSCApproval{}
	for _, l := range logs {
		approval := token.ZapTokenBSCApproval{}
		err := bar.UnpackLog(&approval, "Approval", l)
		if err != nil {
			fmt.Errorf("\U0001F6AB failed to unpack into object: %v", err)
		}
		approvals = append(approvals, approval)
	}

	return approvals
}

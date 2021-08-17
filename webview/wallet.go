package webview

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/webview/webview"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/setup"
	"github.com/zapproject/pythia/webview/util"
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
	transferLogs := util.GetTransferLogs()
	fmt.Println("Transfer logs: ", transferLogs)

	approvalLogs := util.GetApprovalLogs()
	fmt.Println("Approval logs: ", approvalLogs)

	newStakeLogs := util.GetNewStakeLogs()
	fmt.Println("New Stake Logs: ", newStakeLogs)

	stakeWithdrawRequestLogs := util.GetStakeRequestedLogs()
	fmt.Println("Stake withdraw request logs: ", stakeWithdrawRequestLogs)

	stakeWithdrawnLogs := util.GetStakeWithdawLogs()
	fmt.Println("Stake Withdrawn logs: ", stakeWithdrawnLogs)

	minedLogs := util.GetMinedLogs()
	fmt.Println("Mined logs: ", minedLogs)

	newDisputeLogs := util.GetNewDisputeLogs()
	fmt.Println("New Dispute logs: ", newDisputeLogs)

	votedLogs := util.GetVotedLogs()
	fmt.Println("Voted logs: ", votedLogs)
}

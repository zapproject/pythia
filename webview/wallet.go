package webview

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/webview/webview"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/setup"
	"github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/webview/util"
)

type transaction struct {
	to        string
	from      string
	event     string
	timestamp uint64
}

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
	eventLogs := util.GetTransferLogs()

	approvalLogs := util.GetApprovalLogs()
	// fmt.Println("Approval logs: ", approvalLogs)
	eventLogs = append(eventLogs, approvalLogs...)

	newStakeLogs := util.GetNewStakeLogs()
	// fmt.Println("New Stake Logs: ", newStakeLogs)
	eventLogs = append(eventLogs, newStakeLogs...)

	stakeWithdrawRequestLogs := util.GetStakeRequestedLogs()
	// fmt.Println("Stake withdraw request logs: ", stakeWithdrawRequestLogs)
	eventLogs = append(eventLogs, stakeWithdrawRequestLogs...)

	stakeWithdrawnLogs := util.GetStakeWithdawLogs()
	// fmt.Println("Stake Withdrawn logs: ", stakeWithdrawnLogs)
	eventLogs = append(eventLogs, stakeWithdrawnLogs...)

	minedLogs := util.GetMinedLogs()
	// fmt.Println("Mined logs: ", minedLogs)
	eventLogs = append(eventLogs, minedLogs...)

	newDisputeLogs := util.GetNewDisputeLogs()
	// fmt.Println("New Dispute logs: ", newDisputeLogs)
	eventLogs = append(eventLogs, newDisputeLogs...)

	votedLogs := util.GetVotedLogs()
	// fmt.Println("Voted logs: ", votedLogs)
	eventLogs = append(eventLogs, votedLogs...)

	sort.Slice(eventLogs, func(i, j int) bool {
		return eventLogs[i].Timestamp < eventLogs[j].Timestamp
	})

	txs := formatLogs(eventLogs)
	fmt.Println(txs)
}

func formatLogs(eventLogs []util.EventLog) []transaction {
	txs := []transaction{}

	for _, j := range eventLogs {
		tx := transaction{}
		tx.timestamp = j.Timestamp

		logType := fmt.Sprintf("%T", j.Log)
		switch logType {
		case "token.ZapTokenBSCTransfer":
			tx.from = j.Log.(token.ZapTokenBSCTransfer).From.String()
			tx.to = j.Log.(token.ZapTokenBSCTransfer).To.String()
			tx.event = "Transfer"

		case "token.ZapTokenBSCApproval":
			tx.from = j.Log.(token.ZapTokenBSCApproval).Owner.String()
			tx.to = j.Log.(token.ZapTokenBSCApproval).Spender.String()
			tx.event = "Approval"

		case "contracts1.ZapStakeNewStake":
			tx.event = "NewStake"

		case "contracts1.ZapStakeStakeWithdrawRequested":
			tx.event = "StakeWithdrawRequested"

		case "contracts1.ZapStakeStakeWithdrawn":
			tx.event = "StakeWithdrawn"

		case "contracts1.ZapLibraryNonceSubmitted":
			tx.event = "NonceSubmitted"

		case "contracts1.ZapNewDispute":
			tx.event = "NewDispute"

		case "contracts1.ZapDisputeVoted":
			tx.event = "Voted"

		}

		txs = append(txs, tx)
	}

	return txs
}

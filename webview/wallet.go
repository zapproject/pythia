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

type Transaction struct {
	To        string `json:"To"`
	From      string `json:"From"`
	Event     string `json:"Event"`
	Timestamp uint64 `json:"Timestamp"`
}

func showWallet(w webview.WebView) {
	// w.SetTitle("Wallet")
	// w.SetSize(800, 600, webview.HintMin)

	addr := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	instance := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	zapBalance, _ := instance.BalanceOf(nil, addr)

	w.Bind("showConfig", func() {
		showConfig(w)
	})

	w.Bind("showStake", func() {
		showStake(w)
	})

	w.Bind("balance", func() string {
		return zapBalance.String()
	})

	w.Bind("showTxs", func() []Transaction {
		return showTxs()
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/wallet.html")
	p = "file://" + p
	w.Navigate(p)
	// w.Run()
}

func showTxs() []Transaction {
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
	// fmt.Println(txs)
	return txs
}

func formatLogs(eventLogs []util.EventLog) []Transaction {
	txs := []Transaction{}
	// fmt.Println(eventLogs)
	for _, j := range eventLogs {
		tx := Transaction{}
		tx.Timestamp = j.Timestamp

		logType := fmt.Sprintf("%T", j.Log)

		switch logType {
		case "token.ZapTokenBSCTransfer":
			tx.From = j.Log.(token.ZapTokenBSCTransfer).From.String()
			tx.To = j.Log.(token.ZapTokenBSCTransfer).To.String()
			tx.Event = "Transfer"

		case "token.ZapTokenBSCApproval":
			tx.From = j.Log.(token.ZapTokenBSCApproval).Owner.String()
			tx.To = j.Log.(token.ZapTokenBSCApproval).Spender.String()
			tx.Event = "Approval"

		case "contracts1.ZapStakeNewStake":
			tx.Event = "NewStake"

		case "contracts1.ZapStakeStakeWithdrawRequested":
			tx.Event = "StakeWithdrawRequested"

		case "contracts1.ZapStakeStakeWithdrawn":
			tx.Event = "StakeWithdrawn"

		case "contracts1.ZapLibraryNonceSubmitted":
			tx.Event = "NonceSubmitted"

		case "contracts1.ZapNewDispute":
			tx.Event = "NewDispute"

		case "contracts1.ZapDisputeVoted":
			tx.Event = "Voted"

		default:
			tx.Event = logType
		}

		txs = append(txs, tx)
	}

	return txs
}

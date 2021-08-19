package webview

import (
	"os"
	"path/filepath"

	"github.com/webview/webview"
	"github.com/zapproject/pythia/ops"
	"github.com/zapproject/pythia/setup"
)

func showStake(w webview.WebView) {
	w.Bind("newStake", func() string {
		err := newStake()

		if err != nil {
			return err.Error()
		} else {
			return "Successfully desposited stake."
		}
	})

	w.Bind("stakeWithdrawRequest", func() string {
		err := stakeWithdrawRequest()

		if err != nil {
			return err.Error()
		} else {
			return "Successfully requested for a stake withdraw, you will be able to withdraw after 7 days"
		}
	})

	w.Bind("stakeWithdraw", func() string {
		err := stakeWithdraw()

		if err != nil {
			return err.Error()
		} else {
			return "Successfully withdrew stake."
		}
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/stake.html")
	p = "file://" + p
	w.Navigate(p)
}

func newStake() error {
	return ops.Deposit(setup.CTX)
}

func stakeWithdrawRequest() error {
	return ops.RequestStakingWithdraw(setup.CTX)
}

func stakeWithdraw() error {
	return ops.WithdrawStake(setup.CTX)
}

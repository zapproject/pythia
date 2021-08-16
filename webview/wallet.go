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

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/wallet.html")
	p = "file://" + p
	fmt.Println(p)
	w.Navigate(p)
	w.Run()
}

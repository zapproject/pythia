package main

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

	// debug := true
	// w := webview.New(debug)

	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintMin)

	addr := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	instance := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	zapBalance, _ := instance.BalanceOf(nil, addr)
	fmt.Println(zapBalance)

	w.Bind("balance", func() string {
		return zapBalance.String()
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "public/wallet.html")
	p = "file://" + p
	w.Navigate(p)
	w.Run()
}

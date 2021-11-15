package webview

import (
	"os"
	"path/filepath"

	"github.com/webview/webview"
)

func LoadWebview() {

	debug := false
	w := webview.New(debug)

	defer w.Destroy()
	w.SetTitle("Pythia")
	w.SetSize(800, 600, webview.HintMin)

	w.Bind("showWallet", func() {
		showWallet(w)
	})

	w.Bind("getPubKey", getPubKey)
	w.Bind("configWallet", configWallet)
	// w.Bind("loadHtml", loadHtml)
	w.Bind("getBalance", getBalance)
	w.Bind("stakeStatus", stakeStatus)
	w.Bind("startMine", startMine)
	w.Bind("stopMine", stopMine)
	w.Bind("isMining", isMining)
	w.Bind("isInit", isInit)
	w.Bind("showLogs", showLogs)
	w.Bind("requestWithdraw", requestWithdraw)
	w.Bind("withdraw", withdraw)
	w.Bind("canWithdraw", canWithdraw)
	w.Bind("withdrawMsg", withdrawMsg)
	w.Bind("depositStake", depositStake)
	w.Bind("transfer", transfer)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/index.html")
	p = "file://" + p

	w.Navigate(p)

	w.Run()
}

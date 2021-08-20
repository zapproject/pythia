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

	w.Bind("configWallet",configWallet)
	w.Bind("loadHtml",loadHtml)
	w.Bind("getBalance",getBalance)
	w.Bind("stakeStatus",stakeStatus)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/index.html")
	p = "file://" + p

	w.Navigate(p)

	// content, _ := ioutil.ReadFile("./public/wallet.html")
	// w.Navigate(string(content))

	// w.Navigate("./public/wallet.html")
	w.Run()
}

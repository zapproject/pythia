package webview

import (
	"os"
	"path/filepath"

	"github.com/webview/webview"
)

func Start() {
	debug := true
	w := webview.New(debug)

	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintMin)

	w.Bind("showWallet", func() {
		showWallet(w)
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/config.html")
	p = "file://" + p

	w.Navigate(p)
	w.Run()
}

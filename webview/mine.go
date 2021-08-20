package webview

import (
	"os"
	"path/filepath"

	"github.com/webview/webview"
)

func showMine(w webview.WebView) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/mine.html")
	p = "file://" + p
	w.Navigate(p)
}

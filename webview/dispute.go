package webview

import (
	"os"
	"path/filepath"

	"github.com/webview/webview"
)

func showDispute(w webview.WebView) {
	w.Bind("vote", func(selection bool) string {
		return vote(selection)
	})

	w.Bind("newDispute", func(requestID int, timestamp int) bool {
		return newDispute(requestID, timestamp)
	})

	w.Bind("showStatus", func() string {
		return showStatus()
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/dispute.html")
	p = "file://" + p
	w.Navigate(p)
}

func vote(selection bool) string {

	return "vote statistics"
}

func newDispute(requestID int, timestamp int) bool {

	return true
}

func showStatus() string {

	return "status"
}

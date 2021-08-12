package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/webview/webview"
)

func main() {
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
	p := filepath.Join(filepath.Dir(ex), "public/index.html")
	p = "file://" + p
	fmt.Println(p)
	w.Navigate(p)

	// content, _ := ioutil.ReadFile("./public/wallet.html")
	// w.Navigate(string(content))

	// w.Navigate("./public/wallet.html")
	w.Run()
<<<<<<< HEAD
}
=======
}
>>>>>>> cc92b2bc326ef6ebff96ed74a173db300734cb3f

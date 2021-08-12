package webview

import ( 
	"os"
	"github.com/webview/webview"
	"fmt"
	"path/filepath"

)

func LoadWebview(){


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
	p := filepath.Join(filepath.Dir(ex), "webview/public/index.html")
	p = "file://" + p
	fmt.Println(p)
	w.Navigate(p)

	// content, _ := ioutil.ReadFile("./public/wallet.html")
	// w.Navigate(string(content))

	// w.Navigate("./public/wallet.html")
	w.Run()
}

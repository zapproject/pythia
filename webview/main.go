package main 

import(
	"github.com/webview/webview"
	// "io/ioutil"
	"path/filepath"
	"os"
)

func main(){
	debug := true
	w := webview.New(debug)

	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintMin)


	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex),"public/index.html")

	w.Navigate(p)
	w.Run()
}


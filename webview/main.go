package main 

import(
	"github.com/webview/webview"
	"io/ioutil"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main(){
	debug := true
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go func() {
		// Set up your http server here
	   http.HandleFunc("/", serveFiles)
	   log.Fatal(http.Serve(ln, nil))
   }()
	index, err := ioutil.ReadFile("./public/index.html")
	if(err != nil){
		fmt.Println("Alpine not found")
	}
	// alpine, err := ioutil.ReadFile("./public/alpinejs/cdn.js")
	// if(err != nil){
	// 	fmt.Println("Alpine not found")
	// }
	
	w := webview.New(debug)

	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintMin)
	fmt.Println(string(index))

	// w.Init(string(alpine))
	w.Navigate("http://"+ln.Addr().String())
	w.Run()
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := "." + r.URL.Path
    if p == "./" {
        p = "./public/"
    }
    http.ServeFile(w, r, p)
}
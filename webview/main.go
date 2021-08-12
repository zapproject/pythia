package main

<<<<<<< HEAD
import(
	"github.com/webview/webview"
	// "io/ioutil"
	"path/filepath"
	"io"
	// "log"
	"net/rpc/jsonrpc"
	"os"
	"fmt"
)

// type Counter struct {
// 	Value int `json:"value"`
// }

// // Add increases the value of a counter by n
// func (c *Counter) Add(n int) int {
// 	c.Value = c.Value + int(n)
// 	fmt.Println(c.Value)
// 	return c.Value
// }

// // Reset sets the value of a counter back to zero
// func (c *Counter) Reset() {
// 	c.Value = 0
// }

type ReadWriteCloser struct {
	io.ReadCloser
	io.WriteCloser
}

func (rw *ReadWriteCloser) Close() error {
	rw.ReadCloser.Close()
	rw.WriteCloser.Close()
	return nil
}

// Defined for the RPC package
type Adder int

func main(){
=======
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/webview/webview"
)

func main() {
>>>>>>> cc92b2bc326ef6ebff96ed74a173db300734cb3f
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
<<<<<<< HEAD
	p := filepath.Join(filepath.Dir(ex),"public/index.html")

	adder := Adder(1)
	var ret int

	rwc := new(ReadWriteCloser)

	rwc.ReadCloser = os.Stdin
	rwc.WriteCloser = os.Stdout

	client := jsonrpc.NewClient(rwc)

	w.Bind("add", func(a, b int) int {
		fmt.Println("Adding",a,b)
		client.Call("Adder.Add", &adder, &ret)
		return a + b
	})

=======
	p := filepath.Join(filepath.Dir(ex), "public/index.html")
	p = "file://" + p
	fmt.Println(p)
>>>>>>> cc92b2bc326ef6ebff96ed74a173db300734cb3f
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

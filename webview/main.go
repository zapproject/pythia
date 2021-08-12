package main 

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

	w.Navigate(p)
	w.Run()
}
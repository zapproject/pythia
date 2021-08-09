package main

import (

	"github.com/zapproject/pythia/qt"



)



func main() {
	
	// var err error
	// adder := Adder(1)
	// var ret int

	// rwc := new(ReadWriteCloser)

	// rwc.ReadCloser = os.Stdin
	// rwc.WriteCloser = os.Stdout

	// client := jsonrpc.NewClient(rwc)

	// for i := 0; i < 3; i++ {
		// err = client.Call("Adder.Add", &adder, &ret)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	// }

	app := qt.App()

	app.Exec()
}

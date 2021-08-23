package webview

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/webview/webview"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/ops"
	"github.com/zapproject/pythia/setup"
)

var out chan string
var wg *sync.WaitGroup
var reader *os.File

func showMine(w webview.WebView) {

	w.Bind("startMine", func() {
		startMine()
	})

	w.Bind("showLogs", func() string {
		return showLogs()
	})

	// w.Bind("stopMine", func() {
	// 	stopMine()
	// })

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/mine.html")
	p = "file://" + p
	w.Navigate(p)
}

func startMine() {
	var writer *os.File
	var err error
	reader, writer, err = os.Pipe()
	if err != nil {
		panic(err)
	}

	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	// out = make(chan string)

	// wg = new(sync.WaitGroup)
	// wg.Add(1)
	// go func() {
	// 	var buf bytes.Buffer
	// 	wg.Done()
	// 	io.Copy(&buf, reader)
	// 	out <- buf.String()
	// }()
	// wg.Wait()

	//create os kill sig listener
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	exitChannels := make([]*chan os.Signal, 0)

	cfg := config.GetConfig()
	var ds *ops.DataServerOps
	if !cfg.EnablePoolWorker {
		setup.AddDBToCtx(false)
		ch := make(chan os.Signal)
		exitChannels = append(exitChannels, &ch)

		var err error
		ds, err = ops.CreateDataServerOps(setup.CTX, ch)
		if err != nil {
			log.Fatal(err)
		}
		//start and wait for it to be ready
		ds.Start(setup.CTX)
		<-ds.Ready()
	}
	//start miner
	ch := make(chan os.Signal)
	exitChannels = append(exitChannels, &ch)
	miner, err := ops.CreateMiningManager(setup.CTX, ch, ops.NewSubmitter())
	if err != nil {
		log.Fatal(err)
	}
	miner.Start(setup.CTX)

	//now we wait for kill sig
	<-c
	//and then notify exit channels
	for _, ch := range exitChannels {
		*ch <- os.Interrupt
	}
	cnt := 0
	start := time.Now()
	for {
		cnt++
		dsStopped := false
		minerStopped := false

		if ds != nil {
			dsStopped = !ds.Running
		} else {
			dsStopped = true
		}

		if miner != nil {
			minerStopped = !miner.Running
		} else {
			minerStopped = true
		}

		if !dsStopped && !minerStopped && cnt > 60 {
			fmt.Printf("\U000026A0 Taking longer than expected to stop operations. Waited %v so far\n", time.Now().Sub(start))
		} else if dsStopped && minerStopped {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("Main shutdown complete \U00002622\n")

}

func stopMine() {
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

func showLogs() string {
	// wg = new(sync.WaitGroup)
	// wg.Add(1)
	// go func() {
	var buf bytes.Buffer
	// wg.Done()
	io.Copy(&buf, reader)
	// out <- buf.String()
	// }()
	// wg.Wait()
	return buf.String()
}

package webview

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/webview/webview"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/ops"
	"github.com/zapproject/pythia/setup"
	"github.com/zapproject/pythia/util"
)

var stdout *os.File
var stderr *os.File

func showMine(w webview.WebView) {

	w.Bind("startMine", func() {
		go startMine()
	})

	w.Bind("showLogs", func(index int) string {
		return showLogs(index)
	})

	w.Bind("stopMine", func() {
		go stopMine()
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/mine.html")
	p = "file://" + p
	w.Navigate(p)
}

func startMine() {
	file, _ := os.OpenFile("./webview/public/miningLogs.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	stdout = os.Stdout
	stderr = os.Stderr
	os.Stdout = file
	os.Stderr = file
	log.SetOutput(file)
	loggingCfgs := util.GetLoggingConfig()
	util.InitLoggers(loggingCfgs)

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
	e := os.Remove("./webview/public/miningLogs.log")
	if e != nil {
		panic(e)
	}

	// return to original outputs
	os.Stderr = stderr
	os.Stdout = stdout
	log.SetOutput(stdout)
	loggingCfgs := util.GetLoggingConfig()
	util.InitLoggers(loggingCfgs)

	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

func showLogs(index int) string {
	var content string
	file, _ := os.Open("./webview/public/miningLogs.log")
	reader := bufio.NewScanner(file)

	reader.Split(bufio.ScanLines)

	ind := 0
	for reader.Scan() {
		ind = ind + 1

		if ind <= index {
			continue
		}

		content = content + reader.Text() + "\n"
	}

	return content
}

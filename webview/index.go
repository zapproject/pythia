package webview

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/webview/webview"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/setup"
)

func Start() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintMin)

	showConfig(w)
	w.Run()
}

func showConfig(w webview.WebView) {

	w.Bind("showWallet", func() {
		showWallet(w)
	})

	w.Bind("showStake", func() {
		showStake(w)
	})

	w.Bind("showMine", func() {
		showMine(w)
	})

	w.Bind("saveConfigs", func(inputs map[string]interface{}) {
		config.NewConfig()
		config.SetPublicAddress(fmt.Sprintf("%v", inputs["publickey"]))
		config.SetPrivateKey(fmt.Sprintf("%v", inputs["privatekey"]))
		trackers := []string{}

		if inputs["balance"] == true {
			trackers = append(trackers, "balance")
		}

		if inputs["disputeStatus"] == true {
			trackers = append(trackers, "disputeStatus")
		}

		if inputs["gas"] == true {
			trackers = append(trackers, "gas")
		}

		if inputs["tokenBalance"] == true {
			trackers = append(trackers, "tokenBalance")
		}

		if inputs["indexers"] == true {
			trackers = append(trackers, "indexers")
		}

		if inputs["currentVariables"] == true {
			trackers = append(trackers, "currentVariables")
		}

		if inputs["newCurrentVariables"] == true {
			trackers = append(trackers, "newCurrentVariables")
		}

		if inputs["disputeChecker"] == true {
			trackers = append(trackers, "disputeChecker")
		}

		if inputs["tallyVotes"] == true {
			trackers = append(trackers, "tallyVotes")
		}

		config.SetTrackers(trackers)

		trackerCycle := strings.Trim(fmt.Sprintf("%v", inputs["trackerCycle"]), "s")
		converted, _ := strconv.ParseUint(trackerCycle, 10, 32)
		config.SetTrackerSleepCycle(uint(converted))

		converted, _ = strconv.ParseUint(fmt.Sprintf("%v", inputs["clientTimeout"]), 10, 32)
		config.SetEthClientTimeout(uint(converted))

		// other static configs
		config.SetContractAddress("0x7a2088a1bFc9d81c55368AE168C2C02570cB814F")
		config.SetTokenAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
		config.SetVaultAddress("0x09635f643e140090a9a8dcd712ed6285858cebef")
		config.SetNodeURL("http://127.0.0.1:8545/")
		config.SetServerHost("localhost")
		config.SetServerPort(5001)
		config.SetGasMultiplier(1)
		config.SetGasMax(30)
		config.SetUseGPU(false)
		config.SetIndexFolder("indexes")
		config.SetDBFile("zapDB")
		config.SetDisputeTimeDelta(600)

		setup.App()
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/index.html")
	p = "file://" + p

	w.Navigate(p)
}

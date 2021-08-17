package webview

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/webview/webview"
	"github.com/zapproject/pythia/config"
)

func Start() {
	debug := true
	w := webview.New(debug)

	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintMin)

	w.Bind("showWallet", func() {
		showWallet(w)
	})

	// w.Bind("saveConfigs", func(inputs interface{}) {
	w.Bind("saveConfigs", func(inputs map[string]interface{}) {
		// fmt.Println(fmt.Sprintf("inputs: %T", inputs["publickey"]))
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

		trackerCycle := strings.Trim(fmt.Sprintf("%v", inputs["trackerCycle"]), "s")
		converted, _ := strconv.ParseUint(trackerCycle, 10, 32)
		config.SetTrackerSleepCycle(uint(converted))

		converted, _ = strconv.ParseUint(fmt.Sprintf("%v", inputs["clientTimeout"]), 10, 32)
		config.SetEthClientTimeout(uint(converted))

		// other static configs
		config.SetContractAddress("0xCD8a1C3ba11CF5ECfa6267617243239504a98d90")
		config.SetTokenAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
		config.SetVaultAddress("0x82e01223d51Eb87e16A03E24687EDF0F294da6f1")
		config.SetNodeURL("http://127.0.0.1:8545/")
		config.SetServerHost("localhost")
		config.SetServerPort(5001)
		config.SetGasMultiplier(1)
		config.SetGasMax(30)
		config.SetUseGPU(false)
		config.SetIndexFolder("indexes")
		config.SetDBFile("zapDB")
		config.SetDisputeTimeDelta(600)

		cfg := config.GetConfig()
		fmt.Println("CFG: ", cfg.PublicAddress)
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Join(filepath.Dir(ex), "webview/public/index.html")
	p = "file://" + p

	w.Navigate(p)
	w.Run()
}

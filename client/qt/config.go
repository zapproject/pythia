package qt

import (
	"strconv"
	"time"
	"fmt"
	"log"
	"github.com/therecipe/qt/widgets"
	"github.com/zapproject/pythia/config"
)

var trackers map[string]int

func showConfig() {
	config.NewConfig()
	trackers = map[string]int{}

	configWidget := widgets.NewQGroupBox2("Configuration", nil)
	layout := widgets.NewQGridLayout2()

	/**
	Login Group Box
	*/
	LoginRadioGroup := widgets.NewQGroupBox2("Choose Login Type", nil)
	LoginButtonGroup := widgets.NewQButtonGroup(nil)
	LoginTypeNames := [3]string{"Public/Private Key Pair", "Key File", "Mneomonic Phrase"}
	LoginRadioGrid := widgets.NewQGridLayout2()

	for i, name := range LoginTypeNames {
		button := widgets.NewQRadioButton2(name, nil)
		LoginButtonGroup.AddButton(button, i)
		LoginRadioGrid.AddWidget2(button, i, 0, 0)
	}
	LoginRadioGroup.SetLayout(LoginRadioGrid)

	var PrivatePublicGroup *widgets.QGroupBox = nil
	var KeyFileGroup *widgets.QGroupBox = nil
	var MnemonicGroup *widgets.QGroupBox = nil

	LoginButtonGroup.ConnectButtonToggled2(func(id int, checked bool) {

		if id == 0 && checked {

			if MnemonicGroup != nil {
				LoginRadioGrid.RemoveWidget(MnemonicGroup)
				MnemonicGroup.DeleteLater()
			}

			if KeyFileGroup != nil {
				LoginRadioGrid.RemoveWidget(KeyFileGroup)
				KeyFileGroup.DeleteLater()
			}

			PrivatePublicGroup = newPrivatePublicKeyWidget()

			LoginRadioGrid.AddWidget2(PrivatePublicGroup, id, 1, 0)
		}

		if id == 1 && checked {

			if PrivatePublicGroup != nil {
				LoginRadioGrid.RemoveWidget(PrivatePublicGroup)
				PrivatePublicGroup.DeleteLater()
			}

			if MnemonicGroup != nil {
				LoginRadioGrid.RemoveWidget(MnemonicGroup)
				MnemonicGroup.DeleteLater()
			}

			KeyFileGroup = newKeyFileWidget()

			LoginRadioGrid.AddWidget2(KeyFileGroup, id, 1, 0)

		}

		if id == 2 && checked {

			if PrivatePublicGroup != nil {
				LoginRadioGrid.RemoveWidget(PrivatePublicGroup)
				PrivatePublicGroup.DeleteLater()
			}

			if KeyFileGroup != nil {
				LoginRadioGrid.RemoveWidget(KeyFileGroup)
				KeyFileGroup.DeleteLater()
			}

			MnemonicGroup = newMnemonicWidget()

			LoginRadioGrid.AddWidget2(MnemonicGroup, id, 1, 0)

		}

		// window.Show()

	})

	/**
	ENV Config Group Box
	*/
	envWidget := widgets.NewQGroupBox2("App Enviornment Variables", nil)
	nodeURLLabel := widgets.NewQLabel2("Node URL", nil, 0)
	nodeURLBox := widgets.NewQComboBox(nil)
	clientTimeoutLabel := widgets.NewQLabel2("Client Timeout", nil, 0)
	clientTimeoutBox := widgets.NewQLineEdit(nil)
	trackerCycleLabel := widgets.NewQLabel2("Tracker Cycle", nil, 0)
	trackerCycleBox := widgets.NewQComboBox(nil)

	// add on editing finished signals to update configs
	nodeURLBox.ConnectCurrentIndexChanged(func(index int) { nodeURLChanged(index) })
	clientTimeoutBox.ConnectEditingFinished(func() { clientTimeoutChanged(clientTimeoutBox) })
	trackerCycleBox.ConnectCurrentIndexChanged(func(index int) { trackerCycleChanged(index) })

	nodeURLBox.AddItems([]string{"Binance Mainnet", "Binance Testnet", "Local"})
	trackerCycleBox.AddItems([]string{"10s", "30s", "90s", "120s"})
	envLayout := widgets.NewQGridLayout2()
	envLayout.AddWidget2(nodeURLLabel, 1, 0, 0)
	envLayout.AddWidget2(nodeURLBox, 1, 1, 0)
	envLayout.AddWidget2(clientTimeoutLabel, 2, 0, 0)
	envLayout.AddWidget2(clientTimeoutBox, 2, 1, 0)
	envLayout.AddWidget2(trackerCycleLabel, 3, 0, 0)
	envLayout.AddWidget2(trackerCycleBox, 3, 1, 0)

	envWidget.SetLayout(envLayout)

	/**
	Tracker Configs
	*/
	trackerWidget := widgets.NewQGroupBox2("Trackers", nil)
	balanceLabel := widgets.NewQLabel2("Balance", nil, 0)
	balanceBox := widgets.NewQCheckBox(nil)
	disputeStatusLabel := widgets.NewQLabel2("Dispute Status", nil, 0)
	disputeStatusBox := widgets.NewQCheckBox(nil)
	gasLabel := widgets.NewQLabel2("Gas", nil, 0)
	gasBox := widgets.NewQCheckBox(nil)
	tokenBalanceLabel := widgets.NewQLabel2("Token Balance", nil, 0)
	tokenBalanceBox := widgets.NewQCheckBox(nil)
	indexersLabel := widgets.NewQLabel2("Indexers", nil, 0)
	indexersBox := widgets.NewQCheckBox(nil)
	newCurVarsLabel := widgets.NewQLabel2("Current Variables", nil, 0)
	newCurVarsBox := widgets.NewQCheckBox(nil)
	curVarsLabel := widgets.NewQLabel2("Current Variables", nil, 0)
	curVarsBox := widgets.NewQCheckBox(nil)
	disputeCheckerLabel := widgets.NewQLabel2("Dispute Checker", nil, 0)
	disputeCheckerBox := widgets.NewQCheckBox(nil)
	tallyVotesLabel := widgets.NewQLabel2("Tally Votes", nil, 0)
	tallyVotesBox := widgets.NewQCheckBox(nil)

	// tracker signal & slot
	balanceBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "balance") })
	disputeStatusBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "disputeStatus") })
	gasBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "gas") })
	tokenBalanceBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "tokenBalance") })
	indexersBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "indexers") })
	newCurVarsBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "newCurrentVariables") })
	curVarsBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "currentVariables") })
	disputeCheckerBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "disputeChecker") })
	tallyVotesBox.ConnectStateChanged(func(state int) { trackersBoxChanged(state, "tallyVotes") })

	trackerLayout := widgets.NewQGridLayout2()
	trackerLayout.AddWidget2(balanceLabel, 0, 0, 0)
	trackerLayout.AddWidget2(balanceBox, 0, 1, 0)
	trackerLayout.AddWidget2(indexersLabel, 0, 2, 0)
	trackerLayout.AddWidget2(indexersBox, 0, 3, 0)
	trackerLayout.AddWidget2(disputeStatusLabel, 1, 0, 0)
	trackerLayout.AddWidget2(disputeStatusBox, 1, 1, 0)
	trackerLayout.AddWidget2(newCurVarsLabel, 1, 2, 0)
	trackerLayout.AddWidget2(newCurVarsBox, 1, 3, 0)
	trackerLayout.AddWidget2(gasLabel, 2, 0, 0)
	trackerLayout.AddWidget2(gasBox, 2, 1, 0)
	trackerLayout.AddWidget2(curVarsLabel, 2, 2, 0)
	trackerLayout.AddWidget2(curVarsBox, 2, 3, 0)
	trackerLayout.AddWidget2(tokenBalanceLabel, 3, 0, 0)
	trackerLayout.AddWidget2(tokenBalanceBox, 3, 1, 0)
	trackerLayout.AddWidget2(disputeCheckerLabel, 3, 2, 0)
	trackerLayout.AddWidget2(disputeCheckerBox, 3, 3, 0)
	trackerLayout.AddWidget2(tallyVotesLabel, 4, 0, 0)
	trackerLayout.AddWidget2(tallyVotesBox, 4, 1, 0)

	trackerWidget.SetLayout(trackerLayout)

	//Start Mining
	StartButton := widgets.NewQPushButton2("Start Mining", nil)

	//Setup Status Bar
	StatusBar := widgets.NewQStatusBar(nil)
	StatusBar.ShowMessage("Fill out configuration to begin mining", 10000)

	StartButton.ConnectClicked(func(checked bool) {
		fmt.Println("Addding")
		adder := Adder(1)
		var ret int
		err := client.Call("Adder.Add", &adder, &ret)
		if err != nil {
			log.Fatal(err)
		}
		// setupConfig()
		// if configStatus() == false {
		// 	StatusBar.ShowMessage("Config is invalid", 5000)

		// } else {
		// 	StatusBar.ShowMessage("Starting Mining", 5000)
		// }
	})

	/**
	Add grid box layouts to window
	*/
	layout.AddWidget2(LoginRadioGroup, 0, 0, 0)
	layout.AddWidget2(envWidget, 1, 0, 0)
	layout.AddWidget2(trackerWidget, 2, 0, 0)
	layout.AddWidget2(StartButton, 3, 0, 0)
	layout.AddWidget3(StatusBar, 4, 0, 1, 0, 0)
	configWidget.SetLayout(layout)

	widget = configWidget
}

func newPrivatePublicKeyWidget() *widgets.QGroupBox {
	PrivatePublicLayout := widgets.NewQVBoxLayout2(nil)
	PrivatePublicGroup := widgets.NewQGroupBox2("Enter Private Public Keys", nil)
	PublicKeyLabel := widgets.NewQLabel2("Public Key", nil, 0)
	PublicText := widgets.NewQLineEdit(nil)
	PrivateKeyLabel := widgets.NewQLabel2("Private Key", nil, 0)
	PrivateText := widgets.NewQLineEdit(nil)

	PublicText.ConnectEditingFinished(func() { changeStringField(PublicText, "public_key") })
	PrivateText.ConnectEditingFinished(func() { changeStringField(PrivateText, "private_key") })

	PrivatePublicLayout.AddWidget(PublicKeyLabel, 0, 0)
	PrivatePublicLayout.AddWidget(PublicText, 0, 0)
	PrivatePublicLayout.AddWidget(PrivateKeyLabel, 0, 0)
	PrivatePublicLayout.AddWidget(PrivateText, 0, 0)

	PrivatePublicGroup.SetLayout(PrivatePublicLayout)

	return PrivatePublicGroup
}

func newKeyFileWidget() *widgets.QGroupBox {
	KeyFileLayout := widgets.NewQHBoxLayout2(nil)
	KeyFileGroup := widgets.NewQGroupBox2("Enter KeyFile Phrase", nil)

	KeyFileLabel := widgets.NewQLabel2("KeyFile Directory", nil, 0)
	KeyFileText := widgets.NewQLineEdit(nil)
	KeyFileButton := widgets.NewQPushButton2("Browse File", nil)

	fileDialog := widgets.NewQFileDialog2(nil, "Open File...", "", "")

	KeyFileButton.ConnectClicked(func(checked bool) {
		fileDialog.SetAcceptMode(widgets.QFileDialog__AcceptOpen)
		fileDialog.SetFileMode(widgets.QFileDialog__ExistingFile)
		if fileDialog.Exec() != int(widgets.QDialog__Accepted) {
			return
		}
		fn := fileDialog.SelectedFiles()[0]

		KeyFileText.SetText(fn)

	})

	KeyFileLayout.AddWidget(KeyFileLabel, 0, 0)
	KeyFileLayout.AddWidget(KeyFileText, 0, 0)
	KeyFileLayout.AddWidget(KeyFileButton, 0, 0)

	KeyFileGroup.SetLayout(KeyFileLayout)

	return KeyFileGroup
}

func newMnemonicWidget() *widgets.QGroupBox {
	MnemonicLayout := widgets.NewQVBoxLayout2(nil)
	MnemonicGroup := widgets.NewQGroupBox2("Enter Mnemonic Phrase", nil)
	MnemonicLabel := widgets.NewQLabel2("Mnemonic Phrase", nil, 0)
	MnemonicText := widgets.NewQLineEdit(nil)

	MnemonicLayout.AddWidget(MnemonicLabel, 0, 0)
	MnemonicLayout.AddWidget(MnemonicText, 0, 0)

	MnemonicGroup.SetLayout(MnemonicLayout)

	return MnemonicGroup
}

func clientTimeoutChanged(line *widgets.QLineEdit) {
	value, err := strconv.ParseUint(line.Text(), 10, 0)
	if err != nil {

	}

	config.SetEthClientTimeout(uint(value))
	cfg = config.GetConfig()
}

//Used to change config field if it's a string type
func changeStringField(line *widgets.QLineEdit, field string) {
	value := line.Text()

	switch field {
	case "public_key":
		{

			config.SetPublicAddress(value)

			whitelist := []string{value}
			config.SetServerWhiteList(whitelist)

		}
	case "private_key":
		{
			config.SetPrivateKey(value)
		}
	}
	cfg = config.GetConfig()
}

func trackerCycleChanged(index int) {
	switch index {
	case 0:
		{
			config.SetTrackerSleepCycle(uint(10))
		}
	case 1:
		{
			config.SetTrackerSleepCycle(uint(30))
		}
	case 2:
		{
			config.SetTrackerSleepCycle(uint(90))
		}
	case 3:
		{
			config.SetTrackerSleepCycle(uint(120))
		}
	}

	cfg = config.GetConfig()
}

func nodeURLChanged(index int) {
	switch index {
	case 0:
		{
			config.SetNodeURL("https://bsc-dataseed1.ninicoin.io")
		}
	case 1:
		{
			config.SetNodeURL("https://data-seed-prebsc-1-s1.binance.org:8545/")
		}
	case 2:
		{
			config.SetTokenAddress("0x09d8AF358636D9BCC9a3e177B66EB30381a4b1a8")
			config.SetVaultAddress("0x88f2bF033d43DFaF72f1660DaC4625d39C2828EB")
			config.SetContractAddress("0xA2C32d373D6d4d5572CaC38A9fF3CAa20a29eB05")
			config.SetNodeURL("http://127.0.0.1:8545/")
		}
	}

	cfg = config.GetConfig()
}

func trackersBoxChanged(state int, tracker string) {
	switch state {
	case 0:
		{
			trackers["balance"] = 0
			tracks := buildTracker()
			config.SetTrackers(tracks)
		}
	case 2:
		{
			trackers["balance"] = 2
			tracks := buildTracker()
			config.SetTrackers(tracks)
		}
	}
}

func buildTracker() []string {
	var tracks []string

	if trackers["balance"] == 2 {
		tracks = append(tracks, "balance")
	}
	if trackers["disputeStatus"] == 2 {
		tracks = append(tracks, "disputeStatus")
	}
	if trackers["gas"] == 2 {
		tracks = append(tracks, "gas")
	}
	if trackers["tokenBalance"] == 2 {
		tracks = append(tracks, "tokenBalance")
	}
	if trackers["indexers"] == 2 {
		tracks = append(tracks, "indexers")
	}
	if trackers["newCurrentVariables"] == 2 {
		tracks = append(tracks, "newCurrentVariables")
	}
	if trackers["currentVariables"] == 2 {
		tracks = append(tracks, "currentVariables")
	}
	if trackers["disputeChecker"] == 2 {
		tracks = append(tracks, "disputeChecker")
	}
	if trackers["tallyVotes"] == 2 {
		tracks = append(tracks, "tallyVotes")
	}

	return tracks
}

func configStatus() bool {

	parsedNilDuration, _ := time.ParseDuration("0")

	nilDuration := config.Duration{parsedNilDuration}

	conf := config.GetConfig()

	if conf.TokenAddress == "" {
		return false
	}
	if conf.ContractAddress == "" {
		return false
	}
	if conf.VaultAddress == "" {
		return false
	}
	if conf.NodeURL == "" {
		return false
	}
	if conf.PublicAddress == "" {
		return false
	}
	if conf.EthClientTimeout == 0 {
		return false
	}
	if conf.TrackerSleepCycle == nilDuration {
		return false
	}
	if len(config.Trackers) == 0 {
		return false
	}
	if conf.DBFile == "" {
		return false
	}
	if conf.ServerHost == "" {
		return false
	}
	if conf.ServerPort == 0 {
		return false
	}
	if conf.GasMultiplier == 0 {
		return false
	}
	if conf.GasMax == 0 {
		return false
	}
	if len(conf.ServerWhitelist) == 0 {
		return false
	}
	if conf.DisputeTimeDelta == nilDuration {
		return false
	}
	if conf.UseGPU == true {
		return false
	}
	if conf.PrivateKey == "" {
		return false
	}
	return true
}

func setupConfig() {

	config.SetServerHost("localhost")
	config.SetServerPort(5001)
	config.SetUseGPU(false)
	config.SetGasMax(30)
	config.SetGasMultiplier(1)
	config.SetDBFile("zapDB")
	config.SetDisputeTimeDelta(6)
	// return

}

func setupMiner(statusBar *widgets.QStatusBar) {
	// setupConfig()
	//configStatus()
	return
}

// func createStatusWidget(){
// 	return
// }

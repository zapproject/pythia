package qt

import (
	"os"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/zapproject/pythia/config"
)

var cfg *config.Config

var trackers map[string]int

func App() *widgets.QApplication {
	config.NewConfig()

	trackers = map[string]int{}

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)

	window.SetMinimumSize2(980, 800)

	window.SetWindowTitle("Pythia")

	toolbar := startUp()
	widget := widgets.NewQGroupBox2("Configuration", nil)
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
		LoginRadioGrid.AddWidget(button, i, 0, 0)
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

			LoginRadioGrid.AddWidget(PrivatePublicGroup, id, 1, 0)
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

			LoginRadioGrid.AddWidget(KeyFileGroup, id, 1, 0)

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

			LoginRadioGrid.AddWidget(MnemonicGroup, id, 1, 0)

		}

		window.Show()

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
	envLayout.AddWidget(nodeURLLabel, 1, 0, 0)
	envLayout.AddWidget(nodeURLBox, 1, 1, 0)
	envLayout.AddWidget(clientTimeoutLabel, 2, 0, 0)
	envLayout.AddWidget(clientTimeoutBox, 2, 1, 0)
	envLayout.AddWidget(trackerCycleLabel, 3, 0, 0)
	envLayout.AddWidget(trackerCycleBox, 3, 1, 0)

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
	trackerLayout.AddWidget(balanceLabel, 0, 0, 0)
	trackerLayout.AddWidget(balanceBox, 0, 1, 0)
	trackerLayout.AddWidget(indexersLabel, 0, 2, 0)
	trackerLayout.AddWidget(indexersBox, 0, 3, 0)
	trackerLayout.AddWidget(disputeStatusLabel, 1, 0, 0)
	trackerLayout.AddWidget(disputeStatusBox, 1, 1, 0)
	trackerLayout.AddWidget(newCurVarsLabel, 1, 2, 0)
	trackerLayout.AddWidget(newCurVarsBox, 1, 3, 0)
	trackerLayout.AddWidget(gasLabel, 2, 0, 0)
	trackerLayout.AddWidget(gasBox, 2, 1, 0)
	trackerLayout.AddWidget(curVarsLabel, 2, 2, 0)
	trackerLayout.AddWidget(curVarsBox, 2, 3, 0)
	trackerLayout.AddWidget(tokenBalanceLabel, 3, 0, 0)
	trackerLayout.AddWidget(tokenBalanceBox, 3, 1, 0)
	trackerLayout.AddWidget(disputeCheckerLabel, 3, 2, 0)
	trackerLayout.AddWidget(disputeCheckerBox, 3, 3, 0)
	trackerLayout.AddWidget(tallyVotesLabel, 4, 0, 0)
	trackerLayout.AddWidget(tallyVotesBox, 4, 1, 0)

	trackerWidget.SetLayout(trackerLayout)

	/**
	Add grid box layouts to window
	*/
	layout.AddWidget(LoginRadioGroup, 0, 0, 0)
	layout.AddWidget(envWidget, 1, 0, 0)
	layout.AddWidget(trackerWidget, 2, 0, 0)
	// layout.AddLayout(toolbar, 0, 0, 0)
	// layout.AddLayout(widgetLayout, 0, 1, 0)

	widget.SetLayout(layout)
	window.SetCentralWidget(widget)
	window.AddToolBar(core.Qt__LeftToolBarArea, toolbar)

	// set up File menu bar
	fileMenu := window.MenuBar().AddMenu2("&File")
	fileMenu.AddAction("File Option")

	// set up settings menu bar
	settingsMenu := window.MenuBar().AddMenu2("&Settings")
	settingsMenu.AddAction("Settings Option")

	// set up help menu bar
	helpMenu := window.MenuBar().AddMenu2("&Help")
	helpMenu.AddAction("Help Option")

	window.Show()

	return app
}

func newPrivatePublicKeyWidget() *widgets.QGroupBox {
	PrivatePublicLayout := widgets.NewQVBoxLayout2(nil)
	PrivatePublicGroup := widgets.NewQGroupBox2("Enter Private Public Keys", nil)
	PublicKeyLabel := widgets.NewQLabel2("Public Key", nil, 0)
	PublicText := widgets.NewQLineEdit(nil)
	PrivateKeyLabel := widgets.NewQLabel2("Private Key", nil, 0)
	PrivateText := widgets.NewQLineEdit(nil)

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

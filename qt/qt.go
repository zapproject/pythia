package qt

import (
	"os"
	"strconv"

	"github.com/therecipe/qt/widgets"
	"github.com/zapproject/pythia/config"
)

var cfg *config.Config

func App() *widgets.QApplication {
	config.NewConfig()

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)

	window.SetMinimumSize2(980, 800)

	window.SetWindowTitle("Pythia")

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

	LoginButtonGroup.ConnectButtonToggled2(func(id int, checked bool) {

		if id == 1 && checked {
			// LoginRadioGrid.AddWidget2(button, i, 1)
		}

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
	Save Button
	*/
	// saveButton := widgets.NewQPushButton2("Save", nil)

	/**
	Add grid box layouts to window
	*/
	layout.AddWidget(LoginRadioGroup, 0, 0, 0)
	layout.AddWidget(envWidget, 1, 0, 0)
	layout.AddWidget(trackerWidget, 2, 0, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)

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

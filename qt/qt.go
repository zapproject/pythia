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


	var PrivatePublicGroup *widgets.QGroupBox = nil
	var MnemonicGroup *widgets.QGroupBox = nil

	
	LoginButtonGroup.ConnectButtonToggled2(func(id int,checked bool){

		if(id == 0 && checked){

			if MnemonicGroup != nil {
				LoginRadioGrid.RemoveWidget(MnemonicGroup)
				MnemonicGroup.DeleteLater()
			}


			PrivatePublicGroup = newPrivatePublicKeyWidget()

			LoginRadioGrid.AddWidget(PrivatePublicGroup, id, 1,0)
		}
		if(id == 1 && checked){
			
			if PrivatePublicGroup != nil {
				LoginRadioGrid.RemoveWidget(PrivatePublicGroup)
				PrivatePublicGroup.DeleteLater()
			}

			if MnemonicGroup != nil {
				LoginRadioGrid.RemoveWidget(MnemonicGroup)
				MnemonicGroup.DeleteLater()
			}

		}
		if(id == 2 && checked){

			if PrivatePublicGroup != nil {
				LoginRadioGrid.RemoveWidget(PrivatePublicGroup)
				PrivatePublicGroup.DeleteLater()
			}

			MnemonicGroup = newMnemonicWidget()

			LoginRadioGrid.AddWidget(MnemonicGroup, id, 1,0)

			
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
	clientTimeoutBox.ConnectEditingFinished(func() { clientTimeoutChanged(clientTimeoutBox, clientTimeoutLabel) })

	nodeURLBox.AddItems([]string{"Ethereum Mainnet", "Ethereum Testnet", "Binance Mainnet", "Binance Testnet"})
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
	// layout.AddWidget(saveButton, 5, 3, 0)
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

func newPrivatePublicKeyWidget() *widgets.QGroupBox{
	PrivatePublicLayout := widgets.NewQVBoxLayout2(nil)
	PrivatePublicGroup := widgets.NewQGroupBox2("Enter Private Public Keys", nil)
	PublicKeyLabel := widgets.NewQLabel2("Public Key", nil, 0)
	PublicText := widgets.NewQLineEdit(nil)
	PrivateKeyLabel := widgets.NewQLabel2("Private Key", nil, 0)
	PrivateText := widgets.NewQLineEdit(nil)


	PrivatePublicLayout.AddWidget(PublicKeyLabel,0,0)
	PrivatePublicLayout.AddWidget(PublicText,0,0)
	PrivatePublicLayout.AddWidget(PrivateKeyLabel,0,0)
	PrivatePublicLayout.AddWidget(PrivateText,0,0)

	PrivatePublicGroup.SetLayout(PrivatePublicLayout)

	return PrivatePublicGroup
}

func newMnemonicWidget() *widgets.QGroupBox{
	MnemonicLayout := widgets.NewQVBoxLayout2(nil)
	MnemonicGroup := widgets.NewQGroupBox2("Enter Mnemonic Phrase", nil)
	MnemonicLabel := widgets.NewQLabel2("Mnemonic Phrase", nil, 0)
	MnemonicText := widgets.NewQLineEdit(nil)

	MnemonicLayout.AddWidget(MnemonicLabel,0,0)
	MnemonicLayout.AddWidget(MnemonicText,0,0)


	MnemonicGroup.SetLayout(MnemonicLayout)

	return MnemonicGroup
}
func clientTimeoutChanged(line *widgets.QLineEdit, label *widgets.QLabel) {
	value, err := strconv.ParseUint(line.Text(), 10, 0)
	if err != nil {

	}
	// label.SetText(line.Text())
	config.SetEthClientTimeout(uint(value))
	cfg = config.GetConfig()
}

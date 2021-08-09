package qt

import (
	"github.com/therecipe/qt/widgets"
)

// var walletButton widgets.QPushButton

func startUp() *widgets.QToolBar {
	layout := widgets.NewQToolBar2(nil)

	walletButton := widgets.NewQPushButton2("Wallet", nil)
	configButton := widgets.NewQPushButton2("Config", nil)
	mineButton := widgets.NewQPushButton2("Mine", nil)
	disputeButton := widgets.NewQPushButton2("Dispute", nil)

	walletButton.ConnectClicked(func(click bool) { showWallet() })
	configButton.ConnectClicked(func(click bool) { showConfig() })

	layout.AddWidget(walletButton)
	layout.AddWidget(configButton)
	layout.AddWidget(mineButton)
	layout.AddWidget(disputeButton)

	return layout
}

package qt

import (
	
	"github.com/therecipe/qt/widgets"
)

func showWallet() {
	walletWidget := widgets.NewQGroupBox2("Wallet", nil)
	walletLayout := widgets.NewQGridLayout2()

	// for an option to add balances of different tokens, create group
	balanceLabel := widgets.NewQLabel2("Balance: ", nil, 0)
	// addr := ctx.Value(ZapCommon.PublicAddress).(common.Address)
	// instance := ctx.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	// zapBalance, err := instance.BalanceOf(nil, addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(addr)
	// fmt.Println(instance)

	// balanceBox := widgets.NewQLabel2(zapBalance.String(), nil, 0)

	walletLayout.AddWidget2(balanceLabel, 0, 0, 0)
	// walletLayout.AddWidget(balanceBox, 1, 0, 0)

	walletWidget.SetLayout(walletLayout)

	widget = walletWidget
	window.SetCentralWidget(widget)
	widget.Update()
}


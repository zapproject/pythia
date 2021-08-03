package qt

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/contracts2"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/vault"
)


func showWallet() {
	// cfg = config.GetConfig()
	buildContext()

	newWindow := widgets.NewQMainWindow(nil, 0)

	newWindow.SetMinimumSize2(980, 800)

	newWindow.SetWindowTitle("Pythia")

	toolbar := startUp()
	widget := widgets.NewQGroupBox2("Wallet", nil)
	layout := widgets.NewQGridLayout2()

	// for an option to add balances of different tokens, create group
	balanceLabel := widgets.NewQLabel2("Balance: ", nil, 0)
	addr := ctx.Value(ZapCommon.PublicAddress).(common.Address)
	instance := ctx.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	zapBalance, err := instance.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	balanceBox := widgets.NewQLabel2(zapBalance.String(), nil, 0)

	layout.AddWidget(balanceLabel, 0, 0, 0)
	layout.AddWidget(balanceBox, 1, 0, 0)

	widget.SetLayout(layout)
	newWindow.SetCentralWidget(widget)
	newWindow.AddToolBar(core.Qt__LeftToolBarArea, toolbar)

	// set up File menu bar
	fileMenu := newWindow.MenuBar().AddMenu2("&File")
	fileMenu.AddAction("File Option")

	// set up settings menu bar
	settingsMenu := newWindow.MenuBar().AddMenu2("&Settings")
	settingsMenu.AddAction("Settings Option")

	// set up help menu bar
	helpMenu := newWindow.MenuBar().AddMenu2("&Help")
	helpMenu.AddAction("Help Option")

	window.Hide()
	newWindow.Show()

	// window.ActivateWindow()
}

func buildContext() error {
	if !cfg.EnablePoolWorker {
		//create an rpc client
		client, err := rpc.NewClient(cfg.NodeURL)
		if err != nil {
			log.Fatal(err)
		}

		//create an instance of the Zap master contract for on-chain interactions
		tokenAddress := common.HexToAddress(cfg.TokenAddress)
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		vaultAddress := common.HexToAddress(cfg.VaultAddress)
		masterInstance, _ := contracts.NewZapMaster(contractAddress, client)
		transactorInstance, _ := contracts1.NewZapTransactor(contractAddress, client)
		newZapInstance, _ := contracts2.NewZap(contractAddress, client)
		newTransactorInstance, _ := contracts2.NewZapTransactor(contractAddress, client)
		tokenInstance, _ := token.NewZapTokenBSCTransactor(tokenAddress, client)
		vaultInstance, _ := vault.NewVaultTransactor(vaultAddress, client)

		ctx = context.WithValue(context.Background(), ZapCommon.ClientContextKey, client)
		ctx = context.WithValue(ctx, ZapCommon.ContractAddress, contractAddress)
		ctx = context.WithValue(ctx, ZapCommon.MasterContractContextKey, masterInstance)
		ctx = context.WithValue(ctx, ZapCommon.TransactorContractContextKey, transactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.TokenTransactorContractContextKey, tokenInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewZapContractContextKey, newZapInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewTransactorContractContextKey, newTransactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.VaultTransactorContractContextKey, vaultInstance)

		privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
		if err != nil {
			return fmt.Errorf("problem getting private key: %s", err.Error())
		}
		ctx = context.WithValue(ctx, ZapCommon.PrivateKey, privateKey)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		ctx = context.WithValue(ctx, ZapCommon.PublicAddress, publicAddress)

		s, err := client.IsSyncing(ctx)
		if err != nil {
			return fmt.Errorf("could not determine if Ethereum client is syncing: %v\n", err)
		}
		if s {
			return fmt.Errorf("ethereum node is still sycning with the network")
		}
	}
	return nil
}

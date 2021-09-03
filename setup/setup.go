package setup

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	cli "github.com/jawher/mow.cli"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/contracts2"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/util"
	"github.com/zapproject/pythia/vault"
)

var CTX context.Context

func ErrorHandler(err error, operation string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %s\n", operation, err.Error())
		cli.Exit(-1)
	}
}

func buildContext() error {
	cfg := config.GetConfig()
	if !cfg.EnablePoolWorker {
		//create an rpc client
		client, err := rpc.NewClient(cfg.NodeURL)
		if err != nil {
			fmt.Println(err)
			return err
		}

		//create an instance of the Zap master contract for on-chain interactions
		tokenAddress := common.HexToAddress(cfg.TokenAddress)
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		vaultAddress := common.HexToAddress(cfg.VaultAddress)
		masterInstance, _ := contracts.NewZapMaster(contractAddress, client)
		transactorInstance, _ := contracts1.NewZapTransactor(contractAddress, client)
		newZapInstance, _ := contracts2.NewZap(contractAddress, client)
		newTransactorInstance, _ := contracts2.NewZapTransactor(contractAddress, client)
		tokenTransactorInstance, _ := token.NewZapTokenBSCTransactor(tokenAddress, client)
		vaultInstance, _ := vault.NewVaultTransactor(vaultAddress, client)

		CTX = context.WithValue(context.Background(), ZapCommon.ClientContextKey, client)
		CTX = context.WithValue(CTX, ZapCommon.ContractAddress, contractAddress)
		CTX = context.WithValue(CTX, ZapCommon.MasterContractContextKey, masterInstance)
		CTX = context.WithValue(CTX, ZapCommon.TransactorContractContextKey, transactorInstance)
		CTX = context.WithValue(CTX, ZapCommon.TokenTransactorContractContextKey, tokenTransactorInstance)
		CTX = context.WithValue(CTX, ZapCommon.NewZapContractContextKey, newZapInstance)
		CTX = context.WithValue(CTX, ZapCommon.NewTransactorContractContextKey, newTransactorInstance)
		CTX = context.WithValue(CTX, ZapCommon.VaultTransactorContractContextKey, vaultInstance)
		CTX = context.WithValue(CTX, ZapCommon.TokenAddress, tokenAddress)

		privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
		if err != nil {
			return fmt.Errorf("problem getting private key: %s", err.Error())
		}
		CTX = context.WithValue(CTX, ZapCommon.PrivateKey, privateKey)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		CTX = context.WithValue(CTX, ZapCommon.PublicAddress, publicAddress)

		s, err := client.IsSyncing(CTX)
		if err != nil {
			return fmt.Errorf("could not determine if Ethereum client is syncing: %v\n", err)
		}
		if s {
			return fmt.Errorf("ethereum node is still sycning with the network")
		}
	}
	return nil
}

func AddDBToCtx(remote bool) error {
	cfg := config.GetConfig()
	//create a db instance
	os.RemoveAll(cfg.DBFile)
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		return err
	}

	var dataProxy db.DataServerProxy
	if remote {
		proxy, err := db.OpenRemoteDB(DB)
		if err != nil {
			log.Fatal(err)
		}
		dataProxy = proxy
	} else {
		proxy, err := db.OpenLocalProxy(DB)
		if err != nil {
			log.Fatal(err)
		}
		dataProxy = proxy
	}
	CTX = context.WithValue(CTX, ZapCommon.DataProxyKey, dataProxy)
	CTX = context.WithValue(CTX, ZapCommon.DBContextKey, DB)
	return nil
}

var GitTag string
var GitHash string

const versionMessage = `
	The official Pythia %s (%s)
	-----------------------------------------
	Website: https://Zap.org
	Github:  https://github.com/zapproject/pythia
	`

func App() (bool,error) {

	err := util.ParseLoggingConfig("./loggingConfig.json")
	if err != nil{
		fmt.Errorf(err.Error()) 
		return false,err
	}
	err = buildContext()
	if err != nil{
		fmt.Errorf(err.Error()) 
		return false,err
	}

	return true,nil

}


func listenTransfers(client rpc.ETHClient, cfg *config.Config) {
	tokenAddress := common.HexToAddress(cfg.TokenAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tokenAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}

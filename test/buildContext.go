package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/contracts2"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/vault"
)

var ctx context.Context
var minerCtx [6]context.Context

func buildContext() error {
	cfg := config.GetConfig()

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
		tokenListener, _ := token.NewZapTokenBSCFilterer(tokenAddress, client)
		vaultInstance, _ := vault.NewVaultTransactor(vaultAddress, client)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		ctx = context.WithValue(context.Background(), ZapCommon.ClientContextKey, client)
		ctx = context.WithValue(ctx, ZapCommon.ContractAddress, contractAddress)
		ctx = context.WithValue(ctx, ZapCommon.MasterContractContextKey, masterInstance)
		ctx = context.WithValue(ctx, ZapCommon.TransactorContractContextKey, transactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.TokenTransactorContractContextKey, tokenInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewZapContractContextKey, newZapInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewTransactorContractContextKey, newTransactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.TokenFilterContextKey, tokenListener)
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

		//Issue #55, halt if client is still syncing with Ethereum network
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

func buildMinerContext(num int) error {
	cfg := config.GetConfig()

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
		masterInstance, err := contracts.NewZapMaster(contractAddress, client)
		transactorInstance, err := contracts1.NewZapTransactor(contractAddress, client)
		newZapInstance, err := contracts2.NewZap(contractAddress, client)
		newTransactorInstance, err := contracts2.NewZapTransactor(contractAddress, client)
		tokenInstance, err := token.NewZapTokenBSCTransactor(tokenAddress, client)
		tokenListener, err := token.NewZapTokenBSCFilterer(tokenAddress, client)
		vaultInstance, _ := vault.NewVaultTransactor(vaultAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		minerCtx[num] = context.WithValue(context.Background(), ZapCommon.ClientContextKey, client)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.ContractAddress, contractAddress)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.MasterContractContextKey, masterInstance)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.TransactorContractContextKey, transactorInstance)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.TokenTransactorContractContextKey, tokenInstance)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.NewZapContractContextKey, newZapInstance)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.NewTransactorContractContextKey, newTransactorInstance)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.TokenFilterContextKey, tokenListener)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.VaultTransactorContractContextKey, vaultInstance)

		privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
		if err != nil {
			return fmt.Errorf("problem getting private key: %s", err.Error())
		}
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.PrivateKey, privateKey)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		minerCtx[num] = context.WithValue(minerCtx[num], ZapCommon.PublicAddress, publicAddress)

		//Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.IsSyncing(minerCtx[num])
		if err != nil {
			return fmt.Errorf("could not determine if Ethereum client is syncing: %v\n", err)
		}
		if s {
			return fmt.Errorf("ethereum node is still sycning with the network")
		}
	}
	return nil
}

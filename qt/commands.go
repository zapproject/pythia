package qt

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/contracts2"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/vault"
	"github.com/zapproject/pythia/ops"
)


func startMiner(){

}


func buildCtx() error {
	cfg = config.GetConfig()
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


func mineCmd() {
	remoteDS := false
	cmd.Action = func() {
		//create os kill sig listener
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		exitChannels := make([]*chan os.Signal, 0)

		cfg := config.GetConfig()
		var ds *ops.DataServerOps
		if !cfg.EnablePoolWorker {
			ErrorHandler(AddDBToCtx(*remoteDS), "\U0001F5C4 initializing database \U0001F5C4")
			if !*remoteDS {
				ch := make(chan os.Signal)
				exitChannels = append(exitChannels, &ch)

				var err error
				ds, err = ops.CreateDataServerOps(ctx, ch)
				if err != nil {
					log.Fatal(err)
				}
				//start and wait for it to be ready
				ds.Start(ctx)
				<-ds.Ready()
			}
		}
		//start miner
		DB := ctx.Value(ZapCommon.DataProxyKey).(db.DataServerProxy)
		//DB := ctx.Value(ZapCommon.DBContextKey).(db.DB)
		v, err := DB.Get(db.DisputeStatusKey)
		if err != nil {
			fmt.Println("ignoring --- could not get dispute status.  Check if staked")
		}
		status, _ := hexutil.DecodeBig(string(v))
		if status.Cmp(big.NewInt(1)) != 0 {
			log.Fatalf("\U0001F6AB Miner is not able to mine with status %v. Stopping all mining immediately \U00002622", status)
		}
		ch := make(chan os.Signal)
		exitChannels = append(exitChannels, &ch)
		miner, err := ops.CreateMiningManager(ctx, ch, ops.NewSubmitter())
		if err != nil {
			log.Fatal(err)
		}
		miner.Start(ctx)

		//now we wait for kill sig
		<-c
		//and then notify exit channels
		for _, ch := range exitChannels {
			*ch <- os.Interrupt
		}
		cnt := 0
		start := time.Now()
		for {
			cnt++
			dsStopped := false
			minerStopped := false

			if ds != nil {
				dsStopped = !ds.Running
			} else {
				dsStopped = true
			}

			if miner != nil {
				minerStopped = !miner.Running
			} else {
				minerStopped = true
			}

			if !dsStopped && !minerStopped && cnt > 60 {
				fmt.Printf("\U000026A0 Taking longer than expected to stop operations. Waited %v so far\n", time.Now().Sub(start))
			} else if dsStopped && minerStopped {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Printf("Main shutdown complete \U00002622\n")
	}
}
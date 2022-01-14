package tracker

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/util"
)

var currentVarsLog = util.NewLogger("tracker", "CurrentVarsTracker")

//CurrentVariablesTracker concrete tracker type
type CurrentVariablesTracker struct {
}

// Returns the CurrentVariablesTracker name
func (b *CurrentVariablesTracker) String() string {
	return "CurrentVariablesTracker"
}

func PrepareEthTransaction(ctx context.Context) (*bind.TransactOpts, error) {

	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)

	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)

	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		return nil, fmt.Errorf("problem getting pending nonce: %+v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("problem getting gas price: %+v", err)
	}

	ethBalance, err := client.BalanceAt(context.Background(), publicAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("problem getting balance: %+v", err)
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if ethBalance.Cmp(cost) < 0 {
		return nil, fmt.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	privateKey := ctx.Value(zapCommon.PrivateKey).(*ecdsa.PrivateKey)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

//Exec implementation for tracker
func (b *CurrentVariablesTracker) Exec(ctx context.Context) error {
	auth, _ := PrepareEthTransaction(ctx)
	instanceZ := ctx.Value(zapCommon.TransactorContractContextKey).(*contracts1.ZapTransactor)
	instanceM := ctx.Value(zapCommon.ContractAddress).(common.Address)
	instanceT := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instanceT.IncreaseApproval(auth, instanceM, new(big.Int).SetInt64(1000000))
	auth, _ = PrepareEthTransaction(ctx)
	instanceZ.RequestData(auth, "json(https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd).ethereum.usd", "ETH/USD", new(big.Int).SetInt64(1000000), new(big.Int).SetInt64(1))
	//cast client using type assertion since context holds generic interface{}
	DB := ctx.Value(zapCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)

	currentChallenge, requestID, difficulty, queryString, granularity, totalTip, err := instance.GetCurrentVariables(nil)
	if err != nil {
		fmt.Println("Current Variables Retrieval Error")
		return err
	}

	//if we've mined it, don't save it
	myStatus, err := instance.DidMine(nil, currentChallenge, fromAddress)
	if err != nil {
		fmt.Println("My Status Retrieval Error")
		return err
	}
	bitSetVar := []byte{0}
	if myStatus {
		bitSetVar = []byte{1}
	}
	currentVarsLog.Info("Retrieved variables. challengeHash: %x", currentChallenge)
	fmt.Println("VARS", currentVarsLog)

	array := [32]byte{}
	data := []byte("timeOfLastNewValue")
	data = crypto.Keccak256(data)
	for i := 0; i < 32; i++ {
		array[i] = data[i]
	}
	uvar, _ := instance.GetUintVar(nil, array)
	currentVarsLog.Info("TimeStamp: ", uvar)

	err = DB.Put(db.CurrentChallengeKey, currentChallenge[:])
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(requestID)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(difficulty)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.QueryStringKey, []byte(queryString))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.GranularityKey, []byte(hexutil.EncodeBig(granularity)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(totalTip)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}

	return DB.Put(db.MiningStatusKey, bitSetVar)
}

package tracker

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	zap "github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

type TokenTracker struct {
}

func (b *TokenTracker) String() string {
	return "TokenTracker"
}

func (b *TokenTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(zapCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	_conAddress := cfg.ContractAddress

	//convert to address
	contractAddress := common.HexToAddress(_conAddress)

	instance, err := zap.NewZapMaster(contractAddress, client)
	if err != nil {
		fmt.Println("Instance error - TokenBalance")
		return err
	}

	balance, err := instance.BalanceOf(nil, fromAddress)
	balanceInTokens, _ := big.NewFloat(1).SetString(balance.String())
	// this _should_ be unreachable given that there is an erro flag for
	// the balanceOf call
	//if !ok {
	//	fmt.Println("Problem converting tokens.")
	//	balanceInTokens = big.NewFloat(0)
	//}
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	// This is unreachable since it's hardcoded
	//if !ok {
	//	fmt.Println("Could not create token float for computing tokens")
	//	balanceInTokens = big.NewFloat(0)
	//}
	if decimals != nil {
		balanceInTokens = balanceInTokens.Quo(balanceInTokens, decimals)
	}

	//numTokens, _ := balanceInTokens.Float64()
	log.Printf("Token Balance: %v (%v tokens)\n", balance, balanceInTokens)
	if err != nil {
		fmt.Println("Balance Retrieval Error - Token Balance")
		return err
	}
	enc := hexutil.EncodeBig(balance)
	return DB.Put(db.TokenBalanceKey, []byte(enc))
}

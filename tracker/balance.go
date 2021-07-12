package tracker

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

//BalanceTracker concrete tracker type
type BalanceTracker struct {
}

// Returns the BalanceTracker name
func (b *BalanceTracker) String() string {

	return "BalanceTracker"
}

//Exec implementation for tracker
func (b *BalanceTracker) Exec(ctx context.Context) error {

	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(zapCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	// Gets the balance from an address
	balance, err := client.BalanceAt(ctx, fromAddress, nil)

	// If the balance is a negative number it will not return nil and throw an error
	if err != nil {

		fmt.Println("balance Error, balance.go")

		return err
	}

	// Encodes a BigInt balance to a hex string
	enc := hexutil.EncodeBig(balance)

	// log = Prints the date/time
	// Printf = Prints formatted data
	// %v = Used with Printf and gets the value of a variable
	log.Printf("Balance: %v", balance)

	// BalanceKey is the key to store/lookup account balance
	// Stores the converted hex string balance as a slice of bytes
	return DB.Put(db.BalanceKey, []byte(enc))
}

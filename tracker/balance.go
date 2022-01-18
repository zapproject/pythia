package tracker

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/vault"
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
	vaultAddress := ctx.Value((zapCommon.VaultAddress)).(common.Address)
	vault, err := vault.NewVaultCaller(vaultAddress, client)
	if err != nil {
		return fmt.Errorf("error with initialising vault instance to get vault balance: %+v", err)
	}

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

	// Gets the vault balance
	vaultBalance, err := vault.UserBalance(nil, fromAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Encodes a BigInt balance to a hex string
	enc := hexutil.EncodeBig(balance)

	// log = Prints the date/time
	// Printf = Prints formatted data
	// %v = Used with Printf and gets the value of a variable
	log.Printf("Wallet Balance: %v", balance)
	log.Printf("Vault Balance: %v", vaultBalance)

	// BalanceKey is the key to store/lookup account balance
	// Stores the converted hex string balance as a slice of bytes
	return DB.Put(db.BalanceKey, []byte(enc))
}

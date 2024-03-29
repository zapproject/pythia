package ops

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/rpc"
)

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
	// The above statement will be deprecated soon, the following will be supported thereafter
	// may need to add a config for chain id in which case or a conditional based on the nodeURL
	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

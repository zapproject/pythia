package ops

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/pythia/common"
	zap "github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/rpc"
	token "github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/util"
)

/**
 * This is the operational transfer component. Its purpose is to transfer zap tokens
 */

func prepareTransfer(amt *big.Int, ctx context.Context) (*bind.TransactOpts, error) {
	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	senderPubAddr := ctx.Value(zapCommon.PublicAddress).(common.Address)

	balance, err := instance.BalanceOf(nil, senderPubAddr)
	if err != nil {
		return nil, fmt.Errorf("\U0001F6AB failed to get balance: %v", err)
	}
	fmt.Println("My balance", util.FormatERC20Balance(balance))
	if balance.Cmp(amt) < 0 {
		return nil, fmt.Errorf("insufficent balance (%s ZAP), requested %s ZAP",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(amt))
	}
	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("\U0001F6AB failed to prepare ethereum transaction: %s", err.Error())
	}
	return auth, nil
}

func Transfer(toAddress common.Address, amt *big.Int, ctx context.Context) error {
	auth, err := prepareTransfer(amt, ctx)
	if err != nil {
		return err
	}

	instance2 := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	tx, err := instance2.Transfer(auth, toAddress, amt)
	if err != nil {
		return fmt.Errorf("\U0001F6AB contract failed: %s", err.Error())
	}
	fmt.Printf("\U0001F381 Transferred %s to %s \U0001F381... with tx:\n%s\n", util.FormatERC20Balance(amt), toAddress.String()[:12], tx.Hash().Hex())
	return nil
}

func Approve(_spender common.Address, amt *big.Int, ctx context.Context) error {
	auth, err := prepareTransfer(amt, ctx)
	if err != nil {
		return err
	}

	instance2 := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	tx, err := instance2.Approve(auth, _spender, amt)
	if err != nil {
		return err
	}
	fmt.Printf("Approved \U00002705 %s to %s \U00002705... with tx:\n%s\n", util.FormatERC20Balance(amt), _spender.String()[:12], tx.Hash().Hex())
	return nil
}

func Balance(ctx context.Context, addr common.Address) error {
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)

	ethBalance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return fmt.Errorf("\U0001F6AB problem getting balance: %+v \U0001F6AB", err)
	}

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	zapBalance, err := instance.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Lets \U0001F440 how much \U0001F4B0\U0001F4B0\U0001F4B0 %s has...\n", addr.String())
	fmt.Printf("%10s BNB\n", util.FormatERC20Balance(ethBalance))
	fmt.Printf("%10s ZAP\n", util.FormatERC20Balance(zapBalance))
	return nil
}

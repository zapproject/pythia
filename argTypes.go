package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zapproject/zap-miner/util"
)

type ZAPAmount struct {
	*big.Int
}

func (a *ZAPAmount) Set(v string) error {
	f, _, err := big.ParseFloat(v, 10, 256, 0)
	if err != nil {
		return err
	}
	scale := big.NewFloat(1e18)
	f.Mul(f, scale)
	a.Int, _ = f.Int(nil)
	return nil
}

func (a *ZAPAmount) String() string {
	return util.FormatERC20Balance(a.Int)
}

func (a *ZAPAmount) IsDefault() bool {
	return true
}

type ETHAddress struct {
	addr common.Address
}

func (a *ETHAddress) Set(v string) error {
	valid := common.IsHexAddress(v)
	if !valid {
		return fmt.Errorf("%s is not a valid etherum address format", v)
	}
	a.addr = common.HexToAddress(v)
	return nil
}

func (a *ETHAddress) String() string {
	return a.addr.String()
}

func (a *ETHAddress) IsDefault() bool {
	return true
}

type EthereumInt struct {
	*big.Int
}

func (b *EthereumInt) Set(v string) error {
	g := new(big.Int)
	_, ok := g.SetString(v, 10)
	if !ok {
		return fmt.Errorf("%s is not a valid integer", v)
	}
	if len(g.Bytes()) > 32 {
		return fmt.Errorf("%s is larger than 256 bits", v)
	}
	b.Int = g
	return nil
}

func (b *EthereumInt) String() string {
	return b.Int.Text(10)
}

func (b *EthereumInt) IsDefault() bool {
	return true
}

package common

import (
	"github.com/zapproject/pythia/util"
)

var (
	//ClientContextKey is the key used to set the eth client on tracker contexts
	ClientContextKey = util.NewKey("common", "ETHClient")

	//DBContextKey is the shared context key where a DB instance can be found in a context
	DBContextKey = util.NewKey("common", "DB")

	//Zap Contract Address
	ContractAddress = util.NewKey("common", "contractAddress")

	//MasterContractContextKey is the shared context key to get shared master zap contract instance
	MasterContractContextKey = util.NewKey("common", "masterContract")

	NewZapContractContextKey = util.NewKey("common", "newZapContract")

	//TransactorContractContextKey is the shared context key to get shared transactor zap contract instance
	TransactorContractContextKey = util.NewKey("common", "transactorContract")

	TokenAddress = util.NewKey("common", "tokenAddress")

	TokenTransactorContractContextKey = util.NewKey("common", "tokenTransactorContract")

	NewTransactorContractContextKey = util.NewKey("common", "newTransactorContract")

	//DataProxyKey used to access the local or remote data server proxy
	DataProxyKey = util.NewKey("common", "DataServerProxy")

	//Ethereum wallet private key
	PrivateKey = util.NewKey("common", "PrivateKey")

	//Ethereum wallet public address
	PublicAddress = util.NewKey("common", "PublicAddress")

	TokenFilterContextKey = util.NewKey("common", "tokenFilter")

	VaultTransactorContractContextKey = util.NewKey("common", "vaultTransactorContract")
)

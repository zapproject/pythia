// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SafeMathMABI is the input ABI used to generate the binding from.
const SafeMathMABI = "[]"

// SafeMathMBin is the compiled bytecode used for deploying new contracts.
var SafeMathMBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208b6c380155d79e0f0ba989571d3c9436045cf8c69475429f68cbd6e48af3d50f64736f6c63430005100032"

// DeploySafeMathM deploys a new Ethereum contract, binding an instance of SafeMathM to it.
func DeploySafeMathM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathM, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathM{SafeMathMCaller: SafeMathMCaller{contract: contract}, SafeMathMTransactor: SafeMathMTransactor{contract: contract}, SafeMathMFilterer: SafeMathMFilterer{contract: contract}}, nil
}

// SafeMathM is an auto generated Go binding around an Ethereum contract.
type SafeMathM struct {
	SafeMathMCaller     // Read-only binding to the contract
	SafeMathMTransactor // Write-only binding to the contract
	SafeMathMFilterer   // Log filterer for contract events
}

// SafeMathMCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathMSession struct {
	Contract     *SafeMathM        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathMCallerSession struct {
	Contract *SafeMathMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeMathMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathMTransactorSession struct {
	Contract     *SafeMathMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeMathMRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathMRaw struct {
	Contract *SafeMathM // Generic contract binding to access the raw methods on
}

// SafeMathMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathMCallerRaw struct {
	Contract *SafeMathMCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathMTransactorRaw struct {
	Contract *SafeMathMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathM creates a new instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathM(address common.Address, backend bind.ContractBackend) (*SafeMathM, error) {
	contract, err := bindSafeMathM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathM{SafeMathMCaller: SafeMathMCaller{contract: contract}, SafeMathMTransactor: SafeMathMTransactor{contract: contract}, SafeMathMFilterer: SafeMathMFilterer{contract: contract}}, nil
}

// NewSafeMathMCaller creates a new read-only instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathMCaller(address common.Address, caller bind.ContractCaller) (*SafeMathMCaller, error) {
	contract, err := bindSafeMathM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathMCaller{contract: contract}, nil
}

// NewSafeMathMTransactor creates a new write-only instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathMTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathMTransactor, error) {
	contract, err := bindSafeMathM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathMTransactor{contract: contract}, nil
}

// NewSafeMathMFilterer creates a new log filterer instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathMFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathMFilterer, error) {
	contract, err := bindSafeMathM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathMFilterer{contract: contract}, nil
}

// bindSafeMathM binds a generic wrapper to an already deployed contract.
func bindSafeMathM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathM *SafeMathMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathM.Contract.SafeMathMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathM *SafeMathMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathM.Contract.SafeMathMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathM *SafeMathMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathM.Contract.SafeMathMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathM *SafeMathMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathM *SafeMathMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathM *SafeMathMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathM.Contract.contract.Transact(opts, method, params...)
}

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"miniVault\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"miniVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"authorizedUser\",\"type\":\"address\"}],\"name\":\"lockSmith\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zapMaster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = map[string]string{
	"47e7ef24": "deposit(address,uint256)",
	"1709ef07": "hasAccess(address,address)",
	"cd8f942c": "increaseApproval()",
	"7a47328b": "lockSmith(address,address)",
	"0103c92b": "userBalance(address)",
	"f3fef3a3": "withdraw(address,uint256)",
	"dca6d2c5": "zapMaster()",
	"2b56eb56": "zapToken()",
}

// VaultBin is the compiled bytecode used for deploying new contracts.
var VaultBin = "0x608060405234801561001057600080fd5b50604051610a2f380380610a2f8339818101604052604081101561003357600080fd5b508051602091820151600080546001600160a01b038085166001600160a01b0319928316811790935560018054918516919092168117909155604080516024810192909252600019604480840191909152815180840390910181526064909201815294810180516001600160e01b031663095ea7b360e01b178152945181519495939492939192909182918083835b602083106100e15780518252601f1990920191602091820191016100c2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610143576040519150601f19603f3d011682016040523d82523d6000602084013e610148565b606091505b50505050506108d38061015c6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80637a47328b1161005b5780637a47328b14610159578063cd8f942c14610187578063dca6d2c51461018f578063f3fef3a31461019757610088565b80630103c92b1461008d5780631709ef07146100c55780632b56eb561461010757806347e7ef241461012b575b600080fd5b6100b3600480360360208110156100a357600080fd5b50356001600160a01b03166101c3565b60408051918252519081900360200190f35b6100f3600480360360408110156100db57600080fd5b506001600160a01b03813581169160200135166101de565b604080519115158252519081900360200190f35b61010f610262565b604080516001600160a01b039092168252519081900360200190f35b6101576004803603604081101561014157600080fd5b506001600160a01b038135169060200135610271565b005b6101576004803603604081101561016f57600080fd5b506001600160a01b0381358116916020013516610344565b6100f361047e565b61010f610684565b610157600480360360408110156101ad57600080fd5b506001600160a01b038135169060200135610693565b6001600160a01b031660009081526002602052604090205490565b6000331515806101f757506001600160a01b0382163314155b6102325760405162461bcd60e51b81526004018080602001828103825260268152602001806108796026913960400191505060405180910390fd5b506001600160a01b0380821660009081526003602090815260408083209386168352929052205460ff1692915050565b6000546001600160a01b031681565b6001600160a01b0382166102b65760405162461bcd60e51b81526004018080602001828103825260268152602001806108796026913960400191505060405180910390fd5b6102c033836101de565b6102fb5760405162461bcd60e51b815260040180806020018281038252602c815260200180610828602c913960400191505060405180910390fd5b6001600160a01b038216600090815260026020526040902054610324908263ffffffff6107a316565b6001600160a01b0390921660009081526002602052604090209190915550565b336001600160a01b038316146103a1576040805162461bcd60e51b815260206004820152601a60248201527f596f7520646f206e6f74206f776e2074686973207661756c742e000000000000604482015290519081900360640190fd5b331515806103b857506001600160a01b0382163314155b6103f35760405162461bcd60e51b81526004018080602001828103825260258152602001806108546025913960400191505060405180910390fd5b6001600160a01b038216600090815260036020908152604080832033845290915290205460ff1661044a576001600160a01b03821660009081526003602090815260408083209091529020805460ff191660011790555b6001600160a01b0391821660009081526003602090815260408083209390941682529190915220805460ff19166001179055565b60008054600154604080513060248201526001600160a01b0392831660448083019190915282518083039091018152606490910182526020810180516001600160e01b0316636eb1769f60e11b178152915181518695606095169382918083835b602083106104fe5780518252601f1990920191602091820191016104df565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610560576040519150601f19603f3d011682016040523d82523d6000602084013e610565565b606091505b5091509150600061058961057a8360006107b9565b6000199063ffffffff61081516565b60008054600154604080516001600160a01b039283166024820152604480820187905282518083039091018152606490910182526020810180516001600160e01b031663d73dd62360e01b178152915181519697509495606095939094169390929182918083835b602083106106105780518252601f1990920191602091820191016105f1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610672576040519150601f19603f3d011682016040523d82523d6000602084013e610677565b606091505b5090965050505050505090565b6001546001600160a01b031681565b6001600160a01b0382166106d85760405162461bcd60e51b81526004018080602001828103825260268152602001806108796026913960400191505060405180910390fd5b6106e233836101de565b61071d5760405162461bcd60e51b815260040180806020018281038252602c815260200180610828602c913960400191505060405180910390fd5b80610727836101c3565b101561077a576040805162461bcd60e51b815260206004820152601d60248201527f596f75722062616c616e636520697320696e73756666696369656e742e000000604482015290519081900360640190fd5b6001600160a01b038216600090815260026020526040902054610324908263ffffffff61081516565b6000828201838110156107b257fe5b9392505050565b6000816020018351101561080c576040805162461bcd60e51b8152602060048201526015602482015274746f55696e743235365f6f75744f66426f756e647360581b604482015290519081900360640190fd5b50016020015190565b60008282111561082157fe5b5090039056fe596f7520617265206e6f7420617574686f72697a656420746f206163636573732074686973207661756c742e546865207a65726f20616464726573732063616e206e6f74206f776e2061207661756c742e546865207a65726f206164647265737320646f6573206e6f74206f776e2061207661756c742ea265627a7a723158209fe29acaec20079c021202ebd8bd3f572910cf071392c6505663d36a553b4d0764736f6c63430005100032"

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, master common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, token, master)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// HasAccess is a free data retrieval call binding the contract method 0x1709ef07.
//
// Solidity: function hasAccess(address user, address miniVault) view returns(bool)
func (_Vault *VaultCaller) HasAccess(opts *bind.CallOpts, user common.Address, miniVault common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "hasAccess")
	return *ret0, err
}

// HasAccess is a free data retrieval call binding the contract method 0x1709ef07.
//
// Solidity: function hasAccess(address user, address miniVault) view returns(bool)
func (_Vault *VaultSession) HasAccess(user common.Address, miniVault common.Address) (bool, error) {
	return _Vault.Contract.HasAccess(&_Vault.CallOpts, user, miniVault)
}

// HasAccess is a free data retrieval call binding the contract method 0x1709ef07.
//
// Solidity: function hasAccess(address user, address miniVault) view returns(bool)
func (_Vault *VaultCallerSession) HasAccess(user common.Address, miniVault common.Address) (bool, error) {
	return _Vault.Contract.HasAccess(&_Vault.CallOpts, user, miniVault)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address userAddress) view returns(uint256 balance)
func (_Vault *VaultCaller) UserBalance(opts *bind.CallOpts, userAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "userBalance", userAddress)
	return *ret0, err
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address userAddress) view returns(uint256 balance)
func (_Vault *VaultSession) UserBalance(userAddress common.Address) (*big.Int, error) {
	return _Vault.Contract.UserBalance(&_Vault.CallOpts, userAddress)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address userAddress) view returns(uint256 balance)
func (_Vault *VaultCallerSession) UserBalance(userAddress common.Address) (*big.Int, error) {
	return _Vault.Contract.UserBalance(&_Vault.CallOpts, userAddress)
}

// ZapMaster is a free data retrieval call binding the contract method 0xdca6d2c5.
//
// Solidity: function zapMaster() view returns(address)
func (_Vault *VaultCaller) ZapMaster(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "zapMaster")
	return *ret0, err
}

// ZapMaster is a free data retrieval call binding the contract method 0xdca6d2c5.
//
// Solidity: function zapMaster() view returns(address)
func (_Vault *VaultSession) ZapMaster() (common.Address, error) {
	return _Vault.Contract.ZapMaster(&_Vault.CallOpts)
}

// ZapMaster is a free data retrieval call binding the contract method 0xdca6d2c5.
//
// Solidity: function zapMaster() view returns(address)
func (_Vault *VaultCallerSession) ZapMaster() (common.Address, error) {
	return _Vault.Contract.ZapMaster(&_Vault.CallOpts)
}

// ZapToken is a free data retrieval call binding the contract method 0x2b56eb56.
//
// Solidity: function zapToken() view returns(address)
func (_Vault *VaultCaller) ZapToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "zapToken")
	return *ret0, err
}

// ZapToken is a free data retrieval call binding the contract method 0x2b56eb56.
//
// Solidity: function zapToken() view returns(address)
func (_Vault *VaultSession) ZapToken() (common.Address, error) {
	return _Vault.Contract.ZapToken(&_Vault.CallOpts)
}

// ZapToken is a free data retrieval call binding the contract method 0x2b56eb56.
//
// Solidity: function zapToken() view returns(address)
func (_Vault *VaultCallerSession) ZapToken() (common.Address, error) {
	return _Vault.Contract.ZapToken(&_Vault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", userAddress, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address userAddress, uint256 value) returns()
func (_Vault *VaultSession) Deposit(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, userAddress, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactorSession) Deposit(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, userAddress, value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xcd8f942c.
//
// Solidity: function increaseApproval() returns(bool)
func (_Vault *VaultTransactor) IncreaseApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "increaseApproval")
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xcd8f942c.
//
// Solidity: function increaseApproval() returns(bool)
func (_Vault *VaultSession) IncreaseApproval() (*types.Transaction, error) {
	return _Vault.Contract.IncreaseApproval(&_Vault.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xcd8f942c.
//
// Solidity: function increaseApproval() returns(bool)
func (_Vault *VaultTransactorSession) IncreaseApproval() (*types.Transaction, error) {
	return _Vault.Contract.IncreaseApproval(&_Vault.TransactOpts)
}

// LockSmith is a paid mutator transaction binding the contract method 0x7a47328b.
//
// Solidity: function lockSmith(address miniVault, address authorizedUser) returns()
func (_Vault *VaultTransactor) LockSmith(opts *bind.TransactOpts, miniVault common.Address, authorizedUser common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "lockSmith", miniVault, authorizedUser)
}

// LockSmith is a paid mutator transaction binding the contract method 0x7a47328b.
//
// Solidity: function lockSmith(address miniVault, address authorizedUser) returns()
func (_Vault *VaultSession) LockSmith(miniVault common.Address, authorizedUser common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LockSmith(&_Vault.TransactOpts, miniVault, authorizedUser)
}

// LockSmith is a paid mutator transaction binding the contract method 0x7a47328b.
//
// Solidity: function lockSmith(address miniVault, address authorizedUser) returns()
func (_Vault *VaultTransactorSession) LockSmith(miniVault common.Address, authorizedUser common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LockSmith(&_Vault.TransactOpts, miniVault, authorizedUser)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", userAddress, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address userAddress, uint256 value) returns()
func (_Vault *VaultSession) Withdraw(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, userAddress, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactorSession) Withdraw(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, userAddress, value)
}

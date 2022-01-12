// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts1

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b80f64b1e9c68df89ed77756c24ebdda9fb8b2bd0e13dd44adf4ad84eb758ef064736f6c634300080b0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathMMetaData contains all meta data concerning the SafeMathM contract.
var SafeMathMMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122040dcc6bc9661f40904118719192b84d6b2c85eaf081e8ff47659416acea1f67464736f6c634300080b0033",
}

// SafeMathMABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMMetaData.ABI instead.
var SafeMathMABI = SafeMathMMetaData.ABI

// SafeMathMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMMetaData.Bin instead.
var SafeMathMBin = SafeMathMMetaData.Bin

// DeploySafeMathM deploys a new Ethereum contract, binding an instance of SafeMathM to it.
func DeploySafeMathM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathM, error) {
	parsed, err := SafeMathMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathMBin), backend)
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

// UtilitiesMetaData contains all meta data concerning the Utilities contract.
var UtilitiesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122085d6c25242508f07fcf0e261657213e292a00cc59d1851782d2ef4f4e2be196864736f6c634300080b0033",
}

// UtilitiesABI is the input ABI used to generate the binding from.
// Deprecated: Use UtilitiesMetaData.ABI instead.
var UtilitiesABI = UtilitiesMetaData.ABI

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UtilitiesMetaData.Bin instead.
var UtilitiesBin = UtilitiesMetaData.Bin

// DeployUtilities deploys a new Ethereum contract, binding an instance of Utilities to it.
func DeployUtilities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utilities, error) {
	parsed, err := UtilitiesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UtilitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// Utilities is an auto generated Go binding around an Ethereum contract.
type Utilities struct {
	UtilitiesCaller     // Read-only binding to the contract
	UtilitiesTransactor // Write-only binding to the contract
	UtilitiesFilterer   // Log filterer for contract events
}

// UtilitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilitiesSession struct {
	Contract     *Utilities        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilitiesCallerSession struct {
	Contract *UtilitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UtilitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilitiesTransactorSession struct {
	Contract     *UtilitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UtilitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilitiesRaw struct {
	Contract *Utilities // Generic contract binding to access the raw methods on
}

// UtilitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilitiesCallerRaw struct {
	Contract *UtilitiesCaller // Generic read-only contract binding to access the raw methods on
}

// UtilitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilitiesTransactorRaw struct {
	Contract *UtilitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtilities creates a new instance of Utilities, bound to a specific deployed contract.
func NewUtilities(address common.Address, backend bind.ContractBackend) (*Utilities, error) {
	contract, err := bindUtilities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// NewUtilitiesCaller creates a new read-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesCaller(address common.Address, caller bind.ContractCaller) (*UtilitiesCaller, error) {
	contract, err := bindUtilities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesCaller{contract: contract}, nil
}

// NewUtilitiesTransactor creates a new write-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilitiesTransactor, error) {
	contract, err := bindUtilities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesTransactor{contract: contract}, nil
}

// NewUtilitiesFilterer creates a new log filterer instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilitiesFilterer, error) {
	contract, err := bindUtilities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilitiesFilterer{contract: contract}, nil
}

// bindUtilities binds a generic wrapper to an already deployed contract.
func bindUtilities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.UtilitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transact(opts, method, params...)
}

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZapMasterAddress\",\"type\":\"address\"}],\"name\":\"NewZapMasterEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"zapMasterAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"zapTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrateVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"}],\"name\":\"setApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVault\",\"type\":\"address\"}],\"name\":\"setNewVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f2a40db8": "accounts(uint256)",
		"47e7ef24": "deposit(address,uint256)",
		"ffeffee7": "getZM()",
		"9aa22f6f": "getZT()",
		"cd8f942c": "increaseApproval()",
		"ea37c3b5": "migrateVault()",
		"e6c681c3": "setApproval(address)",
		"e7b77f70": "setNewVault(address)",
		"0103c92b": "userBalance(address)",
		"f3fef3a3": "withdraw(address,uint256)",
		"2b56eb56": "zapToken()",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200108c3803806200108c83398101604081905262000034916200012a565b60008054610100600160a81b0319166101006001600160a01b0385811691820292909217909255600280546001600160a01b03191691841691821790556040516024810191909152600019604482015260640160408051601f198184030181529181526020820180516001600160e01b031663095ea7b360e01b17905251620000be919062000162565b6000604051808303816000865af19150503d8060008114620000fd576040519150601f19603f3d011682016040523d82523d6000602084013e62000102565b606091505b5050505050620001a0565b80516001600160a01b03811681146200012557600080fd5b919050565b600080604083850312156200013e57600080fd5b62000149836200010d565b915062000159602084016200010d565b90509250929050565b6000825160005b8181101562000185576020818601810151858301520162000169565b8181111562000195576000828501525b509190910192915050565b610edc80620001b06000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063e6c681c311610071578063e6c681c31461015d578063e7b77f7014610170578063ea37c3b514610183578063f2a40db81461018b578063f3fef3a31461019e578063ffeffee7146101b157600080fd5b80630103c92b146100ae5780632b56eb56146100ea57806347e7ef241461011a5780639aa22f6f1461012f578063cd8f942c14610145575b600080fd5b6100d76100bc366004610c59565b6001600160a01b031660009081526005602052604090205490565b6040519081526020015b60405180910390f35b6000546101029061010090046001600160a01b031681565b6040516001600160a01b0390911681526020016100e1565b61012d610128366004610c7d565b6101c2565b005b60005461010090046001600160a01b0316610102565b61014d610312565b60405190151581526020016100e1565b61012d61016b366004610c59565b610489565b61014d61017e366004610c59565b610555565b61014d610625565b610102610199366004610ca9565b6109e6565b61012d6101ac366004610c7d565b610a10565b6002546001600160a01b0316610102565b6002546001600160a01b03166101d757600080fd5b6002546001600160a01b03163314806101ff57503360009081526006602052604090205460ff165b8061020957503330145b61022e5760405162461bcd60e51b815260040161022590610cc2565b60405180910390fd5b6001600160a01b0382166102545760405162461bcd60e51b815260040161022590610d3d565b6001600160a01b0382166000908152600560205260409020546102ce57600380546001600160a01b0384166000818152600460205260408120839055600183018455929092527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b03191690911790555b6001600160a01b0382166000908152600560205260409020546102f2908290610d99565b6001600160a01b0390921660009081526005602052604090209190915550565b600080546002546040513060248201526001600160a01b039182166044820152839261010090049091169060640160408051601f198184030181529181526020820180516001600160e01b0316636eb1769f60e11b179052516103759190610db1565b6000604051808303816000865af19150503d80600081146103b2576040519150601f19603f3d011682016040523d82523d6000602084013e6103b7565b606091505b5091505060006103c8826000610be3565b6103d490600019610dec565b600080546002546040516001600160a01b0391821660248201526044810185905293945091926101009091049091169060640160408051601f198184030181529181526020820180516001600160e01b031663d73dd62360e01b1790525161043c9190610db1565b6000604051808303816000865af19150503d8060008114610479576040519150601f19603f3d011682016040523d82523d6000602084013e61047e565b606091505b509095945050505050565b6002546001600160a01b031661049e57600080fd5b6002546001600160a01b031633146104c85760405162461bcd60e51b815260040161022590610e03565b60005460ff16156105275760405162461bcd60e51b815260206004820152602360248201527f43616e6e6f742073657420617070726f76616c206166746572206d696772617460448201526234b7b760e91b6064820152608401610225565b6001600160a01b031660009081526006602052604081208054600160ff199182168117909255825416179055565b6002546000906001600160a01b031661056d57600080fd5b6002546001600160a01b031633146105975760405162461bcd60e51b815260040161022590610e03565b6001600160a01b0382166106025760405162461bcd60e51b815260206004820152602c60248201527f43616e6e6f742073657420746865207a65726f2061646472657373206173207460448201526b1a19481b995dc815985d5b1d60a21b6064820152608401610225565b50600180546001600160a01b0383166001600160a01b0319909116178155919050565b6002546000906001600160a01b031661063d57600080fd5b6002546001600160a01b031633146106675760405162461bcd60e51b815260040161022590610e03565b6001546001600160a01b03166106d35760405162461bcd60e51b815260206004820152602b60248201527f43616e27742073657420746865207a65726f206164647265737320617320746860448201526a19481b995dc815985d5b1d60aa1b6064820152608401610225565b6002546001546040805160016210011960e01b0319815290516001600160a01b03938416939092169163ffeffee7916004808201926020929091908290030181865afa158015610727573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074b9190610e58565b6001600160a01b0316146107ba5760405162461bcd60e51b815260206004820152603060248201527f546865206e6577207661756c74206d7573742073686172652073616d65205a6160448201526f1c13585cdd195c8818dbdb9d1c9858dd60821b6064820152608401610225565b60005460015460408051639aa22f6f60e01b815290516101009093046001600160a01b0390811693921691639aa22f6f916004808201926020929091908290030181865afa158015610810573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108349190610e58565b6001600160a01b0316146108a65760405162461bcd60e51b815260206004820152603360248201527f546865206e6577207661756c74206d757374207368617265207468652073616d604482015272194816985c151bdad95b8818dbdb9d1c9858dd606a1b6064820152608401610225565b60008060005b6003548110156109dc57600381815481106108c9576108c9610e75565b60009182526020808320909101546001600160a01b03168083526005909152604090912054909250156109ca5760056000836001600160a01b03166001600160a01b0316815260200190815260200160002054925061096382600560006003858154811061093957610939610e75565b60009182526020808320909101546001600160a01b03168352820192909252604001902054610a10565b6001546040516311f9fbc960e21b81526001600160a01b03848116600483015260248201869052909116906347e7ef2490604401600060405180830381600087803b1580156109b157600080fd5b505af11580156109c5573d6000803e3d6000fd5b505050505b806109d481610e8b565b9150506108ac565b5060019250505090565b600381815481106109f657600080fd5b6000918252602090912001546001600160a01b0316905081565b6002546001600160a01b0316610a2557600080fd5b6002546001600160a01b0316331480610a4d57503360009081526006602052604090205460ff165b80610a5757503330145b610a735760405162461bcd60e51b815260040161022590610cc2565b6001600160a01b038216610a995760405162461bcd60e51b815260040161022590610d3d565b80610ab9836001600160a01b031660009081526005602052604090205490565b1015610b075760405162461bcd60e51b815260206004820152601d60248201527f596f75722062616c616e636520697320696e73756666696369656e742e0000006044820152606401610225565b6001600160a01b038216600090815260056020526040902054610b2b908290610dec565b610bbf576001600160a01b038216600090815260046020526040902054600380549091908110610b5d57610b5d610e75565b6000918252602080832090910180546001600160a01b03191690556001600160a01b03841682526004815260408083208390556005909152902054610ba3908290610dec565b50506001600160a01b0316600090815260056020526040812055565b6001600160a01b0382166000908152600560205260409020546102f2908290610dec565b6000610bf0826020610d99565b83511015610c385760405162461bcd60e51b8152602060048201526015602482015274746f55696e743235365f6f75744f66426f756e647360581b6044820152606401610225565b50016020015190565b6001600160a01b0381168114610c5657600080fd5b50565b600060208284031215610c6b57600080fd5b8135610c7681610c41565b9392505050565b60008060408385031215610c9057600080fd5b8235610c9b81610c41565b946020939093013593505050565b600060208284031215610cbb57600080fd5b5035919050565b60208082526055908201527f5661756c743a204f6e6c7920746865205a61704d617374657220636f6e74726160408201527f6374206f7220616e20617574686f72697a6564205661756c7420436f6e74726160608201527418dd0818d85b881b585ad9481d1a1a5cc818d85b1b605a1b608082015260a00190565b60208082526026908201527f546865207a65726f206164647265737320646f6573206e6f74206f776e2061206040820152653b30bab63a1760d11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60008219821115610dac57610dac610d83565b500190565b6000825160005b81811015610dd25760208186018101518583015201610db8565b81811115610de1576000828501525b509190910192915050565b600082821015610dfe57610dfe610d83565b500390565b60208082526035908201527f5661756c743a204f6e6c7920746865205a61704d617374657220636f6e74726160408201527418dd0818d85b881b585ad9481d1a1a5cc818d85b1b605a1b606082015260800190565b600060208284031215610e6a57600080fd5b8151610c7681610c41565b634e487b7160e01b600052603260045260246000fd5b6000600019821415610e9f57610e9f610d83565b506001019056fea2646970667358221220a88b0b8fbbce1e59270832b8d04681f8b7182e6254d332ccb7ce8acc076815b664736f6c634300080b0033",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Deprecated: Use VaultMetaData.Sigs instead.
// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = VaultMetaData.Sigs

// VaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultMetaData.Bin instead.
var VaultBin = VaultMetaData.Bin

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, master common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultBin), backend, token, master)
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

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Vault *VaultCaller) Accounts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "accounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Vault *VaultSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.Accounts(&_Vault.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Vault *VaultCallerSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.Accounts(&_Vault.CallOpts, arg0)
}

// GetZM is a free data retrieval call binding the contract method 0xffeffee7.
//
// Solidity: function getZM() view returns(address zapMasterAddress)
func (_Vault *VaultCaller) GetZM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getZM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZM is a free data retrieval call binding the contract method 0xffeffee7.
//
// Solidity: function getZM() view returns(address zapMasterAddress)
func (_Vault *VaultSession) GetZM() (common.Address, error) {
	return _Vault.Contract.GetZM(&_Vault.CallOpts)
}

// GetZM is a free data retrieval call binding the contract method 0xffeffee7.
//
// Solidity: function getZM() view returns(address zapMasterAddress)
func (_Vault *VaultCallerSession) GetZM() (common.Address, error) {
	return _Vault.Contract.GetZM(&_Vault.CallOpts)
}

// GetZT is a free data retrieval call binding the contract method 0x9aa22f6f.
//
// Solidity: function getZT() view returns(address zapTokenAddress)
func (_Vault *VaultCaller) GetZT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getZT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZT is a free data retrieval call binding the contract method 0x9aa22f6f.
//
// Solidity: function getZT() view returns(address zapTokenAddress)
func (_Vault *VaultSession) GetZT() (common.Address, error) {
	return _Vault.Contract.GetZT(&_Vault.CallOpts)
}

// GetZT is a free data retrieval call binding the contract method 0x9aa22f6f.
//
// Solidity: function getZT() view returns(address zapTokenAddress)
func (_Vault *VaultCallerSession) GetZT() (common.Address, error) {
	return _Vault.Contract.GetZT(&_Vault.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address userAddress) view returns(uint256 balance)
func (_Vault *VaultCaller) UserBalance(opts *bind.CallOpts, userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "userBalance", userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// ZapToken is a free data retrieval call binding the contract method 0x2b56eb56.
//
// Solidity: function zapToken() view returns(address)
func (_Vault *VaultCaller) ZapToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "zapToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// MigrateVault is a paid mutator transaction binding the contract method 0xea37c3b5.
//
// Solidity: function migrateVault() returns(bool success)
func (_Vault *VaultTransactor) MigrateVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "migrateVault")
}

// MigrateVault is a paid mutator transaction binding the contract method 0xea37c3b5.
//
// Solidity: function migrateVault() returns(bool success)
func (_Vault *VaultSession) MigrateVault() (*types.Transaction, error) {
	return _Vault.Contract.MigrateVault(&_Vault.TransactOpts)
}

// MigrateVault is a paid mutator transaction binding the contract method 0xea37c3b5.
//
// Solidity: function migrateVault() returns(bool success)
func (_Vault *VaultTransactorSession) MigrateVault() (*types.Transaction, error) {
	return _Vault.Contract.MigrateVault(&_Vault.TransactOpts)
}

// SetApproval is a paid mutator transaction binding the contract method 0xe6c681c3.
//
// Solidity: function setApproval(address oldVault) returns()
func (_Vault *VaultTransactor) SetApproval(opts *bind.TransactOpts, oldVault common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setApproval", oldVault)
}

// SetApproval is a paid mutator transaction binding the contract method 0xe6c681c3.
//
// Solidity: function setApproval(address oldVault) returns()
func (_Vault *VaultSession) SetApproval(oldVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetApproval(&_Vault.TransactOpts, oldVault)
}

// SetApproval is a paid mutator transaction binding the contract method 0xe6c681c3.
//
// Solidity: function setApproval(address oldVault) returns()
func (_Vault *VaultTransactorSession) SetApproval(oldVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetApproval(&_Vault.TransactOpts, oldVault)
}

// SetNewVault is a paid mutator transaction binding the contract method 0xe7b77f70.
//
// Solidity: function setNewVault(address _newVault) returns(bool success)
func (_Vault *VaultTransactor) SetNewVault(opts *bind.TransactOpts, _newVault common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setNewVault", _newVault)
}

// SetNewVault is a paid mutator transaction binding the contract method 0xe7b77f70.
//
// Solidity: function setNewVault(address _newVault) returns(bool success)
func (_Vault *VaultSession) SetNewVault(_newVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetNewVault(&_Vault.TransactOpts, _newVault)
}

// SetNewVault is a paid mutator transaction binding the contract method 0xe7b77f70.
//
// Solidity: function setNewVault(address _newVault) returns(bool success)
func (_Vault *VaultTransactorSession) SetNewVault(_newVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetNewVault(&_Vault.TransactOpts, _newVault)
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

// VaultNewZapMasterEventIterator is returned from FilterNewZapMasterEvent and is used to iterate over the raw logs and unpacked data for NewZapMasterEvent events raised by the Vault contract.
type VaultNewZapMasterEventIterator struct {
	Event *VaultNewZapMasterEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultNewZapMasterEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultNewZapMasterEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultNewZapMasterEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultNewZapMasterEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultNewZapMasterEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultNewZapMasterEvent represents a NewZapMasterEvent event raised by the Vault contract.
type VaultNewZapMasterEvent struct {
	NewZapMasterAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewZapMasterEvent is a free log retrieval operation binding the contract event 0xf9ee346654b71cdd8eca09afc1b80db994985b920df52a2956e74d72ccc4a7e4.
//
// Solidity: event NewZapMasterEvent(address _newZapMasterAddress)
func (_Vault *VaultFilterer) FilterNewZapMasterEvent(opts *bind.FilterOpts) (*VaultNewZapMasterEventIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "NewZapMasterEvent")
	if err != nil {
		return nil, err
	}
	return &VaultNewZapMasterEventIterator{contract: _Vault.contract, event: "NewZapMasterEvent", logs: logs, sub: sub}, nil
}

// WatchNewZapMasterEvent is a free log subscription operation binding the contract event 0xf9ee346654b71cdd8eca09afc1b80db994985b920df52a2956e74d72ccc4a7e4.
//
// Solidity: event NewZapMasterEvent(address _newZapMasterAddress)
func (_Vault *VaultFilterer) WatchNewZapMasterEvent(opts *bind.WatchOpts, sink chan<- *VaultNewZapMasterEvent) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "NewZapMasterEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultNewZapMasterEvent)
				if err := _Vault.contract.UnpackLog(event, "NewZapMasterEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewZapMasterEvent is a log parse operation binding the contract event 0xf9ee346654b71cdd8eca09afc1b80db994985b920df52a2956e74d72ccc4a7e4.
//
// Solidity: event NewZapMasterEvent(address _newZapMasterAddress)
func (_Vault *VaultFilterer) ParseNewZapMasterEvent(log types.Log) (*VaultNewZapMasterEvent, error) {
	event := new(VaultNewZapMasterEvent)
	if err := _Vault.contract.UnpackLog(event, "NewZapMasterEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapMetaData contains all meta data concerning the Zap contract.
var ZapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zapTokenBsc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedContract\",\"type\":\"address\"}],\"name\":\"NewForkProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"increaseVaultApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewZapAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"forkedContract\",\"type\":\"uint256\"}],\"name\":\"proposeFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_c_sapi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_c_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"requestData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"752d49a1": "addTip(uint256,uint256)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"8581af19": "beginDispute(uint256,uint256,uint256)",
		"0d2d76a2": "depositStake()",
		"4049f198": "getNewCurrentVariables()",
		"7cf26110": "increaseVaultApproval(address)",
		"8da5cb5b": "owner()",
		"2bd326af": "proposeFork(address,uint256)",
		"3fff2816": "requestData(string,string,uint256,uint256)",
		"28449c3a": "requestStakingWithdraw()",
		"68c180d5": "submitMiningSolution(string,uint256,uint256)",
		"4d318b0e": "tallyVotes(uint256)",
		"fc0c546a": "token()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"c9d27afe": "vote(uint256,bool)",
		"bed9d861": "withdrawStake()",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200313538038062003135833981016040819052620000349162000063565b604980546001600160a01b039092166001600160a01b0319928316179055604a80549091163317905562000095565b6000602082840312156200007657600080fd5b81516001600160a01b03811681146200008e57600080fd5b9392505050565b61309080620000a56000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806370a08231116100a25780638da5cb5b116100715780638da5cb5b14610226578063a9059cbb14610251578063bed9d86114610264578063c9d27afe1461026c578063fc0c546a1461027f57600080fd5b806370a08231146101cc578063752d49a1146101ed5780637cf26110146102005780638581af191461021357600080fd5b80632bd326af116100e95780632bd326af146101685780633fff28161461017b5780634049f1981461018e5780634d318b0e146101a657806368c180d5146101b957600080fd5b8063095ea7b31461011b5780630d2d76a21461014357806323b872dd1461014d57806328449c3a14610160575b600080fd5b61012e61012936600461295b565b610292565b60405190151581526020015b60405180910390f35b61014b61030f565b005b61014b61015b366004612987565b610584565b61014b6105f8565b61014b61017636600461295b565b61065d565b61014b610189366004612a11565b610728565b610196610a2c565b60405161013a9493929190612a8f565b61014b6101b4366004612ad3565b610a51565b61014b6101c7366004612aec565b611164565b6101df6101da366004612b3d565b611434565b60405190815260200161013a565b61014b6101fb366004612b5a565b6114a9565b61012e61020e366004612b3d565b6115a7565b61014b610221366004612b7c565b61160d565b604a54610239906001600160a01b031681565b6040516001600160a01b03909116815260200161013a565b61014b61025f36600461295b565b611d37565b61014b611d64565b61014b61027a366004612bb6565b611f31565b604954610239906001600160a01b031681565b60495460405163095ea7b360e01b81526000916001600160a01b03169063095ea7b3906102c59086908690600401612be6565b6020604051808303816000875af11580156102e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103089190612bff565b9392505050565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b2600052604060208190527fc8e2509b51f0dfee411d330b6e0467af3d0596bb7e449f8db658eed93b9b05c85460495491516370a0823160e01b8152336004820152909182916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156103a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103cb9190612c1c565b101561041e5760405162461bcd60e51b815260206004820152601760248201527f4e6f7420656e6f756768205a415020746f207374616b6500000000000000000060448201526064015b60405180910390fd5b60405163326991a360e01b81526000600482015273__$d4d3c96f4ab0b13728ed5056adc694f4da$__9063326991a39060240160006040518083038186803b15801561046957600080fd5b505af415801561047d573d6000803e3d6000fd5b50507fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb287836600090815260406020527f39efff102fba42ea00eb856629a713fead698b2401b1a7fdf4afafc9ef898fd180548594509092506104de908490612c4b565b9091555050600080516020612f59833981519152600052603f602052600080516020612f99833981519152546001600160a01b03168061051f338285610584565b6040516311f9fbc960e21b81526001600160a01b038216906347e7ef249061054d9033908790600401612be6565b600060405180830381600087803b15801561056757600080fd5b505af115801561057b573d6000803e3d6000fd5b50505050505050565b6049546040516001600160a01b0385811660248301528481166044830152606482018490526105f39216906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261203f565b505050565b604051633c73482760e01b81526000600482015273__$d4d3c96f4ab0b13728ed5056adc694f4da$__90633c7348279060240160006040518083038186803b15801561064357600080fd5b505af4158015610657573d6000803e3d6000fd5b50505050565b604051630d97f62f60e31b8152600060048201526001600160a01b03831660248201526044810182905273__$f08294d4d1c904fbf7968e6213a9d2a1b1$__90636cbfb1789060640160006040518083038186803b1580156106be57600080fd5b505af41580156106d2573d6000803e3d6000fd5b505060008051602061301b83398151915254600080516020612ffb833981519152600052604060205260008051602061303b8339815191525461072493503392506001600160a01b0390911690610584565b5050565b6000821161073557600080fd5b670de0b6b3a764000082111561074a57600080fd5b600086868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250604080516020601f8c018190048102820181019092528a8152959650909493508992508891508190840183828082843760009201919091525050845192935050506107c857600080fd5b60408151106107d657600080fd5b600082856040516020016107eb929190612c8f565b60408051601f19818403018152918152815160209283012060008181526047909352912054909150610a08577f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc98600090815260406020527ff2ef5143d163eb08a31276b5aa31e419ea26b4db10668507f424160a6efc34ab80549161086f83612cb1565b90915550507ff2ef5143d163eb08a31276b5aa31e419ea26b4db10668507f424160a6efc34ab54600081815260466020908152604090912085516108b592870190612852565b50600081815260466020908152604090912084516108db92600190920191860190612852565b506000818152604660208181526040808420600281018790558151858152808401928390529486905292909152915161091a92600390920191906128d6565b5060008181526046602090815260408083207fba3571a50e0c436953d31396edfb65be5925bcc7fef5a3441ed5d43dbce2548f84526004018252808320899055600080516020612fb98339815191528352808320839055600080516020612f7983398151915283528083208390558483526047909152902081905584156109a6576109a6333087610584565b6109b081866120ee565b600081815260466020526040908190209051829133917f6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9916109fa9160018201908c908c90612da7565b60405180910390a350610a21565b600081815260476020526040902054610a2190856114a9565b505050505050505050565b6000610a36612910565b600080610a436000612555565b935093509350935090919293565b33600090815260456020526040902054600114610aa85760405162461bcd60e51b815260206004820152601560248201527410d85b1b195c881b5d5cdd081899481cdd185ad959605a1b6044820152606401610415565b600080516020612f598339815191526000908152603f602052600080516020612f9983398151915254604051633cedafe560e01b815260048101839052602481018490526001600160a01b039091169190819073__$f08294d4d1c904fbf7968e6213a9d2a1b1$__90633cedafe590604401606060405180830381865af4158015610b37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5b9190612de0565b6000878152604460205260409020600381015492955090935091506060908590610e4f576040516311f9fbc960e21b81526001600160a01b038216906347e7ef2490610bad9088908890600401612be6565b600060405180830381600087803b158015610bc757600080fd5b505af1158015610bdb573d6000803e3d6000fd5b5050505060058301546001600160a01b0386811691161415610e4f577f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b2600052604060208190527fc8e2509b51f0dfee411d330b6e0467af3d0596bb7e449f8db658eed93b9b05c854600480860154925163f3fef3a360e01b815291926001600160a01b038086169363f3fef3a393610c7a9390921691869101612be6565b600060405180830381600087803b158015610c9457600080fd5b505af1158015610ca8573d6000803e3d6000fd5b50506040516311f9fbc960e21b81526001600160a01b03851692506347e7ef249150610cda9089908590600401612be6565b600060405180830381600087803b158015610cf457600080fd5b505af1158015610d08573d6000803e3d6000fd5b50505050600484810154604051630103c92b60e01b81526001600160a01b039182169281019290925260009190841690630103c92b90602401602060405180830381865afa158015610d5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d829190612c1c565b60048087015460405163f3fef3a360e01b81529293506001600160a01b038087169363f3fef3a393610dba9390921691869101612be6565b600060405180830381600087803b158015610dd457600080fd5b505af1158015610de8573d6000803e3d6000fd5b50506040516311f9fbc960e21b81526001600160a01b03861692506347e7ef249150610e1a9030908590600401612be6565b600060405180830381600087803b158015610e3457600080fd5b505af1158015610e48573d6000803e3d6000fd5b5050505050505b600283600301541415610f38576049546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa158015610ea5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec99190612c1c565b60068501546040519192506001600160a01b031690610eee9082908490602401612be6565b60408051601f198184030181529190526020810180516001600160e01b031663a9059cbb60e01b179052604954909450610f31906001600160a01b03168561203f565b505061057b565b60038360030154141561057b57600683015460405163e6c681c360e01b81526001600160a01b0388811660048301529091169063e6c681c390602401600060405180830381600087803b158015610f8e57600080fd5b505af1158015610fa2573d6000803e3d6000fd5b50505050610faf866115a7565b5060068301546049546040516370a0823160e01b81526001600160a01b03808a16600483015261102c938a93908216929116906370a0823190602401602060405180830381865afa158015611008573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061015b9190612c1c565b6006830154604051630e7b77f760e41b81526001600160a01b0391821660048201529082169063e7b77f70906024016020604051808303816000875af115801561107a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109e9190612bff565b1561111357806001600160a01b031663ea37c3b56040518163ffffffff1660e01b81526004016020604051808303816000875af11580156110e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111079190612bff565b61111357611113612e23565b6006830154600080516020612f59833981519152600052603f602052600080516020612f9983398151915280546001600160a01b0319166001600160a01b0390921691909117905550505050505050565b60405163d723552b60e01b815273__$7f5e830ec774d87c329c85d735d7e0ce58$__9063d723552b906111a4906000908890889088908890600401612e39565b60006040518083038186803b1580156111bc57600080fd5b505af41580156111d0573d6000803e3d6000fd5b50506040805160a081019091526000925090506035600583835b8282101561122b5760408051808201909152600283028501805482526001908101546001600160a01b031660208084019190915291835290920191016111ea565b5050600080516020612f99833981519152547fdcf69331608ae4837870755f6bd10a16da39d47dbece9e37a5b80a54219d7ed360005260406020527f7cc8b5af461a3e5896300862e4f0c9851aff7b33aa654c13ef6153a737a0e3ad549394506001600160a01b031692839250905080156113dd5760005b600581101561136e5760008582600581106112c0576112c0612e80565b6020020151602001516001600160a01b03161461135c576112e18383611d37565b826001600160a01b03166347e7ef2486836005811061130257611302612e80565b602002015160200151846040518363ffffffff1660e01b8152600401611329929190612be6565b600060405180830381600087803b15801561134357600080fd5b505af1158015611357573d6000803e3d6000fd5b505050505b8061136681612cb1565b9150506112a3565b5060008051602061301b833981519152547f8fe9ded8d7c08f720cf0340699024f83522ea66b2bbfb8f557851cb9ee63b54c60005260406020527f027b1db1b6870da5f1d9eb809b54d688d0629fdcef6070a1226b229dbc899792546113dd916001600160a01b031690611d37565b50507fdcf69331608ae4837870755f6bd10a16da39d47dbece9e37a5b80a54219d7ed3600090815260406020527f7cc8b5af461a3e5896300862e4f0c9851aff7b33aa654c13ef6153a737a0e3ad55505050505050565b6049546040516370a0823160e01b81526001600160a01b03838116600483015260009216906370a0823190602401602060405180830381865afa15801561147f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a39190612c1c565b92915050565b600082116114b657600080fd5b69021e19e0c9bab24000008111156115245760405162461bcd60e51b815260206004820152602b60248201527f5469702063616e6e6f742062652067726561746572207468616e20313030302060448201526a2d30b8102a37b5b2b7399760a91b6064820152608401610415565b801561153557611535333083610584565b61153f82826120ee565b6000828152604660209081526040808320600080516020612f79833981519152845260040182529182902054825184815291820152839133917fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820910160405180910390a35050565b6000816001600160a01b031663cd8f942c6040518163ffffffff1660e01b81526004016020604051808303816000875af11580156115e9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a39190612bff565b600083815260466020908152604080832085845260058101909252909120546170809061163a9043612e96565b111561169a5760405162461bcd60e51b815260206004820152602960248201527f4d75737420646973707574652077697468696e203120646179206f662062656960448201526837339036b4b732b21760b91b6064820152608401610415565b60008381526005820160205260409020546116b457600080fd5b600582106116c157600080fd5b336000908152604560205260409020546001146117205760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79207374616b6572732063616e20626567696e206120646973707574656044820152606401610415565b600080516020612ffb8339815191526000526040602081905260008051602061303b8339815191525460495491516370a0823160e01b815233600482015290916001600160a01b0316906370a0823190602401602060405180830381865afa158015611790573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117b49190612c1c565b10156118105760405162461bcd60e51b815260206004820152602560248201527f596f7520646f206e6f74206861766520612062616c616e636520746f20646973604482015264383aba329760d91b6064820152608401610415565b60008381526008820160205260408120836005811061183157611831612e80565b01546040516bffffffffffffffffffffffff19606083901b16602082015260348101879052605481018690526001600160a01b03909116915060009060740160408051601f19818403018152918152815160209283012060008181526048909352912054909150156118a257600080fd5b600080516020612f9983398151915254600080516020612ffb8339815191526000526040602081905260008051602061303b83398151915254905163f3fef3a360e01b81526001600160a01b03909216918291829163f3fef3a39161190c91339190600401612be6565b600060405180830381600087803b15801561192657600080fd5b505af115801561193a573d6000803e3d6000fd5b505060008051602061301b83398151915254600080516020612ffb8339815191526000526040602081905260008051602061303b8339815191525490516311f9fbc960e21b81526001600160a01b0380871695506347e7ef2494506119a493169190600401612be6565b600060405180830381600087803b1580156119be57600080fd5b505af11580156119d2573d6000803e3d6000fd5b505060008051602061301b83398151915254600080516020612ffb833981519152600052604060205260008051602061303b83398151915254611a2493503392506001600160a01b0390911690610584565b7f475da5340e76792184fb177cb85d21980c2530616313aef501564d484eb5ca1e60005260406020527f0eff6fadb497c4730d398e900b416e7234ae208c37f96b69256cb5312e30893954611a7a906001612c4b565b7f0eff6fadb497c4730d398e900b416e7234ae208c37f96b69256cb5312e308939819055600084815260486020908152604080832084905583835260448252808320878155600381018490556004810180546001600160a01b038b166001600160a01b0319918216179091556005808301805483163317905560068301805490921690915560028201805461ffff19169055600182018590557f31b40192effc42bcf1e4289fe674c678e673a3052992548fef566d8c33a21b918552600790910183528184208d90557f4ebf727c48eac2c66272456b06a885c5cc03e54d140f63b63b6fd10c1227958e84528184208c90558b845260098a019092529091209088908110611b8a57611b8a612e80565b015460008281526044602090815260408083207f81afeeaff0ed5cee7d05a21078399c2f56226b0cd5657062500cef4c4e736f858452600701909152902055611bd64262093a80612c4b565b60008281526044602090815260408083207f74c9bc34b0b2333f1b565fbee67d940cf7d78b5a980c5f23da43f6729965ed408452600701909152808220929092557f6f8f54d1af9b6cb8a219d88672c797f9f3ee97ce5d9369aa897fd0deb5e2dffa81528181204390557f8ef61a1efbc527d6428ff88c95fdff5c6e644b979bfe67e03cbf88c8162c5fac815281812089905560008051602061303b833981519152547f833b9f6abf0b529613680afe2a00fa663cc95cbdc47d726d85a044462eabbf028252919020556002871415611cd15760008981526046602090815260408083208b84526007019091529020805460ff191660011790555b6001600160a01b0385166000818152604560209081526040918290206003905581518b8152908101929092528a9183917feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64910160405180910390a3505050505050505050565b604954604051610724916001600160a01b03169063a9059cbb60e01b906105bc9086908690602401612be6565b6040516378bfa27760e01b81526000600482015273__$d4d3c96f4ab0b13728ed5056adc694f4da$__906378bfa2779060240160006040518083038186803b158015611daf57600080fd5b505af4158015611dc3573d6000803e3d6000fd5b5050600080516020612f598339815191526000908152603f602052600080516020612f9983398151915254604051630103c92b60e01b81523360048201526001600160a01b0390911693508392508290630103c92b90602401602060405180830381865afa158015611e39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e5d9190612c1c565b9050611e6a823383610584565b60405163f3fef3a360e01b81526001600160a01b0383169063f3fef3a390611e989033908590600401612be6565b600060405180830381600087803b158015611eb257600080fd5b505af1158015611ec6573d6000803e3d6000fd5b50507fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb287836600090815260406020527f39efff102fba42ea00eb856629a713fead698b2401b1a7fdf4afafc9ef898fd18054859450909250611f27908490612e96565b9091555050505050565b600080516020612f598339815191526000908152603f602052600080516020612f9983398151915254604051630103c92b60e01b81523360048201526001600160a01b039091169182918290630103c92b90602401602060405180830381865afa158015611fa3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc79190612c1c565b60405163f9f6af1760e01b8152600060048201526024810187905285151560448201526064810182905290915073__$f08294d4d1c904fbf7968e6213a9d2a1b1$__9063f9f6af179060840160006040518083038186803b15801561202b57600080fd5b505af4158015610a21573d6000803e3d6000fd5b600061206f82604051806060016040528060228152602001612fd9602291396001600160a01b0386169190612619565b8051909150156105f3578080602001905181019061208d9190612bff565b6105f35760405162461bcd60e51b815260206004820152602c60248201527f5a6170546f6b656e4253433a204552433230206f7065726174696f6e2064696460448201526b081b9bdd081cdd58d8d9595960a21b6064820152608401610415565b60008281526046602052604081209061210681612630565b9050821561215a57600080516020612f798339815191526000908152600483016020526040902054612139908490612c4b565b600080516020612f7983398151915260009081526004840160205260409020555b600080516020612f7983398151915260009081526004830160209081526040808320547f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b890935290527f87c6cfa4d22e28e1e4c87613f3e10333f705784486dec674bed745a01e72d3c4546123a457600080516020612f798339815191526000908152600484016020908152604080832083905590527f87c6cfa4d22e28e1e4c87613f3e10333f705784486dec674bed745a01e72d3c48690557fd26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfc81527fd9aa51e8155c737ee587210adff11a25c4a42c39cf01ee0e08a2dea74f8f3ab2829055548190612269600143612e96565b60408051602081019490945283019190915240606082015260800160408051601f19818403018152828252805160209182012060008181557f87c6cfa4d22e28e1e4c87613f3e10333f705784486dec674bed745a01e72d3c4547fd44e438305e5a7fe9f690b59139512969b5dae7ef40a81d0078a05c66951d77a54818352604685528583207fba3571a50e0c436953d31396edfb65be5925bcc7fef5a3441ed5d43dbce2548f8452600481018652868420547fd26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfc909452959094527fd9aa51e8155c737ee587210adff11a25c4a42c39cf01ee0e08a2dea74f8f3ab25490957f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42956123979593929091612ead565b60405180910390a261254e565b6000828152604660209081526040808320600080516020612f7983398151915284526004019091529020548111806123da575081155b1561242157600283015460405186917f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c916124189187918690612ee4565b60405180910390a25b600080516020612fb983398151915260009081526004840160205260409020546124fd576040805161066081019182905260009182916124819160019060339082845b81548152602001906001019080831161246457505050505061268c565b909250905081831180612492575081155b156124f65782600182603381106124ab576124ab612e80565b015560008181526043602090815260408083208054845260468352818420600080516020612fb9833981519152855260049081018452828520859055908b9055880190915290208190555b505061254e565b831561254e57600080516020612fb9833981519152600090815260048401602052604090205484906001906033811061253857612538612e80565b0160008282546125489190612c4b565b90915550505b5050505050565b600061255f612910565b60008060005b60058110156125b15785603501816005811061258357612583612e80565b600202015484826005811061259a5761259a612e80565b6020020152806125a981612cb1565b915050612565565b505083547fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e60009081526040808701602052808220547fd26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfc835291205491945091509193509193565b606061262884846000856126c3565b949350505050565b60408051610660810191829052600091829182916126719190600187019060339082845b8154815260200190600101908083116126545750505050506127eb565b60009081526043909501602052505060409092205492915050565b610640810151603260315b80156126bd5760208102840151838110156126b3578093508192505b5060001901612697565b50915091565b6060824710156127245760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610415565b843b6127725760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610415565b600080866001600160a01b0316858760405161278e9190612f09565b60006040518083038185875af1925050503d80600081146127cb576040519150601f19603f3d011682016040523d82523d6000602084013e6127d0565b606091505b50915091506127e0828286612819565b979650505050505050565b60008060005b60338110156126bd576020810284015180841015612810578093508192505b506001016127f1565b60608315612828575081610308565b8251156128385782518084602001fd5b8160405162461bcd60e51b81526004016104159190612f25565b82805461285e90612ccc565b90600052602060002090601f01602090048101928261288057600085556128c6565b82601f1061289957805160ff19168380011785556128c6565b828001600101855582156128c6579182015b828111156128c65782518255916020019190600101906128ab565b506128d292915061292e565b5090565b8280548282559060005260206000209081019282156128c657916020028201828111156128c65782518255916020019190600101906128ab565b6040518060a001604052806005906020820280368337509192915050565b5b808211156128d2576000815560010161292f565b6001600160a01b038116811461295857600080fd5b50565b6000806040838503121561296e57600080fd5b823561297981612943565b946020939093013593505050565b60008060006060848603121561299c57600080fd5b83356129a781612943565b925060208401356129b781612943565b929592945050506040919091013590565b60008083601f8401126129da57600080fd5b50813567ffffffffffffffff8111156129f257600080fd5b602083019150836020828501011115612a0a57600080fd5b9250929050565b60008060008060008060808789031215612a2a57600080fd5b863567ffffffffffffffff80821115612a4257600080fd5b612a4e8a838b016129c8565b90985096506020890135915080821115612a6757600080fd5b50612a7489828a016129c8565b979a9699509760408101359660609091013595509350505050565b848152610100810160208083018660005b6005811015612abd57815183529183019190830190600101612aa0565b5050505060c082019390935260e0015292915050565b600060208284031215612ae557600080fd5b5035919050565b60008060008060608587031215612b0257600080fd5b843567ffffffffffffffff811115612b1957600080fd5b612b25878288016129c8565b90989097506020870135966040013595509350505050565b600060208284031215612b4f57600080fd5b813561030881612943565b60008060408385031215612b6d57600080fd5b50508035926020909101359150565b600080600060608486031215612b9157600080fd5b505081359360208301359350604090920135919050565b801515811461295857600080fd5b60008060408385031215612bc957600080fd5b823591506020830135612bdb81612ba8565b809150509250929050565b6001600160a01b03929092168252602082015260400190565b600060208284031215612c1157600080fd5b815161030881612ba8565b600060208284031215612c2e57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115612c5e57612c5e612c35565b500190565b60005b83811015612c7e578181015183820152602001612c66565b838111156106575750506000910152565b60008351612ca1818460208801612c63565b9190910191825250602001919050565b6000600019821415612cc557612cc5612c35565b5060010190565b600181811c90821680612ce057607f821691505b60208210811415612d0157634e487b7160e01b600052602260045260246000fd5b50919050565b8054600090600181811c9080831680612d2157607f831692505b6020808410821415612d4357634e487b7160e01b600052602260045260246000fd5b83885260208801828015612d5e5760018114612d6f57612d9a565b60ff19871682528282019750612d9a565b60008981526020902060005b87811015612d9457815484820152908601908401612d7b565b83019850505b5050505050505092915050565b608081526000612dba6080830187612d07565b8281036020840152612dcc8187612d07565b604084019590955250506060015292915050565b600080600060608486031215612df557600080fd5b8351612e0081612943565b6020850151909350612e1181612943565b80925050604084015190509250925092565b634e487b7160e01b600052600160045260246000fd5b85815260806020820152836080820152838560a0830137600060a08583010152600060a0601f19601f87011683010190508360408301528260608301529695505050505050565b634e487b7160e01b600052603260045260246000fd5b600082821015612ea857612ea8612c35565b500390565b85815284602082015283604082015260a060608201526000612ed260a0830185612d07565b90508260808301529695505050505050565b606081526000612ef76060830186612d07565b60208301949094525060400152919050565b60008251612f1b818460208701612c63565b9190910192915050565b6020815260008251806020840152612f44816040850160208701612c63565b601f01601f1916919091016040019291505056fe973059b1f52a2bdab7ecb9e4cd36d6cb999848b14684091a5057353deb173ed72a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c38b10842dd64626a76de2c61c929077485e8093b2266c7d122beb0f0edee22ff81e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea5a6170546f6b656e4253433a206c6f772d6c6576656c2063616c6c206661696c65648b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc4923977fe60635f602ff3815e40e1a4ed9026273a2eaf4e65856aaf7dfc8a2f0da6995065e9fbf4cade1e55358c20e9d1f06f94dd7926d857be43e1deff4109dcb3a0a26469706673582212207d432506ca3281c45c52aaca065bc19c646ab20c6b44bf17e22f76c94caecac564736f6c634300080b0033",
}

// ZapABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapMetaData.ABI instead.
var ZapABI = ZapMetaData.ABI

// Deprecated: Use ZapMetaData.Sigs instead.
// ZapFuncSigs maps the 4-byte function signature to its string representation.
var ZapFuncSigs = ZapMetaData.Sigs

// ZapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapMetaData.Bin instead.
var ZapBin = ZapMetaData.Bin

// DeployZap deploys a new Ethereum contract, binding an instance of Zap to it.
func DeployZap(auth *bind.TransactOpts, backend bind.ContractBackend, zapTokenBsc common.Address) (common.Address, *types.Transaction, *Zap, error) {
	parsed, err := ZapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	zapLibraryAddr, _, _, _ := DeployZapLibrary(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$7f5e830ec774d87c329c85d735d7e0ce58$__", zapLibraryAddr.String()[2:], -1)

	zapStakeAddr, _, _, _ := DeployZapStake(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$d4d3c96f4ab0b13728ed5056adc694f4da$__", zapStakeAddr.String()[2:], -1)

	zapDisputeAddr, _, _, _ := DeployZapDispute(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$f08294d4d1c904fbf7968e6213a9d2a1b1$__", zapDisputeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapBin), backend, zapTokenBsc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zap{ZapCaller: ZapCaller{contract: contract}, ZapTransactor: ZapTransactor{contract: contract}, ZapFilterer: ZapFilterer{contract: contract}}, nil
}

// Zap is an auto generated Go binding around an Ethereum contract.
type Zap struct {
	ZapCaller     // Read-only binding to the contract
	ZapTransactor // Write-only binding to the contract
	ZapFilterer   // Log filterer for contract events
}

// ZapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapSession struct {
	Contract     *Zap              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapCallerSession struct {
	Contract *ZapCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTransactorSession struct {
	Contract     *ZapTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapRaw struct {
	Contract *Zap // Generic contract binding to access the raw methods on
}

// ZapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapCallerRaw struct {
	Contract *ZapCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTransactorRaw struct {
	Contract *ZapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZap creates a new instance of Zap, bound to a specific deployed contract.
func NewZap(address common.Address, backend bind.ContractBackend) (*Zap, error) {
	contract, err := bindZap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zap{ZapCaller: ZapCaller{contract: contract}, ZapTransactor: ZapTransactor{contract: contract}, ZapFilterer: ZapFilterer{contract: contract}}, nil
}

// NewZapCaller creates a new read-only instance of Zap, bound to a specific deployed contract.
func NewZapCaller(address common.Address, caller bind.ContractCaller) (*ZapCaller, error) {
	contract, err := bindZap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapCaller{contract: contract}, nil
}

// NewZapTransactor creates a new write-only instance of Zap, bound to a specific deployed contract.
func NewZapTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTransactor, error) {
	contract, err := bindZap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransactor{contract: contract}, nil
}

// NewZapFilterer creates a new log filterer instance of Zap, bound to a specific deployed contract.
func NewZapFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapFilterer, error) {
	contract, err := bindZap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapFilterer{contract: contract}, nil
}

// bindZap binds a generic wrapper to an already deployed contract.
func bindZap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zap *ZapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zap.Contract.ZapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zap *ZapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.Contract.ZapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zap *ZapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zap.Contract.ZapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zap *ZapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zap *ZapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zap *ZapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zap.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 balance)
func (_Zap *ZapCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 balance)
func (_Zap *ZapSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Zap.Contract.BalanceOf(&_Zap.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 balance)
func (_Zap *ZapCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Zap.Contract.BalanceOf(&_Zap.CallOpts, _user)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_Zap *ZapCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficulty *big.Int
		Tip        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenge = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RequestIds = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.Difficulty = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tip = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_Zap *ZapSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	return _Zap.Contract.GetNewCurrentVariables(&_Zap.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_Zap *ZapCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	return _Zap.Contract.GetNewCurrentVariables(&_Zap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zap *ZapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zap *ZapSession) Owner() (common.Address, error) {
	return _Zap.Contract.Owner(&_Zap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zap *ZapCallerSession) Owner() (common.Address, error) {
	return _Zap.Contract.Owner(&_Zap.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Zap *ZapCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Zap *ZapSession) Token() (common.Address, error) {
	return _Zap.Contract.Token(&_Zap.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Zap *ZapCallerSession) Token() (common.Address, error) {
	return _Zap.Contract.Token(&_Zap.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Zap *ZapTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Zap *ZapSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.AddTip(&_Zap.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Zap *ZapTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.AddTip(&_Zap.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Zap *ZapTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Zap *ZapSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Approve(&_Zap.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Zap *ZapTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Approve(&_Zap.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Zap *ZapTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Zap *ZapSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.BeginDispute(&_Zap.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Zap *ZapTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.BeginDispute(&_Zap.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapSession) DepositStake() (*types.Transaction, error) {
	return _Zap.Contract.DepositStake(&_Zap.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Zap.Contract.DepositStake(&_Zap.TransactOpts)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0x7cf26110.
//
// Solidity: function increaseVaultApproval(address vaultAddress) returns(bool)
func (_Zap *ZapTransactor) IncreaseVaultApproval(opts *bind.TransactOpts, vaultAddress common.Address) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "increaseVaultApproval", vaultAddress)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0x7cf26110.
//
// Solidity: function increaseVaultApproval(address vaultAddress) returns(bool)
func (_Zap *ZapSession) IncreaseVaultApproval(vaultAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.IncreaseVaultApproval(&_Zap.TransactOpts, vaultAddress)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0x7cf26110.
//
// Solidity: function increaseVaultApproval(address vaultAddress) returns(bool)
func (_Zap *ZapTransactorSession) IncreaseVaultApproval(vaultAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.IncreaseVaultApproval(&_Zap.TransactOpts, vaultAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x2bd326af.
//
// Solidity: function proposeFork(address _propNewZapAddress, uint256 forkedContract) returns()
func (_Zap *ZapTransactor) ProposeFork(opts *bind.TransactOpts, _propNewZapAddress common.Address, forkedContract *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "proposeFork", _propNewZapAddress, forkedContract)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x2bd326af.
//
// Solidity: function proposeFork(address _propNewZapAddress, uint256 forkedContract) returns()
func (_Zap *ZapSession) ProposeFork(_propNewZapAddress common.Address, forkedContract *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress, forkedContract)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x2bd326af.
//
// Solidity: function proposeFork(address _propNewZapAddress, uint256 forkedContract) returns()
func (_Zap *ZapTransactorSession) ProposeFork(_propNewZapAddress common.Address, forkedContract *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress, forkedContract)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Zap *ZapTransactor) RequestData(opts *bind.TransactOpts, _c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "requestData", _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Zap *ZapSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.RequestData(&_Zap.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Zap *ZapTransactorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.RequestData(&_Zap.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Zap.Contract.RequestStakingWithdraw(&_Zap.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Zap.Contract.RequestStakingWithdraw(&_Zap.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Zap *ZapTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Zap *ZapSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.SubmitMiningSolution(&_Zap.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Zap *ZapTransactorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.SubmitMiningSolution(&_Zap.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Zap *ZapTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Zap *ZapSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TallyVotes(&_Zap.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Zap *ZapTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TallyVotes(&_Zap.TransactOpts, _disputeId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_Zap *ZapTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_Zap *ZapSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns()
func (_Zap *ZapTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns()
func (_Zap *ZapTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns()
func (_Zap *ZapSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns()
func (_Zap *ZapTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Zap *ZapTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Zap *ZapSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.Contract.Vote(&_Zap.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Zap *ZapTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.Contract.Vote(&_Zap.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapSession) WithdrawStake() (*types.Transaction, error) {
	return _Zap.Contract.WithdrawStake(&_Zap.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Zap.Contract.WithdrawStake(&_Zap.TransactOpts)
}

// ZapApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Zap contract.
type ZapApprovalIterator struct {
	Event *ZapApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapApproval represents a Approval event raised by the Zap contract.
type ZapApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Zap *ZapFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ZapApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZapApprovalIterator{contract: _Zap.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Zap *ZapFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZapApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapApproval)
				if err := _Zap.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Zap *ZapFilterer) ParseApproval(log types.Log) (*ZapApproval, error) {
	event := new(ZapApproval)
	if err := _Zap.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDataRequestedIterator is returned from FilterDataRequested and is used to iterate over the raw logs and unpacked data for DataRequested events raised by the Zap contract.
type ZapDataRequestedIterator struct {
	Event *ZapDataRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapDataRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDataRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapDataRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapDataRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDataRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDataRequested represents a DataRequested event raised by the Zap contract.
type ZapDataRequested struct {
	Sender      common.Address
	Query       string
	QuerySymbol string
	Granularity *big.Int
	RequestId   *big.Int
	TotalTips   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRequested is a free log retrieval operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_Zap *ZapFilterer) FilterDataRequested(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapDataRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapDataRequestedIterator{contract: _Zap.contract, event: "DataRequested", logs: logs, sub: sub}, nil
}

// WatchDataRequested is a free log subscription operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_Zap *ZapFilterer) WatchDataRequested(opts *bind.WatchOpts, sink chan<- *ZapDataRequested, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDataRequested)
				if err := _Zap.contract.UnpackLog(event, "DataRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDataRequested is a log parse operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_Zap *ZapFilterer) ParseDataRequested(log types.Log) (*ZapDataRequested, error) {
	event := new(ZapDataRequested)
	if err := _Zap.contract.UnpackLog(event, "DataRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the Zap contract.
type ZapNewChallengeIterator struct {
	Event *ZapNewChallenge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapNewChallenge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapNewChallenge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapNewChallenge represents a NewChallenge event raised by the Zap contract.
type ZapNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId *big.Int
	Difficulty       *big.Int
	Multiplier       *big.Int
	Query            string
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_Zap *ZapFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentRequestId []*big.Int) (*ZapNewChallengeIterator, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapNewChallengeIterator{contract: _Zap.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_Zap *ZapFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *ZapNewChallenge, _currentRequestId []*big.Int) (event.Subscription, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapNewChallenge)
				if err := _Zap.contract.UnpackLog(event, "NewChallenge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewChallenge is a log parse operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_Zap *ZapFilterer) ParseNewChallenge(log types.Log) (*ZapNewChallenge, error) {
	event := new(ZapNewChallenge)
	if err := _Zap.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the Zap contract.
type ZapNewDisputeIterator struct {
	Event *ZapNewDispute // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapNewDispute)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapNewDispute)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapNewDispute represents a NewDispute event raised by the Zap contract.
type ZapNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Zap *ZapFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*ZapNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapNewDisputeIterator{contract: _Zap.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Zap *ZapFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *ZapNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapNewDispute)
				if err := _Zap.contract.UnpackLog(event, "NewDispute", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Zap *ZapFilterer) ParseNewDispute(log types.Log) (*ZapNewDispute, error) {
	event := new(ZapNewDispute)
	if err := _Zap.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapNewForkProposalIterator is returned from FilterNewForkProposal and is used to iterate over the raw logs and unpacked data for NewForkProposal events raised by the Zap contract.
type ZapNewForkProposalIterator struct {
	Event *ZapNewForkProposal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapNewForkProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapNewForkProposal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapNewForkProposal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapNewForkProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapNewForkProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapNewForkProposal represents a NewForkProposal event raised by the Zap contract.
type ZapNewForkProposal struct {
	DisputeId        *big.Int
	Timestamp        *big.Int
	ProposedContract common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewForkProposal is a free log retrieval operation binding the contract event 0x6e5021148ce12eaf364c5747e45b0a331a203cc13c84c4d2c371ebd3872d4112.
//
// Solidity: event NewForkProposal(uint256 indexed _disputeId, uint256 _timestamp, address indexed proposedContract)
func (_Zap *ZapFilterer) FilterNewForkProposal(opts *bind.FilterOpts, _disputeId []*big.Int, proposedContract []common.Address) (*ZapNewForkProposalIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}

	var proposedContractRule []interface{}
	for _, proposedContractItem := range proposedContract {
		proposedContractRule = append(proposedContractRule, proposedContractItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "NewForkProposal", _disputeIdRule, proposedContractRule)
	if err != nil {
		return nil, err
	}
	return &ZapNewForkProposalIterator{contract: _Zap.contract, event: "NewForkProposal", logs: logs, sub: sub}, nil
}

// WatchNewForkProposal is a free log subscription operation binding the contract event 0x6e5021148ce12eaf364c5747e45b0a331a203cc13c84c4d2c371ebd3872d4112.
//
// Solidity: event NewForkProposal(uint256 indexed _disputeId, uint256 _timestamp, address indexed proposedContract)
func (_Zap *ZapFilterer) WatchNewForkProposal(opts *bind.WatchOpts, sink chan<- *ZapNewForkProposal, _disputeId []*big.Int, proposedContract []common.Address) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}

	var proposedContractRule []interface{}
	for _, proposedContractItem := range proposedContract {
		proposedContractRule = append(proposedContractRule, proposedContractItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "NewForkProposal", _disputeIdRule, proposedContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapNewForkProposal)
				if err := _Zap.contract.UnpackLog(event, "NewForkProposal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewForkProposal is a log parse operation binding the contract event 0x6e5021148ce12eaf364c5747e45b0a331a203cc13c84c4d2c371ebd3872d4112.
//
// Solidity: event NewForkProposal(uint256 indexed _disputeId, uint256 _timestamp, address indexed proposedContract)
func (_Zap *ZapFilterer) ParseNewForkProposal(log types.Log) (*ZapNewForkProposal, error) {
	event := new(ZapNewForkProposal)
	if err := _Zap.contract.UnpackLog(event, "NewForkProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapNewRequestOnDeckIterator is returned from FilterNewRequestOnDeck and is used to iterate over the raw logs and unpacked data for NewRequestOnDeck events raised by the Zap contract.
type ZapNewRequestOnDeckIterator struct {
	Event *ZapNewRequestOnDeck // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapNewRequestOnDeckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapNewRequestOnDeck)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapNewRequestOnDeck)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapNewRequestOnDeckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapNewRequestOnDeckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapNewRequestOnDeck represents a NewRequestOnDeck event raised by the Zap contract.
type ZapNewRequestOnDeck struct {
	RequestId       *big.Int
	Query           string
	OnDeckQueryHash [32]byte
	OnDeckTotalTips *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequestOnDeck is a free log retrieval operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_Zap *ZapFilterer) FilterNewRequestOnDeck(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapNewRequestOnDeckIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapNewRequestOnDeckIterator{contract: _Zap.contract, event: "NewRequestOnDeck", logs: logs, sub: sub}, nil
}

// WatchNewRequestOnDeck is a free log subscription operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_Zap *ZapFilterer) WatchNewRequestOnDeck(opts *bind.WatchOpts, sink chan<- *ZapNewRequestOnDeck, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapNewRequestOnDeck)
				if err := _Zap.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRequestOnDeck is a log parse operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_Zap *ZapFilterer) ParseNewRequestOnDeck(log types.Log) (*ZapNewRequestOnDeck, error) {
	event := new(ZapNewRequestOnDeck)
	if err := _Zap.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Zap contract.
type ZapOwnershipTransferredIterator struct {
	Event *ZapOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapOwnershipTransferred represents a OwnershipTransferred event raised by the Zap contract.
type ZapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zap *ZapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZapOwnershipTransferredIterator{contract: _Zap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zap *ZapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapOwnershipTransferred)
				if err := _Zap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zap *ZapFilterer) ParseOwnershipTransferred(log types.Log) (*ZapOwnershipTransferred, error) {
	event := new(ZapOwnershipTransferred)
	if err := _Zap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the Zap contract.
type ZapTipAddedIterator struct {
	Event *ZapTipAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTipAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapTipAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTipAdded represents a TipAdded event raised by the Zap contract.
type ZapTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Zap *ZapFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapTipAddedIterator{contract: _Zap.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Zap *ZapFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *ZapTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTipAdded)
				if err := _Zap.contract.UnpackLog(event, "TipAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTipAdded is a log parse operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Zap *ZapFilterer) ParseTipAdded(log types.Log) (*ZapTipAdded, error) {
	event := new(ZapTipAdded)
	if err := _Zap.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeMetaData contains all meta data concerning the ZapDispute contract.
var ZapDisputeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedContract\",\"type\":\"address\"}],\"name\":\"NewForkProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"6cbfb178": "proposeFork(ZapStorage.ZapStorageStruct storage,address,uint256)",
		"3cedafe5": "tallyVotes(ZapStorage.ZapStorageStruct storage,uint256)",
		"9264b888": "updateDisputeFee(ZapStorage.ZapStorageStruct storage)",
		"f9f6af17": "vote(ZapStorage.ZapStorageStruct storage,uint256,bool,uint256)",
	},
	Bin: "0x61104161003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c80633cedafe51461005b5780636cbfb178146100a55780639264b888146100c7578063f9f6af17146100e7575b600080fd5b81801561006757600080fd5b5061007b610076366004610d9e565b610107565b604080516001600160a01b0394851681529390921660208401529082015260600160405180910390f35b8180156100b157600080fd5b506100c56100c0366004610dc0565b610642565b005b8180156100d357600080fd5b506100c56100e2366004610e04565b6108eb565b8180156100f357600080fd5b506100c5610102366004610e1d565b610a74565b600081815260448301602090815260408083207f31b40192effc42bcf1e4289fe674c678e673a3052992548fef566d8c33a21b918452600781018084528285205485526046870184528285207f833b9f6abf0b529613680afe2a00fa663cc95cbdc47d726d85a044462eabbf028652935290832054600282015484938493929091849060ff16156101df5760405162461bcd60e51b815260206004820152601e60248201527f546869732068617320616c7265616479206265656e206578656375746564000060448201526064015b60405180910390fd5b7f74c9bc34b0b2333f1b565fbee67d940cf7d78b5a980c5f23da43f6729965ed406000908152600785016020526040902054421161025f5760405162461bcd60e51b815260206004820152601960248201527f43616e6e6f7420766f746520617420746869732074696d652e0000000000000060448201526064016101d6565b600384015461046f5760048401546001600160a01b0316600090815260458a01602090815260408083207fa0bc13ce85a2091e950a370bced0825e58ab3a3ffeb709ed50d5562cbd82faab845260078801835281842054600080516020610fec8339815191528552828e0190935290832054909291906102e190600290610e77565b6102eb9190610ea4565b9050600086600101541380156103025750600a8111155b156103e7576103146201518042610eb8565b61031e9042610e77565b6001830155600080516020610fec83398151915260009081526040808d01602052812080549161034d83610ecc565b919050555061035b8b6108eb565b60028601805461ff001916610100179055600080516020610fcc8339815191526000908152600780880160209081526040808420548452918801905290205460ff16156103d257600080516020610fcc83398151915260009081526007870160209081526040808320548352600688019091528120555b60058601546001600160a01b03169250610468565b60018255600080516020610fcc8339815191526000908152600787810160209081526040808420548452918801905290205460ff161561045757600080516020610fcc833981519152600090815260078781016020908152604080842054845291880190529020805460ff191690555b60048601546001600160a01b031692505b50506105b6565b6000846001015413156105b6577fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783660009081526040808b016020529020546064906104bb906023610ee3565b6104c59190610ea4565b7f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b2660009081526007860160205260409020541161050157600080fd5b8360030154600114156105625760068401547f710052ea0d12b2397f41c761f87b3558ef80d996883cf1def28dfbcfc47780236000908152603f8b016020526040902080546001600160a01b0319166001600160a01b039092169190911790555b60028401805461ff00191661010017905560068401546040516001600160a01b0390911681527f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab2709060200160405180910390a15b60028401805460ff19166001908117918290556004860154908601546005870154604080519283526001600160a01b03918216602084015261010090940460ff1615159382019390935291169089907f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e104347399060600160405180910390a33096509450925050509250925092565b6040516bffffffffffffffffffffffff19606084901b16602082015260009060340160408051601f198184030181529181528151602092830120600081815260488801909352912054909150156106e55760405162461bcd60e51b815260206004820152602160248201527f446973707574652048617368206973206e6f7420657175616c20746f207a65726044820152606f60f81b60648201526084016101d6565b7f475da5340e76792184fb177cb85d21980c2530616313aef501564d484eb5ca1e60009081526040808601602052812080549161072183610f02565b90915550507f475da5340e76792184fb177cb85d21980c2530616313aef501564d484eb5ca1e6000908152604085810160208181528284205485855260488901825283852081905580855260448901825283852086815560038101889055600481018054336001600160a01b0319918216811790925560058301805482169092179091556006820180549091166001600160a01b038b1617905560028101805461ffff19169055600181018690557f6f8f54d1af9b6cb8a219d88672c797f9f3ee97ce5d9369aa897fd0deb5e2dffa86526007018083528486204390557f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc492398652928252838520547f833b9f6abf0b529613680afe2a00fa663cc95cbdc47d726d85a044462eabbf0286529290915291909220919091556108644262093a80610f1d565b600082815260448701602090815260408083207f74c9bc34b0b2333f1b565fbee67d940cf7d78b5a980c5f23da43f6729965ed40845260070182529182902092909255514281526001600160a01b0386169183917f6e5021148ce12eaf364c5747e45b0a331a203cc13c84c4d2c371ebd3872d411291015b60405180910390a35050505050565b7fabef544d8048318ece54fb2c6385255cd1b06e176525d149a0338a7acca6deb36000908152604080830160205280822054600080516020610fec83398151915283529120546103e891906109409083610ee3565b61094a9190610ea4565b1015610a3f577fabef544d8048318ece54fb2c6385255cd1b06e176525d149a0338a7acca6deb36000908152604080830160205280822054600080516020610fec8339815191528352912054610a0b91600f916103e8916109fc916109af9084610ee3565b6109b99190610ea4565b6109c5906103e8610e77565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b26000908152604080880160205290205490610d51565b610a069190610ea4565b610d88565b7f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc492396000908152604080840160205290205550565b7f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc492396000908152604091820160205220600f9055565b600083815260448501602090815260408083203384526045880190925290912054600114610afd5760405162461bcd60e51b815260206004820152603060248201527f4f6e6c79205374616b657273207468617420617265206e6f7420756e6465722060448201526f646973707574652063616e20766f746560801b60648201526084016101d6565b33600090815260088201602052604090205460ff16151560011415610b645760405162461bcd60e51b815260206004820152601c60248201527f6d73672e73656e6465722068617320616c726561647920766f7465640000000060448201526064016101d6565b60008211610bc75760405162461bcd60e51b815260206004820152602a60248201527f55736572206d757374206861766520612062616c616e63652067726561746572604482015269207468616e207a65726f60b01b60648201526084016101d6565b60058101546001600160a01b0316331415610c3b5760405162461bcd60e51b815260206004820152602e60248201527f546865207265706f7274696e67207061727479206f662074686520646973707560448201526d74652063616e6e6f7420766f746560901b60648201526084016101d6565b3360009081526008820160209081526040808320805460ff191660019081179091557fa0bc13ce85a2091e950a370bced0825e58ab3a3ffeb709ed50d5562cbd82faab8452600785019092528220805491929091610c9a908490610f1d565b90915550507f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b26600090815260078201602052604081208054849290610ce0908490610f1d565b90915550508215610d0557818160010154610cfb9190610f35565b6001820155610d1b565b818160010154610d159190610f76565b60018201555b6040518315158152339085907f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0906020016108dc565b600080610d5e8385610ee3565b9050831580610d75575082610d738583610ea4565b145b610d8157610d81610fb5565b9392505050565b6000818311610d975781610d81565b5090919050565b60008060408385031215610db157600080fd5b50508035926020909101359150565b600080600060608486031215610dd557600080fd5b8335925060208401356001600160a01b0381168114610df357600080fd5b929592945050506040919091013590565b600060208284031215610e1657600080fd5b5035919050565b60008060008060808587031215610e3357600080fd5b843593506020850135925060408501358015158114610e5157600080fd5b9396929550929360600135925050565b634e487b7160e01b600052601160045260246000fd5b600082821015610e8957610e89610e61565b500390565b634e487b7160e01b600052601260045260246000fd5b600082610eb357610eb3610e8e565b500490565b600082610ec757610ec7610e8e565b500690565b600081610edb57610edb610e61565b506000190190565b6000816000190483118215151615610efd57610efd610e61565b500290565b6000600019821415610f1657610f16610e61565b5060010190565b60008219821115610f3057610f30610e61565b500190565b600080821280156001600160ff1b0384900385131615610f5757610f57610e61565b600160ff1b8390038412811615610f7057610f70610e61565b50500190565b60008083128015600160ff1b850184121615610f9457610f94610e61565b6001600160ff1b0384018313811615610faf57610faf610e61565b50500390565b634e487b7160e01b600052600160045260246000fdfe4ebf727c48eac2c66272456b06a885c5cc03e54d140f63b63b6fd10c1227958eedddb9344bfe0dadc78c558b8ffca446679cbffc17be64eb83973fce7bea5f34a26469706673582212203a1acd61655e0a0814f1ea0926a0fbae78f3ef7181c5b0dbd675d899026a31ed64736f6c634300080b0033",
}

// ZapDisputeABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapDisputeMetaData.ABI instead.
var ZapDisputeABI = ZapDisputeMetaData.ABI

// Deprecated: Use ZapDisputeMetaData.Sigs instead.
// ZapDisputeFuncSigs maps the 4-byte function signature to its string representation.
var ZapDisputeFuncSigs = ZapDisputeMetaData.Sigs

// ZapDisputeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapDisputeMetaData.Bin instead.
var ZapDisputeBin = ZapDisputeMetaData.Bin

// DeployZapDispute deploys a new Ethereum contract, binding an instance of ZapDispute to it.
func DeployZapDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapDispute, error) {
	parsed, err := ZapDisputeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapDisputeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapDispute{ZapDisputeCaller: ZapDisputeCaller{contract: contract}, ZapDisputeTransactor: ZapDisputeTransactor{contract: contract}, ZapDisputeFilterer: ZapDisputeFilterer{contract: contract}}, nil
}

// ZapDispute is an auto generated Go binding around an Ethereum contract.
type ZapDispute struct {
	ZapDisputeCaller     // Read-only binding to the contract
	ZapDisputeTransactor // Write-only binding to the contract
	ZapDisputeFilterer   // Log filterer for contract events
}

// ZapDisputeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapDisputeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapDisputeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapDisputeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapDisputeSession struct {
	Contract     *ZapDispute       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapDisputeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapDisputeCallerSession struct {
	Contract *ZapDisputeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapDisputeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapDisputeTransactorSession struct {
	Contract     *ZapDisputeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapDisputeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapDisputeRaw struct {
	Contract *ZapDispute // Generic contract binding to access the raw methods on
}

// ZapDisputeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapDisputeCallerRaw struct {
	Contract *ZapDisputeCaller // Generic read-only contract binding to access the raw methods on
}

// ZapDisputeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapDisputeTransactorRaw struct {
	Contract *ZapDisputeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapDispute creates a new instance of ZapDispute, bound to a specific deployed contract.
func NewZapDispute(address common.Address, backend bind.ContractBackend) (*ZapDispute, error) {
	contract, err := bindZapDispute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapDispute{ZapDisputeCaller: ZapDisputeCaller{contract: contract}, ZapDisputeTransactor: ZapDisputeTransactor{contract: contract}, ZapDisputeFilterer: ZapDisputeFilterer{contract: contract}}, nil
}

// NewZapDisputeCaller creates a new read-only instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeCaller(address common.Address, caller bind.ContractCaller) (*ZapDisputeCaller, error) {
	contract, err := bindZapDispute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeCaller{contract: contract}, nil
}

// NewZapDisputeTransactor creates a new write-only instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapDisputeTransactor, error) {
	contract, err := bindZapDispute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeTransactor{contract: contract}, nil
}

// NewZapDisputeFilterer creates a new log filterer instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapDisputeFilterer, error) {
	contract, err := bindZapDispute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeFilterer{contract: contract}, nil
}

// bindZapDispute binds a generic wrapper to an already deployed contract.
func bindZapDispute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapDisputeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapDispute *ZapDisputeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapDispute.Contract.ZapDisputeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapDispute *ZapDisputeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapDispute.Contract.ZapDisputeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapDispute *ZapDisputeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapDispute.Contract.ZapDisputeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapDispute *ZapDisputeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapDispute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapDispute *ZapDisputeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapDispute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapDispute *ZapDisputeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapDispute.Contract.contract.Transact(opts, method, params...)
}

// ZapDisputeDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the ZapDispute contract.
type ZapDisputeDisputeVoteTalliedIterator struct {
	Event *ZapDisputeDisputeVoteTallied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapDisputeDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeDisputeVoteTallied)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapDisputeDisputeVoteTallied)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapDisputeDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeDisputeVoteTallied represents a DisputeVoteTallied event raised by the ZapDispute contract.
type ZapDisputeDisputeVoteTallied struct {
	DisputeID      *big.Int
	Result         *big.Int
	ReportedMiner  common.Address
	ReportingParty common.Address
	Active         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeVoteTallied is a free log retrieval operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_ZapDispute *ZapDisputeFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*ZapDisputeDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeDisputeVoteTalliedIterator{contract: _ZapDispute.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_ZapDispute *ZapDisputeFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *ZapDisputeDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeDisputeVoteTallied)
				if err := _ZapDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeVoteTallied is a log parse operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_ZapDispute *ZapDisputeFilterer) ParseDisputeVoteTallied(log types.Log) (*ZapDisputeDisputeVoteTallied, error) {
	event := new(ZapDisputeDisputeVoteTallied)
	if err := _ZapDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the ZapDispute contract.
type ZapDisputeNewDisputeIterator struct {
	Event *ZapDisputeNewDispute // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapDisputeNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeNewDispute)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapDisputeNewDispute)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapDisputeNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeNewDispute represents a NewDispute event raised by the ZapDispute contract.
type ZapDisputeNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ZapDispute *ZapDisputeFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*ZapDisputeNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeNewDisputeIterator{contract: _ZapDispute.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ZapDispute *ZapDisputeFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *ZapDisputeNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeNewDispute)
				if err := _ZapDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ZapDispute *ZapDisputeFilterer) ParseNewDispute(log types.Log) (*ZapDisputeNewDispute, error) {
	event := new(ZapDisputeNewDispute)
	if err := _ZapDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeNewForkProposalIterator is returned from FilterNewForkProposal and is used to iterate over the raw logs and unpacked data for NewForkProposal events raised by the ZapDispute contract.
type ZapDisputeNewForkProposalIterator struct {
	Event *ZapDisputeNewForkProposal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapDisputeNewForkProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeNewForkProposal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapDisputeNewForkProposal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapDisputeNewForkProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeNewForkProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeNewForkProposal represents a NewForkProposal event raised by the ZapDispute contract.
type ZapDisputeNewForkProposal struct {
	DisputeId        *big.Int
	Timestamp        *big.Int
	ProposedContract common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewForkProposal is a free log retrieval operation binding the contract event 0x6e5021148ce12eaf364c5747e45b0a331a203cc13c84c4d2c371ebd3872d4112.
//
// Solidity: event NewForkProposal(uint256 indexed _disputeId, uint256 _timestamp, address indexed proposedContract)
func (_ZapDispute *ZapDisputeFilterer) FilterNewForkProposal(opts *bind.FilterOpts, _disputeId []*big.Int, proposedContract []common.Address) (*ZapDisputeNewForkProposalIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}

	var proposedContractRule []interface{}
	for _, proposedContractItem := range proposedContract {
		proposedContractRule = append(proposedContractRule, proposedContractItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "NewForkProposal", _disputeIdRule, proposedContractRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeNewForkProposalIterator{contract: _ZapDispute.contract, event: "NewForkProposal", logs: logs, sub: sub}, nil
}

// WatchNewForkProposal is a free log subscription operation binding the contract event 0x6e5021148ce12eaf364c5747e45b0a331a203cc13c84c4d2c371ebd3872d4112.
//
// Solidity: event NewForkProposal(uint256 indexed _disputeId, uint256 _timestamp, address indexed proposedContract)
func (_ZapDispute *ZapDisputeFilterer) WatchNewForkProposal(opts *bind.WatchOpts, sink chan<- *ZapDisputeNewForkProposal, _disputeId []*big.Int, proposedContract []common.Address) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}

	var proposedContractRule []interface{}
	for _, proposedContractItem := range proposedContract {
		proposedContractRule = append(proposedContractRule, proposedContractItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "NewForkProposal", _disputeIdRule, proposedContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeNewForkProposal)
				if err := _ZapDispute.contract.UnpackLog(event, "NewForkProposal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewForkProposal is a log parse operation binding the contract event 0x6e5021148ce12eaf364c5747e45b0a331a203cc13c84c4d2c371ebd3872d4112.
//
// Solidity: event NewForkProposal(uint256 indexed _disputeId, uint256 _timestamp, address indexed proposedContract)
func (_ZapDispute *ZapDisputeFilterer) ParseNewForkProposal(log types.Log) (*ZapDisputeNewForkProposal, error) {
	event := new(ZapDisputeNewForkProposal)
	if err := _ZapDispute.contract.UnpackLog(event, "NewForkProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapDispute contract.
type ZapDisputeNewZapAddressIterator struct {
	Event *ZapDisputeNewZapAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapDisputeNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeNewZapAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapDisputeNewZapAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapDisputeNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeNewZapAddress represents a NewZapAddress event raised by the ZapDispute contract.
type ZapDisputeNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapDispute *ZapDisputeFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapDisputeNewZapAddressIterator, error) {

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapDisputeNewZapAddressIterator{contract: _ZapDispute.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapDispute *ZapDisputeFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapDisputeNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeNewZapAddress)
				if err := _ZapDispute.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewZapAddress is a log parse operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapDispute *ZapDisputeFilterer) ParseNewZapAddress(log types.Log) (*ZapDisputeNewZapAddress, error) {
	event := new(ZapDisputeNewZapAddress)
	if err := _ZapDispute.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the ZapDispute contract.
type ZapDisputeVotedIterator struct {
	Event *ZapDisputeVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapDisputeVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapDisputeVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapDisputeVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeVoted represents a Voted event raised by the ZapDispute contract.
type ZapDisputeVoted struct {
	DisputeID *big.Int
	Position  bool
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
func (_ZapDispute *ZapDisputeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address) (*ZapDisputeVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeVotedIterator{contract: _ZapDispute.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
func (_ZapDispute *ZapDisputeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ZapDisputeVoted, _disputeID []*big.Int, _voter []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeVoted)
				if err := _ZapDispute.contract.UnpackLog(event, "Voted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoted is a log parse operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
func (_ZapDispute *ZapDisputeFilterer) ParseVoted(log types.Log) (*ZapDisputeVoted, error) {
	event := new(ZapDisputeVoted)
	if err := _ZapDispute.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapGettersMetaData contains all meta data concerning the ZapGetters contract.
var ZapGettersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zapTokenBsc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"70a08231": "balanceOf(address)",
		"63bb82ad": "didMine(bytes32,address)",
		"a7c438bc": "didVote(uint256,address)",
		"133bee5e": "getAddressVars(bytes32)",
		"af0b1327": "getAllDisputeVars(uint256)",
		"a22e407a": "getCurrentVariables()",
		"da379941": "getDisputeIdByDisputeHash(bytes32)",
		"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
		"fc7cf0a0": "getLastNewValue()",
		"3180f8df": "getLastNewValueById(uint256)",
		"c775b542": "getMinedBlockNum(uint256,uint256)",
		"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
		"46eee1c4": "getNewValueCountbyRequestId(uint256)",
		"1db842f0": "getRequestIdByQueryHash(bytes32)",
		"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
		"0f0b424d": "getRequestIdByTimestamp(uint256)",
		"b5413029": "getRequestQ()",
		"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
		"e1eee6d6": "getRequestVars(uint256)",
		"733bdef0": "getStakerInfo(address)",
		"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
		"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
		"612c8f7f": "getUintVar(bytes32)",
		"19e8e03b": "getVariablesOnDeck()",
		"3df0777b": "isInDispute(uint256,uint256)",
		"93fa4915": "retrieveData(uint256,uint256)",
		"1ca8b6cb": "totalTokenSupply()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161168b38038061168b83398101604081905261002f916100a6565b604980546001600160a01b039092166001600160a01b031992831681179091557f0dbd11667cb3ff39667584f924913ce4dcfc6917743490eff0a0ac17f29b151e600052603f6020527fc5d5df8489ed162ad95aee74a257293c0a4237df2a31e83b0d8cbd05fa533a4880549092161790556100d6565b6000602082840312156100b857600080fd5b81516001600160a01b03811681146100cf57600080fd5b9392505050565b6115a6806100e56000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c8063733bdef0116100f9578063b541302911610097578063dd62ed3e11610071578063dd62ed3e1461046f578063e0ae93c114610482578063e1eee6d614610495578063fc7cf0a0146104ba57600080fd5b8063b541302914610434578063c775b54214610449578063da3799411461045c57600080fd5b806393fa4915116100d357806393fa4915146103cc578063a22e407a146103df578063a7c438bc146103f9578063af0b13271461040c57600080fd5b8063733bdef01461037e57806377fbb663146103a65780637f6fd5d9146103b957600080fd5b80633df0777b116101665780636173c0b8116101405780636173c0b81461032557806363bb82ad1461033857806369026d631461034b57806370a082311461036b57600080fd5b80633df0777b146102dc57806346eee1c4146102ff578063612c8f7f1461031257600080fd5b806319e8e03b116101a257806319e8e03b1461023a5780631ca8b6cb146102515780631db842f0146102a15780633180f8df146102b457600080fd5b80630f0b424d146101c957806311c98512146101ef578063133bee5e1461020f575b600080fd5b6101dc6101d7366004611213565b61050a565b6040519081526020015b60405180910390f35b6102026101fd36600461122c565b610520565b6040516101e6919061124e565b61022261021d366004611213565b61053b565b6040516001600160a01b0390911681526020016101e6565b610242610558565b6040516101e6939291906112cc565b7fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783660005260406020527f39efff102fba42ea00eb856629a713fead698b2401b1a7fdf4afafc9ef898fd1546101dc565b6101dc6102af366004611213565b610577565b6102c76102c2366004611213565b61058b565b604080519283529015156020830152016101e6565b6102ef6102ea36600461122c565b6105a1565b60405190151581526020016101e6565b6101dc61030d366004611213565b6105c6565b6101dc610320366004611213565b6105dd565b6101dc610333366004611213565b6105f1565b6102ef610346366004611310565b6105fd565b61035e61035936600461122c565b610628565b6040516101e6919061133c565b6101dc61037936600461136d565b61063c565b61039161038c36600461136d565b6106ab565b604080519283526020830191909152016101e6565b6101dc6103b436600461122c565b6106d3565b6101dc6103c736600461122c565b6106e0565b6101dc6103da36600461122c565b610702565b6103e7610724565b6040516101e696959493929190611388565b6102ef610407366004611310565b61074b565b61041f61041a366004611213565b610779565b6040516101e6999897969594939291906113c2565b61043c6107b7565b6040516101e69190611443565b6101dc61045736600461122c565b6107c9565b6101dc61046a366004611213565b6107eb565b6101dc61047d36600461146c565b6107ff565b6101dc61049036600461122c565b610876565b6104a86104a3366004611213565b610898565b6040516101e696959493929190611496565b7f56ed6ca285f512c2643e5f72926a108bde86cf97e1be4aabb5c6bf871b3fd62a5460008181526042602090815260408083205483526046825280832093835260069093019052205460016102c7565b6000818152604260205260408120545b92915050565b6105286111b7565b610534600084846108be565b9392505050565b6000818152603f60205260408120546001600160a01b031661051a565b60008060606105676000610918565b925092509250909192565b905090565b60008181526047602052604081205461051a565b6000806105988184610a0f565b91509150915091565b600082815260466020908152604080832084845260070190915281205460ff16610534565b60008181526046602052604081206003015461051a565b60008181526040602081905281205461051a565b600061051a8183610aa2565b60008281526041602090815260408083206001600160a01b038516845290915281205460ff16610534565b6106306111b7565b61053460008484610ac9565b6049546040516370a0823160e01b81526001600160a01b03838116600483015260009216906370a0823190602401602060405180830381865afa158015610687573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051a91906114e1565b6001600160a01b03811660009081526045602052604081208054600190910154829190610598565b6000610534818484610b2d565b6000828152604460209081526040808320848452600701909152812054610534565b6000828152604660209081526040808320848452600601909152812054610534565b600080600060606000806107386000610b66565b949b939a50919850965094509092509050565b60008281526044602090815260408083206001600160a01b038516845260080190915281205460ff16610534565b600080600080600080600061078c6111d5565b6000610798818b610cda565b9850985098509850985098509850985098509193959799909294969850565b6107bf6111f4565b6105726000610f09565b6000828152604660209081526040808320848452600501909152812054610534565b60008181526048602052604081205461051a565b604954604051636eb1769f60e11b81526001600160a01b0384811660048301528381166024830152600092169063dd62ed3e90604401602060405180830381865afa158015610852573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053491906114e1565b6000828152604660209081526040808320848452600401909152812054610534565b60608060008080806108aa8188610f49565b949c939b5091995097509550909350915050565b6108c66111b7565b6000838152604685016020908152604080832085845260090190915290819020815160a08101928390529160059082845b8154815260200190600101908083116108f757505050505090509392505050565b6000806060600061092885611127565b600081815260468701602081815260408084207f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c385526004810183529084205493859052919052805492935083928190610981906114fa565b80601f01602080910402602001604051908101604052809291908181526020018280546109ad906114fa565b80156109fa5780601f106109cf576101008083540402835291602001916109fa565b820191906000526020600020905b8154815290600101906020018083116109dd57829003601f168201915b50505050509050935093509350509193909250565b60008181526046830160205260408120600381015482919015610a9257600381018054610a86918791879190610a4790600190611535565b81548110610a5757610a5761155a565b906000526020600020015460009182526046929092016020908152604080832093835260069093019052205490565b60019250925050610a9b565b60008092509250505b9250929050565b60006032821115610ab257600080fd5b506000908152604391909101602052604090205490565b610ad16111b7565b6000838152604685016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b03168152600190910190602001808311610b0257505050505090509392505050565b60008281526046840160205260408120600301805483908110610b5257610b5261155a565b906000526020600020015490509392505050565b80547f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b8600090815260408084016020908152818320547fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e8452828420548185526046870183528385207fba3571a50e0c436953d31396edfb65be5925bcc7fef5a3441ed5d43dbce2548f865260048101808552858720547f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c3885294529385205484549596879687966060968896879694959094938390610c45906114fa565b80601f0160208091040260200160405190810160405280929190818152602001828054610c71906114fa565b8015610cbe5780601f10610c9357610100808354040283529160200191610cbe565b820191906000526020600020905b815481529060010190602001808311610ca157829003601f168201915b5050505050925095509550955095509550955091939550919395565b6000806000806000806000610ced6111d5565b505050600095865250505050604492909201602090815260408083208054600282015460048301546005840154600685015460038601548751610120810189527f31b40192effc42bcf1e4289fe674c678e673a3052992548fef566d8c33a21b918b5260078801808b52898c205482527f4ebf727c48eac2c66272456b06a885c5cc03e54d140f63b63b6fd10c1227958e8c52808b52898c2054828c01527f81afeeaff0ed5cee7d05a21078399c2f56226b0cd5657062500cef4c4e736f858c52808b52898c2054828b01527f74c9bc34b0b2333f1b565fbee67d940cf7d78b5a980c5f23da43f6729965ed408c52808b52898c205460608301527fa0bc13ce85a2091e950a370bced0825e58ab3a3ffeb709ed50d5562cbd82faab8c52808b52898c205460808301527f6f8f54d1af9b6cb8a219d88672c797f9f3ee97ce5d9369aa897fd0deb5e2dffa8c52808b52898c205460a08301527f8ef61a1efbc527d6428ff88c95fdff5c6e644b979bfe67e03cbf88c8162c5fac8c52808b52898c205460c08301527f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b268c52808b52898c205460e08301527f833b9f6abf0b529613680afe2a00fa663cc95cbdc47d726d85a044462eabbf028c52909952969098205461010080890191909152600190950154939960ff8085169a5095909304909416966001600160a01b03918216969482169591169390929091565b610f116111f4565b6040805161066081019182905290600184019060339082845b815481526020019060010190808311610f2a5750505050509050919050565b6000818152604683016020908152604080832060028101547fba3571a50e0c436953d31396edfb65be5925bcc7fef5a3441ed5d43dbce2548f855260048201909352818420547f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea8552828520547f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c3865292852054825460609687969095869586958695919485946001860194909390918690611004906114fa565b80601f0160208091040260200160405190810160405280929190818152602001828054611030906114fa565b801561107d5780601f106110525761010080835404028352916020019161107d565b820191906000526020600020905b81548152906001019060200180831161106057829003601f168201915b50505050509550848054611090906114fa565b80601f01602080910402602001604051908101604052809291908181526020018280546110bc906114fa565b80156111095780601f106110de57610100808354040283529160200191611109565b820191906000526020600020905b8154815290600101906020018083116110ec57829003601f168201915b50505050509450965096509650965096509650509295509295509295565b60408051610660810191829052600091829182916111689190600187019060339082845b81548152602001906001019080831161114b575050505050611183565b60009081526043909501602052505060409092205492915050565b60008060005b60338110156111b15760208102840151808410156111a8578093508192505b50600101611189565b50915091565b6040518060a001604052806005906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b6040518061066001604052806033906020820280368337509192915050565b60006020828403121561122557600080fd5b5035919050565b6000806040838503121561123f57600080fd5b50508035926020909101359150565b60a08101818360005b6005811015611276578151835260209283019290910190600101611257565b50505092915050565b6000815180845260005b818110156112a557602081850181015186830182015201611289565b818111156112b7576000602083870101525b50601f01601f19169290920160200192915050565b8381528260208201526060604082015260006112eb606083018461127f565b95945050505050565b80356001600160a01b038116811461130b57600080fd5b919050565b6000806040838503121561132357600080fd5b82359150611333602084016112f4565b90509250929050565b60a08101818360005b60058110156112765781516001600160a01b0316835260209283019290910190600101611345565b60006020828403121561137f57600080fd5b610534826112f4565b86815285602082015284604082015260c0606082015260006113ad60c083018661127f565b60808301949094525060a00152949350505050565b89815288151560208083019190915288151560408301526001600160a01b0388811660608401528781166080840152861660a083015260c0820185905261022082019060e083018560005b600981101561142a5781518352918301919083019060010161140d565b50505050826102008301529a9950505050505050505050565b6106608101818360005b603381101561127657815183526020928301929091019060010161144d565b6000806040838503121561147f57600080fd5b611488836112f4565b9150611333602084016112f4565b60c0815260006114a960c083018961127f565b82810360208401526114bb818961127f565b9150508560408301528460608301528360808301528260a0830152979650505050505050565b6000602082840312156114f357600080fd5b5051919050565b600181811c9082168061150e57607f821691505b6020821081141561152f57634e487b7160e01b600052602260045260246000fd5b50919050565b60008282101561155557634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220a1eb7766b6cdbea8d6ee496cf300dc30cad16fe16852ed65233b563d5d25473e64736f6c634300080b0033",
}

// ZapGettersABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapGettersMetaData.ABI instead.
var ZapGettersABI = ZapGettersMetaData.ABI

// Deprecated: Use ZapGettersMetaData.Sigs instead.
// ZapGettersFuncSigs maps the 4-byte function signature to its string representation.
var ZapGettersFuncSigs = ZapGettersMetaData.Sigs

// ZapGettersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapGettersMetaData.Bin instead.
var ZapGettersBin = ZapGettersMetaData.Bin

// DeployZapGetters deploys a new Ethereum contract, binding an instance of ZapGetters to it.
func DeployZapGetters(auth *bind.TransactOpts, backend bind.ContractBackend, zapTokenBsc common.Address) (common.Address, *types.Transaction, *ZapGetters, error) {
	parsed, err := ZapGettersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapGettersBin), backend, zapTokenBsc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapGetters{ZapGettersCaller: ZapGettersCaller{contract: contract}, ZapGettersTransactor: ZapGettersTransactor{contract: contract}, ZapGettersFilterer: ZapGettersFilterer{contract: contract}}, nil
}

// ZapGetters is an auto generated Go binding around an Ethereum contract.
type ZapGetters struct {
	ZapGettersCaller     // Read-only binding to the contract
	ZapGettersTransactor // Write-only binding to the contract
	ZapGettersFilterer   // Log filterer for contract events
}

// ZapGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapGettersSession struct {
	Contract     *ZapGetters       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapGettersCallerSession struct {
	Contract *ZapGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapGettersTransactorSession struct {
	Contract     *ZapGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapGettersRaw struct {
	Contract *ZapGetters // Generic contract binding to access the raw methods on
}

// ZapGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapGettersCallerRaw struct {
	Contract *ZapGettersCaller // Generic read-only contract binding to access the raw methods on
}

// ZapGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapGettersTransactorRaw struct {
	Contract *ZapGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapGetters creates a new instance of ZapGetters, bound to a specific deployed contract.
func NewZapGetters(address common.Address, backend bind.ContractBackend) (*ZapGetters, error) {
	contract, err := bindZapGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapGetters{ZapGettersCaller: ZapGettersCaller{contract: contract}, ZapGettersTransactor: ZapGettersTransactor{contract: contract}, ZapGettersFilterer: ZapGettersFilterer{contract: contract}}, nil
}

// NewZapGettersCaller creates a new read-only instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersCaller(address common.Address, caller bind.ContractCaller) (*ZapGettersCaller, error) {
	contract, err := bindZapGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersCaller{contract: contract}, nil
}

// NewZapGettersTransactor creates a new write-only instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapGettersTransactor, error) {
	contract, err := bindZapGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersTransactor{contract: contract}, nil
}

// NewZapGettersFilterer creates a new log filterer instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapGettersFilterer, error) {
	contract, err := bindZapGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapGettersFilterer{contract: contract}, nil
}

// bindZapGetters binds a generic wrapper to an already deployed contract.
func bindZapGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGetters *ZapGettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGetters.Contract.ZapGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGetters *ZapGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGetters.Contract.ZapGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGetters *ZapGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGetters.Contract.ZapGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGetters *ZapGettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGetters *ZapGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGetters *ZapGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGetters.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapGetters *ZapGettersSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.Allowance(&_ZapGetters.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.Allowance(&_ZapGetters.CallOpts, _user, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapGetters *ZapGettersSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOf(&_ZapGetters.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOf(&_ZapGetters.CallOpts, _user)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapGetters *ZapGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapGetters *ZapGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapGetters.Contract.DidMine(&_ZapGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapGetters *ZapGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapGetters.Contract.DidMine(&_ZapGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapGetters *ZapGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapGetters *ZapGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapGetters.Contract.DidVote(&_ZapGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapGetters *ZapGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapGetters.Contract.DidVote(&_ZapGetters.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapGetters *ZapGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapGetters *ZapGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapGetters.Contract.GetAddressVars(&_ZapGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapGetters *ZapGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapGetters.Contract.GetAddressVars(&_ZapGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, address, address, address, uint256, uint256[9], int256)
func (_ZapGetters *ZapGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, common.Address, common.Address, common.Address, *big.Int, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new(*big.Int), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, address, address, address, uint256, uint256[9], int256)
func (_ZapGetters *ZapGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, common.Address, common.Address, common.Address, *big.Int, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, address, address, address, uint256, uint256[9], int256)
func (_ZapGetters *ZapGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, common.Address, common.Address, common.Address, *big.Int, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getCurrentVariables")

	if err != nil {
		return *new([32]byte), *new(*big.Int), *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetCurrentVariables(&_ZapGetters.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetCurrentVariables(&_ZapGetters.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeIdByDisputeHash(&_ZapGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeIdByDisputeHash(&_ZapGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeUintVars(&_ZapGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeUintVars(&_ZapGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapGetters *ZapGettersCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getLastNewValue")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapGetters *ZapGettersSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValue(&_ZapGetters.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapGetters *ZapGettersCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValue(&_ZapGetters.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapGetters *ZapGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getLastNewValueById", _requestId)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapGetters *ZapGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValueById(&_ZapGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapGetters *ZapGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValueById(&_ZapGetters.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetMinedBlockNum(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetMinedBlockNum(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapGetters *ZapGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapGetters *ZapGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapGetters.Contract.GetMinersByRequestIdAndTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapGetters *ZapGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapGetters.Contract.GetMinersByRequestIdAndTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetNewValueCountbyRequestId(&_ZapGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetNewValueCountbyRequestId(&_ZapGetters.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestIdByQueryHash", _request)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByQueryHash(&_ZapGetters.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByQueryHash(&_ZapGetters.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByRequestQIndex(&_ZapGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByRequestQIndex(&_ZapGetters.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByTimestamp(&_ZapGetters.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByTimestamp(&_ZapGetters.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapGetters *ZapGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapGetters *ZapGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapGetters.Contract.GetRequestQ(&_ZapGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapGetters *ZapGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapGetters.Contract.GetRequestQ(&_ZapGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestUintVars(&_ZapGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestUintVars(&_ZapGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(string), *new(string), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetRequestVars(&_ZapGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetRequestVars(&_ZapGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getStakerInfo", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetStakerInfo(&_ZapGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetStakerInfo(&_ZapGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapGetters *ZapGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapGetters *ZapGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapGetters.Contract.GetSubmissionsByTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapGetters *ZapGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapGetters.Contract.GetSubmissionsByTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetTimestampbyRequestIDandIndex(&_ZapGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetTimestampbyRequestIDandIndex(&_ZapGetters.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetUintVar(&_ZapGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetUintVar(&_ZapGetters.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getVariablesOnDeck")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapGetters.Contract.GetVariablesOnDeck(&_ZapGetters.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapGetters.Contract.GetVariablesOnDeck(&_ZapGetters.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapGetters *ZapGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapGetters *ZapGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapGetters.Contract.IsInDispute(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapGetters *ZapGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapGetters.Contract.IsInDispute(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.RetrieveData(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.RetrieveData(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapGetters *ZapGettersCaller) TotalTokenSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "totalTokenSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapGetters *ZapGettersSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapGetters.Contract.TotalTokenSupply(&_ZapGetters.CallOpts)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapGetters.Contract.TotalTokenSupply(&_ZapGetters.CallOpts)
}

// ZapGettersLibraryMetaData contains all meta data concerning the ZapGettersLibrary contract.
var ZapGettersLibraryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122050bd55867d8c14d726718582a6120e3841578ae3340920fdcd90cc666350bdfa64736f6c634300080b0033",
}

// ZapGettersLibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapGettersLibraryMetaData.ABI instead.
var ZapGettersLibraryABI = ZapGettersLibraryMetaData.ABI

// ZapGettersLibraryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapGettersLibraryMetaData.Bin instead.
var ZapGettersLibraryBin = ZapGettersLibraryMetaData.Bin

// DeployZapGettersLibrary deploys a new Ethereum contract, binding an instance of ZapGettersLibrary to it.
func DeployZapGettersLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapGettersLibrary, error) {
	parsed, err := ZapGettersLibraryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapGettersLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapGettersLibrary{ZapGettersLibraryCaller: ZapGettersLibraryCaller{contract: contract}, ZapGettersLibraryTransactor: ZapGettersLibraryTransactor{contract: contract}, ZapGettersLibraryFilterer: ZapGettersLibraryFilterer{contract: contract}}, nil
}

// ZapGettersLibrary is an auto generated Go binding around an Ethereum contract.
type ZapGettersLibrary struct {
	ZapGettersLibraryCaller     // Read-only binding to the contract
	ZapGettersLibraryTransactor // Write-only binding to the contract
	ZapGettersLibraryFilterer   // Log filterer for contract events
}

// ZapGettersLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapGettersLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapGettersLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapGettersLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapGettersLibrarySession struct {
	Contract     *ZapGettersLibrary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZapGettersLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapGettersLibraryCallerSession struct {
	Contract *ZapGettersLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZapGettersLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapGettersLibraryTransactorSession struct {
	Contract     *ZapGettersLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZapGettersLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapGettersLibraryRaw struct {
	Contract *ZapGettersLibrary // Generic contract binding to access the raw methods on
}

// ZapGettersLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapGettersLibraryCallerRaw struct {
	Contract *ZapGettersLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// ZapGettersLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapGettersLibraryTransactorRaw struct {
	Contract *ZapGettersLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapGettersLibrary creates a new instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibrary(address common.Address, backend bind.ContractBackend) (*ZapGettersLibrary, error) {
	contract, err := bindZapGettersLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibrary{ZapGettersLibraryCaller: ZapGettersLibraryCaller{contract: contract}, ZapGettersLibraryTransactor: ZapGettersLibraryTransactor{contract: contract}, ZapGettersLibraryFilterer: ZapGettersLibraryFilterer{contract: contract}}, nil
}

// NewZapGettersLibraryCaller creates a new read-only instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryCaller(address common.Address, caller bind.ContractCaller) (*ZapGettersLibraryCaller, error) {
	contract, err := bindZapGettersLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryCaller{contract: contract}, nil
}

// NewZapGettersLibraryTransactor creates a new write-only instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapGettersLibraryTransactor, error) {
	contract, err := bindZapGettersLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryTransactor{contract: contract}, nil
}

// NewZapGettersLibraryFilterer creates a new log filterer instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapGettersLibraryFilterer, error) {
	contract, err := bindZapGettersLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryFilterer{contract: contract}, nil
}

// bindZapGettersLibrary binds a generic wrapper to an already deployed contract.
func bindZapGettersLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGettersLibrary *ZapGettersLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGettersLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGettersLibrary *ZapGettersLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGettersLibrary *ZapGettersLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.contract.Transact(opts, method, params...)
}

// ZapGettersLibraryNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapGettersLibrary contract.
type ZapGettersLibraryNewZapAddressIterator struct {
	Event *ZapGettersLibraryNewZapAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapGettersLibraryNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapGettersLibraryNewZapAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapGettersLibraryNewZapAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapGettersLibraryNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapGettersLibraryNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapGettersLibraryNewZapAddress represents a NewZapAddress event raised by the ZapGettersLibrary contract.
type ZapGettersLibraryNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapGettersLibraryNewZapAddressIterator, error) {

	logs, sub, err := _ZapGettersLibrary.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryNewZapAddressIterator{contract: _ZapGettersLibrary.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapGettersLibraryNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapGettersLibrary.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapGettersLibraryNewZapAddress)
				if err := _ZapGettersLibrary.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewZapAddress is a log parse operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) ParseNewZapAddress(log types.Log) (*ZapGettersLibraryNewZapAddress, error) {
	event := new(ZapGettersLibraryNewZapAddress)
	if err := _ZapGettersLibrary.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryMetaData contains all meta data concerning the ZapLibrary contract.
var ZapLibraryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"d723552b": "submitMiningSolution(ZapStorage.ZapStorageStruct storage,string,uint256,uint256)",
	},
	Bin: "0x61185561003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063d723552b1461003a575b600080fd5b81801561004657600080fd5b5061005a610055366004611272565b61005c565b005b3360009081526045850160205260409020546001146100b85760405162461bcd60e51b8152602060048201526013602482015272135a5b995c881a5cc81b9bdd081cdd185ad959606a1b60448201526064015b60405180910390fd5b6000805160206117c083398151915260009081526040808601602052902054821461014b5760405162461bcd60e51b815260206004820152603860248201527f54686520736f6c7574696f6e207375626d6974746564206973206e6f7420666f60448201527f72207468652063757272656e742072657175657374204944000000000000000060648201526084016100af565b836040016000600080516020611780833981519152815260200190815260200160002054600260038660000154338760405160200161018c93929190611369565b60408051601f198184030181528282528051602091820120908301520160408051601f19818403018152908290526101c3916113a8565b602060405180830381855afa1580156101e0573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff19166020820152603401905060408051601f198184030181529082905261021d916113a8565b602060405180830381855afa15801561023a573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061025d91906113c4565b61026791906113f3565b156102b45760405162461bcd60e51b815260206004820152601c60248201527f4368616c6c656e676520696e666f206973206e6f7420756e697175650000000060448201526064016100af565b83546000908152604185016020908152604080832033845290915290205460ff161561032e5760405162461bcd60e51b815260206004820152602360248201527f4d696e65722068617320616c7265616479207375626d697474656420612076616044820152626c756560e81b60648201526084016100af565b7fdcf69331608ae4837870755f6bd10a16da39d47dbece9e37a5b80a54219d7ed3600090815260408086016020528082208290556000805160206117a08339815191528252902054819060358601906005811061038d5761038d611407565b60020201556000805160206117a08339815191526000908152604080860160205290205433906035860190600581106103c8576103c8611407565b6002020160010180546001600160a01b0319166001600160a01b03929092169190911790556000805160206117a083398151915260009081526040858101602052812080549161041783611433565b909155505083546000908152604185016020908152604080832033808552925291829020805460ff191660011790558554915184927fe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc49161047c91889187919061144e565b60405180910390a36000805160206117a083398151915260009081526040808601602052902054600514156104b6576104b68484846104bc565b50505050565b6000818152604684016020908152604080832060008051602061176083398151915284528187019092528220549091906064906104f99042611490565b7fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc336000908152604080890160205290205461053491906114a7565b6000805160206117808339815191526000908152604080890160205290205461055d91906114e6565b60008051602061178083398151915260009081526040808901602052902054610586919061156b565b61059091906115ac565b9050600081136105c157600080516020611780833981519152600090815260408087016020529020600190556105e3565b6000805160206117808339815191526000908152604080870160205290208190555b6105ee603c426113f3565b6105f89042611490565b6000805160206117608339815191526000908152604080880160205280822092909255815160a081019092529060358701600583835b8282101561066f5760408051808201909152600283028501805482526001908101546001600160a01b0316602080840191909152918352909201910161062e565b509293506001925050505b60058110156107f757600082826005811061069757610697611407565b602002015151905060008383600581106106b3576106b3611407565b602002015160200151905060008390505b6000811180156106f45750846106db600183611490565b600581106106eb576106eb611407565b60200201515183105b156107945784610705600183611490565b6005811061071557610715611407565b60200201515185826005811061072d5761072d611407565b6020020151528461073f600183611490565b6005811061074f5761074f611407565b60200201516020015185826005811061076a5761076a611407565b602090810291909101516001600160a01b039092169101528061078c816115da565b9150506106c4565b838110156107e157828582600581106107af576107af611407565b602002015152818582600581106107c8576107c8611407565b602090810291909101516001600160a01b039092169101525b50505080806107ef90611433565b91505061067a565b6000805160206117e083398151915260009081526040808901602052902054610844576000805160206117e08339815191526000908152604080890160205290206753444835ec58000090555b6000805160206117e083398151915260009081526040808901602052902054670de0b6b3a7640000101561094d576000805160206117e083398151915260009081526040808901602052902054670de0b6b3a7640000906108ab90651bd78f205bc66115f1565b6108b59190611610565b6000805160206117e083398151915260009081526040808a016020529020546108de9190611490565b6000805160206117e083398151915260009081526040808a01602052902081905560649061090d9060326115f1565b6109179190611610565b7f8fe9ded8d7c08f720cf0340699024f83522ea66b2bbfb8f557851cb9ee63b54c60009081526040808a01602052902055610977565b6000805160206117e0833981519152600090815260408089016020529020670de0b6b3a764000090555b6000805160206117e0833981519152600090815260408089016020528120546109a990670de0b6b3a764000090611610565b6109bb90670de0b6b3a76400006115f1565b60008051602061180083398151915260009081526040808b016020529020549091506109e990600590611610565b6109f39082611624565b7fdcf69331608ae4837870755f6bd10a16da39d47dbece9e37a5b80a54219d7ed3600090815260408a81016020528082209290925560008051602061176083398151915281528181205485830151516000805160206118008339815191528352929091205488927fc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd169291610a89906005906113f3565b60008051602061180083398151915260009081526040808f01602052902054610ab29190611490565b8c5460408051948552602085019390935291830152606082015260800160405180910390a27fdcf69331608ae4837870755f6bd10a16da39d47dbece9e37a5b80a54219d7ed360009081526040808a01602052902054610b139060056115f1565b7f8fe9ded8d7c08f720cf0340699024f83522ea66b2bbfb8f557851cb9ee63b54c60009081526040808b01602052902054610b4e9190611624565b7fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783660009081526040808b0160205281208054909190610b8e908490611624565b9091555050604080840180515160008051602061176083398151915260009081528b840160208181528583208054845260068c01825286842094909455908152825460038b0180546001810182559084528284200155845160a08101865288518201516001600160a01b039081168252828a015183015181168284015294518201518516818701526060808a01518301518616908201526080808a015183015190951694810194909452915481526008890190915291909120610c529160056111c1565b506040805160a081018252845151815260208086015151818301528583015151828401526060808701515190830152608080870151519083015260008051602061176083398151915260009081528b840182528381205481526009890190915291909120610cc1916005611219565b5060008051602061176083398151915260009081526040808a0160208181528284208054855260058a0182528385204390558054855260428d0182528385208b9055918152905460348c018054600181018255908552918420909101556000805160206117a08339815191528252812055610d3b88611131565b6000805160206117c083398151915260009081526040808b01602052902081905515611103576000805160206117c083398151915260009081526040808a0160208181528284208054855260468d018083528486207f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c3875260049081018452858720546000805160206118008339815191528852948452858720949094559054855281528284207f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea8552909101905281205460018a019060338110610e2257610e22611407565b01556000805160206117c083398151915260009081526040808a0160209081528183208054845260468c018083528385207f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea8087526004918201855285872054875260438f01855285872087905583548752828552858720908752810184528486208690559154855282528284207f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c385520190528120819055610ee489611131565b89549091508890610ef6600143611490565b40604051602001610f099392919061163c565b6040516020818303038152906040528051906020012089600001819055508860400160006000805160206117c08339815191528152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f428a600001548b60400160006000805160206117808339815191528152602001908152602001600020548c60460160008e60400160006000805160206117c0833981519152815260200190815260200160002054815260200190815260200160002060040160007fba3571a50e0c436953d31396edfb65be5925bcc7fef5a3441ed5d43dbce2548f8152602001908152602001600020548d60460160008f60400160006000805160206117c083398151915281526020019081526020016000205481526020019081526020016000206000018e6040016000600080516020611800833981519152815260200190815260200160002054604051611072959493929190611703565b60405180910390a2600081815260468a016020908152604080832060028101547f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c38552600482019093529281902054905184937f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c936110f593919290919061173a565b60405180910390a250611127565b60008051602061180083398151915260009081526040808a01602052812081905588555b5050505050505050565b60408051610660810191829052600091829182916111729190600187019060339082845b81548152602001906001019080831161115557505050505061118d565b60009081526043909501602052505060409092205492915050565b60008060005b60338110156111bb5760208102840151808410156111b2578093508192505b50600101611193565b50915091565b8260058101928215611209579160200282015b8281111561120957825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906111d4565b50611215929150611247565b5090565b8260058101928215611209579160200282015b8281111561120957825182559160200191906001019061122c565b5b808211156112155760008155600101611248565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561128857600080fd5b84359350602085013567ffffffffffffffff808211156112a757600080fd5b818701915087601f8301126112bb57600080fd5b8135818111156112cd576112cd61125c565b604051601f8201601f19908116603f011681019083821181831017156112f5576112f561125c565b816040528281528a602084870101111561130e57600080fd5b826020860160208301376000928101602001929092525095989597505050506040840135936060013592915050565b60005b83811015611358578181015183820152602001611340565b838111156104b65750506000910152565b8381526bffffffffffffffffffffffff198360601b1660208201526000825161139981603485016020870161133d565b91909101603401949350505050565b600082516113ba81846020870161133d565b9190910192915050565b6000602082840312156113d657600080fd5b5051919050565b634e487b7160e01b600052601260045260246000fd5b600082611402576114026113dd565b500690565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114475761144761141d565b5060010190565b606081526000845180606084015261146d81608085016020890161133d565b60208301949094525060408101919091526080601f909201601f19160101919050565b6000828210156114a2576114a261141d565b500390565b60008083128015600160ff1b8501841216156114c5576114c561141d565b6001600160ff1b03840183138116156114e0576114e061141d565b50500390565b60006001600160ff1b038184138284138082168684048611161561150c5761150c61141d565b600160ff1b600087128281168783058912161561152b5761152b61141d565b600087129250878205871284841616156115475761154761141d565b8785058712818416161561155d5761155d61141d565b505050929093029392505050565b600080821280156001600160ff1b038490038513161561158d5761158d61141d565b600160ff1b83900384128116156115a6576115a661141d565b50500190565b6000826115bb576115bb6113dd565b600160ff1b8214600019841416156115d5576115d561141d565b500590565b6000816115e9576115e961141d565b506000190190565b600081600019048311821515161561160b5761160b61141d565b500290565b60008261161f5761161f6113dd565b500490565b600082198211156116375761163761141d565b500190565b6000845161164e81846020890161133d565b91909101928352506020820152604001919050565b8054600090600181811c908083168061167d57607f831692505b602080841082141561169f57634e487b7160e01b600052602260045260246000fd5b838852602088018280156116ba57600181146116cb576116f6565b60ff198716825282820197506116f6565b60008981526020902060005b878110156116f0578154848201529086019084016116d7565b83019850505b5050505050505092915050565b85815284602082015283604082015260a06060820152600061172860a0830185611663565b90508260808301529695505050505050565b60608152600061174d6060830186611663565b6020830194909452506040015291905056fe97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a4b12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e6c505cb2db6644f57b42d87bd9407b0f66788b07d0617a2bc1356a0e69e66f9a7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b89b6853911475b07474368644a0d922ee13bc76a15cd3e97d3e334326424a47d4d26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfca26469706673582212209b006e24ad2790176c4b9fe26cefcc93e5ea3b7b1db14acb703d2f924c7219cf64736f6c634300080b0033",
}

// ZapLibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapLibraryMetaData.ABI instead.
var ZapLibraryABI = ZapLibraryMetaData.ABI

// Deprecated: Use ZapLibraryMetaData.Sigs instead.
// ZapLibraryFuncSigs maps the 4-byte function signature to its string representation.
var ZapLibraryFuncSigs = ZapLibraryMetaData.Sigs

// ZapLibraryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapLibraryMetaData.Bin instead.
var ZapLibraryBin = ZapLibraryMetaData.Bin

// DeployZapLibrary deploys a new Ethereum contract, binding an instance of ZapLibrary to it.
func DeployZapLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapLibrary, error) {
	parsed, err := ZapLibraryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapLibrary{ZapLibraryCaller: ZapLibraryCaller{contract: contract}, ZapLibraryTransactor: ZapLibraryTransactor{contract: contract}, ZapLibraryFilterer: ZapLibraryFilterer{contract: contract}}, nil
}

// ZapLibrary is an auto generated Go binding around an Ethereum contract.
type ZapLibrary struct {
	ZapLibraryCaller     // Read-only binding to the contract
	ZapLibraryTransactor // Write-only binding to the contract
	ZapLibraryFilterer   // Log filterer for contract events
}

// ZapLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapLibrarySession struct {
	Contract     *ZapLibrary       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapLibraryCallerSession struct {
	Contract *ZapLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapLibraryTransactorSession struct {
	Contract     *ZapLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapLibraryRaw struct {
	Contract *ZapLibrary // Generic contract binding to access the raw methods on
}

// ZapLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapLibraryCallerRaw struct {
	Contract *ZapLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// ZapLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapLibraryTransactorRaw struct {
	Contract *ZapLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapLibrary creates a new instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibrary(address common.Address, backend bind.ContractBackend) (*ZapLibrary, error) {
	contract, err := bindZapLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapLibrary{ZapLibraryCaller: ZapLibraryCaller{contract: contract}, ZapLibraryTransactor: ZapLibraryTransactor{contract: contract}, ZapLibraryFilterer: ZapLibraryFilterer{contract: contract}}, nil
}

// NewZapLibraryCaller creates a new read-only instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryCaller(address common.Address, caller bind.ContractCaller) (*ZapLibraryCaller, error) {
	contract, err := bindZapLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryCaller{contract: contract}, nil
}

// NewZapLibraryTransactor creates a new write-only instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapLibraryTransactor, error) {
	contract, err := bindZapLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryTransactor{contract: contract}, nil
}

// NewZapLibraryFilterer creates a new log filterer instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapLibraryFilterer, error) {
	contract, err := bindZapLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryFilterer{contract: contract}, nil
}

// bindZapLibrary binds a generic wrapper to an already deployed contract.
func bindZapLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapLibrary *ZapLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapLibrary.Contract.ZapLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapLibrary *ZapLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapLibrary.Contract.ZapLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapLibrary *ZapLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapLibrary.Contract.ZapLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapLibrary *ZapLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapLibrary *ZapLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapLibrary *ZapLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapLibrary.Contract.contract.Transact(opts, method, params...)
}

// ZapLibraryDataRequestedIterator is returned from FilterDataRequested and is used to iterate over the raw logs and unpacked data for DataRequested events raised by the ZapLibrary contract.
type ZapLibraryDataRequestedIterator struct {
	Event *ZapLibraryDataRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapLibraryDataRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryDataRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapLibraryDataRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapLibraryDataRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryDataRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryDataRequested represents a DataRequested event raised by the ZapLibrary contract.
type ZapLibraryDataRequested struct {
	Sender      common.Address
	Query       string
	QuerySymbol string
	Granularity *big.Int
	RequestId   *big.Int
	TotalTips   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRequested is a free log retrieval operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterDataRequested(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapLibraryDataRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryDataRequestedIterator{contract: _ZapLibrary.contract, event: "DataRequested", logs: logs, sub: sub}, nil
}

// WatchDataRequested is a free log subscription operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchDataRequested(opts *bind.WatchOpts, sink chan<- *ZapLibraryDataRequested, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryDataRequested)
				if err := _ZapLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDataRequested is a log parse operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseDataRequested(log types.Log) (*ZapLibraryDataRequested, error) {
	event := new(ZapLibraryDataRequested)
	if err := _ZapLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the ZapLibrary contract.
type ZapLibraryNewChallengeIterator struct {
	Event *ZapLibraryNewChallenge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapLibraryNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewChallenge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapLibraryNewChallenge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapLibraryNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewChallenge represents a NewChallenge event raised by the ZapLibrary contract.
type ZapLibraryNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId *big.Int
	Difficulty       *big.Int
	Multiplier       *big.Int
	Query            string
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentRequestId []*big.Int) (*ZapLibraryNewChallengeIterator, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewChallengeIterator{contract: _ZapLibrary.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewChallenge, _currentRequestId []*big.Int) (event.Subscription, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewChallenge)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewChallenge is a log parse operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseNewChallenge(log types.Log) (*ZapLibraryNewChallenge, error) {
	event := new(ZapLibraryNewChallenge)
	if err := _ZapLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNewRequestOnDeckIterator is returned from FilterNewRequestOnDeck and is used to iterate over the raw logs and unpacked data for NewRequestOnDeck events raised by the ZapLibrary contract.
type ZapLibraryNewRequestOnDeckIterator struct {
	Event *ZapLibraryNewRequestOnDeck // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapLibraryNewRequestOnDeckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewRequestOnDeck)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapLibraryNewRequestOnDeck)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapLibraryNewRequestOnDeckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewRequestOnDeckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewRequestOnDeck represents a NewRequestOnDeck event raised by the ZapLibrary contract.
type ZapLibraryNewRequestOnDeck struct {
	RequestId       *big.Int
	Query           string
	OnDeckQueryHash [32]byte
	OnDeckTotalTips *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequestOnDeck is a free log retrieval operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewRequestOnDeck(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapLibraryNewRequestOnDeckIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewRequestOnDeckIterator{contract: _ZapLibrary.contract, event: "NewRequestOnDeck", logs: logs, sub: sub}, nil
}

// WatchNewRequestOnDeck is a free log subscription operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewRequestOnDeck(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewRequestOnDeck, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewRequestOnDeck)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRequestOnDeck is a log parse operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseNewRequestOnDeck(log types.Log) (*ZapLibraryNewRequestOnDeck, error) {
	event := new(ZapLibraryNewRequestOnDeck)
	if err := _ZapLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the ZapLibrary contract.
type ZapLibraryNewValueIterator struct {
	Event *ZapLibraryNewValue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapLibraryNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewValue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapLibraryNewValue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapLibraryNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewValue represents a NewValue event raised by the ZapLibrary contract.
type ZapLibraryNewValue struct {
	RequestId        *big.Int
	Time             *big.Int
	Value            *big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewValue(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapLibraryNewValueIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewValueIterator{contract: _ZapLibrary.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewValue, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewValue)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewValue is a log parse operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) ParseNewValue(log types.Log) (*ZapLibraryNewValue, error) {
	event := new(ZapLibraryNewValue)
	if err := _ZapLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the ZapLibrary contract.
type ZapLibraryNonceSubmittedIterator struct {
	Event *ZapLibraryNonceSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapLibraryNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNonceSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapLibraryNonceSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapLibraryNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNonceSubmitted represents a NonceSubmitted event raised by the ZapLibrary contract.
type ZapLibraryNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        *big.Int
	Value            *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _requestId []*big.Int) (*ZapLibraryNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNonceSubmittedIterator{contract: _ZapLibrary.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *ZapLibraryNonceSubmitted, _miner []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNonceSubmitted)
				if err := _ZapLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNonceSubmitted is a log parse operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) ParseNonceSubmitted(log types.Log) (*ZapLibraryNonceSubmitted, error) {
	event := new(ZapLibraryNonceSubmitted)
	if err := _ZapLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZapLibrary contract.
type ZapLibraryOwnershipTransferredIterator struct {
	Event *ZapLibraryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapLibraryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapLibraryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapLibraryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryOwnershipTransferred represents a OwnershipTransferred event raised by the ZapLibrary contract.
type ZapLibraryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ZapLibrary *ZapLibraryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*ZapLibraryOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryOwnershipTransferredIterator{contract: _ZapLibrary.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ZapLibrary *ZapLibraryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZapLibraryOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryOwnershipTransferred)
				if err := _ZapLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ZapLibrary *ZapLibraryFilterer) ParseOwnershipTransferred(log types.Log) (*ZapLibraryOwnershipTransferred, error) {
	event := new(ZapLibraryOwnershipTransferred)
	if err := _ZapLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the ZapLibrary contract.
type ZapLibraryTipAddedIterator struct {
	Event *ZapLibraryTipAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapLibraryTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryTipAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapLibraryTipAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapLibraryTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryTipAdded represents a TipAdded event raised by the ZapLibrary contract.
type ZapLibraryTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapLibraryTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryTipAddedIterator{contract: _ZapLibrary.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *ZapLibraryTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryTipAdded)
				if err := _ZapLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTipAdded is a log parse operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseTipAdded(log types.Log) (*ZapLibraryTipAdded, error) {
	event := new(ZapLibraryTipAdded)
	if err := _ZapLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapMasterMetaData contains all meta data concerning the ZapMaster contract.
var ZapMasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultContract\",\"type\":\"address\"}],\"name\":\"changeVaultContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"70a08231": "balanceOf(address)",
		"af2a5102": "changeVaultContract(address)",
		"63bb82ad": "didMine(bytes32,address)",
		"a7c438bc": "didVote(uint256,address)",
		"133bee5e": "getAddressVars(bytes32)",
		"af0b1327": "getAllDisputeVars(uint256)",
		"a22e407a": "getCurrentVariables()",
		"da379941": "getDisputeIdByDisputeHash(bytes32)",
		"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
		"fc7cf0a0": "getLastNewValue()",
		"3180f8df": "getLastNewValueById(uint256)",
		"c775b542": "getMinedBlockNum(uint256,uint256)",
		"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
		"46eee1c4": "getNewValueCountbyRequestId(uint256)",
		"1db842f0": "getRequestIdByQueryHash(bytes32)",
		"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
		"0f0b424d": "getRequestIdByTimestamp(uint256)",
		"b5413029": "getRequestQ()",
		"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
		"e1eee6d6": "getRequestVars(uint256)",
		"733bdef0": "getStakerInfo(address)",
		"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
		"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
		"612c8f7f": "getUintVar(bytes32)",
		"19e8e03b": "getVariablesOnDeck()",
		"3df0777b": "isInDispute(uint256,uint256)",
		"8da5cb5b": "owner()",
		"93fa4915": "retrieveData(uint256,uint256)",
		"1ca8b6cb": "totalTokenSupply()",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162001bfd38038062001bfd833981016040819052620000349162000237565b604980546001600160a01b0383166001600160a01b031991821681179092557f0dbd11667cb3ff39667584f924913ce4dcfc6917743490eff0a0ac17f29b151e6000908152603f6020527fc5d5df8489ed162ad95aee74a257293c0a4237df2a31e83b0d8cbd05fa533a48805490921690921790556040516347b024eb60e01b815273__$d4d3c96f4ab0b13728ed5056adc694f4da$__916347b024eb91620000e4919060040190815260200190565b60006040518083038186803b158015620000fd57600080fd5b505af415801562000112573d6000803e3d6000fd5b5050603f60209081527f77fe60635f602ff3815e40e1a4ed9026273a2eaf4e65856aaf7dfc8a2f0da6998054336001600160a01b031991821681179092557fe82bad50ed21b0b346eda61b3dfbb0532423ce1e0ded398f9e213767e30bf14780548216831790557f710052ea0d12b2397f41c761f87b3558ef80d996883cf1def28dfbcfc47780236000527f335ff78e0d6ba44a6d9321c4133511a8ad65b142fc6b0abee992f17a25a9b10a805482166001600160a01b038916908117909155604a805490921690921790556040519081527f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270935001905060405180910390a150506200026f565b80516001600160a01b03811681146200023257600080fd5b919050565b600080604083850312156200024b57600080fd5b62000256836200021a565b915062000266602084016200021a565b90509250929050565b61197e806200027f6000396000f3fe6080604052600436106101d15760003560e01c806377fbb663116100f7578063af2a510211610095578063dd62ed3e11610064578063dd62ed3e146106e4578063e0ae93c114610704578063e1eee6d614610724578063fc7cf0a01461075657610210565b8063af2a510214610660578063b541302914610682578063c775b542146106a4578063da379941146106c457610210565b806393fa4915116100d157806393fa4915146105c4578063a22e407a146105e4578063a7c438bc1461060b578063af0b13271461062b57610210565b806377fbb663146105645780637f6fd5d9146105845780638da5cb5b146105a457610210565b80633df0777b1161016f57806363bb82ad1161013e57806363bb82ad146104c257806369026d63146104e257806370a082311461050f578063733bdef01461052f57610210565b80633df0777b1461043257806346eee1c414610462578063612c8f7f146104825780636173c0b8146104a257610210565b806319e8e03b116101ab57806319e8e03b1461035c5780631ca8b6cb146103805780631db842f0146103dd5780633180f8df146103fd57610210565b80630f0b424d146102c457806311c98512146102f7578063133bee5e1461032457610210565b3661021057604080513381523460208201527f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874910160405180910390a1005b7f710052ea0d12b2397f41c761f87b3558ef80d996883cf1def28dfbcfc47780236000908152603f60209081527f335ff78e0d6ba44a6d9321c4133511a8ad65b142fc6b0abee992f17a25a9b10a546040805136601f81018590048502820185019092528181526001600160a01b03909216939283919081908401838280828437600092018290525084519495509384935091505060208401855af43d604051816000823e8280156102c0578282f35b8282fd5b3480156102d057600080fd5b506102e46102df3660046115eb565b6107b3565b6040519081526020015b60405180910390f35b34801561030357600080fd5b50610317610312366004611604565b6107c9565b6040516102ee9190611626565b34801561033057600080fd5b5061034461033f3660046115eb565b6107e4565b6040516001600160a01b0390911681526020016102ee565b34801561036857600080fd5b50610371610801565b6040516102ee939291906116a4565b34801561038c57600080fd5b507fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783660005260406020527f39efff102fba42ea00eb856629a713fead698b2401b1a7fdf4afafc9ef898fd1546102e4565b3480156103e957600080fd5b506102e46103f83660046115eb565b610820565b34801561040957600080fd5b5061041d6104183660046115eb565b610834565b604080519283529015156020830152016102ee565b34801561043e57600080fd5b5061045261044d366004611604565b61084a565b60405190151581526020016102ee565b34801561046e57600080fd5b506102e461047d3660046115eb565b61086f565b34801561048e57600080fd5b506102e461049d3660046115eb565b610886565b3480156104ae57600080fd5b506102e46104bd3660046115eb565b61089a565b3480156104ce57600080fd5b506104526104dd3660046116e8565b6108a6565b3480156104ee57600080fd5b506105026104fd366004611604565b6108d1565b6040516102ee9190611714565b34801561051b57600080fd5b506102e461052a366004611745565b6108e5565b34801561053b57600080fd5b5061054f61054a366004611745565b610954565b604080519283526020830191909152016102ee565b34801561057057600080fd5b506102e461057f366004611604565b61097c565b34801561059057600080fd5b506102e461059f366004611604565b610989565b3480156105b057600080fd5b50604a54610344906001600160a01b031681565b3480156105d057600080fd5b506102e46105df366004611604565b6109ab565b3480156105f057600080fd5b506105f96109cd565b6040516102ee96959493929190611760565b34801561061757600080fd5b506104526106263660046116e8565b6109f4565b34801561063757600080fd5b5061064b6106463660046115eb565b610a22565b6040516102ee9998979695949392919061179a565b34801561066c57600080fd5b5061068061067b366004611745565b610a60565b005b34801561068e57600080fd5b50610697610af6565b6040516102ee919061181b565b3480156106b057600080fd5b506102e46106bf366004611604565b610b08565b3480156106d057600080fd5b506102e46106df3660046115eb565b610b2a565b3480156106f057600080fd5b506102e46106ff366004611844565b610b3e565b34801561071057600080fd5b506102e461071f366004611604565b610bb5565b34801561073057600080fd5b5061074461073f3660046115eb565b610bd7565b6040516102ee9695949392919061186e565b34801561076257600080fd5b507f56ed6ca285f512c2643e5f72926a108bde86cf97e1be4aabb5c6bf871b3fd62a54600081815260426020908152604080832054835260468252808320938352600690930190522054600161041d565b6000818152604260205260408120545b92915050565b6107d161158f565b6107dd60008484610bfd565b9392505050565b6000818152603f60205260408120546001600160a01b03166107c3565b60008060606108106000610c57565b925092509250909192565b905090565b6000818152604760205260408120546107c3565b6000806108418184610d4e565b91509150915091565b600082815260466020908152604080832084845260070190915281205460ff166107dd565b6000818152604660205260408120600301546107c3565b6000818152604060208190528120546107c3565b60006107c38183610de1565b60008281526041602090815260408083206001600160a01b038516845290915281205460ff166107dd565b6108d961158f565b6107dd60008484610e08565b6049546040516370a0823160e01b81526001600160a01b03838116600483015260009216906370a0823190602401602060405180830381865afa158015610930573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c391906118b9565b6001600160a01b03811660009081526045602052604081208054600190910154829190610841565b60006107dd818484610e6c565b60008281526044602090815260408083208484526007019091528120546107dd565b60008281526046602090815260408083208484526006019091528120546107dd565b600080600060606000806109e16000610ea5565b949b939a50919850965094509092509050565b60008281526044602090815260408083206001600160a01b038516845260080190915281205460ff166107dd565b6000806000806000806000610a356115ad565b6000610a41818b611019565b9850985098509850985098509850985098509193959799909294969850565b604a546001600160a01b03163314610abe5760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722063616e207472616e736665722062616c616e63652e604482015260640160405180910390fd5b604a54600160a01b900460ff1615610ad557600080fd5b604a805460ff60a01b1916600160a01b179055610af3600082611248565b50565b610afe6115cc565b61081b60006112e1565b60008281526046602090815260408083208484526005019091528120546107dd565b6000818152604860205260408120546107c3565b604954604051636eb1769f60e11b81526001600160a01b0384811660048301528381166024830152600092169063dd62ed3e90604401602060405180830381865afa158015610b91573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107dd91906118b9565b60008281526046602090815260408083208484526004019091528120546107dd565b6060806000808080610be98188611321565b949c939b5091995097509550909350915050565b610c0561158f565b6000838152604685016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610c3657505050505090509392505050565b60008060606000610c67856114ff565b600081815260468701602081815260408084207f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c385526004810183529084205493859052919052805492935083928190610cc0906118d2565b80601f0160208091040260200160405190810160405280929190818152602001828054610cec906118d2565b8015610d395780601f10610d0e57610100808354040283529160200191610d39565b820191906000526020600020905b815481529060010190602001808311610d1c57829003601f168201915b50505050509050935093509350509193909250565b60008181526046830160205260408120600381015482919015610dd157600381018054610dc5918791879190610d869060019061190d565b81548110610d9657610d96611932565b906000526020600020015460009182526046929092016020908152604080832093835260069093019052205490565b60019250925050610dda565b60008092509250505b9250929050565b60006032821115610df157600080fd5b506000908152604391909101602052604090205490565b610e1061158f565b6000838152604685016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b03168152600190910190602001808311610e4157505050505090509392505050565b60008281526046840160205260408120600301805483908110610e9157610e91611932565b906000526020600020015490509392505050565b80547f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b8600090815260408084016020908152818320547fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e8452828420548185526046870183528385207fba3571a50e0c436953d31396edfb65be5925bcc7fef5a3441ed5d43dbce2548f865260048101808552858720547f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c3885294529385205484549596879687966060968896879694959094938390610f84906118d2565b80601f0160208091040260200160405190810160405280929190818152602001828054610fb0906118d2565b8015610ffd5780601f10610fd257610100808354040283529160200191610ffd565b820191906000526020600020905b815481529060010190602001808311610fe057829003601f168201915b5050505050925095509550955095509550955091939550919395565b600080600080600080600061102c6115ad565b505050600095865250505050604492909201602090815260408083208054600282015460048301546005840154600685015460038601548751610120810189527f31b40192effc42bcf1e4289fe674c678e673a3052992548fef566d8c33a21b918b5260078801808b52898c205482527f4ebf727c48eac2c66272456b06a885c5cc03e54d140f63b63b6fd10c1227958e8c52808b52898c2054828c01527f81afeeaff0ed5cee7d05a21078399c2f56226b0cd5657062500cef4c4e736f858c52808b52898c2054828b01527f74c9bc34b0b2333f1b565fbee67d940cf7d78b5a980c5f23da43f6729965ed408c52808b52898c205460608301527fa0bc13ce85a2091e950a370bced0825e58ab3a3ffeb709ed50d5562cbd82faab8c52808b52898c205460808301527f6f8f54d1af9b6cb8a219d88672c797f9f3ee97ce5d9369aa897fd0deb5e2dffa8c52808b52898c205460a08301527f8ef61a1efbc527d6428ff88c95fdff5c6e644b979bfe67e03cbf88c8162c5fac8c52808b52898c205460c08301527f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b268c52808b52898c205460e08301527f833b9f6abf0b529613680afe2a00fa663cc95cbdc47d726d85a044462eabbf028c52909952969098205461010080890191909152600190950154939960ff8085169a5095909304909416966001600160a01b03918216969482169591169390929091565b7f9dbc393ddc18fd27b1d9b1b129059925688d2f2d5818a5ec3ebb750b7c286ea66000908152603f830160205260409020546001600160a01b0316331461128e57600080fd5b7f973059b1f52a2bdab7ecb9e4cd36d6cb999848b14684091a5057353deb173ed76000908152603f92909201602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6112e96115cc565b6040805161066081019182905290600184019060339082845b8154815260200190600101908083116113025750505050509050919050565b6000818152604683016020908152604080832060028101547fba3571a50e0c436953d31396edfb65be5925bcc7fef5a3441ed5d43dbce2548f855260048201909352818420547f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea8552828520547f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c38652928520548254606096879690958695869586959194859460018601949093909186906113dc906118d2565b80601f0160208091040260200160405190810160405280929190818152602001828054611408906118d2565b80156114555780601f1061142a57610100808354040283529160200191611455565b820191906000526020600020905b81548152906001019060200180831161143857829003601f168201915b50505050509550848054611468906118d2565b80601f0160208091040260200160405190810160405280929190818152602001828054611494906118d2565b80156114e15780601f106114b6576101008083540402835291602001916114e1565b820191906000526020600020905b8154815290600101906020018083116114c457829003601f168201915b50505050509450965096509650965096509650509295509295509295565b60408051610660810191829052600091829182916115409190600187019060339082845b81548152602001906001019080831161152357505050505061155b565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611589576020810284015180841015611580578093508192505b50600101611561565b50915091565b6040518060a001604052806005906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b6040518061066001604052806033906020820280368337509192915050565b6000602082840312156115fd57600080fd5b5035919050565b6000806040838503121561161757600080fd5b50508035926020909101359150565b60a08101818360005b600581101561164e57815183526020928301929091019060010161162f565b50505092915050565b6000815180845260005b8181101561167d57602081850181015186830182015201611661565b8181111561168f576000602083870101525b50601f01601f19169290920160200192915050565b8381528260208201526060604082015260006116c36060830184611657565b95945050505050565b80356001600160a01b03811681146116e357600080fd5b919050565b600080604083850312156116fb57600080fd5b8235915061170b602084016116cc565b90509250929050565b60a08101818360005b600581101561164e5781516001600160a01b031683526020928301929091019060010161171d565b60006020828403121561175757600080fd5b6107dd826116cc565b86815285602082015284604082015260c06060820152600061178560c0830186611657565b60808301949094525060a00152949350505050565b89815288151560208083019190915288151560408301526001600160a01b0388811660608401528781166080840152861660a083015260c0820185905261022082019060e083018560005b6009811015611802578151835291830191908301906001016117e5565b50505050826102008301529a9950505050505050505050565b6106608101818360005b603381101561164e578151835260209283019290910190600101611825565b6000806040838503121561185757600080fd5b611860836116cc565b915061170b602084016116cc565b60c08152600061188160c0830189611657565b82810360208401526118938189611657565b9150508560408301528460608301528360808301528260a0830152979650505050505050565b6000602082840312156118cb57600080fd5b5051919050565b600181811c908216806118e657607f821691505b6020821081141561190757634e487b7160e01b600052602260045260246000fd5b50919050565b60008282101561192d57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fdfea26469706673582212206b8bf7a2733be40aa0319ce151733be53f2004a4b8d5084b84472550f21684e964736f6c634300080b0033",
}

// ZapMasterABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapMasterMetaData.ABI instead.
var ZapMasterABI = ZapMasterMetaData.ABI

// Deprecated: Use ZapMasterMetaData.Sigs instead.
// ZapMasterFuncSigs maps the 4-byte function signature to its string representation.
var ZapMasterFuncSigs = ZapMasterMetaData.Sigs

// ZapMasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapMasterMetaData.Bin instead.
var ZapMasterBin = ZapMasterMetaData.Bin

// DeployZapMaster deploys a new Ethereum contract, binding an instance of ZapMaster to it.
func DeployZapMaster(auth *bind.TransactOpts, backend bind.ContractBackend, _zapContract common.Address, tokenAddress common.Address) (common.Address, *types.Transaction, *ZapMaster, error) {
	parsed, err := ZapMasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	zapStakeAddr, _, _, _ := DeployZapStake(auth, backend)
	ZapMasterBin = strings.Replace(ZapMasterBin, "__$d4d3c96f4ab0b13728ed5056adc694f4da$__", zapStakeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapMasterBin), backend, _zapContract, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapMaster{ZapMasterCaller: ZapMasterCaller{contract: contract}, ZapMasterTransactor: ZapMasterTransactor{contract: contract}, ZapMasterFilterer: ZapMasterFilterer{contract: contract}}, nil
}

// ZapMaster is an auto generated Go binding around an Ethereum contract.
type ZapMaster struct {
	ZapMasterCaller     // Read-only binding to the contract
	ZapMasterTransactor // Write-only binding to the contract
	ZapMasterFilterer   // Log filterer for contract events
}

// ZapMasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapMasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapMasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapMasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapMasterSession struct {
	Contract     *ZapMaster        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapMasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapMasterCallerSession struct {
	Contract *ZapMasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZapMasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapMasterTransactorSession struct {
	Contract     *ZapMasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZapMasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapMasterRaw struct {
	Contract *ZapMaster // Generic contract binding to access the raw methods on
}

// ZapMasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapMasterCallerRaw struct {
	Contract *ZapMasterCaller // Generic read-only contract binding to access the raw methods on
}

// ZapMasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapMasterTransactorRaw struct {
	Contract *ZapMasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapMaster creates a new instance of ZapMaster, bound to a specific deployed contract.
func NewZapMaster(address common.Address, backend bind.ContractBackend) (*ZapMaster, error) {
	contract, err := bindZapMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapMaster{ZapMasterCaller: ZapMasterCaller{contract: contract}, ZapMasterTransactor: ZapMasterTransactor{contract: contract}, ZapMasterFilterer: ZapMasterFilterer{contract: contract}}, nil
}

// NewZapMasterCaller creates a new read-only instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterCaller(address common.Address, caller bind.ContractCaller) (*ZapMasterCaller, error) {
	contract, err := bindZapMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapMasterCaller{contract: contract}, nil
}

// NewZapMasterTransactor creates a new write-only instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapMasterTransactor, error) {
	contract, err := bindZapMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapMasterTransactor{contract: contract}, nil
}

// NewZapMasterFilterer creates a new log filterer instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapMasterFilterer, error) {
	contract, err := bindZapMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapMasterFilterer{contract: contract}, nil
}

// bindZapMaster binds a generic wrapper to an already deployed contract.
func bindZapMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapMasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapMaster *ZapMasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapMaster.Contract.ZapMasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapMaster *ZapMasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapMaster.Contract.ZapMasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapMaster *ZapMasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapMaster.Contract.ZapMasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapMaster *ZapMasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapMaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapMaster *ZapMasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapMaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapMaster *ZapMasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapMaster.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapMaster *ZapMasterSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.Allowance(&_ZapMaster.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.Allowance(&_ZapMaster.CallOpts, _user, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapMaster *ZapMasterSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOf(&_ZapMaster.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOf(&_ZapMaster.CallOpts, _user)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapMaster *ZapMasterCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapMaster *ZapMasterSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapMaster.Contract.DidMine(&_ZapMaster.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapMaster *ZapMasterCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapMaster.Contract.DidMine(&_ZapMaster.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapMaster *ZapMasterCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapMaster *ZapMasterSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapMaster.Contract.DidVote(&_ZapMaster.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapMaster *ZapMasterCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapMaster.Contract.DidVote(&_ZapMaster.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapMaster *ZapMasterCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapMaster *ZapMasterSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapMaster.Contract.GetAddressVars(&_ZapMaster.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapMaster *ZapMasterCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapMaster.Contract.GetAddressVars(&_ZapMaster.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, address, address, address, uint256, uint256[9], int256)
func (_ZapMaster *ZapMasterCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, common.Address, common.Address, common.Address, *big.Int, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new(*big.Int), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, address, address, address, uint256, uint256[9], int256)
func (_ZapMaster *ZapMasterSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, common.Address, common.Address, common.Address, *big.Int, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, address, address, address, uint256, uint256[9], int256)
func (_ZapMaster *ZapMasterCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, common.Address, common.Address, common.Address, *big.Int, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getCurrentVariables")

	if err != nil {
		return *new([32]byte), *new(*big.Int), *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetCurrentVariables(&_ZapMaster.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetCurrentVariables(&_ZapMaster.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeIdByDisputeHash(&_ZapMaster.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeIdByDisputeHash(&_ZapMaster.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeUintVars(&_ZapMaster.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeUintVars(&_ZapMaster.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapMaster *ZapMasterCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getLastNewValue")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapMaster *ZapMasterSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValue(&_ZapMaster.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapMaster *ZapMasterCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValue(&_ZapMaster.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapMaster *ZapMasterCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getLastNewValueById", _requestId)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapMaster *ZapMasterSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValueById(&_ZapMaster.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapMaster *ZapMasterCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValueById(&_ZapMaster.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetMinedBlockNum(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetMinedBlockNum(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapMaster *ZapMasterCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapMaster *ZapMasterSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapMaster.Contract.GetMinersByRequestIdAndTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapMaster *ZapMasterCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapMaster.Contract.GetMinersByRequestIdAndTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetNewValueCountbyRequestId(&_ZapMaster.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetNewValueCountbyRequestId(&_ZapMaster.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestIdByQueryHash", _request)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByQueryHash(&_ZapMaster.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByQueryHash(&_ZapMaster.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByRequestQIndex(&_ZapMaster.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByRequestQIndex(&_ZapMaster.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByTimestamp(&_ZapMaster.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByTimestamp(&_ZapMaster.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapMaster *ZapMasterCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapMaster *ZapMasterSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapMaster.Contract.GetRequestQ(&_ZapMaster.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapMaster *ZapMasterCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapMaster.Contract.GetRequestQ(&_ZapMaster.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestUintVars(&_ZapMaster.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestUintVars(&_ZapMaster.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(string), *new(string), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetRequestVars(&_ZapMaster.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetRequestVars(&_ZapMaster.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getStakerInfo", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetStakerInfo(&_ZapMaster.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetStakerInfo(&_ZapMaster.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapMaster *ZapMasterCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapMaster *ZapMasterSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapMaster.Contract.GetSubmissionsByTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapMaster *ZapMasterCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapMaster.Contract.GetSubmissionsByTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetTimestampbyRequestIDandIndex(&_ZapMaster.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetTimestampbyRequestIDandIndex(&_ZapMaster.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetUintVar(&_ZapMaster.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetUintVar(&_ZapMaster.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getVariablesOnDeck")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapMaster.Contract.GetVariablesOnDeck(&_ZapMaster.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapMaster.Contract.GetVariablesOnDeck(&_ZapMaster.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapMaster *ZapMasterCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapMaster *ZapMasterSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapMaster.Contract.IsInDispute(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapMaster *ZapMasterCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapMaster.Contract.IsInDispute(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapMaster *ZapMasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapMaster *ZapMasterSession) Owner() (common.Address, error) {
	return _ZapMaster.Contract.Owner(&_ZapMaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapMaster *ZapMasterCallerSession) Owner() (common.Address, error) {
	return _ZapMaster.Contract.Owner(&_ZapMaster.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.RetrieveData(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.RetrieveData(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapMaster *ZapMasterCaller) TotalTokenSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "totalTokenSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapMaster *ZapMasterSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapMaster.Contract.TotalTokenSupply(&_ZapMaster.CallOpts)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapMaster.Contract.TotalTokenSupply(&_ZapMaster.CallOpts)
}

// ChangeVaultContract is a paid mutator transaction binding the contract method 0xaf2a5102.
//
// Solidity: function changeVaultContract(address _vaultContract) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeVaultContract(opts *bind.TransactOpts, _vaultContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeVaultContract", _vaultContract)
}

// ChangeVaultContract is a paid mutator transaction binding the contract method 0xaf2a5102.
//
// Solidity: function changeVaultContract(address _vaultContract) returns()
func (_ZapMaster *ZapMasterSession) ChangeVaultContract(_vaultContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeVaultContract(&_ZapMaster.TransactOpts, _vaultContract)
}

// ChangeVaultContract is a paid mutator transaction binding the contract method 0xaf2a5102.
//
// Solidity: function changeVaultContract(address _vaultContract) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeVaultContract(_vaultContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeVaultContract(&_ZapMaster.TransactOpts, _vaultContract)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZapMaster *ZapMasterTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ZapMaster.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZapMaster *ZapMasterSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZapMaster.Contract.Fallback(&_ZapMaster.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZapMaster *ZapMasterTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZapMaster.Contract.Fallback(&_ZapMaster.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZapMaster *ZapMasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapMaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZapMaster *ZapMasterSession) Receive() (*types.Transaction, error) {
	return _ZapMaster.Contract.Receive(&_ZapMaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZapMaster *ZapMasterTransactorSession) Receive() (*types.Transaction, error) {
	return _ZapMaster.Contract.Receive(&_ZapMaster.TransactOpts)
}

// ZapMasterNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapMaster contract.
type ZapMasterNewZapAddressIterator struct {
	Event *ZapMasterNewZapAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapMasterNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapMasterNewZapAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapMasterNewZapAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapMasterNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapMasterNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapMasterNewZapAddress represents a NewZapAddress event raised by the ZapMaster contract.
type ZapMasterNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapMaster *ZapMasterFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapMasterNewZapAddressIterator, error) {

	logs, sub, err := _ZapMaster.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapMasterNewZapAddressIterator{contract: _ZapMaster.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapMaster *ZapMasterFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapMasterNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapMaster.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapMasterNewZapAddress)
				if err := _ZapMaster.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewZapAddress is a log parse operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapMaster *ZapMasterFilterer) ParseNewZapAddress(log types.Log) (*ZapMasterNewZapAddress, error) {
	event := new(ZapMasterNewZapAddress)
	if err := _ZapMaster.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapMasterReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the ZapMaster contract.
type ZapMasterReceivedIterator struct {
	Event *ZapMasterReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapMasterReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapMasterReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapMasterReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapMasterReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapMasterReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapMasterReceived represents a Received event raised by the ZapMaster contract.
type ZapMasterReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ZapMaster *ZapMasterFilterer) FilterReceived(opts *bind.FilterOpts) (*ZapMasterReceivedIterator, error) {

	logs, sub, err := _ZapMaster.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &ZapMasterReceivedIterator{contract: _ZapMaster.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ZapMaster *ZapMasterFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ZapMasterReceived) (event.Subscription, error) {

	logs, sub, err := _ZapMaster.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapMasterReceived)
				if err := _ZapMaster.contract.UnpackLog(event, "Received", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ZapMaster *ZapMasterFilterer) ParseReceived(log types.Log) (*ZapMasterReceived, error) {
	event := new(ZapMasterReceived)
	if err := _ZapMaster.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStakeMetaData contains all meta data concerning the ZapStake contract.
var ZapStakeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"326991a3": "depositStake(ZapStorage.ZapStorageStruct storage)",
		"47b024eb": "init(ZapStorage.ZapStorageStruct storage)",
		"3c734827": "requestStakingWithdraw(ZapStorage.ZapStorageStruct storage)",
		"78bfa277": "withdrawStake(ZapStorage.ZapStorageStruct storage)",
	},
	Bin: "0x6107c061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c8063326991a31461005b5780633c7348271461007d57806347b024eb1461009d57806378bfa277146100bd575b600080fd5b81801561006757600080fd5b5061007b61007636600461070a565b6100dd565b005b81801561008957600080fd5b5061007b61009836600461070a565b61014d565b8180156100a957600080fd5b5061007b6100b836600461070a565b6102a5565b8180156100c957600080fd5b5061007b6100d836600461070a565b61047c565b6100e781336105cc565b60405163124c971160e31b81526004810182905273__$f08294d4d1c904fbf7968e6213a9d2a1b1$__90639264b8889060240160006040518083038186803b15801561013257600080fd5b505af4158015610146573d6000803e3d6000fd5b5050505050565b336000908152604582016020526040902080546001146101aa5760405162461bcd60e51b8152602060048201526013602482015272135a5b995c881a5cc81b9bdd081cdd185ad959606a1b60448201526064015b60405180910390fd5b600281556101bb6201518042610723565b6101c5904261075b565b6001808301919091557fedddb9344bfe0dadc78c558b8ffca446679cbffc17be64eb83973fce7bea5f34600090815260408085016020528120805490919061020e90849061075b565b909155505060405163124c971160e31b81526004810183905273__$f08294d4d1c904fbf7968e6213a9d2a1b1$__90639264b8889060240160006040518083038186803b15801561025e57600080fd5b505af4158015610272573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b7f784c4fb1ab068f6039d5780c68dd0fa2f8742cceb3426d19667778ca7f3518a960009081526040808301602052902054156103195760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b60448201526064016101a1565b7f784c4fb1ab068f6039d5780c68dd0fa2f8742cceb3426d19667778ca7f3518a960009081526040808301602052808220601290557fabef544d8048318ece54fb2c6385255cd1b06e176525d149a0338a7acca6deb3825280822060c890557f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b282528082206969e10de76676d080000090557f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc4923982528082206834957444b840e8000090557fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc3382529020610258908190556104139042610723565b61041d904261075b565b7f97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a460009081526040928301602052828120919091557fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e81522060019055565b3360009081526045820160205260409020600181015462093a80906104a46201518042610723565b6104ae904261075b565b6104b8919061075b565b101561053c5760405162461bcd60e51b815260206004820152604760248201527f43616e2774207769746864726177207965742e204e65656420746f207761697460448201527f206174204c45415354203720646179732066726f6d207374616b6520737461726064820152663a103230ba329760c91b608482015260a4016101a1565b805460021461059b5760405162461bcd60e51b815260206004820152602560248201527f526571756972656420746f2072657175657374207769746864726177206f66206044820152647374616b6560d81b60648201526084016101a1565b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6001600160a01b03811660009081526045830160205260409020541561063f5760405162461bcd60e51b815260206004820152602260248201527f5a61705374616b653a205374616b657220697320616c7265616479207374616b604482015261195960f21b60648201526084016101a1565b7fedddb9344bfe0dadc78c558b8ffca446679cbffc17be64eb83973fce7bea5f34600090815260408084016020528120805460019290610680908490610772565b90915550506040805180820190915260018152602081016106a46201518042610723565b6106ae904261075b565b90526001600160a01b03821660008181526045850160209081526040808320855181559490910151600190940193909355915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a25050565b60006020828403121561071c57600080fd5b5035919050565b60008261074057634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60008282101561076d5761076d610745565b500390565b6000821982111561078557610785610745565b50019056fea2646970667358221220809b8672d75e8611e53c935881e42029190f9246559a7e5824bd9c57c77b487864736f6c634300080b0033",
}

// ZapStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapStakeMetaData.ABI instead.
var ZapStakeABI = ZapStakeMetaData.ABI

// Deprecated: Use ZapStakeMetaData.Sigs instead.
// ZapStakeFuncSigs maps the 4-byte function signature to its string representation.
var ZapStakeFuncSigs = ZapStakeMetaData.Sigs

// ZapStakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapStakeMetaData.Bin instead.
var ZapStakeBin = ZapStakeMetaData.Bin

// DeployZapStake deploys a new Ethereum contract, binding an instance of ZapStake to it.
func DeployZapStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStake, error) {
	parsed, err := ZapStakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	zapDisputeAddr, _, _, _ := DeployZapDispute(auth, backend)
	ZapStakeBin = strings.Replace(ZapStakeBin, "__$f08294d4d1c904fbf7968e6213a9d2a1b1$__", zapDisputeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapStake{ZapStakeCaller: ZapStakeCaller{contract: contract}, ZapStakeTransactor: ZapStakeTransactor{contract: contract}, ZapStakeFilterer: ZapStakeFilterer{contract: contract}}, nil
}

// ZapStake is an auto generated Go binding around an Ethereum contract.
type ZapStake struct {
	ZapStakeCaller     // Read-only binding to the contract
	ZapStakeTransactor // Write-only binding to the contract
	ZapStakeFilterer   // Log filterer for contract events
}

// ZapStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapStakeSession struct {
	Contract     *ZapStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapStakeCallerSession struct {
	Contract *ZapStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZapStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapStakeTransactorSession struct {
	Contract     *ZapStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZapStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapStakeRaw struct {
	Contract *ZapStake // Generic contract binding to access the raw methods on
}

// ZapStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapStakeCallerRaw struct {
	Contract *ZapStakeCaller // Generic read-only contract binding to access the raw methods on
}

// ZapStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapStakeTransactorRaw struct {
	Contract *ZapStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapStake creates a new instance of ZapStake, bound to a specific deployed contract.
func NewZapStake(address common.Address, backend bind.ContractBackend) (*ZapStake, error) {
	contract, err := bindZapStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapStake{ZapStakeCaller: ZapStakeCaller{contract: contract}, ZapStakeTransactor: ZapStakeTransactor{contract: contract}, ZapStakeFilterer: ZapStakeFilterer{contract: contract}}, nil
}

// NewZapStakeCaller creates a new read-only instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeCaller(address common.Address, caller bind.ContractCaller) (*ZapStakeCaller, error) {
	contract, err := bindZapStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStakeCaller{contract: contract}, nil
}

// NewZapStakeTransactor creates a new write-only instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapStakeTransactor, error) {
	contract, err := bindZapStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStakeTransactor{contract: contract}, nil
}

// NewZapStakeFilterer creates a new log filterer instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapStakeFilterer, error) {
	contract, err := bindZapStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapStakeFilterer{contract: contract}, nil
}

// bindZapStake binds a generic wrapper to an already deployed contract.
func bindZapStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStake *ZapStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStake.Contract.ZapStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStake *ZapStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStake.Contract.ZapStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStake *ZapStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStake.Contract.ZapStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStake *ZapStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStake *ZapStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStake *ZapStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStake.Contract.contract.Transact(opts, method, params...)
}

// ZapStakeNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the ZapStake contract.
type ZapStakeNewStakeIterator struct {
	Event *ZapStakeNewStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapStakeNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeNewStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapStakeNewStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapStakeNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeNewStake represents a NewStake event raised by the ZapStake contract.
type ZapStakeNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeNewStakeIterator{contract: _ZapStake.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *ZapStakeNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeNewStake)
				if err := _ZapStake.contract.UnpackLog(event, "NewStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewStake is a log parse operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) ParseNewStake(log types.Log) (*ZapStakeNewStake, error) {
	event := new(ZapStakeNewStake)
	if err := _ZapStake.contract.UnpackLog(event, "NewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStakeStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the ZapStake contract.
type ZapStakeStakeWithdrawRequestedIterator struct {
	Event *ZapStakeStakeWithdrawRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapStakeStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeStakeWithdrawRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapStakeStakeWithdrawRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapStakeStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the ZapStake contract.
type ZapStakeStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeStakeWithdrawRequestedIterator{contract: _ZapStake.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ZapStakeStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeStakeWithdrawRequested)
				if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdrawRequested is a log parse operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) ParseStakeWithdrawRequested(log types.Log) (*ZapStakeStakeWithdrawRequested, error) {
	event := new(ZapStakeStakeWithdrawRequested)
	if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStakeStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the ZapStake contract.
type ZapStakeStakeWithdrawnIterator struct {
	Event *ZapStakeStakeWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZapStakeStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeStakeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZapStakeStakeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZapStakeStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeStakeWithdrawn represents a StakeWithdrawn event raised by the ZapStake contract.
type ZapStakeStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeStakeWithdrawnIterator{contract: _ZapStake.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ZapStakeStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeStakeWithdrawn)
				if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) ParseStakeWithdrawn(log types.Log) (*ZapStakeStakeWithdrawn, error) {
	event := new(ZapStakeStakeWithdrawn)
	if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStorageMetaData contains all meta data concerning the ZapStorage contract.
var ZapStorageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204eafa5cc98cac01f3577a9d0932a8b03e11aa54d93ee24bb331de2f878a2203e64736f6c634300080b0033",
}

// ZapStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ZapStorageMetaData.ABI instead.
var ZapStorageABI = ZapStorageMetaData.ABI

// ZapStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZapStorageMetaData.Bin instead.
var ZapStorageBin = ZapStorageMetaData.Bin

// DeployZapStorage deploys a new Ethereum contract, binding an instance of ZapStorage to it.
func DeployZapStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStorage, error) {
	parsed, err := ZapStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZapStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapStorage{ZapStorageCaller: ZapStorageCaller{contract: contract}, ZapStorageTransactor: ZapStorageTransactor{contract: contract}, ZapStorageFilterer: ZapStorageFilterer{contract: contract}}, nil
}

// ZapStorage is an auto generated Go binding around an Ethereum contract.
type ZapStorage struct {
	ZapStorageCaller     // Read-only binding to the contract
	ZapStorageTransactor // Write-only binding to the contract
	ZapStorageFilterer   // Log filterer for contract events
}

// ZapStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapStorageSession struct {
	Contract     *ZapStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapStorageCallerSession struct {
	Contract *ZapStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapStorageTransactorSession struct {
	Contract     *ZapStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapStorageRaw struct {
	Contract *ZapStorage // Generic contract binding to access the raw methods on
}

// ZapStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapStorageCallerRaw struct {
	Contract *ZapStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ZapStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapStorageTransactorRaw struct {
	Contract *ZapStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapStorage creates a new instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorage(address common.Address, backend bind.ContractBackend) (*ZapStorage, error) {
	contract, err := bindZapStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapStorage{ZapStorageCaller: ZapStorageCaller{contract: contract}, ZapStorageTransactor: ZapStorageTransactor{contract: contract}, ZapStorageFilterer: ZapStorageFilterer{contract: contract}}, nil
}

// NewZapStorageCaller creates a new read-only instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageCaller(address common.Address, caller bind.ContractCaller) (*ZapStorageCaller, error) {
	contract, err := bindZapStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStorageCaller{contract: contract}, nil
}

// NewZapStorageTransactor creates a new write-only instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapStorageTransactor, error) {
	contract, err := bindZapStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStorageTransactor{contract: contract}, nil
}

// NewZapStorageFilterer creates a new log filterer instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapStorageFilterer, error) {
	contract, err := bindZapStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapStorageFilterer{contract: contract}, nil
}

// bindZapStorage binds a generic wrapper to an already deployed contract.
func bindZapStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStorage *ZapStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStorage.Contract.ZapStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStorage *ZapStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStorage.Contract.ZapStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStorage *ZapStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStorage.Contract.ZapStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStorage *ZapStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStorage *ZapStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStorage *ZapStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStorage.Contract.contract.Transact(opts, method, params...)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts1

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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BasicTokenFuncSigs maps the 4-byte function signature to its string representation.
var BasicTokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
var BasicTokenBin = "0x60806040526a52b7d2dcc80cd2e400000060005534801561001f57600080fd5b506102098061002f6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806318160ddd1461004657806370a0823114610060578063a9059cbb14610086575b600080fd5b61004e6100c6565b60408051918252519081900360200190f35b61004e6004803603602081101561007657600080fd5b50356001600160a01b03166100cc565b6100b26004803603604081101561009c57600080fd5b506001600160a01b0381351690602001356100e7565b604080519115158252519081900360200190f35b60005481565b6001600160a01b031660009081526001602052604090205490565b60006001600160a01b0383166100fc57600080fd5b3360009081526001602052604090205461011c908363ffffffff6101ac16565b33600090815260016020526040808220929092556001600160a01b0385168152205461014e908363ffffffff6101be16565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6000828211156101b857fe5b50900390565b6000828201838110156101cd57fe5b939250505056fea265627a7a72315820008e572f461884b34225379af3a7ae2b210c8010adbe9bab6493c4d25d36282b64736f6c63430005100032"

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BasicToken *BasicTokenFilterer) ParseTransfer(log types.Log) (*BasicTokenTransfer, error) {
	event := new(BasicTokenTransfer)
	if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BasicFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BasicFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Basic *ERC20BasicFilterer) ParseTransfer(log types.Log) (*ERC20BasicTransfer, error) {
	event := new(ERC20BasicTransfer)
	if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenABI is the input ABI used to generate the binding from.
const MintableTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MintableTokenFuncSigs maps the 4-byte function signature to its string representation.
var MintableTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"66188463": "decreaseApproval(address,uint256)",
	"7d64bcb4": "finishMinting()",
	"893d20e8": "getOwner()",
	"d73dd623": "increaseApproval(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"05d2035b": "mintingFinished()",
	"8da5cb5b": "owner()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// MintableTokenBin is the compiled bytecode used for deploying new contracts.
var MintableTokenBin = "0x60806040526a52b7d2dcc80cd2e4000000600055600380546001600160a81b0319163317905561097d806100346000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637d64bcb41161008c578063a9059cbb11610066578063a9059cbb14610239578063d73dd62314610265578063dd62ed3e14610291578063f2fde38b146102bf576100ea565b80637d64bcb414610205578063893d20e81461020d5780638da5cb5b14610231576100ea565b806323b872dd116100c857806323b872dd1461015157806340c10f191461018757806366188463146101b357806370a08231146101df576100ea565b806305d2035b146100ef578063095ea7b31461010b57806318160ddd14610137575b600080fd5b6100f76102e7565b604080519115158252519081900360200190f35b6100f76004803603604081101561012157600080fd5b506001600160a01b0381351690602001356102f7565b61013f61035d565b60408051918252519081900360200190f35b6100f76004803603606081101561016757600080fd5b506001600160a01b03813581169160208101359091169060400135610363565b6100f76004803603604081101561019d57600080fd5b506001600160a01b038135169060200135610481565b6100f7600480360360408110156101c957600080fd5b506001600160a01b03813516906020013561058c565b61013f600480360360208110156101f557600080fd5b50356001600160a01b031661067c565b6100f7610697565b6102156106f3565b604080516001600160a01b039092168252519081900360200190f35b610215610702565b6100f76004803603604081101561024f57600080fd5b506001600160a01b038135169060200135610711565b6100f76004803603604081101561027b57600080fd5b506001600160a01b0381351690602001356107d6565b61013f600480360360408110156102a757600080fd5b506001600160a01b038135811691602001351661086f565b6102e5600480360360208110156102d557600080fd5b50356001600160a01b031661089a565b005b600354600160a01b900460ff1681565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60006001600160a01b03831661037857600080fd5b6001600160a01b038416600081815260026020908152604080832033845282528083205493835260019091529020546103b7908463ffffffff61092016565b6001600160a01b0380871660009081526001602052604080822093909355908616815220546103ec908463ffffffff61093216565b6001600160a01b038516600090815260016020526040902055610415818463ffffffff61092016565b6001600160a01b03808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b6003546000906001600160a01b0316331461049b57600080fd5b600354600160a01b900460ff16156104b257600080fd5b6000546104c5908363ffffffff61093216565b60009081556001600160a01b0384168152600160205260409020546104f0908363ffffffff61093216565b6001600160a01b038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a26040805183815290516001600160a01b038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600192915050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054808311156105e1573360009081526002602090815260408083206001600160a01b0388168452909152812055610616565b6105f1818463ffffffff61092016565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526001602052604090205490565b6003546000906001600160a01b031633146106b157600080fd5b6003805460ff60a01b1916600160a01b1790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b6003546001600160a01b031690565b6003546001600160a01b031681565b60006001600160a01b03831661072657600080fd5b33600090815260016020526040902054610746908363ffffffff61092016565b33600090815260016020526040808220929092556001600160a01b03851681522054610778908363ffffffff61093216565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b038616845290915281205461080a908363ffffffff61093216565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6003546001600160a01b031633146108b157600080fd5b6001600160a01b0381166108c457600080fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b60008282111561092c57fe5b50900390565b60008282018381101561094157fe5b939250505056fea265627a7a723158206c78c313c2de952ea406de3b816b2954f7c9591e9df6c18f7b11f87406301bd664736f6c63430005100032"

// DeployMintableToken deploys a new Ethereum contract, binding an instance of MintableToken to it.
func DeployMintableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MintableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MintableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}, MintableTokenFilterer: MintableTokenFilterer{contract: contract}}, nil
}

// MintableToken is an auto generated Go binding around an Ethereum contract.
type MintableToken struct {
	MintableTokenCaller     // Read-only binding to the contract
	MintableTokenTransactor // Write-only binding to the contract
	MintableTokenFilterer   // Log filterer for contract events
}

// MintableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintableTokenSession struct {
	Contract     *MintableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintableTokenCallerSession struct {
	Contract *MintableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MintableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintableTokenTransactorSession struct {
	Contract     *MintableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MintableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintableTokenRaw struct {
	Contract *MintableToken // Generic contract binding to access the raw methods on
}

// MintableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintableTokenCallerRaw struct {
	Contract *MintableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MintableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintableTokenTransactorRaw struct {
	Contract *MintableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintableToken creates a new instance of MintableToken, bound to a specific deployed contract.
func NewMintableToken(address common.Address, backend bind.ContractBackend) (*MintableToken, error) {
	contract, err := bindMintableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}, MintableTokenFilterer: MintableTokenFilterer{contract: contract}}, nil
}

// NewMintableTokenCaller creates a new read-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenCaller(address common.Address, caller bind.ContractCaller) (*MintableTokenCaller, error) {
	contract, err := bindMintableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenCaller{contract: contract}, nil
}

// NewMintableTokenTransactor creates a new write-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MintableTokenTransactor, error) {
	contract, err := bindMintableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransactor{contract: contract}, nil
}

// NewMintableTokenFilterer creates a new log filterer instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MintableTokenFilterer, error) {
	contract, err := bindMintableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintableTokenFilterer{contract: contract}, nil
}

// bindMintableToken binds a generic wrapper to an already deployed contract.
func bindMintableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.MintableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_MintableToken *MintableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_MintableToken *MintableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_MintableToken *MintableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_MintableToken *MintableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_MintableToken *MintableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_MintableToken *MintableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MintableToken *MintableTokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MintableToken *MintableTokenSession) GetOwner() (common.Address, error) {
	return _MintableToken.Contract.GetOwner(&_MintableToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MintableToken *MintableTokenCallerSession) GetOwner() (common.Address, error) {
	return _MintableToken.Contract.GetOwner(&_MintableToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_MintableToken *MintableTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_MintableToken *MintableTokenSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_MintableToken *MintableTokenCallerSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintableToken *MintableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintableToken *MintableTokenSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintableToken *MintableTokenCallerSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableToken *MintableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableToken *MintableTokenSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableToken *MintableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MintableToken *MintableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.DecreaseApproval(&_MintableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.DecreaseApproval(&_MintableToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MintableToken *MintableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.IncreaseApproval(&_MintableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.IncreaseApproval(&_MintableToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_MintableToken *MintableTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_MintableToken *MintableTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintableToken *MintableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintableToken *MintableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintableToken *MintableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// MintableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MintableToken contract.
type MintableTokenApprovalIterator struct {
	Event *MintableTokenApproval // Event containing the contract specifics and raw log

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
func (it *MintableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenApproval)
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
		it.Event = new(MintableTokenApproval)
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
func (it *MintableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenApproval represents a Approval event raised by the MintableToken contract.
type MintableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintableToken *MintableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MintableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenApprovalIterator{contract: _MintableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintableToken *MintableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MintableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenApproval)
				if err := _MintableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MintableToken *MintableTokenFilterer) ParseApproval(log types.Log) (*MintableTokenApproval, error) {
	event := new(MintableTokenApproval)
	if err := _MintableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MintableToken contract.
type MintableTokenMintIterator struct {
	Event *MintableTokenMint // Event containing the contract specifics and raw log

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
func (it *MintableTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenMint)
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
		it.Event = new(MintableTokenMint)
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
func (it *MintableTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenMint represents a Mint event raised by the MintableToken contract.
type MintableTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_MintableToken *MintableTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*MintableTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenMintIterator{contract: _MintableToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_MintableToken *MintableTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MintableTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenMint)
				if err := _MintableToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_MintableToken *MintableTokenFilterer) ParseMint(log types.Log) (*MintableTokenMint, error) {
	event := new(MintableTokenMint)
	if err := _MintableToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the MintableToken contract.
type MintableTokenMintFinishedIterator struct {
	Event *MintableTokenMintFinished // Event containing the contract specifics and raw log

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
func (it *MintableTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenMintFinished)
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
		it.Event = new(MintableTokenMintFinished)
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
func (it *MintableTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenMintFinished represents a MintFinished event raised by the MintableToken contract.
type MintableTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_MintableToken *MintableTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*MintableTokenMintFinishedIterator, error) {

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &MintableTokenMintFinishedIterator{contract: _MintableToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_MintableToken *MintableTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *MintableTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenMintFinished)
				if err := _MintableToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// ParseMintFinished is a log parse operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_MintableToken *MintableTokenFilterer) ParseMintFinished(log types.Log) (*MintableTokenMintFinished, error) {
	event := new(MintableTokenMintFinished)
	if err := _MintableToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MintableToken contract.
type MintableTokenOwnershipTransferredIterator struct {
	Event *MintableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenOwnershipTransferred)
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
		it.Event = new(MintableTokenOwnershipTransferred)
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
func (it *MintableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MintableToken contract.
type MintableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintableToken *MintableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenOwnershipTransferredIterator{contract: _MintableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintableToken *MintableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenOwnershipTransferred)
				if err := _MintableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MintableToken *MintableTokenFilterer) ParseOwnershipTransferred(log types.Log) (*MintableTokenOwnershipTransferred, error) {
	event := new(MintableTokenOwnershipTransferred)
	if err := _MintableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MintableToken contract.
type MintableTokenTransferIterator struct {
	Event *MintableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MintableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenTransfer)
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
		it.Event = new(MintableTokenTransfer)
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
func (it *MintableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenTransfer represents a Transfer event raised by the MintableToken contract.
type MintableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintableToken *MintableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MintableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransferIterator{contract: _MintableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintableToken *MintableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MintableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenTransfer)
				if err := _MintableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MintableToken *MintableTokenFilterer) ParseTransfer(log types.Log) (*MintableTokenTransfer, error) {
	event := new(MintableTokenTransfer)
	if err := _MintableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"893d20e8": "getOwner()",
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610172806100326000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063893d20e8146100465780638da5cb5b1461006a578063f2fde38b14610072575b600080fd5b61004e61009a565b604080516001600160a01b039092168252519081900360200190f35b61004e6100a9565b6100986004803603602081101561008857600080fd5b50356001600160a01b03166100b8565b005b6000546001600160a01b031690565b6000546001600160a01b031681565b6000546001600160a01b031633146100cf57600080fd5b6001600160a01b0381166100e257600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820116978fe3d8d2652ecb0444bc1fa327c421671fad54692c08d2f418df65285ea64736f6c63430005100032"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Ownable *OwnableCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Ownable *OwnableSession) GetOwner() (common.Address, error) {
	return _Ownable.Contract.GetOwner(&_Ownable.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Ownable *OwnableCallerSession) GetOwner() (common.Address, error) {
	return _Ownable.Contract.GetOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820ded23893a148c75765811c96ce4ec71701d71918c0e7dbc873e6deea613caeee64736f6c63430005100032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

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

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StandardTokenFuncSigs maps the 4-byte function signature to its string representation.
var StandardTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"66188463": "decreaseApproval(address,uint256)",
	"d73dd623": "increaseApproval(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
var StandardTokenBin = "0x60806040526a52b7d2dcc80cd2e400000060005534801561001f57600080fd5b506106708061002f6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a9059cbb1461016f578063d73dd6231461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063661884631461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561025b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610261565b6100b96004803603604081101561013357600080fd5b506001600160a01b03813516906020013561037f565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b031661046f565b6100b96004803603604081101561018557600080fd5b506001600160a01b03813516906020013561048a565b6100b9600480360360408110156101b157600080fd5b506001600160a01b03813516906020013561054f565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b03813581169160200135166105e8565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60006001600160a01b03831661027657600080fd5b6001600160a01b038416600081815260026020908152604080832033845282528083205493835260019091529020546102b5908463ffffffff61061316565b6001600160a01b0380871660009081526001602052604080822093909355908616815220546102ea908463ffffffff61062516565b6001600160a01b038516600090815260016020526040902055610313818463ffffffff61061316565b6001600160a01b03808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054808311156103d4573360009081526002602090815260408083206001600160a01b0388168452909152812055610409565b6103e4818463ffffffff61061316565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526001602052604090205490565b60006001600160a01b03831661049f57600080fd5b336000908152600160205260409020546104bf908363ffffffff61061316565b33600090815260016020526040808220929092556001600160a01b038516815220546104f1908363ffffffff61062516565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610583908363ffffffff61062516565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561061f57fe5b50900390565b60008282018381101561063457fe5b939250505056fea265627a7a72315820f70346d88aba88240568286dc9e4f1c0965930a9f47ed07ab7123f770218e26f64736f6c63430005100032"

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StandardToken *StandardTokenFilterer) ParseApproval(log types.Log) (*StandardTokenApproval, error) {
	event := new(StandardTokenApproval)
	if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StandardToken *StandardTokenFilterer) ParseTransfer(log types.Log) (*StandardTokenTransfer, error) {
	event := new(StandardTokenTransfer)
	if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[]"

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
var UtilitiesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820369cc1edbe25d01e1c1f04ab5f7bdbf9dc056bf50b0d65a166a37a9605c0f87464736f6c63430005100032"

// DeployUtilities deploys a new Ethereum contract, binding an instance of Utilities to it.
func DeployUtilities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utilities, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilitiesBin), backend)
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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"miniVault\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"miniVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"authorizedUser\",\"type\":\"address\"}],\"name\":\"lockSmith\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = map[string]string{
	"47e7ef24": "deposit(address,uint256)",
	"1709ef07": "hasAccess(address,address)",
	"cd8f942c": "increaseApproval()",
	"7a47328b": "lockSmith(address,address)",
	"0103c92b": "userBalance(address)",
	"f3fef3a3": "withdraw(address,uint256)",
	"2b56eb56": "zapToken()",
}

// VaultBin is the compiled bytecode used for deploying new contracts.
var VaultBin = "0x608060405234801561001057600080fd5b50604051610a0d380380610a0d8339818101604052604081101561003357600080fd5b508051602091820151600080546001600160a01b038085166001600160a01b0319928316811790935560018054918516919092168117909155604080516024810192909252600019604480840191909152815180840390910181526064909201815294810180516001600160e01b031663095ea7b360e01b178152945181519495939492939192909182918083835b602083106100e15780518252601f1990920191602091820191016100c2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610143576040519150601f19603f3d011682016040523d82523d6000602084013e610148565b606091505b50505050506108b18061015c6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806347e7ef241161005b57806347e7ef24146101205780637a47328b1461014e578063cd8f942c1461017c578063f3fef3a3146101845761007d565b80630103c92b146100825780631709ef07146100ba5780632b56eb56146100fc575b600080fd5b6100a86004803603602081101561009857600080fd5b50356001600160a01b03166101b0565b60408051918252519081900360200190f35b6100e8600480360360408110156100d057600080fd5b506001600160a01b03813581169160200135166101cb565b604080519115158252519081900360200190f35b61010461024f565b604080516001600160a01b039092168252519081900360200190f35b61014c6004803603604081101561013657600080fd5b506001600160a01b03813516906020013561025e565b005b61014c6004803603604081101561016457600080fd5b506001600160a01b0381358116916020013516610331565b6100e861046b565b61014c6004803603604081101561019a57600080fd5b506001600160a01b038135169060200135610671565b6001600160a01b031660009081526002602052604090205490565b6000331515806101e457506001600160a01b0382163314155b61021f5760405162461bcd60e51b81526004018080602001828103825260268152602001806108576026913960400191505060405180910390fd5b506001600160a01b0380821660009081526003602090815260408083209386168352929052205460ff1692915050565b6000546001600160a01b031681565b6001600160a01b0382166102a35760405162461bcd60e51b81526004018080602001828103825260268152602001806108576026913960400191505060405180910390fd5b6102ad33836101cb565b6102e85760405162461bcd60e51b815260040180806020018281038252602c815260200180610806602c913960400191505060405180910390fd5b6001600160a01b038216600090815260026020526040902054610311908263ffffffff61078116565b6001600160a01b0390921660009081526002602052604090209190915550565b336001600160a01b0383161461038e576040805162461bcd60e51b815260206004820152601a60248201527f596f7520646f206e6f74206f776e2074686973207661756c742e000000000000604482015290519081900360640190fd5b331515806103a557506001600160a01b0382163314155b6103e05760405162461bcd60e51b81526004018080602001828103825260258152602001806108326025913960400191505060405180910390fd5b6001600160a01b038216600090815260036020908152604080832033845290915290205460ff16610437576001600160a01b03821660009081526003602090815260408083209091529020805460ff191660011790555b6001600160a01b0391821660009081526003602090815260408083209390941682529190915220805460ff19166001179055565b60008054600154604080513060248201526001600160a01b0392831660448083019190915282518083039091018152606490910182526020810180516001600160e01b0316636eb1769f60e11b178152915181518695606095169382918083835b602083106104eb5780518252601f1990920191602091820191016104cc565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461054d576040519150601f19603f3d011682016040523d82523d6000602084013e610552565b606091505b50915091506000610576610567836000610797565b6000199063ffffffff6107f316565b60008054600154604080516001600160a01b039283166024820152604480820187905282518083039091018152606490910182526020810180516001600160e01b031663d73dd62360e01b178152915181519697509495606095939094169390929182918083835b602083106105fd5780518252601f1990920191602091820191016105de565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461065f576040519150601f19603f3d011682016040523d82523d6000602084013e610664565b606091505b5090965050505050505090565b6001600160a01b0382166106b65760405162461bcd60e51b81526004018080602001828103825260268152602001806108576026913960400191505060405180910390fd5b6106c033836101cb565b6106fb5760405162461bcd60e51b815260040180806020018281038252602c815260200180610806602c913960400191505060405180910390fd5b80610705836101b0565b1015610758576040805162461bcd60e51b815260206004820152601d60248201527f596f75722062616c616e636520697320696e73756666696369656e742e000000604482015290519081900360640190fd5b6001600160a01b038216600090815260026020526040902054610311908263ffffffff6107f316565b60008282018381101561079057fe5b9392505050565b600081602001835110156107ea576040805162461bcd60e51b8152602060048201526015602482015274746f55696e743235365f6f75744f66426f756e647360581b604482015290519081900360640190fd5b50016020015190565b6000828211156107ff57fe5b5090039056fe596f7520617265206e6f7420617574686f72697a656420746f206163636573732074686973207661756c742e546865207a65726f20616464726573732063616e206e6f74206f776e2061207661756c742e546865207a65726f206164647265737320646f6573206e6f74206f776e2061207661756c742ea265627a7a723158204601ec65e24284bd0e92304a81247ef04f1e08b3871bfcea99faa2e7b045de9564736f6c63430005100032"

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
	err := _Vault.contract.Call(opts, out, "hasAccess", user, miniVault)
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

// ZapABI is the input ABI used to generate the binding from.
const ZapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zapTokenBsc\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"doTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"getBalanceAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseVaultApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewZapAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_c_sapi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_c_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"requestData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractZapTokenBSC\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"updateBalanceAtNow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZapFuncSigs maps the 4-byte function signature to its string representation.
var ZapFuncSigs = map[string]string{
	"752d49a1": "addTip(uint256,uint256)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"0d2d76a2": "depositStake()",
	"37751b35": "doTransfer(address,address,uint256)",
	"26ca8adb": "getBalanceAt(address,uint256)",
	"4049f198": "getNewCurrentVariables()",
	"d18ed17b": "increaseVaultApproval()",
	"8da5cb5b": "owner()",
	"26b7d9f6": "proposeFork(address)",
	"3fff2816": "requestData(string,string,uint256,uint256)",
	"28449c3a": "requestStakingWithdraw()",
	"68c180d5": "submitMiningSolution(string,uint256,uint256)",
	"4d318b0e": "tallyVotes(uint256)",
	"fc0c546a": "token()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"d67dcbc5": "updateBalanceAtNow(address,uint256)",
	"c9d27afe": "vote(uint256,bool)",
	"bed9d861": "withdrawStake()",
}

// ZapBin is the compiled bytecode used for deploying new contracts.
var ZapBin = "0x608060405234801561001057600080fd5b5060405161290c38038061290c8339818101604052602081101561003357600080fd5b5051604b80546001600160a01b039092166001600160a01b0319928316179055604c80549091163317905561289f8061006d6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806370a08231116100b8578063a9059cbb1161007c578063a9059cbb146104db578063bed9d86114610507578063c9d27afe1461050f578063d18ed17b14610534578063d67dcbc51461053c578063fc0c546a1461056857610142565b806370a0823114610419578063752d49a11461043f5780638581af19146104625780638da5cb5b1461048b578063999cf26c146104af57610142565b806328449c3a1161010a57806328449c3a1461022b57806337751b35146102335780633fff2816146102695780634049f198146103315780634d318b0e1461038657806368c180d5146103a357610142565b8063095ea7b3146101475780630d2d76a21461018757806323b872dd1461019157806326b7d9f6146101c757806326ca8adb146101ed575b600080fd5b6101736004803603604081101561015d57600080fd5b506001600160a01b038135169060200135610570565b604080519115158252519081900360200190f35b61018f6105fe565b005b610173600480360360608110156101a757600080fd5b506001600160a01b038135811691602081013590911690604001356108ab565b61018f600480360360208110156101dd57600080fd5b50356001600160a01b031661097e565b6102196004803603604081101561020357600080fd5b506001600160a01b0381351690602001356109f8565b60408051918252519081900360200190f35b61018f610b44565b61018f6004803603606081101561024957600080fd5b506001600160a01b03813581169160208101359091169060400135610bae565b61018f6004803603608081101561027f57600080fd5b81019060208101813564010000000081111561029a57600080fd5b8201836020820111156102ac57600080fd5b803590602001918460018302840111640100000000831117156102ce57600080fd5b9193909290916020810190356401000000008111156102ec57600080fd5b8201836020820111156102fe57600080fd5b8035906020019184600183028401116401000000008311171561032057600080fd5b919350915080359060200135610cdb565b610339611124565b604051848152602081018460a080838360005b8381101561036457818101518382015260200161034c565b5050505090500183815260200182815260200194505050505060405180910390f35b61018f6004803603602081101561039c57600080fd5b5035611149565b61018f600480360360608110156103b957600080fd5b8101906020810181356401000000008111156103d457600080fd5b8201836020820111156103e657600080fd5b8035906020019184600183028401116401000000008311171561040857600080fd5b9193509150803590602001356111fd565b6102196004803603602081101561042f57600080fd5b50356001600160a01b0316611589565b61018f6004803603604081101561045557600080fd5b508035906020013561160c565b61018f6004803603606081101561047857600080fd5b50803590602081013590604001356116ab565b610493611b3a565b604080516001600160a01b039092168252519081900360200190f35b610173600480360360408110156104c557600080fd5b506001600160a01b038135169060200135611b49565b610173600480360360408110156104f157600080fd5b506001600160a01b038135169060200135611bb0565b61018f611c7a565b61018f6004803603604081101561052557600080fd5b50803590602001351515611db3565b610173611e2e565b61018f6004803603604081101561055257600080fd5b506001600160a01b038135169060200135611ecb565b610493611fbd565b604b546040805163095ea7b360e01b81526001600160a01b038581166004830152602482018590529151600093929092169163095ea7b39160448082019260209290919082900301818787803b1580156105c957600080fd5b505af11580156105dd573d6000803e3d6000fd5b505050506040513d60208110156105f357600080fd5b505190505b92915050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b018120600090815260208381529083902054604b546370a0823160e01b84523360048501529351909384936001600160a01b03909116926370a0823192602480840193829003018186803b15801561067257600080fd5b505afa158015610686573d6000803e3d6000fd5b505050506040513d602081101561069c57600080fd5b505110156106a957600080fd5b6040805163326991a360e01b8152600060048201819052915173__$341c69a3b8ea65d7eeecd190190dea4d1b$__9263326991a39260248082019391829003018186803b1580156106f957600080fd5b505af415801561070d573d6000803e3d6000fd5b5050604080516517dd985d5b1d60d21b815281519081900360060181206000908152603f602090815283822054604b5463095ea7b360e01b85523060048601526024850189905294516001600160a01b03918216975087965094169363095ea7b393604480820194918390030190829087803b15801561078c57600080fd5b505af11580156107a0573d6000803e3d6000fd5b505050506040513d60208110156107b657600080fd5b5050604b54604080516323b872dd60e01b81523360048201526001600160a01b03858116602483015260448201879052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561081457600080fd5b505af1158015610828573d6000803e3d6000fd5b505050506040513d602081101561083e57600080fd5b5050604080516311f9fbc960e21b81523360048201526024810185905290516001600160a01b038316916347e7ef2491604480830192600092919082900301818387803b15801561088e57600080fd5b505af11580156108a2573d6000803e3d6000fd5b50505050505050565b6000806108b785611589565b90506108c585848303611ecb565b6108ce84611589565b90508083820110156108df57600080fd5b6108eb84848301611ecb565b604b54604080516323b872dd60e01b81526001600160a01b038881166004830152878116602483015260448201879052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561094957600080fd5b505af115801561095d573d6000803e3d6000fd5b505050506040513d602081101561097357600080fd5b505195945050505050565b6040805163804b893160e01b81526000600482018190526001600160a01b0384166024830152915173__$965cfd6d1cee46a73cf1c675b6712824df$__9263804b89319260448082019391829003018186803b1580156109dd57600080fd5b505af41580156109f1573d6000803e3d6000fd5b5050505050565b6001600160a01b03821660009081526045602052604081208054610a205760009150506105f8565b805481906000198101908110610a3257fe5b6000918252602090912001546001600160801b03168310610a8457805481906000198101908110610a5f57fe5b600091825260209091200154600160801b90046001600160801b031691506105f89050565b80600081548110610a9157fe5b6000918252602090912001546001600160801b0316831015610ab75760009150506105f8565b8054600090600019015b81811115610b12576000600260018385010104905085848281548110610ae357fe5b6000918252602090912001546001600160801b031611610b0557809250610b0c565b6001810391505b50610ac1565b828281548110610b1e57fe5b600091825260209091200154600160801b90046001600160801b03169695505050505050565b60408051633c73482760e01b8152600060048201819052915173__$341c69a3b8ea65d7eeecd190190dea4d1b$__92633c7348279260248082019391829003018186803b158015610b9457600080fd5b505af4158015610ba8573d6000803e3d6000fd5b50505050565b60008111610bbb57600080fd5b6001600160a01b038216610bce57600080fd5b610bd88382611b49565b610be157600080fd5b6000610bec84611589565b9050610bf783611589565b9050808282011015610c0857600080fd5b604b54604080516323b872dd60e01b81526001600160a01b038781166004830152868116602483015260448201869052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610c6657600080fd5b505af1158015610c7a573d6000803e3d6000fd5b505050506040513d6020811015610c9057600080fd5b50506040805183815290516001600160a01b0380861692908716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350505050565b60008211610ce857600080fd5b670de0b6b3a7640000821115610cfd57600080fd5b606086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a0181900481028201810190925288815293945060609392508891508790819084018382808284376000920191909152505084519293505050610d7d57600080fd5b6040815110610d8b57600080fd5b600082856040516020018083805190602001908083835b60208310610dc15780518252601f199092019160209182019101610da2565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820181528351938201939093206000818152604990925292902054919350501515905061110057604080516b1c995c5d595cdd10dbdd5b9d60a21b8082528251600c9281900383018120600090815260208581528582208054600101905592825284519182900390930181208352838252838320546080820185528782528183018790528185018690528451848152808401865260608301528084526048835293909220825180519192610ea492849290910190612716565b506020828101518051610ebd9260018501920190612716565b506040820151600282015560608201518051610ee3916003840191602090910190612794565b505050600081815260486020818152604080842081516a6772616e756c617269747960a81b8152825190819003600b018120865260049091018084528286208c90558686528484526f3932b8bab2b9ba28a837b9b4ba34b7b760811b825282519182900360100182208652808452828620869055868652938352670746f74616c5469760c41b8152815190819003600801902084529181528183208390558483526049905290208190558415610f9e57610f9e333087610bae565b610fa88186611fcc565b600081815260486020908152604091829020825192830189905260608301889052608080845281546002600180831615610100026000190190921604918501829052859433947f6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc99493928401928d928d929091829182019060a0830190889080156110745780601f1061104957610100808354040283529160200191611074565b820191906000526020600020905b81548152906001019060200180831161105757829003601f168201915b50508381038252865460026000196101006001841615020190911604808252602090910190879080156110e85780601f106110bd576101008083540402835291602001916110e8565b820191906000526020600020905b8154815290600101906020018083116110cb57829003601f168201915b5050965050505050505060405180910390a350611119565b600081815260496020526040902054611119908561160c565b505050505050505050565b600061112e6127ce565b60008061113b600061257e565b935093509350935090919293565b60008060008073__$965cfd6d1cee46a73cf1c675b6712824df$__633cedafe59091866040518363ffffffff1660e01b8152600401808381526020018281526020019250505060606040518083038186803b1580156111a757600080fd5b505af41580156111bb573d6000803e3d6000fd5b505050506040513d60608110156111d157600080fd5b508051602082015160409092015190945090925090506111f18382610570565b50610ba8838383610bae565b600073__$6f5c017f31ba759198a3415002c70b4aa5$__63d723552b9091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b15801561129357600080fd5b505af41580156112a7573d6000803e3d6000fd5b505050506112b36127ec565b6040805160a08101909152603560056000835b828210156113075760408051808201909152600283028501805482526001908101546001600160a01b031660208084019190915291835290920191016112c6565b5050604080516517dd985d5b1d60d21b815281519081900360060181206000908152603f6020908152838220547118dd5c9c995b9d135a5b995c94995dd85c9960721b8452845193849003601201909320825283905291909120549394506001600160a01b0316928392509050801561154a5760005b600581101561154857600085826005811061139457fe5b6020020151602001516001600160a01b03161461154057604b546040805163095ea7b360e01b81523060048201526024810185905290516001600160a01b039092169163095ea7b3916044808201926020929091908290030181600087803b1580156113ff57600080fd5b505af1158015611413573d6000803e3d6000fd5b505050506040513d602081101561142957600080fd5b5050604b54604080516323b872dd60e01b81523060048201526001600160a01b03868116602483015260448201869052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561148757600080fd5b505af115801561149b573d6000803e3d6000fd5b505050506040513d60208110156114b157600080fd5b50506001600160a01b0383166347e7ef248683600581106114ce57fe5b602002015160200151846040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561152757600080fd5b505af115801561153b573d6000803e3d6000fd5b505050505b60010161137d565b505b5050604080517118dd5c9c995b9d135a5b995c94995dd85c9960721b815281519081900360120190206000908152602082905290812055505050505050565b604b54604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b1580156115da57600080fd5b505afa1580156115ee573d6000803e3d6000fd5b505050506040513d602081101561160457600080fd5b505192915050565b6000821161161957600080fd5b801561162a5761162a333083610bae565b6116348282611fcc565b60008281526048602090815260408083208151670746f74616c5469760c41b815282519081900360080181208552600490910183529281902054848452918301919091528051849233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e882092918290030190a35050565b600083815260486020908152604080832085845260058101909252909120546090439190910311156116dc57600080fd5b60008381526005820160205260409020546116f657600080fd5b6005821061170357600080fd5b60008381526008820160205260408120836005811061171e57fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607490920183528151918101919091206000818152604a909252919020546001600160a01b039092169250901561179057600080fd5b60408051696469737075746546656560b01b8152815190819003600a0190206000908152602082905220546117c89033903090610bae565b604080516b191a5cdc1d5d1950dbdd5b9d60a21b808252825191829003600c90810183206000908152602085815285822054848652865195869003840186208352868252868320600191820190559385528551948590039092018420815284825284812054868252604a8352858220819055610100808601875287865285840183815286880184815260608801858152608089018681526001600160a01b03808e1660a08c019081523360c08d0190815260e08d018a8152898b526044808d528f8c209e518f5597519c8e019c909c55945160028d01805495519451925184166301000000026301000000600160b81b0319931515620100000262ff000019961515909a0261ff001993151560ff19909816979097179290921695909517939093169690961795909516179055516003880180549184166001600160a01b031992831617905595516004880180549190931696169590951790558551681c995c5d595cdd125960ba1b815286519081900360099081018220845260059687018086528885208f905583855295855268074696d657374616d760bc1b8252875191829003810190912083529383528582208b90558a825292880190915292909220908690811061199357fe5b0154600082815260446020818152604080842081516476616c756560d81b8152825190819003600590810182208752909101808452828620969096558685528383526f6d696e457865637574696f6e4461746560801b81528151908190036010018120855285835281852062093a80420190558685528383526a313637b1b5a73ab6b132b960a91b8152815190819003600b0181208552858352818520439055868552838352681b5a5b995c94db1bdd60ba1b8152815190819003600901812085528583528185208b9055696469737075746546656560b01b8152815190819003600a0181208552818352818520548786529383526266656560e81b815281519081900360030190208452939052919020556002851415611ad65760008781526048602090815260408083208984526007019091529020805460ff191660011790555b6001600160a01b038316600081815260476020908152604091829020600390558151898152908101929092528051899284927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a350505050505050565b604c546001600160a01b031681565b6001600160a01b03821660009081526047602052604081205415611b94576000611b8283611b7686611589565b9063ffffffff61262d16565b10611b8f575060016105f8565b6105f8565b6000611ba383611b7686611589565b106105f8575060016105f8565b600080611bbc33611589565b9050611bca33848303611ecb565b611bd384611589565b9050808382011015611be457600080fd5b611bf084848301611ecb565b604b546040805163a9059cbb60e01b81526001600160a01b038781166004830152602482018790529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015611c4657600080fd5b505af1158015611c5a573d6000803e3d6000fd5b505050506040513d6020811015611c7057600080fd5b5051949350505050565b604080516378bfa27760e01b8152600060048201819052915173__$341c69a3b8ea65d7eeecd190190dea4d1b$__926378bfa2779260248082019391829003018186803b158015611cca57600080fd5b505af4158015611cde573d6000803e3d6000fd5b5050604b54604080516517dd985d5b1d60d21b815281519081900360060181206000908152603f6020908152838220546a1cdd185ad9505b5bdd5b9d60aa1b8452845193849003600b0184208352848252848320546323b872dd60e01b85526001600160a01b039182166004860152336024860152604485015293519390941695506323b872dd945060648083019493928390030190829087803b158015611d8557600080fd5b505af1158015611d99573d6000803e3d6000fd5b505050506040513d6020811015611daf57600080fd5b5050565b60408051631798f51560e31b8152600060048201819052602482018590528315156044830152915173__$965cfd6d1cee46a73cf1c675b6712824df$__9263bcc7a8a89260648082019391829003018186803b158015611e1257600080fd5b505af4158015611e26573d6000803e3d6000fd5b505050505050565b604080516517dd985d5b1d60d21b815281519081900360060181206000908152603f602090815283822054633363e50b60e21b8452935191936001600160a01b0316928392839263cd8f942c926004808201939182900301818987803b158015611e9757600080fd5b505af1158015611eab573d6000803e3d6000fd5b505050506040513d6020811015611ec157600080fd5b5051925050505b90565b6001600160a01b038216600090815260456020526040902080541580611f1757508054439082906000198101908110611f0057fe5b6000918252602090912001546001600160801b0316105b15611f7e5780546000908290611f308260018301612819565b81548110611f3a57fe5b600091825260209091200180546001600160801b03858116600160801b024382166fffffffffffffffffffffffffffffffff19909316929092171617905550611fb8565b805460009082906000198101908110611f9357fe5b600091825260209091200180546001600160801b03808616600160801b029116179055505b505050565b604b546001600160a01b031681565b600082815260486020526040812090611fe48161263f565b905082156120535760408051670746f74616c5469760c41b815281519081900360080190206000908152600484016020522054612027908463ffffffff61269b16565b60408051670746f74616c5469760c41b8152815190819003600801902060009081526004850160205220555b60408051670746f74616c5469760c41b815281519081900360080181206000908152600485016020908152838220546f18dd5c9c995b9d14995c5d595cdd125960821b8452845193849003601001909320825283905291909120546123025760408051670746f74616c5469760c41b815281519081900360080181206000908152600480870160209081528483208390556f18dd5c9c995b9d14995c5d595cdd125960821b8085528551948590036010908101862085528683528685208c90556f63757272656e74546f74616c5469707360801b80875287519687900382018720865287845287862089905585548785018a905287890152600019438101406060808a01919091528951808a03820181526080808b01808d52825192890192909220808b55918790528b519a8b90036090018b208a528b88528b8a205469646966666963756c747960b01b8c528c519b8c9003600a018c208b528c89528c8b2054888d528d519c8d900388018d208c528d8a528d8c20548c526048808b528e8d206a6772616e756c617269747960a81b8f528f519e8f9003600b018f208e52909b018a528d8c2054988d528d519c8d900388018d208c528d8a528d8c20548c529989528c8b20958c528c519b8c90039096018b208a528b8852988b902054818b52968a0188905299890185905296880185905260a096880187815282546002610100600183161502909301169190910496880187905291977f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42979096939491939192909160c0830190859080156122eb5780601f106122c0576101008083540402835291602001916122eb565b820191906000526020600020905b8154815290600101906020018083116122ce57829003601f168201915b5050965050505050505060405180910390a26109f1565b60008281526048602090815260408083208151670746f74616c5469760c41b815282519081900360080190208452600401909152902054811180612344575081155b156124115760028084015460408051602081018390529081018490526060808252865460001961010060018316150201169390930492810183905287927f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c928792909186919081906080820190869080156124005780601f106123d557610100808354040283529160200191612400565b820191906000526020600020905b8154815290600101906020018083116123e357829003601f168201915b505094505050505060405180910390a25b604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b815281519081900360100190206000908152600485016020522054612527576040805161066081019182905260009182916124849160019060339082845b8154815260200190600101908083116124675750505050506126b1565b909250905081831180612495575081155b156125205782600182603381106124a857fe5b01556000818152604360208181526040808420805485526048835281852082516f3932b8bab2b9ba28a837b9b4ba34b7b760811b80825284519182900360109081018320895260049384018752858920899055898952968652928e905591825282519182900390940190208452918801905290208190555b50506109f1565b83156109f157604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b81528151908190036010019020600090815260048501602052205484906001906033811061256f57fe5b01805490910190555050505050565b60006125886127ce565b600080805b60058110156125c3578560350181600581106125a557fe5b60020201548482600581106125b657fe5b602002015260010161258d565b505083546040805169646966666963756c747960b01b8152815190819003600a01812060009081528288016020818152848320546f63757272656e74546f74616c5469707360801b8552855194859003601001909420835252919091205491945091509193509193565b60008282111561263957fe5b50900390565b60408051610660810191829052600091829182916126809190600187019060339082845b8154815260200190600101908083116126635750505050506126e8565b60009081526043909501602052505060409092205492915050565b6000828201838110156126aa57fe5b9392505050565b610640810151603260315b80156126e25760208102840151838110156126d8578093508192505b50600019016126bc565b50915091565b60008060005b60338110156126e257602081028401518084101561270d578093508192505b506001016126ee565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061275757805160ff1916838001178555612784565b82800160010185558215612784579182015b82811115612784578251825591602001919060010190612769565b50612790929150612839565b5090565b8280548282559060005260206000209081019282156127845791602002820182811115612784578251825591602001919060010190612769565b6040518060a001604052806005906020820280388339509192915050565b6040518060a001604052806005905b612803612853565b8152602001906001900390816127fb5790505090565b815481835581811115611fb857600083815260209020611fb89181019083015b611ec891905b80821115612790576000815560010161283f565b60408051808201909152600080825260208201529056fea265627a7a723158202a749c3cd71e08ea223e318d24a041cc39c9c339e7360518d2e7844e4a65004764736f6c63430005100032"

// DeployZap deploys a new Ethereum contract, binding an instance of Zap to it.
func DeployZap(auth *bind.TransactOpts, backend bind.ContractBackend, zapTokenBsc common.Address) (common.Address, *types.Transaction, *Zap, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapStakeAddr, _, _, _ := DeployZapStake(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$341c69a3b8ea65d7eeecd190190dea4d1b$__", zapStakeAddr.String()[2:], -1)

	zapLibraryAddr, _, _, _ := DeployZapLibrary(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$6f5c017f31ba759198a3415002c70b4aa5$__", zapLibraryAddr.String()[2:], -1)

	zapDisputeAddr, _, _, _ := DeployZapDispute(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$965cfd6d1cee46a73cf1c675b6712824df$__", zapDisputeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapBin), backend, zapTokenBsc)
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

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Zap *ZapCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Zap.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Zap *ZapSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _Zap.Contract.AllowedToTrade(&_Zap.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Zap *ZapCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _Zap.Contract.AllowedToTrade(&_Zap.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 balance)
func (_Zap *ZapCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Zap.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
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

// GetBalanceAt is a free data retrieval call binding the contract method 0x26ca8adb.
//
// Solidity: function getBalanceAt(address _user, uint256 _block) view returns(uint256)
func (_Zap *ZapCaller) GetBalanceAt(opts *bind.CallOpts, _user common.Address, _block *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Zap.contract.Call(opts, out, "getBalanceAt", _user, _block)
	return *ret0, err
}

// GetBalanceAt is a free data retrieval call binding the contract method 0x26ca8adb.
//
// Solidity: function getBalanceAt(address _user, uint256 _block) view returns(uint256)
func (_Zap *ZapSession) GetBalanceAt(_user common.Address, _block *big.Int) (*big.Int, error) {
	return _Zap.Contract.GetBalanceAt(&_Zap.CallOpts, _user, _block)
}

// GetBalanceAt is a free data retrieval call binding the contract method 0x26ca8adb.
//
// Solidity: function getBalanceAt(address _user, uint256 _block) view returns(uint256)
func (_Zap *ZapCallerSession) GetBalanceAt(_user common.Address, _block *big.Int) (*big.Int, error) {
	return _Zap.Contract.GetBalanceAt(&_Zap.CallOpts, _user, _block)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Zap *ZapCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	out := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficutly *big.Int
		Tip        *big.Int
	})
	err := _Zap.contract.Call(opts, out, "getNewCurrentVariables")
	return *out, err
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Zap *ZapSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Zap.Contract.GetNewCurrentVariables(&_Zap.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Zap *ZapCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Zap.Contract.GetNewCurrentVariables(&_Zap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zap *ZapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zap.contract.Call(opts, out, "owner")
	return *ret0, err

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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zap.contract.Call(opts, out, "token")
	return *ret0, err
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

// DoTransfer is a paid mutator transaction binding the contract method 0x37751b35.
//
// Solidity: function doTransfer(address _from, address _to, uint256 _amount) returns()
func (_Zap *ZapTransactor) DoTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "doTransfer", _from, _to, _amount)
}

// DoTransfer is a paid mutator transaction binding the contract method 0x37751b35.
//
// Solidity: function doTransfer(address _from, address _to, uint256 _amount) returns()
func (_Zap *ZapSession) DoTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.DoTransfer(&_Zap.TransactOpts, _from, _to, _amount)
}

// DoTransfer is a paid mutator transaction binding the contract method 0x37751b35.
//
// Solidity: function doTransfer(address _from, address _to, uint256 _amount) returns()
func (_Zap *ZapTransactorSession) DoTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.DoTransfer(&_Zap.TransactOpts, _from, _to, _amount)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0xd18ed17b.
//
// Solidity: function increaseVaultApproval() returns(bool)
func (_Zap *ZapTransactor) IncreaseVaultApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "increaseVaultApproval")
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0xd18ed17b.
//
// Solidity: function increaseVaultApproval() returns(bool)
func (_Zap *ZapSession) IncreaseVaultApproval() (*types.Transaction, error) {
	return _Zap.Contract.IncreaseVaultApproval(&_Zap.TransactOpts)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0xd18ed17b.
//
// Solidity: function increaseVaultApproval() returns(bool)
func (_Zap *ZapTransactorSession) IncreaseVaultApproval() (*types.Transaction, error) {
	return _Zap.Contract.IncreaseVaultApproval(&_Zap.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewZapAddress) returns()
func (_Zap *ZapTransactor) ProposeFork(opts *bind.TransactOpts, _propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "proposeFork", _propNewZapAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewZapAddress) returns()
func (_Zap *ZapSession) ProposeFork(_propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewZapAddress) returns()
func (_Zap *ZapTransactorSession) ProposeFork(_propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress)
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
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Zap *ZapSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Zap *ZapSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_Zap *ZapTransactor) UpdateBalanceAtNow(opts *bind.TransactOpts, _user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "updateBalanceAtNow", _user, _value)
}

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_Zap *ZapSession) UpdateBalanceAtNow(_user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.UpdateBalanceAtNow(&_Zap.TransactOpts, _user, _value)
}

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_Zap *ZapTransactorSession) UpdateBalanceAtNow(_user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.UpdateBalanceAtNow(&_Zap.TransactOpts, _user, _value)
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

// ZapTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Zap contract.
type ZapTransferIterator struct {
	Event *ZapTransferEvent // Event containing the contract specifics and raw log

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
func (it *ZapTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTransferEvent)
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
		it.Event = new(ZapTransferEvent)
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
func (it *ZapTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTransfer represents a Transfer event raised by the Zap contract.
type ZapTransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Zap *ZapFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ZapTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ZapTransferIterator{contract: _Zap.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Zap *ZapFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZapTransferEvent, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTransferEvent)
				if err := _Zap.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Zap *ZapFilterer) ParseTransfer(log types.Log) (*ZapTransferEvent, error) {
	event := new(ZapTransferEvent)
	if err := _Zap.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeABI is the input ABI used to generate the binding from.
const ZapDisputeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"}]"

// ZapDisputeFuncSigs maps the 4-byte function signature to its string representation.
var ZapDisputeFuncSigs = map[string]string{
	"804b8931": "proposeFork(ZapStorage.ZapStorageStruct storage,address)",
	"3cedafe5": "tallyVotes(ZapStorage.ZapStorageStruct storage,uint256)",
	"9264b888": "updateDisputeFee(ZapStorage.ZapStorageStruct storage)",
	"bcc7a8a8": "vote(ZapStorage.ZapStorageStruct storage,uint256,bool)",
}

// ZapDisputeBin is the compiled bytecode used for deploying new contracts.
var ZapDisputeBin = "0x610ec0610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c80633cedafe51461005b578063804b8931146100b55780639264b888146100f0578063bcc7a8a81461011a575b600080fd5b81801561006757600080fd5b5061008b6004803603604081101561007e57600080fd5b5080359060200135610152565b604080516001600160a01b0394851681529290931660208301528183015290519081900360600190f35b8180156100c157600080fd5b506100ee600480360360408110156100d857600080fd5b50803590602001356001600160a01b0316610695565b005b8180156100fc57600080fd5b506100ee6004803603602081101561011357600080fd5b5035610ad0565b81801561012657600080fd5b506100ee6004803603606081101561013d57600080fd5b50803590602081013590604001351515610c5b565b600081815260448301602090815260408083208151681c995c5d595cdd125960ba1b815282519081900360090181208552600582018085528386205486526048880185528386206266656560e81b8352845192839003600301909220865290935290832054600282015484938493929091849060ff16156101d257600080fd5b604080516f6d696e457865637574696f6e4461746560801b815281519081900360100190206000908152600586016020522054421161021057600080fd5b600284015462010000900460ff166104ff576002840154630100000090046001600160a01b0316600090815260478a0160205260408120600186015490911215610455576000815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0190206000908152818c016020522080546000190190556102a28a610ad0565b73__$4bd25fb8e309492c3f630e57c98565c1cd$__63a93a4d038b8760020160039054906101000a90046001600160a01b03168860030160009054906101000a90046001600160a01b03168e604001600060405180806a1cdd185ad9505b5bdd5b9d60aa1b815250600b01905060405180910390208152602001908152602001600020546040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b15801561039357600080fd5b505af41580156103a7573d6000803e3d6000fd5b50505060028601805461ff001916610100179055506040805168074696d657374616d760bc1b815281519081900360090190206000908152600587016020908152828220548252600787019052205460ff16151560011415610440576040805168074696d657374616d760bc1b815281519081900360090190206000908152600587016020908152828220548252600687019052908120555b60038501546001600160a01b031691506104f9565b60018082556040805168074696d657374616d760bc1b815281519081900360090190206000908152600588016020908152828220548252600788019052205460ff16151514156104e1576040805168074696d657374616d760bc1b81528151908190036009019020600090815260058701602090815282822054825260078701905220805460ff191690555b6002850154630100000090046001600160a01b031691505b5061060d565b60008460010154131561060d57604080516b746f74616c5f737570706c7960a01b8152815190819003600c0190206000908152818b016020522054606490601402604080516571756f72756d60d01b8152815190819003600601902060009081526005880160205220549190041061057657600080fd5b600484018054604080516a1e985c10dbdb9d1c9858dd60aa1b8152815190819003600b0181206000908152603f8e0160209081529083902080546001600160a01b0319166001600160a01b0395861617905560028901805461ff00191661010017905593549092168252517f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270929181900390910190a15b60028401805460ff19166001908117918290558501546003860154604080519283526001600160a01b039182166020840152610100840460ff1615158382015251630100000090930416918a917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a33096509450925050509250925092565b604080516bffffffffffffffffffffffff19606084901b1660208083019190915282518083036014018152603490920183528151918101919091206000818152604a860190925291902054156106ea57600080fd5b60408051696469737075746546656560b01b8152815190819003600a01812060009081528286016020528281205463a93a4d0360e01b8352600483018790523360248401523060448401526064830152915173__$4bd25fb8e309492c3f630e57c98565c1cd$__9263a93a4d039260848082019391829003018186803b15801561077357600080fd5b505af4158015610787573d6000803e3d6000fd5b5050505082604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002060008154809291906001019190505550600083604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508084604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600115158152602001336001600160a01b03168152602001336001600160a01b03168152602001846001600160a01b0316815250846044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555090505043846044016000838152602001908152602001600020600501600060405180806a313637b1b5a73ab6b132b960a91b815250600b01905060405180910390208152602001908152602001600020819055508360400160006040518080696469737075746546656560b01b815250600a0190506040518091039020815260200190815260200160002054846044016000838152602001908152602001600020600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020819055504262093a8001846044016000838152602001908152602001600020600501600060405180806f6d696e457865637574696f6e4461746560801b8152506010019050604051809103902081526020019081526020016000208190555050505050565b604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546103e89190820281610b3557fe5b041015610c2857604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b019094208352529190912054610bf791600f916103e891610bea91830281610bab57fe5b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b019020600090815281890160205220549190046103e80363ffffffff610e4e16565b81610bf157fe5b04610e75565b60408051696469737075746546656560b01b8152815190819003600a01902060009081528184016020522055610c58565b60408051696469737075746546656560b01b8152815190819003600a019020600090815281830160205220600f90555b50565b600082815260448085016020908152604080842081516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260058201845282862054630637bf7f60e51b8252600482018a905233602483015294810194909452905190939273__$4bd25fb8e309492c3f630e57c98565c1cd$__9263c6f7efe092606480840193829003018186803b158015610cf257600080fd5b505af4158015610d06573d6000803e3d6000fd5b505050506040513d6020811015610d1c57600080fd5b505133600090815260068401602052604090205490915060ff16151560011415610d4557600080fd5b60008111610d5257600080fd5b33600090815260478601602052604090205460031415610d7157600080fd5b336000908152600680840160209081526040808420805460ff1916600190811790915581516c6e756d6265724f66566f74657360981b8152825190819003600d01812086526005880180855283872080549093019092556571756f72756d60d01b8152825190819003909401909320845291905290208054820190558215610e025760018201805482019055610e0e565b60018201805482900390555b6040805184151581529051339186917f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae09181900360200190a35050505050565b6000828202831580610e68575082848281610e6557fe5b04145b610e6e57fe5b9392505050565b6000818311610e845781610e6e565b509091905056fea265627a7a723158209d005d2485dc9f078f5f4c6acb56a257c865c35a81fc1862a7bfa7788119273964736f6c63430005100032"

// DeployZapDispute deploys a new Ethereum contract, binding an instance of ZapDispute to it.
func DeployZapDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapTransferAddr, _, _, _ := DeployZapTransfer(auth, backend)
	ZapDisputeBin = strings.Replace(ZapDisputeBin, "__$4bd25fb8e309492c3f630e57c98565c1cd$__", zapTransferAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapDisputeBin), backend)
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

// ZapGettersABI is the input ABI used to generate the binding from.
const ZapGettersABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zapTokenBsc\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZapGettersFuncSigs maps the 4-byte function signature to its string representation.
var ZapGettersFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
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
	"17d7de7c": "getName()",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"1db842f0": "getRequestIdByQueryHash(bytes32)",
	"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
	"0f0b424d": "getRequestIdByTimestamp(uint256)",
	"b5413029": "getRequestQ()",
	"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
	"e1eee6d6": "getRequestVars(uint256)",
	"733bdef0": "getStakerInfo(address)",
	"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
	"15070401": "getSymbol()",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"612c8f7f": "getUintVar(bytes32)",
	"19e8e03b": "getVariablesOnDeck()",
	"3df0777b": "isInDispute(uint256,uint256)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"1ca8b6cb": "totalTokenSupply()",
}

// ZapGettersBin is the compiled bytecode used for deploying new contracts.
var ZapGettersBin = "0x608060405234801561001057600080fd5b506040516119d83803806119d88339818101604052602081101561003357600080fd5b5051604b80546001600160a01b0319166001600160a01b03909216919091179055611975806100636000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806370a082311161010f578063af0b1327116100a2578063dd62ed3e11610071578063dd62ed3e146107c3578063e0ae93c1146107f1578063e1eee6d614610814578063fc7cf0a01461092b576101e5565b8063af0b1327146106c1578063b541302914610765578063c775b54214610783578063da379941146107a6576101e5565b806393fa4915116100de57806393fa4915146105a3578063999cf26c146105c6578063a22e407a146105f2578063a7c438bc14610695576101e5565b806370a08231146104f8578063733bdef01461051e57806377fbb6631461055d5780637f6fd5d914610580576101e5565b80631db842f011610187578063612c8f7f11610156578063612c8f7f1461046f5780636173c0b81461048c57806363bb82ad146104a957806369026d63146104d5576101e5565b80631db842f0146103c85780633180f8df146103e55780633df0777b1461041b57806346eee1c414610452576101e5565b806315070401116101c357806315070401146102ad57806317d7de7c1461032a57806319e8e03b146103325780631ca8b6cb146103c0576101e5565b80630f0b424d146101ea57806311c9851214610219578063133bee5e14610274575b600080fd5b6102076004803603602081101561020057600080fd5b5035610933565b60408051918252519081900360200190f35b61023c6004803603604081101561022f57600080fd5b508035906020013561094b565b604051808260a080838360005b83811015610261578181015183820152602001610249565b5050505090500191505060405180910390f35b6102916004803603602081101561028a57600080fd5b503561096c565b604080516001600160a01b039092168252519081900360200190f35b6102b561097e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102ef5781810151838201526020016102d7565b50505050905090810190601f16801561031c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102b561098f565b61033a61099b565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561038357818101518382015260200161036b565b50505050905090810190601f1680156103b05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6102076109b5565b610207600480360360208110156103de57600080fd5b50356109c1565b610402600480360360208110156103fb57600080fd5b50356109d3565b6040805192835290151560208301528051918290030190f35b61043e6004803603604081101561043157600080fd5b50803590602001356109ef565b604080519115158252519081900360200190f35b6102076004803603602081101561046857600080fd5b5035610a02565b6102076004803603602081101561048557600080fd5b5035610a14565b610207600480360360208110156104a257600080fd5b5035610a26565b61043e600480360360408110156104bf57600080fd5b50803590602001356001600160a01b0316610a38565b61023c600480360360408110156104eb57600080fd5b5080359060200135610a4b565b6102076004803603602081101561050e57600080fd5b50356001600160a01b0316610a65565b6105446004803603602081101561053457600080fd5b50356001600160a01b0316610ae8565b6040805192835260208301919091528051918290030190f35b6102076004803603604081101561057357600080fd5b5080359060200135610afb565b6102076004803603604081101561059657600080fd5b5080359060200135610b0e565b610207600480360360408110156105b957600080fd5b5080359060200135610b21565b61043e600480360360408110156105dc57600080fd5b506001600160a01b038135169060200135610b34565b6105fa610bd1565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561065557818101518382015260200161063d565b50505050905090810190601f1680156106825780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b61043e600480360360408110156106ab57600080fd5b50803590602001356001600160a01b0316610bf8565b6106de600480360360208110156106d757600080fd5b5035610c0b565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b8381101561074457818101518382015260200161072c565b50505050905001828152602001995050505050505050505060405180910390f35b61076d610c4f565b6040518151815280826106608083836020610249565b6102076004803603604081101561079957600080fd5b5080359060200135610c61565b610207600480360360208110156107bc57600080fd5b5035610c74565b610207600480360360408110156107d957600080fd5b506001600160a01b0381358116916020013516610c86565b6102076004803603604081101561080757600080fd5b5080359060200135610cf3565b6108316004803603602081101561082a57600080fd5b5035610d06565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b8381101561088a578181015183820152602001610872565b50505050905090810190601f1680156108b75780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b838110156108ea5781810151838201526020016108d2565b50505050905090810190601f1680156109175780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b610402610d32565b6000610945818363ffffffff610d4716565b92915050565b6109536118e4565b6109656000848463ffffffff610d5d16565b9392505050565b6000610945818363ffffffff610db716565b606061098a6000610dd6565b905090565b606061098a6000610df5565b60008060606109aa6000610e19565b925092509250909192565b600061098a6000610f05565b6000610945818363ffffffff610f3816565b6000806109e6818463ffffffff610f4e16565b91509150915091565b600061096581848463ffffffff610fb416565b6000610945818363ffffffff610fdb16565b6000610945818363ffffffff610ff416565b6000610945818363ffffffff61100616565b600061096581848463ffffffff61102d16565b610a536118e4565b6109656000848463ffffffff61105a16565b604b54604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b158015610ab657600080fd5b505afa158015610aca573d6000803e3d6000fd5b505050506040513d6020811015610ae057600080fd5b505192915050565b6000806109e6818463ffffffff6110be16565b600061096581848463ffffffff6110e516565b600061096581848463ffffffff61111816565b600061096581848463ffffffff61113c16565b6040805163b9290ca560e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$4bd25fb8e309492c3f630e57c98565c1cd$__9163b9290ca5916064808301926020929190829003018186803b158015610b9e57600080fd5b505af4158015610bb2573d6000803e3d6000fd5b505050506040513d6020811015610bc857600080fd5b50519392505050565b60008060006060600080610be56000611160565b949b939a50919850965094509092509050565b600061096581848463ffffffff61131e16565b6000806000806000806000610c1e611902565b6000610c30818b63ffffffff61135016565b9850985098509850985098509850985098509193959799909294969850565b610c57611921565b61098a6000611565565b600061096581848463ffffffff6115a516565b6000610945818363ffffffff6115c916565b604b5460408051636eb1769f60e11b81526001600160a01b03858116600483015284811660248301529151600093929092169163dd62ed3e91604480820192602092909190829003018186803b158015610cdf57600080fd5b505afa158015610bb2573d6000803e3d6000fd5b600061096581848463ffffffff6115df16565b6060806000808080610d1e818863ffffffff61160316565b949c939b5091995097509550909350915050565b600080610d3f60006117de565b915091509091565b6000908152604291909101602052604090205490565b610d656118e4565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610d9657505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b506040805180820190915260048152632d20a82160e11b602082015290565b5060408051808201909152600981526805a61702042455032360bc1b602082015290565b60008060606000610e2985611854565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f810184900484028501840190925281845294955085949291839190830182828015610ef05780601f10610ec557610100808354040283529160200191610ef0565b820191906000526020600020905b815481529060010190602001808311610ed357829003601f168201915b50505050509050935093509350509193909250565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b6000908152604991909101602052604090205490565b60008181526048830160205260408120600381015482919015610fa457600381018054610f989187918791906000198101908110610f8857fe5b906000526020600020015461113c565b60019250925050610fad565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60009081526040918201602052205490565b6000603282111561101657600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b6110626118e4565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161109357505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b6000828152604884016020526040812060030180548390811061110457fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156113025780601f106112d757610100808354040283529160200191611302565b820191906000526020600020905b8154815290600101906020018083116112e557829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b6000806000806000806000611363611902565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b61156d611921565b6040805161066081019182905290600184019060339082845b8154815260200190600101908083116115865750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a999098899889988998919788979388019692959294909188918301828280156117325780601f1061170757610100808354040283529160200191611732565b820191906000526020600020905b81548152906001019060200180831161171557829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a9450925084019050828280156117c05780601f10611795576101008083540402835291602001916117c0565b820191906000526020600020905b8154815290600101906020018083116117a357829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b808252825160129281900383018120600090815285850160208181528683205483526042880181528683205494845286519384900390950190922081529252918120549091829161184a91859161113c565b9360019350915050565b60408051610660810191829052600091829182916118959190600187019060339082845b8154815260200190600101908083116118785750505050506118b0565b60009081526043909501602052505060409092205492915050565b60008060005b60338110156118de5760208102840151808410156118d5578093508192505b506001016118b6565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a723158200402ceb694e4df80893a07180549c04c68ead2322525f20ca2f5cfba1ede8d4164736f6c63430005100032"

// DeployZapGetters deploys a new Ethereum contract, binding an instance of ZapGetters to it.
func DeployZapGetters(auth *bind.TransactOpts, backend bind.ContractBackend, zapTokenBsc common.Address) (common.Address, *types.Transaction, *ZapGetters, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapTransferAddr, _, _, _ := DeployZapTransfer(auth, backend)
	ZapGettersBin = strings.Replace(ZapGettersBin, "__$4bd25fb8e309492c3f630e57c98565c1cd$__", zapTransferAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapGettersBin), backend, zapTokenBsc)
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
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

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ZapGetters *ZapGettersCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ZapGetters *ZapGettersSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapGetters.Contract.AllowedToTrade(&_ZapGetters.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ZapGetters *ZapGettersCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapGetters.Contract.AllowedToTrade(&_ZapGetters.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
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
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _ZapGetters.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapGetters.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapGetters.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapGetters.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
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
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
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

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapGetters *ZapGettersCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapGetters *ZapGettersSession) GetName() (string, error) {
	return _ZapGetters.Contract.GetName(&_ZapGetters.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapGetters *ZapGettersCallerSession) GetName() (string, error) {
	return _ZapGetters.Contract.GetName(&_ZapGetters.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getNewVauleCountbyRequestId", _requestId)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestByQueryHash", _request)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err

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
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
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
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapGetters.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapGetters.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
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
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
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

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapGetters *ZapGettersCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapGetters *ZapGettersSession) GetSymbol() (string, error) {
	return _ZapGetters.Contract.GetSymbol(&_ZapGetters.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapGetters *ZapGettersCallerSession) GetSymbol() (string, error) {
	return _ZapGetters.Contract.GetSymbol(&_ZapGetters.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ZapGetters.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "totalSupply")
	return *ret0, err
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

// ZapGettersLibraryABI is the input ABI used to generate the binding from.
const ZapGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"}]"

// ZapGettersLibraryBin is the compiled bytecode used for deploying new contracts.
var ZapGettersLibraryBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158203fb103180a6c34adaffd950841f4f3a9cf81b436db90edfce002b37c9534295564736f6c63430005100032"

// DeployZapGettersLibrary deploys a new Ethereum contract, binding an instance of ZapGettersLibrary to it.
func DeployZapGettersLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapGettersLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapGettersLibraryBin), backend)
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

// ZapLibraryABI is the input ABI used to generate the binding from.
const ZapLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"}]"

// ZapLibraryFuncSigs maps the 4-byte function signature to its string representation.
var ZapLibraryFuncSigs = map[string]string{
	"d723552b": "submitMiningSolution(ZapStorage.ZapStorageStruct storage,string,uint256,uint256)",
}

// ZapLibraryBin is the compiled bytecode used for deploying new contracts.
var ZapLibraryBin = "0x61199a610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063d723552b1461003a575b600080fd5b81801561004657600080fd5b506100f96004803603608081101561005d57600080fd5b8135919081019060408101602082013564010000000081111561007f57600080fd5b82018360208201111561009157600080fd5b803590602001918460018302840111640100000000831117156100b357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356100fb565b005b33600090815260478501602052604090205460011461011957600080fd5b604080516f18dd5c9c995b9d14995c5d595cdd125960821b8152815190819003601001902060009081528186016020522054821461015657600080fd5b836040016000604051808069646966666963756c747960b01b815250600a0190506040518091039020815260200190815260200160002054600260038660000154338760405160200180848152602001836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b602083106101ec5780518252601f1990920191602091820191016101cd565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106102775780518252601f199092019160209182019101610258565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156102b6573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b6020831061032c5780518252601f19909201916020918201910161030d565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561036b573d6000803e3d6000fd5b5050506040513d602081101561038057600080fd5b50518161038957fe5b061561039457600080fd5b83546000908152604185016020908152604080832033845290915290205460ff16156103bf57600080fd5b604080517118dd5c9c995b9d135a5b995c94995dd85c9960721b81528151908190036012018120600090815282870160208181528483208390556b736c6f7450726f677265737360a01b8452845193849003600c0190932082529091522054819060358601906005811061042f57fe5b6002020155604080516b736c6f7450726f677265737360a01b8152815190819003600c01902060009081528186016020522054339060358601906005811061047357fe5b60020201600190810180546001600160a01b0319166001600160a01b039390931692909217909155604080516b736c6f7450726f677265737360a01b81528151600c91819003919091018120600090815287830160209081528382208054860190558854825260418901815283822033808452908252848320805460ff19169096179095558854818401879052938301849052606080845288519084015287518795947fe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4948a9489949293919283926080840192918801918190849084905b8381101561056a578181015183820152602001610552565b50505050905090810190601f1680156105975780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a3604080516b736c6f7450726f677265737360a01b8152815190819003600c01902060009081528186016020522054600514156105e7576105e78484846105ed565b50505050565b6000818152604884016020908152604080832081517174696d654f664c6173744e657756616c756560701b81528251601291819003919091018120855282880180855283862054691d1a5b5955185c99d95d60b21b83528451600a938190038401812088528287528588205469646966666963756c747960b01b808352875192839003860183208a52848952878a20549083528751928390039095019091208852919095529285205491946064429590950390930302929092059091019081136106e5576040805169646966666963756c747960b01b8152815190819003600a01902060009081528187016020522060019055610714565b6040805169646966666963756c747960b01b8152815190819003600a0190206000908152818701602052208190555b603c42604080517174696d654f664c6173744e657756616c756560701b815281519081900360120190206000908152818901602052209190064203905561075961184e565b6040805160a081019091526035870160056000835b828210156107af5760408051808201909152600283028501805482526001908101546001600160a01b0316602080840191909152918352909201910161076e565b509293506001925050505b60058110156108d95760008282600581106107d157fe5b602002015151905060008383600581106107e757fe5b602002015160200151905060008390505b60008111801561081b575084600182036005811061081257fe5b60200201515183105b1561088d5784600182036005811061082f57fe5b60200201515185826005811061084157fe5b6020020151528460001982016005811061085757fe5b60200201516020015185826005811061086c57fe5b602090810291909101516001600160a01b03909216910152600019016107f8565b838110156108ce57828582600581106108a257fe5b602002015152818582600581106108b557fe5b602090810291909101516001600160a01b039092169101525b5050506001016107ba565b604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d0190206000908152818901602052205461094657604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d0190206000908152818901602052206753444835ec58000090555b604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d01902060009081528189016020522054670de0b6b3a76400001015610a5157604080516c18dd5c9c995b9d14995dd85c99609a1b8082528251600d928190038301812060009081528b850160208181528683205485855287519485900387018520845282825287842054868652885195869003880186208552838352888520670de0b6b3a7640000651bd78f205bc69093028390049091039055948452865193849003909501909220815292529190205460649160329104026040805167646576536861726560c01b815281519081900360080190206000908152818b01602052209190049055610a8b565b604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d019020600090815281890160205220670de0b6b3a764000090555b50604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d01812060009081528883016020818152848320546f63757272656e74546f74616c5469707360801b855285519485900360100185208452828252858420547118dd5c9c995b9d135a5b995c94995dd85c9960721b86528651958690036012019095208452919052928120670de0b6b3a7640000909304600590920482019092555b6005821015610c415773__$4bd25fb8e309492c3f630e57c98565c1cd$__63a93a4d038930868660058110610b5d57fe5b60200201516020015160058d604001600060405180806f63757272656e74546f74616c5469707360801b8152506010019050604051809103902081526020019081526020016000205481610bad57fe5b0486016040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b158015610c1d57600080fd5b505af4158015610c31573d6000803e3d6000fd5b505060019093019250610b2c9050565b604080517174696d654f664c6173744e657756616c756560701b8152815190819003601201812060009081528a830160208181528483205488860151516f63757272656e74546f74616c5469707360801b80875287519687900360109081018820875285855288872054918852885197889003018720865293835293869020548e549186529185019390935260059091069003828401526060820152905187917fc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16916080918190039190910190a2604080516b746f74616c5f737570706c7960a01b8152815190819003600c0181206000908152828b0160208181528483208054610113019055652fb7bbb732b960d11b845284519384900360060184208352603f8d0181528483205467646576536861726560c01b8552855194859003600801852084529190528382205463a93a4d0360e01b8452600484018d90523060248501526001600160a01b0390911660448401526064830152915173__$4bd25fb8e309492c3f630e57c98565c1cd$__9263a93a4d039260848082019391829003018186803b158015610df257600080fd5b505af4158015610e06573d6000803e3d6000fd5b5050505082600260058110610e1757fe5b6020908102919091015151604080517174696d654f664c6173744e657756616c756560701b80825282516012928190038301812060009081528e850180885285822054825260068d018852858220969096558282528451918290038401822081528587528481205460038d018054600181018255908352888320015560a08201855289518701516001600160a01b0390811683528a8801518801518116838901528a8601518801518116838701526060808c01518901518216908401526080808c0151890151909116908301528451928352845192839003909301909120825292845281812054815260088901909352909120610f1591600561187b565b506040805160a081018252845151815260208086015151818301528583015151828401526060808701515190830152608080870151519083015282517174696d654f664c6173744e657756616c756560701b8152835190819003601201902060009081528b840182528381205481526009890190915291909120610f9a9160056118d3565b50604080517174696d654f664c6173744e657756616c756560701b808252825191829003601290810183206000908152848d01602081815286832054835260058c01815286832043905584865286519586900384018620835281815286832054835260428f0181528683208d9055938552855194859003909201842081528183528481205460348e01805460018101825590835284832001556b736c6f7450726f677265737360a01b8452845193849003600c019093208352905290812055611062886117be565b604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282516010928190038301812060009081528d8501602081815286832097909755928252845191829003909301902082529092529020541561177b57604080516f18dd5c9c995b9d14995c5d595cdd125960821b808252825191829003601090810183206000908152848d01602081815286832054835260488f01808252878420670746f74616c5469760c41b88528851978890036008018820855260049081018352888520546f63757272656e74546f74616c5469707360801b8952895198899003870189208652848452898620559587528751968790038501872084529181528683205483529081528582206f3932b8bab2b9ba28a837b9b4ba34b7b760811b8652865195869003909301909420815291019091529081205460018a0190603381106111a657fe5b018190555060008860430160008a60480160008c604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000208190555060008860480160008a604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b8152506010019050604051809103902081526020019081526020016000208190555060008860480160008a604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b81525060100190506040518091039020815260200190815260200160002054815260200190815260200160002060040160006040518080670746f74616c5469760c41b81525060080190506040518091039020815260200190815260200160002081905550600061137f896117be565b905087896000015460014303406040516020018084805190602001908083835b602083106113be5780518252601f19909201916020918201910161139f565b6001836020036101000a038019825116818451168082178552505050505050905001838152602001828152602001935050505060405160208183030381529060405280519060200120896000018190555088604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f428a600001548b6040016000604051808069646966666963756c747960b01b815250600a01905060405180910390208152602001908152602001600020548c60480160008e604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806a6772616e756c617269747960a81b815250600b01905060405180910390208152602001908152602001600020548d60480160008f604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000206000018e604001600060405180806f63757272656e74546f74616c5469707360801b8152506010019050604051809103902081526020019081526020016000205460405180868152602001858152602001848152602001806020018381526020018281038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156116695780601f1061163e57610100808354040283529160200191611669565b820191906000526020600020905b81548152906001019060200180831161164c57829003601f168201915b5050965050505050505060405180910390a2600081815260488a01602090815260408083206002808201548351670746f74616c5469760c41b81528451908190036008018120875260048401865295849020549486018190529285018490526060808652825460001961010060018316150201169190910490850181905285947f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c949293929181906080820190869080156117655780601f1061173a57610100808354040283529160200191611765565b820191906000526020600020905b81548152906001019060200180831161174857829003601f168201915b505094505050505060405180910390a2506117b4565b604080516f63757272656e74546f74616c5469707360801b815281519081900360100190206000908152818a0160205290812081905588555b5050505050505050565b60408051610660810191829052600091829182916117ff9190600187019060339082845b8154815260200190600101908083116117e257505050505061181a565b60009081526043909501602052505060409092205492915050565b60008060005b603381101561184857602081028401518084101561183f578093508192505b50600101611820565b50915091565b6040518060a001604052806005905b61186561190d565b81526020019060019003908161185d5790505090565b82600581019282156118c3579160200282015b828111156118c357825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061188e565b506118cf929150611924565b5090565b8260058101928215611901579160200282015b828111156119015782518255916020019190600101906118e6565b506118cf92915061194b565b604080518082019091526000808252602082015290565b61194891905b808211156118cf5780546001600160a01b031916815560010161192a565b90565b61194891905b808211156118cf576000815560010161195156fea265627a7a72315820e273142c448afc22b99e9e7f62bd1255629f534d23e3574c418c50b39566559364736f6c63430005100032"

// DeployZapLibrary deploys a new Ethereum contract, binding an instance of ZapLibrary to it.
func DeployZapLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapTransferAddr, _, _, _ := DeployZapTransfer(auth, backend)
	ZapLibraryBin = strings.Replace(ZapLibraryBin, "__$4bd25fb8e309492c3f630e57c98565c1cd$__", zapTransferAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapLibraryBin), backend)
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

// ZapMasterABI is the input ABI used to generate the binding from.
const ZapMasterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultContract\",\"type\":\"address\"}],\"name\":\"changeVaultContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"changeZapContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZapMasterFuncSigs maps the 4-byte function signature to its string representation.
var ZapMasterFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"70a08231": "balanceOf(address)",
	"47abd7f1": "changeDeity(address)",
	"af2a5102": "changeVaultContract(address)",
	"e4203f66": "changeZapContract(address)",
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
	"17d7de7c": "getName()",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"1db842f0": "getRequestIdByQueryHash(bytes32)",
	"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
	"0f0b424d": "getRequestIdByTimestamp(uint256)",
	"b5413029": "getRequestQ()",
	"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
	"e1eee6d6": "getRequestVars(uint256)",
	"733bdef0": "getStakerInfo(address)",
	"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
	"15070401": "getSymbol()",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"612c8f7f": "getUintVar(bytes32)",
	"19e8e03b": "getVariablesOnDeck()",
	"3df0777b": "isInDispute(uint256,uint256)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"1ca8b6cb": "totalTokenSupply()",
}

// ZapMasterBin is the compiled bytecode used for deploying new contracts.
var ZapMasterBin = "0x608060405234801561001057600080fd5b506040516120283803806120288339818101604052604081101561003357600080fd5b508051602090910151604b80546001600160a01b0319166001600160a01b038316179055604080516347b024eb60e01b8152600060048201819052915173__$341c69a3b8ea65d7eeecd190190dea4d1b$__926347b024eb9260248082019391829003018186803b1580156100a757600080fd5b505af41580156100bb573d6000803e3d6000fd5b505060408051652fb7bbb732b960d11b8152815190819003600690810182206000908152603f602081815285832080546001600160a01b031990811633908117909255655f646569747960d01b8752875196879003909501862084528282528684208054861690911790556a1e985c10dbdb9d1c9858dd60aa1b8552855194859003600b018520835281815285832080546001600160a01b038c811691871682179092556f1e985c151bdad95b90dbdb9d1c9858dd60821b875287519687900360100187208552928252928690208054938a16939094169290921790925590825291517f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab2709450908190039091019150a15050611e4c806101dc6000396000f3fe6080604052600436106101f95760003560e01c8063733bdef01161010d578063af2a5102116100a0578063dd62ed3e1161006f578063dd62ed3e14610a2c578063e0ae93c114610a67578063e1eee6d614610a97578063e4203f6614610bbb578063fc7cf0a014610bee576101f9565b8063af2a510214610974578063b5413029146109a7578063c775b542146109d2578063da37994114610a02576101f9565b8063999cf26c116100dc578063999cf26c146107a1578063a22e407a146107da578063a7c438bc1461088a578063af0b1327146108c3576101f9565b8063733bdef0146106c557806377fbb663146107115780637f6fd5d91461074157806393fa491514610771576101f9565b80633180f8df11610190578063612c8f7f1161015f578063612c8f7f146105d55780636173c0b8146105ff57806363bb82ad1461062957806369026d631461066257806370a0823114610692576101f9565b80633180f8df146104ef5780633df0777b1461053257806346eee1c41461057657806347abd7f1146105a0576101f9565b806317d7de7c116101cc57806317d7de7c1461040057806319e8e03b146104155780631ca8b6cb146104b05780631db842f0146104c5576101f9565b80630f0b424d1461028c57806311c98512146102c8578063133bee5e146103305780631507040114610376575b604080516a1e985c10dbdb9d1c9858dd60aa1b8152815190819003600b0181206000908152603f602090815283822054601f369081018390048302850183019095528484526001600160a01b03169360609392918190840183828082843760009201829052508451949550938493509150506020840185600019f43d604051816000823e828015610288578282f35b8282fd5b34801561029857600080fd5b506102b6600480360360208110156102af57600080fd5b5035610c03565b60408051918252519081900360200190f35b3480156102d457600080fd5b506102f8600480360360408110156102eb57600080fd5b5080359060200135610c1b565b604051808260a080838360005b8381101561031d578181015183820152602001610305565b5050505090500191505060405180910390f35b34801561033c57600080fd5b5061035a6004803603602081101561035357600080fd5b5035610c3c565b604080516001600160a01b039092168252519081900360200190f35b34801561038257600080fd5b5061038b610c4e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103c55781810151838201526020016103ad565b50505050905090810190601f1680156103f25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040c57600080fd5b5061038b610c5f565b34801561042157600080fd5b5061042a610c6b565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561047357818101518382015260200161045b565b50505050905090810190601f1680156104a05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156104bc57600080fd5b506102b6610c85565b3480156104d157600080fd5b506102b6600480360360208110156104e857600080fd5b5035610c91565b3480156104fb57600080fd5b506105196004803603602081101561051257600080fd5b5035610ca3565b6040805192835290151560208301528051918290030190f35b34801561053e57600080fd5b506105626004803603604081101561055557600080fd5b5080359060200135610cbf565b604080519115158252519081900360200190f35b34801561058257600080fd5b506102b66004803603602081101561059957600080fd5b5035610cd2565b3480156105ac57600080fd5b506105d3600480360360208110156105c357600080fd5b50356001600160a01b0316610ce4565b005b3480156105e157600080fd5b506102b6600480360360208110156105f857600080fd5b5035610cf8565b34801561060b57600080fd5b506102b66004803603602081101561062257600080fd5b5035610d0a565b34801561063557600080fd5b506105626004803603604081101561064c57600080fd5b50803590602001356001600160a01b0316610d1c565b34801561066e57600080fd5b506102f86004803603604081101561068557600080fd5b5080359060200135610d2f565b34801561069e57600080fd5b506102b6600480360360208110156106b557600080fd5b50356001600160a01b0316610d49565b3480156106d157600080fd5b506106f8600480360360208110156106e857600080fd5b50356001600160a01b0316610dcc565b6040805192835260208301919091528051918290030190f35b34801561071d57600080fd5b506102b66004803603604081101561073457600080fd5b5080359060200135610ddf565b34801561074d57600080fd5b506102b66004803603604081101561076457600080fd5b5080359060200135610df2565b34801561077d57600080fd5b506102b66004803603604081101561079457600080fd5b5080359060200135610e05565b3480156107ad57600080fd5b50610562600480360360408110156107c457600080fd5b506001600160a01b038135169060200135610e18565b3480156107e657600080fd5b506107ef610eb5565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561084a578181015183820152602001610832565b50505050905090810190601f1680156108775780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b34801561089657600080fd5b50610562600480360360408110156108ad57600080fd5b50803590602001356001600160a01b0316610edc565b3480156108cf57600080fd5b506108ed600480360360208110156108e657600080fd5b5035610eef565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b8381101561095357818101518382015260200161093b565b50505050905001828152602001995050505050505050505060405180910390f35b34801561098057600080fd5b506105d36004803603602081101561099757600080fd5b50356001600160a01b0316610f33565b3480156109b357600080fd5b506109bc610f44565b6040518151815280826106608083836020610305565b3480156109de57600080fd5b506102b6600480360360408110156109f557600080fd5b5080359060200135610f56565b348015610a0e57600080fd5b506102b660048036036020811015610a2557600080fd5b5035610f69565b348015610a3857600080fd5b506102b660048036036040811015610a4f57600080fd5b506001600160a01b0381358116916020013516610f7b565b348015610a7357600080fd5b506102b660048036036040811015610a8a57600080fd5b5080359060200135610fe8565b348015610aa357600080fd5b50610ac160048036036020811015610aba57600080fd5b5035610ffb565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b83811015610b1a578181015183820152602001610b02565b50505050905090810190601f168015610b475780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610b7a578181015183820152602001610b62565b50505050905090810190601f168015610ba75780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015610bc757600080fd5b506105d360048036036020811015610bde57600080fd5b50356001600160a01b0316611027565b348015610bfa57600080fd5b50610519611038565b6000610c15818363ffffffff61104d16565b92915050565b610c23611dbb565b610c356000848463ffffffff61106316565b9392505050565b6000610c15818363ffffffff6110bd16565b6060610c5a60006110dc565b905090565b6060610c5a60006110fb565b6000806060610c7a600061111f565b925092509250909192565b6000610c5a600061120b565b6000610c15818363ffffffff61123e16565b600080610cb6818463ffffffff61125416565b91509150915091565b6000610c3581848463ffffffff6112ba16565b6000610c15818363ffffffff6112e116565b610cf560008263ffffffff6112fa16565b50565b6000610c15818363ffffffff61138316565b6000610c15818363ffffffff61139516565b6000610c3581848463ffffffff6113bc16565b610d37611dbb565b610c356000848463ffffffff6113e916565b604b54604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b158015610d9a57600080fd5b505afa158015610dae573d6000803e3d6000fd5b505050506040513d6020811015610dc457600080fd5b505192915050565b600080610cb6818463ffffffff61144d16565b6000610c3581848463ffffffff61147416565b6000610c3581848463ffffffff6114a716565b6000610c3581848463ffffffff6114cb16565b6040805163b9290ca560e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$4bd25fb8e309492c3f630e57c98565c1cd$__9163b9290ca5916064808301926020929190829003018186803b158015610e8257600080fd5b505af4158015610e96573d6000803e3d6000fd5b505050506040513d6020811015610eac57600080fd5b50519392505050565b60008060006060600080610ec960006114ef565b949b939a50919850965094509092509050565b6000610c3581848463ffffffff6116ad16565b6000806000806000806000610f02611dd9565b6000610f14818b63ffffffff6116df16565b9850985098509850985098509850985098509193959799909294969850565b610cf560008263ffffffff6118f416565b610f4c611df8565b610c5a600061197d565b6000610c3581848463ffffffff6119bd16565b6000610c15818363ffffffff6119e116565b604b5460408051636eb1769f60e11b81526001600160a01b03858116600483015284811660248301529151600093929092169163dd62ed3e91604480820192602092909190829003018186803b158015610fd457600080fd5b505afa158015610e96573d6000803e3d6000fd5b6000610c3581848463ffffffff6119f716565b6060806000808080611013818863ffffffff611a1b16565b949c939b5091995097509550909350915050565b610cf560008263ffffffff611bf616565b6000806110456000611cb5565b915091509091565b6000908152604291909101602052604090205490565b61106b611dbb565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b81548152602001906001019080831161109c57505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b506040805180820190915260048152632d20a82160e11b602082015290565b5060408051808201909152600981526805a61702042455032360bc1b602082015290565b6000806060600061112f85611d2b565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f8101849004840285018401909252818452949550859492918391908301828280156111f65780601f106111cb576101008083540402835291602001916111f6565b820191906000526020600020905b8154815290600101906020018083116111d957829003601f168201915b50505050509050935093509350509193909250565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b6000908152604991909101602052604090205490565b600081815260488301602052604081206003810154829190156112aa5760038101805461129e918791879190600019810190811061128e57fe5b90600052602060002001546114cb565b600192509250506112b3565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b0316331461133757600080fd5b60408051655f646569747960d01b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b0319909216919091179055565b60009081526040918201602052205490565b600060328211156113a557600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b6113f1611dbb565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161142257505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b6000828152604884016020526040812060030180548390811061149357fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156116915780601f1061166657610100808354040283529160200191611691565b820191906000526020600020905b81548152906001019060200180831161167457829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60008060008060008060006116f2611dd9565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b60408051652fb7bbb732b960d11b815281519081900360060190206000908152603f840160205220546001600160a01b0316331461193157600080fd5b604080516517dd985d5b1d60d21b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b0319909216919091179055565b611985611df8565b6040805161066081019182905290600184019060339082845b81548152602001906001019080831161199e5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a99909889988998899891978897938801969295929490918891830182828015611b4a5780601f10611b1f57610100808354040283529160200191611b4a565b820191906000526020600020905b815481529060010190602001808311611b2d57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a945092508401905082828015611bd85780601f10611bad57610100808354040283529160200191611bd8565b820191906000526020600020905b815481529060010190602001808311611bbb57829003601f168201915b50505050509450965096509650965096509650509295509295509295565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b03163314611c3357600080fd5b604080516a1e985c10dbdb9d1c9858dd60aa1b8152815190819003600b0181206000908152603f850160209081529083902080546001600160a01b0386166001600160a01b03199091168117909155825291517f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270929181900390910190a15050565b604080517174696d654f664c6173744e657756616c756560701b8082528251601292819003830181206000908152858501602081815286832054835260428801815286832054948452865193849003909501909220815292529181205490918291611d219185916114cb565b9360019350915050565b6040805161066081019182905260009182918291611d6c9190600187019060339082845b815481526020019060010190808311611d4f575050505050611d87565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611db5576020810284015180841015611dac578093508192505b50600101611d8d565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a72315820600b4a417988138ff97d0aaf176ae062821169fbb54eb42fe3311dd380de119764736f6c63430005100032"

// DeployZapMaster deploys a new Ethereum contract, binding an instance of ZapMaster to it.
func DeployZapMaster(auth *bind.TransactOpts, backend bind.ContractBackend, _zapContract common.Address, tokenAddress common.Address) (common.Address, *types.Transaction, *ZapMaster, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapMasterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapStakeAddr, _, _, _ := DeployZapStake(auth, backend)
	ZapMasterBin = strings.Replace(ZapMasterBin, "__$341c69a3b8ea65d7eeecd190190dea4d1b$__", zapStakeAddr.String()[2:], -1)

	zapTransferAddr, _, _, _ := DeployZapTransfer(auth, backend)
	ZapMasterBin = strings.Replace(ZapMasterBin, "__$4bd25fb8e309492c3f630e57c98565c1cd$__", zapTransferAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapMasterBin), backend, _zapContract, tokenAddress)
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
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

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ZapMaster *ZapMasterCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ZapMaster *ZapMasterSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapMaster.Contract.AllowedToTrade(&_ZapMaster.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ZapMaster *ZapMasterCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapMaster.Contract.AllowedToTrade(&_ZapMaster.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
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
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _ZapMaster.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapMaster.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapMaster.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapMaster.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
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
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
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

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapMaster *ZapMasterCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapMaster *ZapMasterSession) GetName() (string, error) {
	return _ZapMaster.Contract.GetName(&_ZapMaster.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapMaster *ZapMasterCallerSession) GetName() (string, error) {
	return _ZapMaster.Contract.GetName(&_ZapMaster.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getNewValueCountbyRequestId", _requestId)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestIdByQueryHash", _request)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err
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
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
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
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapMaster.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapMaster.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
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
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
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

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapMaster *ZapMasterCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapMaster *ZapMasterSession) GetSymbol() (string, error) {
	return _ZapMaster.Contract.GetSymbol(&_ZapMaster.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapMaster *ZapMasterCallerSession) GetSymbol() (string, error) {
	return _ZapMaster.Contract.GetSymbol(&_ZapMaster.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ZapMaster.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
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

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "totalTokenSupply")
	return *ret0, err
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

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ZapMaster *ZapMasterSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeDeity(&_ZapMaster.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeDeity(&_ZapMaster.TransactOpts, _newDeity)
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

// ChangeZapContract is a paid mutator transaction binding the contract method 0xe4203f66.
//
// Solidity: function changeZapContract(address _zapContract) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeZapContract(opts *bind.TransactOpts, _zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeZapContract", _zapContract)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xe4203f66.
//
// Solidity: function changeZapContract(address _zapContract) returns()
func (_ZapMaster *ZapMasterSession) ChangeZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeZapContract(&_ZapMaster.TransactOpts, _zapContract)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xe4203f66.
//
// Solidity: function changeZapContract(address _zapContract) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeZapContract(&_ZapMaster.TransactOpts, _zapContract)
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

// ZapStakeABI is the input ABI used to generate the binding from.
const ZapStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"}]"

// ZapStakeFuncSigs maps the 4-byte function signature to its string representation.
var ZapStakeFuncSigs = map[string]string{
	"326991a3": "depositStake(ZapStorage.ZapStorageStruct storage)",
	"47b024eb": "init(ZapStorage.ZapStorageStruct storage)",
	"3c734827": "requestStakingWithdraw(ZapStorage.ZapStorageStruct storage)",
	"78bfa277": "withdrawStake(ZapStorage.ZapStorageStruct storage)",
}

// ZapStakeBin is the compiled bytecode used for deploying new contracts.
var ZapStakeBin = "0x610884610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c8063326991a31461005b5780633c7348271461008757806347b024eb146100b157806378bfa277146100db575b600080fd5b81801561006757600080fd5b506100856004803603602081101561007e57600080fd5b5035610105565b005b81801561009357600080fd5b50610085600480360360208110156100aa57600080fd5b5035610179565b8180156100bd57600080fd5b50610085600480360360208110156100d457600080fd5b5035610270565b8180156100e757600080fd5b50610085600480360360208110156100fe57600080fd5b5035610634565b61010f81336106a2565b73__$965cfd6d1cee46a73cf1c675b6712824df$__639264b888826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561015e57600080fd5b505af4158015610172573d6000803e3d6000fd5b5050505050565b3360009081526047820160205260409020805460011461019857600080fd5b6002815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0181206000908152828501602052828120805460001901905563124c971160e31b825260048201859052915173__$965cfd6d1cee46a73cf1c675b6712824df$__92639264b8889260248082019391829003018186803b15801561022957600080fd5b505af415801561023d573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b6040805167646563696d616c7360c01b8152815190819003600801902060009081528183016020522054156102a457600080fd5b30600090815260458201602052604080822081516305bfb12160e31b81526004810191909152629896806024820152905173__$4bd25fb8e309492c3f630e57c98565c1cd$__92632dfd89089260448082019391829003018186803b15801561030c57600080fd5b505af4158015610320573d6000803e3d6000fd5b5050505061032c610831565b506040805160c08101825273f39fd6e51aad88f6f4ce6ab8827279cfffb92266815273cd3b766ccdd6ae721141f452c550ca635964ce716020820152732546bcd3c84621e976d8185a91a922ae77ecec309181019190915273bda5747bfd65f08deb54cb465eb87d40e51b197e606082015273dd2fd4581271e230360230f9337d5c0430bf44c06080820152738626f6940e2eb28930efb4cef49b2d1f2c9c119960a082015260005b60068110156104a85773__$4bd25fb8e309492c3f630e57c98565c1cd$__632dfd890884604501600085856006811061040a57fe5b60200201516001600160a01b03166001600160a01b031681526020019081526020016000206207a1206040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561046e57600080fd5b505af4158015610482573d6000803e3d6000fd5b505050506104a08383836006811061049657fe5b60200201516106a2565b6001016103d5565b50604080516b746f74616c5f737570706c7960a01b8152815190819003600c9081018220600090815283860160208181528583208054622dc6c001905567646563696d616c7360c01b855285519485900360080185208352818152858320601290556b7461726765744d696e65727360a01b85528551948590039093018420825280835284822060c890556a1cdd185ad9505b5bdd5b9d60aa1b8452845193849003600b01842082528083528482206207a1209055696469737075746546656560b01b8452845193849003600a908101852083528184528583206103ca9055691d1a5b5955185c99d95d60b21b80865286519586900382018620845282855286842061025890558552855194859003019093208152919052205442816105ca57fe5b604080517174696d654f664c6173744e657756616c756560701b815281519081900360120181206000908152958201602081815283882095909406420390945569646966666963756c747960b01b8152815190819003600a01902085529190529091206001905550565b3360009081526047820160205260409020600181015462093a8090620151804206420303101561066357600080fd5b805460021461067157600080fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6001600160a01b038116600090815260478301602052604090205415806106e357506001600160a01b03811660009081526047830160205260409020546002145b6106ec57600080fd5b604080516a1cdd185ad95c90dbdd5b9d60aa1b81528151600b918190038201812060009081528584016020818152858320805460019081019091558487018752808552620151804290810690038286019081526001600160a01b038916855260478a018352878520955186555194019390935560458701835284822085516a1cdd185ad9505b5bdd5b9d60aa1b8152865190819003909501852083529252838120546305bfb12160e31b845260048401929092526024830191909152915173__$4bd25fb8e309492c3f630e57c98565c1cd$__92632dfd89089260448082019391829003018186803b1580156107e157600080fd5b505af41580156107f5573d6000803e3d6000fd5b50506040516001600160a01b03841692507f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e29150600090a25050565b6040518060c00160405280600690602082028038833950919291505056fea265627a7a72315820dca7e5cea8d646727cbfbc2a5668503098a757b258ee93a44fd1edb838e745dc64736f6c63430005100032"

// DeployZapStake deploys a new Ethereum contract, binding an instance of ZapStake to it.
func DeployZapStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStake, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapTransferAddr, _, _, _ := DeployZapTransfer(auth, backend)
	ZapStakeBin = strings.Replace(ZapStakeBin, "__$4bd25fb8e309492c3f630e57c98565c1cd$__", zapTransferAddr.String()[2:], -1)

	zapDisputeAddr, _, _, _ := DeployZapDispute(auth, backend)
	ZapStakeBin = strings.Replace(ZapStakeBin, "__$965cfd6d1cee46a73cf1c675b6712824df$__", zapDisputeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapStakeBin), backend)
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

// ZapStorageABI is the input ABI used to generate the binding from.
const ZapStorageABI = "[]"

// ZapStorageBin is the compiled bytecode used for deploying new contracts.
var ZapStorageBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158207645f4e57c8b377a841d1518626850c6dabbc5ece038b74563628da20c13ba0664736f6c63430005100032"

// DeployZapStorage deploys a new Ethereum contract, binding an instance of ZapStorage to it.
func DeployZapStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapStorageBin), backend)
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

// ZapTokenBSCABI is the input ABI used to generate the binding from.
const ZapTokenBSCABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"allocate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZapTokenBSCFuncSigs maps the 4-byte function signature to its string representation.
var ZapTokenBSCFuncSigs = map[string]string{
	"b78b52df": "allocate(address,uint256)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"66188463": "decreaseApproval(address,uint256)",
	"7d64bcb4": "finishMinting()",
	"893d20e8": "getOwner()",
	"d73dd623": "increaseApproval(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"05d2035b": "mintingFinished()",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// ZapTokenBSCBin is the compiled bytecode used for deploying new contracts.
var ZapTokenBSCBin = "0x6a52b7d2dcc80cd2e40000006000556003805460ff60a01b1916905560c0604052600960808190526805a61702042455032360bc1b60a090815261004691600491906100ed565b50604080518082019091526004808252632d20a82160e11b6020909201918252610072916005916100ed565b506006805460ff1916601217905534801561008c57600080fd5b50600380546001600160a01b03191633908117909155600080548282526001602090815260408084208390558051928352517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a3610188565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012e57805160ff191683800117855561015b565b8280016001018555821561015b579182015b8281111561015b578251825591602001919060010190610140565b5061016792915061016b565b5090565b61018591905b808211156101675760008155600101610171565b90565b610b79806101976000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80637d64bcb4116100a2578063a9059cbb11610071578063a9059cbb14610308578063b78b52df14610334578063d73dd62314610362578063dd62ed3e1461038e578063f2fde38b146103bc57610116565b80637d64bcb4146102cc578063893d20e8146102d45780638da5cb5b146102f857806395d89b411461030057610116565b806323b872dd116100e957806323b872dd146101fa578063313ce5671461023057806340c10f191461024e578063661884631461027a57806370a08231146102a657610116565b806305d2035b1461011b57806306fdde0314610137578063095ea7b3146101b457806318160ddd146101e0575b600080fd5b6101236103e2565b604080519115158252519081900360200190f35b61013f6103f2565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610179578181015183820152602001610161565b50505050905090810190601f1680156101a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610123600480360360408110156101ca57600080fd5b506001600160a01b038135169060200135610480565b6101e86104e6565b60408051918252519081900360200190f35b6101236004803603606081101561021057600080fd5b506001600160a01b038135811691602081013590911690604001356104ec565b61023861060a565b6040805160ff9092168252519081900360200190f35b6101236004803603604081101561026457600080fd5b506001600160a01b038135169060200135610613565b6101236004803603604081101561029057600080fd5b506001600160a01b03813516906020013561071e565b6101e8600480360360208110156102bc57600080fd5b50356001600160a01b031661080e565b610123610829565b6102dc610885565b604080516001600160a01b039092168252519081900360200190f35b6102dc610894565b61013f6108a3565b6101236004803603604081101561031e57600080fd5b506001600160a01b0381351690602001356108fe565b6103606004803603604081101561034a57600080fd5b506001600160a01b0381351690602001356109c3565b005b6101236004803603604081101561037857600080fd5b506001600160a01b0381351690602001356109d2565b6101e8600480360360408110156103a457600080fd5b506001600160a01b0381358116916020013516610a6b565b610360600480360360208110156103d257600080fd5b50356001600160a01b0316610a96565b600354600160a01b900460ff1681565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104785780601f1061044d57610100808354040283529160200191610478565b820191906000526020600020905b81548152906001019060200180831161045b57829003601f168201915b505050505081565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60006001600160a01b03831661050157600080fd5b6001600160a01b03841660008181526002602090815260408083203384528252808320549383526001909152902054610540908463ffffffff610b1c16565b6001600160a01b038087166000908152600160205260408082209390935590861681522054610575908463ffffffff610b2e16565b6001600160a01b03851660009081526001602052604090205561059e818463ffffffff610b1c16565b6001600160a01b03808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b60065460ff1681565b6003546000906001600160a01b0316331461062d57600080fd5b600354600160a01b900460ff161561064457600080fd5b600054610657908363ffffffff610b2e16565b60009081556001600160a01b038416815260016020526040902054610682908363ffffffff610b2e16565b6001600160a01b038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a26040805183815290516001600160a01b038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600192915050565b3360009081526002602090815260408083206001600160a01b038616845290915281205480831115610773573360009081526002602090815260408083206001600160a01b03881684529091528120556107a8565b610783818463ffffffff610b1c16565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526001602052604090205490565b6003546000906001600160a01b0316331461084357600080fd5b6003805460ff60a01b1916600160a01b1790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b6003546001600160a01b031690565b6003546001600160a01b031681565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104785780601f1061044d57610100808354040283529160200191610478565b60006001600160a01b03831661091357600080fd5b33600090815260016020526040902054610933908363ffffffff610b1c16565b33600090815260016020526040808220929092556001600160a01b03851681522054610965908363ffffffff610b2e16565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6109cd8282610613565b505050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610a06908363ffffffff610b2e16565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6003546001600160a01b03163314610aad57600080fd5b6001600160a01b038116610ac057600080fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b600082821115610b2857fe5b50900390565b600082820183811015610b3d57fe5b939250505056fea265627a7a723158202203c5ca61aa4fffb4df251c495d89294cbe34802b8f07643edefda78650fc5564736f6c63430005100032"

// DeployZapTokenBSC deploys a new Ethereum contract, binding an instance of ZapTokenBSC to it.
func DeployZapTokenBSC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapTokenBSC, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTokenBSCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapTokenBSCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapTokenBSC{ZapTokenBSCCaller: ZapTokenBSCCaller{contract: contract}, ZapTokenBSCTransactor: ZapTokenBSCTransactor{contract: contract}, ZapTokenBSCFilterer: ZapTokenBSCFilterer{contract: contract}}, nil
}

// ZapTokenBSC is an auto generated Go binding around an Ethereum contract.
type ZapTokenBSC struct {
	ZapTokenBSCCaller     // Read-only binding to the contract
	ZapTokenBSCTransactor // Write-only binding to the contract
	ZapTokenBSCFilterer   // Log filterer for contract events
}

// ZapTokenBSCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapTokenBSCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTokenBSCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTokenBSCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTokenBSCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapTokenBSCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTokenBSCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapTokenBSCSession struct {
	Contract     *ZapTokenBSC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapTokenBSCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapTokenBSCCallerSession struct {
	Contract *ZapTokenBSCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZapTokenBSCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTokenBSCTransactorSession struct {
	Contract     *ZapTokenBSCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZapTokenBSCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapTokenBSCRaw struct {
	Contract *ZapTokenBSC // Generic contract binding to access the raw methods on
}

// ZapTokenBSCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapTokenBSCCallerRaw struct {
	Contract *ZapTokenBSCCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTokenBSCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTokenBSCTransactorRaw struct {
	Contract *ZapTokenBSCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapTokenBSC creates a new instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSC(address common.Address, backend bind.ContractBackend) (*ZapTokenBSC, error) {
	contract, err := bindZapTokenBSC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSC{ZapTokenBSCCaller: ZapTokenBSCCaller{contract: contract}, ZapTokenBSCTransactor: ZapTokenBSCTransactor{contract: contract}, ZapTokenBSCFilterer: ZapTokenBSCFilterer{contract: contract}}, nil
}

// NewZapTokenBSCCaller creates a new read-only instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSCCaller(address common.Address, caller bind.ContractCaller) (*ZapTokenBSCCaller, error) {
	contract, err := bindZapTokenBSC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCCaller{contract: contract}, nil
}

// NewZapTokenBSCTransactor creates a new write-only instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSCTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTokenBSCTransactor, error) {
	contract, err := bindZapTokenBSC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCTransactor{contract: contract}, nil
}

// NewZapTokenBSCFilterer creates a new log filterer instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSCFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapTokenBSCFilterer, error) {
	contract, err := bindZapTokenBSC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCFilterer{contract: contract}, nil
}

// bindZapTokenBSC binds a generic wrapper to an already deployed contract.
func bindZapTokenBSC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTokenBSCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTokenBSC *ZapTokenBSCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTokenBSC.Contract.ZapTokenBSCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTokenBSC *ZapTokenBSCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.ZapTokenBSCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTokenBSC *ZapTokenBSCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.ZapTokenBSCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTokenBSC *ZapTokenBSCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTokenBSC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTokenBSC *ZapTokenBSCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTokenBSC *ZapTokenBSCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ZapTokenBSC *ZapTokenBSCCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ZapTokenBSC *ZapTokenBSCSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.Allowance(&_ZapTokenBSC.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.Allowance(&_ZapTokenBSC.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ZapTokenBSC *ZapTokenBSCCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ZapTokenBSC *ZapTokenBSCSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.BalanceOf(&_ZapTokenBSC.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.BalanceOf(&_ZapTokenBSC.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZapTokenBSC *ZapTokenBSCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZapTokenBSC *ZapTokenBSCSession) Decimals() (uint8, error) {
	return _ZapTokenBSC.Contract.Decimals(&_ZapTokenBSC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Decimals() (uint8, error) {
	return _ZapTokenBSC.Contract.Decimals(&_ZapTokenBSC.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCSession) GetOwner() (common.Address, error) {
	return _ZapTokenBSC.Contract.GetOwner(&_ZapTokenBSC.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) GetOwner() (common.Address, error) {
	return _ZapTokenBSC.Contract.GetOwner(&_ZapTokenBSC.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_ZapTokenBSC *ZapTokenBSCCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) MintingFinished() (bool, error) {
	return _ZapTokenBSC.Contract.MintingFinished(&_ZapTokenBSC.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) MintingFinished() (bool, error) {
	return _ZapTokenBSC.Contract.MintingFinished(&_ZapTokenBSC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCSession) Name() (string, error) {
	return _ZapTokenBSC.Contract.Name(&_ZapTokenBSC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Name() (string, error) {
	return _ZapTokenBSC.Contract.Name(&_ZapTokenBSC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCSession) Owner() (common.Address, error) {
	return _ZapTokenBSC.Contract.Owner(&_ZapTokenBSC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Owner() (common.Address, error) {
	return _ZapTokenBSC.Contract.Owner(&_ZapTokenBSC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCSession) Symbol() (string, error) {
	return _ZapTokenBSC.Contract.Symbol(&_ZapTokenBSC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Symbol() (string, error) {
	return _ZapTokenBSC.Contract.Symbol(&_ZapTokenBSC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZapTokenBSC *ZapTokenBSCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapTokenBSC.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZapTokenBSC *ZapTokenBSCSession) TotalSupply() (*big.Int, error) {
	return _ZapTokenBSC.Contract.TotalSupply(&_ZapTokenBSC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) TotalSupply() (*big.Int, error) {
	return _ZapTokenBSC.Contract.TotalSupply(&_ZapTokenBSC.CallOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xb78b52df.
//
// Solidity: function allocate(address to, uint256 amount) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactor) Allocate(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "allocate", to, amount)
}

// Allocate is a paid mutator transaction binding the contract method 0xb78b52df.
//
// Solidity: function allocate(address to, uint256 amount) returns()
func (_ZapTokenBSC *ZapTokenBSCSession) Allocate(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Allocate(&_ZapTokenBSC.TransactOpts, to, amount)
}

// Allocate is a paid mutator transaction binding the contract method 0xb78b52df.
//
// Solidity: function allocate(address to, uint256 amount) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Allocate(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Allocate(&_ZapTokenBSC.TransactOpts, to, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Approve(&_ZapTokenBSC.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Approve(&_ZapTokenBSC.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.DecreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.DecreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) FinishMinting() (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.FinishMinting(&_ZapTokenBSC.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.FinishMinting(&_ZapTokenBSC.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.IncreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.IncreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Mint(&_ZapTokenBSC.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Mint(&_ZapTokenBSC.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Transfer(&_ZapTokenBSC.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Transfer(&_ZapTokenBSC.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferFrom(&_ZapTokenBSC.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferFrom(&_ZapTokenBSC.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZapTokenBSC *ZapTokenBSCSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferOwnership(&_ZapTokenBSC.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferOwnership(&_ZapTokenBSC.TransactOpts, newOwner)
}

// ZapTokenBSCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZapTokenBSC contract.
type ZapTokenBSCApprovalIterator struct {
	Event *ZapTokenBSCApproval // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCApproval)
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
		it.Event = new(ZapTokenBSCApproval)
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
func (it *ZapTokenBSCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCApproval represents a Approval event raised by the ZapTokenBSC contract.
type ZapTokenBSCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZapTokenBSCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCApprovalIterator{contract: _ZapTokenBSC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCApproval)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseApproval(log types.Log) (*ZapTokenBSCApproval, error) {
	event := new(ZapTokenBSCApproval)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ZapTokenBSC contract.
type ZapTokenBSCMintIterator struct {
	Event *ZapTokenBSCMint // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCMint)
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
		it.Event = new(ZapTokenBSCMint)
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
func (it *ZapTokenBSCMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCMint represents a Mint event raised by the ZapTokenBSC contract.
type ZapTokenBSCMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*ZapTokenBSCMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCMintIterator{contract: _ZapTokenBSC.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCMint)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseMint(log types.Log) (*ZapTokenBSCMint, error) {
	event := new(ZapTokenBSCMint)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the ZapTokenBSC contract.
type ZapTokenBSCMintFinishedIterator struct {
	Event *ZapTokenBSCMintFinished // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCMintFinished)
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
		it.Event = new(ZapTokenBSCMintFinished)
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
func (it *ZapTokenBSCMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCMintFinished represents a MintFinished event raised by the ZapTokenBSC contract.
type ZapTokenBSCMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterMintFinished(opts *bind.FilterOpts) (*ZapTokenBSCMintFinishedIterator, error) {

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCMintFinishedIterator{contract: _ZapTokenBSC.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCMintFinished) (event.Subscription, error) {

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCMintFinished)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// ParseMintFinished is a log parse operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseMintFinished(log types.Log) (*ZapTokenBSCMintFinished, error) {
	event := new(ZapTokenBSCMintFinished)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "MintFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZapTokenBSC contract.
type ZapTokenBSCOwnershipTransferredIterator struct {
	Event *ZapTokenBSCOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCOwnershipTransferred)
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
		it.Event = new(ZapTokenBSCOwnershipTransferred)
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
func (it *ZapTokenBSCOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCOwnershipTransferred represents a OwnershipTransferred event raised by the ZapTokenBSC contract.
type ZapTokenBSCOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZapTokenBSCOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCOwnershipTransferredIterator{contract: _ZapTokenBSC.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCOwnershipTransferred)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseOwnershipTransferred(log types.Log) (*ZapTokenBSCOwnershipTransferred, error) {
	event := new(ZapTokenBSCOwnershipTransferred)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZapTokenBSC contract.
type ZapTokenBSCTransferIterator struct {
	Event *ZapTokenBSCTransfer // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCTransfer)
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
		it.Event = new(ZapTokenBSCTransfer)
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
func (it *ZapTokenBSCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCTransfer represents a Transfer event raised by the ZapTokenBSC contract.
type ZapTokenBSCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZapTokenBSCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCTransferIterator{contract: _ZapTokenBSC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCTransfer)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseTransfer(log types.Log) (*ZapTokenBSCTransfer, error) {
	event := new(ZapTokenBSCTransfer)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTransferABI is the input ABI used to generate the binding from.
const ZapTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ZapTransferFuncSigs maps the 4-byte function signature to its string representation.
var ZapTransferFuncSigs = map[string]string{
	"fade3342": "allowance(ZapStorage.ZapStorageStruct storage,address,address)",
	"b9290ca5": "allowedToTrade(ZapStorage.ZapStorageStruct storage,address,uint256)",
	"c6f7efe0": "balanceOfAt(ZapStorage.ZapStorageStruct storage,address,uint256)",
	"a93a4d03": "doTransfer(ZapStorage.ZapStorageStruct storage,address,address,uint256)",
	"82c1fecb": "getBalanceAt(ZapStorage.Checkpoint[] storage,uint256)",
	"2dfd8908": "updateBalanceAtNow(ZapStorage.Checkpoint[] storage,uint256)",
}

// ZapTransferBin is the compiled bytecode used for deploying new contracts.
var ZapTransferBin = "0x6106af610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c80632dfd89081461007157806382c1fecb146100a3578063a93a4d03146100d8578063b9290ca514610121578063c6f7efe014610167578063fade334214610199575b600080fd5b81801561007d57600080fd5b506100a16004803603604081101561009457600080fd5b50803590602001356101cd565b005b6100c6600480360360408110156100b957600080fd5b50803590602001356102a6565b60408051918252519081900360200190f35b8180156100e457600080fd5b506100a1600480360360808110156100fb57600080fd5b508035906001600160a01b036020820135811691604081013590911690606001356103d8565b6101536004803603606081101561013757600080fd5b508035906001600160a01b0360208201351690604001356104e3565b604080519115158252519081900360200190f35b6100c66004803603606081101561017d57600080fd5b508035906001600160a01b03602082013516906040013561055b565b6100c6600480360360608110156101af57600080fd5b508035906001600160a01b03602082013581169160400135166105f1565b81541580610201575081544390839060001981019081106101ea57fe5b6000918252602090912001546001600160801b0316105b15610268578154600090839061021a8260018301610630565b8154811061022457fe5b600091825260209091200180546001600160801b03848116600160801b024382166fffffffffffffffffffffffffffffffff199093169290921716179055506102a2565b81546000908390600019810190811061027d57fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b5050565b81546000906102b7575060006103d2565b8254839060001981019081106102c957fe5b6000918252602090912001546001600160801b03168210610319578254839060001981019081106102f657fe5b600091825260209091200154600160801b90046001600160801b031690506103d2565b8260008154811061032657fe5b6000918252602090912001546001600160801b031682101561034a575060006103d2565b8254600090600019015b818111156103a557600060026001838501010490508486828154811061037657fe5b6000918252602090912001546001600160801b0316116103985780925061039f565b6001810391505b50610354565b8482815481106103b157fe5b600091825260209091200154600160801b90046001600160801b0316925050505b92915050565b600081116103e557600080fd5b6001600160a01b0382166103f857600080fd5b6104038484836104e3565b61040c57600080fd5b600061041985854361055b565b6001600160a01b03851660009081526045870160205260409020909150610442908383036101cd565b61044d85844361055b565b905080828201101561045e57600080fd5b6001600160a01b03831660009081526045860160205260409020610484908284016101cd565b61048f85844361055b565b9050826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b6001600160a01b0382166000908152604784016020526040812054156105325760006105208361051487874361055b565b9063ffffffff61061e16565b1061052d57506001610554565b610550565b60006105438361051487874361055b565b1061055057506001610554565b5060005b9392505050565b6001600160a01b038216600090815260458401602052604081205415806105b957506001600160a01b0383166000908152604585016020526040812080548492906105a257fe5b6000918252602090912001546001600160801b0316115b156105c657506000610554565b6001600160a01b038316600090815260458501602052604090206105ea90836102a6565b9050610554565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b60008282111561062a57fe5b50900390565b81548183558181111561065457600083815260209020610654918101908301610659565b505050565b61067791905b80821115610673576000815560010161065f565b5090565b9056fea265627a7a723158204805f51316edc81d40547772b4189c6300fbfc88482516f938d346df6da2c42964736f6c63430005100032"

// DeployZapTransfer deploys a new Ethereum contract, binding an instance of ZapTransfer to it.
func DeployZapTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapTransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapTransfer{ZapTransferCaller: ZapTransferCaller{contract: contract}, ZapTransferTransactor: ZapTransferTransactor{contract: contract}, ZapTransferFilterer: ZapTransferFilterer{contract: contract}}, nil
}

// ZapTransfer is an auto generated Go binding around an Ethereum contract.
type ZapTransfer struct {
	ZapTransferCaller     // Read-only binding to the contract
	ZapTransferTransactor // Write-only binding to the contract
	ZapTransferFilterer   // Log filterer for contract events
}

// ZapTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapTransferSession struct {
	Contract     *ZapTransfer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapTransferCallerSession struct {
	Contract *ZapTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZapTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTransferTransactorSession struct {
	Contract     *ZapTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZapTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapTransferRaw struct {
	Contract *ZapTransfer // Generic contract binding to access the raw methods on
}

// ZapTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapTransferCallerRaw struct {
	Contract *ZapTransferCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTransferTransactorRaw struct {
	Contract *ZapTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapTransfer creates a new instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransfer(address common.Address, backend bind.ContractBackend) (*ZapTransfer, error) {
	contract, err := bindZapTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapTransfer{ZapTransferCaller: ZapTransferCaller{contract: contract}, ZapTransferTransactor: ZapTransferTransactor{contract: contract}, ZapTransferFilterer: ZapTransferFilterer{contract: contract}}, nil
}

// NewZapTransferCaller creates a new read-only instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferCaller(address common.Address, caller bind.ContractCaller) (*ZapTransferCaller, error) {
	contract, err := bindZapTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransferCaller{contract: contract}, nil
}

// NewZapTransferTransactor creates a new write-only instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTransferTransactor, error) {
	contract, err := bindZapTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransferTransactor{contract: contract}, nil
}

// NewZapTransferFilterer creates a new log filterer instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapTransferFilterer, error) {
	contract, err := bindZapTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapTransferFilterer{contract: contract}, nil
}

// bindZapTransfer binds a generic wrapper to an already deployed contract.
func bindZapTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTransfer *ZapTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTransfer.Contract.ZapTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTransfer *ZapTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTransfer.Contract.ZapTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTransfer *ZapTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTransfer.Contract.ZapTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTransfer *ZapTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTransfer *ZapTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTransfer *ZapTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTransfer.Contract.contract.Transact(opts, method, params...)
}

// ZapTransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZapTransfer contract.
type ZapTransferApprovalIterator struct {
	Event *ZapTransferApproval // Event containing the contract specifics and raw log

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
func (it *ZapTransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTransferApproval)
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
		it.Event = new(ZapTransferApproval)
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
func (it *ZapTransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTransferApproval represents a Approval event raised by the ZapTransfer contract.
type ZapTransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ZapTransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ZapTransfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZapTransferApprovalIterator{contract: _ZapTransfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZapTransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ZapTransfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTransferApproval)
				if err := _ZapTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ZapTransfer *ZapTransferFilterer) ParseApproval(log types.Log) (*ZapTransferApproval, error) {
	event := new(ZapTransferApproval)
	if err := _ZapTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTransferTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZapTransfer contract.
type ZapTransferTransferIterator struct {
	Event *ZapTransferTransfer // Event containing the contract specifics and raw log

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
func (it *ZapTransferTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTransferTransfer)
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
		it.Event = new(ZapTransferTransfer)
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
func (it *ZapTransferTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTransferTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTransferTransfer represents a Transfer event raised by the ZapTransfer contract.
type ZapTransferTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ZapTransferTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ZapTransfer.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ZapTransferTransferIterator{contract: _ZapTransfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZapTransferTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ZapTransfer.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTransferTransfer)
				if err := _ZapTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) ParseTransfer(log types.Log) (*ZapTransferTransfer, error) {
	event := new(ZapTransferTransfer)
	if err := _ZapTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

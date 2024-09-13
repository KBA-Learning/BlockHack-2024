// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

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
	_ = abi.ConvertType
)

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Stored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"collection\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506107ae8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063dc642f3014610038578063f641090c14610068575b5f80fd5b610052600480360381019061004d91906102ff565b610084565b60405161005f91906103c0565b60405180910390f35b610082600480360381019061007d91906103e0565b610136565b005b5f818051602081018201805184825260208301602085012081835280955050505050505f9150905080546100b790610483565b80601f01602080910402602001604051908101604052809291908181526020018280546100e390610483565b801561012e5780601f106101055761010080835404028352916020019161012e565b820191905f5260205f20905b81548152906001019060200180831161011157829003601f168201915b505050505081565b805f8360405161014691906104ed565b9081526020016040518091039020908161016091906106a9565b508160405161016f91906104ed565b60405180910390207f05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99826040516101a691906103c0565b60405180910390a25050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610211826101cb565b810181811067ffffffffffffffff821117156102305761022f6101db565b5b80604052505050565b5f6102426101b2565b905061024e8282610208565b919050565b5f67ffffffffffffffff82111561026d5761026c6101db565b5b610276826101cb565b9050602081019050919050565b828183375f83830152505050565b5f6102a361029e84610253565b610239565b9050828152602081018484840111156102bf576102be6101c7565b5b6102ca848285610283565b509392505050565b5f82601f8301126102e6576102e56101c3565b5b81356102f6848260208601610291565b91505092915050565b5f60208284031215610314576103136101bb565b5b5f82013567ffffffffffffffff811115610331576103306101bf565b5b61033d848285016102d2565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561037d578082015181840152602081019050610362565b5f8484015250505050565b5f61039282610346565b61039c8185610350565b93506103ac818560208601610360565b6103b5816101cb565b840191505092915050565b5f6020820190508181035f8301526103d88184610388565b905092915050565b5f80604083850312156103f6576103f56101bb565b5b5f83013567ffffffffffffffff811115610413576104126101bf565b5b61041f858286016102d2565b925050602083013567ffffffffffffffff8111156104405761043f6101bf565b5b61044c858286016102d2565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061049a57607f821691505b6020821081036104ad576104ac610456565b5b50919050565b5f81905092915050565b5f6104c782610346565b6104d181856104b3565b93506104e1818560208601610360565b80840191505092915050565b5f6104f882846104bd565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261055f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610524565b6105698683610524565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6105ad6105a86105a384610581565b61058a565b610581565b9050919050565b5f819050919050565b6105c683610593565b6105da6105d2826105b4565b848454610530565b825550505050565b5f90565b6105ee6105e2565b6105f98184846105bd565b505050565b5b8181101561061c576106115f826105e6565b6001810190506105ff565b5050565b601f8211156106615761063281610503565b61063b84610515565b8101602085101561064a578190505b61065e61065685610515565b8301826105fe565b50505b505050565b5f82821c905092915050565b5f6106815f1984600802610666565b1980831691505092915050565b5f6106998383610672565b9150826002028217905092915050565b6106b282610346565b67ffffffffffffffff8111156106cb576106ca6101db565b5b6106d58254610483565b6106e0828285610620565b5f60209050601f831160018114610711575f84156106ff578287015190505b610709858261068e565b865550610770565b601f19841661071f86610503565b5f5b8281101561074657848901518255600182019150602085019450602081019050610721565b86831015610763578489015161075f601f891682610672565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220db6c60d8d8359a6f32ab1ee5cac909917d78fe1828b1459d0f3aec545e4316d464736f6c63430008180033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// Collection is a free data retrieval call binding the contract method 0xdc642f30.
//
// Solidity: function collection(string ) view returns(string)
func (_Storage *StorageCaller) Collection(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "collection", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Collection is a free data retrieval call binding the contract method 0xdc642f30.
//
// Solidity: function collection(string ) view returns(string)
func (_Storage *StorageSession) Collection(arg0 string) (string, error) {
	return _Storage.Contract.Collection(&_Storage.CallOpts, arg0)
}

// Collection is a free data retrieval call binding the contract method 0xdc642f30.
//
// Solidity: function collection(string ) view returns(string)
func (_Storage *StorageCallerSession) Collection(arg0 string) (string, error) {
	return _Storage.Contract.Collection(&_Storage.CallOpts, arg0)
}

// Store is a paid mutator transaction binding the contract method 0xf641090c.
//
// Solidity: function store(string _key, string _value) returns()
func (_Storage *StorageTransactor) Store(opts *bind.TransactOpts, _key string, _value string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "store", _key, _value)
}

// Store is a paid mutator transaction binding the contract method 0xf641090c.
//
// Solidity: function store(string _key, string _value) returns()
func (_Storage *StorageSession) Store(_key string, _value string) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, _key, _value)
}

// Store is a paid mutator transaction binding the contract method 0xf641090c.
//
// Solidity: function store(string _key, string _value) returns()
func (_Storage *StorageTransactorSession) Store(_key string, _value string) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, _key, _value)
}

// StorageStoredIterator is returned from FilterStored and is used to iterate over the raw logs and unpacked data for Stored events raised by the Storage contract.
type StorageStoredIterator struct {
	Event *StorageStored // Event containing the contract specifics and raw log

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
func (it *StorageStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageStored)
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
		it.Event = new(StorageStored)
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
func (it *StorageStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageStored represents a Stored event raised by the Storage contract.
type StorageStored struct {
	Key   common.Hash
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStored is a free log retrieval operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string indexed key, string value)
func (_Storage *StorageFilterer) FilterStored(opts *bind.FilterOpts, key []string) (*StorageStoredIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Stored", keyRule)
	if err != nil {
		return nil, err
	}
	return &StorageStoredIterator{contract: _Storage.contract, event: "Stored", logs: logs, sub: sub}, nil
}

// WatchStored is a free log subscription operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string indexed key, string value)
func (_Storage *StorageFilterer) WatchStored(opts *bind.WatchOpts, sink chan<- *StorageStored, key []string) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Stored", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageStored)
				if err := _Storage.contract.UnpackLog(event, "Stored", log); err != nil {
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

// ParseStored is a log parse operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string indexed key, string value)
func (_Storage *StorageFilterer) ParseStored(log types.Log) (*StorageStored, error) {
	event := new(StorageStored)
	if err := _Storage.contract.UnpackLog(event, "Stored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

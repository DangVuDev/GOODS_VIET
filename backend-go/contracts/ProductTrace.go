	// Code generated - DO NOT EDIT.
	// This file is a generated binding and any manual changes will be lost.

	package contracts

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

	// ProductTraceStatus is an auto generated low-level Go binding around an user-defined struct.
	type ProductTraceStatus struct {
		Status    string
		Details   string
		Timestamp *big.Int
	}

	// ContractsMetaData contains all meta data concerning the Contracts contract.
	var ContractsMetaData = &bind.MetaData{
		ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"manufacturer\",\"type\":\"string\"}],\"name\":\"ProductAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_manufacturer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_initialStatus\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_initialDetails\",\"type\":\"string\"}],\"name\":\"addProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_productId\",\"type\":\"uint256\"}],\"name\":\"getCurrentStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_productId\",\"type\":\"uint256\"}],\"name\":\"getProduct\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"manufacturer\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_productId\",\"type\":\"uint256\"}],\"name\":\"getProductStatuses\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structProductTrace.Status[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"productStatuses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"manufacturer\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_productId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_details\",\"type\":\"string\"}],\"name\":\"updateStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	}

	// ContractsABI is the input ABI used to generate the binding from.
	// Deprecated: Use ContractsMetaData.ABI instead.
	var ContractsABI = ContractsMetaData.ABI

	// Contracts is an auto generated Go binding around an Ethereum contract.
	type Contracts struct {
		ContractsCaller     // Read-only binding to the contract
		ContractsTransactor // Write-only binding to the contract
		ContractsFilterer   // Log filterer for contract events
	}

	// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ContractsCaller struct {
		contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ContractsTransactor struct {
		contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ContractsFilterer struct {
		contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ContractsSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ContractsSession struct {
		Contract     *Contracts        // Generic contract binding to set the session for
		CallOpts     bind.CallOpts     // Call options to use throughout this session
		TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ContractsCallerSession struct {
		Contract *ContractsCaller // Generic contract caller binding to set the session for
		CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ContractsTransactorSession struct {
		Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
		TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ContractsRaw struct {
		Contract *Contracts // Generic contract binding to access the raw methods on
	}

	// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ContractsCallerRaw struct {
		Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
	}

	// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ContractsTransactorRaw struct {
		Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
	func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
		contract, err := bindContracts(address, backend, backend, backend)
		if err != nil {
			return nil, err
		}
		return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
	}

	// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
	func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
		contract, err := bindContracts(address, caller, nil, nil)
		if err != nil {
			return nil, err
		}
		return &ContractsCaller{contract: contract}, nil
	}

	// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
	func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
		contract, err := bindContracts(address, nil, transactor, nil)
		if err != nil {
			return nil, err
		}
		return &ContractsTransactor{contract: contract}, nil
	}

	// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
	func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
		contract, err := bindContracts(address, nil, nil, filterer)
		if err != nil {
			return nil, err
		}
		return &ContractsFilterer{contract: contract}, nil
	}

	// bindContracts binds a generic wrapper to an already deployed contract.
	func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
		parsed, err := ContractsMetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Contracts.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Contracts.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Contracts.Contract.contract.Transact(opts, method, params...)
	}

	// GetCurrentStatus is a free data retrieval call binding the contract method 0xa0aa13c6.
	//
	// Solidity: function getCurrentStatus(uint256 _productId) view returns(string status, string details, uint256 timestamp)
	func (_Contracts *ContractsCaller) GetCurrentStatus(opts *bind.CallOpts, _productId *big.Int) (struct {
		Status    string
		Details   string
		Timestamp *big.Int
	}, error) {
		var out []interface{}
		err := _Contracts.contract.Call(opts, &out, "getCurrentStatus", _productId)

		outstruct := new(struct {
			Status    string
			Details   string
			Timestamp *big.Int
		})
		if err != nil {
			return *outstruct, err
		}

		outstruct.Status = *abi.ConvertType(out[0], new(string)).(*string)
		outstruct.Details = *abi.ConvertType(out[1], new(string)).(*string)
		outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

		return *outstruct, err

	}

	// GetCurrentStatus is a free data retrieval call binding the contract method 0xa0aa13c6.
	//
	// Solidity: function getCurrentStatus(uint256 _productId) view returns(string status, string details, uint256 timestamp)
	func (_Contracts *ContractsSession) GetCurrentStatus(_productId *big.Int) (struct {
		Status    string
		Details   string
		Timestamp *big.Int
	}, error) {
		return _Contracts.Contract.GetCurrentStatus(&_Contracts.CallOpts, _productId)
	}

	// GetCurrentStatus is a free data retrieval call binding the contract method 0xa0aa13c6.
	//
	// Solidity: function getCurrentStatus(uint256 _productId) view returns(string status, string details, uint256 timestamp)
	func (_Contracts *ContractsCallerSession) GetCurrentStatus(_productId *big.Int) (struct {
		Status    string
		Details   string
		Timestamp *big.Int
	}, error) {
		return _Contracts.Contract.GetCurrentStatus(&_Contracts.CallOpts, _productId)
	}

	// GetProduct is a free data retrieval call binding the contract method 0xb9db15b4.
	//
	// Solidity: function getProduct(uint256 _productId) view returns(string name, string manufacturer)
	func (_Contracts *ContractsCaller) GetProduct(opts *bind.CallOpts, _productId *big.Int) (struct {
		Name         string
		Manufacturer string
	}, error) {
		var out []interface{}
		err := _Contracts.contract.Call(opts, &out, "getProduct", _productId)

		outstruct := new(struct {
			Name         string
			Manufacturer string
		})
		if err != nil {
			return *outstruct, err
		}

		outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
		outstruct.Manufacturer = *abi.ConvertType(out[1], new(string)).(*string)

		return *outstruct, err

	}

	// GetProduct is a free data retrieval call binding the contract method 0xb9db15b4.
	//
	// Solidity: function getProduct(uint256 _productId) view returns(string name, string manufacturer)
	func (_Contracts *ContractsSession) GetProduct(_productId *big.Int) (struct {
		Name         string
		Manufacturer string
	}, error) {
		return _Contracts.Contract.GetProduct(&_Contracts.CallOpts, _productId)
	}

	// GetProduct is a free data retrieval call binding the contract method 0xb9db15b4.
	//
	// Solidity: function getProduct(uint256 _productId) view returns(string name, string manufacturer)
	func (_Contracts *ContractsCallerSession) GetProduct(_productId *big.Int) (struct {
		Name         string
		Manufacturer string
	}, error) {
		return _Contracts.Contract.GetProduct(&_Contracts.CallOpts, _productId)
	}

	// GetProductStatuses is a free data retrieval call binding the contract method 0x9fbecd86.
	//
	// Solidity: function getProductStatuses(uint256 _productId) view returns((string,string,uint256)[])
	func (_Contracts *ContractsCaller) GetProductStatuses(opts *bind.CallOpts, _productId *big.Int) ([]ProductTraceStatus, error) {
		var out []interface{}
		err := _Contracts.contract.Call(opts, &out, "getProductStatuses", _productId)

		if err != nil {
			return *new([]ProductTraceStatus), err
		}

		out0 := *abi.ConvertType(out[0], new([]ProductTraceStatus)).(*[]ProductTraceStatus)

		return out0, err

	}

	// GetProductStatuses is a free data retrieval call binding the contract method 0x9fbecd86.
	//
	// Solidity: function getProductStatuses(uint256 _productId) view returns((string,string,uint256)[])
	func (_Contracts *ContractsSession) GetProductStatuses(_productId *big.Int) ([]ProductTraceStatus, error) {
		return _Contracts.Contract.GetProductStatuses(&_Contracts.CallOpts, _productId)
	}

	// GetProductStatuses is a free data retrieval call binding the contract method 0x9fbecd86.
	//
	// Solidity: function getProductStatuses(uint256 _productId) view returns((string,string,uint256)[])
	func (_Contracts *ContractsCallerSession) GetProductStatuses(_productId *big.Int) ([]ProductTraceStatus, error) {
		return _Contracts.Contract.GetProductStatuses(&_Contracts.CallOpts, _productId)
	}

	// ProductCount is a free data retrieval call binding the contract method 0xe0f6ef87.
	//
	// Solidity: function productCount() view returns(uint256)
	func (_Contracts *ContractsCaller) ProductCount(opts *bind.CallOpts) (*big.Int, error) {
		var out []interface{}
		err := _Contracts.contract.Call(opts, &out, "productCount")

		if err != nil {
			return *new(*big.Int), err
		}

		out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return out0, err

	}

	// ProductCount is a free data retrieval call binding the contract method 0xe0f6ef87.
	//
	// Solidity: function productCount() view returns(uint256)
	func (_Contracts *ContractsSession) ProductCount() (*big.Int, error) {
		return _Contracts.Contract.ProductCount(&_Contracts.CallOpts)
	}

	// ProductCount is a free data retrieval call binding the contract method 0xe0f6ef87.
	//
	// Solidity: function productCount() view returns(uint256)
	func (_Contracts *ContractsCallerSession) ProductCount() (*big.Int, error) {
		return _Contracts.Contract.ProductCount(&_Contracts.CallOpts)
	}

	// ProductStatuses is a free data retrieval call binding the contract method 0x82370194.
	//
	// Solidity: function productStatuses(uint256 , uint256 ) view returns(string status, string details, uint256 timestamp)
	func (_Contracts *ContractsCaller) ProductStatuses(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
		Status    string
		Details   string
		Timestamp *big.Int
	}, error) {
		var out []interface{}
		err := _Contracts.contract.Call(opts, &out, "productStatuses", arg0, arg1)

		outstruct := new(struct {
			Status    string
			Details   string
			Timestamp *big.Int
		})
		if err != nil {
			return *outstruct, err
		}

		outstruct.Status = *abi.ConvertType(out[0], new(string)).(*string)
		outstruct.Details = *abi.ConvertType(out[1], new(string)).(*string)
		outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

		return *outstruct, err

	}

	// ProductStatuses is a free data retrieval call binding the contract method 0x82370194.
	//
	// Solidity: function productStatuses(uint256 , uint256 ) view returns(string status, string details, uint256 timestamp)
	func (_Contracts *ContractsSession) ProductStatuses(arg0 *big.Int, arg1 *big.Int) (struct {
		Status    string
		Details   string
		Timestamp *big.Int
	}, error) {
		return _Contracts.Contract.ProductStatuses(&_Contracts.CallOpts, arg0, arg1)
	}

	// ProductStatuses is a free data retrieval call binding the contract method 0x82370194.
	//
	// Solidity: function productStatuses(uint256 , uint256 ) view returns(string status, string details, uint256 timestamp)
	func (_Contracts *ContractsCallerSession) ProductStatuses(arg0 *big.Int, arg1 *big.Int) (struct {
		Status    string
		Details   string
		Timestamp *big.Int
	}, error) {
		return _Contracts.Contract.ProductStatuses(&_Contracts.CallOpts, arg0, arg1)
	}

	// Products is a free data retrieval call binding the contract method 0x7acc0b20.
	//
	// Solidity: function products(uint256 ) view returns(uint256 productId, string name, string manufacturer)
	func (_Contracts *ContractsCaller) Products(opts *bind.CallOpts, arg0 *big.Int) (struct {
		ProductId    *big.Int
		Name         string
		Manufacturer string
	}, error) {
		var out []interface{}
		err := _Contracts.contract.Call(opts, &out, "products", arg0)

		outstruct := new(struct {
			ProductId    *big.Int
			Name         string
			Manufacturer string
		})
		if err != nil {
			return *outstruct, err
		}

		outstruct.ProductId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
		outstruct.Manufacturer = *abi.ConvertType(out[2], new(string)).(*string)

		return *outstruct, err

	}

	// Products is a free data retrieval call binding the contract method 0x7acc0b20.
	//
	// Solidity: function products(uint256 ) view returns(uint256 productId, string name, string manufacturer)
	func (_Contracts *ContractsSession) Products(arg0 *big.Int) (struct {
		ProductId    *big.Int
		Name         string
		Manufacturer string
	}, error) {
		return _Contracts.Contract.Products(&_Contracts.CallOpts, arg0)
	}

	// Products is a free data retrieval call binding the contract method 0x7acc0b20.
	//
	// Solidity: function products(uint256 ) view returns(uint256 productId, string name, string manufacturer)
	func (_Contracts *ContractsCallerSession) Products(arg0 *big.Int) (struct {
		ProductId    *big.Int
		Name         string
		Manufacturer string
	}, error) {
		return _Contracts.Contract.Products(&_Contracts.CallOpts, arg0)
	}

	// AddProduct is a paid mutator transaction binding the contract method 0xfde74e81.
	//
	// Solidity: function addProduct(string _name, string _manufacturer, string _initialStatus, string _initialDetails) returns()
	func (_Contracts *ContractsTransactor) AddProduct(opts *bind.TransactOpts, _name string, _manufacturer string, _initialStatus string, _initialDetails string) (*types.Transaction, error) {
		return _Contracts.contract.Transact(opts, "addProduct", _name, _manufacturer, _initialStatus, _initialDetails)
	}

	// AddProduct is a paid mutator transaction binding the contract method 0xfde74e81.
	//
	// Solidity: function addProduct(string _name, string _manufacturer, string _initialStatus, string _initialDetails) returns()
	func (_Contracts *ContractsSession) AddProduct(_name string, _manufacturer string, _initialStatus string, _initialDetails string) (*types.Transaction, error) {
		return _Contracts.Contract.AddProduct(&_Contracts.TransactOpts, _name, _manufacturer, _initialStatus, _initialDetails)
	}

	// AddProduct is a paid mutator transaction binding the contract method 0xfde74e81.
	//
	// Solidity: function addProduct(string _name, string _manufacturer, string _initialStatus, string _initialDetails) returns()
	func (_Contracts *ContractsTransactorSession) AddProduct(_name string, _manufacturer string, _initialStatus string, _initialDetails string) (*types.Transaction, error) {
		return _Contracts.Contract.AddProduct(&_Contracts.TransactOpts, _name, _manufacturer, _initialStatus, _initialDetails)
	}

	// UpdateStatus is a paid mutator transaction binding the contract method 0xb49dbfc2.
	//
	// Solidity: function updateStatus(uint256 _productId, string _status, string _details) returns()
	func (_Contracts *ContractsTransactor) UpdateStatus(opts *bind.TransactOpts, _productId *big.Int, _status string, _details string) (*types.Transaction, error) {
		return _Contracts.contract.Transact(opts, "updateStatus", _productId, _status, _details)
	}

	// UpdateStatus is a paid mutator transaction binding the contract method 0xb49dbfc2.
	//
	// Solidity: function updateStatus(uint256 _productId, string _status, string _details) returns()
	func (_Contracts *ContractsSession) UpdateStatus(_productId *big.Int, _status string, _details string) (*types.Transaction, error) {
		return _Contracts.Contract.UpdateStatus(&_Contracts.TransactOpts, _productId, _status, _details)
	}

	// UpdateStatus is a paid mutator transaction binding the contract method 0xb49dbfc2.
	//
	// Solidity: function updateStatus(uint256 _productId, string _status, string _details) returns()
	func (_Contracts *ContractsTransactorSession) UpdateStatus(_productId *big.Int, _status string, _details string) (*types.Transaction, error) {
		return _Contracts.Contract.UpdateStatus(&_Contracts.TransactOpts, _productId, _status, _details)
	}

	// ContractsProductAddedIterator is returned from FilterProductAdded and is used to iterate over the raw logs and unpacked data for ProductAdded events raised by the Contracts contract.
	type ContractsProductAddedIterator struct {
		Event *ContractsProductAdded // Event containing the contract specifics and raw log

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
	func (it *ContractsProductAddedIterator) Next() bool {
		// If the iterator failed, stop iterating
		if it.fail != nil {
			return false
		}
		// If the iterator completed, deliver directly whatever's available
		if it.done {
			select {
			case log := <-it.logs:
				it.Event = new(ContractsProductAdded)
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
			it.Event = new(ContractsProductAdded)
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
	func (it *ContractsProductAddedIterator) Error() error {
		return it.fail
	}

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	func (it *ContractsProductAddedIterator) Close() error {
		it.sub.Unsubscribe()
		return nil
	}

	// ContractsProductAdded represents a ProductAdded event raised by the Contracts contract.
	type ContractsProductAdded struct {
		ProductId    *big.Int
		Name         string
		Manufacturer string
		Raw          types.Log // Blockchain specific contextual infos
	}

	// FilterProductAdded is a free log retrieval operation binding the contract event 0xb0b1b5b61bc0dbb800dfdecd756ac304fe33d3a63c2ac5dc419452555c99a829.
	//
	// Solidity: event ProductAdded(uint256 indexed productId, string name, string manufacturer)
	func (_Contracts *ContractsFilterer) FilterProductAdded(opts *bind.FilterOpts, productId []*big.Int) (*ContractsProductAddedIterator, error) {

		var productIdRule []interface{}
		for _, productIdItem := range productId {
			productIdRule = append(productIdRule, productIdItem)
		}

		logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProductAdded", productIdRule)
		if err != nil {
			return nil, err
		}
		return &ContractsProductAddedIterator{contract: _Contracts.contract, event: "ProductAdded", logs: logs, sub: sub}, nil
	}

	// WatchProductAdded is a free log subscription operation binding the contract event 0xb0b1b5b61bc0dbb800dfdecd756ac304fe33d3a63c2ac5dc419452555c99a829.
	//
	// Solidity: event ProductAdded(uint256 indexed productId, string name, string manufacturer)
	func (_Contracts *ContractsFilterer) WatchProductAdded(opts *bind.WatchOpts, sink chan<- *ContractsProductAdded, productId []*big.Int) (event.Subscription, error) {

		var productIdRule []interface{}
		for _, productIdItem := range productId {
			productIdRule = append(productIdRule, productIdItem)
		}

		logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProductAdded", productIdRule)
		if err != nil {
			return nil, err
		}
		return event.NewSubscription(func(quit <-chan struct{}) error {
			defer sub.Unsubscribe()
			for {
				select {
				case log := <-logs:
					// New log arrived, parse the event and forward to the user
					event := new(ContractsProductAdded)
					if err := _Contracts.contract.UnpackLog(event, "ProductAdded", log); err != nil {
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

	// ParseProductAdded is a log parse operation binding the contract event 0xb0b1b5b61bc0dbb800dfdecd756ac304fe33d3a63c2ac5dc419452555c99a829.
	//
	// Solidity: event ProductAdded(uint256 indexed productId, string name, string manufacturer)
	func (_Contracts *ContractsFilterer) ParseProductAdded(log types.Log) (*ContractsProductAdded, error) {
		event := new(ContractsProductAdded)
		if err := _Contracts.contract.UnpackLog(event, "ProductAdded", log); err != nil {
			return nil, err
		}
		event.Raw = log
		return event, nil
	}

	// ContractsStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the Contracts contract.
	type ContractsStatusUpdatedIterator struct {
		Event *ContractsStatusUpdated // Event containing the contract specifics and raw log

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
	func (it *ContractsStatusUpdatedIterator) Next() bool {
		// If the iterator failed, stop iterating
		if it.fail != nil {
			return false
		}
		// If the iterator completed, deliver directly whatever's available
		if it.done {
			select {
			case log := <-it.logs:
				it.Event = new(ContractsStatusUpdated)
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
			it.Event = new(ContractsStatusUpdated)
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
	func (it *ContractsStatusUpdatedIterator) Error() error {
		return it.fail
	}

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	func (it *ContractsStatusUpdatedIterator) Close() error {
		it.sub.Unsubscribe()
		return nil
	}

	// ContractsStatusUpdated represents a StatusUpdated event raised by the Contracts contract.
	type ContractsStatusUpdated struct {
		ProductId *big.Int
		Status    string
		Details   string
		Timestamp *big.Int
		Raw       types.Log // Blockchain specific contextual infos
	}

	// FilterStatusUpdated is a free log retrieval operation binding the contract event 0xf09e75b124e0bcc67c43558bce3f0b9e4a56686362aae44f0ba1a4ebf79b37d0.
	//
	// Solidity: event StatusUpdated(uint256 indexed productId, string status, string details, uint256 timestamp)
	func (_Contracts *ContractsFilterer) FilterStatusUpdated(opts *bind.FilterOpts, productId []*big.Int) (*ContractsStatusUpdatedIterator, error) {

		var productIdRule []interface{}
		for _, productIdItem := range productId {
			productIdRule = append(productIdRule, productIdItem)
		}

		logs, sub, err := _Contracts.contract.FilterLogs(opts, "StatusUpdated", productIdRule)
		if err != nil {
			return nil, err
		}
		return &ContractsStatusUpdatedIterator{contract: _Contracts.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
	}

	// WatchStatusUpdated is a free log subscription operation binding the contract event 0xf09e75b124e0bcc67c43558bce3f0b9e4a56686362aae44f0ba1a4ebf79b37d0.
	//
	// Solidity: event StatusUpdated(uint256 indexed productId, string status, string details, uint256 timestamp)
	func (_Contracts *ContractsFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractsStatusUpdated, productId []*big.Int) (event.Subscription, error) {

		var productIdRule []interface{}
		for _, productIdItem := range productId {
			productIdRule = append(productIdRule, productIdItem)
		}

		logs, sub, err := _Contracts.contract.WatchLogs(opts, "StatusUpdated", productIdRule)
		if err != nil {
			return nil, err
		}
		return event.NewSubscription(func(quit <-chan struct{}) error {
			defer sub.Unsubscribe()
			for {
				select {
				case log := <-logs:
					// New log arrived, parse the event and forward to the user
					event := new(ContractsStatusUpdated)
					if err := _Contracts.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
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

	// ParseStatusUpdated is a log parse operation binding the contract event 0xf09e75b124e0bcc67c43558bce3f0b9e4a56686362aae44f0ba1a4ebf79b37d0.
	//
	// Solidity: event StatusUpdated(uint256 indexed productId, string status, string details, uint256 timestamp)
	func (_Contracts *ContractsFilterer) ParseStatusUpdated(log types.Log) (*ContractsStatusUpdated, error) {
		event := new(ContractsStatusUpdated)
		if err := _Contracts.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
			return nil, err
		}
		event.Raw = log
		return event, nil
	}

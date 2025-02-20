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
)

// YRegistryV4MetaData contains all meta data concerning the YRegistryV4 contract.
var YRegistryV4MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_releaseRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"}],\"name\":\"NewEndorsedStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"}],\"name\":\"NewEndorsedVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"}],\"name\":\"RemovedStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"}],\"name\":\"RemovedVault\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetIsUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deploymentTimestamp\",\"type\":\"uint256\"}],\"name\":\"endorseStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"endorseStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deploymentTimestamp\",\"type\":\"uint256\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllEndorsedStrategies\",\"outputs\":[{\"internalType\":\"address[][]\",\"name\":\"allEndorsedStrategies\",\"type\":\"address[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllEndorsedVaults\",\"outputs\":[{\"internalType\":\"address[][]\",\"name\":\"allEndorsedVaults\",\"type\":\"address[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getEndorsedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getEndorsedVaults\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"info\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deploymentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_roleManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_profitMaxUnlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"}],\"name\":\"newEndorsedVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"numEndorsedStrategies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"numEndorsedVaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"}],\"name\":\"tagVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YRegistryV4ABI is the input ABI used to generate the binding from.
// Deprecated: Use YRegistryV4MetaData.ABI instead.
var YRegistryV4ABI = YRegistryV4MetaData.ABI

// YRegistryV4 is an auto generated Go binding around an Ethereum contract.
type YRegistryV4 struct {
	YRegistryV4Caller     // Read-only binding to the contract
	YRegistryV4Transactor // Write-only binding to the contract
	YRegistryV4Filterer   // Log filterer for contract events
}

// YRegistryV4Caller is an auto generated read-only Go binding around an Ethereum contract.
type YRegistryV4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type YRegistryV4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YRegistryV4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YRegistryV4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YRegistryV4Session struct {
	Contract     *YRegistryV4      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YRegistryV4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YRegistryV4CallerSession struct {
	Contract *YRegistryV4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YRegistryV4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YRegistryV4TransactorSession struct {
	Contract     *YRegistryV4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YRegistryV4Raw is an auto generated low-level Go binding around an Ethereum contract.
type YRegistryV4Raw struct {
	Contract *YRegistryV4 // Generic contract binding to access the raw methods on
}

// YRegistryV4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YRegistryV4CallerRaw struct {
	Contract *YRegistryV4Caller // Generic read-only contract binding to access the raw methods on
}

// YRegistryV4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YRegistryV4TransactorRaw struct {
	Contract *YRegistryV4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYRegistryV4 creates a new instance of YRegistryV4, bound to a specific deployed contract.
func NewYRegistryV4(address common.Address, backend bind.ContractBackend) (*YRegistryV4, error) {
	contract, err := bindYRegistryV4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4{YRegistryV4Caller: YRegistryV4Caller{contract: contract}, YRegistryV4Transactor: YRegistryV4Transactor{contract: contract}, YRegistryV4Filterer: YRegistryV4Filterer{contract: contract}}, nil
}

// NewYRegistryV4Caller creates a new read-only instance of YRegistryV4, bound to a specific deployed contract.
func NewYRegistryV4Caller(address common.Address, caller bind.ContractCaller) (*YRegistryV4Caller, error) {
	contract, err := bindYRegistryV4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4Caller{contract: contract}, nil
}

// NewYRegistryV4Transactor creates a new write-only instance of YRegistryV4, bound to a specific deployed contract.
func NewYRegistryV4Transactor(address common.Address, transactor bind.ContractTransactor) (*YRegistryV4Transactor, error) {
	contract, err := bindYRegistryV4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4Transactor{contract: contract}, nil
}

// NewYRegistryV4Filterer creates a new log filterer instance of YRegistryV4, bound to a specific deployed contract.
func NewYRegistryV4Filterer(address common.Address, filterer bind.ContractFilterer) (*YRegistryV4Filterer, error) {
	contract, err := bindYRegistryV4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4Filterer{contract: contract}, nil
}

// bindYRegistryV4 binds a generic wrapper to an already deployed contract.
func bindYRegistryV4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YRegistryV4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV4 *YRegistryV4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV4.Contract.YRegistryV4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV4 *YRegistryV4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV4.Contract.YRegistryV4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV4 *YRegistryV4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV4.Contract.YRegistryV4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YRegistryV4 *YRegistryV4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YRegistryV4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YRegistryV4 *YRegistryV4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YRegistryV4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YRegistryV4 *YRegistryV4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YRegistryV4.Contract.contract.Transact(opts, method, params...)
}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YRegistryV4 *YRegistryV4Caller) AssetIsUsed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "assetIsUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YRegistryV4 *YRegistryV4Session) AssetIsUsed(arg0 common.Address) (bool, error) {
	return _YRegistryV4.Contract.AssetIsUsed(&_YRegistryV4.CallOpts, arg0)
}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YRegistryV4 *YRegistryV4CallerSession) AssetIsUsed(arg0 common.Address) (bool, error) {
	return _YRegistryV4.Contract.AssetIsUsed(&_YRegistryV4.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YRegistryV4 *YRegistryV4Caller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "assets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YRegistryV4 *YRegistryV4Session) Assets(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV4.Contract.Assets(&_YRegistryV4.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YRegistryV4 *YRegistryV4CallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _YRegistryV4.Contract.Assets(&_YRegistryV4.CallOpts, arg0)
}

// GetAllEndorsedStrategies is a free data retrieval call binding the contract method 0x06a70f3d.
//
// Solidity: function getAllEndorsedStrategies() view returns(address[][] allEndorsedStrategies)
func (_YRegistryV4 *YRegistryV4Caller) GetAllEndorsedStrategies(opts *bind.CallOpts) ([][]common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "getAllEndorsedStrategies")

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// GetAllEndorsedStrategies is a free data retrieval call binding the contract method 0x06a70f3d.
//
// Solidity: function getAllEndorsedStrategies() view returns(address[][] allEndorsedStrategies)
func (_YRegistryV4 *YRegistryV4Session) GetAllEndorsedStrategies() ([][]common.Address, error) {
	return _YRegistryV4.Contract.GetAllEndorsedStrategies(&_YRegistryV4.CallOpts)
}

// GetAllEndorsedStrategies is a free data retrieval call binding the contract method 0x06a70f3d.
//
// Solidity: function getAllEndorsedStrategies() view returns(address[][] allEndorsedStrategies)
func (_YRegistryV4 *YRegistryV4CallerSession) GetAllEndorsedStrategies() ([][]common.Address, error) {
	return _YRegistryV4.Contract.GetAllEndorsedStrategies(&_YRegistryV4.CallOpts)
}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YRegistryV4 *YRegistryV4Caller) GetAllEndorsedVaults(opts *bind.CallOpts) ([][]common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "getAllEndorsedVaults")

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YRegistryV4 *YRegistryV4Session) GetAllEndorsedVaults() ([][]common.Address, error) {
	return _YRegistryV4.Contract.GetAllEndorsedVaults(&_YRegistryV4.CallOpts)
}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YRegistryV4 *YRegistryV4CallerSession) GetAllEndorsedVaults() ([][]common.Address, error) {
	return _YRegistryV4.Contract.GetAllEndorsedVaults(&_YRegistryV4.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YRegistryV4 *YRegistryV4Caller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YRegistryV4 *YRegistryV4Session) GetAssets() ([]common.Address, error) {
	return _YRegistryV4.Contract.GetAssets(&_YRegistryV4.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YRegistryV4 *YRegistryV4CallerSession) GetAssets() ([]common.Address, error) {
	return _YRegistryV4.Contract.GetAssets(&_YRegistryV4.CallOpts)
}

// GetEndorsedStrategies is a free data retrieval call binding the contract method 0x153a5b16.
//
// Solidity: function getEndorsedStrategies(address _asset) view returns(address[])
func (_YRegistryV4 *YRegistryV4Caller) GetEndorsedStrategies(opts *bind.CallOpts, _asset common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "getEndorsedStrategies", _asset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEndorsedStrategies is a free data retrieval call binding the contract method 0x153a5b16.
//
// Solidity: function getEndorsedStrategies(address _asset) view returns(address[])
func (_YRegistryV4 *YRegistryV4Session) GetEndorsedStrategies(_asset common.Address) ([]common.Address, error) {
	return _YRegistryV4.Contract.GetEndorsedStrategies(&_YRegistryV4.CallOpts, _asset)
}

// GetEndorsedStrategies is a free data retrieval call binding the contract method 0x153a5b16.
//
// Solidity: function getEndorsedStrategies(address _asset) view returns(address[])
func (_YRegistryV4 *YRegistryV4CallerSession) GetEndorsedStrategies(_asset common.Address) ([]common.Address, error) {
	return _YRegistryV4.Contract.GetEndorsedStrategies(&_YRegistryV4.CallOpts, _asset)
}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YRegistryV4 *YRegistryV4Caller) GetEndorsedVaults(opts *bind.CallOpts, _asset common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "getEndorsedVaults", _asset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YRegistryV4 *YRegistryV4Session) GetEndorsedVaults(_asset common.Address) ([]common.Address, error) {
	return _YRegistryV4.Contract.GetEndorsedVaults(&_YRegistryV4.CallOpts, _asset)
}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YRegistryV4 *YRegistryV4CallerSession) GetEndorsedVaults(_asset common.Address) ([]common.Address, error) {
	return _YRegistryV4.Contract.GetEndorsedVaults(&_YRegistryV4.CallOpts, _asset)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV4 *YRegistryV4Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV4 *YRegistryV4Session) Governance() (common.Address, error) {
	return _YRegistryV4.Contract.Governance(&_YRegistryV4.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YRegistryV4 *YRegistryV4CallerSession) Governance() (common.Address, error) {
	return _YRegistryV4.Contract.Governance(&_YRegistryV4.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address ) view returns(address asset, uint256 releaseVersion, uint256 deploymentTimestamp, string tag)
func (_YRegistryV4 *YRegistryV4Caller) Info(opts *bind.CallOpts, arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	DeploymentTimestamp *big.Int
	Tag                 string
}, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "info", arg0)

	outstruct := new(struct {
		Asset               common.Address
		ReleaseVersion      *big.Int
		DeploymentTimestamp *big.Int
		Tag                 string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ReleaseVersion = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeploymentTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tag = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address ) view returns(address asset, uint256 releaseVersion, uint256 deploymentTimestamp, string tag)
func (_YRegistryV4 *YRegistryV4Session) Info(arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	DeploymentTimestamp *big.Int
	Tag                 string
}, error) {
	return _YRegistryV4.Contract.Info(&_YRegistryV4.CallOpts, arg0)
}

// Info is a free data retrieval call binding the contract method 0x0aae7a6b.
//
// Solidity: function info(address ) view returns(address asset, uint256 releaseVersion, uint256 deploymentTimestamp, string tag)
func (_YRegistryV4 *YRegistryV4CallerSession) Info(arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	DeploymentTimestamp *big.Int
	Tag                 string
}, error) {
	return _YRegistryV4.Contract.Info(&_YRegistryV4.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YRegistryV4 *YRegistryV4Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YRegistryV4 *YRegistryV4Session) Name() (string, error) {
	return _YRegistryV4.Contract.Name(&_YRegistryV4.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YRegistryV4 *YRegistryV4CallerSession) Name() (string, error) {
	return _YRegistryV4.Contract.Name(&_YRegistryV4.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YRegistryV4 *YRegistryV4Caller) NumAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "numAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YRegistryV4 *YRegistryV4Session) NumAssets() (*big.Int, error) {
	return _YRegistryV4.Contract.NumAssets(&_YRegistryV4.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YRegistryV4 *YRegistryV4CallerSession) NumAssets() (*big.Int, error) {
	return _YRegistryV4.Contract.NumAssets(&_YRegistryV4.CallOpts)
}

// NumEndorsedStrategies is a free data retrieval call binding the contract method 0x3e3c7b60.
//
// Solidity: function numEndorsedStrategies(address _asset) view returns(uint256)
func (_YRegistryV4 *YRegistryV4Caller) NumEndorsedStrategies(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "numEndorsedStrategies", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumEndorsedStrategies is a free data retrieval call binding the contract method 0x3e3c7b60.
//
// Solidity: function numEndorsedStrategies(address _asset) view returns(uint256)
func (_YRegistryV4 *YRegistryV4Session) NumEndorsedStrategies(_asset common.Address) (*big.Int, error) {
	return _YRegistryV4.Contract.NumEndorsedStrategies(&_YRegistryV4.CallOpts, _asset)
}

// NumEndorsedStrategies is a free data retrieval call binding the contract method 0x3e3c7b60.
//
// Solidity: function numEndorsedStrategies(address _asset) view returns(uint256)
func (_YRegistryV4 *YRegistryV4CallerSession) NumEndorsedStrategies(_asset common.Address) (*big.Int, error) {
	return _YRegistryV4.Contract.NumEndorsedStrategies(&_YRegistryV4.CallOpts, _asset)
}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YRegistryV4 *YRegistryV4Caller) NumEndorsedVaults(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "numEndorsedVaults", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YRegistryV4 *YRegistryV4Session) NumEndorsedVaults(_asset common.Address) (*big.Int, error) {
	return _YRegistryV4.Contract.NumEndorsedVaults(&_YRegistryV4.CallOpts, _asset)
}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YRegistryV4 *YRegistryV4CallerSession) NumEndorsedVaults(_asset common.Address) (*big.Int, error) {
	return _YRegistryV4.Contract.NumEndorsedVaults(&_YRegistryV4.CallOpts, _asset)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YRegistryV4 *YRegistryV4Caller) ReleaseRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YRegistryV4.contract.Call(opts, &out, "releaseRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YRegistryV4 *YRegistryV4Session) ReleaseRegistry() (common.Address, error) {
	return _YRegistryV4.Contract.ReleaseRegistry(&_YRegistryV4.CallOpts)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YRegistryV4 *YRegistryV4CallerSession) ReleaseRegistry() (common.Address, error) {
	return _YRegistryV4.Contract.ReleaseRegistry(&_YRegistryV4.CallOpts)
}

// EndorseStrategy is a paid mutator transaction binding the contract method 0x0724b07b.
//
// Solidity: function endorseStrategy(address _strategy, uint256 _releaseDelta, uint256 _deploymentTimestamp) returns()
func (_YRegistryV4 *YRegistryV4Transactor) EndorseStrategy(opts *bind.TransactOpts, _strategy common.Address, _releaseDelta *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "endorseStrategy", _strategy, _releaseDelta, _deploymentTimestamp)
}

// EndorseStrategy is a paid mutator transaction binding the contract method 0x0724b07b.
//
// Solidity: function endorseStrategy(address _strategy, uint256 _releaseDelta, uint256 _deploymentTimestamp) returns()
func (_YRegistryV4 *YRegistryV4Session) EndorseStrategy(_strategy common.Address, _releaseDelta *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseStrategy(&_YRegistryV4.TransactOpts, _strategy, _releaseDelta, _deploymentTimestamp)
}

// EndorseStrategy is a paid mutator transaction binding the contract method 0x0724b07b.
//
// Solidity: function endorseStrategy(address _strategy, uint256 _releaseDelta, uint256 _deploymentTimestamp) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) EndorseStrategy(_strategy common.Address, _releaseDelta *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseStrategy(&_YRegistryV4.TransactOpts, _strategy, _releaseDelta, _deploymentTimestamp)
}

// EndorseStrategy0 is a paid mutator transaction binding the contract method 0x0a225ecf.
//
// Solidity: function endorseStrategy(address _strategy) returns()
func (_YRegistryV4 *YRegistryV4Transactor) EndorseStrategy0(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "endorseStrategy0", _strategy)
}

// EndorseStrategy0 is a paid mutator transaction binding the contract method 0x0a225ecf.
//
// Solidity: function endorseStrategy(address _strategy) returns()
func (_YRegistryV4 *YRegistryV4Session) EndorseStrategy0(_strategy common.Address) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseStrategy0(&_YRegistryV4.TransactOpts, _strategy)
}

// EndorseStrategy0 is a paid mutator transaction binding the contract method 0x0a225ecf.
//
// Solidity: function endorseStrategy(address _strategy) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) EndorseStrategy0(_strategy common.Address) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseStrategy0(&_YRegistryV4.TransactOpts, _strategy)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address _vault) returns()
func (_YRegistryV4 *YRegistryV4Transactor) EndorseVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "endorseVault", _vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address _vault) returns()
func (_YRegistryV4 *YRegistryV4Session) EndorseVault(_vault common.Address) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseVault(&_YRegistryV4.TransactOpts, _vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x29b2e0c6.
//
// Solidity: function endorseVault(address _vault) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) EndorseVault(_vault common.Address) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseVault(&_YRegistryV4.TransactOpts, _vault)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0x931074ba.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _deploymentTimestamp) returns()
func (_YRegistryV4 *YRegistryV4Transactor) EndorseVault0(opts *bind.TransactOpts, _vault common.Address, _releaseDelta *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "endorseVault0", _vault, _releaseDelta, _deploymentTimestamp)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0x931074ba.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _deploymentTimestamp) returns()
func (_YRegistryV4 *YRegistryV4Session) EndorseVault0(_vault common.Address, _releaseDelta *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseVault0(&_YRegistryV4.TransactOpts, _vault, _releaseDelta, _deploymentTimestamp)
}

// EndorseVault0 is a paid mutator transaction binding the contract method 0x931074ba.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _deploymentTimestamp) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) EndorseVault0(_vault common.Address, _releaseDelta *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.EndorseVault0(&_YRegistryV4.TransactOpts, _vault, _releaseDelta, _deploymentTimestamp)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YRegistryV4 *YRegistryV4Transactor) NewEndorsedVault(opts *bind.TransactOpts, _asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "newEndorsedVault", _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YRegistryV4 *YRegistryV4Session) NewEndorsedVault(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.NewEndorsedVault(&_YRegistryV4.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YRegistryV4 *YRegistryV4TransactorSession) NewEndorsedVault(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.NewEndorsedVault(&_YRegistryV4.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4Transactor) RemoveAsset(opts *bind.TransactOpts, _asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "removeAsset", _asset, _index)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4Session) RemoveAsset(_asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.RemoveAsset(&_YRegistryV4.TransactOpts, _asset, _index)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) RemoveAsset(_asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.RemoveAsset(&_YRegistryV4.TransactOpts, _asset, _index)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0xea682eeb.
//
// Solidity: function removeStrategy(address _strategy, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4Transactor) RemoveStrategy(opts *bind.TransactOpts, _strategy common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "removeStrategy", _strategy, _index)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0xea682eeb.
//
// Solidity: function removeStrategy(address _strategy, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4Session) RemoveStrategy(_strategy common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.RemoveStrategy(&_YRegistryV4.TransactOpts, _strategy, _index)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0xea682eeb.
//
// Solidity: function removeStrategy(address _strategy, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) RemoveStrategy(_strategy common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.RemoveStrategy(&_YRegistryV4.TransactOpts, _strategy, _index)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xb37c61cd.
//
// Solidity: function removeVault(address _vault, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4Transactor) RemoveVault(opts *bind.TransactOpts, _vault common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "removeVault", _vault, _index)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xb37c61cd.
//
// Solidity: function removeVault(address _vault, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4Session) RemoveVault(_vault common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.RemoveVault(&_YRegistryV4.TransactOpts, _vault, _index)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xb37c61cd.
//
// Solidity: function removeVault(address _vault, uint256 _index) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) RemoveVault(_vault common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YRegistryV4.Contract.RemoveVault(&_YRegistryV4.TransactOpts, _vault, _index)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YRegistryV4 *YRegistryV4Transactor) TagVault(opts *bind.TransactOpts, _vault common.Address, _tag string) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "tagVault", _vault, _tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YRegistryV4 *YRegistryV4Session) TagVault(_vault common.Address, _tag string) (*types.Transaction, error) {
	return _YRegistryV4.Contract.TagVault(&_YRegistryV4.TransactOpts, _vault, _tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) TagVault(_vault common.Address, _tag string) (*types.Transaction, error) {
	return _YRegistryV4.Contract.TagVault(&_YRegistryV4.TransactOpts, _vault, _tag)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YRegistryV4 *YRegistryV4Transactor) TransferGovernance(opts *bind.TransactOpts, _newGovernance common.Address) (*types.Transaction, error) {
	return _YRegistryV4.contract.Transact(opts, "transferGovernance", _newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YRegistryV4 *YRegistryV4Session) TransferGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _YRegistryV4.Contract.TransferGovernance(&_YRegistryV4.TransactOpts, _newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YRegistryV4 *YRegistryV4TransactorSession) TransferGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _YRegistryV4.Contract.TransferGovernance(&_YRegistryV4.TransactOpts, _newGovernance)
}

// YRegistryV4GovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the YRegistryV4 contract.
type YRegistryV4GovernanceTransferredIterator struct {
	Event *YRegistryV4GovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *YRegistryV4GovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV4GovernanceTransferred)
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
		it.Event = new(YRegistryV4GovernanceTransferred)
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
func (it *YRegistryV4GovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV4GovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV4GovernanceTransferred represents a GovernanceTransferred event raised by the YRegistryV4 contract.
type YRegistryV4GovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YRegistryV4 *YRegistryV4Filterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*YRegistryV4GovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _YRegistryV4.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4GovernanceTransferredIterator{contract: _YRegistryV4.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YRegistryV4 *YRegistryV4Filterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *YRegistryV4GovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _YRegistryV4.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV4GovernanceTransferred)
				if err := _YRegistryV4.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YRegistryV4 *YRegistryV4Filterer) ParseGovernanceTransferred(log types.Log) (*YRegistryV4GovernanceTransferred, error) {
	event := new(YRegistryV4GovernanceTransferred)
	if err := _YRegistryV4.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV4NewEndorsedStrategyIterator is returned from FilterNewEndorsedStrategy and is used to iterate over the raw logs and unpacked data for NewEndorsedStrategy events raised by the YRegistryV4 contract.
type YRegistryV4NewEndorsedStrategyIterator struct {
	Event *YRegistryV4NewEndorsedStrategy // Event containing the contract specifics and raw log

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
func (it *YRegistryV4NewEndorsedStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV4NewEndorsedStrategy)
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
		it.Event = new(YRegistryV4NewEndorsedStrategy)
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
func (it *YRegistryV4NewEndorsedStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV4NewEndorsedStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV4NewEndorsedStrategy represents a NewEndorsedStrategy event raised by the YRegistryV4 contract.
type YRegistryV4NewEndorsedStrategy struct {
	Strategy       common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewEndorsedStrategy is a free log retrieval operation binding the contract event 0x2ab8c6c9129a30daa2d2add32ab462e1b35dd8fc42a473ea04380a25b5cc9d3a.
//
// Solidity: event NewEndorsedStrategy(address indexed strategy, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) FilterNewEndorsedStrategy(opts *bind.FilterOpts, strategy []common.Address, asset []common.Address) (*YRegistryV4NewEndorsedStrategyIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.FilterLogs(opts, "NewEndorsedStrategy", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4NewEndorsedStrategyIterator{contract: _YRegistryV4.contract, event: "NewEndorsedStrategy", logs: logs, sub: sub}, nil
}

// WatchNewEndorsedStrategy is a free log subscription operation binding the contract event 0x2ab8c6c9129a30daa2d2add32ab462e1b35dd8fc42a473ea04380a25b5cc9d3a.
//
// Solidity: event NewEndorsedStrategy(address indexed strategy, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) WatchNewEndorsedStrategy(opts *bind.WatchOpts, sink chan<- *YRegistryV4NewEndorsedStrategy, strategy []common.Address, asset []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.WatchLogs(opts, "NewEndorsedStrategy", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV4NewEndorsedStrategy)
				if err := _YRegistryV4.contract.UnpackLog(event, "NewEndorsedStrategy", log); err != nil {
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

// ParseNewEndorsedStrategy is a log parse operation binding the contract event 0x2ab8c6c9129a30daa2d2add32ab462e1b35dd8fc42a473ea04380a25b5cc9d3a.
//
// Solidity: event NewEndorsedStrategy(address indexed strategy, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) ParseNewEndorsedStrategy(log types.Log) (*YRegistryV4NewEndorsedStrategy, error) {
	event := new(YRegistryV4NewEndorsedStrategy)
	if err := _YRegistryV4.contract.UnpackLog(event, "NewEndorsedStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV4NewEndorsedVaultIterator is returned from FilterNewEndorsedVault and is used to iterate over the raw logs and unpacked data for NewEndorsedVault events raised by the YRegistryV4 contract.
type YRegistryV4NewEndorsedVaultIterator struct {
	Event *YRegistryV4NewEndorsedVault // Event containing the contract specifics and raw log

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
func (it *YRegistryV4NewEndorsedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV4NewEndorsedVault)
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
		it.Event = new(YRegistryV4NewEndorsedVault)
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
func (it *YRegistryV4NewEndorsedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV4NewEndorsedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV4NewEndorsedVault represents a NewEndorsedVault event raised by the YRegistryV4 contract.
type YRegistryV4NewEndorsedVault struct {
	Vault          common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewEndorsedVault is a free log retrieval operation binding the contract event 0x5bf19cf6c9f6c9210bc8cfecb4fda8057ebe0c41e300b60c5efa3de7f98f2f35.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) FilterNewEndorsedVault(opts *bind.FilterOpts, vault []common.Address, asset []common.Address) (*YRegistryV4NewEndorsedVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.FilterLogs(opts, "NewEndorsedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4NewEndorsedVaultIterator{contract: _YRegistryV4.contract, event: "NewEndorsedVault", logs: logs, sub: sub}, nil
}

// WatchNewEndorsedVault is a free log subscription operation binding the contract event 0x5bf19cf6c9f6c9210bc8cfecb4fda8057ebe0c41e300b60c5efa3de7f98f2f35.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) WatchNewEndorsedVault(opts *bind.WatchOpts, sink chan<- *YRegistryV4NewEndorsedVault, vault []common.Address, asset []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.WatchLogs(opts, "NewEndorsedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV4NewEndorsedVault)
				if err := _YRegistryV4.contract.UnpackLog(event, "NewEndorsedVault", log); err != nil {
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

// ParseNewEndorsedVault is a log parse operation binding the contract event 0x5bf19cf6c9f6c9210bc8cfecb4fda8057ebe0c41e300b60c5efa3de7f98f2f35.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) ParseNewEndorsedVault(log types.Log) (*YRegistryV4NewEndorsedVault, error) {
	event := new(YRegistryV4NewEndorsedVault)
	if err := _YRegistryV4.contract.UnpackLog(event, "NewEndorsedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV4RemovedStrategyIterator is returned from FilterRemovedStrategy and is used to iterate over the raw logs and unpacked data for RemovedStrategy events raised by the YRegistryV4 contract.
type YRegistryV4RemovedStrategyIterator struct {
	Event *YRegistryV4RemovedStrategy // Event containing the contract specifics and raw log

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
func (it *YRegistryV4RemovedStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV4RemovedStrategy)
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
		it.Event = new(YRegistryV4RemovedStrategy)
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
func (it *YRegistryV4RemovedStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV4RemovedStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV4RemovedStrategy represents a RemovedStrategy event raised by the YRegistryV4 contract.
type YRegistryV4RemovedStrategy struct {
	Strategy       common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemovedStrategy is a free log retrieval operation binding the contract event 0xd7cdfeb9ab09d32e715ad82cdf2440286c5d462f66b98170a579195945800953.
//
// Solidity: event RemovedStrategy(address indexed strategy, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) FilterRemovedStrategy(opts *bind.FilterOpts, strategy []common.Address, asset []common.Address) (*YRegistryV4RemovedStrategyIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.FilterLogs(opts, "RemovedStrategy", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4RemovedStrategyIterator{contract: _YRegistryV4.contract, event: "RemovedStrategy", logs: logs, sub: sub}, nil
}

// WatchRemovedStrategy is a free log subscription operation binding the contract event 0xd7cdfeb9ab09d32e715ad82cdf2440286c5d462f66b98170a579195945800953.
//
// Solidity: event RemovedStrategy(address indexed strategy, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) WatchRemovedStrategy(opts *bind.WatchOpts, sink chan<- *YRegistryV4RemovedStrategy, strategy []common.Address, asset []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.WatchLogs(opts, "RemovedStrategy", strategyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV4RemovedStrategy)
				if err := _YRegistryV4.contract.UnpackLog(event, "RemovedStrategy", log); err != nil {
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

// ParseRemovedStrategy is a log parse operation binding the contract event 0xd7cdfeb9ab09d32e715ad82cdf2440286c5d462f66b98170a579195945800953.
//
// Solidity: event RemovedStrategy(address indexed strategy, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) ParseRemovedStrategy(log types.Log) (*YRegistryV4RemovedStrategy, error) {
	event := new(YRegistryV4RemovedStrategy)
	if err := _YRegistryV4.contract.UnpackLog(event, "RemovedStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YRegistryV4RemovedVaultIterator is returned from FilterRemovedVault and is used to iterate over the raw logs and unpacked data for RemovedVault events raised by the YRegistryV4 contract.
type YRegistryV4RemovedVaultIterator struct {
	Event *YRegistryV4RemovedVault // Event containing the contract specifics and raw log

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
func (it *YRegistryV4RemovedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YRegistryV4RemovedVault)
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
		it.Event = new(YRegistryV4RemovedVault)
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
func (it *YRegistryV4RemovedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YRegistryV4RemovedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YRegistryV4RemovedVault represents a RemovedVault event raised by the YRegistryV4 contract.
type YRegistryV4RemovedVault struct {
	Vault          common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemovedVault is a free log retrieval operation binding the contract event 0x90f0b136797e67e4987983b7f7326571b458a7fd5cb6fb3e1209fdea9483b4cf.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) FilterRemovedVault(opts *bind.FilterOpts, vault []common.Address, asset []common.Address) (*YRegistryV4RemovedVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.FilterLogs(opts, "RemovedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YRegistryV4RemovedVaultIterator{contract: _YRegistryV4.contract, event: "RemovedVault", logs: logs, sub: sub}, nil
}

// WatchRemovedVault is a free log subscription operation binding the contract event 0x90f0b136797e67e4987983b7f7326571b458a7fd5cb6fb3e1209fdea9483b4cf.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) WatchRemovedVault(opts *bind.WatchOpts, sink chan<- *YRegistryV4RemovedVault, vault []common.Address, asset []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YRegistryV4.contract.WatchLogs(opts, "RemovedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YRegistryV4RemovedVault)
				if err := _YRegistryV4.contract.UnpackLog(event, "RemovedVault", log); err != nil {
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

// ParseRemovedVault is a log parse operation binding the contract event 0x90f0b136797e67e4987983b7f7326571b458a7fd5cb6fb3e1209fdea9483b4cf.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion)
func (_YRegistryV4 *YRegistryV4Filterer) ParseRemovedVault(log types.Log) (*YRegistryV4RemovedVault, error) {
	event := new(YRegistryV4RemovedVault)
	if err := _YRegistryV4.contract.UnpackLog(event, "RemovedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

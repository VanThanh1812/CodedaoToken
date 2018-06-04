// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prnetworkcontract

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

// BaseContractABI is the input ABI used to generate the binding from.
const BaseContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"default_gateway\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"mapLinkContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"super_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractLink\",\"type\":\"address\"},{\"name\":\"_ownerContractLink\",\"type\":\"address\"}],\"name\":\"saveContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addrContract\",\"type\":\"address\"},{\"name\":\"_addrRecive\",\"type\":\"address\"},{\"name\":\"_idSublink\",\"type\":\"uint16\"}],\"name\":\"forwardLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"updateExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listMessageContract\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"idSublink\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalLinkContract\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"mapContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_new_gateway\",\"type\":\"string\"}],\"name\":\"updateGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mapMember\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_defautGateway\",\"type\":\"string\"},{\"name\":\"_rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdateExchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_gateway\",\"type\":\"string\"}],\"name\":\"UpdateGateway\",\"type\":\"event\"}]"

// BaseContractBin is the compiled bytecode used for deploying new contracts.
const BaseContractBin = `0x608060405234801561001057600080fd5b5060405161090e38038061090e83398101604052805160208083015160008054600160a060020a0319163317905591909201805190926100559160019185019061005f565b50600255506100fa565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a057805160ff19168380011785556100cd565b828001600101855582156100cd579182015b828111156100cd5782518255916020019190600101906100b2565b506100d99291506100dd565b5090565b6100f791905b808211156100d957600081556001016100e3565b90565b610805806101096000396000f3006080604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663027e749981146100c65780630b8f6aa1146101505780630df4c1e3146101945780632c4e722e146101c0578063322d5a8e146101e7578063652db798146101fc5780637f7c2c0714610223578063816b18d61461025157806382dce8f914610269578063972cc2e7146102b3578063c659c8b1146102c8578063e7261687146102e4578063ebc0a6481461033d575b005b3480156100d257600080fd5b506100db610355565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101155781810151838201526020016100fd565b50505050905090810190601f1680156101425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015c57600080fd5b50610178600160a060020a036004351661ffff602435166103e2565b60408051600160a060020a039092168252519081900360200190f35b3480156101a057600080fd5b506101a9610408565b6040805161ffff9092168252519081900360200190f35b3480156101cc57600080fd5b506101d5610418565b60408051918252519081900360200190f35b3480156101f357600080fd5b5061017861041e565b34801561020857600080fd5b506100c4600160a060020a036004358116906024351661042d565b34801561022f57600080fd5b506100c4600160a060020a036004358116906024351661ffff604435166104ab565b34801561025d57600080fd5b506100c460043561055a565b34801561027557600080fd5b5061028d600160a060020a03600435166024356105ac565b60408051600160a060020a03909316835261ffff90911660208301528051918290030190f35b3480156102bf57600080fd5b506101a9610600565b3480156102d457600080fd5b5061017861ffff6004351661060a565b3480156102f057600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100c49436949293602493928401919081908401838280828437509497506106259650505050505050565b34801561034957600080fd5b506101786004356106ff565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103da5780601f106103af576101008083540402835291602001916103da565b820191906000526020600020905b8154815290600101906020018083116103bd57829003601f168201915b505050505081565b6005602090815260009283526040808420909152908252902054600160a060020a031681565b60035462010000900461ffff1681565b60025481565b600054600160a060020a031681565b6003805461ffff198116600161ffff9283160182161780835581166000908152600660209081526040808320805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03998a16908117909255969097168352600582528083209454909316825292909252902080549091169091179055565b6104b3610727565b50604080518082018252600160a060020a03948516815261ffff92831660208083019182529486166000908152600786529283208054600181018255908452949092209051930180549151909216740100000000000000000000000000000000000000000275ffff0000000000000000000000000000000000000000199390941673ffffffffffffffffffffffffffffffffffffffff199091161791909116919091179055565b600054600160a060020a0316331461057157600080fd5b60028190556040805182815290517f17b9b1d64f4a94661439177f6b95bce47cec5dc9916abf651cbca2c20ce38ca49181900360200190a150565b6007602052816000526040600020818154811015156105c757fe5b600091825260209091200154600160a060020a038116925074010000000000000000000000000000000000000000900461ffff16905082565b60035461ffff1681565b600660205260009081526040902054600160a060020a031681565b600054600160a060020a0316331461063c57600080fd5b805161064f90600190602084019061073e565b50604080516020808252600180546002600019610100838516150201909116049183018290527f2239fb6cf9cb0d991f7b802000e32aedca08db23cf392b94b50772d63f76272e939092918291820190849080156106ee5780601f106106c3576101008083540402835291602001916106ee565b820191906000526020600020905b8154815290600101906020018083116106d157829003601f168201915b50509250505060405180910390a150565b600480548290811061070d57fe5b600091825260209091200154600160a060020a0316905081565b604080518082019091526000808252602082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061077f57805160ff19168380011785556107ac565b828001600101855582156107ac579182015b828111156107ac578251825591602001919060010190610791565b506107b89291506107bc565b5090565b6107d691905b808211156107b857600081556001016107c2565b905600a165627a7a723058201956f433d5cf2cb0427f00ebcacebfa51a689b24a61a77464f7a2cf0315354140029`

// DeployBaseContract deploys a new Ethereum contract, binding an instance of BaseContract to it.
func DeployBaseContract(auth *bind.TransactOpts, backend bind.ContractBackend, _defautGateway string, _rate *big.Int) (common.Address, *types.Transaction, *BaseContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseContractBin), backend, _defautGateway, _rate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseContract{BaseContractCaller: BaseContractCaller{contract: contract}, BaseContractTransactor: BaseContractTransactor{contract: contract}, BaseContractFilterer: BaseContractFilterer{contract: contract}}, nil
}

// BaseContract is an auto generated Go binding around an Ethereum contract.
type BaseContract struct {
	BaseContractCaller     // Read-only binding to the contract
	BaseContractTransactor // Write-only binding to the contract
	BaseContractFilterer   // Log filterer for contract events
}

// BaseContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseContractSession struct {
	Contract     *BaseContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseContractCallerSession struct {
	Contract *BaseContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseContractTransactorSession struct {
	Contract     *BaseContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseContractRaw struct {
	Contract *BaseContract // Generic contract binding to access the raw methods on
}

// BaseContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseContractCallerRaw struct {
	Contract *BaseContractCaller // Generic read-only contract binding to access the raw methods on
}

// BaseContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseContractTransactorRaw struct {
	Contract *BaseContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseContract creates a new instance of BaseContract, bound to a specific deployed contract.
func NewBaseContract(address common.Address, backend bind.ContractBackend) (*BaseContract, error) {
	contract, err := bindBaseContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseContract{BaseContractCaller: BaseContractCaller{contract: contract}, BaseContractTransactor: BaseContractTransactor{contract: contract}, BaseContractFilterer: BaseContractFilterer{contract: contract}}, nil
}

// NewBaseContractCaller creates a new read-only instance of BaseContract, bound to a specific deployed contract.
func NewBaseContractCaller(address common.Address, caller bind.ContractCaller) (*BaseContractCaller, error) {
	contract, err := bindBaseContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContractCaller{contract: contract}, nil
}

// NewBaseContractTransactor creates a new write-only instance of BaseContract, bound to a specific deployed contract.
func NewBaseContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseContractTransactor, error) {
	contract, err := bindBaseContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContractTransactor{contract: contract}, nil
}

// NewBaseContractFilterer creates a new log filterer instance of BaseContract, bound to a specific deployed contract.
func NewBaseContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseContractFilterer, error) {
	contract, err := bindBaseContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseContractFilterer{contract: contract}, nil
}

// bindBaseContract binds a generic wrapper to an already deployed contract.
func bindBaseContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContract *BaseContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContract.Contract.BaseContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContract *BaseContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContract.Contract.BaseContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContract *BaseContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContract.Contract.BaseContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContract *BaseContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContract *BaseContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContract *BaseContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContract.Contract.contract.Transact(opts, method, params...)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x027e7499.
//
// Solidity: function default_gateway() constant returns(string)
func (_BaseContract *BaseContractCaller) DefaultGateway(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "default_gateway")
	return *ret0, err
}

// DefaultGateway is a free data retrieval call binding the contract method 0x027e7499.
//
// Solidity: function default_gateway() constant returns(string)
func (_BaseContract *BaseContractSession) DefaultGateway() (string, error) {
	return _BaseContract.Contract.DefaultGateway(&_BaseContract.CallOpts)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x027e7499.
//
// Solidity: function default_gateway() constant returns(string)
func (_BaseContract *BaseContractCallerSession) DefaultGateway() (string, error) {
	return _BaseContract.Contract.DefaultGateway(&_BaseContract.CallOpts)
}

// ListMessageContract is a free data retrieval call binding the contract method 0x82dce8f9.
//
// Solidity: function listMessageContract( address,  uint256) constant returns(addr address, idSublink uint16)
func (_BaseContract *BaseContractCaller) ListMessageContract(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Addr      common.Address
	IdSublink uint16
}, error) {
	ret := new(struct {
		Addr      common.Address
		IdSublink uint16
	})
	out := ret
	err := _BaseContract.contract.Call(opts, out, "listMessageContract", arg0, arg1)
	return *ret, err
}

// ListMessageContract is a free data retrieval call binding the contract method 0x82dce8f9.
//
// Solidity: function listMessageContract( address,  uint256) constant returns(addr address, idSublink uint16)
func (_BaseContract *BaseContractSession) ListMessageContract(arg0 common.Address, arg1 *big.Int) (struct {
	Addr      common.Address
	IdSublink uint16
}, error) {
	return _BaseContract.Contract.ListMessageContract(&_BaseContract.CallOpts, arg0, arg1)
}

// ListMessageContract is a free data retrieval call binding the contract method 0x82dce8f9.
//
// Solidity: function listMessageContract( address,  uint256) constant returns(addr address, idSublink uint16)
func (_BaseContract *BaseContractCallerSession) ListMessageContract(arg0 common.Address, arg1 *big.Int) (struct {
	Addr      common.Address
	IdSublink uint16
}, error) {
	return _BaseContract.Contract.ListMessageContract(&_BaseContract.CallOpts, arg0, arg1)
}

// MapContract is a free data retrieval call binding the contract method 0xc659c8b1.
//
// Solidity: function mapContract( uint16) constant returns(address)
func (_BaseContract *BaseContractCaller) MapContract(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "mapContract", arg0)
	return *ret0, err
}

// MapContract is a free data retrieval call binding the contract method 0xc659c8b1.
//
// Solidity: function mapContract( uint16) constant returns(address)
func (_BaseContract *BaseContractSession) MapContract(arg0 uint16) (common.Address, error) {
	return _BaseContract.Contract.MapContract(&_BaseContract.CallOpts, arg0)
}

// MapContract is a free data retrieval call binding the contract method 0xc659c8b1.
//
// Solidity: function mapContract( uint16) constant returns(address)
func (_BaseContract *BaseContractCallerSession) MapContract(arg0 uint16) (common.Address, error) {
	return _BaseContract.Contract.MapContract(&_BaseContract.CallOpts, arg0)
}

// MapLinkContract is a free data retrieval call binding the contract method 0x0b8f6aa1.
//
// Solidity: function mapLinkContract( address,  uint16) constant returns(address)
func (_BaseContract *BaseContractCaller) MapLinkContract(opts *bind.CallOpts, arg0 common.Address, arg1 uint16) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "mapLinkContract", arg0, arg1)
	return *ret0, err
}

// MapLinkContract is a free data retrieval call binding the contract method 0x0b8f6aa1.
//
// Solidity: function mapLinkContract( address,  uint16) constant returns(address)
func (_BaseContract *BaseContractSession) MapLinkContract(arg0 common.Address, arg1 uint16) (common.Address, error) {
	return _BaseContract.Contract.MapLinkContract(&_BaseContract.CallOpts, arg0, arg1)
}

// MapLinkContract is a free data retrieval call binding the contract method 0x0b8f6aa1.
//
// Solidity: function mapLinkContract( address,  uint16) constant returns(address)
func (_BaseContract *BaseContractCallerSession) MapLinkContract(arg0 common.Address, arg1 uint16) (common.Address, error) {
	return _BaseContract.Contract.MapLinkContract(&_BaseContract.CallOpts, arg0, arg1)
}

// MapMember is a free data retrieval call binding the contract method 0xebc0a648.
//
// Solidity: function mapMember( uint256) constant returns(address)
func (_BaseContract *BaseContractCaller) MapMember(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "mapMember", arg0)
	return *ret0, err
}

// MapMember is a free data retrieval call binding the contract method 0xebc0a648.
//
// Solidity: function mapMember( uint256) constant returns(address)
func (_BaseContract *BaseContractSession) MapMember(arg0 *big.Int) (common.Address, error) {
	return _BaseContract.Contract.MapMember(&_BaseContract.CallOpts, arg0)
}

// MapMember is a free data retrieval call binding the contract method 0xebc0a648.
//
// Solidity: function mapMember( uint256) constant returns(address)
func (_BaseContract *BaseContractCallerSession) MapMember(arg0 *big.Int) (common.Address, error) {
	return _BaseContract.Contract.MapMember(&_BaseContract.CallOpts, arg0)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BaseContract *BaseContractCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BaseContract *BaseContractSession) Rate() (*big.Int, error) {
	return _BaseContract.Contract.Rate(&_BaseContract.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BaseContract *BaseContractCallerSession) Rate() (*big.Int, error) {
	return _BaseContract.Contract.Rate(&_BaseContract.CallOpts)
}

// SuperOwner is a free data retrieval call binding the contract method 0x322d5a8e.
//
// Solidity: function super_owner() constant returns(address)
func (_BaseContract *BaseContractCaller) SuperOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "super_owner")
	return *ret0, err
}

// SuperOwner is a free data retrieval call binding the contract method 0x322d5a8e.
//
// Solidity: function super_owner() constant returns(address)
func (_BaseContract *BaseContractSession) SuperOwner() (common.Address, error) {
	return _BaseContract.Contract.SuperOwner(&_BaseContract.CallOpts)
}

// SuperOwner is a free data retrieval call binding the contract method 0x322d5a8e.
//
// Solidity: function super_owner() constant returns(address)
func (_BaseContract *BaseContractCallerSession) SuperOwner() (common.Address, error) {
	return _BaseContract.Contract.SuperOwner(&_BaseContract.CallOpts)
}

// TotalLinkContract is a free data retrieval call binding the contract method 0x972cc2e7.
//
// Solidity: function totalLinkContract() constant returns(uint16)
func (_BaseContract *BaseContractCaller) TotalLinkContract(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "totalLinkContract")
	return *ret0, err
}

// TotalLinkContract is a free data retrieval call binding the contract method 0x972cc2e7.
//
// Solidity: function totalLinkContract() constant returns(uint16)
func (_BaseContract *BaseContractSession) TotalLinkContract() (uint16, error) {
	return _BaseContract.Contract.TotalLinkContract(&_BaseContract.CallOpts)
}

// TotalLinkContract is a free data retrieval call binding the contract method 0x972cc2e7.
//
// Solidity: function totalLinkContract() constant returns(uint16)
func (_BaseContract *BaseContractCallerSession) TotalLinkContract() (uint16, error) {
	return _BaseContract.Contract.TotalLinkContract(&_BaseContract.CallOpts)
}

// TotalMessage is a free data retrieval call binding the contract method 0x0df4c1e3.
//
// Solidity: function totalMessage() constant returns(uint16)
func (_BaseContract *BaseContractCaller) TotalMessage(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _BaseContract.contract.Call(opts, out, "totalMessage")
	return *ret0, err
}

// TotalMessage is a free data retrieval call binding the contract method 0x0df4c1e3.
//
// Solidity: function totalMessage() constant returns(uint16)
func (_BaseContract *BaseContractSession) TotalMessage() (uint16, error) {
	return _BaseContract.Contract.TotalMessage(&_BaseContract.CallOpts)
}

// TotalMessage is a free data retrieval call binding the contract method 0x0df4c1e3.
//
// Solidity: function totalMessage() constant returns(uint16)
func (_BaseContract *BaseContractCallerSession) TotalMessage() (uint16, error) {
	return _BaseContract.Contract.TotalMessage(&_BaseContract.CallOpts)
}

// ForwardLink is a paid mutator transaction binding the contract method 0x7f7c2c07.
//
// Solidity: function forwardLink(_addrContract address, _addrRecive address, _idSublink uint16) returns()
func (_BaseContract *BaseContractTransactor) ForwardLink(opts *bind.TransactOpts, _addrContract common.Address, _addrRecive common.Address, _idSublink uint16) (*types.Transaction, error) {
	return _BaseContract.contract.Transact(opts, "forwardLink", _addrContract, _addrRecive, _idSublink)
}

// ForwardLink is a paid mutator transaction binding the contract method 0x7f7c2c07.
//
// Solidity: function forwardLink(_addrContract address, _addrRecive address, _idSublink uint16) returns()
func (_BaseContract *BaseContractSession) ForwardLink(_addrContract common.Address, _addrRecive common.Address, _idSublink uint16) (*types.Transaction, error) {
	return _BaseContract.Contract.ForwardLink(&_BaseContract.TransactOpts, _addrContract, _addrRecive, _idSublink)
}

// ForwardLink is a paid mutator transaction binding the contract method 0x7f7c2c07.
//
// Solidity: function forwardLink(_addrContract address, _addrRecive address, _idSublink uint16) returns()
func (_BaseContract *BaseContractTransactorSession) ForwardLink(_addrContract common.Address, _addrRecive common.Address, _idSublink uint16) (*types.Transaction, error) {
	return _BaseContract.Contract.ForwardLink(&_BaseContract.TransactOpts, _addrContract, _addrRecive, _idSublink)
}

// SaveContract is a paid mutator transaction binding the contract method 0x652db798.
//
// Solidity: function saveContract(_contractLink address, _ownerContractLink address) returns()
func (_BaseContract *BaseContractTransactor) SaveContract(opts *bind.TransactOpts, _contractLink common.Address, _ownerContractLink common.Address) (*types.Transaction, error) {
	return _BaseContract.contract.Transact(opts, "saveContract", _contractLink, _ownerContractLink)
}

// SaveContract is a paid mutator transaction binding the contract method 0x652db798.
//
// Solidity: function saveContract(_contractLink address, _ownerContractLink address) returns()
func (_BaseContract *BaseContractSession) SaveContract(_contractLink common.Address, _ownerContractLink common.Address) (*types.Transaction, error) {
	return _BaseContract.Contract.SaveContract(&_BaseContract.TransactOpts, _contractLink, _ownerContractLink)
}

// SaveContract is a paid mutator transaction binding the contract method 0x652db798.
//
// Solidity: function saveContract(_contractLink address, _ownerContractLink address) returns()
func (_BaseContract *BaseContractTransactorSession) SaveContract(_contractLink common.Address, _ownerContractLink common.Address) (*types.Transaction, error) {
	return _BaseContract.Contract.SaveContract(&_BaseContract.TransactOpts, _contractLink, _ownerContractLink)
}

// UpdateExchange is a paid mutator transaction binding the contract method 0x816b18d6.
//
// Solidity: function updateExchange(_newRate uint256) returns()
func (_BaseContract *BaseContractTransactor) UpdateExchange(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _BaseContract.contract.Transact(opts, "updateExchange", _newRate)
}

// UpdateExchange is a paid mutator transaction binding the contract method 0x816b18d6.
//
// Solidity: function updateExchange(_newRate uint256) returns()
func (_BaseContract *BaseContractSession) UpdateExchange(_newRate *big.Int) (*types.Transaction, error) {
	return _BaseContract.Contract.UpdateExchange(&_BaseContract.TransactOpts, _newRate)
}

// UpdateExchange is a paid mutator transaction binding the contract method 0x816b18d6.
//
// Solidity: function updateExchange(_newRate uint256) returns()
func (_BaseContract *BaseContractTransactorSession) UpdateExchange(_newRate *big.Int) (*types.Transaction, error) {
	return _BaseContract.Contract.UpdateExchange(&_BaseContract.TransactOpts, _newRate)
}

// UpdateGateway is a paid mutator transaction binding the contract method 0xe7261687.
//
// Solidity: function updateGateway(_new_gateway string) returns()
func (_BaseContract *BaseContractTransactor) UpdateGateway(opts *bind.TransactOpts, _new_gateway string) (*types.Transaction, error) {
	return _BaseContract.contract.Transact(opts, "updateGateway", _new_gateway)
}

// UpdateGateway is a paid mutator transaction binding the contract method 0xe7261687.
//
// Solidity: function updateGateway(_new_gateway string) returns()
func (_BaseContract *BaseContractSession) UpdateGateway(_new_gateway string) (*types.Transaction, error) {
	return _BaseContract.Contract.UpdateGateway(&_BaseContract.TransactOpts, _new_gateway)
}

// UpdateGateway is a paid mutator transaction binding the contract method 0xe7261687.
//
// Solidity: function updateGateway(_new_gateway string) returns()
func (_BaseContract *BaseContractTransactorSession) UpdateGateway(_new_gateway string) (*types.Transaction, error) {
	return _BaseContract.Contract.UpdateGateway(&_BaseContract.TransactOpts, _new_gateway)
}

// BaseContractUpdateExchangeIterator is returned from FilterUpdateExchange and is used to iterate over the raw logs and unpacked data for UpdateExchange events raised by the BaseContract contract.
type BaseContractUpdateExchangeIterator struct {
	Event *BaseContractUpdateExchange // Event containing the contract specifics and raw log

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
func (it *BaseContractUpdateExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContractUpdateExchange)
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
		it.Event = new(BaseContractUpdateExchange)
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
func (it *BaseContractUpdateExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContractUpdateExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContractUpdateExchange represents a UpdateExchange event raised by the BaseContract contract.
type BaseContractUpdateExchange struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateExchange is a free log retrieval operation binding the contract event 0x17b9b1d64f4a94661439177f6b95bce47cec5dc9916abf651cbca2c20ce38ca4.
//
// Solidity: e UpdateExchange(_rate uint256)
func (_BaseContract *BaseContractFilterer) FilterUpdateExchange(opts *bind.FilterOpts) (*BaseContractUpdateExchangeIterator, error) {

	logs, sub, err := _BaseContract.contract.FilterLogs(opts, "UpdateExchange")
	if err != nil {
		return nil, err
	}
	return &BaseContractUpdateExchangeIterator{contract: _BaseContract.contract, event: "UpdateExchange", logs: logs, sub: sub}, nil
}

// WatchUpdateExchange is a free log subscription operation binding the contract event 0x17b9b1d64f4a94661439177f6b95bce47cec5dc9916abf651cbca2c20ce38ca4.
//
// Solidity: e UpdateExchange(_rate uint256)
func (_BaseContract *BaseContractFilterer) WatchUpdateExchange(opts *bind.WatchOpts, sink chan<- *BaseContractUpdateExchange) (event.Subscription, error) {

	logs, sub, err := _BaseContract.contract.WatchLogs(opts, "UpdateExchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContractUpdateExchange)
				if err := _BaseContract.contract.UnpackLog(event, "UpdateExchange", log); err != nil {
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

// BaseContractUpdateGatewayIterator is returned from FilterUpdateGateway and is used to iterate over the raw logs and unpacked data for UpdateGateway events raised by the BaseContract contract.
type BaseContractUpdateGatewayIterator struct {
	Event *BaseContractUpdateGateway // Event containing the contract specifics and raw log

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
func (it *BaseContractUpdateGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContractUpdateGateway)
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
		it.Event = new(BaseContractUpdateGateway)
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
func (it *BaseContractUpdateGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContractUpdateGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContractUpdateGateway represents a UpdateGateway event raised by the BaseContract contract.
type BaseContractUpdateGateway struct {
	Gateway string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateGateway is a free log retrieval operation binding the contract event 0x2239fb6cf9cb0d991f7b802000e32aedca08db23cf392b94b50772d63f76272e.
//
// Solidity: e UpdateGateway(_gateway string)
func (_BaseContract *BaseContractFilterer) FilterUpdateGateway(opts *bind.FilterOpts) (*BaseContractUpdateGatewayIterator, error) {

	logs, sub, err := _BaseContract.contract.FilterLogs(opts, "UpdateGateway")
	if err != nil {
		return nil, err
	}
	return &BaseContractUpdateGatewayIterator{contract: _BaseContract.contract, event: "UpdateGateway", logs: logs, sub: sub}, nil
}

// WatchUpdateGateway is a free log subscription operation binding the contract event 0x2239fb6cf9cb0d991f7b802000e32aedca08db23cf392b94b50772d63f76272e.
//
// Solidity: e UpdateGateway(_gateway string)
func (_BaseContract *BaseContractFilterer) WatchUpdateGateway(opts *bind.WatchOpts, sink chan<- *BaseContractUpdateGateway) (event.Subscription, error) {

	logs, sub, err := _BaseContract.contract.WatchLogs(opts, "UpdateGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContractUpdateGateway)
				if err := _BaseContract.contract.UnpackLog(event, "UpdateGateway", log); err != nil {
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

// LinkContractABI is the input ABI used to generate the binding from.
const LinkContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"default_gateway\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subLinks\",\"outputs\":[{\"name\":\"id\",\"type\":\"int32\"},{\"name\":\"addr_create\",\"type\":\"address\"},{\"name\":\"addr_parent\",\"type\":\"address\"},{\"name\":\"current_contract\",\"type\":\"address\"},{\"name\":\"link\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_child\",\"type\":\"address\"},{\"name\":\"_parent\",\"type\":\"address\"}],\"name\":\"onNewClick\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"invest_eth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"link\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"super_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxClick\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"members\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalClicked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalLinkGenerared\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"mapLinkGenerated\",\"outputs\":[{\"name\":\"id\",\"type\":\"int32\"},{\"name\":\"addr_create\",\"type\":\"address\"},{\"name\":\"addr_parent\",\"type\":\"address\"},{\"name\":\"current_contract\",\"type\":\"address\"},{\"name\":\"link\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mapStatisticalClick\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_base\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parent\",\"type\":\"address\"},{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"generateLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"process\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mapClick\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addrRecive\",\"type\":\"address\"},{\"name\":\"_linkId\",\"type\":\"uint16\"}],\"name\":\"forwardLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_contract_base\",\"type\":\"address\"},{\"name\":\"_link\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addrRecive\",\"type\":\"address\"}],\"name\":\"ForwardTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpContractLink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_parent\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_link\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_gateway\",\"type\":\"string\"}],\"name\":\"GenerateLinkContract\",\"type\":\"event\"}]"

// LinkContractBin is the compiled bytecode used for deploying new contracts.
const LinkContractBin = `0x60806040523480156200001157600080fd5b50604051620015ca380380620015ca83398101604052805160208083015160038054600160a060020a038516600160a060020a0319918216179091556001805490911633179055909201805191929091620000739160059190840190620000b6565b50600a805462ff00001916620100001790556040517f784a9cc22598ebcd47d2d416a4fb20c43dfbbce49ad2368a7e2fd4a8a78dc2ba90600090a150506200015b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f957805160ff191683800117855562000129565b8280016001018555821562000129579182015b82811115620001295782518255916020019190600101906200010c565b50620001379291506200013b565b5090565b6200015891905b8082111562000137576000815560010162000142565b90565b61145f806200016b6000396000f3006080604052600436106101115763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663027e7499811461044e57806302b06ab1146104d8578063039243a5146105c057806305ef5229146105e95780631c4695f414610610578063200d2ed214610625578063322d5a8e1461064e57806342f6487a1461067f5780635410fcc6146106945780635daf08ca146106a957806369f3619c146106c157806371d10ee4146106d65780638da5cb5b1461070257806394fa6da5146107175780639e5e865714610733578063ad291c7b14610754578063b7b18a8814610769578063c33fb87714610790578063d238011e146107a5578063f6c655df146107c6575b6001546000908190600160a060020a0316331461012d57600080fd5b6003546000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392831617808255604080517f322d5a8e0000000000000000000000000000000000000000000000000000000081529051919093169263322d5a8e9260048083019360209390929083900390910190829087803b1580156101b257600080fd5b505af11580156101c6573d6000803e3d6000fd5b505050506040513d60208110156101dc57600080fd5b50516002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392831617905560008054604080517f027e74990000000000000000000000000000000000000000000000000000000081529051919093169263027e749992600480830193919282900301818387803b15801561025c57600080fd5b505af1158015610270573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561029957600080fd5b8101908080516401000000008111156102b157600080fd5b820160208101848111156102c457600080fd5b81516401000000008111828201871017156102de57600080fd5b505080516102f79450600693506020909101915061136b565b506000809054906101000a9004600160a060020a0316600160a060020a0316632c4e722e6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561036357600080fd5b505af1158015610377573d6000803e3d6000fd5b505050506040513d602081101561038d57600080fd5b505160008054604080517f652db7980000000000000000000000000000000000000000000000000000000081523060048201523360248201529051939550600160a060020a039091169263652db7989260448084019391929182900301818387803b1580156103fb57600080fd5b505af115801561040f573d6000803e3d6000fd5b50506004805434019081905560649250670de0b6b3a764000090048402605a02905004600955600454606490600a0204905061044a816107ee565b5050005b34801561045a57600080fd5b5061046361082c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561049d578181015183820152602001610485565b50505050905090810190601f1680156104ca5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104e457600080fd5b506104f06004356108ba565b604051808660030b60030b815260200185600160a060020a0316600160a060020a0316815260200184600160a060020a0316600160a060020a0316815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610581578181015183820152602001610569565b50505050905090810190601f1680156105ae5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156105cc57600080fd5b506105e7600160a060020a036004358116906024351661099f565b005b3480156105f557600080fd5b506105fe610a50565b60408051918252519081900360200190f35b34801561061c57600080fd5b50610463610a56565b34801561063157600080fd5b5061063a610ab1565b604080519115158252519081900360200190f35b34801561065a57600080fd5b50610663610ac0565b60408051600160a060020a039092168252519081900360200190f35b34801561068b57600080fd5b506105e7610acf565b3480156106a057600080fd5b506105fe610c41565b3480156106b557600080fd5b50610663600435610c47565b3480156106cd57600080fd5b506105fe610c6f565b3480156106e257600080fd5b506106eb610c75565b6040805161ffff9092168252519081900360200190f35b34801561070e57600080fd5b50610663610c7f565b34801561072357600080fd5b506104f061ffff60043516610c8e565b34801561073f57600080fd5b506105fe600160a060020a0360043516610d27565b34801561076057600080fd5b50610663610d39565b34801561077557600080fd5b506105e7600160a060020a0360043581169060243516610d48565b34801561079c57600080fd5b506105fe611214565b3480156107b157600080fd5b506105fe600160a060020a036004351661121a565b3480156107d257600080fd5b506105e7600160a060020a036004351661ffff6024351661122c565b600254604051600160a060020a039091169082156108fc029083906000818181858888f19350505050158015610828573d6000803e3d6000fd5b5050565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108b25780601f10610887576101008083540402835291602001916108b2565b820191906000526020600020905b81548152906001019060200180831161089557829003601f168201915b505050505081565b600c8054829081106108c857fe5b6000918252602091829020600491909102018054600180830154600280850154600380870180546040805161010098831615989098026000190190911694909404601f81018a90048a0287018a019094528386529086900b9850640100000000909504600160a060020a0390811697938116969116949093929091908301828280156109955780601f1061096a57610100808354040283529160200191610995565b820191906000526020600020905b81548152906001019060200180831161097857829003601f168201915b5050505050905085565b600254600160a060020a031633146109b657600080fd5b600a5462010000900460ff1615156001146109d057600080fd5b600160a060020a038281166000818152600d602090815260408083208054600a0181559486168084528184208054600190810182559585529554600e90935281842092909255908252925492902091909155600880549091019055610a33611333565b1561082857610a40610acf565b600a805462ff0000191690555050565b60045481565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108b25780601f10610887576101008083540402835291602001916108b2565b600a5462010000900460ff1681565b600254600160a060020a031681565b600b5460095430319060009081901515610ae857610c3b565b600091505b83821015610c3b57610b44600954600a0284600d6000600b87815481101515610b1257fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902054029063ffffffff61135416565b9050600b82815481101515610b5557fe5b6000918252602082200154604051600160a060020a039091169183156108fc02918491818181858888f19350505050158015610b95573d6000803e3d6000fd5b506000600d6000600b85815481101515610bab57fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902055600b80547fd4f43975feb89f48dd30cabbb32011045be187d1e11c8ea9faa43efc35282519919084908110610c0157fe5b6000918252602091829020015460408051600160a060020a03909216825291810184905281519081900390910190a1600190910190610aed565b50505050565b60095481565b600b805482908110610c5557fe5b600091825260209091200154600160a060020a0316905081565b60085481565b600a5461ffff1681565b600154600160a060020a031681565b600f602090815260009182526040918290208054600180830154600280850154600380870180548a516101009782161597909702600019011693909304601f810189900489028601890190995288855285900b97640100000000909504600160a060020a0390811697938116969116949093928301828280156109955780601f1061096a57610100808354040283529160200191610995565b600e6020526000908152604090205481565b600354600160a060020a031681565b610d506113e9565b600154600160a060020a0316331415610d6857600080fd5b600a5462010000900460ff161515600114610d8257600080fd5b6040805160a081018252600a5461ffff1660030b815233602080830191909152600160a060020a03868116838501528516606083015260058054845160026101006001841615026000190190921691909104601f81018490048402820184019095528481529293608085019392830182828015610e405780601f10610e1557610100808354040283529160200191610e40565b820191906000526020600020905b815481529060010190602001808311610e2357829003601f168201915b505050919092525050600a5461ffff166000908152600f6020908152604091829020835181548584015163ffffffff1990911663ffffffff600393840b161777ffffffffffffffffffffffffffffffffffffffff000000001916640100000000600160a060020a03928316021783559385015160018301805473ffffffffffffffffffffffffffffffffffffffff1990811692871692909217905560608601516002840180549092169516949094179093556080840151805194955085949193610f10939085019291019061136b565b5050600b805460018082019092557f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db90180543373ffffffffffffffffffffffffffffffffffffffff1991821617909155600c805492830180825560009190915284517fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c76004909402938401805460208089015163ffffffff1990921663ffffffff60039590950b949094169390931777ffffffffffffffffffffffffffffffffffffffff000000001916640100000000600160a060020a039283160217825560408801517fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c887018054871691831691909117905560608801517fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c987018054909616911617909355608086015180519295508694611096937fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8ca90910192919091019061136b565b5050600a805461ffff198116600161ffff9283168101909216179091556040805133808252600160a060020a03898116602084015288169282019290925260a06060820181815260058054600261010097821615979097026000190116959095049183018290527f0a13894e900be98ffdd0efbe75397f8621ce5762f64a7f2a25795a797b7445de9650929450889388939092600692909190608083019060c0840190869080156111885780601f1061115d57610100808354040283529160200191611188565b820191906000526020600020905b81548152906001019060200180831161116b57829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156111fc5780601f106111d1576101008083540402835291602001916111fc565b820191906000526020600020905b8154815290600101906020018083116111df57829003601f168201915b505097505050505050505060405180910390a1505050565b60075481565b600d6020526000908152604090205481565b600154600160a060020a031633141561124457600080fd5b6003546000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392831617808255604080517f7f7c2c07000000000000000000000000000000000000000000000000000000008152306004820152868516602482015261ffff8616604482015290519190931692637f7c2c0792606480830193919282900301818387803b1580156112da57600080fd5b505af11580156112ee573d6000803e3d6000fd5b505060408051600160a060020a038616815290517fb708fb2f4fb705069ac2d34f4057a64fa1fc9246fdc075f7cdf818a25e80415e9350908190036020019150a15050565b6000600954600a0260085410151561134d57506001611351565b5060005b90565b600080828481151561136257fe5b04949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113ac57805160ff19168380011785556113d9565b828001600101855582156113d9579182015b828111156113d95782518255916020019190600101906113be565b506113e5929150611419565b5090565b6040805160a081018252600080825260208201819052918101829052606080820192909252608081019190915290565b61135191905b808211156113e5576000815560010161141f5600a165627a7a723058208c4e162d4557d1c3bf1832eb87c717657dd3bf76da874ba82d579ef2aeeeb4c00029`

// DeployLinkContract deploys a new Ethereum contract, binding an instance of LinkContract to it.
func DeployLinkContract(auth *bind.TransactOpts, backend bind.ContractBackend, _contract_base common.Address, _link string) (common.Address, *types.Transaction, *LinkContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LinkContractBin), backend, _contract_base, _link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LinkContract{LinkContractCaller: LinkContractCaller{contract: contract}, LinkContractTransactor: LinkContractTransactor{contract: contract}, LinkContractFilterer: LinkContractFilterer{contract: contract}}, nil
}

// LinkContract is an auto generated Go binding around an Ethereum contract.
type LinkContract struct {
	LinkContractCaller     // Read-only binding to the contract
	LinkContractTransactor // Write-only binding to the contract
	LinkContractFilterer   // Log filterer for contract events
}

// LinkContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinkContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinkContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinkContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinkContractSession struct {
	Contract     *LinkContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LinkContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinkContractCallerSession struct {
	Contract *LinkContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LinkContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinkContractTransactorSession struct {
	Contract     *LinkContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LinkContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinkContractRaw struct {
	Contract *LinkContract // Generic contract binding to access the raw methods on
}

// LinkContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinkContractCallerRaw struct {
	Contract *LinkContractCaller // Generic read-only contract binding to access the raw methods on
}

// LinkContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinkContractTransactorRaw struct {
	Contract *LinkContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinkContract creates a new instance of LinkContract, bound to a specific deployed contract.
func NewLinkContract(address common.Address, backend bind.ContractBackend) (*LinkContract, error) {
	contract, err := bindLinkContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkContract{LinkContractCaller: LinkContractCaller{contract: contract}, LinkContractTransactor: LinkContractTransactor{contract: contract}, LinkContractFilterer: LinkContractFilterer{contract: contract}}, nil
}

// NewLinkContractCaller creates a new read-only instance of LinkContract, bound to a specific deployed contract.
func NewLinkContractCaller(address common.Address, caller bind.ContractCaller) (*LinkContractCaller, error) {
	contract, err := bindLinkContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkContractCaller{contract: contract}, nil
}

// NewLinkContractTransactor creates a new write-only instance of LinkContract, bound to a specific deployed contract.
func NewLinkContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkContractTransactor, error) {
	contract, err := bindLinkContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkContractTransactor{contract: contract}, nil
}

// NewLinkContractFilterer creates a new log filterer instance of LinkContract, bound to a specific deployed contract.
func NewLinkContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkContractFilterer, error) {
	contract, err := bindLinkContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkContractFilterer{contract: contract}, nil
}

// bindLinkContract binds a generic wrapper to an already deployed contract.
func bindLinkContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkContract *LinkContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkContract.Contract.LinkContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkContract *LinkContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkContract.Contract.LinkContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkContract *LinkContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkContract.Contract.LinkContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkContract *LinkContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkContract *LinkContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkContract *LinkContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkContract.Contract.contract.Transact(opts, method, params...)
}

// ContractBase is a free data retrieval call binding the contract method 0xad291c7b.
//
// Solidity: function contract_base() constant returns(address)
func (_LinkContract *LinkContractCaller) ContractBase(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "contract_base")
	return *ret0, err
}

// ContractBase is a free data retrieval call binding the contract method 0xad291c7b.
//
// Solidity: function contract_base() constant returns(address)
func (_LinkContract *LinkContractSession) ContractBase() (common.Address, error) {
	return _LinkContract.Contract.ContractBase(&_LinkContract.CallOpts)
}

// ContractBase is a free data retrieval call binding the contract method 0xad291c7b.
//
// Solidity: function contract_base() constant returns(address)
func (_LinkContract *LinkContractCallerSession) ContractBase() (common.Address, error) {
	return _LinkContract.Contract.ContractBase(&_LinkContract.CallOpts)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x027e7499.
//
// Solidity: function default_gateway() constant returns(string)
func (_LinkContract *LinkContractCaller) DefaultGateway(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "default_gateway")
	return *ret0, err
}

// DefaultGateway is a free data retrieval call binding the contract method 0x027e7499.
//
// Solidity: function default_gateway() constant returns(string)
func (_LinkContract *LinkContractSession) DefaultGateway() (string, error) {
	return _LinkContract.Contract.DefaultGateway(&_LinkContract.CallOpts)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x027e7499.
//
// Solidity: function default_gateway() constant returns(string)
func (_LinkContract *LinkContractCallerSession) DefaultGateway() (string, error) {
	return _LinkContract.Contract.DefaultGateway(&_LinkContract.CallOpts)
}

// InvestEth is a free data retrieval call binding the contract method 0x05ef5229.
//
// Solidity: function invest_eth() constant returns(uint256)
func (_LinkContract *LinkContractCaller) InvestEth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "invest_eth")
	return *ret0, err
}

// InvestEth is a free data retrieval call binding the contract method 0x05ef5229.
//
// Solidity: function invest_eth() constant returns(uint256)
func (_LinkContract *LinkContractSession) InvestEth() (*big.Int, error) {
	return _LinkContract.Contract.InvestEth(&_LinkContract.CallOpts)
}

// InvestEth is a free data retrieval call binding the contract method 0x05ef5229.
//
// Solidity: function invest_eth() constant returns(uint256)
func (_LinkContract *LinkContractCallerSession) InvestEth() (*big.Int, error) {
	return _LinkContract.Contract.InvestEth(&_LinkContract.CallOpts)
}

// Link is a free data retrieval call binding the contract method 0x1c4695f4.
//
// Solidity: function link() constant returns(string)
func (_LinkContract *LinkContractCaller) Link(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "link")
	return *ret0, err
}

// Link is a free data retrieval call binding the contract method 0x1c4695f4.
//
// Solidity: function link() constant returns(string)
func (_LinkContract *LinkContractSession) Link() (string, error) {
	return _LinkContract.Contract.Link(&_LinkContract.CallOpts)
}

// Link is a free data retrieval call binding the contract method 0x1c4695f4.
//
// Solidity: function link() constant returns(string)
func (_LinkContract *LinkContractCallerSession) Link() (string, error) {
	return _LinkContract.Contract.Link(&_LinkContract.CallOpts)
}

// MapClick is a free data retrieval call binding the contract method 0xd238011e.
//
// Solidity: function mapClick( address) constant returns(uint256)
func (_LinkContract *LinkContractCaller) MapClick(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "mapClick", arg0)
	return *ret0, err
}

// MapClick is a free data retrieval call binding the contract method 0xd238011e.
//
// Solidity: function mapClick( address) constant returns(uint256)
func (_LinkContract *LinkContractSession) MapClick(arg0 common.Address) (*big.Int, error) {
	return _LinkContract.Contract.MapClick(&_LinkContract.CallOpts, arg0)
}

// MapClick is a free data retrieval call binding the contract method 0xd238011e.
//
// Solidity: function mapClick( address) constant returns(uint256)
func (_LinkContract *LinkContractCallerSession) MapClick(arg0 common.Address) (*big.Int, error) {
	return _LinkContract.Contract.MapClick(&_LinkContract.CallOpts, arg0)
}

// MapLinkGenerated is a free data retrieval call binding the contract method 0x94fa6da5.
//
// Solidity: function mapLinkGenerated( uint16) constant returns(id int32, addr_create address, addr_parent address, current_contract address, link string)
func (_LinkContract *LinkContractCaller) MapLinkGenerated(opts *bind.CallOpts, arg0 uint16) (struct {
	Id              int32
	AddrCreate      common.Address
	AddrParent      common.Address
	CurrentContract common.Address
	Link            string
}, error) {
	ret := new(struct {
		Id              int32
		AddrCreate      common.Address
		AddrParent      common.Address
		CurrentContract common.Address
		Link            string
	})
	out := ret
	err := _LinkContract.contract.Call(opts, out, "mapLinkGenerated", arg0)
	return *ret, err
}

// MapLinkGenerated is a free data retrieval call binding the contract method 0x94fa6da5.
//
// Solidity: function mapLinkGenerated( uint16) constant returns(id int32, addr_create address, addr_parent address, current_contract address, link string)
func (_LinkContract *LinkContractSession) MapLinkGenerated(arg0 uint16) (struct {
	Id              int32
	AddrCreate      common.Address
	AddrParent      common.Address
	CurrentContract common.Address
	Link            string
}, error) {
	return _LinkContract.Contract.MapLinkGenerated(&_LinkContract.CallOpts, arg0)
}

// MapLinkGenerated is a free data retrieval call binding the contract method 0x94fa6da5.
//
// Solidity: function mapLinkGenerated( uint16) constant returns(id int32, addr_create address, addr_parent address, current_contract address, link string)
func (_LinkContract *LinkContractCallerSession) MapLinkGenerated(arg0 uint16) (struct {
	Id              int32
	AddrCreate      common.Address
	AddrParent      common.Address
	CurrentContract common.Address
	Link            string
}, error) {
	return _LinkContract.Contract.MapLinkGenerated(&_LinkContract.CallOpts, arg0)
}

// MapStatisticalClick is a free data retrieval call binding the contract method 0x9e5e8657.
//
// Solidity: function mapStatisticalClick( address) constant returns(uint256)
func (_LinkContract *LinkContractCaller) MapStatisticalClick(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "mapStatisticalClick", arg0)
	return *ret0, err
}

// MapStatisticalClick is a free data retrieval call binding the contract method 0x9e5e8657.
//
// Solidity: function mapStatisticalClick( address) constant returns(uint256)
func (_LinkContract *LinkContractSession) MapStatisticalClick(arg0 common.Address) (*big.Int, error) {
	return _LinkContract.Contract.MapStatisticalClick(&_LinkContract.CallOpts, arg0)
}

// MapStatisticalClick is a free data retrieval call binding the contract method 0x9e5e8657.
//
// Solidity: function mapStatisticalClick( address) constant returns(uint256)
func (_LinkContract *LinkContractCallerSession) MapStatisticalClick(arg0 common.Address) (*big.Int, error) {
	return _LinkContract.Contract.MapStatisticalClick(&_LinkContract.CallOpts, arg0)
}

// MaxClick is a free data retrieval call binding the contract method 0x5410fcc6.
//
// Solidity: function maxClick() constant returns(uint256)
func (_LinkContract *LinkContractCaller) MaxClick(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "maxClick")
	return *ret0, err
}

// MaxClick is a free data retrieval call binding the contract method 0x5410fcc6.
//
// Solidity: function maxClick() constant returns(uint256)
func (_LinkContract *LinkContractSession) MaxClick() (*big.Int, error) {
	return _LinkContract.Contract.MaxClick(&_LinkContract.CallOpts)
}

// MaxClick is a free data retrieval call binding the contract method 0x5410fcc6.
//
// Solidity: function maxClick() constant returns(uint256)
func (_LinkContract *LinkContractCallerSession) MaxClick() (*big.Int, error) {
	return _LinkContract.Contract.MaxClick(&_LinkContract.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members( uint256) constant returns(address)
func (_LinkContract *LinkContractCaller) Members(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "members", arg0)
	return *ret0, err
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members( uint256) constant returns(address)
func (_LinkContract *LinkContractSession) Members(arg0 *big.Int) (common.Address, error) {
	return _LinkContract.Contract.Members(&_LinkContract.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members( uint256) constant returns(address)
func (_LinkContract *LinkContractCallerSession) Members(arg0 *big.Int) (common.Address, error) {
	return _LinkContract.Contract.Members(&_LinkContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LinkContract *LinkContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LinkContract *LinkContractSession) Owner() (common.Address, error) {
	return _LinkContract.Contract.Owner(&_LinkContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LinkContract *LinkContractCallerSession) Owner() (common.Address, error) {
	return _LinkContract.Contract.Owner(&_LinkContract.CallOpts)
}

// Process is a free data retrieval call binding the contract method 0xc33fb877.
//
// Solidity: function process() constant returns(uint256)
func (_LinkContract *LinkContractCaller) Process(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "process")
	return *ret0, err
}

// Process is a free data retrieval call binding the contract method 0xc33fb877.
//
// Solidity: function process() constant returns(uint256)
func (_LinkContract *LinkContractSession) Process() (*big.Int, error) {
	return _LinkContract.Contract.Process(&_LinkContract.CallOpts)
}

// Process is a free data retrieval call binding the contract method 0xc33fb877.
//
// Solidity: function process() constant returns(uint256)
func (_LinkContract *LinkContractCallerSession) Process() (*big.Int, error) {
	return _LinkContract.Contract.Process(&_LinkContract.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_LinkContract *LinkContractCaller) Status(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_LinkContract *LinkContractSession) Status() (bool, error) {
	return _LinkContract.Contract.Status(&_LinkContract.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_LinkContract *LinkContractCallerSession) Status() (bool, error) {
	return _LinkContract.Contract.Status(&_LinkContract.CallOpts)
}

// SubLinks is a free data retrieval call binding the contract method 0x02b06ab1.
//
// Solidity: function subLinks( uint256) constant returns(id int32, addr_create address, addr_parent address, current_contract address, link string)
func (_LinkContract *LinkContractCaller) SubLinks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              int32
	AddrCreate      common.Address
	AddrParent      common.Address
	CurrentContract common.Address
	Link            string
}, error) {
	ret := new(struct {
		Id              int32
		AddrCreate      common.Address
		AddrParent      common.Address
		CurrentContract common.Address
		Link            string
	})
	out := ret
	err := _LinkContract.contract.Call(opts, out, "subLinks", arg0)
	return *ret, err
}

// SubLinks is a free data retrieval call binding the contract method 0x02b06ab1.
//
// Solidity: function subLinks( uint256) constant returns(id int32, addr_create address, addr_parent address, current_contract address, link string)
func (_LinkContract *LinkContractSession) SubLinks(arg0 *big.Int) (struct {
	Id              int32
	AddrCreate      common.Address
	AddrParent      common.Address
	CurrentContract common.Address
	Link            string
}, error) {
	return _LinkContract.Contract.SubLinks(&_LinkContract.CallOpts, arg0)
}

// SubLinks is a free data retrieval call binding the contract method 0x02b06ab1.
//
// Solidity: function subLinks( uint256) constant returns(id int32, addr_create address, addr_parent address, current_contract address, link string)
func (_LinkContract *LinkContractCallerSession) SubLinks(arg0 *big.Int) (struct {
	Id              int32
	AddrCreate      common.Address
	AddrParent      common.Address
	CurrentContract common.Address
	Link            string
}, error) {
	return _LinkContract.Contract.SubLinks(&_LinkContract.CallOpts, arg0)
}

// SuperOwner is a free data retrieval call binding the contract method 0x322d5a8e.
//
// Solidity: function super_owner() constant returns(address)
func (_LinkContract *LinkContractCaller) SuperOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "super_owner")
	return *ret0, err
}

// SuperOwner is a free data retrieval call binding the contract method 0x322d5a8e.
//
// Solidity: function super_owner() constant returns(address)
func (_LinkContract *LinkContractSession) SuperOwner() (common.Address, error) {
	return _LinkContract.Contract.SuperOwner(&_LinkContract.CallOpts)
}

// SuperOwner is a free data retrieval call binding the contract method 0x322d5a8e.
//
// Solidity: function super_owner() constant returns(address)
func (_LinkContract *LinkContractCallerSession) SuperOwner() (common.Address, error) {
	return _LinkContract.Contract.SuperOwner(&_LinkContract.CallOpts)
}

// TotalClicked is a free data retrieval call binding the contract method 0x69f3619c.
//
// Solidity: function totalClicked() constant returns(uint256)
func (_LinkContract *LinkContractCaller) TotalClicked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "totalClicked")
	return *ret0, err
}

// TotalClicked is a free data retrieval call binding the contract method 0x69f3619c.
//
// Solidity: function totalClicked() constant returns(uint256)
func (_LinkContract *LinkContractSession) TotalClicked() (*big.Int, error) {
	return _LinkContract.Contract.TotalClicked(&_LinkContract.CallOpts)
}

// TotalClicked is a free data retrieval call binding the contract method 0x69f3619c.
//
// Solidity: function totalClicked() constant returns(uint256)
func (_LinkContract *LinkContractCallerSession) TotalClicked() (*big.Int, error) {
	return _LinkContract.Contract.TotalClicked(&_LinkContract.CallOpts)
}

// TotalLinkGenerared is a free data retrieval call binding the contract method 0x71d10ee4.
//
// Solidity: function totalLinkGenerared() constant returns(uint16)
func (_LinkContract *LinkContractCaller) TotalLinkGenerared(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _LinkContract.contract.Call(opts, out, "totalLinkGenerared")
	return *ret0, err
}

// TotalLinkGenerared is a free data retrieval call binding the contract method 0x71d10ee4.
//
// Solidity: function totalLinkGenerared() constant returns(uint16)
func (_LinkContract *LinkContractSession) TotalLinkGenerared() (uint16, error) {
	return _LinkContract.Contract.TotalLinkGenerared(&_LinkContract.CallOpts)
}

// TotalLinkGenerared is a free data retrieval call binding the contract method 0x71d10ee4.
//
// Solidity: function totalLinkGenerared() constant returns(uint16)
func (_LinkContract *LinkContractCallerSession) TotalLinkGenerared() (uint16, error) {
	return _LinkContract.Contract.TotalLinkGenerared(&_LinkContract.CallOpts)
}

// ForwardLink is a paid mutator transaction binding the contract method 0xf6c655df.
//
// Solidity: function forwardLink(_addrRecive address, _linkId uint16) returns()
func (_LinkContract *LinkContractTransactor) ForwardLink(opts *bind.TransactOpts, _addrRecive common.Address, _linkId uint16) (*types.Transaction, error) {
	return _LinkContract.contract.Transact(opts, "forwardLink", _addrRecive, _linkId)
}

// ForwardLink is a paid mutator transaction binding the contract method 0xf6c655df.
//
// Solidity: function forwardLink(_addrRecive address, _linkId uint16) returns()
func (_LinkContract *LinkContractSession) ForwardLink(_addrRecive common.Address, _linkId uint16) (*types.Transaction, error) {
	return _LinkContract.Contract.ForwardLink(&_LinkContract.TransactOpts, _addrRecive, _linkId)
}

// ForwardLink is a paid mutator transaction binding the contract method 0xf6c655df.
//
// Solidity: function forwardLink(_addrRecive address, _linkId uint16) returns()
func (_LinkContract *LinkContractTransactorSession) ForwardLink(_addrRecive common.Address, _linkId uint16) (*types.Transaction, error) {
	return _LinkContract.Contract.ForwardLink(&_LinkContract.TransactOpts, _addrRecive, _linkId)
}

// GenerateLink is a paid mutator transaction binding the contract method 0xb7b18a88.
//
// Solidity: function generateLink(_parent address, _contract address) returns()
func (_LinkContract *LinkContractTransactor) GenerateLink(opts *bind.TransactOpts, _parent common.Address, _contract common.Address) (*types.Transaction, error) {
	return _LinkContract.contract.Transact(opts, "generateLink", _parent, _contract)
}

// GenerateLink is a paid mutator transaction binding the contract method 0xb7b18a88.
//
// Solidity: function generateLink(_parent address, _contract address) returns()
func (_LinkContract *LinkContractSession) GenerateLink(_parent common.Address, _contract common.Address) (*types.Transaction, error) {
	return _LinkContract.Contract.GenerateLink(&_LinkContract.TransactOpts, _parent, _contract)
}

// GenerateLink is a paid mutator transaction binding the contract method 0xb7b18a88.
//
// Solidity: function generateLink(_parent address, _contract address) returns()
func (_LinkContract *LinkContractTransactorSession) GenerateLink(_parent common.Address, _contract common.Address) (*types.Transaction, error) {
	return _LinkContract.Contract.GenerateLink(&_LinkContract.TransactOpts, _parent, _contract)
}

// OnNewClick is a paid mutator transaction binding the contract method 0x039243a5.
//
// Solidity: function onNewClick(_child address, _parent address) returns()
func (_LinkContract *LinkContractTransactor) OnNewClick(opts *bind.TransactOpts, _child common.Address, _parent common.Address) (*types.Transaction, error) {
	return _LinkContract.contract.Transact(opts, "onNewClick", _child, _parent)
}

// OnNewClick is a paid mutator transaction binding the contract method 0x039243a5.
//
// Solidity: function onNewClick(_child address, _parent address) returns()
func (_LinkContract *LinkContractSession) OnNewClick(_child common.Address, _parent common.Address) (*types.Transaction, error) {
	return _LinkContract.Contract.OnNewClick(&_LinkContract.TransactOpts, _child, _parent)
}

// OnNewClick is a paid mutator transaction binding the contract method 0x039243a5.
//
// Solidity: function onNewClick(_child address, _parent address) returns()
func (_LinkContract *LinkContractTransactorSession) OnNewClick(_child common.Address, _parent common.Address) (*types.Transaction, error) {
	return _LinkContract.Contract.OnNewClick(&_LinkContract.TransactOpts, _child, _parent)
}

// Payment is a paid mutator transaction binding the contract method 0x42f6487a.
//
// Solidity: function payment() returns()
func (_LinkContract *LinkContractTransactor) Payment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkContract.contract.Transact(opts, "payment")
}

// Payment is a paid mutator transaction binding the contract method 0x42f6487a.
//
// Solidity: function payment() returns()
func (_LinkContract *LinkContractSession) Payment() (*types.Transaction, error) {
	return _LinkContract.Contract.Payment(&_LinkContract.TransactOpts)
}

// Payment is a paid mutator transaction binding the contract method 0x42f6487a.
//
// Solidity: function payment() returns()
func (_LinkContract *LinkContractTransactorSession) Payment() (*types.Transaction, error) {
	return _LinkContract.Contract.Payment(&_LinkContract.TransactOpts)
}

// LinkContractForwardToIterator is returned from FilterForwardTo and is used to iterate over the raw logs and unpacked data for ForwardTo events raised by the LinkContract contract.
type LinkContractForwardToIterator struct {
	Event *LinkContractForwardTo // Event containing the contract specifics and raw log

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
func (it *LinkContractForwardToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkContractForwardTo)
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
		it.Event = new(LinkContractForwardTo)
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
func (it *LinkContractForwardToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkContractForwardToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkContractForwardTo represents a ForwardTo event raised by the LinkContract contract.
type LinkContractForwardTo struct {
	AddrRecive common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterForwardTo is a free log retrieval operation binding the contract event 0xb708fb2f4fb705069ac2d34f4057a64fa1fc9246fdc075f7cdf818a25e80415e.
//
// Solidity: e ForwardTo(addrRecive address)
func (_LinkContract *LinkContractFilterer) FilterForwardTo(opts *bind.FilterOpts) (*LinkContractForwardToIterator, error) {

	logs, sub, err := _LinkContract.contract.FilterLogs(opts, "ForwardTo")
	if err != nil {
		return nil, err
	}
	return &LinkContractForwardToIterator{contract: _LinkContract.contract, event: "ForwardTo", logs: logs, sub: sub}, nil
}

// WatchForwardTo is a free log subscription operation binding the contract event 0xb708fb2f4fb705069ac2d34f4057a64fa1fc9246fdc075f7cdf818a25e80415e.
//
// Solidity: e ForwardTo(addrRecive address)
func (_LinkContract *LinkContractFilterer) WatchForwardTo(opts *bind.WatchOpts, sink chan<- *LinkContractForwardTo) (event.Subscription, error) {

	logs, sub, err := _LinkContract.contract.WatchLogs(opts, "ForwardTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkContractForwardTo)
				if err := _LinkContract.contract.UnpackLog(event, "ForwardTo", log); err != nil {
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

// LinkContractGenerateLinkContractIterator is returned from FilterGenerateLinkContract and is used to iterate over the raw logs and unpacked data for GenerateLinkContract events raised by the LinkContract contract.
type LinkContractGenerateLinkContractIterator struct {
	Event *LinkContractGenerateLinkContract // Event containing the contract specifics and raw log

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
func (it *LinkContractGenerateLinkContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkContractGenerateLinkContract)
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
		it.Event = new(LinkContractGenerateLinkContract)
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
func (it *LinkContractGenerateLinkContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkContractGenerateLinkContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkContractGenerateLinkContract represents a GenerateLinkContract event raised by the LinkContract contract.
type LinkContractGenerateLinkContract struct {
	Sender   common.Address
	Parent   common.Address
	Contract common.Address
	Link     string
	Gateway  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGenerateLinkContract is a free log retrieval operation binding the contract event 0x0a13894e900be98ffdd0efbe75397f8621ce5762f64a7f2a25795a797b7445de.
//
// Solidity: e GenerateLinkContract(_sender address, _parent address, _contract address, _link string, _gateway string)
func (_LinkContract *LinkContractFilterer) FilterGenerateLinkContract(opts *bind.FilterOpts) (*LinkContractGenerateLinkContractIterator, error) {

	logs, sub, err := _LinkContract.contract.FilterLogs(opts, "GenerateLinkContract")
	if err != nil {
		return nil, err
	}
	return &LinkContractGenerateLinkContractIterator{contract: _LinkContract.contract, event: "GenerateLinkContract", logs: logs, sub: sub}, nil
}

// WatchGenerateLinkContract is a free log subscription operation binding the contract event 0x0a13894e900be98ffdd0efbe75397f8621ce5762f64a7f2a25795a797b7445de.
//
// Solidity: e GenerateLinkContract(_sender address, _parent address, _contract address, _link string, _gateway string)
func (_LinkContract *LinkContractFilterer) WatchGenerateLinkContract(opts *bind.WatchOpts, sink chan<- *LinkContractGenerateLinkContract) (event.Subscription, error) {

	logs, sub, err := _LinkContract.contract.WatchLogs(opts, "GenerateLinkContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkContractGenerateLinkContract)
				if err := _LinkContract.contract.UnpackLog(event, "GenerateLinkContract", log); err != nil {
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

// LinkContractPaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the LinkContract contract.
type LinkContractPaymentIterator struct {
	Event *LinkContractPayment // Event containing the contract specifics and raw log

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
func (it *LinkContractPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkContractPayment)
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
		it.Event = new(LinkContractPayment)
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
func (it *LinkContractPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkContractPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkContractPayment represents a Payment event raised by the LinkContract contract.
type LinkContractPayment struct {
	Addr  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0xd4f43975feb89f48dd30cabbb32011045be187d1e11c8ea9faa43efc35282519.
//
// Solidity: e Payment(addr address, value uint256)
func (_LinkContract *LinkContractFilterer) FilterPayment(opts *bind.FilterOpts) (*LinkContractPaymentIterator, error) {

	logs, sub, err := _LinkContract.contract.FilterLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return &LinkContractPaymentIterator{contract: _LinkContract.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0xd4f43975feb89f48dd30cabbb32011045be187d1e11c8ea9faa43efc35282519.
//
// Solidity: e Payment(addr address, value uint256)
func (_LinkContract *LinkContractFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *LinkContractPayment) (event.Subscription, error) {

	logs, sub, err := _LinkContract.contract.WatchLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkContractPayment)
				if err := _LinkContract.contract.UnpackLog(event, "Payment", log); err != nil {
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

// LinkContractUpContractLinkIterator is returned from FilterUpContractLink and is used to iterate over the raw logs and unpacked data for UpContractLink events raised by the LinkContract contract.
type LinkContractUpContractLinkIterator struct {
	Event *LinkContractUpContractLink // Event containing the contract specifics and raw log

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
func (it *LinkContractUpContractLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkContractUpContractLink)
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
		it.Event = new(LinkContractUpContractLink)
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
func (it *LinkContractUpContractLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkContractUpContractLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkContractUpContractLink represents a UpContractLink event raised by the LinkContract contract.
type LinkContractUpContractLink struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpContractLink is a free log retrieval operation binding the contract event 0x784a9cc22598ebcd47d2d416a4fb20c43dfbbce49ad2368a7e2fd4a8a78dc2ba.
//
// Solidity: e UpContractLink()
func (_LinkContract *LinkContractFilterer) FilterUpContractLink(opts *bind.FilterOpts) (*LinkContractUpContractLinkIterator, error) {

	logs, sub, err := _LinkContract.contract.FilterLogs(opts, "UpContractLink")
	if err != nil {
		return nil, err
	}
	return &LinkContractUpContractLinkIterator{contract: _LinkContract.contract, event: "UpContractLink", logs: logs, sub: sub}, nil
}

// WatchUpContractLink is a free log subscription operation binding the contract event 0x784a9cc22598ebcd47d2d416a4fb20c43dfbbce49ad2368a7e2fd4a8a78dc2ba.
//
// Solidity: e UpContractLink()
func (_LinkContract *LinkContractFilterer) WatchUpContractLink(opts *bind.WatchOpts, sink chan<- *LinkContractUpContractLink) (event.Subscription, error) {

	logs, sub, err := _LinkContract.contract.WatchLogs(opts, "UpContractLink")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkContractUpContractLink)
				if err := _LinkContract.contract.UnpackLog(event, "UpContractLink", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582083d3fe39622063d79e1206629eb428041db92d81c2c25536ddd53065c58863620029`

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
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

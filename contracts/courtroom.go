// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CaseContractABI is the input ABI used to generate the binding from.
const CaseContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"},{\"name\":\"_evident\",\"type\":\"bytes32\"}],\"name\":\"submitEvidence\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"claimId\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"OpenClaims\",\"outputs\":[{\"name\":\"claimId\",\"type\":\"bytes32\"},{\"name\":\"plaintiff\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint256\"},{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"setClaimValid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"resolveClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"name\":\"plaintiff\",\"type\":\"address\"},{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plaintiff\",\"type\":\"address\"},{\"name\":\"_evidence\",\"type\":\"bytes32\"}],\"name\":\"newClaim\",\"outputs\":[{\"name\":\"claimId\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// CaseContractBin is the compiled bytecode used for deploying new contracts.
const CaseContractBin = `0x6060604052341561000f57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b5b5b6106628061003e6000396000f300606060405236156100965763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317dfa44c811461009b5780635de28ae0146100c65780637f46848b146100ee5780638da5cb5b14610138578063a388246714610167578063a6f9dae11461017f578063c1203cb8146101a0578063c9100bcb146101b8578063e391c4c0146101f2575b600080fd5b34156100a657600080fd5b6100b4600435602435610226565b60405190815260200160405180910390f35b34156100d157600080fd5b6100b4600435610296565b60405190815260200160405180910390f35b34156100f957600080fd5b6101046004356102ae565b604051938452600160a060020a03909216602084015260408084019190915290151560608301526080909101905180910390f35b341561014357600080fd5b61014b6102e3565b604051600160a060020a03909116815260200160405180910390f35b341561017257600080fd5b61017d6004356102f2565b005b341561018a57600080fd5b61017d600160a060020a0360043516610315565b005b34156101ab57600080fd5b61017d60043561035d565b005b34156101c357600080fd5b6101ce6004356103f5565b604051600160a060020a039092168252151560208201526040908101905180910390f35b34156101fd57600080fd5b6100b4600160a060020a0360043516602435610425565b60405190815260200160405180910390f35b6000828152600160205260408120600201546004901061024557600080fd5b60008381526001602081905260409091206002018054909181016102698382610570565b916000526020600020900160005b5083905550506000828152600160205260409020600301545b92915050565b6000818152600160205260409020600301545b919050565b60016020819052600091825260409091208054918101546003820154600490920154600160a060020a03909116919060ff1684565b600054600160a060020a031681565b6000818152600160208190526040909120600401805460ff191690911790555b50565b60005433600160a060020a0390811691161461033057600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60008181526001602052604090205415806103875750600081815260016020526040902060030154155b1561039157600080fd5b60008181526001602081905260408220828155908101805473ffffffffffffffffffffffffffffffffffffffff191690556103cf9060020182610570565b5060008181526001602052604081206003810191909155600401805460ff191690555b50565b600081815260016020819052604090912090810154600490910154600160a060020a039091169060ff165b915091565b600080838342604051600160a060020a03939093166c010000000000000000000000000283526014830191909152603482015260540160405190819003902081556001808201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038716179055600282018054909181016104a58382610570565b916000526020600020900160005b5084905550600160038083018290558254600090815260209290925260409091200154156104e057600080fd5b8054600081815260016020819052604090912091825580830154908201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790556002808301805484939261053e92908401916105c4565b50600382810154908201556004918201549101805460ff191660ff9092161515919091179055805491505b5092915050565b81548183558181151161059457600083815260209020610594918101908301610615565b5b505050565b81548183558181151161059457600083815260209020610594918101908301610615565b5b505050565b8280548282559060005260206000209081019282156106045760005260206000209182015b828111156106045782548255916001019190600101906105e9565b5b50610611929150610615565b5090565b61063391905b80821115610611576000815560010161061b565b5090565b905600a165627a7a7230582074b4d83c3e51643085389a7ad83cd5b4c65f60c28aacbf414799961f1f6d73b30029`

// DeployCaseContract deploys a new Ethereum contract, binding an instance of CaseContract to it.
func DeployCaseContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CaseContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CaseContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CaseContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CaseContract{CaseContractCaller: CaseContractCaller{contract: contract}, CaseContractTransactor: CaseContractTransactor{contract: contract}}, nil
}

// CaseContract is an auto generated Go binding around an Ethereum contract.
type CaseContract struct {
	CaseContractCaller     // Read-only binding to the contract
	CaseContractTransactor // Write-only binding to the contract
}

// CaseContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CaseContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaseContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CaseContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaseContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CaseContractSession struct {
	Contract     *CaseContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CaseContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CaseContractCallerSession struct {
	Contract *CaseContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CaseContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CaseContractTransactorSession struct {
	Contract     *CaseContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CaseContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CaseContractRaw struct {
	Contract *CaseContract // Generic contract binding to access the raw methods on
}

// CaseContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CaseContractCallerRaw struct {
	Contract *CaseContractCaller // Generic read-only contract binding to access the raw methods on
}

// CaseContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CaseContractTransactorRaw struct {
	Contract *CaseContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaseContract creates a new instance of CaseContract, bound to a specific deployed contract.
func NewCaseContract(address common.Address, backend bind.ContractBackend) (*CaseContract, error) {
	contract, err := bindCaseContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CaseContract{CaseContractCaller: CaseContractCaller{contract: contract}, CaseContractTransactor: CaseContractTransactor{contract: contract}}, nil
}

// NewCaseContractCaller creates a new read-only instance of CaseContract, bound to a specific deployed contract.
func NewCaseContractCaller(address common.Address, caller bind.ContractCaller) (*CaseContractCaller, error) {
	contract, err := bindCaseContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CaseContractCaller{contract: contract}, nil
}

// NewCaseContractTransactor creates a new write-only instance of CaseContract, bound to a specific deployed contract.
func NewCaseContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CaseContractTransactor, error) {
	contract, err := bindCaseContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CaseContractTransactor{contract: contract}, nil
}

// bindCaseContract binds a generic wrapper to an already deployed contract.
func bindCaseContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CaseContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaseContract *CaseContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CaseContract.Contract.CaseContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaseContract *CaseContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaseContract.Contract.CaseContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaseContract *CaseContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaseContract.Contract.CaseContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaseContract *CaseContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CaseContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaseContract *CaseContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaseContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaseContract *CaseContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaseContract.Contract.contract.Transact(opts, method, params...)
}

// OpenClaims is a free data retrieval call binding the contract method 0x7f46848b.
//
// Solidity: function OpenClaims( bytes32) constant returns(claimId bytes32, plaintiff address, status uint256, valid bool)
func (_CaseContract *CaseContractCaller) OpenClaims(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ClaimId   [32]byte
	Plaintiff common.Address
	Status    *big.Int
	Valid     bool
}, error) {
	ret := new(struct {
		ClaimId   [32]byte
		Plaintiff common.Address
		Status    *big.Int
		Valid     bool
	})
	out := ret
	err := _CaseContract.contract.Call(opts, out, "OpenClaims", arg0)
	return *ret, err
}

// OpenClaims is a free data retrieval call binding the contract method 0x7f46848b.
//
// Solidity: function OpenClaims( bytes32) constant returns(claimId bytes32, plaintiff address, status uint256, valid bool)
func (_CaseContract *CaseContractSession) OpenClaims(arg0 [32]byte) (struct {
	ClaimId   [32]byte
	Plaintiff common.Address
	Status    *big.Int
	Valid     bool
}, error) {
	return _CaseContract.Contract.OpenClaims(&_CaseContract.CallOpts, arg0)
}

// OpenClaims is a free data retrieval call binding the contract method 0x7f46848b.
//
// Solidity: function OpenClaims( bytes32) constant returns(claimId bytes32, plaintiff address, status uint256, valid bool)
func (_CaseContract *CaseContractCallerSession) OpenClaims(arg0 [32]byte) (struct {
	ClaimId   [32]byte
	Plaintiff common.Address
	Status    *big.Int
	Valid     bool
}, error) {
	return _CaseContract.Contract.OpenClaims(&_CaseContract.CallOpts, arg0)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(claimId bytes32) constant returns(status uint256)
func (_CaseContract *CaseContractCaller) GetStatus(opts *bind.CallOpts, claimId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CaseContract.contract.Call(opts, out, "getStatus", claimId)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(claimId bytes32) constant returns(status uint256)
func (_CaseContract *CaseContractSession) GetStatus(claimId [32]byte) (*big.Int, error) {
	return _CaseContract.Contract.GetStatus(&_CaseContract.CallOpts, claimId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(claimId bytes32) constant returns(status uint256)
func (_CaseContract *CaseContractCallerSession) GetStatus(claimId [32]byte) (*big.Int, error) {
	return _CaseContract.Contract.GetStatus(&_CaseContract.CallOpts, claimId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CaseContract *CaseContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CaseContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CaseContract *CaseContractSession) Owner() (common.Address, error) {
	return _CaseContract.Contract.Owner(&_CaseContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CaseContract *CaseContractCallerSession) Owner() (common.Address, error) {
	return _CaseContract.Contract.Owner(&_CaseContract.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_CaseContract *CaseContractTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_CaseContract *CaseContractSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CaseContract.Contract.ChangeOwner(&_CaseContract.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_CaseContract *CaseContractTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CaseContract.Contract.ChangeOwner(&_CaseContract.TransactOpts, _newOwner)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_claimId bytes32) returns(plaintiff address, valid bool)
func (_CaseContract *CaseContractTransactor) GetClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "getClaim", _claimId)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_claimId bytes32) returns(plaintiff address, valid bool)
func (_CaseContract *CaseContractSession) GetClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.GetClaim(&_CaseContract.TransactOpts, _claimId)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_claimId bytes32) returns(plaintiff address, valid bool)
func (_CaseContract *CaseContractTransactorSession) GetClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.GetClaim(&_CaseContract.TransactOpts, _claimId)
}

// NewClaim is a paid mutator transaction binding the contract method 0xe391c4c0.
//
// Solidity: function newClaim(_plaintiff address, _evidence bytes32) returns(claimId bytes32)
func (_CaseContract *CaseContractTransactor) NewClaim(opts *bind.TransactOpts, _plaintiff common.Address, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "newClaim", _plaintiff, _evidence)
}

// NewClaim is a paid mutator transaction binding the contract method 0xe391c4c0.
//
// Solidity: function newClaim(_plaintiff address, _evidence bytes32) returns(claimId bytes32)
func (_CaseContract *CaseContractSession) NewClaim(_plaintiff common.Address, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.NewClaim(&_CaseContract.TransactOpts, _plaintiff, _evidence)
}

// NewClaim is a paid mutator transaction binding the contract method 0xe391c4c0.
//
// Solidity: function newClaim(_plaintiff address, _evidence bytes32) returns(claimId bytes32)
func (_CaseContract *CaseContractTransactorSession) NewClaim(_plaintiff common.Address, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.NewClaim(&_CaseContract.TransactOpts, _plaintiff, _evidence)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0xc1203cb8.
//
// Solidity: function resolveClaim(_claimId bytes32) returns()
func (_CaseContract *CaseContractTransactor) ResolveClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "resolveClaim", _claimId)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0xc1203cb8.
//
// Solidity: function resolveClaim(_claimId bytes32) returns()
func (_CaseContract *CaseContractSession) ResolveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.ResolveClaim(&_CaseContract.TransactOpts, _claimId)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0xc1203cb8.
//
// Solidity: function resolveClaim(_claimId bytes32) returns()
func (_CaseContract *CaseContractTransactorSession) ResolveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.ResolveClaim(&_CaseContract.TransactOpts, _claimId)
}

// SetClaimValid is a paid mutator transaction binding the contract method 0xa3882467.
//
// Solidity: function setClaimValid(_claimId bytes32) returns()
func (_CaseContract *CaseContractTransactor) SetClaimValid(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "setClaimValid", _claimId)
}

// SetClaimValid is a paid mutator transaction binding the contract method 0xa3882467.
//
// Solidity: function setClaimValid(_claimId bytes32) returns()
func (_CaseContract *CaseContractSession) SetClaimValid(_claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SetClaimValid(&_CaseContract.TransactOpts, _claimId)
}

// SetClaimValid is a paid mutator transaction binding the contract method 0xa3882467.
//
// Solidity: function setClaimValid(_claimId bytes32) returns()
func (_CaseContract *CaseContractTransactorSession) SetClaimValid(_claimId [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SetClaimValid(&_CaseContract.TransactOpts, _claimId)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x17dfa44c.
//
// Solidity: function submitEvidence(_claimId bytes32, _evident bytes32) returns(status uint256)
func (_CaseContract *CaseContractTransactor) SubmitEvidence(opts *bind.TransactOpts, _claimId [32]byte, _evident [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "submitEvidence", _claimId, _evident)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x17dfa44c.
//
// Solidity: function submitEvidence(_claimId bytes32, _evident bytes32) returns(status uint256)
func (_CaseContract *CaseContractSession) SubmitEvidence(_claimId [32]byte, _evident [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SubmitEvidence(&_CaseContract.TransactOpts, _claimId, _evident)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x17dfa44c.
//
// Solidity: function submitEvidence(_claimId bytes32, _evident bytes32) returns(status uint256)
func (_CaseContract *CaseContractTransactorSession) SubmitEvidence(_claimId [32]byte, _evident [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SubmitEvidence(&_CaseContract.TransactOpts, _claimId, _evident)
}

// ENSAbstractABI is the input ABI used to generate the binding from.
const ENSAbstractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"

// ENSAbstractBin is the compiled bytecode used for deploying new contracts.
const ENSAbstractBin = `0x`

// DeployENSAbstract deploys a new Ethereum contract, binding an instance of ENSAbstract to it.
func DeployENSAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ENSAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ENSAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENSAbstract{ENSAbstractCaller: ENSAbstractCaller{contract: contract}, ENSAbstractTransactor: ENSAbstractTransactor{contract: contract}}, nil
}

// ENSAbstract is an auto generated Go binding around an Ethereum contract.
type ENSAbstract struct {
	ENSAbstractCaller     // Read-only binding to the contract
	ENSAbstractTransactor // Write-only binding to the contract
}

// ENSAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSAbstractSession struct {
	Contract     *ENSAbstract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSAbstractCallerSession struct {
	Contract *ENSAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ENSAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSAbstractTransactorSession struct {
	Contract     *ENSAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ENSAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSAbstractRaw struct {
	Contract *ENSAbstract // Generic contract binding to access the raw methods on
}

// ENSAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSAbstractCallerRaw struct {
	Contract *ENSAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// ENSAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSAbstractTransactorRaw struct {
	Contract *ENSAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENSAbstract creates a new instance of ENSAbstract, bound to a specific deployed contract.
func NewENSAbstract(address common.Address, backend bind.ContractBackend) (*ENSAbstract, error) {
	contract, err := bindENSAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENSAbstract{ENSAbstractCaller: ENSAbstractCaller{contract: contract}, ENSAbstractTransactor: ENSAbstractTransactor{contract: contract}}, nil
}

// NewENSAbstractCaller creates a new read-only instance of ENSAbstract, bound to a specific deployed contract.
func NewENSAbstractCaller(address common.Address, caller bind.ContractCaller) (*ENSAbstractCaller, error) {
	contract, err := bindENSAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ENSAbstractCaller{contract: contract}, nil
}

// NewENSAbstractTransactor creates a new write-only instance of ENSAbstract, bound to a specific deployed contract.
func NewENSAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSAbstractTransactor, error) {
	contract, err := bindENSAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ENSAbstractTransactor{contract: contract}, nil
}

// bindENSAbstract binds a generic wrapper to an already deployed contract.
func bindENSAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSAbstract *ENSAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSAbstract.Contract.ENSAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSAbstract *ENSAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSAbstract.Contract.ENSAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSAbstract *ENSAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSAbstract.Contract.ENSAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSAbstract *ENSAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSAbstract *ENSAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSAbstract *ENSAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSAbstract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSAbstract.contract.Call(opts, out, "owner", node)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Owner(&_ENSAbstract.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Owner(&_ENSAbstract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSAbstract.contract.Call(opts, out, "resolver", node)
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Resolver(&_ENSAbstract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Resolver(&_ENSAbstract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_ENSAbstract *ENSAbstractCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ENSAbstract.contract.Call(opts, out, "ttl", node)
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_ENSAbstract *ENSAbstractSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSAbstract.Contract.Ttl(&_ENSAbstract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_ENSAbstract *ENSAbstractCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSAbstract.Contract.Ttl(&_ENSAbstract.CallOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetOwner(&_ENSAbstract.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetOwner(&_ENSAbstract.TransactOpts, node, owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENSAbstract *ENSAbstractSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetResolver(&_ENSAbstract.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetResolver(&_ENSAbstract.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetSubnodeOwner(&_ENSAbstract.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetSubnodeOwner(&_ENSAbstract.TransactOpts, node, label, owner)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_ENSAbstract *ENSAbstractSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetTTL(&_ENSAbstract.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetTTL(&_ENSAbstract.TransactOpts, node, ttl)
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b61015c8061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610048578063a6f9dae114610084575b600080fd5b341561005357600080fd5b61005b6100b2565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008f57600080fd5b6100b073ffffffffffffffffffffffffffffffffffffffff600435166100ce565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a72305820efbec96f65b3d63408929c9c69a78f3654b8c7e1a37346a5e6fff6c7a1cb19520029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// ResolverAbstractABI is the input ABI used to generate the binding from.
const ResolverAbstractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"content\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"}]"

// ResolverAbstractBin is the compiled bytecode used for deploying new contracts.
const ResolverAbstractBin = `0x`

// DeployResolverAbstract deploys a new Ethereum contract, binding an instance of ResolverAbstract to it.
func DeployResolverAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ResolverAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ResolverAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ResolverAbstract{ResolverAbstractCaller: ResolverAbstractCaller{contract: contract}, ResolverAbstractTransactor: ResolverAbstractTransactor{contract: contract}}, nil
}

// ResolverAbstract is an auto generated Go binding around an Ethereum contract.
type ResolverAbstract struct {
	ResolverAbstractCaller     // Read-only binding to the contract
	ResolverAbstractTransactor // Write-only binding to the contract
}

// ResolverAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverAbstractSession struct {
	Contract     *ResolverAbstract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverAbstractCallerSession struct {
	Contract *ResolverAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ResolverAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverAbstractTransactorSession struct {
	Contract     *ResolverAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ResolverAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverAbstractRaw struct {
	Contract *ResolverAbstract // Generic contract binding to access the raw methods on
}

// ResolverAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverAbstractCallerRaw struct {
	Contract *ResolverAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverAbstractTransactorRaw struct {
	Contract *ResolverAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverAbstract creates a new instance of ResolverAbstract, bound to a specific deployed contract.
func NewResolverAbstract(address common.Address, backend bind.ContractBackend) (*ResolverAbstract, error) {
	contract, err := bindResolverAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverAbstract{ResolverAbstractCaller: ResolverAbstractCaller{contract: contract}, ResolverAbstractTransactor: ResolverAbstractTransactor{contract: contract}}, nil
}

// NewResolverAbstractCaller creates a new read-only instance of ResolverAbstract, bound to a specific deployed contract.
func NewResolverAbstractCaller(address common.Address, caller bind.ContractCaller) (*ResolverAbstractCaller, error) {
	contract, err := bindResolverAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverAbstractCaller{contract: contract}, nil
}

// NewResolverAbstractTransactor creates a new write-only instance of ResolverAbstract, bound to a specific deployed contract.
func NewResolverAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverAbstractTransactor, error) {
	contract, err := bindResolverAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ResolverAbstractTransactor{contract: contract}, nil
}

// bindResolverAbstract binds a generic wrapper to an already deployed contract.
func bindResolverAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverAbstract *ResolverAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverAbstract.Contract.ResolverAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverAbstract *ResolverAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.ResolverAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverAbstract *ResolverAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.ResolverAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverAbstract *ResolverAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverAbstract *ResolverAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverAbstract *ResolverAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.contract.Transact(opts, method, params...)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(content bytes32)
func (_ResolverAbstract *ResolverAbstractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ResolverAbstract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(content bytes32)
func (_ResolverAbstract *ResolverAbstractSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverAbstract.Contract.Content(&_ResolverAbstract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(content bytes32)
func (_ResolverAbstract *ResolverAbstractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverAbstract.Contract.Content(&_ResolverAbstract.CallOpts, node)
}

// SampleTokenABI is the input ABI used to generate the binding from.
const SampleTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"ref\",\"type\":\"bytes32\"}],\"name\":\"createTokens\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"ref\",\"type\":\"bytes32\"}],\"name\":\"TokenMined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SampleTokenBin is the compiled bytecode used for deploying new contracts.
const SampleTokenBin = `0x6060604052341561000f57600080fd5b60405160208061069d833981016040528080519150505b5b60038054600160a060020a03191633600160a060020a03161790555b600160a060020a03331660009081526001602052604081208290558190555b505b61062a806100736000396000f300606060405236156100965763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009b57806318160ddd146100d157806323b872dd146100f657806370a08231146101325780638da5cb5b14610163578063a3aff92314610192578063a6f9dae1146101b9578063a9059cbb146101da578063dd62ed3e14610210575b600080fd5b34156100a657600080fd5b6100bd600160a060020a0360043516602435610247565b604051901515815260200160405180910390f35b34156100dc57600080fd5b6100e46102b4565b60405190815260200160405180910390f35b341561010157600080fd5b6100bd600160a060020a03600435811690602435166044356102ba565b604051901515815260200160405180910390f35b341561013d57600080fd5b6100e4600160a060020a03600435166102fd565b60405190815260200160405180910390f35b341561016e57600080fd5b61017661031c565b604051600160a060020a03909116815260200160405180910390f35b341561019d57600080fd5b6101b7600160a060020a036004351660243560443561032b565b005b34156101c457600080fd5b6101b7600160a060020a03600435166103e8565b005b34156101e557600080fd5b6100bd600160a060020a0360043516602435610430565b604051901515815260200160405180910390f35b341561021b57600080fd5b6100e4600160a060020a03600435811690602435166104d7565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0380841660009081526002602090815260408083203390941683529290529081208054830190556102f3848484610504565b90505b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600354600160a060020a031681565b60035433600160a060020a0390811691161461034657600080fd5b600160a060020a038316600081815260016020526040808220805486019055815485019091558291907fdfb81fb379557413b0a951b4d7bf7a9df393801d8c539d5e201d6a8daeb913b99085905190815260200160405180910390a382600160a060020a031660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505050565b60035433600160a060020a0390811691161461040357600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600160a060020a0333166000908152600160205260408120548290108015906104595750600082115b1561009657600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102ae565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b92915050565b600160a060020a0383166000908152600160205260408120548290108015906105545750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b80156105605750600082115b1561009657600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102f6565b600080fd5b5b93925050505600a165627a7a72305820f441ab9704383ceac725b7d8a49b8f3e7f15d02d3a606fc590f113a79c2b3c3f0029`

// DeploySampleToken deploys a new Ethereum contract, binding an instance of SampleToken to it.
func DeploySampleToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *SampleToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleTokenBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleToken{SampleTokenCaller: SampleTokenCaller{contract: contract}, SampleTokenTransactor: SampleTokenTransactor{contract: contract}}, nil
}

// SampleToken is an auto generated Go binding around an Ethereum contract.
type SampleToken struct {
	SampleTokenCaller     // Read-only binding to the contract
	SampleTokenTransactor // Write-only binding to the contract
}

// SampleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleTokenSession struct {
	Contract     *SampleToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleTokenCallerSession struct {
	Contract *SampleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SampleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleTokenTransactorSession struct {
	Contract     *SampleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SampleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleTokenRaw struct {
	Contract *SampleToken // Generic contract binding to access the raw methods on
}

// SampleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleTokenCallerRaw struct {
	Contract *SampleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SampleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleTokenTransactorRaw struct {
	Contract *SampleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleToken creates a new instance of SampleToken, bound to a specific deployed contract.
func NewSampleToken(address common.Address, backend bind.ContractBackend) (*SampleToken, error) {
	contract, err := bindSampleToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleToken{SampleTokenCaller: SampleTokenCaller{contract: contract}, SampleTokenTransactor: SampleTokenTransactor{contract: contract}}, nil
}

// NewSampleTokenCaller creates a new read-only instance of SampleToken, bound to a specific deployed contract.
func NewSampleTokenCaller(address common.Address, caller bind.ContractCaller) (*SampleTokenCaller, error) {
	contract, err := bindSampleToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SampleTokenCaller{contract: contract}, nil
}

// NewSampleTokenTransactor creates a new write-only instance of SampleToken, bound to a specific deployed contract.
func NewSampleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleTokenTransactor, error) {
	contract, err := bindSampleToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SampleTokenTransactor{contract: contract}, nil
}

// bindSampleToken binds a generic wrapper to an already deployed contract.
func bindSampleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleToken *SampleTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleToken.Contract.SampleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleToken *SampleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleToken.Contract.SampleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleToken *SampleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleToken.Contract.SampleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleToken *SampleTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleToken *SampleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleToken *SampleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SampleToken *SampleTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SampleToken *SampleTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SampleToken.Contract.Allowance(&_SampleToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SampleToken *SampleTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SampleToken.Contract.Allowance(&_SampleToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleToken *SampleTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleToken *SampleTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SampleToken.Contract.BalanceOf(&_SampleToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleToken *SampleTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SampleToken.Contract.BalanceOf(&_SampleToken.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleToken *SampleTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleToken *SampleTokenSession) Owner() (common.Address, error) {
	return _SampleToken.Contract.Owner(&_SampleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleToken *SampleTokenCallerSession) Owner() (common.Address, error) {
	return _SampleToken.Contract.Owner(&_SampleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleToken *SampleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleToken *SampleTokenSession) TotalSupply() (*big.Int, error) {
	return _SampleToken.Contract.TotalSupply(&_SampleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleToken *SampleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SampleToken.Contract.TotalSupply(&_SampleToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Approve(&_SampleToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Approve(&_SampleToken.TransactOpts, _spender, _value)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SampleToken *SampleTokenTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SampleToken *SampleTokenSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SampleToken.Contract.ChangeOwner(&_SampleToken.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SampleToken *SampleTokenTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SampleToken.Contract.ChangeOwner(&_SampleToken.TransactOpts, _newOwner)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xa3aff923.
//
// Solidity: function createTokens(beneficiary address, amount uint256, ref bytes32) returns()
func (_SampleToken *SampleTokenTransactor) CreateTokens(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int, ref [32]byte) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "createTokens", beneficiary, amount, ref)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xa3aff923.
//
// Solidity: function createTokens(beneficiary address, amount uint256, ref bytes32) returns()
func (_SampleToken *SampleTokenSession) CreateTokens(beneficiary common.Address, amount *big.Int, ref [32]byte) (*types.Transaction, error) {
	return _SampleToken.Contract.CreateTokens(&_SampleToken.TransactOpts, beneficiary, amount, ref)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xa3aff923.
//
// Solidity: function createTokens(beneficiary address, amount uint256, ref bytes32) returns()
func (_SampleToken *SampleTokenTransactorSession) CreateTokens(beneficiary common.Address, amount *big.Int, ref [32]byte) (*types.Transaction, error) {
	return _SampleToken.Contract.CreateTokens(&_SampleToken.TransactOpts, beneficiary, amount, ref)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Transfer(&_SampleToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Transfer(&_SampleToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.TransferFrom(&_SampleToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.TransferFrom(&_SampleToken.TransactOpts, _from, _to, _value)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b5b61043b8061001f6000396000f300606060405236156100755763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007a57806318160ddd146100b057806323b872dd146100d557806370a0823114610111578063a9059cbb14610142578063dd62ed3e14610178575b600080fd5b341561008557600080fd5b61009c600160a060020a03600435166024356101af565b604051901515815260200160405180910390f35b34156100bb57600080fd5b6100c361021c565b60405190815260200160405180910390f35b34156100e057600080fd5b61009c600160a060020a0360043581169060243516604435610222565b604051901515815260200160405180910390f35b341561011c57600080fd5b6100c3600160a060020a036004351661031c565b60405190815260200160405180910390f35b341561014d57600080fd5b61009c600160a060020a036004351660243561033b565b604051901515815260200160405180910390f35b341561018357600080fd5b6100c3600160a060020a03600435811690602435166103e2565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0383166000908152600160205260408120548290108015906102725750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b801561027e5750600082115b1561007557600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610314565b600080fd5b5b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600160a060020a0333166000908152600160205260408120548290108015906103645750600082115b1561007557600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610216565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b929150505600a165627a7a723058204c52979f15737467b39822bd03fcfac79cd8b6389b15e50bf0190ab4a4a3a6b80029`

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
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
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
	contract, err := bindStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
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
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
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
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
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
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// SwearGameABI is the input ABI used to generate the binding from.
const SwearGameABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clientsClaimsIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"leaveGame\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registeredPlayersCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredPlayers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ensAddr\",\"type\":\"address\"}],\"name\":\"setENSAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"claimId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardCompensation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sampleToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evidence\",\"type\":\"bytes32\"}],\"name\":\"openNewClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"caseContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountStaked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_caseContract\",\"type\":\"address\"},{\"name\":\"_sampleToken\",\"type\":\"address\"},{\"name\":\"_rewardCompensation\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"decide\",\"type\":\"string\"}],\"name\":\"Decision\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountStaked\",\"type\":\"uint256\"}],\"name\":\"DepositStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"rewardCompensation\",\"type\":\"uint256\"}],\"name\":\"Compensate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"NewPlayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"PlayerLeftGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caseId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewClaimOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewEvidenceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"rewardCompensation\",\"type\":\"uint256\"}],\"name\":\"ClaimResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amountStaked\",\"type\":\"uint256\"}],\"name\":\"AdditionalDepositRequired\",\"type\":\"event\"}]"

// SwearGameBin is the compiled bytecode used for deploying new contracts.
const SwearGameBin = `0x6060604052600a8054600160a060020a031916738163bc885c2b14478b75f178ca76f31581dc967f179055341561003557600080fd5b604051606080610f6d8339810160405280805191906020018051919060200180519150505b5b60008054600160a060020a03191633600160a060020a03161790555b60068054600160a060020a03808616600160a060020a0319928316179092556007805492851692909116919091179055600281905560006003555b5050505b610ea8806100c56000396000f300606060405236156100ca5763ffffffff60e060020a6000350416633258e6ca81146100cf5780633e95d955146101035780634420e486146101245780634a3ebbed146101575780634bafce221461017c5780634cad1ce9146101af5780635c25fd13146101e257806366be2236146102075780637222e9e31461022c5780638da5cb5b1461025b578063a19e95da1461028a578063a6b694f2146102b4578063a6f9dae1146102e3578063b6b55f2514610304578063d556e8e614610323578063fab1e74714610352575b600080fd5b34156100da57600080fd5b6100f1600160a060020a0360043516602435610377565b60405190815260200160405180910390f35b341561010e57600080fd5b610122600160a060020a03600435166103a9565b005b341561012f57600080fd5b610143600160a060020a0360043516610457565b604051901515815260200160405180910390f35b341561016257600080fd5b6100f1610584565b60405190815260200160405180910390f35b341561018757600080fd5b610143600160a060020a036004351661058a565b604051901515815260200160405180910390f35b34156101ba57600080fd5b610143600160a060020a036004351661059f565b604051901515815260200160405180910390f35b34156101ed57600080fd5b6100f16105cf565b60405190815260200160405180910390f35b341561021257600080fd5b6100f16105d5565b60405190815260200160405180910390f35b341561023757600080fd5b61023f6105db565b604051600160a060020a03909116815260200160405180910390f35b341561026657600080fd5b61023f6105ea565b604051600160a060020a03909116815260200160405180910390f35b341561029557600080fd5b6101436004356105f9565b604051901515815260200160405180910390f35b34156102bf57600080fd5b61023f6106ef565b604051600160a060020a03909116815260200160405180910390f35b34156102ee57600080fd5b610122600160a060020a03600435166106fe565b005b610143600435610746565b604051901515815260200160405180910390f35b341561032e57600080fd5b61023f6108d9565b604051600160a060020a03909116815260200160405180910390f35b341561035d57600080fd5b6100f16108e8565b60405190815260200160405180910390f35b60056020528160005260406000208181548110151561039257fe5b906000526020600020900160005b91509150505481565b60005433600160a060020a039081169116146103c457600080fd5b600160a060020a03811660009081526004602052604090205460ff1615156103eb57600080fd5b7f3def0aea61e344e47b98a0523a5825c0919e88f7b33d0ac7f889a7654880f9a581604051600160a060020a03909116815260200160405180910390a1600160a060020a0381166000908152600460205260409020805460ff19169055600380546000190190555b5b50565b6000805433600160a060020a0390811691161461047357600080fd5b600160a060020a03821660009081526004602052604090205460ff161561049957600080fd5b60035415156104b85760025460015410156104b357600080fd5b61050b565b6002546003546001548115156104ca57fe5b04101561050b577fa8ff0fbf8dc82e5914523ed3a828920f55ab93494c59933e8c45d79f8bba70d760015460405190815260200160405180910390a1600080fd5b5b600160a060020a03821660009081526004602052604090819020805460ff191660019081179091556003805490910190557f52e92d4898337244a39bd42674ac561eadfd3959e947deec1c0ab82dd58b5a7590839051600160a060020a03909116815260200160405180910390a15060015b5b919050565b60035481565b60046020526000908152604090205460ff1681565b600a805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831617905560015b919050565b60085481565b60025481565b600754600160a060020a031681565b600054600160a060020a031681565b600160a060020a03331660009081526004602052604081205460ff16151561062057600080fd5b600654600160a060020a031663e391c4c0338460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561067f57600080fd5b6102c65a03f1151561069057600080fd5b505050604051805160085550600160a060020a03331660009081526005602052604090208054600181016106c48382610e31565b916000526020600020900160005b50600854908190556106e491506108ee565b50600190505b919050565b600a54600160a060020a031681565b60005433600160a060020a0390811691161461071957600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b6000805433600160a060020a0390811691161461076257600080fd5b600754600080548492600160a060020a03908116926370a0823192909116906040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156107c657600080fd5b6102c65a03f115156107d757600080fd5b50505060405180519050101515156107ee57600080fd5b60075460008054600160a060020a03928316926323b872dd92911690309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561086457600080fd5b6102c65a03f1151561087557600080fd5b50505060405180519050151561088a57600080fd5b600180548301908190557f3a7e173a9698235104076a0f536c5169527adecd20d0938c7c156ddf776c859190839060405191825260208201526040908101905180910390a15060015b5b919050565b600654600160a060020a031681565b60015481565b6006546000908190819081908190600160a060020a0316635de28ae087836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561094957600080fd5b6102c65a03f1151561095a57600080fd5b50505060405180511515905061096f57600080fd5b600654600160a060020a031663c9100bcb8760006040516040015260405160e060020a63ffffffff841602815260048101919091526024016040805180830381600087803b15156109bf57600080fd5b6102c65a03f115156109d057600080fd5b50505060405180519060200180519194509092505081156109f057600080fd5b6109f8610b25565b600654909150600160a060020a031663a38824678760405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b1515610a4357600080fd5b6102c65a03f11515610a5457600080fd5b505050801515610ac757610a6783610c93565b600654909450600160a060020a031663c1203cb88760405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b1515610ab257600080fd5b6102c65a03f11515610ac357600080fd5b5050505b7f62c32f975d56ca09b1255432b0bbe6b55f0cea57bf9efab6714409e47a5bf6838684600254604051928352600160a060020a0390911660208301526040808301919091526060909101905180910390a18394505b50505050919050565b600a546009805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790556000808080610b827f94c4860d894e91f2df683b61455630d721209c6265d2e80c86a1f92cab14b370610d42565b9250610bad7facd7f5ed7d93b1526477b93e6c7def60c40420a868e7f694a7671413d89bb9a5610d42565b9150508181148015610c23577f3cf8334d363b8556cf92ddf8f803d5074c4f884fc6b92484e8f2b22fe3d5fc8d6040516020808252600a908201527f6e6f74206775696c7479000000000000000000000000000000000000000000006040808301919091526060909101905180910390a1610c89565b7f3cf8334d363b8556cf92ddf8f803d5074c4f884fc6b92484e8f2b22fe3d5fc8d60405160208082526006908201527f6775696c747900000000000000000000000000000000000000000000000000006040808301919091526060909101905180910390a15b8093505b50505090565b600754600254600091600160a060020a0316906323b872dd9030908590856040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610d0757600080fd5b6102c65a03f11515610d1857600080fd5b5050506040518051915050801515610d2f57600080fd5b600254600180549190910390555b919050565b600954600090819081908190600160a060020a0316630178b8bf86836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610d9b57600080fd5b6102c65a03f11515610dac57600080fd5b5050506040518051935083925050600160a060020a038216632dff69418660006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610e0957600080fd5b6102c65a03f11515610e1a57600080fd5b50505060405180519450849150505b505050919050565b815481835581811511610e5557600083815260209020610e55918101908301610e5b565b5b505050565b610e7991905b80821115610e755760008155600101610e61565b5090565b905600a165627a7a72305820d6f20a27063866e8e6bbcda3e265c116a6c54bcc90be3a7c42d227dea40eb5ef0029`

// DeploySwearGame deploys a new Ethereum contract, binding an instance of SwearGame to it.
func DeploySwearGame(auth *bind.TransactOpts, backend bind.ContractBackend, _caseContract common.Address, _sampleToken common.Address, _rewardCompensation *big.Int) (common.Address, *types.Transaction, *SwearGame, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearGameABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwearGameBin), backend, _caseContract, _sampleToken, _rewardCompensation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwearGame{SwearGameCaller: SwearGameCaller{contract: contract}, SwearGameTransactor: SwearGameTransactor{contract: contract}}, nil
}

// SwearGame is an auto generated Go binding around an Ethereum contract.
type SwearGame struct {
	SwearGameCaller     // Read-only binding to the contract
	SwearGameTransactor // Write-only binding to the contract
}

// SwearGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwearGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwearGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwearGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwearGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwearGameSession struct {
	Contract     *SwearGame        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwearGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwearGameCallerSession struct {
	Contract *SwearGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SwearGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwearGameTransactorSession struct {
	Contract     *SwearGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SwearGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwearGameRaw struct {
	Contract *SwearGame // Generic contract binding to access the raw methods on
}

// SwearGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwearGameCallerRaw struct {
	Contract *SwearGameCaller // Generic read-only contract binding to access the raw methods on
}

// SwearGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwearGameTransactorRaw struct {
	Contract *SwearGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwearGame creates a new instance of SwearGame, bound to a specific deployed contract.
func NewSwearGame(address common.Address, backend bind.ContractBackend) (*SwearGame, error) {
	contract, err := bindSwearGame(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwearGame{SwearGameCaller: SwearGameCaller{contract: contract}, SwearGameTransactor: SwearGameTransactor{contract: contract}}, nil
}

// NewSwearGameCaller creates a new read-only instance of SwearGame, bound to a specific deployed contract.
func NewSwearGameCaller(address common.Address, caller bind.ContractCaller) (*SwearGameCaller, error) {
	contract, err := bindSwearGame(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SwearGameCaller{contract: contract}, nil
}

// NewSwearGameTransactor creates a new write-only instance of SwearGame, bound to a specific deployed contract.
func NewSwearGameTransactor(address common.Address, transactor bind.ContractTransactor) (*SwearGameTransactor, error) {
	contract, err := bindSwearGame(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SwearGameTransactor{contract: contract}, nil
}

// bindSwearGame binds a generic wrapper to an already deployed contract.
func bindSwearGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwearGame *SwearGameRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwearGame.Contract.SwearGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwearGame *SwearGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwearGame.Contract.SwearGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwearGame *SwearGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwearGame.Contract.SwearGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwearGame *SwearGameCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwearGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwearGame *SwearGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwearGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwearGame *SwearGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwearGame.Contract.contract.Transact(opts, method, params...)
}

// AmountStaked is a free data retrieval call binding the contract method 0xfab1e747.
//
// Solidity: function amountStaked() constant returns(uint256)
func (_SwearGame *SwearGameCaller) AmountStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "amountStaked")
	return *ret0, err
}

// AmountStaked is a free data retrieval call binding the contract method 0xfab1e747.
//
// Solidity: function amountStaked() constant returns(uint256)
func (_SwearGame *SwearGameSession) AmountStaked() (*big.Int, error) {
	return _SwearGame.Contract.AmountStaked(&_SwearGame.CallOpts)
}

// AmountStaked is a free data retrieval call binding the contract method 0xfab1e747.
//
// Solidity: function amountStaked() constant returns(uint256)
func (_SwearGame *SwearGameCallerSession) AmountStaked() (*big.Int, error) {
	return _SwearGame.Contract.AmountStaked(&_SwearGame.CallOpts)
}

// CaseContract is a free data retrieval call binding the contract method 0xd556e8e6.
//
// Solidity: function caseContract() constant returns(address)
func (_SwearGame *SwearGameCaller) CaseContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "caseContract")
	return *ret0, err
}

// CaseContract is a free data retrieval call binding the contract method 0xd556e8e6.
//
// Solidity: function caseContract() constant returns(address)
func (_SwearGame *SwearGameSession) CaseContract() (common.Address, error) {
	return _SwearGame.Contract.CaseContract(&_SwearGame.CallOpts)
}

// CaseContract is a free data retrieval call binding the contract method 0xd556e8e6.
//
// Solidity: function caseContract() constant returns(address)
func (_SwearGame *SwearGameCallerSession) CaseContract() (common.Address, error) {
	return _SwearGame.Contract.CaseContract(&_SwearGame.CallOpts)
}

// ClaimId is a free data retrieval call binding the contract method 0x5c25fd13.
//
// Solidity: function claimId() constant returns(bytes32)
func (_SwearGame *SwearGameCaller) ClaimId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "claimId")
	return *ret0, err
}

// ClaimId is a free data retrieval call binding the contract method 0x5c25fd13.
//
// Solidity: function claimId() constant returns(bytes32)
func (_SwearGame *SwearGameSession) ClaimId() ([32]byte, error) {
	return _SwearGame.Contract.ClaimId(&_SwearGame.CallOpts)
}

// ClaimId is a free data retrieval call binding the contract method 0x5c25fd13.
//
// Solidity: function claimId() constant returns(bytes32)
func (_SwearGame *SwearGameCallerSession) ClaimId() ([32]byte, error) {
	return _SwearGame.Contract.ClaimId(&_SwearGame.CallOpts)
}

// ClientsClaimsIds is a free data retrieval call binding the contract method 0x3258e6ca.
//
// Solidity: function clientsClaimsIds( address,  uint256) constant returns(bytes32)
func (_SwearGame *SwearGameCaller) ClientsClaimsIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "clientsClaimsIds", arg0, arg1)
	return *ret0, err
}

// ClientsClaimsIds is a free data retrieval call binding the contract method 0x3258e6ca.
//
// Solidity: function clientsClaimsIds( address,  uint256) constant returns(bytes32)
func (_SwearGame *SwearGameSession) ClientsClaimsIds(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _SwearGame.Contract.ClientsClaimsIds(&_SwearGame.CallOpts, arg0, arg1)
}

// ClientsClaimsIds is a free data retrieval call binding the contract method 0x3258e6ca.
//
// Solidity: function clientsClaimsIds( address,  uint256) constant returns(bytes32)
func (_SwearGame *SwearGameCallerSession) ClientsClaimsIds(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _SwearGame.Contract.ClientsClaimsIds(&_SwearGame.CallOpts, arg0, arg1)
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_SwearGame *SwearGameCaller) EnsAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "ensAddress")
	return *ret0, err
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_SwearGame *SwearGameSession) EnsAddress() (common.Address, error) {
	return _SwearGame.Contract.EnsAddress(&_SwearGame.CallOpts)
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_SwearGame *SwearGameCallerSession) EnsAddress() (common.Address, error) {
	return _SwearGame.Contract.EnsAddress(&_SwearGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearGame *SwearGameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearGame *SwearGameSession) Owner() (common.Address, error) {
	return _SwearGame.Contract.Owner(&_SwearGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearGame *SwearGameCallerSession) Owner() (common.Address, error) {
	return _SwearGame.Contract.Owner(&_SwearGame.CallOpts)
}

// RegisteredPlayers is a free data retrieval call binding the contract method 0x4bafce22.
//
// Solidity: function registeredPlayers( address) constant returns(bool)
func (_SwearGame *SwearGameCaller) RegisteredPlayers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "registeredPlayers", arg0)
	return *ret0, err
}

// RegisteredPlayers is a free data retrieval call binding the contract method 0x4bafce22.
//
// Solidity: function registeredPlayers( address) constant returns(bool)
func (_SwearGame *SwearGameSession) RegisteredPlayers(arg0 common.Address) (bool, error) {
	return _SwearGame.Contract.RegisteredPlayers(&_SwearGame.CallOpts, arg0)
}

// RegisteredPlayers is a free data retrieval call binding the contract method 0x4bafce22.
//
// Solidity: function registeredPlayers( address) constant returns(bool)
func (_SwearGame *SwearGameCallerSession) RegisteredPlayers(arg0 common.Address) (bool, error) {
	return _SwearGame.Contract.RegisteredPlayers(&_SwearGame.CallOpts, arg0)
}

// RegisteredPlayersCounter is a free data retrieval call binding the contract method 0x4a3ebbed.
//
// Solidity: function registeredPlayersCounter() constant returns(uint256)
func (_SwearGame *SwearGameCaller) RegisteredPlayersCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "registeredPlayersCounter")
	return *ret0, err
}

// RegisteredPlayersCounter is a free data retrieval call binding the contract method 0x4a3ebbed.
//
// Solidity: function registeredPlayersCounter() constant returns(uint256)
func (_SwearGame *SwearGameSession) RegisteredPlayersCounter() (*big.Int, error) {
	return _SwearGame.Contract.RegisteredPlayersCounter(&_SwearGame.CallOpts)
}

// RegisteredPlayersCounter is a free data retrieval call binding the contract method 0x4a3ebbed.
//
// Solidity: function registeredPlayersCounter() constant returns(uint256)
func (_SwearGame *SwearGameCallerSession) RegisteredPlayersCounter() (*big.Int, error) {
	return _SwearGame.Contract.RegisteredPlayersCounter(&_SwearGame.CallOpts)
}

// RewardCompensation is a free data retrieval call binding the contract method 0x66be2236.
//
// Solidity: function rewardCompensation() constant returns(uint256)
func (_SwearGame *SwearGameCaller) RewardCompensation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "rewardCompensation")
	return *ret0, err
}

// RewardCompensation is a free data retrieval call binding the contract method 0x66be2236.
//
// Solidity: function rewardCompensation() constant returns(uint256)
func (_SwearGame *SwearGameSession) RewardCompensation() (*big.Int, error) {
	return _SwearGame.Contract.RewardCompensation(&_SwearGame.CallOpts)
}

// RewardCompensation is a free data retrieval call binding the contract method 0x66be2236.
//
// Solidity: function rewardCompensation() constant returns(uint256)
func (_SwearGame *SwearGameCallerSession) RewardCompensation() (*big.Int, error) {
	return _SwearGame.Contract.RewardCompensation(&_SwearGame.CallOpts)
}

// SampleToken is a free data retrieval call binding the contract method 0x7222e9e3.
//
// Solidity: function sampleToken() constant returns(address)
func (_SwearGame *SwearGameCaller) SampleToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "sampleToken")
	return *ret0, err
}

// SampleToken is a free data retrieval call binding the contract method 0x7222e9e3.
//
// Solidity: function sampleToken() constant returns(address)
func (_SwearGame *SwearGameSession) SampleToken() (common.Address, error) {
	return _SwearGame.Contract.SampleToken(&_SwearGame.CallOpts)
}

// SampleToken is a free data retrieval call binding the contract method 0x7222e9e3.
//
// Solidity: function sampleToken() constant returns(address)
func (_SwearGame *SwearGameCallerSession) SampleToken() (common.Address, error) {
	return _SwearGame.Contract.SampleToken(&_SwearGame.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearGame *SwearGameTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearGame *SwearGameSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.ChangeOwner(&_SwearGame.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearGame *SwearGameTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.ChangeOwner(&_SwearGame.TransactOpts, _newOwner)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_depositAmount uint256) returns(bool)
func (_SwearGame *SwearGameTransactor) Deposit(opts *bind.TransactOpts, _depositAmount *big.Int) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "deposit", _depositAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_depositAmount uint256) returns(bool)
func (_SwearGame *SwearGameSession) Deposit(_depositAmount *big.Int) (*types.Transaction, error) {
	return _SwearGame.Contract.Deposit(&_SwearGame.TransactOpts, _depositAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_depositAmount uint256) returns(bool)
func (_SwearGame *SwearGameTransactorSession) Deposit(_depositAmount *big.Int) (*types.Transaction, error) {
	return _SwearGame.Contract.Deposit(&_SwearGame.TransactOpts, _depositAmount)
}

// LeaveGame is a paid mutator transaction binding the contract method 0x3e95d955.
//
// Solidity: function leaveGame(_player address) returns()
func (_SwearGame *SwearGameTransactor) LeaveGame(opts *bind.TransactOpts, _player common.Address) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "leaveGame", _player)
}

// LeaveGame is a paid mutator transaction binding the contract method 0x3e95d955.
//
// Solidity: function leaveGame(_player address) returns()
func (_SwearGame *SwearGameSession) LeaveGame(_player common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.LeaveGame(&_SwearGame.TransactOpts, _player)
}

// LeaveGame is a paid mutator transaction binding the contract method 0x3e95d955.
//
// Solidity: function leaveGame(_player address) returns()
func (_SwearGame *SwearGameTransactorSession) LeaveGame(_player common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.LeaveGame(&_SwearGame.TransactOpts, _player)
}

// OpenNewClaim is a paid mutator transaction binding the contract method 0xa19e95da.
//
// Solidity: function openNewClaim(_evidence bytes32) returns(bool)
func (_SwearGame *SwearGameTransactor) OpenNewClaim(opts *bind.TransactOpts, _evidence [32]byte) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "openNewClaim", _evidence)
}

// OpenNewClaim is a paid mutator transaction binding the contract method 0xa19e95da.
//
// Solidity: function openNewClaim(_evidence bytes32) returns(bool)
func (_SwearGame *SwearGameSession) OpenNewClaim(_evidence [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.OpenNewClaim(&_SwearGame.TransactOpts, _evidence)
}

// OpenNewClaim is a paid mutator transaction binding the contract method 0xa19e95da.
//
// Solidity: function openNewClaim(_evidence bytes32) returns(bool)
func (_SwearGame *SwearGameTransactorSession) OpenNewClaim(_evidence [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.OpenNewClaim(&_SwearGame.TransactOpts, _evidence)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(registered bool)
func (_SwearGame *SwearGameTransactor) Register(opts *bind.TransactOpts, _player common.Address) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "register", _player)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(registered bool)
func (_SwearGame *SwearGameSession) Register(_player common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.Register(&_SwearGame.TransactOpts, _player)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(registered bool)
func (_SwearGame *SwearGameTransactorSession) Register(_player common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.Register(&_SwearGame.TransactOpts, _player)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_SwearGame *SwearGameTransactor) SetENSAddress(opts *bind.TransactOpts, _ensAddr common.Address) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "setENSAddress", _ensAddr)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_SwearGame *SwearGameSession) SetENSAddress(_ensAddr common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.SetENSAddress(&_SwearGame.TransactOpts, _ensAddr)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_SwearGame *SwearGameTransactorSession) SetENSAddress(_ensAddr common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.SetENSAddress(&_SwearGame.TransactOpts, _ensAddr)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

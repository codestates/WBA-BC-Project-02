// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package credit

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

// CreditMetaData contains all meta data concerning the Credit contract.
var CreditMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CustomTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"DexAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"}],\"name\":\"addDexAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"name\":\"checkTokenPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"creditToWemix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"name\":\"mintSelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"}],\"name\":\"reduceDexAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"}],\"name\":\"setDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600681526020017f43726564697400000000000000000000000000000000000000000000000000008152506040518060400160405280600681526020017f7057454d49580000000000000000000000000000000000000000000000000000815250816003908162000081919062000404565b50806004908162000093919062000404565b505050620000b6620000aa620000bc60201b60201c565b620000c460201b60201c565b620004eb565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200020c57607f821691505b602082108103620002225762000221620001c4565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200028c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200024d565b6200029886836200024d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620002e5620002df620002d984620002b0565b620002ba565b620002b0565b9050919050565b6000819050919050565b6200030183620002c4565b620003196200031082620002ec565b8484546200025a565b825550505050565b600090565b6200033062000321565b6200033d818484620002f6565b505050565b5b8181101562000365576200035960008262000326565b60018101905062000343565b5050565b601f821115620003b4576200037e8162000228565b62000389846200023d565b8101602085101562000399578190505b620003b1620003a8856200023d565b83018262000342565b50505b505050565b600082821c905092915050565b6000620003d960001984600802620003b9565b1980831691505092915050565b6000620003f48383620003c6565b9150826002028217905092915050565b6200040f826200018a565b67ffffffffffffffff8111156200042b576200042a62000195565b5b620004378254620001f3565b6200044482828562000369565b600060209050601f8311600181146200047c576000841562000467578287015190505b620004738582620003e6565b865550620004e3565b601f1984166200048c8662000228565b60005b82811015620004b6578489015182556001820191506020850194506020810190506200048f565b86831015620004d65784890151620004d2601f891682620003c6565b8355505b6001600288020188555050505b505050505050565b6126bb80620004fb6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063715018a6116100b8578063a457c2d71161007c578063a457c2d71461038b578063a49c70f2146103bb578063a9059cbb146103eb578063beabacc81461041b578063dd62ed3e1461044b578063f2fde38b1461047b57610142565b8063715018a6146102e5578063781ea389146102ef5780638da5cb5b1461031f578063949fded21461033d57806395d89b411461036d57610142565b8063395093511161010a5780633950935114610201578063447fa8b714610231578063485699c71461024d57806350f85d341461026957806365dd6f2b1461028557806370a08231146102b557610142565b806306fdde0314610147578063095ea7b31461016557806318160ddd1461019557806323b872dd146101b3578063313ce567146101e3575b600080fd5b61014f610497565b60405161015c91906116fb565b60405180910390f35b61017f600480360381019061017a91906117c5565b610529565b60405161018c9190611820565b60405180910390f35b61019d61054c565b6040516101aa919061184a565b60405180910390f35b6101cd60048036038101906101c89190611865565b610556565b6040516101da9190611820565b60405180910390f35b6101eb610585565b6040516101f891906118d4565b60405180910390f35b61021b600480360381019061021691906117c5565b61058e565b6040516102289190611820565b60405180910390f35b61024b600480360381019061024691906118ef565b6105c5565b005b6102676004803603810190610262919061195a565b610611565b005b610283600480360381019061027e9190611acf565b61072c565b005b61029f600480360381019061029a9190611b2b565b6107c7565b6040516102ac919061184a565b60405180910390f35b6102cf60048036038101906102ca91906118ef565b6107f5565b6040516102dc919061184a565b60405180910390f35b6102ed61083d565b005b61030960048036038101906103049190611b74565b610851565b604051610316919061184a565b60405180910390f35b610327610909565b6040516103349190611bdf565b60405180910390f35b61035760048036038101906103529190611b2b565b610933565b604051610364919061184a565b60405180910390f35b61037561095b565b60405161038291906116fb565b60405180910390f35b6103a560048036038101906103a091906117c5565b6109ed565b6040516103b29190611820565b60405180910390f35b6103d560048036038101906103d09190611b74565b610a64565b6040516103e2919061184a565b60405180910390f35b610405600480360381019061040091906117c5565b610b1c565b6040516104129190611820565b60405180910390f35b61043560048036038101906104309190611865565b610b3f565b6040516104429190611820565b60405180910390f35b61046560048036038101906104609190611bfa565b610bf2565b604051610472919061184a565b60405180910390f35b610495600480360381019061049091906118ef565b610c79565b005b6060600380546104a690611c69565b80601f01602080910402602001604051908101604052809291908181526020018280546104d290611c69565b801561051f5780601f106104f45761010080835404028352916020019161051f565b820191906000526020600020905b81548152906001019060200180831161050257829003601f168201915b5050505050905090565b600080610534610cfc565b9050610541818585610d04565b600191505092915050565b6000600254905090565b600080610561610cfc565b905061056e858285610ecd565b610579858585610f59565b60019150509392505050565b60006012905090565b600080610599610cfc565b90506105ba8185856105ab8589610bf2565b6105b59190611cc9565b610d04565b600191505092915050565b6105cd6111cf565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066b57600080fd5b61067d826106788361124d565b611278565b8173ffffffffffffffffffffffffffffffffffffffff166108fc6106a08361124d565b9081150290604051600060405180830381858888f193505050501580156106cb573d6000803e3d6000fd5b507f5620e542b9ce6a03dbe5af2c6483e766af6c96439562b26a0dcd2edd7099524b82826106f8856107f5565b8573ffffffffffffffffffffffffffffffffffffffff16316040516107209493929190611d5c565b60405180910390a15050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461078657600080fd5b610798336107938461124d565b611445565b6107a18261124d565b6007826040516107b19190611ddd565b9081526020016040518091039020819055505050565b6007818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6108456111cf565b61084f600061159b565b565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ad57600080fd5b816007846040516108be9190611ddd565b908152602001604051809103902060008282546108db9190611df4565b925050819055506007836040516108f29190611ddd565b908152602001604051809103902054905092915050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006007826040516109459190611ddd565b9081526020016040518091039020549050919050565b60606004805461096a90611c69565b80601f016020809104026020016040519081016040528092919081815260200182805461099690611c69565b80156109e35780601f106109b8576101008083540402835291602001916109e3565b820191906000526020600020905b8154815290600101906020018083116109c657829003601f168201915b5050505050905090565b6000806109f8610cfc565b90506000610a068286610bf2565b905083811015610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4290611e9a565b60405180910390fd5b610a588286868403610d04565b60019250505092915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac057600080fd5b81600784604051610ad19190611ddd565b90815260200160405180910390206000828254610aee9190611cc9565b92505081905550600783604051610b059190611ddd565b908152602001604051809103902054905092915050565b600080610b27610cfc565b9050610b34818585610f59565b600191505092915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b9b57600080fd5b6000849050610bab818585610f59565b7fe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be858585604051610bde93929190611eba565b60405180910390a160019150509392505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610c816111cf565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cf0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce790611f63565b60405180910390fd5b610cf98161159b565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6a90611ff5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610de2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd990612087565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610ec0919061184a565b60405180910390a3505050565b6000610ed98484610bf2565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610f535781811015610f45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3c906120f3565b60405180910390fd5b610f528484848403610d04565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610fc8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbf90612185565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611037576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102e90612217565b60405180910390fd5b611042838383611661565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156110c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110bf906122a9565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111b6919061184a565b60405180910390a36111c9848484611666565b50505050565b6111d7610cfc565b73ffffffffffffffffffffffffffffffffffffffff166111f5610909565b73ffffffffffffffffffffffffffffffffffffffff161461124b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124290612315565b60405180910390fd5b565b6000611257610585565b60ff16600a6112669190612468565b8261127191906124b3565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036112e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112de90612567565b60405180910390fd5b6112f382600083611661565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611379576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611370906125f9565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161142c919061184a565b60405180910390a361144083600084611666565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036114b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ab90612665565b60405180910390fd5b6114c060008383611661565b80600260008282546114d29190611cc9565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051611583919061184a565b60405180910390a361159760008383611666565b5050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156116a557808201518184015260208101905061168a565b60008484015250505050565b6000601f19601f8301169050919050565b60006116cd8261166b565b6116d78185611676565b93506116e7818560208601611687565b6116f0816116b1565b840191505092915050565b6000602082019050818103600083015261171581846116c2565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061175c82611731565b9050919050565b61176c81611751565b811461177757600080fd5b50565b60008135905061178981611763565b92915050565b6000819050919050565b6117a28161178f565b81146117ad57600080fd5b50565b6000813590506117bf81611799565b92915050565b600080604083850312156117dc576117db611727565b5b60006117ea8582860161177a565b92505060206117fb858286016117b0565b9150509250929050565b60008115159050919050565b61181a81611805565b82525050565b60006020820190506118356000830184611811565b92915050565b6118448161178f565b82525050565b600060208201905061185f600083018461183b565b92915050565b60008060006060848603121561187e5761187d611727565b5b600061188c8682870161177a565b935050602061189d8682870161177a565b92505060406118ae868287016117b0565b9150509250925092565b600060ff82169050919050565b6118ce816118b8565b82525050565b60006020820190506118e960008301846118c5565b92915050565b60006020828403121561190557611904611727565b5b60006119138482850161177a565b91505092915050565b600061192782611731565b9050919050565b6119378161191c565b811461194257600080fd5b50565b6000813590506119548161192e565b92915050565b6000806040838503121561197157611970611727565b5b600061197f85828601611945565b9250506020611990858286016117b0565b9150509250929050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6119dc826116b1565b810181811067ffffffffffffffff821117156119fb576119fa6119a4565b5b80604052505050565b6000611a0e61171d565b9050611a1a82826119d3565b919050565b600067ffffffffffffffff821115611a3a57611a396119a4565b5b611a43826116b1565b9050602081019050919050565b82818337600083830152505050565b6000611a72611a6d84611a1f565b611a04565b905082815260208101848484011115611a8e57611a8d61199f565b5b611a99848285611a50565b509392505050565b600082601f830112611ab657611ab561199a565b5b8135611ac6848260208601611a5f565b91505092915050565b60008060408385031215611ae657611ae5611727565b5b6000611af4858286016117b0565b925050602083013567ffffffffffffffff811115611b1557611b1461172c565b5b611b2185828601611aa1565b9150509250929050565b600060208284031215611b4157611b40611727565b5b600082013567ffffffffffffffff811115611b5f57611b5e61172c565b5b611b6b84828501611aa1565b91505092915050565b60008060408385031215611b8b57611b8a611727565b5b600083013567ffffffffffffffff811115611ba957611ba861172c565b5b611bb585828601611aa1565b9250506020611bc6858286016117b0565b9150509250929050565b611bd981611751565b82525050565b6000602082019050611bf46000830184611bd0565b92915050565b60008060408385031215611c1157611c10611727565b5b6000611c1f8582860161177a565b9250506020611c308582860161177a565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611c8157607f821691505b602082108103611c9457611c93611c3a565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611cd48261178f565b9150611cdf8361178f565b9250828201905080821115611cf757611cf6611c9a565b5b92915050565b6000819050919050565b6000611d22611d1d611d1884611731565b611cfd565b611731565b9050919050565b6000611d3482611d07565b9050919050565b6000611d4682611d29565b9050919050565b611d5681611d3b565b82525050565b6000608082019050611d716000830187611d4d565b611d7e602083018661183b565b611d8b604083018561183b565b611d98606083018461183b565b95945050505050565b600081905092915050565b6000611db78261166b565b611dc18185611da1565b9350611dd1818560208601611687565b80840191505092915050565b6000611de98284611dac565b915081905092915050565b6000611dff8261178f565b9150611e0a8361178f565b9250828203905081811115611e2257611e21611c9a565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611e84602583611676565b9150611e8f82611e28565b604082019050919050565b60006020820190508181036000830152611eb381611e77565b9050919050565b6000606082019050611ecf6000830186611bd0565b611edc6020830185611bd0565b611ee9604083018461183b565b949350505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611f4d602683611676565b9150611f5882611ef1565b604082019050919050565b60006020820190508181036000830152611f7c81611f40565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611fdf602483611676565b9150611fea82611f83565b604082019050919050565b6000602082019050818103600083015261200e81611fd2565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000612071602283611676565b915061207c82612015565b604082019050919050565b600060208201905081810360008301526120a081612064565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b60006120dd601d83611676565b91506120e8826120a7565b602082019050919050565b6000602082019050818103600083015261210c816120d0565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061216f602583611676565b915061217a82612113565b604082019050919050565b6000602082019050818103600083015261219e81612162565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000612201602383611676565b915061220c826121a5565b604082019050919050565b60006020820190508181036000830152612230816121f4565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000612293602683611676565b915061229e82612237565b604082019050919050565b600060208201905081810360008301526122c281612286565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006122ff602083611676565b915061230a826122c9565b602082019050919050565b6000602082019050818103600083015261232e816122f2565b9050919050565b60008160011c9050919050565b6000808291508390505b600185111561238c5780860481111561236857612367611c9a565b5b60018516156123775780820291505b808102905061238585612335565b945061234c565b94509492505050565b6000826123a55760019050612461565b816123b35760009050612461565b81600181146123c957600281146123d357612402565b6001915050612461565b60ff8411156123e5576123e4611c9a565b5b8360020a9150848211156123fc576123fb611c9a565b5b50612461565b5060208310610133831016604e8410600b84101617156124375782820a90508381111561243257612431611c9a565b5b612461565b6124448484846001612342565b9250905081840481111561245b5761245a611c9a565b5b81810290505b9392505050565b60006124738261178f565b915061247e8361178f565b92506124ab7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612395565b905092915050565b60006124be8261178f565b91506124c98361178f565b92508282026124d78161178f565b915082820484148315176124ee576124ed611c9a565b5b5092915050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000612551602183611676565b915061255c826124f5565b604082019050919050565b6000602082019050818103600083015261258081612544565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b60006125e3602283611676565b91506125ee82612587565b604082019050919050565b60006020820190508181036000830152612612816125d6565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b600061264f601f83611676565b915061265a82612619565b602082019050919050565b6000602082019050818103600083015261267e81612642565b905091905056fea264697066735822122060f7e408c08bdcd3327f9530c77edc91d82acec3392e8561c2f13302e46d0c1564736f6c63430008110033",
}

// CreditABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditMetaData.ABI instead.
var CreditABI = CreditMetaData.ABI

// CreditBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreditMetaData.Bin instead.
var CreditBin = CreditMetaData.Bin

// DeployCredit deploys a new Ethereum contract, binding an instance of Credit to it.
func DeployCredit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Credit, error) {
	parsed, err := CreditMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreditBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Credit{CreditCaller: CreditCaller{contract: contract}, CreditTransactor: CreditTransactor{contract: contract}, CreditFilterer: CreditFilterer{contract: contract}}, nil
}

// Credit is an auto generated Go binding around an Ethereum contract.
type Credit struct {
	CreditCaller     // Read-only binding to the contract
	CreditTransactor // Write-only binding to the contract
	CreditFilterer   // Log filterer for contract events
}

// CreditCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditSession struct {
	Contract     *Credit           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditCallerSession struct {
	Contract *CreditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CreditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditTransactorSession struct {
	Contract     *CreditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditRaw struct {
	Contract *Credit // Generic contract binding to access the raw methods on
}

// CreditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditCallerRaw struct {
	Contract *CreditCaller // Generic read-only contract binding to access the raw methods on
}

// CreditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditTransactorRaw struct {
	Contract *CreditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredit creates a new instance of Credit, bound to a specific deployed contract.
func NewCredit(address common.Address, backend bind.ContractBackend) (*Credit, error) {
	contract, err := bindCredit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Credit{CreditCaller: CreditCaller{contract: contract}, CreditTransactor: CreditTransactor{contract: contract}, CreditFilterer: CreditFilterer{contract: contract}}, nil
}

// NewCreditCaller creates a new read-only instance of Credit, bound to a specific deployed contract.
func NewCreditCaller(address common.Address, caller bind.ContractCaller) (*CreditCaller, error) {
	contract, err := bindCredit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditCaller{contract: contract}, nil
}

// NewCreditTransactor creates a new write-only instance of Credit, bound to a specific deployed contract.
func NewCreditTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditTransactor, error) {
	contract, err := bindCredit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditTransactor{contract: contract}, nil
}

// NewCreditFilterer creates a new log filterer instance of Credit, bound to a specific deployed contract.
func NewCreditFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFilterer, error) {
	contract, err := bindCredit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFilterer{contract: contract}, nil
}

// bindCredit binds a generic wrapper to an already deployed contract.
func bindCredit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credit *CreditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credit.Contract.CreditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credit *CreditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.Contract.CreditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credit *CreditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credit.Contract.CreditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credit *CreditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credit *CreditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credit *CreditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credit.Contract.contract.Transact(opts, method, params...)
}

// DexAmount is a free data retrieval call binding the contract method 0x65dd6f2b.
//
// Solidity: function DexAmount(string ) view returns(uint256)
func (_Credit *CreditCaller) DexAmount(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "DexAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DexAmount is a free data retrieval call binding the contract method 0x65dd6f2b.
//
// Solidity: function DexAmount(string ) view returns(uint256)
func (_Credit *CreditSession) DexAmount(arg0 string) (*big.Int, error) {
	return _Credit.Contract.DexAmount(&_Credit.CallOpts, arg0)
}

// DexAmount is a free data retrieval call binding the contract method 0x65dd6f2b.
//
// Solidity: function DexAmount(string ) view returns(uint256)
func (_Credit *CreditCallerSession) DexAmount(arg0 string) (*big.Int, error) {
	return _Credit.Contract.DexAmount(&_Credit.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credit *CreditCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credit *CreditSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Credit.Contract.Allowance(&_Credit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credit *CreditCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Credit.Contract.Allowance(&_Credit.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credit *CreditCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credit *CreditSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Credit.Contract.BalanceOf(&_Credit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credit *CreditCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Credit.Contract.BalanceOf(&_Credit.CallOpts, account)
}

// CheckTokenPool is a free data retrieval call binding the contract method 0x949fded2.
//
// Solidity: function checkTokenPool(string tokenName) view returns(uint256)
func (_Credit *CreditCaller) CheckTokenPool(opts *bind.CallOpts, tokenName string) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "checkTokenPool", tokenName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckTokenPool is a free data retrieval call binding the contract method 0x949fded2.
//
// Solidity: function checkTokenPool(string tokenName) view returns(uint256)
func (_Credit *CreditSession) CheckTokenPool(tokenName string) (*big.Int, error) {
	return _Credit.Contract.CheckTokenPool(&_Credit.CallOpts, tokenName)
}

// CheckTokenPool is a free data retrieval call binding the contract method 0x949fded2.
//
// Solidity: function checkTokenPool(string tokenName) view returns(uint256)
func (_Credit *CreditCallerSession) CheckTokenPool(tokenName string) (*big.Int, error) {
	return _Credit.Contract.CheckTokenPool(&_Credit.CallOpts, tokenName)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Credit *CreditCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Credit *CreditSession) Decimals() (uint8, error) {
	return _Credit.Contract.Decimals(&_Credit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Credit *CreditCallerSession) Decimals() (uint8, error) {
	return _Credit.Contract.Decimals(&_Credit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Credit *CreditCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Credit *CreditSession) Name() (string, error) {
	return _Credit.Contract.Name(&_Credit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Credit *CreditCallerSession) Name() (string, error) {
	return _Credit.Contract.Name(&_Credit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credit *CreditCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credit *CreditSession) Owner() (common.Address, error) {
	return _Credit.Contract.Owner(&_Credit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credit *CreditCallerSession) Owner() (common.Address, error) {
	return _Credit.Contract.Owner(&_Credit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Credit *CreditCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Credit *CreditSession) Symbol() (string, error) {
	return _Credit.Contract.Symbol(&_Credit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Credit *CreditCallerSession) Symbol() (string, error) {
	return _Credit.Contract.Symbol(&_Credit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credit *CreditCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credit *CreditSession) TotalSupply() (*big.Int, error) {
	return _Credit.Contract.TotalSupply(&_Credit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credit *CreditCallerSession) TotalSupply() (*big.Int, error) {
	return _Credit.Contract.TotalSupply(&_Credit.CallOpts)
}

// AddDexAmount is a paid mutator transaction binding the contract method 0xa49c70f2.
//
// Solidity: function addDexAmount(string tokenName, uint256 creditAmount) returns(uint256)
func (_Credit *CreditTransactor) AddDexAmount(opts *bind.TransactOpts, tokenName string, creditAmount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "addDexAmount", tokenName, creditAmount)
}

// AddDexAmount is a paid mutator transaction binding the contract method 0xa49c70f2.
//
// Solidity: function addDexAmount(string tokenName, uint256 creditAmount) returns(uint256)
func (_Credit *CreditSession) AddDexAmount(tokenName string, creditAmount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.AddDexAmount(&_Credit.TransactOpts, tokenName, creditAmount)
}

// AddDexAmount is a paid mutator transaction binding the contract method 0xa49c70f2.
//
// Solidity: function addDexAmount(string tokenName, uint256 creditAmount) returns(uint256)
func (_Credit *CreditTransactorSession) AddDexAmount(tokenName string, creditAmount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.AddDexAmount(&_Credit.TransactOpts, tokenName, creditAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Credit *CreditTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Credit *CreditSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Approve(&_Credit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Credit *CreditTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Approve(&_Credit.TransactOpts, spender, amount)
}

// CreditToWemix is a paid mutator transaction binding the contract method 0x485699c7.
//
// Solidity: function creditToWemix(address to, uint256 amount) returns()
func (_Credit *CreditTransactor) CreditToWemix(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "creditToWemix", to, amount)
}

// CreditToWemix is a paid mutator transaction binding the contract method 0x485699c7.
//
// Solidity: function creditToWemix(address to, uint256 amount) returns()
func (_Credit *CreditSession) CreditToWemix(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.CreditToWemix(&_Credit.TransactOpts, to, amount)
}

// CreditToWemix is a paid mutator transaction binding the contract method 0x485699c7.
//
// Solidity: function creditToWemix(address to, uint256 amount) returns()
func (_Credit *CreditTransactorSession) CreditToWemix(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.CreditToWemix(&_Credit.TransactOpts, to, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Credit *CreditTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Credit *CreditSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.DecreaseAllowance(&_Credit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Credit *CreditTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.DecreaseAllowance(&_Credit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Credit *CreditTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Credit *CreditSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.IncreaseAllowance(&_Credit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Credit *CreditTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.IncreaseAllowance(&_Credit.TransactOpts, spender, addedValue)
}

// MintSelf is a paid mutator transaction binding the contract method 0x50f85d34.
//
// Solidity: function mintSelf(uint256 amount, string tokenName) returns()
func (_Credit *CreditTransactor) MintSelf(opts *bind.TransactOpts, amount *big.Int, tokenName string) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "mintSelf", amount, tokenName)
}

// MintSelf is a paid mutator transaction binding the contract method 0x50f85d34.
//
// Solidity: function mintSelf(uint256 amount, string tokenName) returns()
func (_Credit *CreditSession) MintSelf(amount *big.Int, tokenName string) (*types.Transaction, error) {
	return _Credit.Contract.MintSelf(&_Credit.TransactOpts, amount, tokenName)
}

// MintSelf is a paid mutator transaction binding the contract method 0x50f85d34.
//
// Solidity: function mintSelf(uint256 amount, string tokenName) returns()
func (_Credit *CreditTransactorSession) MintSelf(amount *big.Int, tokenName string) (*types.Transaction, error) {
	return _Credit.Contract.MintSelf(&_Credit.TransactOpts, amount, tokenName)
}

// ReduceDexAmount is a paid mutator transaction binding the contract method 0x781ea389.
//
// Solidity: function reduceDexAmount(string tokenName, uint256 creditAmount) returns(uint256)
func (_Credit *CreditTransactor) ReduceDexAmount(opts *bind.TransactOpts, tokenName string, creditAmount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "reduceDexAmount", tokenName, creditAmount)
}

// ReduceDexAmount is a paid mutator transaction binding the contract method 0x781ea389.
//
// Solidity: function reduceDexAmount(string tokenName, uint256 creditAmount) returns(uint256)
func (_Credit *CreditSession) ReduceDexAmount(tokenName string, creditAmount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.ReduceDexAmount(&_Credit.TransactOpts, tokenName, creditAmount)
}

// ReduceDexAmount is a paid mutator transaction binding the contract method 0x781ea389.
//
// Solidity: function reduceDexAmount(string tokenName, uint256 creditAmount) returns(uint256)
func (_Credit *CreditTransactorSession) ReduceDexAmount(tokenName string, creditAmount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.ReduceDexAmount(&_Credit.TransactOpts, tokenName, creditAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credit *CreditTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credit *CreditSession) RenounceOwnership() (*types.Transaction, error) {
	return _Credit.Contract.RenounceOwnership(&_Credit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credit *CreditTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Credit.Contract.RenounceOwnership(&_Credit.TransactOpts)
}

// SetDex is a paid mutator transaction binding the contract method 0x447fa8b7.
//
// Solidity: function setDex(address dex) returns()
func (_Credit *CreditTransactor) SetDex(opts *bind.TransactOpts, dex common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "setDex", dex)
}

// SetDex is a paid mutator transaction binding the contract method 0x447fa8b7.
//
// Solidity: function setDex(address dex) returns()
func (_Credit *CreditSession) SetDex(dex common.Address) (*types.Transaction, error) {
	return _Credit.Contract.SetDex(&_Credit.TransactOpts, dex)
}

// SetDex is a paid mutator transaction binding the contract method 0x447fa8b7.
//
// Solidity: function setDex(address dex) returns()
func (_Credit *CreditTransactorSession) SetDex(dex common.Address) (*types.Transaction, error) {
	return _Credit.Contract.SetDex(&_Credit.TransactOpts, dex)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Credit *CreditSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Transfer(&_Credit.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Transfer(&_Credit.TransactOpts, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactor) Transfer0(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "transfer0", from, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditSession) Transfer0(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Transfer0(&_Credit.TransactOpts, from, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactorSession) Transfer0(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Transfer0(&_Credit.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.TransferFrom(&_Credit.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.TransferFrom(&_Credit.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credit *CreditTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credit *CreditSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Credit.Contract.TransferOwnership(&_Credit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credit *CreditTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Credit.Contract.TransferOwnership(&_Credit.TransactOpts, newOwner)
}

// CreditApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Credit contract.
type CreditApprovalIterator struct {
	Event *CreditApproval // Event containing the contract specifics and raw log

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
func (it *CreditApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditApproval)
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
		it.Event = new(CreditApproval)
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
func (it *CreditApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditApproval represents a Approval event raised by the Credit contract.
type CreditApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Credit *CreditFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CreditApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CreditApprovalIterator{contract: _Credit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Credit *CreditFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CreditApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditApproval)
				if err := _Credit.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Credit *CreditFilterer) ParseApproval(log types.Log) (*CreditApproval, error) {
	event := new(CreditApproval)
	if err := _Credit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditCustomTransferIterator is returned from FilterCustomTransfer and is used to iterate over the raw logs and unpacked data for CustomTransfer events raised by the Credit contract.
type CreditCustomTransferIterator struct {
	Event *CreditCustomTransfer // Event containing the contract specifics and raw log

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
func (it *CreditCustomTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditCustomTransfer)
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
		it.Event = new(CreditCustomTransfer)
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
func (it *CreditCustomTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditCustomTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditCustomTransfer represents a CustomTransfer event raised by the Credit contract.
type CreditCustomTransfer struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCustomTransfer is a free log retrieval operation binding the contract event 0xe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be.
//
// Solidity: event CustomTransfer(address arg0, address arg1, uint256 arg2)
func (_Credit *CreditFilterer) FilterCustomTransfer(opts *bind.FilterOpts) (*CreditCustomTransferIterator, error) {

	logs, sub, err := _Credit.contract.FilterLogs(opts, "CustomTransfer")
	if err != nil {
		return nil, err
	}
	return &CreditCustomTransferIterator{contract: _Credit.contract, event: "CustomTransfer", logs: logs, sub: sub}, nil
}

// WatchCustomTransfer is a free log subscription operation binding the contract event 0xe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be.
//
// Solidity: event CustomTransfer(address arg0, address arg1, uint256 arg2)
func (_Credit *CreditFilterer) WatchCustomTransfer(opts *bind.WatchOpts, sink chan<- *CreditCustomTransfer) (event.Subscription, error) {

	logs, sub, err := _Credit.contract.WatchLogs(opts, "CustomTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditCustomTransfer)
				if err := _Credit.contract.UnpackLog(event, "CustomTransfer", log); err != nil {
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

// ParseCustomTransfer is a log parse operation binding the contract event 0xe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be.
//
// Solidity: event CustomTransfer(address arg0, address arg1, uint256 arg2)
func (_Credit *CreditFilterer) ParseCustomTransfer(log types.Log) (*CreditCustomTransfer, error) {
	event := new(CreditCustomTransfer)
	if err := _Credit.contract.UnpackLog(event, "CustomTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Credit contract.
type CreditMintIterator struct {
	Event *CreditMint // Event containing the contract specifics and raw log

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
func (it *CreditMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditMint)
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
		it.Event = new(CreditMint)
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
func (it *CreditMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditMint represents a Mint event raised by the Credit contract.
type CreditMint struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address arg0, uint256 arg1)
func (_Credit *CreditFilterer) FilterMint(opts *bind.FilterOpts) (*CreditMintIterator, error) {

	logs, sub, err := _Credit.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &CreditMintIterator{contract: _Credit.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address arg0, uint256 arg1)
func (_Credit *CreditFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CreditMint) (event.Subscription, error) {

	logs, sub, err := _Credit.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditMint)
				if err := _Credit.contract.UnpackLog(event, "Mint", log); err != nil {
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
// Solidity: event Mint(address arg0, uint256 arg1)
func (_Credit *CreditFilterer) ParseMint(log types.Log) (*CreditMint, error) {
	event := new(CreditMint)
	if err := _Credit.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Credit contract.
type CreditOwnershipTransferredIterator struct {
	Event *CreditOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CreditOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditOwnershipTransferred)
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
		it.Event = new(CreditOwnershipTransferred)
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
func (it *CreditOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditOwnershipTransferred represents a OwnershipTransferred event raised by the Credit contract.
type CreditOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Credit *CreditFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CreditOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreditOwnershipTransferredIterator{contract: _Credit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Credit *CreditFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CreditOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditOwnershipTransferred)
				if err := _Credit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Credit *CreditFilterer) ParseOwnershipTransferred(log types.Log) (*CreditOwnershipTransferred, error) {
	event := new(CreditOwnershipTransferred)
	if err := _Credit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Credit contract.
type CreditSwapIterator struct {
	Event *CreditSwap // Event containing the contract specifics and raw log

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
func (it *CreditSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditSwap)
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
		it.Event = new(CreditSwap)
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
func (it *CreditSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditSwap represents a Swap event raised by the Credit contract.
type CreditSwap struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Arg3 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x5620e542b9ce6a03dbe5af2c6483e766af6c96439562b26a0dcd2edd7099524b.
//
// Solidity: event Swap(address arg0, uint256 arg1, uint256 arg2, uint256 arg3)
func (_Credit *CreditFilterer) FilterSwap(opts *bind.FilterOpts) (*CreditSwapIterator, error) {

	logs, sub, err := _Credit.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &CreditSwapIterator{contract: _Credit.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x5620e542b9ce6a03dbe5af2c6483e766af6c96439562b26a0dcd2edd7099524b.
//
// Solidity: event Swap(address arg0, uint256 arg1, uint256 arg2, uint256 arg3)
func (_Credit *CreditFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *CreditSwap) (event.Subscription, error) {

	logs, sub, err := _Credit.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditSwap)
				if err := _Credit.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x5620e542b9ce6a03dbe5af2c6483e766af6c96439562b26a0dcd2edd7099524b.
//
// Solidity: event Swap(address arg0, uint256 arg1, uint256 arg2, uint256 arg3)
func (_Credit *CreditFilterer) ParseSwap(log types.Log) (*CreditSwap, error) {
	event := new(CreditSwap)
	if err := _Credit.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Credit contract.
type CreditTransferIterator struct {
	Event *CreditTransfer // Event containing the contract specifics and raw log

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
func (it *CreditTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditTransfer)
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
		it.Event = new(CreditTransfer)
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
func (it *CreditTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditTransfer represents a Transfer event raised by the Credit contract.
type CreditTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Credit *CreditFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CreditTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditTransferIterator{contract: _Credit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Credit *CreditFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CreditTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditTransfer)
				if err := _Credit.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Credit *CreditFilterer) ParseTransfer(log types.Log) (*CreditTransfer, error) {
	event := new(CreditTransfer)
	if err := _Credit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

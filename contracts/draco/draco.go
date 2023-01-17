// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package draco

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

// DracoMetaData contains all meta data concerning the Draco contract.
var DracoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CustomTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dexAddress\",\"type\":\"address\"}],\"name\":\"setDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600581526020017f447261636f0000000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f445241000000000000000000000000000000000000000000000000000000000081525081600390816200008f919062000412565b508060049081620000a1919062000412565b505050620000c4620000b8620000ca60201b60201c565b620000d260201b60201c565b620004f9565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200021a57607f821691505b60208210810362000230576200022f620001d2565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200029a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200025b565b620002a686836200025b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620002f3620002ed620002e784620002be565b620002c8565b620002be565b9050919050565b6000819050919050565b6200030f83620002d2565b620003276200031e82620002fa565b84845462000268565b825550505050565b600090565b6200033e6200032f565b6200034b81848462000304565b505050565b5b8181101562000373576200036760008262000334565b60018101905062000351565b5050565b601f821115620003c2576200038c8162000236565b62000397846200024b565b81016020851015620003a7578190505b620003bf620003b6856200024b565b83018262000350565b50505b505050565b600082821c905092915050565b6000620003e760001984600802620003c7565b1980831691505092915050565b6000620004028383620003d4565b9150826002028217905092915050565b6200041d8262000198565b67ffffffffffffffff811115620004395762000438620001a3565b5b62000445825462000201565b6200045282828562000377565b600060209050601f8311600181146200048a576000841562000475578287015190505b620004818582620003f4565b865550620004f1565b601f1984166200049a8662000236565b60005b82811015620004c4578489015182556001820191506020850194506020810190506200049d565b86831015620004e45784890151620004e0601f891682620003d4565b8355505b6001600288020188555050505b505050505050565b611c2e80620005096000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806370a08231116100a2578063a457c2d711610071578063a457c2d7146102a8578063a9059cbb146102d8578063beabacc814610308578063dd62ed3e14610338578063f2fde38b146103685761010b565b806370a0823114610232578063715018a6146102625780638da5cb5b1461026c57806395d89b411461028a5761010b565b8063313ce567116100de578063313ce567146101ac57806339509351146101ca57806340c10f19146101fa578063447fa8b7146102165761010b565b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015e57806323b872dd1461017c575b600080fd5b610118610384565b6040516101259190611157565b60405180910390f35b61014860048036038101906101439190611212565b610416565b604051610155919061126d565b60405180910390f35b610166610439565b6040516101739190611297565b60405180910390f35b610196600480360381019061019191906112b2565b610443565b6040516101a3919061126d565b60405180910390f35b6101b4610472565b6040516101c19190611321565b60405180910390f35b6101e460048036038101906101df9190611212565b61047b565b6040516101f1919061126d565b60405180910390f35b610214600480360381019061020f9190611212565b6104b2565b005b610230600480360381019061022b919061133c565b610595565b005b61024c6004803603810190610247919061133c565b6105e1565b6040516102599190611297565b60405180910390f35b61026a610629565b005b61027461063d565b6040516102819190611378565b60405180910390f35b610292610667565b60405161029f9190611157565b60405180910390f35b6102c260048036038101906102bd9190611212565b6106f9565b6040516102cf919061126d565b60405180910390f35b6102f260048036038101906102ed9190611212565b610770565b6040516102ff919061126d565b60405180910390f35b610322600480360381019061031d91906112b2565b610793565b60405161032f919061126d565b60405180910390f35b610352600480360381019061034d9190611393565b610846565b60405161035f9190611297565b60405180910390f35b610382600480360381019061037d919061133c565b6108cd565b005b60606003805461039390611402565b80601f01602080910402602001604051908101604052809291908181526020018280546103bf90611402565b801561040c5780601f106103e15761010080835404028352916020019161040c565b820191906000526020600020905b8154815290600101906020018083116103ef57829003601f168201915b5050505050905090565b600080610421610950565b905061042e818585610958565b600191505092915050565b6000600254905090565b60008061044e610950565b905061045b858285610b21565b610466858585610bad565b60019150509392505050565b60006012905090565b600080610486610950565b90506104a78185856104988589610846565b6104a29190611462565b610958565b600191505092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461050c57600080fd5b61053782610518610472565b60ff16600a61052791906115c9565b836105329190611614565b610e23565b7f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688582610561610472565b60ff16600a61057091906115c9565b8361057b9190611614565b604051610589929190611656565b60405180910390a15050565b61059d610f79565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610631610f79565b61063b6000610ff7565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461067690611402565b80601f01602080910402602001604051908101604052809291908181526020018280546106a290611402565b80156106ef5780601f106106c4576101008083540402835291602001916106ef565b820191906000526020600020905b8154815290600101906020018083116106d257829003601f168201915b5050505050905090565b600080610704610950565b905060006107128286610846565b905083811015610757576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074e906116f1565b60405180910390fd5b6107648286868403610958565b60019250505092915050565b60008061077b610950565b9050610788818585610bad565b600191505092915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107ef57600080fd5b60008490506107ff818585610bad565b7fe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be85858560405161083293929190611711565b60405180910390a160019150509392505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6108d5610f79565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093b906117ba565b60405180910390fd5b61094d81610ff7565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109be9061184c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2d906118de565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610b149190611297565b60405180910390a3505050565b6000610b2d8484610846565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610ba75781811015610b99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b909061194a565b60405180910390fd5b610ba68484848403610958565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c13906119dc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8290611a6e565b60405180910390fd5b610c968383836110bd565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610d1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1390611b00565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e0a9190611297565b60405180910390a3610e1d8484846110c2565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8990611b6c565b60405180910390fd5b610e9e600083836110bd565b8060026000828254610eb09190611462565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f619190611297565b60405180910390a3610f75600083836110c2565b5050565b610f81610950565b73ffffffffffffffffffffffffffffffffffffffff16610f9f61063d565b73ffffffffffffffffffffffffffffffffffffffff1614610ff5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fec90611bd8565b60405180910390fd5b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111015780820151818401526020810190506110e6565b60008484015250505050565b6000601f19601f8301169050919050565b6000611129826110c7565b61113381856110d2565b93506111438185602086016110e3565b61114c8161110d565b840191505092915050565b60006020820190508181036000830152611171818461111e565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006111a98261117e565b9050919050565b6111b98161119e565b81146111c457600080fd5b50565b6000813590506111d6816111b0565b92915050565b6000819050919050565b6111ef816111dc565b81146111fa57600080fd5b50565b60008135905061120c816111e6565b92915050565b6000806040838503121561122957611228611179565b5b6000611237858286016111c7565b9250506020611248858286016111fd565b9150509250929050565b60008115159050919050565b61126781611252565b82525050565b6000602082019050611282600083018461125e565b92915050565b611291816111dc565b82525050565b60006020820190506112ac6000830184611288565b92915050565b6000806000606084860312156112cb576112ca611179565b5b60006112d9868287016111c7565b93505060206112ea868287016111c7565b92505060406112fb868287016111fd565b9150509250925092565b600060ff82169050919050565b61131b81611305565b82525050565b60006020820190506113366000830184611312565b92915050565b60006020828403121561135257611351611179565b5b6000611360848285016111c7565b91505092915050565b6113728161119e565b82525050565b600060208201905061138d6000830184611369565b92915050565b600080604083850312156113aa576113a9611179565b5b60006113b8858286016111c7565b92505060206113c9858286016111c7565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061141a57607f821691505b60208210810361142d5761142c6113d3565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061146d826111dc565b9150611478836111dc565b92508282019050808211156114905761148f611433565b5b92915050565b60008160011c9050919050565b6000808291508390505b60018511156114ed578086048111156114c9576114c8611433565b5b60018516156114d85780820291505b80810290506114e685611496565b94506114ad565b94509492505050565b60008261150657600190506115c2565b8161151457600090506115c2565b816001811461152a576002811461153457611563565b60019150506115c2565b60ff84111561154657611545611433565b5b8360020a91508482111561155d5761155c611433565b5b506115c2565b5060208310610133831016604e8410600b84101617156115985782820a90508381111561159357611592611433565b5b6115c2565b6115a584848460016114a3565b925090508184048111156115bc576115bb611433565b5b81810290505b9392505050565b60006115d4826111dc565b91506115df836111dc565b925061160c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846114f6565b905092915050565b600061161f826111dc565b915061162a836111dc565b9250828202611638816111dc565b9150828204841483151761164f5761164e611433565b5b5092915050565b600060408201905061166b6000830185611369565b6116786020830184611288565b9392505050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006116db6025836110d2565b91506116e68261167f565b604082019050919050565b6000602082019050818103600083015261170a816116ce565b9050919050565b60006060820190506117266000830186611369565b6117336020830185611369565b6117406040830184611288565b949350505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006117a46026836110d2565b91506117af82611748565b604082019050919050565b600060208201905081810360008301526117d381611797565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b60006118366024836110d2565b9150611841826117da565b604082019050919050565b6000602082019050818103600083015261186581611829565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006118c86022836110d2565b91506118d38261186c565b604082019050919050565b600060208201905081810360008301526118f7816118bb565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611934601d836110d2565b915061193f826118fe565b602082019050919050565b6000602082019050818103600083015261196381611927565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006119c66025836110d2565b91506119d18261196a565b604082019050919050565b600060208201905081810360008301526119f5816119b9565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611a586023836110d2565b9150611a63826119fc565b604082019050919050565b60006020820190508181036000830152611a8781611a4b565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611aea6026836110d2565b9150611af582611a8e565b604082019050919050565b60006020820190508181036000830152611b1981611add565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611b56601f836110d2565b9150611b6182611b20565b602082019050919050565b60006020820190508181036000830152611b8581611b49565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611bc26020836110d2565b9150611bcd82611b8c565b602082019050919050565b60006020820190508181036000830152611bf181611bb5565b905091905056fea264697066735822122055993e11cb0350bfbcbd986fdb410922661972cd769a68f77fd27d5e43a4194c64736f6c63430008110033",
}

// DracoABI is the input ABI used to generate the binding from.
// Deprecated: Use DracoMetaData.ABI instead.
var DracoABI = DracoMetaData.ABI

// DracoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DracoMetaData.Bin instead.
var DracoBin = DracoMetaData.Bin

// DeployDraco deploys a new Ethereum contract, binding an instance of Draco to it.
func DeployDraco(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Draco, error) {
	parsed, err := DracoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DracoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Draco{DracoCaller: DracoCaller{contract: contract}, DracoTransactor: DracoTransactor{contract: contract}, DracoFilterer: DracoFilterer{contract: contract}}, nil
}

// Draco is an auto generated Go binding around an Ethereum contract.
type Draco struct {
	DracoCaller     // Read-only binding to the contract
	DracoTransactor // Write-only binding to the contract
	DracoFilterer   // Log filterer for contract events
}

// DracoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DracoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DracoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DracoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DracoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DracoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DracoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DracoSession struct {
	Contract     *Draco            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DracoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DracoCallerSession struct {
	Contract *DracoCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DracoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DracoTransactorSession struct {
	Contract     *DracoTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DracoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DracoRaw struct {
	Contract *Draco // Generic contract binding to access the raw methods on
}

// DracoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DracoCallerRaw struct {
	Contract *DracoCaller // Generic read-only contract binding to access the raw methods on
}

// DracoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DracoTransactorRaw struct {
	Contract *DracoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDraco creates a new instance of Draco, bound to a specific deployed contract.
func NewDraco(address common.Address, backend bind.ContractBackend) (*Draco, error) {
	contract, err := bindDraco(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Draco{DracoCaller: DracoCaller{contract: contract}, DracoTransactor: DracoTransactor{contract: contract}, DracoFilterer: DracoFilterer{contract: contract}}, nil
}

// NewDracoCaller creates a new read-only instance of Draco, bound to a specific deployed contract.
func NewDracoCaller(address common.Address, caller bind.ContractCaller) (*DracoCaller, error) {
	contract, err := bindDraco(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DracoCaller{contract: contract}, nil
}

// NewDracoTransactor creates a new write-only instance of Draco, bound to a specific deployed contract.
func NewDracoTransactor(address common.Address, transactor bind.ContractTransactor) (*DracoTransactor, error) {
	contract, err := bindDraco(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DracoTransactor{contract: contract}, nil
}

// NewDracoFilterer creates a new log filterer instance of Draco, bound to a specific deployed contract.
func NewDracoFilterer(address common.Address, filterer bind.ContractFilterer) (*DracoFilterer, error) {
	contract, err := bindDraco(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DracoFilterer{contract: contract}, nil
}

// bindDraco binds a generic wrapper to an already deployed contract.
func bindDraco(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DracoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Draco *DracoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Draco.Contract.DracoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Draco *DracoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Draco.Contract.DracoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Draco *DracoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Draco.Contract.DracoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Draco *DracoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Draco.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Draco *DracoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Draco.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Draco *DracoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Draco.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Draco *DracoCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Draco.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Draco *DracoSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Draco.Contract.Allowance(&_Draco.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Draco *DracoCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Draco.Contract.Allowance(&_Draco.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Draco *DracoCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Draco.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Draco *DracoSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Draco.Contract.BalanceOf(&_Draco.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Draco *DracoCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Draco.Contract.BalanceOf(&_Draco.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Draco *DracoCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Draco.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Draco *DracoSession) Decimals() (uint8, error) {
	return _Draco.Contract.Decimals(&_Draco.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Draco *DracoCallerSession) Decimals() (uint8, error) {
	return _Draco.Contract.Decimals(&_Draco.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Draco *DracoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Draco.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Draco *DracoSession) Name() (string, error) {
	return _Draco.Contract.Name(&_Draco.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Draco *DracoCallerSession) Name() (string, error) {
	return _Draco.Contract.Name(&_Draco.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Draco *DracoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Draco.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Draco *DracoSession) Owner() (common.Address, error) {
	return _Draco.Contract.Owner(&_Draco.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Draco *DracoCallerSession) Owner() (common.Address, error) {
	return _Draco.Contract.Owner(&_Draco.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Draco *DracoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Draco.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Draco *DracoSession) Symbol() (string, error) {
	return _Draco.Contract.Symbol(&_Draco.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Draco *DracoCallerSession) Symbol() (string, error) {
	return _Draco.Contract.Symbol(&_Draco.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Draco *DracoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Draco.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Draco *DracoSession) TotalSupply() (*big.Int, error) {
	return _Draco.Contract.TotalSupply(&_Draco.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Draco *DracoCallerSession) TotalSupply() (*big.Int, error) {
	return _Draco.Contract.TotalSupply(&_Draco.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Draco *DracoTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Draco *DracoSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Approve(&_Draco.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Draco *DracoTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Approve(&_Draco.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Draco *DracoTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Draco *DracoSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.DecreaseAllowance(&_Draco.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Draco *DracoTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.DecreaseAllowance(&_Draco.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Draco *DracoTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Draco *DracoSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.IncreaseAllowance(&_Draco.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Draco *DracoTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.IncreaseAllowance(&_Draco.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Draco *DracoTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Draco *DracoSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Mint(&_Draco.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Draco *DracoTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Mint(&_Draco.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Draco *DracoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Draco *DracoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Draco.Contract.RenounceOwnership(&_Draco.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Draco *DracoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Draco.Contract.RenounceOwnership(&_Draco.TransactOpts)
}

// SetDex is a paid mutator transaction binding the contract method 0x447fa8b7.
//
// Solidity: function setDex(address dexAddress) returns()
func (_Draco *DracoTransactor) SetDex(opts *bind.TransactOpts, dexAddress common.Address) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "setDex", dexAddress)
}

// SetDex is a paid mutator transaction binding the contract method 0x447fa8b7.
//
// Solidity: function setDex(address dexAddress) returns()
func (_Draco *DracoSession) SetDex(dexAddress common.Address) (*types.Transaction, error) {
	return _Draco.Contract.SetDex(&_Draco.TransactOpts, dexAddress)
}

// SetDex is a paid mutator transaction binding the contract method 0x447fa8b7.
//
// Solidity: function setDex(address dexAddress) returns()
func (_Draco *DracoTransactorSession) SetDex(dexAddress common.Address) (*types.Transaction, error) {
	return _Draco.Contract.SetDex(&_Draco.TransactOpts, dexAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Draco *DracoTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Draco *DracoSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Transfer(&_Draco.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Draco *DracoTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Transfer(&_Draco.TransactOpts, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_Draco *DracoTransactor) Transfer0(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "transfer0", from, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_Draco *DracoSession) Transfer0(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Transfer0(&_Draco.TransactOpts, from, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns(bool)
func (_Draco *DracoTransactorSession) Transfer0(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.Transfer0(&_Draco.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Draco *DracoTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Draco *DracoSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.TransferFrom(&_Draco.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Draco *DracoTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Draco.Contract.TransferFrom(&_Draco.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Draco *DracoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Draco.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Draco *DracoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Draco.Contract.TransferOwnership(&_Draco.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Draco *DracoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Draco.Contract.TransferOwnership(&_Draco.TransactOpts, newOwner)
}

// DracoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Draco contract.
type DracoApprovalIterator struct {
	Event *DracoApproval // Event containing the contract specifics and raw log

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
func (it *DracoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DracoApproval)
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
		it.Event = new(DracoApproval)
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
func (it *DracoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DracoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DracoApproval represents a Approval event raised by the Draco contract.
type DracoApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Draco *DracoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DracoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Draco.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DracoApprovalIterator{contract: _Draco.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Draco *DracoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DracoApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Draco.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DracoApproval)
				if err := _Draco.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Draco *DracoFilterer) ParseApproval(log types.Log) (*DracoApproval, error) {
	event := new(DracoApproval)
	if err := _Draco.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DracoCustomTransferIterator is returned from FilterCustomTransfer and is used to iterate over the raw logs and unpacked data for CustomTransfer events raised by the Draco contract.
type DracoCustomTransferIterator struct {
	Event *DracoCustomTransfer // Event containing the contract specifics and raw log

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
func (it *DracoCustomTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DracoCustomTransfer)
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
		it.Event = new(DracoCustomTransfer)
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
func (it *DracoCustomTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DracoCustomTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DracoCustomTransfer represents a CustomTransfer event raised by the Draco contract.
type DracoCustomTransfer struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCustomTransfer is a free log retrieval operation binding the contract event 0xe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be.
//
// Solidity: event CustomTransfer(address arg0, address arg1, uint256 arg2)
func (_Draco *DracoFilterer) FilterCustomTransfer(opts *bind.FilterOpts) (*DracoCustomTransferIterator, error) {

	logs, sub, err := _Draco.contract.FilterLogs(opts, "CustomTransfer")
	if err != nil {
		return nil, err
	}
	return &DracoCustomTransferIterator{contract: _Draco.contract, event: "CustomTransfer", logs: logs, sub: sub}, nil
}

// WatchCustomTransfer is a free log subscription operation binding the contract event 0xe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be.
//
// Solidity: event CustomTransfer(address arg0, address arg1, uint256 arg2)
func (_Draco *DracoFilterer) WatchCustomTransfer(opts *bind.WatchOpts, sink chan<- *DracoCustomTransfer) (event.Subscription, error) {

	logs, sub, err := _Draco.contract.WatchLogs(opts, "CustomTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DracoCustomTransfer)
				if err := _Draco.contract.UnpackLog(event, "CustomTransfer", log); err != nil {
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
func (_Draco *DracoFilterer) ParseCustomTransfer(log types.Log) (*DracoCustomTransfer, error) {
	event := new(DracoCustomTransfer)
	if err := _Draco.contract.UnpackLog(event, "CustomTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DracoMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Draco contract.
type DracoMintIterator struct {
	Event *DracoMint // Event containing the contract specifics and raw log

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
func (it *DracoMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DracoMint)
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
		it.Event = new(DracoMint)
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
func (it *DracoMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DracoMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DracoMint represents a Mint event raised by the Draco contract.
type DracoMint struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address arg0, uint256 arg1)
func (_Draco *DracoFilterer) FilterMint(opts *bind.FilterOpts) (*DracoMintIterator, error) {

	logs, sub, err := _Draco.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &DracoMintIterator{contract: _Draco.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address arg0, uint256 arg1)
func (_Draco *DracoFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DracoMint) (event.Subscription, error) {

	logs, sub, err := _Draco.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DracoMint)
				if err := _Draco.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_Draco *DracoFilterer) ParseMint(log types.Log) (*DracoMint, error) {
	event := new(DracoMint)
	if err := _Draco.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DracoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Draco contract.
type DracoOwnershipTransferredIterator struct {
	Event *DracoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DracoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DracoOwnershipTransferred)
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
		it.Event = new(DracoOwnershipTransferred)
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
func (it *DracoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DracoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DracoOwnershipTransferred represents a OwnershipTransferred event raised by the Draco contract.
type DracoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Draco *DracoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DracoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Draco.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DracoOwnershipTransferredIterator{contract: _Draco.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Draco *DracoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DracoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Draco.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DracoOwnershipTransferred)
				if err := _Draco.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Draco *DracoFilterer) ParseOwnershipTransferred(log types.Log) (*DracoOwnershipTransferred, error) {
	event := new(DracoOwnershipTransferred)
	if err := _Draco.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DracoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Draco contract.
type DracoTransferIterator struct {
	Event *DracoTransfer // Event containing the contract specifics and raw log

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
func (it *DracoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DracoTransfer)
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
		it.Event = new(DracoTransfer)
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
func (it *DracoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DracoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DracoTransfer represents a Transfer event raised by the Draco contract.
type DracoTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Draco *DracoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DracoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Draco.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DracoTransferIterator{contract: _Draco.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Draco *DracoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DracoTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Draco.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DracoTransfer)
				if err := _Draco.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Draco *DracoFilterer) ParseTransfer(log types.Log) (*DracoTransfer, error) {
	event := new(DracoTransfer)
	if err := _Draco.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

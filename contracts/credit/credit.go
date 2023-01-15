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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CustomTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600681526020017f43726564697400000000000000000000000000000000000000000000000000008152506040518060400160405280600681526020017f7057454d4958000000000000000000000000000000000000000000000000000081525081600390816200008f919062000412565b508060049081620000a1919062000412565b505050620000c4620000b8620000ca60201b60201c565b620000d260201b60201c565b620004f9565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200021a57607f821691505b60208210810362000230576200022f620001d2565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200029a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200025b565b620002a686836200025b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620002f3620002ed620002e784620002be565b620002c8565b620002be565b9050919050565b6000819050919050565b6200030f83620002d2565b620003276200031e82620002fa565b84845462000268565b825550505050565b600090565b6200033e6200032f565b6200034b81848462000304565b505050565b5b8181101562000373576200036760008262000334565b60018101905062000351565b5050565b601f821115620003c2576200038c8162000236565b62000397846200024b565b81016020851015620003a7578190505b620003bf620003b6856200024b565b83018262000350565b50505b505050565b600082821c905092915050565b6000620003e760001984600802620003c7565b1980831691505092915050565b6000620004028383620003d4565b9150826002028217905092915050565b6200041d8262000198565b67ffffffffffffffff811115620004395762000438620001a3565b5b62000445825462000201565b6200045282828562000377565b600060209050601f8311600181146200048a576000841562000475578287015190505b620004818582620003f4565b865550620004f1565b601f1984166200049a8662000236565b60005b82811015620004c4578489015182556001820191506020850194506020810190506200049d565b86831015620004e45784890151620004e0601f891682620003d4565b8355505b6001600288020188555050505b505050505050565b611b0e80620005096000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a611610097578063a9059cbb11610066578063a9059cbb146102b1578063beabacc8146102e1578063dd62ed3e14610311578063f2fde38b1461034157610100565b8063715018a61461023b5780638da5cb5b1461024557806395d89b4114610263578063a457c2d71461028157610100565b8063313ce567116100d3578063313ce567146101a157806339509351146101bf57806340c10f19146101ef57806370a082311461020b57610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d61035d565b60405161011a9190611037565b60405180910390f35b61013d600480360381019061013891906110f2565b6103ef565b60405161014a919061114d565b60405180910390f35b61015b610412565b6040516101689190611177565b60405180910390f35b61018b60048036038101906101869190611192565b61041c565b604051610198919061114d565b60405180910390f35b6101a961044b565b6040516101b69190611201565b60405180910390f35b6101d960048036038101906101d491906110f2565b610454565b6040516101e6919061114d565b60405180910390f35b610209600480360381019061020491906110f2565b61048b565b005b6102256004803603810190610220919061121c565b61051c565b6040516102329190611177565b60405180910390f35b610243610564565b005b61024d610578565b60405161025a9190611258565b60405180910390f35b61026b6105a2565b6040516102789190611037565b60405180910390f35b61029b600480360381019061029691906110f2565b610634565b6040516102a8919061114d565b60405180910390f35b6102cb60048036038101906102c691906110f2565b6106ab565b6040516102d8919061114d565b60405180910390f35b6102fb60048036038101906102f69190611192565b6106ce565b604051610308919061114d565b60405180910390f35b61032b60048036038101906103269190611273565b610726565b6040516103389190611177565b60405180910390f35b61035b6004803603810190610356919061121c565b6107ad565b005b60606003805461036c906112e2565b80601f0160208091040260200160405190810160405280929190818152602001828054610398906112e2565b80156103e55780601f106103ba576101008083540402835291602001916103e5565b820191906000526020600020905b8154815290600101906020018083116103c857829003601f168201915b5050505050905090565b6000806103fa610830565b9050610407818585610838565b600191505092915050565b6000600254905090565b600080610427610830565b9050610434858285610a01565b61043f858585610a8d565b60019150509392505050565b60006012905090565b60008061045f610830565b90506104808185856104718589610726565b61047b9190611342565b610838565b600191505092915050565b610493610d03565b6104be8261049f61044b565b60ff16600a6104ae91906114a9565b836104b991906114f4565b610d81565b7f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826104e861044b565b60ff16600a6104f791906114a9565b8361050291906114f4565b604051610510929190611536565b60405180910390a15050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61056c610d03565b6105766000610ed7565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546105b1906112e2565b80601f01602080910402602001604051908101604052809291908181526020018280546105dd906112e2565b801561062a5780601f106105ff5761010080835404028352916020019161062a565b820191906000526020600020905b81548152906001019060200180831161060d57829003601f168201915b5050505050905090565b60008061063f610830565b9050600061064d8286610726565b905083811015610692576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610689906115d1565b60405180910390fd5b61069f8286868403610838565b60019250505092915050565b6000806106b6610830565b90506106c3818585610a8d565b600191505092915050565b6000808490506106df818585610a8d565b7fe5e0aab7feb98b539542ebe6f8ae203f58e3ca9648e12808e9298825debc15be858585604051610712939291906115f1565b60405180910390a160019150509392505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6107b5610d03565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610824576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081b9061169a565b60405180910390fd5b61082d81610ed7565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089e9061172c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610916576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090d906117be565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516109f49190611177565b60405180910390a3505050565b6000610a0d8484610726565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610a875781811015610a79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a709061182a565b60405180910390fd5b610a868484848403610838565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610afc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af3906118bc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b629061194e565b60405180910390fd5b610b76838383610f9d565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610bfc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf3906119e0565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610cea9190611177565b60405180910390a3610cfd848484610fa2565b50505050565b610d0b610830565b73ffffffffffffffffffffffffffffffffffffffff16610d29610578565b73ffffffffffffffffffffffffffffffffffffffff1614610d7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7690611a4c565b60405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610df0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de790611ab8565b60405180910390fd5b610dfc60008383610f9d565b8060026000828254610e0e9190611342565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610ebf9190611177565b60405180910390a3610ed360008383610fa2565b5050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610fe1578082015181840152602081019050610fc6565b60008484015250505050565b6000601f19601f8301169050919050565b600061100982610fa7565b6110138185610fb2565b9350611023818560208601610fc3565b61102c81610fed565b840191505092915050565b600060208201905081810360008301526110518184610ffe565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006110898261105e565b9050919050565b6110998161107e565b81146110a457600080fd5b50565b6000813590506110b681611090565b92915050565b6000819050919050565b6110cf816110bc565b81146110da57600080fd5b50565b6000813590506110ec816110c6565b92915050565b6000806040838503121561110957611108611059565b5b6000611117858286016110a7565b9250506020611128858286016110dd565b9150509250929050565b60008115159050919050565b61114781611132565b82525050565b6000602082019050611162600083018461113e565b92915050565b611171816110bc565b82525050565b600060208201905061118c6000830184611168565b92915050565b6000806000606084860312156111ab576111aa611059565b5b60006111b9868287016110a7565b93505060206111ca868287016110a7565b92505060406111db868287016110dd565b9150509250925092565b600060ff82169050919050565b6111fb816111e5565b82525050565b600060208201905061121660008301846111f2565b92915050565b60006020828403121561123257611231611059565b5b6000611240848285016110a7565b91505092915050565b6112528161107e565b82525050565b600060208201905061126d6000830184611249565b92915050565b6000806040838503121561128a57611289611059565b5b6000611298858286016110a7565b92505060206112a9858286016110a7565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806112fa57607f821691505b60208210810361130d5761130c6112b3565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061134d826110bc565b9150611358836110bc565b92508282019050808211156113705761136f611313565b5b92915050565b60008160011c9050919050565b6000808291508390505b60018511156113cd578086048111156113a9576113a8611313565b5b60018516156113b85780820291505b80810290506113c685611376565b945061138d565b94509492505050565b6000826113e657600190506114a2565b816113f457600090506114a2565b816001811461140a576002811461141457611443565b60019150506114a2565b60ff84111561142657611425611313565b5b8360020a91508482111561143d5761143c611313565b5b506114a2565b5060208310610133831016604e8410600b84101617156114785782820a90508381111561147357611472611313565b5b6114a2565b6114858484846001611383565b9250905081840481111561149c5761149b611313565b5b81810290505b9392505050565b60006114b4826110bc565b91506114bf836110bc565b92506114ec7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846113d6565b905092915050565b60006114ff826110bc565b915061150a836110bc565b9250828202611518816110bc565b9150828204841483151761152f5761152e611313565b5b5092915050565b600060408201905061154b6000830185611249565b6115586020830184611168565b9392505050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006115bb602583610fb2565b91506115c68261155f565b604082019050919050565b600060208201905081810360008301526115ea816115ae565b9050919050565b60006060820190506116066000830186611249565b6116136020830185611249565b6116206040830184611168565b949350505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611684602683610fb2565b915061168f82611628565b604082019050919050565b600060208201905081810360008301526116b381611677565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611716602483610fb2565b9150611721826116ba565b604082019050919050565b6000602082019050818103600083015261174581611709565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006117a8602283610fb2565b91506117b38261174c565b604082019050919050565b600060208201905081810360008301526117d78161179b565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611814601d83610fb2565b915061181f826117de565b602082019050919050565b6000602082019050818103600083015261184381611807565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006118a6602583610fb2565b91506118b18261184a565b604082019050919050565b600060208201905081810360008301526118d581611899565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611938602383610fb2565b9150611943826118dc565b604082019050919050565b600060208201905081810360008301526119678161192b565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006119ca602683610fb2565b91506119d58261196e565b604082019050919050565b600060208201905081810360008301526119f9816119bd565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611a36602083610fb2565b9150611a4182611a00565b602082019050919050565b60006020820190508181036000830152611a6581611a29565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611aa2601f83610fb2565b9150611aad82611a6c565b602082019050919050565b60006020820190508181036000830152611ad181611a95565b905091905056fea2646970667358221220dda58b21447e56df3cef8ffe833b4e5b8899a4c769382cdc3a8ba049cd037adf64736f6c63430008110033",
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

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Credit *CreditTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Credit *CreditSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Mint(&_Credit.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Credit *CreditTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Mint(&_Credit.TransactOpts, to, amount)
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
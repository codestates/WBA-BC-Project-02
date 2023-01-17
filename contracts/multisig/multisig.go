// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisig

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

// MultisigMetaData contains all meta data concerning the Multisig contract.
var MultisigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_numConfirmationsRequired\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ExecuteTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"RevokeConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SubmitTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"getTransaction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfirmationsRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620026d8380380620026d883398181016040528101906200003791906200050d565b60008251116200007e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007590620005d4565b60405180910390fd5b60008111801562000090575081518111155b620000d2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000c9906200066c565b60405180910390fd5b60005b8251811015620002d4576000838281518110620000f757620000f66200068e565b5b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160362000173576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200016a906200070d565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161562000203576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001fa906200077f565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508080620002cb90620007d0565b915050620000d5565b508060028190555050506200081d565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200034882620002fd565b810181811067ffffffffffffffff821117156200036a57620003696200030e565b5b80604052505050565b60006200037f620002e4565b90506200038d82826200033d565b919050565b600067ffffffffffffffff821115620003b057620003af6200030e565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003f382620003c6565b9050919050565b6200040581620003e6565b81146200041157600080fd5b50565b6000815190506200042581620003fa565b92915050565b6000620004426200043c8462000392565b62000373565b90508083825260208201905060208402830185811115620004685762000467620003c1565b5b835b8181101562000495578062000480888262000414565b8452602084019350506020810190506200046a565b5050509392505050565b600082601f830112620004b757620004b6620002f8565b5b8151620004c98482602086016200042b565b91505092915050565b6000819050919050565b620004e781620004d2565b8114620004f357600080fd5b50565b6000815190506200050781620004dc565b92915050565b60008060408385031215620005275762000526620002ee565b5b600083015167ffffffffffffffff811115620005485762000547620002f3565b5b62000556858286016200049f565b92505060206200056985828601620004f6565b9150509250929050565b600082825260208201905092915050565b7f6f776e6572732072657175697265640000000000000000000000000000000000600082015250565b6000620005bc600f8362000573565b9150620005c98262000584565b602082019050919050565b60006020820190508181036000830152620005ef81620005ad565b9050919050565b7f696e76616c6964206e756d626572206f6620726571756972656420636f6e666960008201527f726d6174696f6e73000000000000000000000000000000000000000000000000602082015250565b60006200065460288362000573565b91506200066182620005f6565b604082019050919050565b60006020820190508181036000830152620006878162000645565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f696e76616c6964206f776e657200000000000000000000000000000000000000600082015250565b6000620006f5600d8362000573565b91506200070282620006bd565b602082019050919050565b600060208201905081810360008301526200072881620006e6565b9050919050565b7f6f776e6572206e6f7420756e6971756500000000000000000000000000000000600082015250565b60006200076760108362000573565b915062000774826200072f565b602082019050919050565b600060208201905081810360008301526200079a8162000758565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620007dd82620004d2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620008125762000811620007a1565b5b600182019050919050565b611eab806200082d6000396000f3fe6080604052600436106100ab5760003560e01c806380f59a651161006457806380f59a65146102535780639ace38c214610290578063a0e67e2b146102d1578063c01a8c84146102fc578063d0549b8514610325578063ee22610b1461035057610102565b8063025e7c27146101075780630d59b5641461014457806320ea8d86146101815780632e7700f0146101aa5780632f54bf6e146101d557806333ea3dc81461021257610102565b36610102573373ffffffffffffffffffffffffffffffffffffffff167f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a1534476040516100f892919061118b565b60405180910390a2005b600080fd5b34801561011357600080fd5b5061012e600480360381019061012991906111f4565b610379565b60405161013b9190611262565b60405180910390f35b34801561015057600080fd5b5061016b600480360381019061016691906113ef565b6103b8565b6040516101789190611472565b60405180910390f35b34801561018d57600080fd5b506101a860048036038101906101a391906111f4565b6105c5565b005b3480156101b657600080fd5b506101bf61089f565b6040516101cc9190611472565b60405180910390f35b3480156101e157600080fd5b506101fc60048036038101906101f7919061148d565b6108ac565b60405161020991906114d5565b60405180910390f35b34801561021e57600080fd5b50610239600480360381019061023491906111f4565b6108cc565b60405161024a95949392919061156f565b60405180910390f35b34801561025f57600080fd5b5061027a600480360381019061027591906115c9565b6109df565b60405161028791906114d5565b60405180910390f35b34801561029c57600080fd5b506102b760048036038101906102b291906111f4565b610a0e565b6040516102c895949392919061156f565b60405180910390f35b3480156102dd57600080fd5b506102e6610b09565b6040516102f391906116c7565b60405180910390f35b34801561030857600080fd5b50610323600480360381019061031e91906111f4565b610b97565b005b34801561033157600080fd5b5061033a610e74565b6040516103479190611472565b60405180910390f35b34801561035c57600080fd5b50610377600480360381019061037291906111f4565b610e7a565b005b6000818154811061038957600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610446576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043d90611746565b60405180910390fd5b6000600480549050905060046040518060a001604052808873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018681526020016000151581526020016000815250908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020190816105229190611972565b5060608201518160030160006101000a81548160ff0219169083151502179055506080820151816004015550508573ffffffffffffffffffffffffffffffffffffffff16813373ffffffffffffffffffffffffffffffffffffffff167fac8d3dc7760c97f73680644024f53b5129d6e052a3838e83a4e019708817200c8888886040516105b193929190611a44565b60405180910390a480915050949350505050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610651576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064890611746565b60405180910390fd5b806004805490508110610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069090611ace565b60405180910390fd5b81600481815481106106ae576106ad611aee565b5b906000526020600020906005020160030160009054906101000a900460ff161561070d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070490611b69565b60405180910390fd5b60006004848154811061072357610722611aee565b5b906000526020600020906005020190506003600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c790611bd5565b60405180910390fd5b60018160040160008282546107e59190611c24565b9250508190555060006003600086815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550833373ffffffffffffffffffffffffffffffffffffffff167ff0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd5560405160405180910390a350505050565b6000600480549050905090565b60016020528060005260406000206000915054906101000a900460ff1681565b60008060606000806000600487815481106108ea576108e9611aee565b5b906000526020600020906005020190508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010154826002018360030160009054906101000a900460ff16846004015482805461094b90611795565b80601f016020809104026020016040519081016040528092919081815260200182805461097790611795565b80156109c45780601f10610999576101008083540402835291602001916109c4565b820191906000526020600020905b8154815290600101906020018083116109a757829003601f168201915b50505050509250955095509550955095505091939590929450565b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60048181548110610a1e57600080fd5b90600052602060002090600502016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054610a6d90611795565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9990611795565b8015610ae65780601f10610abb57610100808354040283529160200191610ae6565b820191906000526020600020905b815481529060010190602001808311610ac957829003601f168201915b5050505050908060030160009054906101000a900460ff16908060040154905085565b60606000805480602002602001604051908101604052809291908181526020018280548015610b8d57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610b43575b5050505050905090565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1a90611746565b60405180910390fd5b806004805490508110610c6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6290611ace565b60405180910390fd5b8160048181548110610c8057610c7f611aee565b5b906000526020600020906005020160030160009054906101000a900460ff1615610cdf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd690611b69565b60405180910390fd5b826003600082815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610d7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7590611ca4565b60405180910390fd5b600060048581548110610d9457610d93611aee565b5b906000526020600020906005020190506001816004016000828254610db99190611cc4565b9250508190555060016003600087815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550843373ffffffffffffffffffffffffffffffffffffffff167f5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb6339004160405160405180910390a35050505050565b60025481565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610f06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610efd90611746565b60405180910390fd5b806004805490508110610f4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4590611ace565b60405180910390fd5b8160048181548110610f6357610f62611aee565b5b906000526020600020906005020160030160009054906101000a900460ff1615610fc2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb990611b69565b60405180910390fd5b600060048481548110610fd857610fd7611aee565b5b9060005260206000209060050201905060025481600401541015611031576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102890611d44565b60405180910390fd5b60018160030160006101000a81548160ff02191690831515021790555060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168260010154836002016040516110a19190611df2565b60006040518083038185875af1925050503d80600081146110de576040519150601f19603f3d011682016040523d82523d6000602084013e6110e3565b606091505b5050905080611127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111e90611e55565b60405180910390fd5b843373ffffffffffffffffffffffffffffffffffffffff167f5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac60405160405180910390a35050505050565b6000819050919050565b61118581611172565b82525050565b60006040820190506111a0600083018561117c565b6111ad602083018461117c565b9392505050565b6000604051905090565b600080fd5b600080fd5b6111d181611172565b81146111dc57600080fd5b50565b6000813590506111ee816111c8565b92915050565b60006020828403121561120a576112096111be565b5b6000611218848285016111df565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061124c82611221565b9050919050565b61125c81611241565b82525050565b60006020820190506112776000830184611253565b92915050565b61128681611241565b811461129157600080fd5b50565b6000813590506112a38161127d565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6112fc826112b3565b810181811067ffffffffffffffff8211171561131b5761131a6112c4565b5b80604052505050565b600061132e6111b4565b905061133a82826112f3565b919050565b600067ffffffffffffffff82111561135a576113596112c4565b5b611363826112b3565b9050602081019050919050565b82818337600083830152505050565b600061139261138d8461133f565b611324565b9050828152602081018484840111156113ae576113ad6112ae565b5b6113b9848285611370565b509392505050565b600082601f8301126113d6576113d56112a9565b5b81356113e684826020860161137f565b91505092915050565b60008060008060808587031215611409576114086111be565b5b600061141787828801611294565b9450506020611428878288016111df565b935050604085013567ffffffffffffffff811115611449576114486111c3565b5b611455878288016113c1565b9250506060611466878288016111df565b91505092959194509250565b6000602082019050611487600083018461117c565b92915050565b6000602082840312156114a3576114a26111be565b5b60006114b184828501611294565b91505092915050565b60008115159050919050565b6114cf816114ba565b82525050565b60006020820190506114ea60008301846114c6565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561152a57808201518184015260208101905061150f565b60008484015250505050565b6000611541826114f0565b61154b81856114fb565b935061155b81856020860161150c565b611564816112b3565b840191505092915050565b600060a0820190506115846000830188611253565b611591602083018761117c565b81810360408301526115a38186611536565b90506115b260608301856114c6565b6115bf608083018461117c565b9695505050505050565b600080604083850312156115e0576115df6111be565b5b60006115ee858286016111df565b92505060206115ff85828601611294565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61163e81611241565b82525050565b60006116508383611635565b60208301905092915050565b6000602082019050919050565b600061167482611609565b61167e8185611614565b935061168983611625565b8060005b838110156116ba5781516116a18882611644565b97506116ac8361165c565b92505060018101905061168d565b5085935050505092915050565b600060208201905081810360008301526116e18184611669565b905092915050565b600082825260208201905092915050565b7f6e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b60006117306009836116e9565b915061173b826116fa565b602082019050919050565b6000602082019050818103600083015261175f81611723565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806117ad57607f821691505b6020821081036117c0576117bf611766565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118287fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826117eb565b61183286836117eb565b95508019841693508086168417925050509392505050565b6000819050919050565b600061186f61186a61186584611172565b61184a565b611172565b9050919050565b6000819050919050565b61188983611854565b61189d61189582611876565b8484546117f8565b825550505050565b600090565b6118b26118a5565b6118bd818484611880565b505050565b5b818110156118e1576118d66000826118aa565b6001810190506118c3565b5050565b601f821115611926576118f7816117c6565b611900846117db565b8101602085101561190f578190505b61192361191b856117db565b8301826118c2565b50505b505050565b600082821c905092915050565b60006119496000198460080261192b565b1980831691505092915050565b60006119628383611938565b9150826002028217905092915050565b61197b826114f0565b67ffffffffffffffff811115611994576119936112c4565b5b61199e8254611795565b6119a98282856118e5565b600060209050601f8311600181146119dc57600084156119ca578287015190505b6119d48582611956565b865550611a3c565b601f1984166119ea866117c6565b60005b82811015611a12578489015182556001820191506020850194506020810190506119ed565b86831015611a2f5784890151611a2b601f891682611938565b8355505b6001600288020188555050505b505050505050565b6000606082019050611a59600083018661117c565b8181036020830152611a6b8185611536565b9050611a7a604083018461117c565b949350505050565b7f747820646f6573206e6f74206578697374000000000000000000000000000000600082015250565b6000611ab86011836116e9565b9150611ac382611a82565b602082019050919050565b60006020820190508181036000830152611ae781611aab565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f747820616c726561647920657865637574656400000000000000000000000000600082015250565b6000611b536013836116e9565b9150611b5e82611b1d565b602082019050919050565b60006020820190508181036000830152611b8281611b46565b9050919050565b7f7478206e6f7420636f6e6669726d656400000000000000000000000000000000600082015250565b6000611bbf6010836116e9565b9150611bca82611b89565b602082019050919050565b60006020820190508181036000830152611bee81611bb2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611c2f82611172565b9150611c3a83611172565b9250828203905081811115611c5257611c51611bf5565b5b92915050565b7f747820616c726561647920636f6e6669726d6564000000000000000000000000600082015250565b6000611c8e6014836116e9565b9150611c9982611c58565b602082019050919050565b60006020820190508181036000830152611cbd81611c81565b9050919050565b6000611ccf82611172565b9150611cda83611172565b9250828201905080821115611cf257611cf1611bf5565b5b92915050565b7f63616e6e6f742065786563757465207478000000000000000000000000000000600082015250565b6000611d2e6011836116e9565b9150611d3982611cf8565b602082019050919050565b60006020820190508181036000830152611d5d81611d21565b9050919050565b600081905092915050565b60008154611d7c81611795565b611d868186611d64565b94506001821660008114611da15760018114611db657611de9565b60ff1983168652811515820286019350611de9565b611dbf856117c6565b60005b83811015611de157815481890152600182019150602081019050611dc2565b838801955050505b50505092915050565b6000611dfe8284611d6f565b915081905092915050565b7f7478206661696c65640000000000000000000000000000000000000000000000600082015250565b6000611e3f6009836116e9565b9150611e4a82611e09565b602082019050919050565b60006020820190508181036000830152611e6e81611e32565b905091905056fea26469706673582212208ccd14230d9d8712e8f6ee3e84dafa755a66149fce29536c9eecead73cfc856164736f6c63430008110033",
}

// MultisigABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisigMetaData.ABI instead.
var MultisigABI = MultisigMetaData.ABI

// MultisigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultisigMetaData.Bin instead.
var MultisigBin = MultisigMetaData.Bin

// DeployMultisig deploys a new Ethereum contract, binding an instance of Multisig to it.
func DeployMultisig(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _numConfirmationsRequired *big.Int) (common.Address, *types.Transaction, *Multisig, error) {
	parsed, err := MultisigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultisigBin), backend, _owners, _numConfirmationsRequired)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multisig{MultisigCaller: MultisigCaller{contract: contract}, MultisigTransactor: MultisigTransactor{contract: contract}, MultisigFilterer: MultisigFilterer{contract: contract}}, nil
}

// Multisig is an auto generated Go binding around an Ethereum contract.
type Multisig struct {
	MultisigCaller     // Read-only binding to the contract
	MultisigTransactor // Write-only binding to the contract
	MultisigFilterer   // Log filterer for contract events
}

// MultisigCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigSession struct {
	Contract     *Multisig         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigCallerSession struct {
	Contract *MultisigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MultisigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigTransactorSession struct {
	Contract     *MultisigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MultisigRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigRaw struct {
	Contract *Multisig // Generic contract binding to access the raw methods on
}

// MultisigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigCallerRaw struct {
	Contract *MultisigCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigTransactorRaw struct {
	Contract *MultisigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisig creates a new instance of Multisig, bound to a specific deployed contract.
func NewMultisig(address common.Address, backend bind.ContractBackend) (*Multisig, error) {
	contract, err := bindMultisig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisig{MultisigCaller: MultisigCaller{contract: contract}, MultisigTransactor: MultisigTransactor{contract: contract}, MultisigFilterer: MultisigFilterer{contract: contract}}, nil
}

// NewMultisigCaller creates a new read-only instance of Multisig, bound to a specific deployed contract.
func NewMultisigCaller(address common.Address, caller bind.ContractCaller) (*MultisigCaller, error) {
	contract, err := bindMultisig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigCaller{contract: contract}, nil
}

// NewMultisigTransactor creates a new write-only instance of Multisig, bound to a specific deployed contract.
func NewMultisigTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigTransactor, error) {
	contract, err := bindMultisig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigTransactor{contract: contract}, nil
}

// NewMultisigFilterer creates a new log filterer instance of Multisig, bound to a specific deployed contract.
func NewMultisigFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigFilterer, error) {
	contract, err := bindMultisig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigFilterer{contract: contract}, nil
}

// bindMultisig binds a generic wrapper to an already deployed contract.
func bindMultisig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisig *MultisigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisig.Contract.MultisigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisig *MultisigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.Contract.MultisigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisig *MultisigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisig.Contract.MultisigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisig *MultisigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisig *MultisigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisig *MultisigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisig.Contract.contract.Transact(opts, method, params...)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisig *MultisigCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisig *MultisigSession) GetOwners() ([]common.Address, error) {
	return _Multisig.Contract.GetOwners(&_Multisig.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisig *MultisigCallerSession) GetOwners() ([]common.Address, error) {
	return _Multisig.Contract.GetOwners(&_Multisig.CallOpts)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCaller) GetTransaction(opts *bind.CallOpts, _txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getTransaction", _txIndex)

	outstruct := new(struct {
		To               common.Address
		Value            *big.Int
		Data             []byte
		Executed         bool
		NumConfirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.NumConfirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigSession) GetTransaction(_txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.GetTransaction(&_Multisig.CallOpts, _txIndex)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCallerSession) GetTransaction(_txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.GetTransaction(&_Multisig.CallOpts, _txIndex)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Multisig *MultisigCaller) GetTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getTransactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Multisig *MultisigSession) GetTransactionCount() (*big.Int, error) {
	return _Multisig.Contract.GetTransactionCount(&_Multisig.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Multisig *MultisigCallerSession) GetTransactionCount() (*big.Int, error) {
	return _Multisig.Contract.GetTransactionCount(&_Multisig.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Multisig *MultisigCaller) IsConfirmed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "isConfirmed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Multisig *MultisigSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Multisig.Contract.IsConfirmed(&_Multisig.CallOpts, arg0, arg1)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Multisig *MultisigCallerSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Multisig.Contract.IsConfirmed(&_Multisig.CallOpts, arg0, arg1)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisig *MultisigCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisig *MultisigSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Multisig.Contract.IsOwner(&_Multisig.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisig *MultisigCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Multisig.Contract.IsOwner(&_Multisig.CallOpts, arg0)
}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_Multisig *MultisigCaller) NumConfirmationsRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "numConfirmationsRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_Multisig *MultisigSession) NumConfirmationsRequired() (*big.Int, error) {
	return _Multisig.Contract.NumConfirmationsRequired(&_Multisig.CallOpts)
}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_Multisig *MultisigCallerSession) NumConfirmationsRequired() (*big.Int, error) {
	return _Multisig.Contract.NumConfirmationsRequired(&_Multisig.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisig *MultisigCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisig *MultisigSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Multisig.Contract.Owners(&_Multisig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisig *MultisigCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Multisig.Contract.Owners(&_Multisig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		To               common.Address
		Value            *big.Int
		Data             []byte
		Executed         bool
		NumConfirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.NumConfirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigSession) Transactions(arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.Transactions(&_Multisig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCallerSession) Transactions(arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.Transactions(&_Multisig.CallOpts, arg0)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactor) ConfirmTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "confirmTransaction", _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ConfirmTransaction(&_Multisig.TransactOpts, _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactorSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ConfirmTransaction(&_Multisig.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactor) ExecuteTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "executeTransaction", _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ExecuteTransaction(&_Multisig.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactorSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ExecuteTransaction(&_Multisig.TransactOpts, _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactor) RevokeConfirmation(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "revokeConfirmation", _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_Multisig *MultisigSession) RevokeConfirmation(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.RevokeConfirmation(&_Multisig.TransactOpts, _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactorSession) RevokeConfirmation(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.RevokeConfirmation(&_Multisig.TransactOpts, _txIndex)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0x0d59b564.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data, uint256 nonce) returns(uint256)
func (_Multisig *MultisigTransactor) SubmitTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte, nonce *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "submitTransaction", _to, _value, _data, nonce)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0x0d59b564.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data, uint256 nonce) returns(uint256)
func (_Multisig *MultisigSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte, nonce *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.SubmitTransaction(&_Multisig.TransactOpts, _to, _value, _data, nonce)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0x0d59b564.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data, uint256 nonce) returns(uint256)
func (_Multisig *MultisigTransactorSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte, nonce *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.SubmitTransaction(&_Multisig.TransactOpts, _to, _value, _data, nonce)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigSession) Receive() (*types.Transaction, error) {
	return _Multisig.Contract.Receive(&_Multisig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigTransactorSession) Receive() (*types.Transaction, error) {
	return _Multisig.Contract.Receive(&_Multisig.TransactOpts)
}

// MultisigConfirmTransactionIterator is returned from FilterConfirmTransaction and is used to iterate over the raw logs and unpacked data for ConfirmTransaction events raised by the Multisig contract.
type MultisigConfirmTransactionIterator struct {
	Event *MultisigConfirmTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigConfirmTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigConfirmTransaction)
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
		it.Event = new(MultisigConfirmTransaction)
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
func (it *MultisigConfirmTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigConfirmTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigConfirmTransaction represents a ConfirmTransaction event raised by the Multisig contract.
type MultisigConfirmTransaction struct {
	Owner   common.Address
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransaction is a free log retrieval operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) FilterConfirmTransaction(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int) (*MultisigConfirmTransactionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "ConfirmTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return &MultisigConfirmTransactionIterator{contract: _Multisig.contract, event: "ConfirmTransaction", logs: logs, sub: sub}, nil
}

// WatchConfirmTransaction is a free log subscription operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) WatchConfirmTransaction(opts *bind.WatchOpts, sink chan<- *MultisigConfirmTransaction, owner []common.Address, txIndex []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "ConfirmTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigConfirmTransaction)
				if err := _Multisig.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
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

// ParseConfirmTransaction is a log parse operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) ParseConfirmTransaction(log types.Log) (*MultisigConfirmTransaction, error) {
	event := new(MultisigConfirmTransaction)
	if err := _Multisig.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Multisig contract.
type MultisigDepositIterator struct {
	Event *MultisigDeposit // Event containing the contract specifics and raw log

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
func (it *MultisigDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigDeposit)
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
		it.Event = new(MultisigDeposit)
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
func (it *MultisigDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigDeposit represents a Deposit event raised by the Multisig contract.
type MultisigDeposit struct {
	Sender  common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 balance)
func (_Multisig *MultisigFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultisigDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultisigDepositIterator{contract: _Multisig.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 balance)
func (_Multisig *MultisigFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultisigDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigDeposit)
				if err := _Multisig.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 balance)
func (_Multisig *MultisigFilterer) ParseDeposit(log types.Log) (*MultisigDeposit, error) {
	event := new(MultisigDeposit)
	if err := _Multisig.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigExecuteTransactionIterator is returned from FilterExecuteTransaction and is used to iterate over the raw logs and unpacked data for ExecuteTransaction events raised by the Multisig contract.
type MultisigExecuteTransactionIterator struct {
	Event *MultisigExecuteTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigExecuteTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigExecuteTransaction)
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
		it.Event = new(MultisigExecuteTransaction)
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
func (it *MultisigExecuteTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigExecuteTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigExecuteTransaction represents a ExecuteTransaction event raised by the Multisig contract.
type MultisigExecuteTransaction struct {
	Owner   common.Address
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuteTransaction is a free log retrieval operation binding the contract event 0x5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac.
//
// Solidity: event ExecuteTransaction(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) FilterExecuteTransaction(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int) (*MultisigExecuteTransactionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "ExecuteTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return &MultisigExecuteTransactionIterator{contract: _Multisig.contract, event: "ExecuteTransaction", logs: logs, sub: sub}, nil
}

// WatchExecuteTransaction is a free log subscription operation binding the contract event 0x5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac.
//
// Solidity: event ExecuteTransaction(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) WatchExecuteTransaction(opts *bind.WatchOpts, sink chan<- *MultisigExecuteTransaction, owner []common.Address, txIndex []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "ExecuteTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigExecuteTransaction)
				if err := _Multisig.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
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

// ParseExecuteTransaction is a log parse operation binding the contract event 0x5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac.
//
// Solidity: event ExecuteTransaction(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) ParseExecuteTransaction(log types.Log) (*MultisigExecuteTransaction, error) {
	event := new(MultisigExecuteTransaction)
	if err := _Multisig.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigRevokeConfirmationIterator is returned from FilterRevokeConfirmation and is used to iterate over the raw logs and unpacked data for RevokeConfirmation events raised by the Multisig contract.
type MultisigRevokeConfirmationIterator struct {
	Event *MultisigRevokeConfirmation // Event containing the contract specifics and raw log

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
func (it *MultisigRevokeConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigRevokeConfirmation)
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
		it.Event = new(MultisigRevokeConfirmation)
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
func (it *MultisigRevokeConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigRevokeConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigRevokeConfirmation represents a RevokeConfirmation event raised by the Multisig contract.
type MultisigRevokeConfirmation struct {
	Owner   common.Address
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevokeConfirmation is a free log retrieval operation binding the contract event 0xf0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd55.
//
// Solidity: event RevokeConfirmation(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) FilterRevokeConfirmation(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int) (*MultisigRevokeConfirmationIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "RevokeConfirmation", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return &MultisigRevokeConfirmationIterator{contract: _Multisig.contract, event: "RevokeConfirmation", logs: logs, sub: sub}, nil
}

// WatchRevokeConfirmation is a free log subscription operation binding the contract event 0xf0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd55.
//
// Solidity: event RevokeConfirmation(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) WatchRevokeConfirmation(opts *bind.WatchOpts, sink chan<- *MultisigRevokeConfirmation, owner []common.Address, txIndex []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "RevokeConfirmation", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigRevokeConfirmation)
				if err := _Multisig.contract.UnpackLog(event, "RevokeConfirmation", log); err != nil {
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

// ParseRevokeConfirmation is a log parse operation binding the contract event 0xf0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd55.
//
// Solidity: event RevokeConfirmation(address indexed owner, uint256 indexed txIndex)
func (_Multisig *MultisigFilterer) ParseRevokeConfirmation(log types.Log) (*MultisigRevokeConfirmation, error) {
	event := new(MultisigRevokeConfirmation)
	if err := _Multisig.contract.UnpackLog(event, "RevokeConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigSubmitTransactionIterator is returned from FilterSubmitTransaction and is used to iterate over the raw logs and unpacked data for SubmitTransaction events raised by the Multisig contract.
type MultisigSubmitTransactionIterator struct {
	Event *MultisigSubmitTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigSubmitTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigSubmitTransaction)
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
		it.Event = new(MultisigSubmitTransaction)
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
func (it *MultisigSubmitTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigSubmitTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigSubmitTransaction represents a SubmitTransaction event raised by the Multisig contract.
type MultisigSubmitTransaction struct {
	Owner   common.Address
	TxIndex *big.Int
	To      common.Address
	Value   *big.Int
	Data    []byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitTransaction is a free log retrieval operation binding the contract event 0xac8d3dc7760c97f73680644024f53b5129d6e052a3838e83a4e019708817200c.
//
// Solidity: event SubmitTransaction(address indexed owner, uint256 indexed txIndex, address indexed to, uint256 value, bytes data, uint256 nonce)
func (_Multisig *MultisigFilterer) FilterSubmitTransaction(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int, to []common.Address) (*MultisigSubmitTransactionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "SubmitTransaction", ownerRule, txIndexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultisigSubmitTransactionIterator{contract: _Multisig.contract, event: "SubmitTransaction", logs: logs, sub: sub}, nil
}

// WatchSubmitTransaction is a free log subscription operation binding the contract event 0xac8d3dc7760c97f73680644024f53b5129d6e052a3838e83a4e019708817200c.
//
// Solidity: event SubmitTransaction(address indexed owner, uint256 indexed txIndex, address indexed to, uint256 value, bytes data, uint256 nonce)
func (_Multisig *MultisigFilterer) WatchSubmitTransaction(opts *bind.WatchOpts, sink chan<- *MultisigSubmitTransaction, owner []common.Address, txIndex []*big.Int, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "SubmitTransaction", ownerRule, txIndexRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigSubmitTransaction)
				if err := _Multisig.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
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

// ParseSubmitTransaction is a log parse operation binding the contract event 0xac8d3dc7760c97f73680644024f53b5129d6e052a3838e83a4e019708817200c.
//
// Solidity: event SubmitTransaction(address indexed owner, uint256 indexed txIndex, address indexed to, uint256 value, bytes data, uint256 nonce)
func (_Multisig *MultisigFilterer) ParseSubmitTransaction(log types.Log) (*MultisigSubmitTransaction, error) {
	event := new(MultisigSubmitTransaction)
	if err := _Multisig.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

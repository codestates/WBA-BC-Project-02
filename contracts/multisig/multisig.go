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
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_numConfirmationsRequired\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ExecuteTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"RevokeConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SubmitTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"getTransaction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfirmationsRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200279a3803806200279a83398181016040528101906200003791906200050d565b60008251116200007e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007590620005d4565b60405180910390fd5b60008111801562000090575081518111155b620000d2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000c9906200066c565b60405180910390fd5b60005b8251811015620002d4576000838281518110620000f757620000f66200068e565b5b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160362000173576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200016a906200070d565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161562000203576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001fa906200077f565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508080620002cb90620007d0565b915050620000d5565b508060028190555050506200081d565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200034882620002fd565b810181811067ffffffffffffffff821117156200036a57620003696200030e565b5b80604052505050565b60006200037f620002e4565b90506200038d82826200033d565b919050565b600067ffffffffffffffff821115620003b057620003af6200030e565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003f382620003c6565b9050919050565b6200040581620003e6565b81146200041157600080fd5b50565b6000815190506200042581620003fa565b92915050565b6000620004426200043c8462000392565b62000373565b90508083825260208201905060208402830185811115620004685762000467620003c1565b5b835b8181101562000495578062000480888262000414565b8452602084019350506020810190506200046a565b5050509392505050565b600082601f830112620004b757620004b6620002f8565b5b8151620004c98482602086016200042b565b91505092915050565b6000819050919050565b620004e781620004d2565b8114620004f357600080fd5b50565b6000815190506200050781620004dc565b92915050565b60008060408385031215620005275762000526620002ee565b5b600083015167ffffffffffffffff811115620005485762000547620002f3565b5b62000556858286016200049f565b92505060206200056985828601620004f6565b9150509250929050565b600082825260208201905092915050565b7f6f776e6572732072657175697265640000000000000000000000000000000000600082015250565b6000620005bc600f8362000573565b9150620005c98262000584565b602082019050919050565b60006020820190508181036000830152620005ef81620005ad565b9050919050565b7f696e76616c6964206e756d626572206f6620726571756972656420636f6e666960008201527f726d6174696f6e73000000000000000000000000000000000000000000000000602082015250565b60006200065460288362000573565b91506200066182620005f6565b604082019050919050565b60006020820190508181036000830152620006878162000645565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f696e76616c6964206f776e657200000000000000000000000000000000000000600082015250565b6000620006f5600d8362000573565b91506200070282620006bd565b602082019050919050565b600060208201905081810360008301526200072881620006e6565b9050919050565b7f6f776e6572206e6f7420756e6971756500000000000000000000000000000000600082015250565b60006200076760108362000573565b915062000774826200072f565b602082019050919050565b600060208201905081810360008301526200079a8162000758565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620007dd82620004d2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620008125762000811620007a1565b5b600182019050919050565b611f6d806200082d6000396000f3fe6080604052600436106100c65760003560e01c806333ea3dc81161007f578063a0e67e2b11610059578063a0e67e2b14610303578063c01a8c841461032e578063d0549b8514610357578063ee22610b146103825761011d565b806333ea3dc81461024457806380f59a65146102855780639ace38c2146102c25761011d565b8063025e7c27146101225780630d59b5641461015f57806320ea8d861461019c5780632b68b9c6146101c55780632e7700f0146101dc5780632f54bf6e146102075761011d565b3661011d573373ffffffffffffffffffffffffffffffffffffffff167f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15344760405161011392919061124d565b60405180910390a2005b600080fd5b34801561012e57600080fd5b50610149600480360381019061014491906112b6565b6103ab565b6040516101569190611324565b60405180910390f35b34801561016b57600080fd5b50610186600480360381019061018191906114b1565b6103ea565b6040516101939190611534565b60405180910390f35b3480156101a857600080fd5b506101c360048036038101906101be91906112b6565b6105f7565b005b3480156101d157600080fd5b506101da6108d1565b005b3480156101e857600080fd5b506101f1610961565b6040516101fe9190611534565b60405180910390f35b34801561021357600080fd5b5061022e6004803603810190610229919061154f565b61096e565b60405161023b9190611597565b60405180910390f35b34801561025057600080fd5b5061026b600480360381019061026691906112b6565b61098e565b60405161027c959493929190611631565b60405180910390f35b34801561029157600080fd5b506102ac60048036038101906102a7919061168b565b610aa1565b6040516102b99190611597565b60405180910390f35b3480156102ce57600080fd5b506102e960048036038101906102e491906112b6565b610ad0565b6040516102fa959493929190611631565b60405180910390f35b34801561030f57600080fd5b50610318610bcb565b6040516103259190611789565b60405180910390f35b34801561033a57600080fd5b50610355600480360381019061035091906112b6565b610c59565b005b34801561036357600080fd5b5061036c610f36565b6040516103799190611534565b60405180910390f35b34801561038e57600080fd5b506103a960048036038101906103a491906112b6565b610f3c565b005b600081815481106103bb57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610478576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046f90611808565b60405180910390fd5b6000600480549050905060046040518060a001604052808873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018681526020016000151581526020016000815250908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020190816105549190611a34565b5060608201518160030160006101000a81548160ff0219169083151502179055506080820151816004015550508573ffffffffffffffffffffffffffffffffffffffff16813373ffffffffffffffffffffffffffffffffffffffff167fac8d3dc7760c97f73680644024f53b5129d6e052a3838e83a4e019708817200c8888886040516105e393929190611b06565b60405180910390a480915050949350505050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067a90611808565b60405180910390fd5b8060048054905081106106cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c290611b90565b60405180910390fd5b81600481815481106106e0576106df611bb0565b5b906000526020600020906005020160030160009054906101000a900460ff161561073f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073690611c2b565b60405180910390fd5b60006004848154811061075557610754611bb0565b5b906000526020600020906005020190506003600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610802576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f990611c97565b60405180910390fd5b60018160040160008282546108179190611ce6565b9250508190555060006003600086815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550833373ffffffffffffffffffffffffffffffffffffffff167ff0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd5560405160405180910390a350505050565b60006002815481106108e6576108e5611bb0565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461094857600080fd5b3073ffffffffffffffffffffffffffffffffffffffff16ff5b6000600480549050905090565b60016020528060005260406000206000915054906101000a900460ff1681565b60008060606000806000600487815481106109ac576109ab611bb0565b5b906000526020600020906005020190508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010154826002018360030160009054906101000a900460ff168460040154828054610a0d90611857565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3990611857565b8015610a865780601f10610a5b57610100808354040283529160200191610a86565b820191906000526020600020905b815481529060010190602001808311610a6957829003601f168201915b50505050509250955095509550955095505091939590929450565b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60048181548110610ae057600080fd5b90600052602060002090600502016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054610b2f90611857565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5b90611857565b8015610ba85780601f10610b7d57610100808354040283529160200191610ba8565b820191906000526020600020905b815481529060010190602001808311610b8b57829003601f168201915b5050505050908060030160009054906101000a900460ff16908060040154905085565b60606000805480602002602001604051908101604052809291908181526020018280548015610c4f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610c05575b5050505050905090565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610ce5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cdc90611808565b60405180910390fd5b806004805490508110610d2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2490611b90565b60405180910390fd5b8160048181548110610d4257610d41611bb0565b5b906000526020600020906005020160030160009054906101000a900460ff1615610da1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9890611c2b565b60405180910390fd5b826003600082815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610e40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3790611d66565b60405180910390fd5b600060048581548110610e5657610e55611bb0565b5b906000526020600020906005020190506001816004016000828254610e7b9190611d86565b9250508190555060016003600087815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550843373ffffffffffffffffffffffffffffffffffffffff167f5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb6339004160405160405180910390a35050505050565b60025481565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610fc8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbf90611808565b60405180910390fd5b806004805490508110611010576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100790611b90565b60405180910390fd5b816004818154811061102557611024611bb0565b5b906000526020600020906005020160030160009054906101000a900460ff1615611084576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107b90611c2b565b60405180910390fd5b60006004848154811061109a57611099611bb0565b5b90600052602060002090600502019050600254816004015410156110f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ea90611e06565b60405180910390fd5b60018160030160006101000a81548160ff02191690831515021790555060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168260010154836002016040516111639190611eb4565b60006040518083038185875af1925050503d80600081146111a0576040519150601f19603f3d011682016040523d82523d6000602084013e6111a5565b606091505b50509050806111e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e090611f17565b60405180910390fd5b843373ffffffffffffffffffffffffffffffffffffffff167f5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac60405160405180910390a35050505050565b6000819050919050565b61124781611234565b82525050565b6000604082019050611262600083018561123e565b61126f602083018461123e565b9392505050565b6000604051905090565b600080fd5b600080fd5b61129381611234565b811461129e57600080fd5b50565b6000813590506112b08161128a565b92915050565b6000602082840312156112cc576112cb611280565b5b60006112da848285016112a1565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061130e826112e3565b9050919050565b61131e81611303565b82525050565b60006020820190506113396000830184611315565b92915050565b61134881611303565b811461135357600080fd5b50565b6000813590506113658161133f565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6113be82611375565b810181811067ffffffffffffffff821117156113dd576113dc611386565b5b80604052505050565b60006113f0611276565b90506113fc82826113b5565b919050565b600067ffffffffffffffff82111561141c5761141b611386565b5b61142582611375565b9050602081019050919050565b82818337600083830152505050565b600061145461144f84611401565b6113e6565b9050828152602081018484840111156114705761146f611370565b5b61147b848285611432565b509392505050565b600082601f8301126114985761149761136b565b5b81356114a8848260208601611441565b91505092915050565b600080600080608085870312156114cb576114ca611280565b5b60006114d987828801611356565b94505060206114ea878288016112a1565b935050604085013567ffffffffffffffff81111561150b5761150a611285565b5b61151787828801611483565b9250506060611528878288016112a1565b91505092959194509250565b6000602082019050611549600083018461123e565b92915050565b60006020828403121561156557611564611280565b5b600061157384828501611356565b91505092915050565b60008115159050919050565b6115918161157c565b82525050565b60006020820190506115ac6000830184611588565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115ec5780820151818401526020810190506115d1565b60008484015250505050565b6000611603826115b2565b61160d81856115bd565b935061161d8185602086016115ce565b61162681611375565b840191505092915050565b600060a0820190506116466000830188611315565b611653602083018761123e565b818103604083015261166581866115f8565b90506116746060830185611588565b611681608083018461123e565b9695505050505050565b600080604083850312156116a2576116a1611280565b5b60006116b0858286016112a1565b92505060206116c185828601611356565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61170081611303565b82525050565b600061171283836116f7565b60208301905092915050565b6000602082019050919050565b6000611736826116cb565b61174081856116d6565b935061174b836116e7565b8060005b8381101561177c5781516117638882611706565b975061176e8361171e565b92505060018101905061174f565b5085935050505092915050565b600060208201905081810360008301526117a3818461172b565b905092915050565b600082825260208201905092915050565b7f6e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b60006117f26009836117ab565b91506117fd826117bc565b602082019050919050565b60006020820190508181036000830152611821816117e5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061186f57607f821691505b60208210810361188257611881611828565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118ea7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826118ad565b6118f486836118ad565b95508019841693508086168417925050509392505050565b6000819050919050565b600061193161192c61192784611234565b61190c565b611234565b9050919050565b6000819050919050565b61194b83611916565b61195f61195782611938565b8484546118ba565b825550505050565b600090565b611974611967565b61197f818484611942565b505050565b5b818110156119a35761199860008261196c565b600181019050611985565b5050565b601f8211156119e8576119b981611888565b6119c28461189d565b810160208510156119d1578190505b6119e56119dd8561189d565b830182611984565b50505b505050565b600082821c905092915050565b6000611a0b600019846008026119ed565b1980831691505092915050565b6000611a2483836119fa565b9150826002028217905092915050565b611a3d826115b2565b67ffffffffffffffff811115611a5657611a55611386565b5b611a608254611857565b611a6b8282856119a7565b600060209050601f831160018114611a9e5760008415611a8c578287015190505b611a968582611a18565b865550611afe565b601f198416611aac86611888565b60005b82811015611ad457848901518255600182019150602085019450602081019050611aaf565b86831015611af15784890151611aed601f8916826119fa565b8355505b6001600288020188555050505b505050505050565b6000606082019050611b1b600083018661123e565b8181036020830152611b2d81856115f8565b9050611b3c604083018461123e565b949350505050565b7f747820646f6573206e6f74206578697374000000000000000000000000000000600082015250565b6000611b7a6011836117ab565b9150611b8582611b44565b602082019050919050565b60006020820190508181036000830152611ba981611b6d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f747820616c726561647920657865637574656400000000000000000000000000600082015250565b6000611c156013836117ab565b9150611c2082611bdf565b602082019050919050565b60006020820190508181036000830152611c4481611c08565b9050919050565b7f7478206e6f7420636f6e6669726d656400000000000000000000000000000000600082015250565b6000611c816010836117ab565b9150611c8c82611c4b565b602082019050919050565b60006020820190508181036000830152611cb081611c74565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611cf182611234565b9150611cfc83611234565b9250828203905081811115611d1457611d13611cb7565b5b92915050565b7f747820616c726561647920636f6e6669726d6564000000000000000000000000600082015250565b6000611d506014836117ab565b9150611d5b82611d1a565b602082019050919050565b60006020820190508181036000830152611d7f81611d43565b9050919050565b6000611d9182611234565b9150611d9c83611234565b9250828201905080821115611db457611db3611cb7565b5b92915050565b7f63616e6e6f742065786563757465207478000000000000000000000000000000600082015250565b6000611df06011836117ab565b9150611dfb82611dba565b602082019050919050565b60006020820190508181036000830152611e1f81611de3565b9050919050565b600081905092915050565b60008154611e3e81611857565b611e488186611e26565b94506001821660008114611e635760018114611e7857611eab565b60ff1983168652811515820286019350611eab565b611e8185611888565b60005b83811015611ea357815481890152600182019150602081019050611e84565b838801955050505b50505092915050565b6000611ec08284611e31565b915081905092915050565b7f7478206661696c65640000000000000000000000000000000000000000000000600082015250565b6000611f016009836117ab565b9150611f0c82611ecb565b602082019050919050565b60006020820190508181036000830152611f3081611ef4565b905091905056fea2646970667358221220c16c9c8d0ecc02a3c35f6dfe92223dac2b53bfb39b08c65a0aa9a0e119ce1ff264736f6c63430008110033",
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

// Destruct is a paid mutator transaction binding the contract method 0x2b68b9c6.
//
// Solidity: function destruct() returns()
func (_Multisig *MultisigTransactor) Destruct(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "destruct")
}

// Destruct is a paid mutator transaction binding the contract method 0x2b68b9c6.
//
// Solidity: function destruct() returns()
func (_Multisig *MultisigSession) Destruct() (*types.Transaction, error) {
	return _Multisig.Contract.Destruct(&_Multisig.TransactOpts)
}

// Destruct is a paid mutator transaction binding the contract method 0x2b68b9c6.
//
// Solidity: function destruct() returns()
func (_Multisig *MultisigTransactorSession) Destruct() (*types.Transaction, error) {
	return _Multisig.Contract.Destruct(&_Multisig.TransactOpts)
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

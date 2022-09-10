// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// DocumentMetaData contains all meta data concerning the Document contract.
var DocumentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userInfoHash\",\"type\":\"bytes32\"}],\"name\":\"CreateAccountSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"NumberOfOwnerDocument\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"compareBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareString\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatureA\",\"type\":\"bytes\"}],\"name\":\"createMultipleSignerDocument\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createSingleSignerDocument\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"name\":\"getHashUserInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"getSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"getSinger\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"DocID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatureB\",\"type\":\"bytes\"}],\"name\":\"signMultipleSignerDocument\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"digest\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061266b806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063a7bb580311610097578063c552e61a11610066578063c552e61a1461034a578063cfd6baa21461037a578063ef1bade3146103ab578063fa540801146103db57610100565b8063a7bb580314610288578063a88c6e45146102ba578063b26feeb9146102ea578063b446f3b21461031a57610100565b806397aba7f9116100d357806397aba7f9146101c85780639a77e1c1146101f85780639ae5c2ee14610228578063a23f12281461025857610100565b80631c3c6d2b1461010557806369fc026b1461013657806375b83187146101685780639201de5514610198575b600080fd5b61011f600480360381019061011a91906114dd565b61040b565b60405161012d929190611563565b60405180910390f35b610150600480360381019061014b91906115ea565b610579565b60405161015f93929190611646565b60405180910390f35b610182600480360381019061017d919061171e565b6106f1565b60405161018f9190611796565b60405180910390f35b6101b260048036038101906101ad91906117e7565b61075d565b6040516101bf9190611893565b60405180910390f35b6101e260048036038101906101dd91906118b5565b6108f8565b6040516101ef9190611920565b60405180910390f35b610212600480360381019061020d919061193b565b610967565b60405161021f9190611a63565b60405180910390f35b610242600480360381019061023d9190611a7e565b610ba7565b60405161024f9190611796565b60405180910390f35b610272600480360381019061026d9190611b09565b610c1e565b60405161027f9190611a63565b60405180910390f35b6102a2600480360381019061029d9190611b52565b610cc0565b6040516102b193929190611bb7565b60405180910390f35b6102d460048036038101906102cf919061193b565b610d28565b6040516102e19190611796565b60405180910390f35b61030460048036038101906102ff9190611bee565b610dae565b6040516103119190611cb2565b60405180910390f35b610334600480360381019061032f9190611b09565b610e0d565b6040516103419190611a63565b60405180910390f35b610364600480360381019061035f9190611cd4565b610e3d565b6040516103719190611a63565b60405180910390f35b610394600480360381019061038f9190611d4c565b611014565b6040516103a2929190611da8565b60405180910390f35b6103c560048036038101906103c09190611dd8565b611255565b6040516103d29190611796565b60405180910390f35b6103f560048036038101906103f091906117e7565b6112ae565b6040516104029190611a63565b60405180910390f35b60008060018054905084101561056a573373ffffffffffffffffffffffffffffffffffffffff166001858154811061044657610445611e18565b5b906000526020600020906008020160040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361055e5782600185815481106104a8576104a7611e18565b5b906000526020600020906008020160060190816104c59190612053565b5060018085815481106104db576104da611e18565b5b906000526020600020906008020160070160006101000a81548160ff021916908315150217905550426001858154811061051857610517611e18565b5b9060005260206000209060080201600201819055506001848154811061054157610540611e18565b5b906000526020600020906008020160020154600191509150610572565b60008091509150610572565b600080915091505b9250929050565b60008060003373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16036105c15760008060009250925092506106ea565b6000600180549050905060018081600181540180825580915050039060005260206000209050506000600182815481106105fe576105fd611e18565b5b90600052602060002090600802019050818160000181905550338160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550868160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550858160050190816106ae9190612053565b5042816001018190555060008160070160006101000a81548160ff02191690831515021790555080600001548160010154600194509450945050505b9250925092565b600081518351146107055760009050610757565b816040516020016107169190612161565b604051602081830303815290604052805190602001208360405160200161073d9190612161565b604051602081830303815290604052805190602001201490505b92915050565b606060005b60208160ff161080156107b45750600060f81b838260ff166020811061078b5761078a611e18565b5b1a60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b156107cc5780806107c4906121a7565b915050610762565b60008160ff1667ffffffffffffffff8111156107eb576107ea6113b2565b5b6040519080825280601f01601f19166020018201604052801561081d5781602001600182028036833780820191505090505b509050600091505b60208260ff161080156108775750600060f81b848360ff166020811061084e5761084d611e18565b5b1a60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b156108ee57838260ff166020811061089257610891611e18565b5b1a60f81b818360ff16815181106108ac576108ab611e18565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535081806108e6906121a7565b925050610825565b8092505050919050565b60008060008061090785610cc0565b9250925092506001868285856040516000815260200160405260405161093094939291906121d0565b6020604051602081039080840390855afa158015610952573d6000803e3d6000fd5b50505060206040510351935050505092915050565b6000806000610975866112de565b905060005b600080549050811015610a96576000818154811061099b5761099a611e18565b5b90600052602060002090600602019250818360020154036109f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e890612261565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff168360030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610a83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7a906122cd565b60405180910390fd5b8080610a8e906122ed565b91505061097a565b506000610aa78a8a8a8a8a8a61130e565b905060008080549050905060006001816001815401808255809150500390600052602060002090505060008181548110610ae457610ae3611e18565b5b90600052602060002090600602019350818460010181905550858460030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610b49886112de565b8460020181905550600084600401819055507f2b34a43219ad79d5196b0e8e41e2591a7180f53aad0fcde99f4f8aac1002e66c82604051610b8a9190611a63565b60405180910390a183600201549450505050509695505050505050565b600060606000610bb78685611014565b80925081935050506000610bd2610bcd87610e0d565b6112ae565b90506000610be082856108f8565b90508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16149450505050509392505050565b600080610c2a836112de565b905060005b600080549050811015610cb257610c6b60008281548110610c5357610c52611e18565b5b90600052602060002090600602016002015483611255565b15610c9f5760008181548110610c8457610c83611e18565b5b90600052602060002090600602016001015492505050610cbb565b8080610caa906122ed565b915050610c2f565b6000801b925050505b919050565b60008060006041845114610d09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0090612381565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b600080610d3988888888888861130e565b905060005b600080549050811015610d9d57610d7a60008281548110610d6257610d61611e18565b5b90600052602060002090600602016001015483611255565b15610d8a57600192505050610da4565b8080610d95906122ed565b915050610d3e565b6000925050505b9695505050505050565b6060806000610dbd8685611014565b80925081935050506000610dd0866112ae565b90506000610dde82856108f8565b604051602001610dee91906123e9565b6040516020818303038152906040529050809450505050509392505050565b600081604051602001610e209190612161565b604051602081830303815290604052805190602001209050919050565b600080610e49846112de565b905060005b600080549050811015610fd357610e8a60008281548110610e7257610e71611e18565b5b90600052602060002090600602016002015483611255565b15610fc0576000808281548110610ea457610ea3611e18565b5b90600052602060002090600602016004015490507f6fc5b69e508af84c5302ddb2ca1e6965cac21ddcad708dce220155101a5db90481604051610ee79190612404565b60405180910390a18460008381548110610f0457610f03611e18565b5b90600052602060002090600602016005016000808581548110610f2a57610f29611e18565b5b90600052602060002090600602016004015481526020019081526020016000206000019081610f599190612053565b50600160008381548110610f7057610f6f611e18565b5b906000526020600020906006020160040154610f8c919061241f565b60008381548110610fa057610f9f611e18565b5b90600052602060002090600602016004018190555082935050505061100e565b8080610fcb906122ed565b915050610e4e565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110059061249f565b60405180910390fd5b92915050565b60606000806110488560405160200161102d9190612161565b6040516020818303038152906040528051906020012061075d565b905060005b600080549050811015611212578160405160200161106b9190612161565b604051602081830303815290604052805190602001206000828154811061109557611094611e18565b5b9060005260206000209060060201600201546040516020016110b791906124e0565b60405160208183030381529060405280519060200120036111ff5760008082815481106110e7576110e6611e18565b5b9060005260206000209060060201600501600087815260200190815260200160002060405180602001604052908160008201805461112490611e76565b80601f016020809104026020016040519081016040528092919081815260200182805461115090611e76565b801561119d5780601f106111725761010080835404028352916020019161119d565b820191906000526020600020905b81548152906001019060200180831161118057829003601f168201915b50505050508152505090508060000151600083815481106111c1576111c0611e18565b5b906000526020600020906006020160030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169450945050505061124e565b808061120a906122ed565b91505061104d565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124590612547565b60405180910390fd5b9250929050565b60008160405160200161126891906124e0565b604051602081830303815290604052805190602001208360405160200161128f91906124e0565b6040516020818303038152906040528051906020012014905092915050565b6000816040516020016112c191906125b3565b604051602081830303815290604052805190602001209050919050565b6000816040516020016112f19190612161565b604051602081830303815290604052805190602001209050919050565b600086868686868660405160200161132b969594939291906125d9565b6040516020818303038152906040528051906020012090509695505050505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61137481611361565b811461137f57600080fd5b50565b6000813590506113918161136b565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6113ea826113a1565b810181811067ffffffffffffffff82111715611409576114086113b2565b5b80604052505050565b600061141c61134d565b905061142882826113e1565b919050565b600067ffffffffffffffff821115611448576114476113b2565b5b611451826113a1565b9050602081019050919050565b82818337600083830152505050565b600061148061147b8461142d565b611412565b90508281526020810184848401111561149c5761149b61139c565b5b6114a784828561145e565b509392505050565b600082601f8301126114c4576114c3611397565b5b81356114d484826020860161146d565b91505092915050565b600080604083850312156114f4576114f3611357565b5b600061150285828601611382565b925050602083013567ffffffffffffffff8111156115235761152261135c565b5b61152f858286016114af565b9150509250929050565b61154281611361565b82525050565b60008115159050919050565b61155d81611548565b82525050565b60006040820190506115786000830185611539565b6115856020830184611554565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115b78261158c565b9050919050565b6115c7816115ac565b81146115d257600080fd5b50565b6000813590506115e4816115be565b92915050565b6000806040838503121561160157611600611357565b5b600061160f858286016115d5565b925050602083013567ffffffffffffffff8111156116305761162f61135c565b5b61163c858286016114af565b9150509250929050565b600060608201905061165b6000830186611539565b6116686020830185611539565b6116756040830184611554565b949350505050565b600067ffffffffffffffff821115611698576116976113b2565b5b6116a1826113a1565b9050602081019050919050565b60006116c16116bc8461167d565b611412565b9050828152602081018484840111156116dd576116dc61139c565b5b6116e884828561145e565b509392505050565b600082601f83011261170557611704611397565b5b81356117158482602086016116ae565b91505092915050565b6000806040838503121561173557611734611357565b5b600083013567ffffffffffffffff8111156117535761175261135c565b5b61175f858286016116f0565b925050602083013567ffffffffffffffff8111156117805761177f61135c565b5b61178c858286016116f0565b9150509250929050565b60006020820190506117ab6000830184611554565b92915050565b6000819050919050565b6117c4816117b1565b81146117cf57600080fd5b50565b6000813590506117e1816117bb565b92915050565b6000602082840312156117fd576117fc611357565b5b600061180b848285016117d2565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561184e578082015181840152602081019050611833565b60008484015250505050565b600061186582611814565b61186f818561181f565b935061187f818560208601611830565b611888816113a1565b840191505092915050565b600060208201905081810360008301526118ad818461185a565b905092915050565b600080604083850312156118cc576118cb611357565b5b60006118da858286016117d2565b925050602083013567ffffffffffffffff8111156118fb576118fa61135c565b5b611907858286016114af565b9150509250929050565b61191a816115ac565b82525050565b60006020820190506119356000830184611911565b92915050565b60008060008060008060c0878903121561195857611957611357565b5b600087013567ffffffffffffffff8111156119765761197561135c565b5b61198289828a016116f0565b965050602087013567ffffffffffffffff8111156119a3576119a261135c565b5b6119af89828a016116f0565b955050604087013567ffffffffffffffff8111156119d0576119cf61135c565b5b6119dc89828a016116f0565b945050606087013567ffffffffffffffff8111156119fd576119fc61135c565b5b611a0989828a016116f0565b935050608087013567ffffffffffffffff811115611a2a57611a2961135c565b5b611a3689828a016116f0565b92505060a0611a4789828a016115d5565b9150509295509295509295565b611a5d816117b1565b82525050565b6000602082019050611a786000830184611a54565b92915050565b600080600060608486031215611a9757611a96611357565b5b600084013567ffffffffffffffff811115611ab557611ab461135c565b5b611ac1868287016116f0565b935050602084013567ffffffffffffffff811115611ae257611ae161135c565b5b611aee868287016116f0565b9250506040611aff86828701611382565b9150509250925092565b600060208284031215611b1f57611b1e611357565b5b600082013567ffffffffffffffff811115611b3d57611b3c61135c565b5b611b49848285016116f0565b91505092915050565b600060208284031215611b6857611b67611357565b5b600082013567ffffffffffffffff811115611b8657611b8561135c565b5b611b92848285016114af565b91505092915050565b600060ff82169050919050565b611bb181611b9b565b82525050565b6000606082019050611bcc6000830186611a54565b611bd96020830185611a54565b611be66040830184611ba8565b949350505050565b600080600060608486031215611c0757611c06611357565b5b600084013567ffffffffffffffff811115611c2557611c2461135c565b5b611c31868287016116f0565b9350506020611c42868287016117d2565b9250506040611c5386828701611382565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000611c8482611c5d565b611c8e8185611c68565b9350611c9e818560208601611830565b611ca7816113a1565b840191505092915050565b60006020820190508181036000830152611ccc8184611c79565b905092915050565b60008060408385031215611ceb57611cea611357565b5b600083013567ffffffffffffffff811115611d0957611d0861135c565b5b611d15858286016116f0565b925050602083013567ffffffffffffffff811115611d3657611d3561135c565b5b611d42858286016114af565b9150509250929050565b60008060408385031215611d6357611d62611357565b5b600083013567ffffffffffffffff811115611d8157611d8061135c565b5b611d8d858286016116f0565b9250506020611d9e85828601611382565b9150509250929050565b60006040820190508181036000830152611dc28185611c79565b9050611dd16020830184611911565b9392505050565b60008060408385031215611def57611dee611357565b5b6000611dfd858286016117d2565b9250506020611e0e858286016117d2565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611e8e57607f821691505b602082108103611ea157611ea0611e47565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611f097fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611ecc565b611f138683611ecc565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611f50611f4b611f4684611361565b611f2b565b611361565b9050919050565b6000819050919050565b611f6a83611f35565b611f7e611f7682611f57565b848454611ed9565b825550505050565b600090565b611f93611f86565b611f9e818484611f61565b505050565b5b81811015611fc257611fb7600082611f8b565b600181019050611fa4565b5050565b601f82111561200757611fd881611ea7565b611fe184611ebc565b81016020851015611ff0578190505b612004611ffc85611ebc565b830182611fa3565b50505b505050565b600082821c905092915050565b600061202a6000198460080261200c565b1980831691505092915050565b60006120438383612019565b9150826002028217905092915050565b61205c82611c5d565b67ffffffffffffffff811115612075576120746113b2565b5b61207f8254611e76565b61208a828285611fc6565b600060209050601f8311600181146120bd57600084156120ab578287015190505b6120b58582612037565b86555061211d565b601f1984166120cb86611ea7565b60005b828110156120f3578489015182556001820191506020850194506020810190506120ce565b86831015612110578489015161210c601f891682612019565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b600061213b82611814565b6121458185612125565b9350612155818560208601611830565b80840191505092915050565b600061216d8284612130565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006121b282611b9b565b915060ff82036121c5576121c4612178565b5b600182019050919050565b60006080820190506121e56000830187611a54565b6121f26020830186611ba8565b6121ff6040830185611a54565b61220c6060830184611a54565b95945050505050565b7f5468652070686f6e65206e756d62657220697320657869737465640000000000600082015250565b600061224b601b8361181f565b915061225682612215565b602082019050919050565b6000602082019050818103600083015261227a8161223e565b9050919050565b7f54686520757365722773207075626c6963206b65792069732065786973746564600082015250565b60006122b760208361181f565b91506122c282612281565b602082019050919050565b600060208201905081810360008301526122e6816122aa565b9050919050565b60006122f882611361565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361232a57612329612178565b5b600182019050919050565b7f696e76616c6964207369676e6174757265206c656e6774680000000000000000600082015250565b600061236b60188361181f565b915061237682612335565b602082019050919050565b6000602082019050818103600083015261239a8161235e565b9050919050565b60008160601b9050919050565b60006123b9826123a1565b9050919050565b60006123cb826123ae565b9050919050565b6123e36123de826115ac565b6123c0565b82525050565b60006123f582846123d2565b60148201915081905092915050565b60006020820190506124196000830184611539565b92915050565b600061242a82611361565b915061243583611361565b925082820190508082111561244d5761244c612178565b5b92915050565b7f7369676e206973206e6f74207375636365737300000000000000000000000000600082015250565b600061248960138361181f565b915061249482612453565b602082019050919050565b600060208201905081810360008301526124b88161247c565b9050919050565b6000819050919050565b6124da6124d5826117b1565b6124bf565b82525050565b60006124ec82846124c9565b60208201915081905092915050565b7f4e6f7420666f756e640000000000000000000000000000000000000000000000600082015250565b600061253160098361181f565b915061253c826124fb565b602082019050919050565b6000602082019050818103600083015261256081612524565b9050919050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b600061259d601c83612125565b91506125a882612567565b601c82019050919050565b60006125be82612590565b91506125ca82846124c9565b60208201915081905092915050565b60006125e58289612130565b91506125f18288612130565b91506125fd8287612130565b91506126098286612130565b91506126158285612130565b915061262182846123d2565b60148201915081905097965050505050505056fea2646970667358221220a69398f4185418b891d983c8c68ab3428d2f4627eeb8786b494a8e6e29cded8f64736f6c63430008100033",
}

// DocumentABI is the input ABI used to generate the binding from.
// Deprecated: Use DocumentMetaData.ABI instead.
var DocumentABI = DocumentMetaData.ABI

// DocumentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DocumentMetaData.Bin instead.
var DocumentBin = DocumentMetaData.Bin

// DeployDocument deploys a new Ethereum contract, binding an instance of Document to it.
func DeployDocument(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Document, error) {
	parsed, err := DocumentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DocumentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Document{DocumentCaller: DocumentCaller{contract: contract}, DocumentTransactor: DocumentTransactor{contract: contract}, DocumentFilterer: DocumentFilterer{contract: contract}}, nil
}

// Document is an auto generated Go binding around an Ethereum contract.
type Document struct {
	DocumentCaller     // Read-only binding to the contract
	DocumentTransactor // Write-only binding to the contract
	DocumentFilterer   // Log filterer for contract events
}

// DocumentCaller is an auto generated read-only Go binding around an Ethereum contract.
type DocumentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DocumentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DocumentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DocumentSession struct {
	Contract     *Document         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DocumentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DocumentCallerSession struct {
	Contract *DocumentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DocumentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DocumentTransactorSession struct {
	Contract     *DocumentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DocumentRaw is an auto generated low-level Go binding around an Ethereum contract.
type DocumentRaw struct {
	Contract *Document // Generic contract binding to access the raw methods on
}

// DocumentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DocumentCallerRaw struct {
	Contract *DocumentCaller // Generic read-only contract binding to access the raw methods on
}

// DocumentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DocumentTransactorRaw struct {
	Contract *DocumentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDocument creates a new instance of Document, bound to a specific deployed contract.
func NewDocument(address common.Address, backend bind.ContractBackend) (*Document, error) {
	contract, err := bindDocument(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Document{DocumentCaller: DocumentCaller{contract: contract}, DocumentTransactor: DocumentTransactor{contract: contract}, DocumentFilterer: DocumentFilterer{contract: contract}}, nil
}

// NewDocumentCaller creates a new read-only instance of Document, bound to a specific deployed contract.
func NewDocumentCaller(address common.Address, caller bind.ContractCaller) (*DocumentCaller, error) {
	contract, err := bindDocument(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DocumentCaller{contract: contract}, nil
}

// NewDocumentTransactor creates a new write-only instance of Document, bound to a specific deployed contract.
func NewDocumentTransactor(address common.Address, transactor bind.ContractTransactor) (*DocumentTransactor, error) {
	contract, err := bindDocument(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DocumentTransactor{contract: contract}, nil
}

// NewDocumentFilterer creates a new log filterer instance of Document, bound to a specific deployed contract.
func NewDocumentFilterer(address common.Address, filterer bind.ContractFilterer) (*DocumentFilterer, error) {
	contract, err := bindDocument(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DocumentFilterer{contract: contract}, nil
}

// bindDocument binds a generic wrapper to an already deployed contract.
func bindDocument(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DocumentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Document *DocumentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Document.Contract.DocumentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Document *DocumentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Document.Contract.DocumentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Document *DocumentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Document.Contract.DocumentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Document *DocumentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Document.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Document *DocumentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Document.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Document *DocumentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Document.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Document *DocumentCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Document *DocumentSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Document.Contract.Bytes32ToString(&_Document.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Document *DocumentCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Document.Contract.Bytes32ToString(&_Document.CallOpts, _bytes32)
}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Document *DocumentCaller) CompareBytes(opts *bind.CallOpts, a [32]byte, b [32]byte) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "compareBytes", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Document *DocumentSession) CompareBytes(a [32]byte, b [32]byte) (bool, error) {
	return _Document.Contract.CompareBytes(&_Document.CallOpts, a, b)
}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Document *DocumentCallerSession) CompareBytes(a [32]byte, b [32]byte) (bool, error) {
	return _Document.Contract.CompareBytes(&_Document.CallOpts, a, b)
}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Document *DocumentCaller) CompareString(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "compareString", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Document *DocumentSession) CompareString(a string, b string) (bool, error) {
	return _Document.Contract.CompareString(&_Document.CallOpts, a, b)
}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Document *DocumentCallerSession) CompareString(a string, b string) (bool, error) {
	return _Document.Contract.CompareString(&_Document.CallOpts, a, b)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Document *DocumentCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Document *DocumentSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Document.Contract.GetEthSignedMessageHash(&_Document.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Document *DocumentCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Document.Contract.GetEthSignedMessageHash(&_Document.CallOpts, _messageHash)
}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Document *DocumentCaller) GetHashUserInfo(opts *bind.CallOpts, phone string) ([32]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getHashUserInfo", phone)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Document *DocumentSession) GetHashUserInfo(phone string) ([32]byte, error) {
	return _Document.Contract.GetHashUserInfo(&_Document.CallOpts, phone)
}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Document *DocumentCallerSession) GetHashUserInfo(phone string) ([32]byte, error) {
	return _Document.Contract.GetHashUserInfo(&_Document.CallOpts, phone)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Document *DocumentCaller) GetMessageHash(opts *bind.CallOpts, _message string) ([32]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getMessageHash", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Document *DocumentSession) GetMessageHash(_message string) ([32]byte, error) {
	return _Document.Contract.GetMessageHash(&_Document.CallOpts, _message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Document *DocumentCallerSession) GetMessageHash(_message string) ([32]byte, error) {
	return _Document.Contract.GetMessageHash(&_Document.CallOpts, _message)
}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Document *DocumentCaller) GetSignature(opts *bind.CallOpts, phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getSignature", phone, indexDoc)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Document *DocumentSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	return _Document.Contract.GetSignature(&_Document.CallOpts, phone, indexDoc)
}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Document *DocumentCallerSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	return _Document.Contract.GetSignature(&_Document.CallOpts, phone, indexDoc)
}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentCaller) GetSinger(opts *bind.CallOpts, phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getSinger", phone, digest, indexDoc)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentSession) GetSinger(phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	return _Document.Contract.GetSinger(&_Document.CallOpts, phone, digest, indexDoc)
}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentCallerSession) GetSinger(phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	return _Document.Contract.GetSinger(&_Document.CallOpts, phone, digest, indexDoc)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Document *DocumentCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Document *DocumentSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Document.Contract.RecoverSigner(&_Document.CallOpts, _ethSignedMessageHash, _signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Document *DocumentCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Document.Contract.RecoverSigner(&_Document.CallOpts, _ethSignedMessageHash, _signature)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Document *DocumentCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Document *DocumentSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Document.Contract.SplitSignature(&_Document.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Document *DocumentCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Document.Contract.SplitSignature(&_Document.CallOpts, sig)
}

// VerifyDoc is a free data retrieval call binding the contract method 0x9ae5c2ee.
//
// Solidity: function verifyDoc(string phone, string digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentCaller) VerifyDoc(opts *bind.CallOpts, phone string, digest string, indexDoc *big.Int) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "verifyDoc", phone, digest, indexDoc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDoc is a free data retrieval call binding the contract method 0x9ae5c2ee.
//
// Solidity: function verifyDoc(string phone, string digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentSession) VerifyDoc(phone string, digest string, indexDoc *big.Int) (bool, error) {
	return _Document.Contract.VerifyDoc(&_Document.CallOpts, phone, digest, indexDoc)
}

// VerifyDoc is a free data retrieval call binding the contract method 0x9ae5c2ee.
//
// Solidity: function verifyDoc(string phone, string digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentCallerSession) VerifyDoc(phone string, digest string, indexDoc *big.Int) (bool, error) {
	return _Document.Contract.VerifyDoc(&_Document.CallOpts, phone, digest, indexDoc)
}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Document *DocumentCaller) VerifyUser(opts *bind.CallOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "verifyUser", name, cmnd, dateOB, phone, gmail, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Document *DocumentSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	return _Document.Contract.VerifyUser(&_Document.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Document *DocumentCallerSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	return _Document.Contract.VerifyUser(&_Document.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// CreateMultipleSignerDocument is a paid mutator transaction binding the contract method 0x69fc026b.
//
// Solidity: function createMultipleSignerDocument(address partner, bytes signatureA) returns(uint256, uint256, bool)
func (_Document *DocumentTransactor) CreateMultipleSignerDocument(opts *bind.TransactOpts, partner common.Address, signatureA []byte) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "createMultipleSignerDocument", partner, signatureA)
}

// CreateMultipleSignerDocument is a paid mutator transaction binding the contract method 0x69fc026b.
//
// Solidity: function createMultipleSignerDocument(address partner, bytes signatureA) returns(uint256, uint256, bool)
func (_Document *DocumentSession) CreateMultipleSignerDocument(partner common.Address, signatureA []byte) (*types.Transaction, error) {
	return _Document.Contract.CreateMultipleSignerDocument(&_Document.TransactOpts, partner, signatureA)
}

// CreateMultipleSignerDocument is a paid mutator transaction binding the contract method 0x69fc026b.
//
// Solidity: function createMultipleSignerDocument(address partner, bytes signatureA) returns(uint256, uint256, bool)
func (_Document *DocumentTransactorSession) CreateMultipleSignerDocument(partner common.Address, signatureA []byte) (*types.Transaction, error) {
	return _Document.Contract.CreateMultipleSignerDocument(&_Document.TransactOpts, partner, signatureA)
}

// CreateSingleSignerDocument is a paid mutator transaction binding the contract method 0xc552e61a.
//
// Solidity: function createSingleSignerDocument(string phone, bytes signature) returns(bytes32)
func (_Document *DocumentTransactor) CreateSingleSignerDocument(opts *bind.TransactOpts, phone string, signature []byte) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "createSingleSignerDocument", phone, signature)
}

// CreateSingleSignerDocument is a paid mutator transaction binding the contract method 0xc552e61a.
//
// Solidity: function createSingleSignerDocument(string phone, bytes signature) returns(bytes32)
func (_Document *DocumentSession) CreateSingleSignerDocument(phone string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.CreateSingleSignerDocument(&_Document.TransactOpts, phone, signature)
}

// CreateSingleSignerDocument is a paid mutator transaction binding the contract method 0xc552e61a.
//
// Solidity: function createSingleSignerDocument(string phone, bytes signature) returns(bytes32)
func (_Document *DocumentTransactorSession) CreateSingleSignerDocument(phone string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.CreateSingleSignerDocument(&_Document.TransactOpts, phone, signature)
}

// SignMultipleSignerDocument is a paid mutator transaction binding the contract method 0x1c3c6d2b.
//
// Solidity: function signMultipleSignerDocument(uint256 DocID, bytes signatureB) returns(uint256, bool)
func (_Document *DocumentTransactor) SignMultipleSignerDocument(opts *bind.TransactOpts, DocID *big.Int, signatureB []byte) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "signMultipleSignerDocument", DocID, signatureB)
}

// SignMultipleSignerDocument is a paid mutator transaction binding the contract method 0x1c3c6d2b.
//
// Solidity: function signMultipleSignerDocument(uint256 DocID, bytes signatureB) returns(uint256, bool)
func (_Document *DocumentSession) SignMultipleSignerDocument(DocID *big.Int, signatureB []byte) (*types.Transaction, error) {
	return _Document.Contract.SignMultipleSignerDocument(&_Document.TransactOpts, DocID, signatureB)
}

// SignMultipleSignerDocument is a paid mutator transaction binding the contract method 0x1c3c6d2b.
//
// Solidity: function signMultipleSignerDocument(uint256 DocID, bytes signatureB) returns(uint256, bool)
func (_Document *DocumentTransactorSession) SignMultipleSignerDocument(DocID *big.Int, signatureB []byte) (*types.Transaction, error) {
	return _Document.Contract.SignMultipleSignerDocument(&_Document.TransactOpts, DocID, signatureB)
}

// StoreUser is a paid mutator transaction binding the contract method 0x9a77e1c1.
//
// Solidity: function storeUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Document *DocumentTransactor) StoreUser(opts *bind.TransactOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "storeUser", name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x9a77e1c1.
//
// Solidity: function storeUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Document *DocumentSession) StoreUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.Contract.StoreUser(&_Document.TransactOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x9a77e1c1.
//
// Solidity: function storeUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Document *DocumentTransactorSession) StoreUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.Contract.StoreUser(&_Document.TransactOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// DocumentCreateAccountSuccessIterator is returned from FilterCreateAccountSuccess and is used to iterate over the raw logs and unpacked data for CreateAccountSuccess events raised by the Document contract.
type DocumentCreateAccountSuccessIterator struct {
	Event *DocumentCreateAccountSuccess // Event containing the contract specifics and raw log

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
func (it *DocumentCreateAccountSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DocumentCreateAccountSuccess)
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
		it.Event = new(DocumentCreateAccountSuccess)
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
func (it *DocumentCreateAccountSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DocumentCreateAccountSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DocumentCreateAccountSuccess represents a CreateAccountSuccess event raised by the Document contract.
type DocumentCreateAccountSuccess struct {
	UserInfoHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateAccountSuccess is a free log retrieval operation binding the contract event 0x2b34a43219ad79d5196b0e8e41e2591a7180f53aad0fcde99f4f8aac1002e66c.
//
// Solidity: event CreateAccountSuccess(bytes32 userInfoHash)
func (_Document *DocumentFilterer) FilterCreateAccountSuccess(opts *bind.FilterOpts) (*DocumentCreateAccountSuccessIterator, error) {

	logs, sub, err := _Document.contract.FilterLogs(opts, "CreateAccountSuccess")
	if err != nil {
		return nil, err
	}
	return &DocumentCreateAccountSuccessIterator{contract: _Document.contract, event: "CreateAccountSuccess", logs: logs, sub: sub}, nil
}

// WatchCreateAccountSuccess is a free log subscription operation binding the contract event 0x2b34a43219ad79d5196b0e8e41e2591a7180f53aad0fcde99f4f8aac1002e66c.
//
// Solidity: event CreateAccountSuccess(bytes32 userInfoHash)
func (_Document *DocumentFilterer) WatchCreateAccountSuccess(opts *bind.WatchOpts, sink chan<- *DocumentCreateAccountSuccess) (event.Subscription, error) {

	logs, sub, err := _Document.contract.WatchLogs(opts, "CreateAccountSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DocumentCreateAccountSuccess)
				if err := _Document.contract.UnpackLog(event, "CreateAccountSuccess", log); err != nil {
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

// ParseCreateAccountSuccess is a log parse operation binding the contract event 0x2b34a43219ad79d5196b0e8e41e2591a7180f53aad0fcde99f4f8aac1002e66c.
//
// Solidity: event CreateAccountSuccess(bytes32 userInfoHash)
func (_Document *DocumentFilterer) ParseCreateAccountSuccess(log types.Log) (*DocumentCreateAccountSuccess, error) {
	event := new(DocumentCreateAccountSuccess)
	if err := _Document.contract.UnpackLog(event, "CreateAccountSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DocumentNumberOfOwnerDocumentIterator is returned from FilterNumberOfOwnerDocument and is used to iterate over the raw logs and unpacked data for NumberOfOwnerDocument events raised by the Document contract.
type DocumentNumberOfOwnerDocumentIterator struct {
	Event *DocumentNumberOfOwnerDocument // Event containing the contract specifics and raw log

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
func (it *DocumentNumberOfOwnerDocumentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DocumentNumberOfOwnerDocument)
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
		it.Event = new(DocumentNumberOfOwnerDocument)
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
func (it *DocumentNumberOfOwnerDocumentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DocumentNumberOfOwnerDocumentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DocumentNumberOfOwnerDocument represents a NumberOfOwnerDocument event raised by the Document contract.
type DocumentNumberOfOwnerDocument struct {
	Num *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNumberOfOwnerDocument is a free log retrieval operation binding the contract event 0x6fc5b69e508af84c5302ddb2ca1e6965cac21ddcad708dce220155101a5db904.
//
// Solidity: event NumberOfOwnerDocument(uint256 num)
func (_Document *DocumentFilterer) FilterNumberOfOwnerDocument(opts *bind.FilterOpts) (*DocumentNumberOfOwnerDocumentIterator, error) {

	logs, sub, err := _Document.contract.FilterLogs(opts, "NumberOfOwnerDocument")
	if err != nil {
		return nil, err
	}
	return &DocumentNumberOfOwnerDocumentIterator{contract: _Document.contract, event: "NumberOfOwnerDocument", logs: logs, sub: sub}, nil
}

// WatchNumberOfOwnerDocument is a free log subscription operation binding the contract event 0x6fc5b69e508af84c5302ddb2ca1e6965cac21ddcad708dce220155101a5db904.
//
// Solidity: event NumberOfOwnerDocument(uint256 num)
func (_Document *DocumentFilterer) WatchNumberOfOwnerDocument(opts *bind.WatchOpts, sink chan<- *DocumentNumberOfOwnerDocument) (event.Subscription, error) {

	logs, sub, err := _Document.contract.WatchLogs(opts, "NumberOfOwnerDocument")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DocumentNumberOfOwnerDocument)
				if err := _Document.contract.UnpackLog(event, "NumberOfOwnerDocument", log); err != nil {
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

// ParseNumberOfOwnerDocument is a log parse operation binding the contract event 0x6fc5b69e508af84c5302ddb2ca1e6965cac21ddcad708dce220155101a5db904.
//
// Solidity: event NumberOfOwnerDocument(uint256 num)
func (_Document *DocumentFilterer) ParseNumberOfOwnerDocument(log types.Log) (*DocumentNumberOfOwnerDocument, error) {
	event := new(DocumentNumberOfOwnerDocument)
	if err := _Document.contract.UnpackLog(event, "NumberOfOwnerDocument", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

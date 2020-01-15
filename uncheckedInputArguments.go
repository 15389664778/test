package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init Success!"))
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	res, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	} else {
		return shim.Success(res)
	}
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start Chaincode Failed, error: %s", err)
	}
}

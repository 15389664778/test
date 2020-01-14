package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

var globalVar string

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init Success!"))
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()

	if len(args) != 0 {
		globalVar = args[0]
		stub.PutState("key", []byte(globalVar))
		return shim.Success([]byte("Set Value Success!"))
	}

	return shim.Error("Invalid Function Name!")
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start chaincode failed, error: %s", err)
	}
}

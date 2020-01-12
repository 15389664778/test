package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type BadChaincode struct {
}

func (t *BadChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

var globalValue = ""

func (t BadChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "setValue" {
		globalValue = args[0]
		stub.PutState("key", []byte(globalValue))
		return shim.Success([]byte("success"))
	} else if fn == "getValue" {
		value, _ := stub.GetState("key")
		return shim.Success(value)
	}
	return shim.Error("not a valid function")
}

func main() {
	if err := shim.Start(new(BadChaincode)); err != nil {
		fmt.Printf("Error starting BadChaincode chaincode: %s", err)
	}
}

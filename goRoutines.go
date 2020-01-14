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
	go writeToLedger(stub, "a")
	return shim.Success([]byte("Invoke Success!"))
}

func writeToLedger(stub shim.ChaincodeStubInterface, data string) {
	stub.PutState("key", []byte(data))
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start chaincode failed, error: %s", err)
	}
}

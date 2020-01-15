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
	iterator, _ := stub.GetHistoryForKey("key")
	value, _ := iterator.Next()

	err := stub.PutState("key", value.Value)
	if err != nil {
		return shim.Error("Write data to ledger failed.")
	}

	return shim.Success([]byte("Write data to ledger success!"))
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start Chaincode Failed, error: %s", err)
	}
}

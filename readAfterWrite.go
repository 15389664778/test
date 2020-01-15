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
	key := "key"
	value := "value"

	err := stub.PutState(key, []byte(value))
	if err != nil {
		return shim.Error("Write data to ledger failed.")
	}

	res, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Read data from ledger failed.")
	}

	return shim.Success([]byte(res))
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start Chaincode Failed, error: %s", err)
	}
}

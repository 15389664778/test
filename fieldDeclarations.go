package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
	val string
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init Success!"))
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if len(function) == 1 {
		if function == "setValue" {
			if len(args) != 0 {
				c.val = args[0]
				stub.PutState("key", []byte(c.val))
				return shim.Success([]byte("Set Value Success!"))
			}
		}
	}
	return shim.Error("Invalid Function Name")
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start chaincode failed, error: %s", err)
	}
}

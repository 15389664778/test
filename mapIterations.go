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
var myMap = map[int]string{
		0: "小明",
		1: "张三",
		2: "李四",
		3: "王五",
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	names := ""
	for _, v := range myMap {
		names += v
		names += " "
	}
	return shim.Success([]byte("All names: " + names))
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start Chaincode Failed, error: %s", err)
	}
}

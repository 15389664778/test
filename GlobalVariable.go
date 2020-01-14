//Explaination: https://gist.github.com/arnabkaycee/d4c10a7f5c01f349632b42b67cee46db
//Reason for double spending
//Solution: Make sure not usring global variable inside the operation. Can be used as constant?
//Process: detect if a given path cfg blocks containts the global variables, and then if they are read/write by any operation

//global vairable vulnerability that cannot be detected by chaincode scanner
package main

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type MotivatingChaincode struct {
	dummyv string
	dummyw string
}

//cache of blocks id
var blocks = map[int]int{
	1: 1,
	2: 5,
	3: 10,
	4: 50,
}

var totalNumberOfMarbles0 int = 111
var totalNumberOfMarbles1 = 222

func writeToLedger(stub shim.ChaincodeStubInterface, data string) (err0 error, err1 error) {
	// stub.PutState("key", []byte(data))
	// stub.GetState("key")
	stub.GetHistoryForKey("key")
	stub.PutState("key", []byte(data))
	stub.GetState("key")
	blocks = map[int]int{
		1: 1,
		2: 5,
		3: 10,
		4: 50,
	}
	returnValue := 0
	for i, ii := range blocks {
		returnValue = returnValue*i - ii
	}

	args1 := stub.GetStringArgs()
	if len(args1) < 1 {
		fmt.Println("working guard")
		return err1, err0
	}
	_, err1 = stub.GetState(args1[2])

	return err0, err1
}

func (t *MotivatingChaincode) initMarble(stub shim.ChaincodeStubInterface) {
	args0 := stub.GetStringArgs()
	go writeToLedger(stub, "data1")
	go writeToLedger(stub, "data2")

	go fmt.Println("nothing worry 1") //here we did not check if read and/or write to variables because simutaneously doing so will give defintely dif results. So we just ignore the concurrency operation only if there are no variables involved
	go fmt.Println("nothing worry 2")

	fmt.Sprintf("MARBLE_%06d", totalNumberOfMarbles0)
	fmt.Println(time.Now())
	//--------------CODE SMELL----------------
	//BIG source of Non-determinism as well as performance hit.
	totalNumberOfMarbles0 = totalNumberOfMarbles0 + 1
	//--------------CODE SMELL----------------
	totalNumberOfMarbles1 = 2
	fmt.Sprintf("MARBLE_%06d", totalNumberOfMarbles0)
	resultT, _ := stub.GetState(args0[0])
	_, err1 := writeToLedger(stub, "data")
	// _, err = stub.GetState("key")
	fmt.Println(resultT)
	fmt.Println(err1)
}

//Initialisation of the Chaincode
func (m *MotivatingChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Successfully initialized Chaincode."))
}

//Entry Point of an invocation
func (m *MotivatingChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	m.initMarble(stub)
	return shim.Success([]byte("initMarble function invoked"))
}

func main() {
	if err := shim.Start(new(MotivatingChaincode)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

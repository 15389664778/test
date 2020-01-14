//Explaination:
//Random number for non-determinism. In chaincode scanner, called blacklisted imports
//Solution: No rand-like function used
//Process: Check blacklisted imports collection and make sure these imported function are used. Then give out warning
//Blacklisted imports collection: crypto/rand, math/rand etc

package main

import (
	"fmt"
	// "math/rand"
	// "time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ApartementRegister struct {
}

func (t *ApartementRegister) initMarble(stub shim.ChaincodeStubInterface) {
	var myMap = map[int]int{
		1: 1,
		2: 5,
		3: 10,
		4: 50,
	}
	returnValue := 0
	for i, ii := range myMap {
		returnValue = returnValue*i - ii
	}
	// return shim.Success([]byte("value: " + string(returnValue)))
}

//Initialisation of the Chaincode
func (m *ApartementRegister) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// blocks = make(map[string]bool)
	m.initMarble(stub)
	return shim.Success([]byte("Successfully initialized Chaincode."))
}

//Entry Point of an invocation
func (m *ApartementRegister) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, _ := stub.GetFunctionAndParameters()

	switch function {
	case "initMarble":
		m.initMarble(stub)
	case "registerRenter":
		return shim.Error("not enough arguments for rentersCount. 2 required")
	}
	return shim.Error(fmt.Sprintf("No function %s implemented", function))
}

func main() {
	if err := shim.Start(new(ApartementRegister)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

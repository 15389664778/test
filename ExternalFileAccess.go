//Need to ensure that the results of calling the web service are not different among peers.
////Explaination: When chaincode exhibits non-deterministic behavior, endorsing peers will not compute the same read and write sets. With inconsistent computation, transactions will always be marked as invalid.
//Solution: No os/open package is imported
//Process: Check blacklisted imports collection and make sure these imported function are used. Then give out warning
//Blacklisted imports collection: os/open etc

package main

import (
	"fmt"
	"os"

	// "math/rand"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ApartementRegister struct {
}

func (m *ApartementRegister) initMarble(stub shim.ChaincodeStubInterface) {
	_, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
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

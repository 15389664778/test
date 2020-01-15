//AKA Range Query Risk
//Explaination: Reading the ledger with GetHistoryOfKey or GetQueryResult do not pass the versioning control of the system. Data received from phantom reads should therefore not be used to write new data or update data on the ledger.
// Unexpected behavior might lead to premature termination of a transaction or unintended accesses to the ledger. Thus rendering parts of the chaincode useless.
//Solution: No os/open package is imported
//Process: Check blacklisted imports collection and make sure these imported function are used. Then give out warning
//Blacklisted imports collection: os/open etc

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ApartementRegister struct {
}

//Initialisation of the Chaincode
func (m *ApartementRegister) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// blocks = make(map[string]bool)
	return shim.Success([]byte("Successfully initialized Chaincode."))
}

//Entry Point of an invocation
func (m *ApartementRegister) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	iterator, _ := stub.GetHistoryForKey("key")
	data, _ := iterator.Next()

	err := stub.PutState("key", data.Value)
	if err != nil {
		return shim.Error("could not write new data")
	}
	return shim.Success([]byte("stored"))
}

func main() {
	if err := shim.Start(new(ApartementRegister)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

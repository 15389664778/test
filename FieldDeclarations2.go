//Explaination: The chaincode object is created once during startup of the peer. With field declarations it is possible to sustain a global state in the program. Since these global states are not stored on the ledger, the chaincode object should not accommodate any fields.
//Global states in chaincode should only be managed by the ledger and no acceses to the ledger should be dependent on global states. Since not every peer executes every transaction, states not maintained on the ledger might diverge, resulting in different accesses to the ledger. This leads to inconsistent write and read sets of transactions, and they will always be marked as invalid.
//Solution:There should be no member in the struct of the chaincode
//Process:

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ApartementRegister struct {
	globalValue string
}

//Initialisation of the Chaincode
func (m *ApartementRegister) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// blocks = make(map[string]bool)
	return shim.Success([]byte("Successfully initialized Chaincode."))
}

//Entry Point of an invocation
func (m *ApartementRegister) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	m.globalValue = args[0]
	return shim.Success([]byte("stored"))
}

func main() {
	if err := shim.Start(new(ApartementRegister)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

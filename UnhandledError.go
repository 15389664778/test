//Explaination: For a write statement to take effect, a transaction first has to be committed and written to the ledger. Therefore, when reading a value which has already been written ot during the same transaction, its old value will be retrieved from the ledger. This is usually is not the expected behavior. Unexpected behavior might lead to premature termination of a transaction or unintended accesses to the ledger. Thus rendering parts of the chaincode useless.
////Global states in chaincode should only be managed by the ledger and no acceses to the ledger should be dependent on global states. Since not every peer executes every transaction, states not maintained on the ledger might diverge, resulting in different accesses to the ledger. This leads to inconsistent write and read sets of transactions, and they will always be marked as invalid.
//Solution: Check if W-R data relation exists, especially PutState and GetState
//Process:

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
	args := stub.GetStringArgs()
	if len(args) == 0 {
		return shim.Error("No argument sepcified.")
	}
	result, _ := stub.GetState(args[0])
	return shim.Success([]byte(result))
}

func main() {
	if err := shim.Start(new(ApartementRegister)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

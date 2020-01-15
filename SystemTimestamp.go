//Explaination:

package main

import (
	"fmt"
	// "math/rand"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ApartementRegister struct {
}

func (t *ApartementRegister) initMarble(stub shim.ChaincodeStubInterface) {
	// var seed int64 = 1
	// time.Now().UnixNano()
	// rand.Seed(seed)
	// sel11 := rand.Intn(10)
	sel11 := time.Now().UnixNano
	// PayPrizeToUser(user, prize)
	fmt.Sprintf("user prize ", sel11)
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

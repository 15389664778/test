// Disclaimer
//
// THIS SAMPLE CODE MAY BE USED SOLELY AS PART OF THE TEST AND EVALUATION OF THE SAP CLOUD PLATFORM BLOCKCHAIN SERVICE (THE “SERVICE”)
// AND IN ACCORDANCE WITH THE TERMS OF THE TEST AND EVALUATION AGREEMENT FOR THE SERVICE. THIS SAMPLE CODE PROVIDED “AS IS”, WITHOUT
// ANY WARRANTY, ESCROW, TRAINING, MAINTENANCE, OR SERVICE OBLIGATIONS WHATSOEVER ON THE PART OF SAP.

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MessageStore struct{}

type message struct {
	ID    string `json:"ID"`
	Value string `json:"value"`
}

// Main function starts up the chaincode in the container during instantiate
//
func main() {
	if err := shim.Start(new(MessageStore)); err != nil {
		fmt.Printf("Main: Error starting MessageStore chaincode: %s", err)
	}
}

// Init is called during Instantiate transaction after the chaincode container
// has been established for the first time, allowing the chaincode to
// initialize its internal data. Note that chaincode upgrade also calls this
// function to reset or to migrate data, so be careful to avoid a scenario
// where you inadvertently clobber your ledger's data!
//
func (cc *MessageStore) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// Validate supplied init parameters, in this case zero arguments!
	if _, args := stub.GetFunctionAndParameters(); len(args) > 0 {

		return shim.Error("Init: Incorrect number of arguments; no arguments were expected and none should have been supplied.")
	}
	return shim.Success(nil)
}

// Invoke is called to update or query the ledger in a proposal transaction.
// Updated state variables are not committed to the ledger until the
// transaction is committed.
//
func (cc *MessageStore) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Which function is been called?
	function, args := stub.GetFunctionAndParameters()

	// Turn arguments to lower case
	function = strings.ToLower(function)

	// alternative implemention for the switch statement
	/*
		if function == "write" {
			return cc.write(stub, args)
		} else if function == "read" {
			return cc.read(stub, args)
		} else {
			return shim.Error("Invalid method! Valid methods are (Invoke) 'write' or (Query) 'read'!")
		}
	*/

	// Route call to the correct function
	switch function {
	case "write":
		return cc.write(stub, args)
	case "read":
		return cc.read(stub, args)
	default:
		return shim.Error("Invalid method! Valid methods are (Invoke) 'write' or (Query) 'read'!")
	}
}

// Write an ID and string to the blockchain
//
func (cc *MessageStore) write(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// extract the ID and value from the arguments the content and object id
	if len(args) != 2 {
		return shim.Error("Write: incorrect number of arguments; expecting an ID and the value to be written.")
	}

	// key should be lowercase
	id := strings.ToLower(args[0])

	// create a message struct
	msg := &message{ID: id, Value: args[1]}

	// marshal message struct to json
	// can be returned directly
	// allows more sophisticated search using couchdb /mq
	msgJSON, _ := json.Marshal(msg)

	// alternative implemention for the error check
	/*
		messageAsBytes, err := stub.GetState(id)
		if err != nil {
			return shim.Error("Failed to get record: " + err.Error())
		} else if messageAsBytes != nil {
			return pd.Response{Status: 409,
				Message: "Write: this ID already has a message assigned: " + id,
			}
		}
	*/

	// Validate that this ID does not yet exist
	if messageAsBytes, err := stub.GetState(id); err != nil || messageAsBytes != nil {
		return shim.Error("Write: this ID already has a message assigned.")
	}

	// Write the message
	if err := stub.PutState(id, msgJSON); err != nil {
		return shim.Error(err.Error())
	} else {
		return shim.Success(nil)
	}
}

// Read a string from the blockchain, given its ID
//
func (cc *MessageStore) read(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Read: incorrect number of arguments; expecting only the ID to be read.")
	}
	id := strings.ToLower(args[0])

	if value, err := stub.GetState(id); err != nil || value == nil {
		return shim.Error("Read: invalid ID supplied.")
	} else {
		return shim.Success(value)
	}
}

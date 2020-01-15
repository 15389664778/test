package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

/*
 * The Init method is called when the Smart Contract "simpleEx" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "simpleEx"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryBook" {
		return s.queryBook(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createBook" {
		return s.createBook(APIstub, args)
	} else if function == "queryAllBooks" {
		return s.queryAllBooks(APIstub)
	} else if function == "changeBookTitle" {
		return s.changeBookTitle(APIstub, args)
	} else if function == "getHistoryForKey" {
		return s.getHistoryForKey(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name")
}

func (s *SmartContract) queryBook(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	bookAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(bookAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	books := []Book{
		Book{Title: "Mithila", Author: "Amish"},
		Book{Title: "Alchemist", Author: "Paulo"},
	}

	i := 0
	for i < len(books) {
		fmt.Println("i is ", i)
		bookAsBytes, _ := json.Marshal(books[i])
		APIstub.PutState("Book"+strconv.Itoa(i), bookAsBytes)
		fmt.Println("Added", books[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createBook(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var book = Book{Title: args[1], Author: args[2]}

	bookAsBytes, _ := json.Marshal(book)
	APIstub.PutState(args[0], bookAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllBooks(APIstub shim.ChaincodeStubInterface) sc.Response {
	startKey := "Book0"
	endKey := "Book99"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	fmt.Println(resultsIterator, "----------")
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("- queryAllBooks:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func (s *SmartContract) changeBookTitle(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	bookAsBytes, _ := APIstub.GetState(args[0])
	book := Book{}

	json.Unmarshal(bookAsBytes, &book)
	book.Title = args[1]

	bookAsBytes, _ = json.Marshal(book)
	APIstub.PutState(args[0], bookAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getHistoryForKey(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	resultsIterator, err := APIstub.GetHistoryForKey(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	fmt.Println(resultsIterator)

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.TxId)
		buffer.WriteString("\"")
		// buffer.WriteString("{\"Timestamp\":")
		// buffer.WriteString("\"")
		// buffer.WriteString(queryResponse.Timestamp)
		// buffer.WriteString("\"")
		buffer.WriteString(", \"Value\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("- queryAllBooks:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.

func main() {
	// Create a new Smart Contract

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Println("Error creating new Smart Contract: %s", err)
	}
}

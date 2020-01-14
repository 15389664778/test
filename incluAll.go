package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	// Blacklisted Import
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
	// Field Declaration
	ChaincodeId int
}

type Bank struct {
	Users map[string]User
}

type User struct {
	Name    string
	ID      string
	Balance float32
}

// Global Variable Usage
var bank_1 Bank
var id_library []string

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	c.setChaincodeId()
	return shim.Success([]byte("Init Success!"))
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	// Unchecked Argument && Unhandled Error
	res, _ := stub.GetState(args[0])
	fmt.Println(res)

	user, err := newUser("深信科创", "123", 1000000.00)
	if err != nil {
		return shim.Error(err.Error())
	}

	bank_1.Users = make(map[string]User)

	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "newUser":
		if len(args) == 3 {
			balanceFloat, err := strconv.ParseFloat(args[2], 32)
			if err != nil {
				return shim.Error("输入参数类型有误，导致类型转换失败！")
			}
			user, err := newUser(args[0], args[1], float32(balanceFloat))
			if err != nil {
				return shim.Error("创建用户失败！原因可能是此id已经存在！")
			}
			fmt.Printf("用户%v创建成功！", user)
		} else {
			fmt.Println("参数个数有误！")
		}
	case "joinBank":
		err := joinBank(user)
		if err != nil {
			return shim.Error("添加用户失败！")
		}

		bank_1_json, err := json.Marshal(bank_1)
		if err != nil {
			return shim.Error("Marshal Error！")
		}
		// Read After Write
		err = stub.PutState("bank_1", bank_1_json)
		if err != nil {
			return shim.Error("Write data to ledger failed.")
		}

		bytes, err := stub.GetState("bank_1")
		if err != nil {
			return shim.Error(err.Error())
		}

		bank := new(Bank)
		err = json.Unmarshal(bytes, bank)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println(bank)
	case "getBalance":
		if len(args) == 1 {
			balance, err := getBalance(stub, args[0])
			if err != nil {
				return shim.Error("获取余额失败！")
			}
			balance_json, err := json.Marshal(balance)
			if err != nil {
				return shim.Error("Marshal Error！")
			}
			return shim.Success([]byte(balance_json))
		}
	case "getAllUsers":
		getAllUsers(stub)
	case "getChaincodeId":
		// Concurrency Usage
		c.getChaincodeId()
	case "crossChannelInvock":
		// crossChannelInvock:
		crossChannelInvock(stub)
	default:
		go switchDefault()
		go switchDefault()
	}

	// Phantom Read
	iterator, _ := stub.GetHistoryForKey("bank_1")
	value, _ := iterator.Next()

	err = stub.PutState("bank_1", value.Value)
	if err != nil {
		return shim.Error("Write data to ledger failed.")
	}
	return shim.Success([]byte("Write data to ledger success!"))
}

// 创建新用户
func newUser(name, id string, balance float32) (interface{}, error) {
	for _, v := range id_library {
		if id == v {
			fmt.Println("此id已存在，请更换id!")
		}
	}

	if len(name) == 0 || len(id) == 0 {
		return nil, fmt.Errorf("输入参数有误，用户名或id不能为空！")
	} else {
		user := User{name, id, balance}
		id_library = append(id_library, id)
		return user, nil
	}
}

// 将被正确创建的用户的信息加入银行数据库
func joinBank(user interface{}) error {
	if user != nil {
		switch t := user.(type) {
		case User:
			var new_user User = t
			bank_1.Users[t.ID] = new_user
			fmt.Printf("用户%v添加成功，创建时间：%v。\n", user, time.Now())
			return nil
		default:
			return fmt.Errorf("用户参数类型有误！")
		}
	} else {
		return fmt.Errorf("此用户为无效用户！")
	}
}

// 查询账户余额
func getBalance(stub shim.ChaincodeStubInterface, id string) (float32, error) {
	for _, v := range bank_1.Users {
		if id == v.ID {
			return v.Balance, nil
		}
	}

	return 0, fmt.Errorf("此账户不存在于本行，请在线下进行确认！")
}

// 设置链码id
func (c *Chaincode) setChaincodeId() {
	fmt.Printf("请输入链码ID：")
	c.ChaincodeId = 1
}

// 获取链码id
func (c *Chaincode) getChaincodeId() int {
	return c.ChaincodeId
}

// 跨链调用
func crossChannelInvock(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetArgs()
	response := stub.InvokeChaincode("othercc", args, "otherchannel")
	if response.Status != shim.OK {
		err := fmt.Sprintln("Invoke Chaincode Failed, error: %s", response.Payload)
		return shim.Error(err)
	}
	return shim.Success([]byte("Invoke Chaincode on another Channel Success!"))
}

// 查询所有用户
func getAllUsers(stub shim.ChaincodeStubInterface) peer.Response {
	bytes, err := stub.GetState("bank_1")
	if err != nil {
		return shim.Error(err.Error())
	}
	bank := new(Bank)
	err = json.Unmarshal(bytes, bank)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Map Range Iteration
	if len(bank.Users) != 0 {
		for _, v := range bank.Users {
			fmt.Println(v)
		}
		return shim.Success(nil)
	} else {
		return shim.Error("银行用户数据库空！")
	}
}

func switchDefault() {
	fmt.Println("嘿嘿！函数名无效，请重新输入！")
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Start Chaincode Failed, error: %s", err)
	}
}

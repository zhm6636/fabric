package chaincode

import (
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Test struct {
	contractapi.Contract
	// 测试区块蓝数据的读和写
}

//func (this *Test) Invoke(stub contractapi.TransactionContextInterface) peer.Response {
//	// 入口,更新 添加 查询都可以走这个方法
//	// 依据传递的数据区分调用的方法
//	// 传递:调用方法的名称(get  set)
//	// 传递:依据不同的方法传递参数( get key)(set key value)
//	// 使用ChaincodeStubInterface的GetFunctionAndParameters
//	// 约定:Parameters,如果是get方法index为0中存储key,如果是set方法index为0中存储key,1中存储value
//
//	function, parameters := stub.GetFunctionAndParameters()
//
//	if function == "get" {
//		return this.get(stub, parameters[0])
//
//	} else if function == "set" {
//		return this.set(stub, parameters[0], []byte(parameters[1]))
//	}
//
//	// 方法参数传递错误
//	return shim.Error("Invalid Smart Contract function name.")
//}

/*
读取数据:
读取是一条依据key获取到的内容
*/
func (this *Test) Get(stub contractapi.TransactionContextInterface, key string) (string, error) {

	// 读数据
	// 读数据的结果处理:error  nil
	// 返回读取数据

	data, err := stub.GetStub().GetState(key)

	// 处理异常
	if err != nil {
		return "", err
	}

	// 处理nil的data
	if data == nil {
		// 数据在ledger中不存在
		return "", errors.New("Data not Available")
	}

	return string(data), nil
}

/*
写数据方法
*/
func (this *Test) Set(stub contractapi.TransactionContextInterface, key string, value []byte) error {

	// 写数据 上链
	err := stub.GetStub().PutState(key, value)

	if err != nil {
		return err
	}

	return nil
}

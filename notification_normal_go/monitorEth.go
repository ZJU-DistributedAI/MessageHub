package main

import (
	"./utils"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/garyburd/redigo/redis"
	"log"
	"strings"
	"time"
)

type TransactionReceipt struct {
	TransactionHash string
	BlockHash       string
}

type ModelClientPullDataReceipt struct {
	Metadata string
	From string
}

type DataClientAddMetaDataReceipt struct {

	MetaDataIpfsHash string
	From string

}

type ComputingGetModelReceipt struct {
	ModelIpfsHash string
	ContractHash string
}

type ComputingGetDataReceipt struct {
	DataIpfsHash string
	ModelAddress string
	ComputingHash string
	From string
}

type DataAggreeReceipt struct {
	DataIpfsHash string
	From string
}


type DataClientMonitorComputingAggreeReciept struct {
	ComputingHash string
	From string
}

type DataClientIsAggreeReceipt struct {
	Message string
	From string
}



var mindedTransactionHashChannel chan TransactionReceipt
var modelClientPullDataChannel chan ModelClientPullDataReceipt
var dataAggreeChannel chan DataAggreeReceipt
var computingGetDataChannel chan ComputingGetDataReceipt
var computingGetModelChannel chan ComputingGetModelReceipt
var dataClientMonitorComputingAggreeChannel chan DataClientMonitorComputingAggreeReciept
var dataClientIsAggreeChannel chan DataClientIsAggreeReceipt



func InitChannels(){
	mindedTransactionHashChannel = make(chan TransactionReceipt)
	modelClientPullDataChannel = make(chan ModelClientPullDataReceipt)
	dataAggreeChannel = make(chan DataAggreeReceipt)
	computingGetModelChannel = make(chan ComputingGetModelReceipt)
	computingGetDataChannel = make(chan ComputingGetDataReceipt)
	dataClientMonitorComputingAggreeChannel = make(chan DataClientMonitorComputingAggreeReciept)
	dataClientIsAggreeChannel = make(chan DataClientIsAggreeReceipt)

}

//为每次获取到处于pending状态的交易时，对其创建协程进行监听
//当有交易被验证并封装到区块时，将其塞到mindedTransactionHashChannel里面
func createGoRoutine(client *rpc.Client, txHashes []string) {

	//用来统计未验证交易列表有多少交易被验证了，如果全部被验证则退出协程
	isDone := 0
	for {
		if isDone == len(txHashes) {
			fmt.Println("结束该协程")
			return
		}
		for i := 0; i < len(txHashes); i++ {

			if txHashes[i] == "" {
				continue
			}

			result := utils.GetTransactionReceipt(client, txHashes[i])

			if result.BlockHash != "" {
				mindedTransactionHashChannel <- TransactionReceipt{TransactionHash: result.TransactionHash, BlockHash: result.BlockHash}
				txHashes[i] = ""
				isDone++
			}
		}

		time.Sleep(100)
	}

}

//开启的第一个协程为监听是否有交易已经被封装到区块里，然后将其插入到redis
//开启的第二个协程为监听是否有交易(pending交易)提交到以太坊

func dealNewTransactions(client *rpc.Client, filterId string, conn redis.Conn) {



	//监听以太坊
	go func(client *rpc.Client){
		log.Println("开始监听以太坊")
		for {
			result, _ := utils.GetFilterChanges(client, filterId)

			//获取到处于pending状态的交易Hash
			if len(result) != 0 {
				fmt.Println(result)
				go createGoRoutine(client, result)
			}

			time.Sleep(100)
		}
	}(client)

	go func(conn redis.Conn) {

		log.Println("开始监听通道是否有完成交易")
		for {
			flag := <-mindedTransactionHashChannel
			fmt.Println("接受到交易: ", flag.TransactionHash)
			if flag.TransactionHash != "" {
				//获得交易内容
				transaction,err:=utils.GetTransactionByHash(client,flag.TransactionHash)
				if err!=nil{
					fmt.Println("获得交易内容失败: ",err)
					return
				}
				if transaction.Input!=""{

					distributeTransactionByInput(transaction.From,utils.DecryptTransactionInput(transaction.Input),conn)
				}
			}
			time.Sleep(100)

		}

	}(conn)




}


func InitEthereumTransactionMonitor(){

	client := utils.Connect2Eth()

	conn := utils.Connect2Redis()

	filterId, err := utils.CreateNewPendingTransactionFilter(client)

	if err != nil {
		fmt.Println("创建交易过滤器失败")
		return
	}

	//不使用协程池，监听以太坊
	dealNewTransactions(client, filterId, conn)
}




func GetModelClientPullDataReceipt()(ModelClientPullDataReceipt){


	modelClientPullDataReceipt:= <-modelClientPullDataChannel
	if &modelClientPullDataReceipt != nil{
		return modelClientPullDataReceipt
	}
	return ModelClientPullDataReceipt{}
}



func GetComputingGetDataReceipt()(ComputingGetDataReceipt){



	dataAskComputingReceipt := <- computingGetDataChannel

	if &dataAskComputingReceipt != nil {
		return dataAskComputingReceipt
	}

	return ComputingGetDataReceipt{}

}


func GetComputingGetModelReceipt()(ComputingGetModelReceipt){

	computingGetModelReceipt := <- computingGetModelChannel

	if &computingGetModelReceipt != nil {
		return computingGetModelReceipt
	}
	return ComputingGetModelReceipt{}
}


func GetDataClientMonitorComputingAggreeReceipt()(DataClientMonitorComputingAggreeReciept){


	dataClientMonitorComputingAggreeReciept := <- dataClientMonitorComputingAggreeChannel

	if &dataClientMonitorComputingAggreeReciept != nil {
		return dataClientMonitorComputingAggreeReciept
	}

	return DataClientMonitorComputingAggreeReciept{}


}

func GetDataClientIsAggreeReceipt()(DataClientIsAggreeReceipt){


	dataClientIsAggreeReceipt := <- dataClientIsAggreeChannel

	if &dataClientIsAggreeReceipt != nil {
		return dataClientIsAggreeReceipt
	}

	return DataClientIsAggreeReceipt{}

}

func distributeTransactionByInput(from string,input string,conn redis.Conn){

	splits:=strings.Split(input,":")
	if splits[0] == "dadd"{ //数据方上传元数据
		utils.Sadd2Redis(conn,"metaDataHash", from, splits[1])
	}else if splits[0] == "dpush"{
		computingGetDataChannel <- ComputingGetDataReceipt{DataIpfsHash:splits[1], ModelAddress:splits[2],
			ComputingHash:"", From:splits[3]}
	}else if splits[0] == "ddelete" {
		utils.SremFromRedis(conn, "metaDataHash", from, splits[1])
	}else if splits[0] == "mpull" {//模型方请求数据方数据
		modelClientPullDataChannel <- ModelClientPullDataReceipt{Metadata: splits[1], From: from}
	}else if splits[0] == "daggree" {//数据方同意模型方的请求
		dataAggreeChannel <- DataAggreeReceipt{DataIpfsHash: splits[1], From: from}
	}else if splits[0]=="madd" { //模型方上传模型Hash
		computingGetModelChannel <- ComputingGetModelReceipt{ModelIpfsHash: splits[1], ContractHash:splits[2]}
		//utils.Sadd2Redis(conn,"modelHash",from,splits[1])
	}else if splits[0]=="cadd" { //运算方上传运算资源Hash
		utils.Sadd2Redis(conn,"computingHash",from,splits[1])
	}else if splits[0]== "mask" { //模型方请求数据，根据上传的metaDataHash参数
		//askDataClient(conn,splits[1])
	}else if splits[0]== "dask" { //数据方寻找合适的运算资源，根据上传的运算资源描述Hash
		//askComputing(conn,splits[1])
	}
}


/*

//使用协程池的交易处理函数（暂时未用）
func dealTransactionWithPool(client *rpc.Client, filterId string, pool *Pool) {

	for {
		result, _ := utils.GetFilterChanges(client, filterId)
		if len(result) > 0 {
			//fmt.Println("提交任务")
			fmt.Println(result)
			pool.submit(func() ([]string, *rpc.Client) {
				//fmt.Println("txHash2",result[i])
				return result, client
			})
		}
		time.Sleep(100)
	}

}*/





/*
func main() {

	client := utils.Connect2Eth()

	if client == nil {
		fmt.Println("获得以太坊连接失败")
		return
	}
	defer client.Close()

	conn := utils.Connect2Redis()

	if conn == nil {
		fmt.Println("获得redis连接失败")
		return
	}
	defer conn.Close()

	mindedTransactionHashChannel = make(chan TransactionReceipt)
	modelClientPullDataChannel = make(chan ModelClientPullDataReceipt)




	//使用协程池
	//pool:=NewPool(20)
	//dealTransactionWithPool(client,filterId,pool)
	//wg.Wait()

}*/

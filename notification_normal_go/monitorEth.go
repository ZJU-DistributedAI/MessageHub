package main

import (
	"./utils"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

type ResultReceipt struct {
	TransactionHash string
	BlockHash       string
}

//该通道用来存储已经被验证的交易
var mindedTransactionHashChannel chan ResultReceipt

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
				mindedTransactionHashChannel <- ResultReceipt{TransactionHash: result.TransactionHash, BlockHash: result.BlockHash}
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

	go func(conn redis.Conn) {

		fmt.Println("通道开始监听是否有完成交易")
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

	//监听以太坊
	for {
		result, _ := utils.GetFilterChanges(client, filterId)

		//获取到处于pending状态的交易Hash
		if len(result) != 0 {
			fmt.Println(result)
			go createGoRoutine(client, result)
		}

		time.Sleep(100)
	}

}

func distributeTransactionByInput(from string,input string,conn redis.Conn){

	splits:=strings.Split(input,":")
	if splits[0]=="dadd"{ //数据方上传元数据
		utils.Sadd2Redis(conn,"metadata",from,splits[1])
	}else if splits[0]=="madd"{ //模型方上传模型Hash
		utils.Sadd2Redis(conn,"model",from,splits[1])
	}else if splits[0]=="cadd"{ //运算方上传运算资源Hash
		utils.Sadd2Redis(conn,"computing",from,splits[1])
	}else if splits[0]== "mask"{ //模型方请求数据，根据上传的metaDataHash参数
		askDataClient(conn,splits[1])
	}else if splits[0]== "dask"{ //数据方寻找合适的运算资源，根据上传的运算资源描述Hash
		askComputing(conn,splits[1])
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

	mindedTransactionHashChannel = make(chan ResultReceipt)

	filterId, err := utils.CreateNewPendingTransactionFilter(client)

	if err != nil {
		fmt.Println("创建交易过滤器失败")
		return
	}

	//不使用协程池，监听以太坊
	dealNewTransactions(client, filterId, conn)









	//使用协程池
	//pool:=NewPool(20)
	//dealTransactionWithPool(client,filterId,pool)
	//wg.Wait()

}

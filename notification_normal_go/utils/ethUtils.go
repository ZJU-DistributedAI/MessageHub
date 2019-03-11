package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"time"
)


const(

	SERVERPATH = "http://10.5.19.174:8545"
	LOCALPATH = "http://127.0.0.1:8545"


)


type Client struct {
	rpcClient *rpc.Client
	EthClient *ethclient.Client
}

type Message struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Value    string          `json:"value"`
	Data     string          `json:"data"`
	GasLimit string          `json:"gasLimit"`
	GasPrice string          `json:"gasPrice"`
}

type Transaction struct {
	From string `json:"from"`
	Hash string `json:"hash"`
	Input string `json:"input"`
}

type Receipt struct {
	TransactionHash  string          `json:"transactionHash"`
	TransactionIndex string          `json:"transactionIndex"`
	BlockHash        string          `json:"blockHash"`
	BlockNumber      string          `json:"blockNumber"`
	To               *common.Address `json:"to"`
	From             common.Address  `json:"from"`
}

type Block struct {
	Number string `json:"number"`
	Hash   string `json:"hash"`
	//ParentHash string `json:"parentHash"`
	//nonce string `json:"nonce"`
	//sha3Uncles string `json:"sha3Uncles"`
	//logsBloom string `json:"logsBloom"`
	//TransactionsRoot string `json:"transactionsRoot"`
	//stateRoot string `json:"stateRoot"`
	//miner string `json:"miner"`
	//difficulty string `json:"difficulty"`
	//totalDifficulty string  `json:"totalDifficulty"`
	//extraData string `json:"extraData"`
	//size string  `json:"size"`
	//gasLimit string `json:"gasLimit"`
	//gasUsed string `json:"gasUsed"`
	//timestamp string `json:timestamp`
	Transactions []Transaction `json:transactions`
	//uncles []string `json:uncles`
}
var clientConnect *rpc.Client

func Connect2Eth() (*rpc.Client) {

	for {

		if clientConnect == nil {
			mutex.Lock()
			if clientConnect == nil{
				client, err := rpc.Dial(LOCALPATH)
				if err != nil {
					ErrorPanic(err)
					time.Sleep(500) //等待0.5秒后重新尝试连接
				} else {
					clientConnect = client
				}
			}
			mutex.Unlock()
			if clientConnect != nil{
				break
			}
		}else{
			break
		}
	}
	return clientConnect

}

func NewMessage(from common.Address, to *common.Address, value string, data string, gasLimit string, gasPrice string) (Message) {

	return Message{From: from, To: to, Value: value, Data: data, GasLimit: gasLimit, GasPrice: gasPrice}

}

func GetEthCoinbase(client *rpc.Client) (string, error) {

	var result string

	err := client.Call(&result, "eth_coinbase")

	return result, err

}

func CreateAccount(client *rpc.Client, password string) (string) {
	var result string
	err := client.Call(&result, "personal_newAccount",password)
	if err != nil{
		ErrorPanic(err)
		return ""
	}
	return result
}

func UnlockAccount(client *rpc.Client, account string, password string) (error) {

	var result bool
	err := client.Call(&result, "personal_unlockAccount", account, password)
	return err
}

func SendTransaction(client *rpc.Client, tx *Message, password string, ctx context.Context) (string, error) {

	var txHash string
	err := client.CallContext(ctx, &txHash, "personal_signAndSendTransaction", tx, password)
	//err:=client.rpcClient.Call(&result,"eth_sendTransaction",tx)
	return txHash, err
}

func SendRawTransaction(client *rpc.Client, data string) (string, error) {
	var result string

	err := client.Call(&result, "eth_sendRawTransaction", data)

	return result, err
}


func GetTransactionByHash(client *rpc.Client, transationHash string)(Transaction,error){
	var result Transaction
	err:=client.Call(&result, "eth_getTransactionByHash",transationHash)
	return result,err
}



func CreateNewPendingTransactionFilter(client *rpc.Client) (string, error) {

	var filterId string
	err := client.Call(&filterId, "eth_newPendingTransactionFilter")
	return filterId, err
}

func GetFilterChanges(client *rpc.Client, filterId string) ([]string, error) {
	var result []string

	err := client.Call(&result, "eth_getFilterChanges", filterId)
	return result, err
}

func GetTransactionReceipt(client *rpc.Client, txHash string) (Receipt) {

	var result Receipt
	_ = client.Call(&result, "eth_getTransactionReceipt", txHash)
	//fmt.Println(result)
	return result
}

func SetEtherBase(client *rpc.Client, account string)bool{
	var result bool
	err := client.Call(&result, "miner_setEtherbase", account)
	if err != nil{
		log.Fatal(err)
	}
	return result
}

func StartMiner(client *rpc.Client, number int) bool {
	var result bool
	err := client.Call(&result, "miner_start", number)
	if err != nil{
		log.Fatal(err)
	}
	return result
}

func GetBalance(client *rpc.Client, account string) string{
	var result string
	err := client.Call(&result, "eth_getBalance", account, "latest")
	if err != nil{
		log.Fatal(err)
	}
	return result
}
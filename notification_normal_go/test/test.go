//package notification_normal_go

package main

import (
	"fmt"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"../utils"
	)


func sendMultiTransaction(client *rpc.Client,from string,to common.Address,number int){

	for i:=0;i<number;i++{

		data:=utils.EncryptTransactionInput("cadd:"+"hash"+string(i)+":"+"contractHash")

		fmt.Println(data)

		message:=utils.NewMessage(common.HexToAddress(from),&to,"0x10","0x"+data,"0x295f05","0x77359400")

		txHash,err:=utils.SendTransaction(client,&message,context.TODO())
		//txHash,err:=sendRawTransaction(client,"")
		if err!=nil{
			fmt.Println("send new transaction fail: ",err)
			break
		}
		fmt.Println(txHash)

	}



}


func main(){

	client:=utils.Connect2Eth()



	if client==nil {
		fmt.Println("获得以太坊连接失败")
		return
	}
	defer client.Close()



	var account[]string
	err := client.Call(&account, "eth_accounts")

	if err!=nil{
		fmt.Println("获得账户失败! ",err)
	}

	//fmt.Println("account:",account)

	to:=common.HexToAddress(account[1])

	fmt.Println("to: ",to)


	err=utils.UnlockAccount(client,account[0],"abc")

	if err!=nil{
		fmt.Println("unlockAccount fail: ",err)
	}
	sendMultiTransaction(client,account[0],to,2)


}
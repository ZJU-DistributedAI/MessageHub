//package notification_normal_go

package main

import (
	"../utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func sendMultiTransaction(client *rpc.Client, from string, to common.Address, number int) {

	for i := 0; i < number; i++ {

		data := utils.EncryptTransactionInput("cadd:" + "hash" + string(i) + ":" + "contractHash")

		fmt.Println(data)

		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10", "0x"+data, "0x295f05", "0x77359400")

		txHash, err := utils.SendTransaction(client, &message, "abc", context.TODO())
		//txHash,err:=sendRawTransaction(client,"")
		if err != nil {
			fmt.Println("send new transaction fail: ", err)
			break
		}
		fmt.Println(txHash)

	}

}

func main() {

	client := utils.Connect2Eth()

	if client == nil {
		fmt.Println("获得以太坊连接失败")
		return
	}
	defer client.Close()

	// create account
	myaccount := utils.CreateAccount(client, "abc")
	if myaccount == "" {
		fmt.Println("account create failed")
	}

	to := common.HexToAddress(myaccount)
	fmt.Println("to: ", to)

	err := utils.UnlockAccount(client, myaccount, "abc")

	if err != nil {
		fmt.Println("unlockAccount fail: ", err)
	}
	sendMultiTransaction(client, myaccount, to, 2)

}

package main

import (

	"log"
	"net/http"
)

func startMessageWatchingService(){



	http.Handle("./template",http.FileServer(http.Dir("template")))

	http.HandleFunc("/",LoginHandler)
	http.HandleFunc("/listWholeMetaData",ListMetaData)
	http.HandleFunc("/listWholeComputing",ListComputing)
	http.HandleFunc("/listAskedMetaData",ListAskedMetaData)
	http.HandleFunc("/listAskedComputing",ListAskedComputing)

	//user
	http.HandleFunc("/user/login",LoginHandler)
	http.HandleFunc("/user/createwallet",CreateWalletHandler)


	//data client
	http.HandleFunc("/dataclient/adddata",DataClientAddDataHandler)
	http.HandleFunc("/dataclient/agreerequest",DataClientAggreeRequestHandler)
	http.HandleFunc("/dataclient/askcomputing",DataClientAskComputingHandler)
	http.HandleFunc("/dataclient/deletedata",DataClientDeleteDataHandler)

	//model client
	http.HandleFunc("/modelclient/askdata",ModelClientAskDataHandler)
	http.HandleFunc("/modelclient/createcontract",ModelClientCreateContractHandler)
	http.HandleFunc("/modelclient/uploadmodel",ModelClientUploadModelHandler)
	http.HandleFunc("/modelclient/uploadresult",ModelClientUploadResultHandler)

	//computing client
	http.HandleFunc("/computingclient/adddata",ComputingClientAddDataHandler)
	http.HandleFunc("/computingclient/agreerequest",ComputingClientAggreeRequestHandler)
	http.HandleFunc("/computingclient/deletedata",ComputingClientDeleteDataHandler)
	http.HandleFunc("/computingclient/train",ComputingClientTrainHandler)
	http.HandleFunc("/computingclient/uploadencrypteddata",ComputingClientUploadEncryptedDataHandler)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}








func main(){

	//监听Model Client和Data Client传过来的http请求
	//_:= handler.Handlers{conn: utils.Connect2Eth()}
	startMessageWatchingService()

}

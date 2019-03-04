package main

import (
	"./utils"
	"github.com/getsentry/raven-go"
	"log"
	"net/http"
)

func startWebService(){



	http.Handle("./template",http.FileServer(http.Dir("template")))

	http.HandleFunc("/", LoginHandler)
	//http.HandleFunc("/listWholeMetaData", ListMetaData)
	//http.HandleFunc("/listWholeComputing", ListComputing)
	//http.HandleFunc("/listAskedMetaData", ListAskedMetaData)
	//http.HandleFunc("/listAskedComputing", ListAskedComputing)

	//user
	http.HandleFunc("/user/login", LoginHandler)
	http.HandleFunc("/user/createwallet", CreateWalletHandler)


	//data client
	http.HandleFunc("/dataclient/adddata", DataClientAddDataHandler)
	http.HandleFunc("/dataclient/pushdatatocomputing", DataClientPushDataToComputingHandler)
	http.HandleFunc("/dataclient/aggreemodelclient", DataClientAggreeModelClientHandler)
	http.HandleFunc("/dataclient/askcomputing", DataClientAskComputingHandler)
	http.HandleFunc("/dataclient/deletedata", DataClientDeleteDataHandler)
	http.HandleFunc("/dataclient/monitormetadata", DataClientMonitorMetaDataHandler)
	http.HandleFunc("/data/client/monitorcomputingaggree", DataClientMonitorComputingAggreeHandler)

	//model client
	http.HandleFunc("/modelclient/askdata", ModelClientAskDataHandler)
	http.HandleFunc("/modelclient/createcontract", ModelClientCreateContractHandler)
	http.HandleFunc("/modelclient/uploadmodel", ModelClientUploadModelHandler)
	http.HandleFunc("/modelclient/uploadresult", ModelClientUploadResultHandler)
	http.HandleFunc("/modelclient/monitordataclient", ModelClientMonitorDataClientResultHandler)

	//computing client
	http.HandleFunc("/computingclient/adddata", ComputingClientAddDataHandler)
	http.HandleFunc("/computingclient/agreerequest", ComputingClientAggreeRequestHandler)
	http.HandleFunc("/computingclient/deletedata", ComputingClientDeleteDataHandler)
	http.HandleFunc("/computingclient/train", ComputingClientTrainHandler)
	http.HandleFunc("/computingclient/uploadencrypteddata", ComputingClientUploadEncryptedDataHandler)
	http.HandleFunc("/computingclient/monitordataclient", ComputingClientMonitorDataClientHandler)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}




func init(){

	//连接到sentry
	raven.SetDSN("http://e8f71faeb2d043fd96058c57a481434c:b7afc614ef7149c997cb9b316d7f2eaf@212.64.85.208:9000/3")
	//连接以太坊
	utils.Connect2Eth()
	//连接redis
	utils.Connect2Redis()
	//获取配置文件对象
	utils.InitConfig("")

	InitChannels()


}



func main(){

	//监听Model Client和Data Client传过来的http请求
	//_:= handler.Handlers{conn: utils.Connect2Eth()}
	startWebService()

}

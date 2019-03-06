package main

import (
	"./utils"
	"github.com/getsentry/raven-go"
	"log"
	"net/http"
)

func startWebService() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("./template", http.FileServer(http.Dir("template")))

	http.HandleFunc("/", LoginHandler)
	//http.HandleFunc("/listWholeMetaData", ListMetaData)
	//http.HandleFunc("/listWholeComputing", ListComputing)
	//http.HandleFunc("/listAskedMetaData", ListAskedMetaData)
	//http.HandleFunc("/listAskedComputing", ListAskedComputing)

	//user
	http.HandleFunc("/user/login", LoginHandler)
	http.HandleFunc("/user/checklogin", CheckLoginHandler)
	http.HandleFunc("/user/createwallet", CreateWalletHandler)
	http.HandleFunc("/user/createwalletpage", CreateWalletPageHandler)
	http.HandleFunc("/user/downloadtool", DownloadToolPageHandle)

	//data client

	http.HandleFunc("/dataclient/addmetadata", DataClientAddMetaDataHandler)
	http.HandleFunc("/dataclient/pushdatatocomputing", DataClientPushDataToComputingHandler) //todo fontend
	http.HandleFunc("/dataclient/aggreemodelclient", DataClientAggreeModelClientHandler)     //todo fontend
	http.HandleFunc("/dataclient/index", IndexDataHandle)                                    //font OK, back ok
	http.HandleFunc("/dataclient/walletpage", DataClientWalletPageHandler)                   //font OK, back ok
	http.HandleFunc("/dataclient/availablecomputingpage", DataClientAvaCompHandle)           //font OK, back ok
	http.HandleFunc("/dataclient/askcomputing", DataClientAskComputingHandler)
	http.HandleFunc("/dataclient/deletedata", DataClientDeleteDataHandler)
	http.HandleFunc("/dataclient/monitormetadata", DataClientMonitorMetaDataHandler)
	http.HandleFunc("/data/client/monitorcomputingaggree", DataClientMonitorComputingAggreeHandler)

	//model client
	http.HandleFunc("/modelclient/index", IndexModelHandle)                  //font OK, back ok
	http.HandleFunc("/modelclient/walletpage", ModelClientWalletPageHandler) //font OK, back ok
	http.HandleFunc("/modelclient/availabledatapage", ModelClientAvaDataHandle)
	http.HandleFunc("/modelclient/askdata", ModelClientAskDataHandler)
	http.HandleFunc("/modelclient/createcontract", ModelClientCreateContractHandler)
	http.HandleFunc("/modelclient/uploadmodel", ModelClientUploadModelHandler)
	http.HandleFunc("/modelclient/uploadresult", ModelClientUploadResultHandler)
	http.HandleFunc("/modelclient/monitordataclient", ModelClientMonitorDataClientResultHandler)

	//computing client

	http.HandleFunc("/computingclient/index", IndexComputingHandle)                       //font OK, back ok
	http.HandleFunc("/computingclient/walletpage", ComputingClientWalletPageHandler)      //font OK, back ok
	http.HandleFunc("/computingclient/adddata", ComputingClientAddDataHandler)            //font OK,
	http.HandleFunc("/computingclient/agreerequest", ComputingClientAggreeRequestHandler) //font OK,
	http.HandleFunc("/computingclient/deletedcomputing", ComputingClientDeleteComputingHashHandler)

	http.HandleFunc("/computingclient/train", ComputingClientTrainHandler)
	http.HandleFunc("/computingclient/uploadencrypteddata", ComputingClientUploadEncryptedDataHandler)
	http.HandleFunc("/computingclient/monitordataclient", ComputingClientMonitorDataClientHandler)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func init() {

	//连接到sentry
	raven.SetDSN("http://1b94cad9e5df4ef2b743b5fc43ad3e6f:4808eab3818b4ef1bc2f068d4278d207@212.64.85.208:9000/2")
	//连接以太坊
	utils.Connect2Eth()
	//连接redis
	utils.Connect2Redis()
	//获取配置文件对象
	// utils.InitConfig("")

	InitChannels()

}

func main() {

	//监听Model Client和Data Client传过来的http请求
	//_:= handler.Handlers{conn: utils.Connect2Eth()}
	startWebService()

}

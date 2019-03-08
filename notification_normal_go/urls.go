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

	http.HandleFunc("/", HomeHandler)
	//http.HandleFunc("/listWholeMetaData", ListMetaData)
	//http.HandleFunc("/listWholeComputing", ListComputing)
	//http.HandleFunc("/listAskedMetaData", ListAskedMetaData)
	//http.HandleFunc("/listAskedComputing", ListAskedComputing)

	//user
	http.HandleFunc("/user/register", RegisterHandler)
	http.HandleFunc("/user/login", LoginHandler)
	http.HandleFunc("/user/checklogin", CheckLoginHandler)
	http.HandleFunc("/user/createwallet", CreateWalletHandler)
	http.HandleFunc("/user/createwalletpage", CreateWalletPageHandler)
	http.HandleFunc("/user/downloadtool", DownloadToolPageHandle)

	//data client
	http.HandleFunc("/dataclient/uploadfile", DataClientUplodFile)
	http.HandleFunc("/dataclient/addmetadata", DataClientAddMetaDataHandler)                 //font OK,
	http.HandleFunc("/dataclient/pushdatatocomputing", DataClientPushDataToComputingHandler) //font OK,
	http.HandleFunc("/dataclient/aggreemodelclient", DataClientAggreeModelClientHandler)     //font OK,
	http.HandleFunc("/dataclient/index", IndexDataHandle)                                    //font OK, back ok
	http.HandleFunc("/dataclient/walletpage", DataClientWalletPageHandler)                   //font OK, back ok
	http.HandleFunc("/dataclient/availablecomputingpage", DataClientAvaCompHandle)           //font OK, back ok
	http.HandleFunc("/dataclient/modelaskingpage", DataClientModelAskHandle)           //font OK, back ok

	http.HandleFunc("/dataclient/askcomputing", DataClientAskComputingHandler)               //font OK,
	http.HandleFunc("/dataclient/deletedata", DataClientDeleteDataHandler)                   //font OK,
	http.HandleFunc("/dataclient/monitormetadata", DataClientMonitorMetaDataHandler)
	http.HandleFunc("/data/client/monitorcomputingaggree", DataClientMonitorComputingAggreeHandler)

	//model client
	http.HandleFunc("/modelclient/uploadfile", ModelClientUploadFile)
	http.HandleFunc("/modelclient/index", IndexModelHandle)                     //font OK, back ok
	http.HandleFunc("/modelclient/walletpage", ModelClientWalletPageHandler)    //font OK, back ok
	http.HandleFunc("/modelclient/availabledatapage", ModelClientAvaDataHandle) //font OK, back ok
	http.HandleFunc("/modelclient/askdata", ModelClientAskDataHandler)          //font OK,
	http.HandleFunc("/modelclient/createcontract", ModelClientCreateContractHandler)
	http.HandleFunc("/modelclient/uploadmodel", ModelClientUploadModelHandler) //font OK,
	http.HandleFunc("/modelclient/uploadresult", ModelClientUploadResultHandler)
	http.HandleFunc("/modelclient/monitordataclient", ModelClientMonitorDataClientResultHandler)

	//computing client
	http.HandleFunc("/computingclient/uploadfile", ComputingClientUploadFile)
	http.HandleFunc("/computingclient/index", IndexComputingHandle)                                 //font OK, back ok
	http.HandleFunc("/computingclient/walletpage", ComputingClientWalletPageHandler)                //font OK, back ok
	http.HandleFunc("/computingclient/trainpage", ComputingClientTrainPageHandle)                                 //font OK, back ok
	http.HandleFunc("/computingclient/dataaskingpage", ComputingClientDataAskHandle)           //font OK, back ok
	http.HandleFunc("/computingclient/adddata", ComputingClientAddDataHandler)                      //font OK,
	http.HandleFunc("/computingclient/agreerequest", ComputingClientAggreeRequestHandler)           //font OK,
	http.HandleFunc("/computingclient/deletedcomputing", ComputingClientDeleteComputingHashHandler) //font OK,

	http.HandleFunc("/computingclient/trainrequest", ComputingClientTrainReceiptHandler)//font OK,
	http.HandleFunc("/computingclient/train", ComputingClientTrainHandler)//font OK,
	http.HandleFunc("/computingclient/uploadencrypteddata", ComputingClientUploadEncryptedDataHandler)
	http.HandleFunc("/computingclient/monitordataclient", ComputingClientMonitorDataClientHandler)
	http.HandleFunc("/computingclient/getdockerstatus", ComputingClientGetDockerStatus)

	//DockerBackend CallBack
	http.HandleFunc("/callback/updatedockerstatus", UpdateDockerStatusHandler)


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
	//连接mysql
	utils.InitMysqlConnection()
	//获取配置文件对象
	// utils.InitConfig("")

	InitChannels()

}

func main() {


	startWebService()

}

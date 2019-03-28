package main

import (
	"log"
	"net/http"

	"./utils"
	raven "github.com/getsentry/raven-go"
)

func startWebService() {

	//data client
	http.HandleFunc("/dataclient/uploadfile", DataClientUplodFileHandler)

	http.HandleFunc("/dataclient/monitormetadata", DataClientMonitorMetaDataHandler)                //font ok
	http.HandleFunc("/data/client/monitorcomputingaggree", DataClientMonitorComputingAggreeHandler) //font ok

	//model client
	http.HandleFunc("/modelclient/uploadfile", ModelClientUploadFileHandler)
	http.HandleFunc("/modelclient/createcontract", ModelClientCreateContractHandler) //font ok

	http.HandleFunc("/modelclient/monitordataclient", ModelClientMonitorDataClientResultHandler)
	http.HandleFunc("/modelclient/monitorparameter", ModelClientMonitorParamterHandler)

	//computing client
	http.HandleFunc("/computingclient/uploadfile", ComputingClientUploadFileHandler)

	http.HandleFunc("/computingclient/trainrequest", ComputingClientTrainReceiptHandler) //font OK,
	http.HandleFunc("/computingclient/train", ComputingClientTrainHandler)               //font OK,
	http.HandleFunc("/computingclient/uploadencrypteddata", ComputingClientUploadEncryptedDataHandler)
	http.HandleFunc("/computingclient/monitordata", ComputingClientMonitorDataHandler)   //font OK,
	http.HandleFunc("/computingclient/monitormodel", ComputingClientMonitorModelHandler) //font OK,
	http.HandleFunc("/computingclient/getdockerstatus", ComputingClientGetDockerStatus)

	//DockerBackend CallBack
	http.HandleFunc("/callback/updatedockerstatus", UpdateDockerStatusHandler)

	// 启动web服务，监听9090端口
	log.Println("Start Listen")
	err := http.ListenAndServe(":9092", nil)
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
	//初始化以太坊监听器
	InitEthereumTransactionMonitor()
	//初始化通道
	InitChannels()

}

func main() {

	startWebService()

}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"

	"./abi"
	"./utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MetaDataList struct {
	MetaList []string
}

type ComputingList struct {
	Computings []string
}

type Data struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type Handlers struct {
	conn *rpc.Client
}

// TODO: test
func matchData(splits []string) {
	//获得所有的数据描述计算数据定价
	conn := utils.Connect2Redis()
	defer conn.Close()

	//TODO 通过计算DQ来获得最佳数据,这里暂时用result[0]来代表
	metaData := splits[1]

	//请求数据方询问是否同意吧最佳数据的Hash给模型方
	resp, err := http.Get("url?data_hash=" + metaData)
	if err != nil {
		log.Println("请求数据方出错：%v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取response数据失败：%v", err)
	}

	//TODO 万一最佳数据的提供方不允许提供最佳数据则需要一个机制来选择不是最佳的数据
	if string(body) == "ok" {
		//fmt.Fprint(w,bestData)
		fmt.Println("数据方同意并把真实数据Hash存入智能合约")
	} else {
		fmt.Println("sorry the no best data")
	}

}

// TODO: test
func matchComputing(splits []string) {
	//获得所有的数据描述计算数据定价
	conn := utils.Connect2Redis()
	defer conn.Close()

	computingMetaHash := splits[1]

	//请求运算方询问是否同意运算资源Hash给区块链
	resp, err := http.Get("url?computing_meta_hash=" + computingMetaHash)
	if err != nil {
		log.Println("请求运算方出错：%v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取response数据失败：%v", err)
	}

	//TODO 万一最佳数据的提供方不允许提供最佳数据则需要一个机制来选择不是最佳的数据
	if string(body) == "ok" {
		fmt.Println("运算方同意并把运算资源Hash存入以太坊")
	} else {
		fmt.Println("sorry the no best data")
	}
}

//response.setHeader("Access-Control-Allow-Origin", "*");
//response.setHeader("Access-Control-Allow-Method", "POST,GET");

// TODO: test
//TODO 返回数据描述信息列表，让ModelClient挑选最佳数据
func ListMetaData(w http.ResponseWriter, r *http.Request) {
	//这里我们用json的格式返回所有data描述信息

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	conn := utils.Connect2Redis()
	defer conn.Close()

	result := utils.SmembersFromRedis(conn, "metadata")

	metaDataList := MetaDataList{
		MetaList: result,
	}

	js, _ := json.Marshal(metaDataList)

	w.Write([]byte(string(js)))

}

// TODO: test
//TODO 返回运算资源描述信息列表，让DataClient挑选最佳运算资源
func ListComputing(w http.ResponseWriter, r *http.Request) {
	//这里我们用json的格式返回所有computing描述信息

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	conn := utils.Connect2Redis()
	defer conn.Close()

	result := utils.SmembersFromRedis(conn, "computing")

	metaDataList := MetaDataList{
		MetaList: result,
	}

	js, _ := json.Marshal(metaDataList)

	w.Write([]byte(string(js)))

}

// TODO: test
func ListAskedMetaData(w http.ResponseWriter, r *http.Request) {

	conn := utils.Connect2Redis()
	defer conn.Close()

	result := utils.SmembersFromRedis(conn, r.FormValue("address"))

	metaDataList := MetaDataList{
		MetaList: result,
	}

	js, _ := json.Marshal(metaDataList)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	w.Write([]byte(string(js)))

}

// TODO: test
func ListAskedComputing(w http.ResponseWriter, r *http.Request) {
	conn := utils.Connect2Redis()
	defer conn.Close()

	result := utils.SmembersFromRedis(conn, r.FormValue("address"))

	computingList := ComputingList{
		Computings: result,
	}

	js, _ := json.Marshal(computingList)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	w.Write([]byte(string(js)))
}

//
//func IndexHandler(w http.ResponseWriter, r *http.Request) {
//
//	// view
//	t, err := template.ParseFiles("template/index.html")
//	if err != nil {
//		utils.ErrorPanic(err)
//	}
//
//	// header
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
//	// body
//	t.Execute(w, nil)
//}

// TODO: test
func DataClientMonitorMetaDataHandler(w http.ResponseWriter, request *http.Request) {
	/**
	接受ajax长轮询
	监听是否有模型方提交的metadata交易
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	log.Println("收到请求DataClientMonitorMetaDataHandler")
	var data Data

	modelClientPullDataReceipt := GetModelClientPullDataReceipt()

	data = Data{Msg: modelClientPullDataReceipt.Metadata + ":" + modelClientPullDataReceipt.From, Code: 200}
	js, _ := json.Marshal(data)
	w.Write(js)
}

// TODO: test
func ModelClientMonitorDataClientResultHandler(w http.ResponseWriter, request *http.Request) {

	/**
	模型方监听数据方的响应（是否同意模型方的数据请求）
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexmodel.html")

	dataClientIsAggreeReceipt := GetDataClientIsAggreeReceipt()

	data = Data{Msg: dataClientIsAggreeReceipt.Message + ":" + dataClientIsAggreeReceipt.From, Code: 200}
	js, _ := json.Marshal(data)
	t.Execute(w, js)

}

// TODO: test
func DataClientMonitorComputingAggreeHandler(w http.ResponseWriter, request *http.Request) {

	/**
	ajax长连接
	数据方监听是否有运算方的同意交易
	*/

	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	log.Println("收到请求DataClientMonitorComputingAggreeHandler")

	// handle
	var data Data
	dataClientMonitorComputingAggreeReceipt := GetDataClientMonitorComputingAggreeReceipt()
	data = Data{Msg: dataClientMonitorComputingAggreeReceipt.ComputingHash + ":" + dataClientMonitorComputingAggreeReceipt.From,
		Code: 200}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

const key = `{"address":"f448d0ae08287173002d06093abdab2ac1d7ce9a", "crypto":{"cipher":"aes-128-ctr", "ciphertext":"b8ed8722c4bce17a035d14d3134c9f49a25476f8de84872bdb8921cdaf418fed", "cipherparams":{"iv":"ad9f92893a3b2ea5a4fc8b038a7c79cb"}, "kdf":"scrypt", "kdfparams":{"dklen":32, "n":262144, "p":1, "r":8, "salt":"73e2d83f69e7afbb323cd73831a50004501366c687698818e4ce9440ca657a9b"}, "mac":"f404ba734343531b6346506f363343151ae67dd26229ae1748c53e116c5939bd"}, "id":"00e53020-8296-4ffb-9d6b-c56d2d246b2c", "version":3}`

// TODO: test
func ModelClientCreateContractHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	rpcConn := utils.Connect2Eth()
	ethConn := ethclient.NewClient(rpcConn)
	password := request.PostFormValue("password")
	auth, err := bind.NewTransactor(strings.NewReader(key), password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Deploy a new contract
	address, tx, _, err := abi.DeployMain(auth, ethConn)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, ethConn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}

	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}

	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexmodel.html")
}

func ModelClientMonitorParamterHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	modelAddress := request.PostFormValue("modeladdress")

	result := utils.GetFederateLearningResult(modelAddress)

	response := Data{Msg: result, Code: 200}

	js, _ := json.Marshal(response)
	w.Write(js)

}

// TODO: test
func ModelClientUploadResultHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data

}

// ComputingClientMonitorDataHandler ComputingClient listen data client request
func ComputingClientMonitorDataHandler(w http.ResponseWriter, request *http.Request) {
	/**
	ajax长连接
	运算方监听是否有数据方的运算资源请求
	*/
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var response Data
	computingGetDataReceipt := GetComputingGetDataReceipt()
	response = Data{Msg: computingGetDataReceipt.DataIpfsHash + ":" + computingGetDataReceipt.From + ":" + computingGetDataReceipt.ModelAddress, Code: 200}

	// response
	js, _ := json.Marshal(response)
	w.Write(js)
}

// ComputingClientMonitorModelHandler ComputingClient listen model client request
func ComputingClientMonitorModelHandler(w http.ResponseWriter, request *http.Request) {
	/**
	ajax长连接
	运算方监听是否有model方的运算资源请求
	*/
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var response Data
	computingGetModelReceipt := GetComputingGetModelReceipt()
	response = Data{Msg: computingGetModelReceipt.ModelIpfsHash + ":" + computingGetModelReceipt.ContractHash, Code: 200}

	// response
	js, _ := json.Marshal(response)
	w.Write(js)
}

// ComputingClientTrainReceiptHandler Computing client send train
func ComputingClientTrainReceiptHandler(w http.ResponseWriter, request *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var response Data
	_, err := http.Get("http://127.0.0.1:9091/dockerbackend/trainrequest")
	if err != nil {
		log.Println("与容器后端通信失败", err)
		// response = Data{Msg: "与容器后端通信失败", Code: 500}
		response = Data{Msg: "send message docker backend failed", Code: 500}
	} else {
		// response = Data{Msg: "接受到训练请求", Code: 200}
		response = Data{Msg: "receipt train request", Code: 200}
	}

	//response
	js, _ := json.Marshal(response)
	w.Write(js)
}

// ComputingClientTrainHandler operate docker backend to train model
func ComputingClientTrainHandler(w http.ResponseWriter, request *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	/**
	调用容器后端接口训练模型
	*/
	var response Data

	modelIpfsHash := request.PostFormValue("modelIpfsHash")
	dataIpfsHash := request.PostFormValue("dataIpfshash")

	modelAddress := request.PostFormValue("train_modelFrom")

	fmt.Println("modelAddress: ", modelAddress)

	uploadPath, directoryPath := utils.MakeDirectory("train_" + modelAddress)
	if uploadPath != "" {
		utils.DownloadFile(modelIpfsHash, uploadPath+"modelFile.json")
		utils.DownloadFile(dataIpfsHash, uploadPath+"dataFile.json")
		utils.CopyTrainCode(directoryPath)
	}

	_, err := http.Get("http://127.0.0.1:9093/dockerbackend/starttrain?from=" + modelAddress + "&directorypath=" + directoryPath)
	if err != nil {
		log.Println("运算方调用容器后端失败", err)
		// response = Data{Msg: "运算方调用容器后端失败", Code: 500}
		response = Data{Msg: "Computing client contact docker backend", Code: 500}
	} else {
		//result := utils.ReadFile(directoryPath+"//parameters.json")

		response = Data{Msg: "Train success", Code: 200}
	}

	// response
	js, _ := json.Marshal(response)
	w.Write(js)
}

// ComputingClientGetDockerStatus Computing client get docker status
func ComputingClientGetDockerStatus(w http.ResponseWriter, request *http.Request) {

	/**
	获取docker状态
	*/

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	useraccount := request.PostFormValue("from")

	status := utils.GetDockerStatus(useraccount)

	var response Data

	if status == 0 {
		// response = Data{Msg: "获取状态失败", Code: 500}
		response = Data{Msg: "get docker status failed", Code: 500}
	} else {
		response = Data{Msg: string(status), Code: 200}
	}
	js, _ := json.Marshal(response)
	w.Write(js)

}

// UpdateDockerStatusHandler update docker status
func UpdateDockerStatusHandler(w http.ResponseWriter, request *http.Request) {
	/**
	DockerBackend回调函数
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	useraccount := request.FormValue("from")
	dockerstatus := request.FormValue("dockerstatus")

	status, err := strconv.Atoi(dockerstatus)
	if err != nil {
		log.Println("数据类型转化失败")

	}
	utils.UpdateDockerStatus(useraccount, status)

}

// TODO: test
func ComputingClientUploadEncryptedDataHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//
	//t, _ = template.ParseFiles("template/indexcomputer.html")
}

func DataClientUplodFileHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var response Data
	fileName, err := utils.SaveFileToLocal(r)
	if err != nil {
		response = Data{Msg: "文件保存失败", Code: 500}
	} else {
		hash, err := utils.UploadFile(fileName)
		if err != nil {
			response = Data{Msg: "上传文件至ipfs失败", Code: 500}
		} else {
			response = Data{Msg: hash, Code: 200}
		}
	}

	// response
	jsonStr, _ := json.Marshal(response)
	log.Println(string(jsonStr))
	w.Write(jsonStr)
}

func ModelClientUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var response Data
	fileName, err := utils.SaveFileToLocal(r)
	if err != nil {
		response = Data{Msg: "文件保存失败", Code: 500}
	} else {
		hash, err := utils.UploadFile(fileName)
		if err != nil {
			response = Data{Msg: "上传文件至ipfs失败", Code: 500}
		} else {
			response = Data{Msg: hash, Code: 200}
		}
	}

	// response
	jsonStr, _ := json.Marshal(response)
	log.Println(string(jsonStr))
	w.Write(jsonStr)
}

func ComputingClientUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var response Data
	fileName, err := utils.SaveFileToLocal(r)
	if err != nil {
		// response = Data{Msg: "文件保存失败", Code: 500}
		response = Data{Msg: "file save failed", Code: 500}
	} else {
		hash, err := utils.UploadFile(fileName)
		if err != nil {
			// response = Data{Msg: "上传文件至ipfs失败", Code: 500}
			response = Data{Msg: "upload file to ipfs failed", Code: 500}
		} else {
			response = Data{Msg: hash, Code: 200}
		}
	}

	// response
	jsonStr, _ := json.Marshal(response)
	log.Println(string(jsonStr))
	w.Write(jsonStr)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}
	t, err := template.ParseFiles("template/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

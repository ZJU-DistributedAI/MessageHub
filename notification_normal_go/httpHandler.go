package main

import (
	"./abi"
	"./utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
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

func IndexDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	t, _ := template.ParseFiles("template/indexdata.html")
	t.Execute(w, nil)

}

func IndexModelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	t, _ := template.ParseFiles("template/indexmodel.html")
	t.Execute(w, nil)
}

func IndexComputingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	t, _ := template.ParseFiles("template/indexcomputer.html")
	t.Execute(w, nil)
}

func CheckLoginHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	userType := r.PostFormValue("userType")
	var err error
	if username == "dcd" && password == "123456" {
		if userType == "0" {
			_, err = w.Write([]byte("/dataclient/index"))
		} else if userType == "1" {
			_, err = w.Write([]byte("/modelclient/index"))
		} else {
			_, err = w.Write([]byte("/computingclient/index"))
		}
	} else {
		_, err = w.Write([]byte("/user/login"))
	}

	// error handle
	if err != nil {
		panic(err)
	}
}

func DataClientAvaCompHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// login view
	var t *template.Template
	var err error
	t, err = template.ParseFiles("template/data_computing_agree.html")
	if err != nil {
		utils.ErrorPanic(err)
		return
	}
	t.Execute(w, "")
}

func DataClientModelAskHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// login view
	var t *template.Template
	var err error
	t, err = template.ParseFiles("template/data_model_ask.html")
	if err != nil {
		utils.ErrorPanic(err)
		return
	}
	t.Execute(w, "")
}

func ModelClientAvaDataHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// login view
	var t *template.Template
	var err error
	t, err = template.ParseFiles("template/model_available_data.html")
	if err != nil {
		utils.ErrorPanic(err)
		return
	}
	t.Execute(w, "")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// login view
	var t *template.Template
	var err error
	t, err = template.ParseFiles("template/register.html")
	if err != nil {
		utils.ErrorPanic(err)
		return
	}
	t.Execute(w, "")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// login view
	var t *template.Template
	var err error
	t, err = template.ParseFiles("template/home.html")
	if err != nil {
		utils.ErrorPanic(err)
		return
	}
	t.Execute(w, "")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// login view
	var t *template.Template
	var err error
	t, err = template.ParseFiles("template/login.html")
	if err != nil {
		utils.ErrorPanic(err)
		return
	}
	t.Execute(w, "")
}

func CreateWalletPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template

	t, _ = template.ParseFiles("template/createwallet.html")
	t.Execute(w, nil)
}

func DownloadToolPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template

	t, _ = template.ParseFiles("template/download_tool.html")
	t.Execute(w, nil)
}

func ComputingClientTrainPageHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	t, _ := template.ParseFiles("template/computing_train.html")
	t.Execute(w, nil)
}

func CreateWalletHandler(w http.ResponseWriter, r *http.Request) {

	/**
	为该用户创建以太坊钱包
	@params: password string
	*/

	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := r.PostFormValue("password")

	if password == "" {
		data = Data{Msg: "未找到password参数", Code: 201}

	} else {
		// connect and create
		conn := utils.Connect2Eth()
		account := utils.CreateAccount(conn, password)

		if account != "" {
			data = Data{Msg: account, Code: 200}
		} else {
			data = Data{Msg: "创建账户失败", Code: 500}
		}
	}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

func DataClientWalletPageHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	t, _ := template.ParseFiles("template/dataclientwallet.html")
	t.Execute(w, nil)
}

func ModelClientWalletPageHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	t, _ := template.ParseFiles("template/modelclientwallet.html")
	t.Execute(w, nil)
}

func ComputingClientWalletPageHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	t, _ := template.ParseFiles("template/computingclientwallet.html")
	t.Execute(w, nil)
}

func ComputingClientDataAskHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	t, _ := template.ParseFiles("template/computing_data_ask.html")
	t.Execute(w, nil)
}

// TODO: test SendTransaction
func DataClientAddMetaDataHandler(w http.ResponseWriter, request *http.Request) {

	/**
	数据方账户将数据通过以太坊保存到消息服务器
	@Params: from string
	@Params: password string
	@Params: metaDataIpfsHash string
	*/

	// headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	metaDataIpfsHash := request.PostFormValue("metaDataIpfsHash")
	// fmt.Println(password, from, metaDataIpfsHash)

	if password == "" || metaDataIpfsHash == "" {
		data = Data{Msg: "参数不完全", Code: 201}
	} else {
		value := "dadd:" + metaDataIpfsHash
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("数据方上传数据到区块链失败", err)
			data = Data{Msg: "数据方上传数据到区块链失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)

}

// TODO: test
func DataClientMonitorMetaDataHandler(w http.ResponseWriter, request *http.Request) {
	/**
	接受ajax长轮询
	监听是否有模型方提交的metadata交易
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")

	modelClientPullDataReceipt := GetModelClientPullDataReceipt()

	data = Data{Msg: modelClientPullDataReceipt.Metadata + ":" + modelClientPullDataReceipt.From, Code: 200}
	js, _ := json.Marshal(data)
	t.Execute(w, js)
}

// TODO: test SendTransaction
func DataClientAggreeModelClientHandler(w http.ResponseWriter, request *http.Request) {
	/**
	数据方同意模型范改的请求
	@Params: from string
	@Params: password string
	@Params: metadata hash
	*/

	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	metaDataIpfsHash := request.PostFormValue("metaDataIpfsHash")
	modelAddress := request.PostFormValue("modelAddress")

	if password == "" || modelAddress == "" || metaDataIpfsHash == "" {
		data = Data{Msg: "参数不完全", Code: 201}
	} else {
		value := "dagree:" + metaDataIpfsHash + ":" + modelAddress
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")

		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())
		if err != nil {
			log.Println("数据方同意交易创建失败", err)
			data = Data{Msg: "数据方上传数据到区块链失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}
	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

// TODO: test SendTransaction
func ModelClientAskDataHandler(w http.ResponseWriter, request *http.Request) {
	/**
	模型方使用metadata请求数据方的数据
	@param password string
	@param metaDataInfo string
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var data Data

	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	metaDataInfo := request.PostFormValue("metaDataInfo")

	if password == "" || metaDataInfo == ""{
		data = Data{Msg: "参数不完全", Code: 500}
	} else {
		value := "mpull:" + metaDataInfo
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")

		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("模型方使用metadata请求数据方的数据失败", err)
			data = Data{Msg: "模型方使用metadata请求数据方的数据失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}
	// response
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

// TODO: test SendTransaction
func DataClientPushDataToComputingHandler(w http.ResponseWriter, request *http.Request) {
	/**
	数据方将原数据传递给计算方法（未加密）
	@param: password string
	@param: from string
	@param: dataIpfsHash string
	*/

	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	dataIpfsHash := request.PostFormValue("dataIpfsHash")
	modelAddress := request.PostFormValue("modelAddress")
	dataMetadataIpfsHash := request.PostFormValue("dataMetadataIpfsHash")

	if password == "" || dataIpfsHash == "" || dataMetadataIpfsHash == "" || modelAddress == "" {
		data = Data{Msg: "参数不完全", Code: 500}
	} else {
		value := "dpush:" + dataIpfsHash + ":" + modelAddress + ":" + dataMetadataIpfsHash
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("数据方将数据传给运算方交易创建失败", err)
			data = Data{Msg: "数据方将数据传给运算方交易创建失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}
	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

// TODO: test SendTransaction
func DataClientAskComputingHandler(w http.ResponseWriter, request *http.Request) {
	/**
	数据方请求运算方的资源
	@param: password string 数据方密码
	@param: computingHash string 运算方资源的IpfsHash
	@param: from string  数据方的账户
	*/

	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	computingHash := request.PostFormValue("computingHash")
	from := request.PostFormValue("from")
	modelAddress := request.PostFormValue("modelAddress")

	if password == "" || computingHash == "" || modelAddress == "" {
		data = Data{Msg: "参数不完全", Code: 500}
	} else {
		value := "dcomputing:" + computingHash + ":" + modelAddress
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("数据方请求运算方的运算资源失败", err)
			data = Data{Msg: "数据方请求运算方的运算资源失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}
	js, _ := json.Marshal(data)
	w.Write(js)
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

// TODO: test SendTransaction
func DataClientDeleteDataHandler(w http.ResponseWriter, request *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	metadataHash := request.PostFormValue("metadataHash")
	from := request.PostFormValue("from")

	if password == "" || metadataHash == "" {
		data = Data{Msg: "参数不完全", Code: 500}
	} else {
		value := "ddelete:" + metadataHash
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("数据方删除metadDataHash失败", err)
			data = Data{Msg: "数据方删除metadDataHash失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}
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


// TODO: test SendTransaction
func ModelClientUploadModelHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexmodel.html")
	var data Data
	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	modelIpfsHash := request.PostFormValue("modelIpfsHash")
	contractHash := request.PostFormValue("contractHash")

	// fmt.Println(password, from, metaDataIpfsHash)

	if password == "" || modelIpfsHash == "" {
		data = Data{Msg: "参数不完全", Code: 201}
	} else {
		value := "madd:" + modelIpfsHash + ":" + contractHash
		to := common.HexToAddress("")

		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("模型方到区块链失败", err)
			data = Data{Msg: "模型方上传数据到区块链失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}
	// response
	js, _ := json.Marshal(data)
	w.Write(js)

}

// TODO: test
func ModelClientUploadResultHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexmodel.html")

}

// TODO: test
func ComputingClientMonitorDataClientHandler(w http.ResponseWriter, request *http.Request) {
	/**
	ajax长连接
	运算方监听是否有数据方的运算资源请求
	*/
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	dataAskComputingReceipt := GetDataAskComputingReceipt()
	data = Data{Msg: dataAskComputingReceipt.ComputingHash + ":" + dataAskComputingReceipt.From, Code: 200}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

// TODO: test SendTransaction
func ComputingClientAddDataHandler(w http.ResponseWriter, request *http.Request) {
	/**
	运算方将运算资源提交的以太坊
	@param: computingIpfsHash string
	@param: from string
	@param: password string
	*/

	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	computingIpfsHash := request.PostFormValue("computingIpfsHash")
	from := request.PostFormValue("from")
	password := request.PostFormValue("password")

	if computingIpfsHash == "" {
		data = Data{Msg: "参数不完全", Code: 201}
	} else {
		value := "cadd:" + computingIpfsHash
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")

		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("运算方上传运算资源到区块链失败", err)
			data = Data{Msg: "运算方上传运算资源到区块链失败", Code: 500}

		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

// TODO: test SendTransaction
func ComputingClientAggreeRequestHandler(w http.ResponseWriter, request *http.Request) {
	/**
	运算方同意数据方的请求
	发起运算方同意交易让数据方监听到后，上传数据
	@param: from string
	@param: password string
	*/

	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// model address
	var data Data
	from := request.PostFormValue("from")
	password := request.PostFormValue("password")
	dataAddress := request.PostFormValue("dataAddress")
	modelAddress := request.PostFormValue("modelAddress")

	if dataAddress == "" || modelAddress == "" {
		data = Data{Msg: "参数不完全", Code: 201}
	} else {
		value := "caggree:" + dataAddress + ":" + modelAddress
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("运算方的同意交易生成失败", err)
			data = Data{Msg: "运算方的同意交易生成失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

// TODO: test SendTransaction
func ComputingClientDeleteComputingHashHandler(w http.ResponseWriter, request *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	computingHash := request.PostFormValue("computinghash")
	from := request.PostFormValue("from")
	if computingHash == "" {
		data = Data{Msg: "参数不完全", Code: 201}
	} else {
		value := "ddelete:" + computingHash
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value), "0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash, err := utils.SendTransaction(conn, &message, password, context.TODO())

		if err != nil {
			log.Println("运算方删除computingHash失败", err)
			data = Data{Msg: "运算方删除computingHash失败", Code: 500}
		} else {
			data = Data{Msg: txHash, Code: 200}
		}
	}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}

// todo test
func ComputingClientTrainReceiptHandler(w http.ResponseWriter, request *http.Request){
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	_, err := http.Get("http://127.0.0.1:9091/dockerbackend/trainrequest")
	if err != nil {
		log.Println("与容器后端通信失败", err)
		data = Data{Msg: "与容器后端通信失败", Code: 500}
	}else{
		data = Data{Msg: "接受到训练请求", Code: 200}
	}

	//response
	js, _ := json.Marshal(data)
	w.Write(js)
}


// TODO: test
func ComputingClientTrainHandler(w http.ResponseWriter, request *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	/**
	调用容器后端接口训练模型
	*/
	var data Data

	modelIpfsHash := request.PostFormValue("modelIpfsHash")
	dataIpfsHash := request.PostFormValue("dataIpfshash")

	computingfrom := request.PostFormValue("computingfrom")

	uploadPath, directoryPath:= utils.MakeDirectory("train_"+computingfrom)
	if uploadPath != "" {
		utils.DownloadFile(modelIpfsHash, uploadPath+"modelFile.json")
		utils.DownloadFile(dataIpfsHash, uploadPath+"dataFile.json")
		utils.CopyTrainCode(directoryPath)
	}




	_, err := http.Get("http://127.0.0.1:9091/dockerbackend/starttrain?from="+computingfrom+"&directorypath="+directoryPath)
	if err != nil {
		log.Println("运算方删除computingHash失败", err)
		data = Data{Msg: "运算方删除computingHash失败", Code: 500}
	} else {
		result := utils.ReadFile("//root//MachineLearning//parameters.json")
		data = Data{Msg: result, Code: 200}
	}

	// response
	js, _ := json.Marshal(data)
	w.Write(js)
}




func ComputingClientGetDockerStatus(w http.ResponseWriter, request *http.Request){

	/**
		获取docker状态
	 */

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	useraccount := request.PostFormValue("from")

	status := utils.GetDockerStatus(useraccount)

	var data Data

	if status == 0 {
		data = Data{Msg: "获取状态失败", Code:500}
	}else {
		data = Data{Msg: string(status), Code:200}
	}
	js, _ := json.Marshal(data)
	w.Write(js)

}



func UpdateDockerStatusHandler(w http.ResponseWriter, request *http.Request){
	/**
		DockerBackend回调函数
	 */
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	useraccount := request.FormValue("from")
	dockerstatus := request.FormValue("dockerstatus")

	status, err:= strconv.Atoi(dockerstatus)
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
	var data Data
	fileName, err := utils.SaveFileToLocal(r)
	if err != nil {
		data = Data{Msg: "文件保存失败", Code:500}
	} else {
		hash, err := utils.UploadFile(fileName)
		if err != nil {
			data = Data{Msg:"上传文件至ipfs失败", Code: 500}
		} else {
			data = Data{Msg: hash, Code: 200}
		}
	}

	// response
	jsonStr, _ := json.Marshal(data)
	log.Println(string(jsonStr))
	w.Write(jsonStr)
}

func ModelClientUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	fileName, err := utils.SaveFileToLocal(r)
	if err != nil {
		data = Data{Msg: "文件保存失败", Code:500}
	} else {
		hash, err := utils.UploadFile(fileName)
		if err != nil {
			data = Data{Msg:"上传文件至ipfs失败", Code: 500}
		} else {
			data = Data{Msg: hash, Code: 200}
		}
	}

	// response
	jsonStr, _ := json.Marshal(data)
	log.Println(string(jsonStr))
	w.Write(jsonStr)
}


func ComputingClientUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	fileName, err := utils.SaveFileToLocal(r)
	if err != nil {
		data = Data{Msg: "文件保存失败", Code:500}
	} else {
		hash, err := utils.UploadFile(fileName)
		if err != nil {
			data = Data{Msg:"上传文件至ipfs失败", Code: 500}
		} else {
			data = Data{Msg: hash, Code: 200}
		}
	}

	// response
	jsonStr, _ := json.Marshal(data)
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



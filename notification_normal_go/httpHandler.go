package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"./utils"
	"net/rpc"

)

type MetaDataList struct{
	MetaList []string
}

type ComputingList struct{
	Computings []string
}

type Data struct{
	Msg string  `json:"msg"`
	Code int	`json:"code"`
}

type Handlers struct{
	conn *rpc.Client
}


// TODO: test
func matchData(splits []string){
	//获得所有的数据描述计算数据定价
	conn:=utils.Connect2Redis()
	defer conn.Close()

	//TODO 通过计算DQ来获得最佳数据,这里暂时用result[0]来代表
	metaData:=splits[1]

	//请求数据方询问是否同意吧最佳数据的Hash给模型方
	resp,err:=http.Get("url?data_hash="+metaData)
	if err!=nil{
		log.Fatal("请求数据方出错：%v",err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取response数据失败：%v",err)
	}

	//TODO 万一最佳数据的提供方不允许提供最佳数据则需要一个机制来选择不是最佳的数据
	if(string(body)=="ok"){
		//fmt.Fprint(w,bestData)
		fmt.Println("数据方同意并把真实数据Hash存入智能合约")
	}else{
		fmt.Println("sorry the no best data")
	}

}

// TODO: test
func matchComputing(splits []string){
	//获得所有的数据描述计算数据定价
	conn:=utils.Connect2Redis()
	defer conn.Close()

	computingMetaHash:=splits[1]

	//请求运算方询问是否同意运算资源Hash给区块链
	resp,err:=http.Get("url?computing_meta_hash="+computingMetaHash)
	if err!=nil{
		log.Fatal("请求运算方出错：%v",err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取response数据失败：%v",err)
	}

	//TODO 万一最佳数据的提供方不允许提供最佳数据则需要一个机制来选择不是最佳的数据
	if(string(body)=="ok"){
		fmt.Println("运算方同意并把运算资源Hash存入以太坊")
	}else{
		fmt.Println("sorry the no best data")
	}
}

//response.setHeader("Access-Control-Allow-Origin", "*");
//response.setHeader("Access-Control-Allow-Method", "POST,GET");


// TODO: test
//TODO 返回数据描述信息列表，让ModelClient挑选最佳数据
func ListMetaData(w http.ResponseWriter,r *http.Request){
	//这里我们用json的格式返回所有data描述信息

	conn:=utils.Connect2Redis()
	defer conn.Close()

	result:=utils.SmembersFromRedis(conn,"metadata")

	metaDataList:=MetaDataList{
		MetaList:result,
	}

	js,_:=json.Marshal(metaDataList)

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	w.Write([]byte(string(js)))

}

// TODO: test
//TODO 返回运算资源描述信息列表，让DataClient挑选最佳运算资源
func ListComputing(w http.ResponseWriter,r *http.Request){
	//这里我们用json的格式返回所有computing描述信息
	conn:=utils.Connect2Redis()
	defer conn.Close()

	result:=utils.SmembersFromRedis(conn,"computing")

	metaDataList:=MetaDataList{
		MetaList:result,
	}

	js,_:=json.Marshal(metaDataList)

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	w.Write([]byte(string(js)))

}

// TODO: test
func ListAskedMetaData(w http.ResponseWriter,r *http.Request){

	conn:=utils.Connect2Redis()
	defer conn.Close()

	result:=utils.SmembersFromRedis(conn,r.FormValue("address"))

	metaDataList:=MetaDataList{
		MetaList:result,
	}

	js,_:=json.Marshal(metaDataList)

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	w.Write([]byte(string(js)))


}

// TODO: test
func ListAskedComputing(w http.ResponseWriter,r *http.Request){
	conn:=utils.Connect2Redis()
	defer conn.Close()

	result:=utils.SmembersFromRedis(conn,r.FormValue("address"))

	computingList:=ComputingList{
		Computings:result,
	}

	js,_:=json.Marshal(computingList)

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	w.Write([]byte(string(js)))
}

func IndexHandler(w http.ResponseWriter,r *http.Request){

	// view
	t, err := template.ParseFiles("template/index.html")
	if err!= nil {
		utils.ErrorPanic(err)
	}

	// header
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	// body
	t.Execute(w, nil)
}

func IndexDataHandle(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	t, _ := template.ParseFiles("template/indexdata.html")
	t.Execute(w, nil)

}

func IndexModelHandle(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	t, _ := template.ParseFiles("template/indexmodel.html")
	t.Execute(w, nil)
}

func IndexComputingHandle(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	t, _ := template.ParseFiles("template/indexcomputer.html")
	t.Execute(w, nil)
}

func CheckLoginHandler(w http.ResponseWriter,r *http.Request){
	// header
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	userType := r.PostFormValue("userType")
	var err error
	if username == "dcd" && password == "123456" {
		if userType == "0"{
			_, err = w.Write([]byte("/dataclient/index"))
		}else if userType == "1"{
			_, err = w.Write([]byte("/modelclient/index"))
		}else{
			_, err = w.Write([]byte("/computingclient/index"))
		}
	}else{
		_, err = w.Write([]byte("/user/login"))
	}

	// error handle
	if err != nil{
		panic(err)
	}
}

func LoginHandler(w http.ResponseWriter,r *http.Request){
	// header
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// login view
	var t *template.Template
	var err error
	t, err = template.ParseFiles("template/login.html")
	if err!=nil {
		utils.ErrorPanic(err)
		return
	}
	t.Execute(w, "")
}

func CreateWalletPageHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template

	t, _ = template.ParseFiles("template/createwallet.html")
	t.Execute(w, nil)

}

func CreateWalletHandler(w http.ResponseWriter,r *http.Request){

	/**
		为该用户创建以太坊钱包
		@params: password string
	 */

 	// header
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password:=r.PostFormValue("password")

	if password == ""{
		data = Data{Msg:"未找到password参数", Code:500}
	} else {
		// connect and create
		conn := utils.Connect2Eth()
		account := utils.CreateAccount(conn,password)

		if account != ""{
			data = Data{Msg:account, Code:200}
		}else{
			data = Data{Msg:"创建账户失败", Code:500}
		}
	}

	// response
	js,_:= json.Marshal(data)
	w.Write(js)
}

func DataClientWalletPageHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	t, _ := template.ParseFiles("template/dataclientwallet.html")
	t.Execute(w, nil)
}

// TODO: test SendTransaction
func DataClientAddDataHandler(w http.ResponseWriter, request *http.Request) {

	/**
		数据方账户将数据通过以太坊保存到消息服务器

		@Params: from string
		@Params: password string
		@Params: dataIpfsHash string
	 */

	// headers
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	metaDataIpfsHash := request.PostFormValue("metaDataIpfsHash")
	// fmt.Println(password, from, metaDataIpfsHash)

	if password == "" || metaDataIpfsHash == ""{
		data = Data{Msg:"参数不完全", Code:500}
	} else {
		value := "dadd:"+metaDataIpfsHash
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")
		conn := utils.Connect2Eth()
		txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

		if  err!= nil{
			log.Println("数据方上传数据到区块链失败", err)
			data = Data{Msg:"数据方上传数据到区块链失败", Code:500}
		}else{
			data = Data{Msg:txHash, Code:200}
		}
	}

	// response
	js,_:=json.Marshal(data)
	w.Write(js)

}

// TODO: test
func DataClientMonitorMetaDataHandler(w http.ResponseWriter, request *http.Request){
	/**
		接受ajax长轮询
		监听是否有模型方提交的metadata交易
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")

	modelClientPullDataReceipt := GetModelClientPullDataReceipt()

	data = Data{Msg: modelClientPullDataReceipt.Metadata+":"+modelClientPullDataReceipt.From, Code:200}
	js, _ := json.Marshal(data)
	t.Execute(w, js)
}

// TODO: test SendTransaction
func DataClientAggreeModelClientHandler(w http.ResponseWriter, request *http.Request){
	/**
		数据方同意模型范改的请求
		@Params: from string
		@Params: password string
		@Params: metadata hash
	 */

	// header
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	// handle
	var data Data
	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	metaDataIpfsHash := request.PostFormValue("metaDataIpfsHash")

	if password == "" {
		data = Data{Msg:"参数不完全", Code:500}
	} else {
		value := "dagree:" + metaDataIpfsHash
		to := common.HexToAddress("")
		//发起交易到以太坊
		message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
			"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")

		conn := utils.Connect2Eth()
		txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())
		if  err!= nil{
			log.Println("数据方同意交易创建失败", err)
			data = Data{Msg:"数据方上传数据到区块链失败", Code:500}
		}else{
			data = Data{Msg:txHash, Code:200}
		}
	}
	// response
	js,_:=json.Marshal(data)
	w.Write(js)
}

// TODO: test
func ModelClientAskDataHandler(w http.ResponseWriter, request *http.Request){
	/**
		模型方使用metadata请求数据方的数据
		@param password string
		@param metaDataInfo string
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")

	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	metaDataInfo := request.PostFormValue("metaDataInfo")

	if password == "" || metaDataInfo == "" {
		data = Data{Msg:"参数不完全", Code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

	value := "mpull:"+metaDataInfo
	to := common.HexToAddress("")
	//发起交易到以太坊
	message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
		"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")

	conn := utils.Connect2Eth()
	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

	if  err!= nil{
		log.Fatal("模型方使用metadata请求数据方的数据失败",err)
		data = Data{Msg:"模型方使用metadata请求数据方的数据失败", Code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}else{
		data = Data{Msg:txHash, Code:200}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

}

// TODO: test
func ModelClientMonitorDataClientResultHandler(w http.ResponseWriter, request *http.Request){

	/**
		模型方监听数据方的响应（是否同意模型方的数据请求）
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexmodel.html")

	dataClientIsAggreeReceipt:=GetDataClientIsAggreeReceipt()

	data = Data{Msg: dataClientIsAggreeReceipt.Message+":"+dataClientIsAggreeReceipt.From, Code:200}
	js, _ := json.Marshal(data)
	t.Execute(w, js)


}

// TODO: test
func DataClientPushDataToComputingHandler(w http.ResponseWriter, request *http.Request) {
	/**
		数据方将原数据传递给计算方法（未加密）
		@param: password string
		@param: from string
		@param: dataIpfsHash string
	 */

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")
	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	dataIpfsHash := request.PostFormValue("dataIpfsHash")

	value := "dpush:"+dataIpfsHash
	to := common.HexToAddress("")
	//发起交易到以太坊
	message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
		"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")
	conn := utils.Connect2Eth()
	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

	if  err!= nil{
		log.Fatal("数据方将数据传给运算方交易创建失败",err)
		data = Data{Msg:"数据方将数据传给运算方交易创建失败", Code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}else{
		data = Data{Msg:txHash, Code:200}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

}

// TODO: test
func DataClientAskComputingHandler(w http.ResponseWriter, request *http.Request) {
	/**
		数据方请求运算方的资源
		@param: password string 数据方密码
		@param: computingHash string 运算方资源的IpfsHash
		@param: from string  数据方的账户
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")
	password := request.PostFormValue("password")
	computingHash := request.PostFormValue("computingHash")
	from := request.PostFormValue("from")

	value := "dcomputing:"+computingHash
	to := common.HexToAddress("")
	//发起交易到以太坊
	message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
		"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")
	conn := utils.Connect2Eth()
	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

	if  err!= nil{
		log.Fatal("数据方请求运算方的运算资源失败",err)
		data = Data{Msg:"数据方请求运算方的运算资源失败", Code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}else{
		data = Data{Msg:txHash, Code:200}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}
}

// TODO: test
func DataClientMonitorComputingAggreeHandler(w http.ResponseWriter, request *http.Request){

	/**
		ajax长连接
		数据方监听是否有运算方的同意交易
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")

	dataClientMonitorComputingAggreeReceipt := GetDataClientMonitorComputingAggreeReceipt()

	data = Data{Msg: dataClientMonitorComputingAggreeReceipt.ComputingHash+":"+dataClientMonitorComputingAggreeReceipt.From,
		Code:200}
	js, _ := json.Marshal(data)
	t.Execute(w, js)

}

// TODO: test
func DataClientDeleteDataHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")


	//password := request.PostFormValue("password")
	//from := request.PostFormValue("from")
	//metaDataIpfsHash := request.PostFormValue("metaDataIpfsHash")

	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexdata.html")


}

// TODO: test
func ModelClientCreateContractHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexmodel.html")

}


// TODO: test
func ModelClientUploadModelHandler(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexmodel.html")

}

// TODO: test
func ModelClientUploadResultHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexmodel.html")

}

// TODO: test
func ComputingClientMonitorDataClientHandler(w http.ResponseWriter, request *http.Request){
	/**
		ajax长连接
		运算方监听是否有数据方的运算资源请求
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexcomputer.html")

	dataAskComputingReceipt := GetDataAskComputingReceipt()

	data = Data{Msg: dataAskComputingReceipt.ComputingHash+":"+dataAskComputingReceipt.From, Code:200}

	js, _ := json.Marshal(data)

	t.Execute(w, js)


}

// TODO: test
func ComputingClientAddDataHandler(w http.ResponseWriter, request *http.Request) {
	/**
		运算方将运算资源提交的以太坊
		@param: computingIpfsHash string
		@param: from string
		@param: password string
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexcomputer.html")

	computingIpfsHash := request.PostFormValue("computingIpfsHash")
	from := request.PostFormValue("from")
	password := request.PostFormValue("password")

	value := "cadd:" + computingIpfsHash
	to := common.HexToAddress("")
	//发起交易到以太坊
	message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
		"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")

	conn := utils.Connect2Eth()
	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

	if  err!= nil{
		log.Fatal("运算方上传运算资源到区块链失败",err)
		data = Data{Msg:"运算方上传运算资源到区块链失败", Code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}else{
		data = Data{Msg:txHash, Code:200}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

}

// TODO: test
func ComputingClientAggreeRequestHandler(w http.ResponseWriter, request *http.Request) {
	/**
		运算方同意数据方的请求
		发起运算方同意交易让数据方监听到后，上传数据
		@param: from string
		@param: password string
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexcomputer.html")

	from := request.PostFormValue("from")
	password := request.PostFormValue("password")

	value := "caggree:"
	to := common.HexToAddress("")
	//发起交易到以太坊
	message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
		"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")

	conn := utils.Connect2Eth()
	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

	if  err!= nil{
		log.Fatal("运算方的同意交易生成失败",err)
		data = Data{Msg:"运算方的同意交易生成失败", Code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}else{
		data = Data{Msg:txHash, Code:200}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}
}

// TODO: test
func ComputingClientDeleteDataHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexcomputer.html")
}

// TODO: test
func ComputingClientTrainHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//t, _ = template.ParseFiles("template/indexcomputer.html")
}

// TODO: test
func ComputingClientUploadEncryptedDataHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	//var t *template.Template
	//var data Data
	//
	//t, _ = template.ParseFiles("template/indexcomputer.html")
}

// TODO: test
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	t, err := template.ParseFiles("template/404.html")
	if (err != nil) {
		log.Println(err)
	}
	t.Execute(w, nil)
}








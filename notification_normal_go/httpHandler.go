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
	"context"
)

type MetaDataList struct{
	MetaList []string
}

type ComputingList struct{
	Computings []string
}

type Data struct{
	msg string
	code int
}

type Handlers struct{
	conn *rpc.Client
}



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

	t, err := template.ParseFiles("template/index.html")
	if err!= nil {
		utils.ErrorPanic(err)
	}

	//js,_:=json.Marshal(serverTest)

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	t.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	userType := r.PostFormValue("userType")

	var t *template.Template
	var err error
	if username == "dcd" && password == "123456" {
		if userType == "0"{
			t, err = template.ParseFiles("template/indexdata.html")
		}else if userType == "1"{
			t, err = template.ParseFiles("template/indexmodel.html")
		}else{
			t, err = template.ParseFiles("template/indexcomputer.html")
		}
		if err!=nil {
			utils.ErrorPanic(err)
			return
		}
		data := Data{msg:"", code:200}
		t.Execute(w, data)

	}else{
		t, err = template.ParseFiles("template/login.html")
		if err!=nil {
			utils.ErrorPanic(err)
			return
		}

		data := Data{msg:"账号或密码错误", code:200}
		t.Execute(w, data)
	}






}

func CreateWalletHandler(w http.ResponseWriter,r *http.Request){

	/**
<<<<<<< HEAD
		为该用户创建以太坊钱包
=======
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
		@params: password string
	 */

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template

	t, _ = template.ParseFiles("template/createwallet.html")

	var data Data
	password:=r.PostFormValue("password")

	if password == ""{
		data = Data{msg:"未找到password参数", code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

	conn := utils.Connect2Eth()
	account := utils.CreateAccount(conn,password)

	if account != ""{
		data = Data{msg:account, code:200}
	}else{
		data = Data{msg:"创建账户失败", code:500}
	}
	js,_:=json.Marshal(data)

	t.Execute(w, js)

}

func DataClientAddDataHandler(w http.ResponseWriter, request *http.Request) {

	/**
<<<<<<< HEAD
		数据方账户将数据通过以太坊保存到消息服务器
=======
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
		@Params: from string
		@Params: password string
		@Params: dataIpfsHash string
	 */

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data

	t, _ = template.ParseFiles("template/indexdata.html")

	password := request.PostFormValue("password")
	from := request.PostFormValue("from")
	metaDataIpfsHash := request.PostFormValue("metaDataIpfsHash")

	if password == "" || metaDataIpfsHash == ""{
		data = Data{msg:"参数不完全", code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

	value := "dadd:"+metaDataIpfsHash
	to := common.HexToAddress("")
	//发起交易到以太坊
	message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
		"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")
<<<<<<< HEAD
	conn := utils.Connect2Eth()
	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

	if  err!= nil{
		log.Fatal("数据方上传数据到区块链失败",err)
		data = Data{msg:"数据方上传数据到区块链失败", code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}else{
		data = Data{msg:txHash, code:200}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

}
=======

	conn := utils.Connect2Eth()

	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())
	if err!= nil{
		utils.ErrorPanic(err)
		data = Data{msg:"提交交易出错", code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}
	data = Data{msg:txHash, code:500}
	js,_:=json.Marshal(data)
	t.Execute(w, js)

}

//获取模型方传过来的metedata响应给数据方
func ModelClientMetaDataInfoHandler(w http.ResponseWriter, request *http.Request){
	/**
		@param password string
		@param metaDataInfo string
	 */
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")

	password := request.PostFormValue("password")
	metaDataInfo := request.PostFormValue("metaDataInfo")

	if password == "" || metaDataInfo == "" {
		data = Data{msg:"参数不完全", code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}

	//生成交易提交到区块链
}



func DataClientAggreeRequestHandler(w http.ResponseWriter, request *http.Request) {
	//
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")

	password := request.PostFormValue("password")
	metaDataInfo := request.PostFormValue("metaDataInfo")





>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b


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

	data = Data{msg: modelClientPullDataReceipt.Metadata+":"+modelClientPullDataReceipt.From, code:200}
	js, _ := json.Marshal(data)
	t.Execute(w, js)
}
<<<<<<< HEAD



func DataClientAggreeRequestHandler(w http.ResponseWriter, request *http.Request) {
	/**
		数据方同意模型方的请求将原数据传递给计算方法（不加密实现）
	 */
=======
func DataClientAskComputingHandler(w http.ResponseWriter, request *http.Request) {
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")

<<<<<<< HEAD
	password := request.PostFormValue("password")
	dataIpfsHash := request.PostFormValue("dataIpfsHash")
	from := request.PostFormValue("from")

	value := "dagree:"+dataIpfsHash
	to := common.HexToAddress("")
	//发起交易到以太坊
	message := utils.NewMessage(common.HexToAddress(from), &to, "0x10",
		"0x"+utils.EncryptTransactionInput(value),"0x295f05", "0x77359400")
	conn := utils.Connect2Eth()
	txHash,err := utils.SendTransaction(conn, &message, password, context.TODO())

	if  err!= nil{
		log.Fatal("数据方上传数据到区块链失败",err)
		data = Data{msg:"数据方上传数据到区块链失败", code:500}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}else{
		data = Data{msg:txHash, code:200}
		js,_:=json.Marshal(data)
		t.Execute(w, js)
	}




}
func DataClientAskComputingHandler(w http.ResponseWriter, request *http.Request) {
=======

}
func DataClientDeleteDataHandler(w http.ResponseWriter, request *http.Request) {
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexdata.html")


}
<<<<<<< HEAD
func DataClientDeleteDataHandler(w http.ResponseWriter, request *http.Request) {
=======
func ModelClientAskDataHandler(w http.ResponseWriter, request *http.Request) {
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
<<<<<<< HEAD
	t, _ = template.ParseFiles("template/indexdata.html")


}
func ModelClientAskDataHandler(w http.ResponseWriter, request *http.Request) {
=======
	t, _ = template.ParseFiles("template/indexmodel.html")

}
func ModelClientCreateContractHandler(w http.ResponseWriter, request *http.Request) {
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexmodel.html")

}
<<<<<<< HEAD
func ModelClientCreateContractHandler(w http.ResponseWriter, request *http.Request) {
=======
func ModelClientUploadModelHandler(w http.ResponseWriter, request *http.Request) {
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexmodel.html")

}
<<<<<<< HEAD
func ModelClientUploadModelHandler(w http.ResponseWriter, request *http.Request) {
=======
func ModelClientUploadResultHandler(w http.ResponseWriter, request *http.Request) {
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexmodel.html")

}
<<<<<<< HEAD
func ModelClientUploadResultHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexmodel.html")

}
=======
>>>>>>> e9a91867ca374c278fae2b73b523e665656e346b
func ComputingClientAddDataHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexcomputer.html")

}
func ComputingClientAggreeRequestHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexcomputer.html")

}
func ComputingClientDeleteDataHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexcomputer.html")

}
func ComputingClientTrainHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data
	t, _ = template.ParseFiles("template/indexcomputer.html")

}
func ComputingClientUploadEncryptedDataHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")
	var t *template.Template
	var data Data

	t, _ = template.ParseFiles("template/indexcomputer.html")

}






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








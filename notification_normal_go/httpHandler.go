package main

import (
	"encoding/json"
	"fmt"
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
		utils.ErrorInfo(err)
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
			utils.ErrorInfo(err)
			return
		}
		data := Data{msg:"", code:200}
		t.Execute(w, data)

	}else{
		t, err = template.ParseFiles("template/login.html")
		if err!=nil {
			utils.ErrorInfo(err)
			return
		}

		data := Data{msg:"账号或密码错误", code:200}
		t.Execute(w, data)
	}






}

func CreateWalletHandler(w http.ResponseWriter,r *http.Request){

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

func DataClientAddDataHandler(writer http.ResponseWriter, request *http.Request) {

}
func DataClientAggreeRequestHandler(writer http.ResponseWriter, request *http.Request) {

}
func DataClientAskComputingHandler(writer http.ResponseWriter, request *http.Request) {

}
func DataClientDeleteDataHandler(writer http.ResponseWriter, request *http.Request) {

}
func ModelClientAskDataHandler(writer http.ResponseWriter, request *http.Request) {

}
func ModelClientCreateContractHandler(writer http.ResponseWriter, request *http.Request) {

}
func ModelClientUploadModelHandler(writer http.ResponseWriter, request *http.Request) {

}
func ModelClientUploadResultHandler(writer http.ResponseWriter, request *http.Request) {

}
func ComputingClientAddDataHandler(writer http.ResponseWriter, request *http.Request) {

}
func ComputingClientAggreeRequestHandler(writer http.ResponseWriter, request *http.Request) {

}
func ComputingClientDeleteDataHandler(writer http.ResponseWriter, request *http.Request) {

}
func ComputingClientTrainHandler(writer http.ResponseWriter, request *http.Request) {

}
func ComputingClientUploadEncryptedDataHandler(writer http.ResponseWriter, request *http.Request) {

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








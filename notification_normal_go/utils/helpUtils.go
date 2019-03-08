package utils

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)


var myMap map[string]string

func DecryptTransactionInput(input string)(string){

	input=input[2:]

	test, _ := hex.DecodeString(input)

	return string(test)

}

func EncryptTransactionInput(input string)string{

	test:=hex.EncodeToString([]byte(input))

	return test
}

func DownloadFile(hash string, filename string) {
	myhash := strings.Split(hash, "\000")
	finalhash := myhash[0]
	cmd := exec.Command("ipfs", "get", finalhash, "-o="+filename)
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
}

func MakeDirectory(dirname string)(path string){

	cmd := exec.Command("mkdir","-p", "//root//"+dirname)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return "//root//"+dirname+"//"
}


func ReadFile(filepath string)(string){

	f, err := os.Open(filepath)

	if err != nil{
		log.Printf("打开文件失败: ", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	b, _, err := r.ReadLine()

	if err != io.EOF {
		log.Println("读取文件失败: ",err)
	}

	return string(b)

}


func InitConfig(path string) map[string]string {
	//初始化

	if myMap != nil {
		return myMap
	}

	myMap = make(map[string]string)
	//打开文件指定目录，返回一个文件f和错误信息
	f, err := os.Open(path)

	//异常处理 以及确保函数结尾关闭文件流
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//创建一个输出流向该文件的缓冲流*Reader
	r := bufio.NewReader(f)
	for {
		//读取，返回[]byte 单行切片给b
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//去除单行属性两端的空格
		s := strings.TrimSpace(string(b))
		//fmt.Println(s)

		//判断等号=在该行的位置
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		//取得等号左边的key值，判断是否为空
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}

		//取得等号右边的value值，判断是否为空
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		//这样就成功吧配置文件里的属性key=value对，成功载入到内存中c对象里
		myMap[key] = value
	}
	return myMap
}



//
//func main(){
//
//	encryption:=encryptTransactionInput("abc")
//	fmt.Println("encrypt: "+encryption)
//	fmt.Println("decrypt: "+decryptTransactionInput(encryption))
//}

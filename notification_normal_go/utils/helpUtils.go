package utils

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	WINdOSPATH     = "D://distribute_ai_users//"
	WINDOSCODEPATH = "D://MNISTCode//."

	LINUXPATH     = "//root//distribute_ai_users//"
	LINUXCODEPATH = "//root//MNISTCode//."
)

var myMap map[string]string

func DecryptTransactionInput(input string) string {

	input = input[2:]

	test, _ := hex.DecodeString(input)

	return string(test)

}

func EncryptTransactionInput(input string) string {

	test := hex.EncodeToString([]byte(input))

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

func MakeDirectory(dirname string) (userPath string, directortPath string) {

	cmd := exec.Command("mkdir", "-p", LINUXPATH+dirname)
	err := cmd.Run()
	if err != nil {
		fmt.Println("create user directory fail: ", err)
		return "", ""
	}
	cmd = exec.Command("mkdir", "-p", LINUXPATH+dirname+"//upload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("create upload directory fail: ", err)
		return "", ""
	}
	return LINUXPATH + dirname + "//upload//", LINUXPATH + dirname + "//"
}

func CopyTrainCode(directoryPath string) {

	cmd := exec.Command("cp", "-r", LINUXCODEPATH, directoryPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("copy train code fial: ", err)
	}

}

// get unique timestamp string
var time_mutex sync.Mutex

func getTimeStamp() string {
	time_mutex.Lock()
	time_stamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	time.Sleep(1)
	time_mutex.Unlock()
	return time_stamp
}

func UploadFile(filename string) (string, error) {
	// run ipfs add -r filename
	cmd := exec.Command("ipfs", "add", "-r", filename)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// run: ipfs daemon
		cmdIpfsDaemon := exec.Command("ipfs", "daemon")
		cmdIpfsDaemon.Run()
		// try again
		cmd := exec.Command("ipfs", "add", "-r", filename)
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			log.Print(err)
			os.Remove(filename)
			return "", err
		}
	}
	out_str := strings.Split(out.String(), " ")
	hash := out_str[1]
	os.Remove(filename)
	return hash, nil
}

func SaveFileToLocal(r *http.Request) (string, error) {
	fileName := "file_" + getTimeStamp()
	// 根据字段名获取表单文件
	formFile, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return "", err
	}
	defer formFile.Close()

	// 创建保存文件
	destFile, err := os.Create("upload_file/" + fileName)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return "", err
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return "", err
	}
	return "./upload_file/" + fileName, nil
}

func ReadFile(filepath string) string {

	f, err := os.Open(filepath)

	if err != nil {
		log.Printf("打开文件失败: ", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	b, _, err := r.ReadLine()

	if err != io.EOF {
		log.Println("读取文件失败: ", err)
	}

	return string(b)

}

// ------ start GetFederating ------
type weightsStruct struct {
	num int
	W1  [][]float64
	b1  [][]float64
	W2  [][]float64
	b2  [][]float64
	W3  [][]float64
	b3  [][]float64
}

func GetFederateLearningResult() (result string) {
	// 模型文件 参数 路径
	modelfiles := []string{"./parameters.json", "./parameters.json"}
	// var allWeights weightsStruct = weightsStruct{}
	allWeights := new(weightsStruct)
	// fmt.Println(allWeights.W1)
	countNum := 0
	for index := 0; index < len(modelfiles); index++ {
		tempWeights := getModelFileWeights(modelfiles[index])
		if index == 0 {
			allWeights.num = tempWeights.num
			countNum = tempWeights.num
			allWeights.W1 = tempWeights.W1
			allWeights.b1 = tempWeights.b1
			allWeights.W2 = tempWeights.W2
			allWeights.b2 = tempWeights.b2
			allWeights.W3 = tempWeights.W3
			allWeights.b3 = tempWeights.b3
			// fmt.Println(tempWeights.W2)
			break
		}

		allWeights.W1 = matrixAdd(numMul(allWeights.W1, allWeights.num), numMul(tempWeights.W1, tempWeights.num))
		allWeights.b1 = matrixAdd(numMul(allWeights.b1, allWeights.num), numMul(tempWeights.b1, tempWeights.num))
		allWeights.W2 = matrixAdd(numMul(allWeights.W2, allWeights.num), numMul(tempWeights.W2, tempWeights.num))
		allWeights.b2 = matrixAdd(numMul(allWeights.b2, allWeights.num), numMul(tempWeights.b2, tempWeights.num))
		allWeights.W3 = matrixAdd(numMul(allWeights.W3, allWeights.num), numMul(tempWeights.W3, tempWeights.num))
		allWeights.b3 = matrixAdd(numMul(allWeights.b3, allWeights.num), numMul(tempWeights.b3, tempWeights.num))
		countNum += tempWeights.num
	}

	allWeights.W1 = matrixDiv(allWeights.W1, countNum)
	allWeights.b1 = matrixDiv(allWeights.b1, countNum)
	allWeights.W2 = matrixDiv(allWeights.W2, countNum)
	allWeights.b2 = matrixDiv(allWeights.b2, countNum)
	allWeights.W3 = matrixDiv(allWeights.W3, countNum)
	allWeights.b3 = matrixDiv(allWeights.b3, countNum)
	// fmt.Println(allWeights.W1)
	// allWeights := getModelFileWeights("./parameters.json")
	// fmt.Println(len(allWeights.W1))
	str, _ := json.Marshal(allWeights)
	write2json(str)
	fmt.Printf("%s\n", str)
	return "ss"
}

func getModelFileWeights(modelfile string) (weights weightsStruct) {
	JsonParse := NewJsonStruct()
	weights = weightsStruct{}
	JsonParse.Load(modelfile, &weights)
	// fmt.Println(weights)
	return weights
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

// matrix add
func matrixAdd(a [][]float64, b [][]float64) (c [][]float64) {
	for i := 0; i < len(a); i++ {
		t := []float64{}
		for j := 0; j < len(a[0]); j++ {
			temp := a[i][j] + b[i][j]
			t = append(t, temp)
		}
		c = append(c, t)
	}
	return c
}

// matrix div
func matrixDiv(a [][]float64, num int) (c [][]float64) {
	if num == 0 {
		return
	}
	// fmt.Println(len(a))
	// fmt.Println(len(a[0]))
	for i := 0; i < len(a); i++ {
		t := []float64{}
		for j := 0; j < len(a[0]); j++ {
			temp := a[i][j] / float64(num)
			t = append(t, temp)
		}
		c = append(c, t)
	}
	return c
}

func numMul(a [][]float64, num int) (c [][]float64) {
	if num == 0 {
		return
	}
	for i := 0; i < len(a); i++ {
		t := []float64{}
		for j := 0; j < len(a[0]); j++ {
			temp := a[i][j] * float64(num)
			t = append(t, temp)
		}
		c = append(c, t)
	}
	return c
}
func write2json(data []byte) {
	fp, err := os.OpenFile("allWeights.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

// ------ end GetFederating ------

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

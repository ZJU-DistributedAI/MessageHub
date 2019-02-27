package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"log"
	"os"
)


func init() {

	raven.SetDSN("http://e8f71faeb2d043fd96058c57a481434c:b7afc614ef7149c997cb9b316d7f2eaf@212.64.85.208:9000/3")

}


func main(){

	f, err := os.Open("abc.txt")

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		raven.CapturePanic(func(){
			log.Panic(err)
		},nil)

	}

	fmt.Println("file: ",f)


}





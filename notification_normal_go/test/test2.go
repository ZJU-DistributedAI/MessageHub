package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"log"
	"os"
)


func init() {

	raven.SetDSN("http://1186ef58d90f4b1992df3673c6c85857:f6a3101d6fdc49d594f7f1d89a4dd6e5@212.64.85.208:9000/4")

}


func main(){

	f, err := os.Open("123.txt")

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}

	fmt.Println("file: ",f)


}





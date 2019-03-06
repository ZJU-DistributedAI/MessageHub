package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"log"
	"os"
)


func init() {

	raven.SetDSN("http://1b94cad9e5df4ef2b743b5fc43ad3e6f:4808eab3818b4ef1bc2f068d4278d207@212.64.85.208:9000/2")

}


func main(){

	f, err := os.Open("123.txt")

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}

	fmt.Println("file: ",f)


}





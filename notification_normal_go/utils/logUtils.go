package utils

import (
	"github.com/getsentry/raven-go"
	"log"
)


func DebugInfo(info string){
	raven.CaptureMessageAndWait(info, map[string]string{"category": "[DEBUG]"})
}


func WarnningInfo(info string){
	raven.CaptureMessageAndWait(info, map[string]string{"category": "[WARNING]"})
}

func ErrorPanic(e error){
	raven.CaptureError(e, nil)
	raven.CapturePanic(func(){
		log.Panic(e)
	}, nil)
}




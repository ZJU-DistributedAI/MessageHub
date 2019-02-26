package utils

import "log"


func DebugInfo(info string){
	log.Println(" [DEBUG]: ", info)
}


func WarnningInfo(info string){
	log.Println(" [INFO]: ", info)
}

func ErrorInfo(e error){
	log.Printf(" [ERROR]: %v\n", e)
}




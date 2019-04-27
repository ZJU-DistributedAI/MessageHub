package main

import (
	"fmt"
	"os/exec"
)




func main(){

	cmd := exec.Command( "mkdir", "-p", "D://distribute_ai_users//abc")
	err := cmd.Run()
	if err!= nil{
		fmt.Println(err)
	}

}





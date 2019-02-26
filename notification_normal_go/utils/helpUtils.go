package utils

import (
	"encoding/hex"
)

func DecryptTransactionInput(input string)(string){

	input=input[2:]

	test, _ := hex.DecodeString(input)

	return string(test)

}

func EncryptTransactionInput(input string)string{



	test:=hex.EncodeToString([]byte(input))

	return test
}

//
//func main(){
//
//	encryption:=encryptTransactionInput("abc")
//	fmt.Println("encrypt: "+encryption)
//	fmt.Println("decrypt: "+decryptTransactionInput(encryption))
//}

package main

import (
	"GoMing/utils"
	"fmt"
	"log"
)

func main() {
	base64 := utils.Base64{}
	encodeStr := base64.EncodeStr("hello")
	fmt.Println("编码:", encodeStr)
	decodeStr, err := base64.DecodeStr(encodeStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("解码:", decodeStr)
}

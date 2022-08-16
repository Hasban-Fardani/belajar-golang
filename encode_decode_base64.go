package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var str string
	fmt.Print("ketik sesuatu: ")
	fmt.Scanln(&str)
	var encoding_str = base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("encoding:", encoding_str)
	var decoding_str, _ = base64.StdEncoding.DecodeString(encoding_str)
	fmt.Println("decode  :", string(decoding_str))
}

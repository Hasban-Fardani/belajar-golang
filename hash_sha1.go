package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func hash_salting(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: %s, salt: %d", text, salt)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)
	return string(encrypted), salt
}

func main() {
	var text string
	fmt.Print("ketik sesuatu: ")
	fmt.Scanln(&text)
	fmt.Println("text:", text)

	var byteText = []byte(text)
	var sha = sha1.New()
	sha.Write(byteText)

	var encryptedText = sha.Sum(nil)
	var stringEncText = string(encryptedText)
	fmt.Println("normal encrypted:", stringEncText)

	var saltEncrypted, salt = hash_salting(text)
	fmt.Println("salting encrypted:", saltEncrypted, "salt:", salt)
}

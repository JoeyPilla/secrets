package main

import (
	"fmt"

	"./encrypt"
)

func main() {
	key, value := "fdsafsdafsa", "hello world"
	encryptedValue, _ := encrypt.Encrypt(key, value)
	decryptedValue, _ := encrypt.Decrypt(key, encryptedValue)
	fmt.Println(encryptedValue)
	fmt.Println(decryptedValue)
}

package main

import (
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter11/encryptString/utils"
	"log"
)

func main() {
	key := "111023043350789514532147"
	message := "I am a mesage"

	log.Println("Original mesage: ", message)
	encryptedString := utils.EncryptString(key, message)
	log.Println("Encrypted message: ", encryptedString)
	decryptedString := utils.DecryptString(key, encryptedString)
	log.Println("Decrypted message: ", decryptedString)
}
package main

import (
	"fmt"

	"github.com/anushasgorawar/cryptit/decrypt"
	"github.com/anushasgorawar/cryptit/encrypt"
)

func main() {
	input := "HelloThere"
	encrypted := encrypt.Encrypt(input)
	fmt.Println("Encrypted value of ", input, " is ", encrypted)
	decrypted := decrypt.Decrypt(encrypted)
	fmt.Println("Decrypted value of ", encrypted, " is ", decrypted)
}

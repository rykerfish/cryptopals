package aes

import (
	"crypto/aes"
	"log"
)

// EcbDecrypt ...
func EcbDecrypt(cipher []byte, key []byte) []byte {

	var c, err = aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	plaintext := make([]byte, 0, len(cipher))
	c.Decrypt(plaintext, cipher)

	return plaintext

}

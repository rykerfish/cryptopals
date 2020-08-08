package aes

import (
	"bytes"
	"crypto/aes"
	"log"
)

// EcbDecrypt decrypts a cipher using ECB mode of AES.
// NOTE: This function does not currently account for padding.
func EcbDecrypt(cipher []byte, key []byte) []byte {

	var aesBlock, err = aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	plaintext := make([]byte, 0, len(cipher))

	for i := 0; i < len(cipher); i += aesBlock.BlockSize() {
		plaintextBlock := make([]byte, aesBlock.BlockSize())
		aesBlock.Decrypt(plaintextBlock, cipher[i:i+aesBlock.BlockSize()])
		plaintext = append(plaintext, plaintextBlock...)

	}

	return plaintext

}

// DetectEcb ...
func DetectEcb(cipher []byte) bool {

	blockLen := 16

	for i := 0; i < len(cipher); i += blockLen {
		startBlock := cipher[i : i+blockLen]
		for j := 0; j < len(cipher); j += blockLen {
			// skips over startBlock when comparing
			if i == j {
				continue
			}
			compareBlock := cipher[j : j+blockLen]
			if bytes.Equal(startBlock, compareBlock) {
				return true
			}
		}
	}
	return false
}

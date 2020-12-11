package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
	"padding"
)

// EcbEncryptBlock will encrypt one block of text using AES ECB mode.
func EcbEncryptBlock(plaintextBlock []byte, aesBlock cipher.Block) []byte {

	cipherBlock := make([]byte, aesBlock.BlockSize())
	aesBlock.Encrypt(cipherBlock, plaintextBlock)

	return cipherBlock

}

// EcbEncrypt will encrypt text using the ECB mode of AES.
func EcbEncrypt(text []byte, key []byte) []byte {

	var aesBlock, err = aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// pads the message to be a multiple of the block size
	text = padding.PadMsg(text, aesBlock.BlockSize())

	cipher := make([]byte, 0, len(text))

	// encryption loop for each block
	for i := 0; i < len(text); i += aesBlock.BlockSize() {
		cipherBlock := EcbEncryptBlock(text[i:i+aesBlock.BlockSize()], aesBlock)
		cipher = append(cipher, cipherBlock...)
	}

	return cipher

}

// EcbDecryptBlock will decrypt one block of text using AES ECB mode.
func EcbDecryptBlock(cipherBlock []byte, aesBlock cipher.Block) []byte {

	textBlock := make([]byte, aesBlock.BlockSize())
	aesBlock.Decrypt(textBlock, cipherBlock)

	return textBlock

}

// EcbDecrypt decrypts a cipher using ECB mode of AES.
func EcbDecrypt(cipher []byte, key []byte) []byte {

	var aesBlock, err = aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	plaintext := make([]byte, 0, len(cipher))

	for i := 0; i < len(cipher); i += aesBlock.BlockSize() {
		plaintextBlock := EcbDecryptBlock(cipher[i:i+aesBlock.BlockSize()], aesBlock)
		plaintext = append(plaintext, plaintextBlock...)
	}

	// strips padding
	plaintext = padding.Strip(plaintext, aesBlock.BlockSize())

	return plaintext

}

// EcbDetect looks for repeating bytes of length blockLen to see if ECB mode of
// AES was used to encrypt the cipher.
func EcbDetect(cipher []byte) bool {

	blockLen := 16

	for i := 0; i < len(cipher); i += blockLen {
		if i+blockLen >= len(cipher) {
			break
		}
		startBlock := cipher[i : i+blockLen]
		for j := 0; j < len(cipher); j += blockLen {
			// skips over startBlock when checking for equality
			if i == j {
				continue
			}
			if j+blockLen >= len(cipher) {
				break
			}
			compareBlock := cipher[j : j+blockLen]
			if bytes.Equal(startBlock, compareBlock) {
				return true
			}
		}
	}
	return false
}

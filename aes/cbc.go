package aes

import (
	"crypto/aes"
	"log"
	"xor"
)

// CbcEncrypt implements AES CBC mode encryption.
// NOTE: Does not account for padding so text length must be a multiple
// of the key length.
func CbcEncrypt(text []byte, key []byte, iv []byte) []byte {

	cipher := make([]byte, 0, len(text))
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// initialize cbcBlock by xoring the first block of text with the IV
	cbcBlock := xor.Bytes(text[0:aesBlock.BlockSize()], iv)

	// encrypt each block and prep the next block with cipher block chaining by
	// xoring the following block with the newly generated cipher text
	for block := aesBlock.BlockSize(); block <= len(text); block += aesBlock.BlockSize() {
		cipherBlock := EcbEncryptBlock(cbcBlock, aesBlock)
		cipher = append(cipher, cipherBlock...)

		if block != len(text) {
			cbcBlock = xor.Bytes(cipherBlock, text[block:block+aesBlock.BlockSize()])
		}
	}

	return cipher

}

// CbcDecrypt implements AES CBC mode decryption. len(key) must equal len(iv)
// and len(cipher) must be a multiple of len(key)
func CbcDecrypt(cipher []byte, key []byte, iv []byte) []byte {

	// cipher.Block initialization
	plaintext := make([]byte, 0, len(cipher))
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// runs the first round of decryption with the IV
	prevCipherBlock := cipher[0:aesBlock.BlockSize()]
	textBlock := xor.Bytes(EcbDecryptBlock(prevCipherBlock, aesBlock), iv)
	plaintext = append(plaintext, textBlock...)

	// every subsequent run decrypts the next block then xors it with the previous
	// block of ciphertext as part of the CBC algorithm.
	for block := aesBlock.BlockSize(); block < len(cipher); block += aesBlock.BlockSize() {
		decryptBlock := EcbDecryptBlock(cipher[block:block+aesBlock.BlockSize()], aesBlock)
		textBlock := xor.Bytes(decryptBlock, prevCipherBlock)
		plaintext = append(plaintext, textBlock...)

		prevCipherBlock = cipher[block : block+aesBlock.BlockSize()]
	}

	return plaintext

}

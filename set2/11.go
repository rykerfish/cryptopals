package setTwo

import (
	"aes"
	"fmt"
)

// ProblemEleven from Cryptopals Set 2
func ProblemEleven() {
	text := []byte("this is not an even sixteen characters")
	key := []byte("YELLOW SUBMARINE")

	cipher := aes.EcbEncrypt(text, key)

	fmt.Println(string(cipher))
	fmt.Println(len(cipher))

	plaintext := aes.EcbDecrypt(cipher, key)
	fmt.Println(string(plaintext))
}

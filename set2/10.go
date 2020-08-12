package setTwo

import (
	"aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ProblemTen from Set 2 of Cryptopals
func ProblemTen() {

	// FileIO stuff
	file, err := os.Open("./files/10.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Cipher specs
	cipher, _ := base64.StdEncoding.DecodeString(string(data))
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)

	// this string was already padded when given to us to decrypt,
	// hence the 4 bytes of 0x04 at the end of the string.
	plaintext := aes.CbcDecrypt(cipher, key, iv)
	fmt.Println(string(plaintext))
}

package setOne

import (
	"aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ProblemSeven from Cryptopals Set 1
func ProblemSeven() {

	key := "YELLOW SUBMARINE"

	file, err := os.Open("./files/set1/7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	cipher, _ := base64.StdEncoding.DecodeString(string(text))

	plaintext := aes.EcbDecrypt(cipher, []byte(key))

	fmt.Println(string(plaintext))

}

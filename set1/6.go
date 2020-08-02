package setOne

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"xor"
)

// ProblemSix from Cryptopals Set 1
func ProblemSix() {

	// file stuff
	file, err := os.Open("./files/6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	cipher, _ := base64.StdEncoding.DecodeString(string(text))

	keyLength := xor.GuessRepeatingKeyLength(cipher)
	key := xor.BreakRepeatingKey(cipher, keyLength)
	plaintext := xor.RepeatingKeyXor(cipher, key)

	fmt.Printf("Decrypt: %s\nKey: %s\n", plaintext, key)

}

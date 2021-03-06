package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/rykerfish/cryptopals/xor"
)

// ProblemSix from Cryptopals Set 1
func main() {

	// file stuff
	file, err := os.Open("../files/set1/6.txt")
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

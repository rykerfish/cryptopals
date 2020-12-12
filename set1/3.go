package main

import (
	hex "encoding/hex"
	"fmt"
	"github.com/rykerfish/cryptopals/xor"
)

// ProblemThree from Cryptopals Set 1
func main() {

	str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, _ := hex.DecodeString(str)

	plaintext, key := xor.SolveSingleCharXor(bytes)

	fmt.Println("Decrypt:", plaintext)
	fmt.Println("Key:", string(key))

}

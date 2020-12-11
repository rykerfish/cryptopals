package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/rykerfish/cryptopals/aes"
)

// ProblemEleven from Cryptopals Set 2.
func main() {

	inReader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	text, _ := inReader.ReadString('\n')

	cipher, mode := aes.EncryptionOracle([]byte(text))

	if aes.EcbDetect(cipher) {
		fmt.Println("ECB mode detected.")
	} else {
		fmt.Println("CBC mode (probably)")
	}

	if mode == 1 {
		fmt.Println("Real mode: ECB")
	} else {
		fmt.Println("Real mode: CBC")
	}

}

package setTwo

import (
	"aes"
	"bufio"
	"fmt"
	"os"
)

// ProblemEleven from Cryptopals Set 2.
func ProblemEleven() {

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

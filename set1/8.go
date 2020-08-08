package setOne

import (
	"aes"
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

// ProblemEight from Cryptopals Set 1
func ProblemEight() {

	// File IO stuff
	file, err := os.Open("./files/8.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var ecbBytes []byte
	var lineNum = 1
	for scanner.Scan() {
		hexBytes, _ := hex.DecodeString(scanner.Text())
		var same bool = aes.DetectEcb(hexBytes)
		if same {
			ecbBytes = hexBytes
			break
		}
		lineNum++
	}

	fmt.Printf("ECB was used to encrypt:\n%s\n", ecbBytes)
	fmt.Printf("Line number: %d\n", lineNum)

}

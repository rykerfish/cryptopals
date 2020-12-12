package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"github.com/rykerfish/cryptopals/aes"
)

// ProblemEight from Cryptopals Set 1
func main() {

	// File IO stuff
	file, err := os.Open("../files/set1/8.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Looks through every line in the file and uses aes.DetectEcb to find
	// the line with repeating 16 byte blocks
	var ecbBytes []byte
	var lineNum = 1
	for scanner.Scan() {
		hexBytes, _ := hex.DecodeString(scanner.Text())
		var same bool = aes.EcbDetect(hexBytes)
		if same {
			ecbBytes = hexBytes
			break
		}
		lineNum++
	}

	fmt.Printf("ECB was used to encrypt:\n%s\n", ecbBytes)
	fmt.Printf("Line number: %d\n", lineNum)

}

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"github.com/rykerfish/cryptopals/freq"
	"github.com/rykerfish/cryptopals/xor"
)

// ProblemFour from Cryptopals Set 1
func main() {

	file, err := os.Open("../files/set1/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var decrypt, key string
	var maxScore float64

	// searches for the highest scoring line in the file by brute
	// forcing each line and comparing it to the current hig score
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexBytes, _ := hex.DecodeString(scanner.Text())
		plaintext, lineKey := xor.SolveSingleCharXor(hexBytes)
		lineScore := freq.ScoreEngText(plaintext)

		if lineScore > maxScore {
			decrypt = plaintext
			key = lineKey
			maxScore = lineScore
		}
	}

	fmt.Printf("Decrypt: %sKey: %s \n", decrypt, key)

}

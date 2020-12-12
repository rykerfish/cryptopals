package main

import (
	"fmt"
	"github.com/rykerfish/cryptopals/padding"
)

// ProblemNine from Set 2 of Cryptopals
func main() {

	message := "YELLOW SUBMARINE"
	blockLen := 20

	padMsg := padding.PadMsg([]byte(message), blockLen)

	fmt.Println(string(padMsg))

}

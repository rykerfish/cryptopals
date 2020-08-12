package setTwo

import (
	"fmt"
	"padding"
)

// ProblemNine from Set 2 of Cryptopals
func ProblemNine() {

	message := "YELLOW SUBMARINE"
	blockLen := 20

	padMsg := padding.PadMsg([]byte(message), blockLen)

	fmt.Println(string(padMsg))

}

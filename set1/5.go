package main

import (
	"encoding/hex"
	"fmt"
	"github.com/rykerfish/cryptopals/xor"
)

// ProblemFive from Cryptopals Set 1
func main() {

	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"

	encrypt := hex.EncodeToString(xor.RepeatingKeyXor([]byte(text), []byte(key)))

	fmt.Println(encrypt)

	ans := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if encrypt == ans {
		fmt.Println("Strings match!")
	}

}

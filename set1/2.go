package main

import (
	hex "encoding/hex"
	"fmt"
	"github.com/rykerfish/cryptopals/xor"
)

// ProblemTwo from Cryptopals Set 1
func main() {

	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"

	bytes1, _ := hex.DecodeString(s1)
	bytes2, _ := hex.DecodeString(s2)

	xor := hex.EncodeToString(xor.Bytes(bytes1, bytes2))
	fmt.Println(xor)

	ans := "746865206b696420646f6e277420706c6179"
	if xor == ans {
		fmt.Println("Strings match.")
	}

}

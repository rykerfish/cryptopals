package setOne

import (
	"encoding/base64"
	hex "encoding/hex"
	"fmt"
)

// ProblemOne from Cryptopals Set 1
func ProblemOne() {

	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	ans := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	bytes, _ := hex.DecodeString(hexStr)

	encoded := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(encoded)

	if encoded == ans {
		fmt.Println("Strings match.")
	}

}

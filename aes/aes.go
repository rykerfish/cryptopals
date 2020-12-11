package aes

import (
	"math/rand"
	"time"
)

// GenerateKey returns a key of random bytes of length size
func GenerateKey(size int) []byte {
	rand.Seed(time.Now().UnixNano())
	key := make([]byte, size)

	// uses a not cryptographically secure PRNG
	for i := 0; i < size; i++ {
		key[i] = byte(rand.Intn(256))
	}

	return key
}

// EncryptionOracle will randomly decide between encrypting with ECB or CBC mode
// and will add 5-10 random bytes to the front and back of the cipher.
// Returns the cipher as well as int = 1 if ECB and int = 2 if CBC
func EncryptionOracle(text []byte) ([]byte, int) {

	rand.Seed(time.Now().UnixNano())

	key := GenerateKey(16)

	// +20 since we are going to add 5-10 random bytes to the front and back
	cipher := make([]byte, 0, len(text)+20)
	var mode int

	// adds 5-10 random bytes to the front of cipher
	cipher = append(cipher, GenerateKey(5+rand.Intn(5))...)

	// 50/50 chance between using ECB or CBC mode of AES
	if rand.Intn(2) == 0 {
		cipher = append(cipher, EcbEncrypt(text, key)...)
		mode = 1
	} else {
		// CBC mode uses a randomly generated IV
		iv := GenerateKey(16)
		cipher = append(cipher, CbcEncrypt(text, key, iv)...)
		mode = 2
	}

	// adds 5-10 random bytes to the back of the cipher
	cipher = append(cipher, GenerateKey(5+rand.Intn(5))...)

	return cipher, mode
}

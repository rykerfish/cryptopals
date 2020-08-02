package xor

import (
	"distance"
	"fmt"
	"freq"
	"math"
)

// Bytes xors two equal length byte slices together. If the two slices aren't
// equal lengths, returns an empty byte slice.
func Bytes(bytes1 []byte, bytes2 []byte) []byte {

	if len(bytes1) != len(bytes2) {
		fmt.Println("Byte arrays of non-equal length used in xor.Bytes")
		return make([]byte, 0)
	}

	xor := make([]byte, len(bytes1))

	for i := range bytes1 {
		xor[i] = bytes1[i] ^ bytes2[i]
	}

	return xor
}

// SingleCharXor encrypts a byte slice with a single character. If the key is
// more than a single character, returns an empty byte slice.
func SingleCharXor(bytes []byte, key string) []byte {

	keyByte := []byte(key)

	encrypt := make([]byte, len(bytes))

	for i, byte := range bytes {
		encrypt[i] = byte ^ keyByte[0]
	}
	return encrypt
}

// SolveSingleCharXor will solve a ciphertext if it has been encrypted with a single character xor.
// Returns the plaintext and the encryption key
func SolveSingleCharXor(ciphertext []byte) (plaintext string, key string) {

	var maxScore float64

	for i := 0; i < 256; i++ {
		decrypt := SingleCharXor(ciphertext, string(i))
		score := freq.ScoreEngText(string(decrypt))

		if score > maxScore {
			plaintext = string(decrypt)
			maxScore = score
			key = string(i)
		}
	}

	return plaintext, key

}

// RepeatingKeyXor takes in a text and a key and encrypts the text using
// a repeating key encryption.
func RepeatingKeyXor(plaintext []byte, key []byte) (ciphertext []byte) {

	cipher := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		cipher[i] = plaintext[i] ^ key[i%len(key)]
	}

	return cipher

}

// GuessRepeatingKeyLength finds the most likely key size for a repeating key
// xor by finding the key length with the smallest normalized edit (hamming) distance.
func GuessRepeatingKeyLength(cipher []byte) int {
	// Why this works: https://crypto.stackexchange.com/questions/8115/repeating-key-xor-and-hamming-distance/8118#8118
	var minDist float64 = math.MaxInt32
	minKey := -1

	// Upper and lower limits for the guess on the key size.
	const minGuess, maxGuess = 2, 40
	for size := minGuess; size <= maxGuess; size++ {

		var hamDist float64
		blockCount := len(cipher) / size

		// Sums the normalized hamming distance between every adjacent block of
		for i := 0; i < blockCount-1; i++ {
			blockOne := cipher[i*size : i*size+size]
			blockTwo := cipher[(i+1)*size : (i+2)*size]
			hamDist += float64(distance.Hamming(blockOne, blockTwo)) / float64(size)
		}
		// Averages hamDist with the number of blocks used to calculate  hamDist
		hamDist = hamDist / float64(blockCount)

		// Always keep the key size with the smallest aggregate edit distance
		if hamDist < minDist {
			minKey = size
			minDist = hamDist
		}
	}

	return minKey
}

// BreakRepeatingKey returns a key to a repeating xor cipher, provided the
// key's length.
func BreakRepeatingKey(cipher []byte, keyLen int) []byte {

	splitCipher := make([][]byte, keyLen)
	for slice := range splitCipher {
		splitCipher[slice] = make([]byte, 0, len(cipher)/keyLen+1)
	}

	// stores every nth byte in a different slice
	for index, byte := range cipher {
		splitCipher[index%keyLen] = append(splitCipher[index%keyLen], byte)
	}

	// solves each slice as a single character xor since every nth character
	// from the cipher was encrypted with the same character as a key
	var key string
	for _, slice := range splitCipher {
		_, keyChar := SolveSingleCharXor(slice)
		key += keyChar
	}

	return []byte(key)

}

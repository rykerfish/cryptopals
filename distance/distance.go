package distance

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Hamming finds the hamming distance, or the number of differeing bits, between
// two slices.
func Hamming(a []byte, b []byte) int {

	if len(a) != len(b) {
		fmt.Println("Error: Can't compute hamming distance between objects of different lengths")
		os.Exit(1)
	}

	var dist int

	for i := range a {
		binaryA := strconv.FormatInt(int64(a[i]), 2)
		binaryB := strconv.FormatInt(int64(b[i]), 2)

		// zero extends binA and binB so that they are each 8 characters long to prevent
		// out of range errors.
		if len(binaryA) < 8 || len(binaryB) < 8 {
			binaryA = strings.Repeat("0", 8-len(binaryA)) + binaryA
			binaryB = strings.Repeat("0", 8-len(binaryB)) + binaryB
		}

		for i := range binaryA {
			if binaryA[i] != binaryB[i] {
				dist++
			}
		}

	}
	return dist
}

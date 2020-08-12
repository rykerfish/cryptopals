package padding

// PadMsg pads the length of a message message to be an even multiple of blockLen
// using PKCS#7 padding.
func PadMsg(message []byte, blockLen int) []byte {

	padLen := (blockLen - (len(message) % blockLen))

	if padLen == 0 {
		return message
	}

	for i := 0; i < padLen; i++ {
		message = append(message, byte(padLen))
	}

	return message

}

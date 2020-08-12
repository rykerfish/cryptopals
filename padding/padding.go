package padding

// PadMsg pads the length of a message message to be an even multiple of blockLen
// using PKCS#7 padding.
func PadMsg(message []byte, blockLen int) []byte {

	padLen := (blockLen - (len(message) % blockLen))

	if padLen == 0 {
		return message
	}

	padding := make([]byte, padLen)
	for i := 0; i < padLen; i++ {
		padding[i] = byte(padLen)
	}

	message = append(message, padding...)

	return message

}

// Strip removes PKCS#7 padding from a message.
func Strip(message []byte, blockLen int) []byte {

	b := message[len(message)-1]
	if int(b) > blockLen {
		return message
	}

	for i := len(message) - int(b); i < len(message); i++ {
		if message[i] != b {
			return message
		}
	}

	return message[:len(message)-int(b)]

}

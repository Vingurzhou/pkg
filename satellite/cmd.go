package satellite

import (
	"encoding/hex"
	"strings"
)

// ScrambleHexString scrambles a hex string and returns the scrambled result as a hex string
func ScrambleHexString(hexInput string) (string, error) {
	decodedBytes, err := hex.DecodeString(strings.ReplaceAll(hexInput, " ", ""))
	if err != nil {
		return "", err
	}
	scrambledBytes := ScrambleBytes(decodedBytes, byte(0xff))
	scrambledHex := hex.EncodeToString(scrambledBytes)
	return scrambledHex, nil
}

// generateNextLFSRByte generates the next byte from a Linear Feedback Shift Register
func generateNextLFSRByte(registerState *byte) byte {
	var outputByte byte = 0
	for i := 0; i < 8; i++ {
		leastSignificantBit := *registerState & 1
		outputByte = (outputByte << 1) | leastSignificantBit

		var feedbackMask byte = 0x5F
		var feedbackBit byte = 0
		tempState := *registerState & feedbackMask
		for tempState > 0 {
			feedbackBit ^= (tempState & 1)
			tempState >>= 1
		}
		*registerState = (feedbackBit << 7) | (*registerState >> 1)
	}
	return outputByte
}

// ScrambleBytes applies LFSR-based scrambling to the input bytes
func ScrambleBytes(inputBytes []byte, initialSeed byte) []byte {
	registerState := initialSeed
	scrambledBytes := make([]byte, len(inputBytes))
	for i, inputByte := range inputBytes {
		pseudoRandomByte := generateNextLFSRByte(&registerState)
		scrambledBytes[i] = inputByte ^ pseudoRandomByte
	}
	return scrambledBytes
}

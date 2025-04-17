package crypto

func EncryptSymmetrically(input string, key rune) string {
	var output []rune
	for _, c := range input {
		// 自动转二进制进行异或，相同为0，不同为1，最后转10进制
		output = append(output, c^key)
	}
	return string(output)
}

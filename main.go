package main

import (
	"fmt"
)

// lfsrNextByte 根据当前 LFSR 状态生成一个伪随机字节，并更新状态。
// 参数 state 是指向 LFSR 状态的指针，初始种子为 0xff。
func lfsrNextByte(state *byte) byte {
	var out byte = 0
	// 每次生成一个字节，需要 8 次迭代，每次产生一个 bit，第一产生的 bit 为最高位
	for i := 0; i < 8; i++ {
		// 取当前状态的最低位作为输出 bit
		bit := *state & 1
		out = (out << 1) | bit

		// 根据多项式 H(x)=x^8+x^6+x^4+x^3+x^2+x+1，tap 位为 {6,4,3,2,1,0}，对应掩码 0x5F（二进制 0101 1111）
		var mask byte = 0x5F
		var feedback byte = 0
		temp := *state & mask
		// 计算选定位的奇偶性（奇数则反馈为 1，偶数为 0）
		for temp > 0 {
			feedback ^= (temp & 1)
			temp >>= 1
		}
		// 更新状态：右移 1 位，并将反馈位移入最高位
		*state = (feedback << 7) | (*state >> 1)
	}
	return out
}

// scramble 对输入字节数组进行加扰，seed 为 LFSR 初始状态。
func scramble(input []byte, seed byte) []byte {
	state := seed
	output := make([]byte, len(input))
	for i, b := range input {
		fmt.Println(b)
		prByte := lfsrNextByte(&state)
		output[i] = b ^ prByte
	}
	return output
}

func main() {
	// 输入字节：3d 0d ca
	// input := []byte{0x3d, 0x0d, 0xca, 0x4a, 0x68, 0xf2}
	input := []byte{0xc2, 0x34, 0x54, 0x10, 0x00, 0x1b}
	// 期望加扰后的输出：c2 34 54
	// expected := []byte{0xc2, 0x34, 0x54}
	seed := byte(0xff)
	output := scramble(input, seed)

	fmt.Printf("输入:    % x\n", input)
	// fmt.Printf("期望输出: % x\n", expected)
	fmt.Printf("实际输出: % x\n", output)

	// // 验证加扰结果是否符合预期
	// if len(output) == len(expected) {
	// 	valid := true
	// 	for i := 0; i < len(output); i++ {
	// 		if output[i] != expected[i] {
	// 			valid = false
	// 			break
	// 		}
	// 	}
	// 	if valid {
	// 		fmt.Println("加扰验证通过!")
	// 	} else {
	// 		fmt.Println("加扰验证失败!")
	// 	}
	// } else {
	// 	fmt.Println("输出长度不匹配!")
	// }
}

package util

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	// 测试字符串
	strs := []string{"go", "java", "python"}
	fmt.Println(Contains(strs, "go"))     // true
	fmt.Println(Contains(strs, "golang")) // false

	// 测试整数
	ints := []int{1, 2, 3, 4}
	fmt.Println(Contains(ints, 3)) // true
	fmt.Println(Contains(ints, 5)) // false
}

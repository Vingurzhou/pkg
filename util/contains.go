package util

// 泛型 contains 函数，适用于任何可比较类型
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

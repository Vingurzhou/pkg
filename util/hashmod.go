package util

import (
	"hash/fnv"
)

// 表数量 / 库数量（一般 2、4、8、16、32、64…）
func HashMod(key string, mod uint32) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32() % mod
}

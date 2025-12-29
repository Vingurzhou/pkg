package cache

import (
	"context"
	"fmt"
	"hash/fnv"
	"math"

	"github.com/redis/go-redis/v9"
)

type BloomFilter struct {
	client *redis.Client // Redis 客户端
	key    string        // 集合
	size   int64         // 最大数量
	hashes int           // 哈希函数数量，多哈希函数降低误判率
}

func NewBloomFilter(client *redis.Client, key string, errorRate float64, capacity int) *BloomFilter {
	m := -float64(capacity) * math.Log(errorRate) / (math.Ln2 * math.Ln2) //m = -(n * ln(p)) / (ln2)^2
	k := (m / float64(capacity)) * math.Ln2                               //k = (m/n) * ln2
	hashes := int(math.Ceil(k))
	return &BloomFilter{
		client: client,
		key:    key,
		size:   int64(math.Ceil(m)),
		hashes: hashes,
	}
}

func (bf *BloomFilter) Add(ctx context.Context, value string) error {
	// TODO 使用 Redis Lua 脚本减少网络开销
	for i := 0; i < bf.hashes; i++ {
		idx := bf.hashIndex(value, i)
		if err := bf.client.SetBit(ctx, bf.key, idx, 1).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (bf *BloomFilter) Exists(ctx context.Context, value string) error {
	for i := 0; i < bf.hashes; i++ {
		idx := bf.hashIndex(value, i)
		//领取->领取
		//未领取->未领取/领取
		bit, err := bf.client.GetBit(ctx, bf.key, idx).Result()
		if err != nil {
			return err
		}
		if bit == 0 {
			err := fmt.Errorf("元素不存在")
			return err
		}
	}
	return nil
}
func (bf *BloomFilter) hashIndex(value string, i int) int64 {
	h := fnv.New64a()
	h.Write([]byte(value))
	h.Write([]byte{byte(i)})
	return int64(h.Sum64() % uint64(bf.size))
}

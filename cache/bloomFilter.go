package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// BloomFilter 结构体
type BloomFilter struct {
	client    *redis.Client
	key       string
	size      int64
	errorRate float64
}

// NewBloomFilter 创建新的布隆过滤器
func NewBloomFilter(client *redis.Client, key string, size int64, errorRate float64) *BloomFilter {
	return &BloomFilter{
		client:    client,
		key:       key,
		size:      size,
		errorRate: errorRate,
	}
}

// Add 添加元素到布隆过滤器
func (bf *BloomFilter) Add(ctx context.Context, value string) error {
	// BF.ADD 命令添加单个元素
	cmd := bf.client.BFAdd(ctx, bf.key, value)
	return cmd.Err()
}

// Exists 检查元素是否存在
func (bf *BloomFilter) Exists(ctx context.Context, value string) (bool, error) {
	// BF.EXISTS 命令检查元素是否存在
	cmd := bf.client.BFExists(ctx, bf.key, value)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val(), nil
}

// Init 初始化布隆过滤器（可选）
func (bf *BloomFilter) Init(ctx context.Context) error {
	// BF.RESERVE 创建具有指定参数的布隆过滤器
	cmd := bf.client.BFReserve(ctx, bf.key, bf.errorRate, bf.size)
	return cmd.Err()
}

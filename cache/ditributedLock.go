package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type DistributedLock struct {
	client     *redis.Client
	lockKey    string
	lockValue  string // 唯一标识，通常使用随机值
	expiration time.Duration
}

// NewDistributedLock 创建新的分布式锁实例
func NewDistributedLock(client *redis.Client, key string, expiration time.Duration) *DistributedLock {
	return &DistributedLock{
		client:     client,
		lockKey:    key,
		lockValue:  fmt.Sprintf("%d", time.Now().UnixNano()), // 生成唯一值
		expiration: expiration,
	}
}

// TryLock 尝试获取锁
func (dl *DistributedLock) TryLock(ctx context.Context) (bool, error) {
	// 使用 SETNX 命令，只有当 key 不存在时才设置成功
	success, err := dl.client.SetNX(ctx, dl.lockKey, dl.lockValue, dl.expiration).Result()
	if err != nil {
		return false, err
	}
	return success, nil
}

// Unlock 释放锁
func (dl *DistributedLock) Unlock(ctx context.Context) (bool, error) {
	// 使用 Lua 脚本确保只有持有锁的客户端才能释放锁
	script := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `

	result, err := dl.client.Eval(ctx, script, []string{dl.lockKey}, dl.lockValue).Int64()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

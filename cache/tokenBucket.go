package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// TokenBucket 令牌桶结构
type TokenBucket struct {
	redisClient *redis.Client
	key         string    // redis中的key
	capacity    int64     // 桶容量
	rate        float64   // 令牌生成速率(每秒)
	lastRefill  time.Time // 上次填充时间
}

// NewTokenBucket 创建新的令牌桶
func NewTokenBucket(client *redis.Client, key string, capacity int64, rate float64) *TokenBucket {
	return &TokenBucket{
		redisClient: client,
		key:         key,
		capacity:    capacity,
		rate:        rate,
	}
}

// Allow 检查是否允许通过
func (tb *TokenBucket) Allow(ctx context.Context, tokens int64) (bool, error) {
	// 使用Lua脚本保证原子性
	luaScript := `
        local key = KEYS[1]
        local now = tonumber(ARGV[1])
        local capacity = tonumber(ARGV[2])
        local rate = tonumber(ARGV[3])
        local tokens = tonumber(ARGV[4])
        
        -- 获取当前令牌数和上次更新时间
        local current = redis.call('HGET', key, 'tokens')
        local last_refill = redis.call('HGET', key, 'last_refill')
        
        if current == false then
            current = capacity
        else
            current = tonumber(current)
        end
        
        if last_refill == false then
            last_refill = now
        else
            last_refill = tonumber(last_refill)
        end
        
        -- 计算新生成的令牌
        local elapsed = now - last_refill
        local new_tokens = math.floor(elapsed * rate / 1000000000)
        current = math.min(capacity, current + new_tokens)
        
        -- 检查是否有足够的令牌
        if current >= tokens then
            current = current - tokens
            redis.call('HSET', key, 'tokens', current)
            redis.call('HSET', key, 'last_refill', now)
            return 1
        end
        return 0
    `

	now := time.Now().UnixNano()
	result, err := tb.redisClient.Eval(
		ctx,
		luaScript,
		[]string{tb.key},
		now,
		tb.capacity,
		tb.rate,
		tokens,
	).Int()

	if err != nil {
		return false, err
	}

	return result == 1, nil
}

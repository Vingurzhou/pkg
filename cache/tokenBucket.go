package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenBucket struct {
	redisClient *redis.Client
	key         string  // redis中的key
	capacity    int64   // 桶容量
	rate        float64 // 令牌生成速率
}

func NewTokenBucket(client *redis.Client, key string, capacity int64, rate float64) *TokenBucket {
	return &TokenBucket{
		redisClient: client,
		key:         key,
		capacity:    capacity,
		rate:        rate,
	}
}

func (tb *TokenBucket) Decrease(ctx context.Context, tokens int64) error {
	script := `
local key = KEYS[1]
local now = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local rate = tonumber(ARGV[3])
local tokens = tonumber(ARGV[4])

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

local elapsed = now - last_refill
local new_tokens = math.floor(elapsed * rate / 1000000000)
current = math.min(capacity, current + new_tokens)

if current >= tokens then
    current = current - tokens
    redis.call('HSET', key, 'tokens', current)
    redis.call('HSET', key, 'last_refill', now)
    return 1
end
error(string.format("%s 不够,想扣减%d,但仅剩 %d", KEYS[1],tokens,current))
    `

	return tb.redisClient.Eval(
		ctx,
		script,
		[]string{tb.key},
		time.Now().UnixNano(),
		tb.capacity,
		tb.rate,
		tokens,
	).Err()

}

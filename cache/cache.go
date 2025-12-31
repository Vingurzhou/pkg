package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisOptions struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisCli(r RedisOptions) *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	return redisCli
}

type RedisClusterOptions struct {
	Addrs []string
}

func NewRedisClusterCli(r RedisClusterOptions) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    r.Addrs,
		Password: "", // 如果 Redis 有密码，设置这里
		// 其他可选配置
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})
}

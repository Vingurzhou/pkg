package cache

import "github.com/redis/go-redis/v9"

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

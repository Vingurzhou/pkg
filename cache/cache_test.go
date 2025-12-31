package cache

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx             context.Context
	redisCli        *redis.Client
	redisClusterCli *redis.ClusterClient
)

func TestMain(m *testing.M) {
	ctx = context.Background()

	redisCli = NewRedisCli(RedisOptions{
		Addr: "127.0.0.1:7001",
	})
	redisClusterCli = NewRedisClusterCli(RedisClusterOptions{
		Addrs: []string{
			"127.0.0.1:7001",
			"127.0.0.1:7002",
			"127.0.0.1:7003",
		},
	})
	code := m.Run()

	os.Exit(code)
}
func TestNewRedisCli(t *testing.T) {
	t.Log(redisCli.Ping(ctx))
}

func TestDistributedLock(t *testing.T) {
	lock := NewDistributedLock(redisCli, "lock", "user1", 3*time.Second)
	t.Log(lock.TryLock(ctx))
	t.Log(lock.TryLock(ctx))

	time.Sleep(4 * time.Second)

	t.Log(lock.Unlock(ctx))
	t.Log(lock.Unlock(ctx))
}
func TestBloomFilter(t *testing.T) {
	recordList := NewBloomFilter(redisCli, "record", 0.99, 1000)
	err := recordList.Add(ctx, "user1")
	if err != nil {
		t.Fatal(err)
	}
	err = recordList.Add(ctx, "user2")
	if err != nil {
		t.Fatal(err)
	}
	err = recordList.Add(ctx, "user3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(recordList.Exists(ctx, "user1"))
	t.Log(recordList.Exists(ctx, "user7"))
	t.Log(recordList.Exists(ctx, "user8"))
	t.Log(recordList.Exists(ctx, "user9"))
}
func TestTokenBucket(t *testing.T) {
	stock := NewTokenBucket(redisCli, "stock", 6, 1)
	t.Log(stock.Decrease(ctx, 3))
	t.Log(stock.Decrease(ctx, 3))
	t.Log(stock.Decrease(ctx, 3))
	time.Sleep(time.Second)
	t.Log(stock.Decrease(ctx, 3))
	time.Sleep(time.Second)
	t.Log(stock.Decrease(ctx, 3))
	time.Sleep(time.Second)
	t.Log(stock.Decrease(ctx, 3))
}

func TestNewRedisClusterCli(t *testing.T) {
	slots, err := redisClusterCli.ClusterSlots(ctx).Result()
	if err != nil {
		t.Fatal(err)
	}
	for _, slot := range slots {
		t.Log(slot)
	}
	t.Log(redisClusterCli.Set(ctx, "1{packet}", "1", 3*time.Second))
	t.Log(redisClusterCli.Set(ctx, "2{packet}", "2", 3*time.Second))
	t.Log(redisClusterCli.Set(ctx, "3{packet2}", "3", 3*time.Second))
	t.Log(redisClusterCli.Get(ctx, "1{packet}"))
	t.Log(redisClusterCli.Get(ctx, "2{packet}"))
	t.Log(redisClusterCli.Get(ctx, "3{packet2}"))
	t.Log(redisClusterCli.ClusterKeySlot(ctx, "1{packet}"))
	t.Log(redisClusterCli.ClusterKeySlot(ctx, "2{packet}"))
	t.Log(redisClusterCli.ClusterKeySlot(ctx, "3{packet2}"))
}

package cache

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	ctx      context.Context
	redisCli *redis.Client
)

func TestMain(m *testing.M) {
	ctx = context.Background()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := strconv.ParseInt(os.Getenv("DB"), 10, 64)
	if err != nil {
		panic(err)
	}

	redisCli = NewRedisCli(RedisOptions{
		Addr:     os.Getenv("Addr"),
		Password: os.Getenv("Password"),
		DB:       int(db),
	})

	code := m.Run()

	os.Exit(code)
}
func TestNewRedisCli(t *testing.T) {
	for _, v := range redisCli.Keys(ctx, "*").Val() {
		t.Log(v)
	}

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

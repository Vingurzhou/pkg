package cache

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewRedisCli(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
	db, err := strconv.ParseInt((os.Getenv("DB")), 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	redisCli := NewRedisCli(RedisOptions{
		Addr:     os.Getenv("Addr"),
		Password: os.Getenv("Password"),
		DB:       int(db),
	})
	for _, v := range redisCli.Keys(ctx, "*").Val() {
		t.Log(v)
	}
}

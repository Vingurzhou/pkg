package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// TODO
type RedLock struct {
	clientList []*redis.Client //实例，非集群
}
type DistributedLock struct {
	client          *redis.Client //redis客户端
	lockKey         string        //锁
	lockValue       string        //钥匙持有者
	redisExpiration int           //自动解锁时间，防止解锁失败死锁,纳秒
	goExpiration    time.Duration //自动解锁时间，防止解锁失败死锁,毫秒
}

func NewDistributedLock(client *redis.Client, key string, value string, expiration time.Duration) *DistributedLock {
	return &DistributedLock{
		client:          client,
		lockKey:         key,
		lockValue:       value,
		goExpiration:    expiration,
		redisExpiration: int(expiration / time.Millisecond),
	}
}

func (dl *DistributedLock) TryLock(ctx context.Context) error {
	script := `
if redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2]) then
	return 1
else
	error(string.format("%s 已经上锁", KEYS[1]))
end
    `

	err := dl.client.Eval(ctx, script, []string{dl.lockKey}, dl.lockValue, dl.redisExpiration).Err()
	if err != nil {
		return err
	}
	go dl.watchDog(ctx)
	return nil
}

// 看门狗续锁，防止自动解锁<手动解锁时间
func (dl *DistributedLock) watchDog(ctx context.Context) {
	ticker := time.NewTicker(dl.goExpiration / 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			script := `
if redis.call('get', KEYS[1]) == ARGV[1] then
	redis.call("PEXPIRE", KEYS[1], ARGV[2])
	return 1
else
	error(string.format("%s 没上锁或 %s 非钥匙持有者", KEYS[1],ARGV[1]))
end
    			`
			err := dl.client.Eval(ctx, script, []string{dl.lockKey}, dl.lockValue, dl.redisExpiration).Err()
			if err != nil {
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

func (dl *DistributedLock) Unlock(ctx context.Context) error {
	script := `
if redis.call('get', KEYS[1]) == ARGV[1] then
    redis.call('del', KEYS[1])
    return 1
else
	error(string.format("%s 没上锁或 %s 非钥匙持有者", KEYS[1],ARGV[1]))
end
    `
	err := dl.client.Eval(ctx, script, []string{dl.lockKey}, dl.lockValue).Err()
	if err != nil {
		return err
	}
	return nil
}

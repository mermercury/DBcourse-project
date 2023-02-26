package redis

import (
	"Back_End/conf"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

const (
	FALSE string = "false"
	TRUE  string = "true"
)

var globalRedisPool *redis.Client

func init() {
	addrStr := fmt.Sprintf(
		"%s:%d",
		conf.GlobalConfig.RedisConf.Addr,
		conf.GlobalConfig.RedisConf.Port,
	)

	opts := redis.Options{
		Addr:     addrStr,
		Password: conf.GlobalConfig.RedisConf.Password,
		DB:       conf.GlobalConfig.RedisConf.DBIndex,
		
	}
	globalRedisPool = redis.NewClient(
		&opts,
	)
	//globalRedisPool.Conn().Auth(context.Background(),conf.GlobalConfig.RedisConf.Password)
	fmt.Println("redis password : ",conf.GlobalConfig.RedisConf.Password)
}

func Get(key string) (ans string, contain bool) {
	ctx := context.Background()
	val := globalRedisPool.Get(ctx, key)

	if val == nil || val.Err() != nil {
		ans = ""
		contain = false
		return
	}

	ans = val.String()
	contain = true
	return
}

func Set(key string, val string, expire time.Duration) (err error) {
	ctx := context.Background()
	state := globalRedisPool.Set(ctx, key, val, expire)
	if state.Err() != nil {
		err = state.Err()
		return
	}
	err = nil
	return
}

func SetHash(key string, kv map[string]string, expire time.Duration) (err error) {
	ctx := context.Background()
	var kvSlice []string
	for k, v := range kv {
		kvSlice = append(kvSlice, k)
		kvSlice = append(kvSlice, v)
	}

	addState := globalRedisPool.HSet(ctx, key, kvSlice)

	if addState.Err() != nil {
		err = addState.Err()
		return
	}

	expireState := globalRedisPool.Expire(ctx, key, expire)

	if expireState.Err() != nil {
		globalRedisPool.Del(ctx, key)
		err = expireState.Err()
		return
	}

	return

}

func GetHash(key string) (result map[string]string, err error) {
	ctx := context.Background()
	getARes := globalRedisPool.HGetAll(ctx, key)

	if getARes.Err() != nil {
		result = nil
		err = nil
		return
	}

	result = getARes.Val()
	err = nil
	return
}

func UpdateHash(key, field, val string) (err error) {
	ctx := context.Background()
	status := globalRedisPool.HSet(ctx, key, field, val)
	err = status.Err()
	return
}

func UnsetKey(key string) (err error) {
	ctx := context.Background()
	status := globalRedisPool.Del(ctx, key)
	err = status.Err()
	return
}

package redis

import (
	"context"
	"encoding/json"
	"fmt"
	providers_2 "github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/app"
	redis2 "github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

func Client() *redis2.Client {
	return providers_2.GetBean("redis").(app.Bean).GetBean("redis").(*services.Redis).Client
}

func Exists(key string) bool {
	ret, err := Client().Exists(context.Background(), key).Result()
	if err != nil {
		log.Error(err)
		return false
	}
	return ret == 1
}

func Transactions(fn func(pipe redis2.Pipeliner) error) ([]redis2.Cmder, error) {
	client := Client()
	return client.TxPipelined(client.Context(), fn)
}

func HExists(key, field string) bool {
	ret, err := Client().HExists(context.Background(), key, field).Result()
	if err != nil {
		log.Error(err)
		return false
	}
	return ret
}

func Set(key string, value interface{}, expiration time.Duration) {
	ret := Client().Set(context.Background(), key, value, expiration)
	if ret.Err() != nil {
		log.Error(ret.Err())
	}
}

func HSet(key string, values ...interface{}) int64 {
	ret, err := Client().HSet(context.Background(), key, values...).Result()
	if err != nil {
		log.Error(err)
	}
	return ret
}
func HDel(key string, fields ...string) int64 {
	ret, err := Client().HDel(context.Background(), key, fields...).Result()
	if err != nil {
		log.Error(err)
	}
	return ret
}

func HGet(key, field string) (string, bool) {
	ret, err := Client().HGet(context.Background(), key, field).Result()
	if err != nil {
		if err != redis2.Nil {
			log.Error(err)
		}
		return "", false
	}
	return ret, true
}

func HGetAll(key string) map[string]string {
	ret, err := Client().HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Error(err)
	}
	return ret
}

func Del(key string) {
	ret := Client().Del(context.Background(), key)
	if ret.Err() != nil {
		log.Error(ret.Err())
	}
}

// SAdd 集合
func SAdd(key string, members ...interface{}) {
	ret := Client().SAdd(context.Background(), key, members...)

	if ret.Err() != nil {
		log.Error(ret.Err())
	}
}

// SRem 删除 集合 key 中的元素 member1、member2、member3
func SRem(key string, members ...interface{}) {
	ret := Client().SRem(context.Background(), key, members...)

	if ret.Err() != nil {
		log.Error(ret.Err())
	}
}

func SetNx(key string, value interface{}, expiration time.Duration) bool {
	return Client().SetNX(context.Background(), key, value, expiration).Val()
}

func GetInt(key string) (int, bool) {
	i, err := Client().Get(context.Background(), key).Int()
	if err != nil {
		if err == redis2.Nil {
			return 0, false
		}
		log.Errorf("GetInt %v", err)
	}
	return i, true
}

func GetString(key string) string {
	cmd := Client().Get(context.Background(), key)
	err := cmd.Err()
	if err == redis2.Nil {
		log.Debug("key" + key + " does not exist")
		return ""
	} else if err != nil {
		log.Error(err)
	}
	return cmd.Val()
}

// Incr 自增1
func Incr(key string) int64 {
	cmd := Client().Incr(context.Background(), key)
	err := cmd.Err()
	if err == redis2.Nil {
		log.Debug("key" + key + " does not exist")
		return 0
	} else if err != nil {
		log.Error(err)
	}
	return cmd.Val()
}

// IncrBy 自增指定数字
func IncrBy(key string, value int64) int64 {
	cmd := Client().IncrBy(context.Background(), key, value)
	err := cmd.Err()
	if err == redis2.Nil {
		log.Debug("key" + key + " does not exist")
		return 0
	} else if err != nil {
		log.Error(err)
	}
	return cmd.Val()
}

func HIncrBy(key, field string, incr int64) int64 {
	cmd, err := Client().HIncrBy(context.Background(), key, field, incr).Result()
	if err != nil {
		log.Error(err)
	}
	return cmd
}

func HIncrByFloat(key, field string, incr float64) float64 {
	cmd, err := Client().HIncrByFloat(context.Background(), key, field, incr).Result()
	if err != nil {
		log.Error(err)
	}
	return cmd
}

func Expire(key string, expiration time.Duration) {
	Client().Expire(context.Background(), key, expiration)
}

func GetFloat64(key string) (float64, bool) {
	i, err := Client().Get(context.Background(), key).Float64()
	if err != nil {
		if err == redis2.Nil {
			return 0, false
		}
		log.Errorf("GetFloat64 %v", err)
	}
	return i, true
}

// CacheIsEqual 便捷判断缓存字符串是否相等, 只能判断map
func CacheIsEqual(data map[string]interface{}, key string) bool {
	cache := GetString(key)
	if cache == "" {
		return false
	}

	cacheMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(cache), &cacheMap)
	if err == nil {
		eq := true
		for s, i := range data {
			e1 := fmt.Sprintf("%v", cacheMap[s])
			e2 := fmt.Sprintf("%v", i)
			if e1 != e2 {
				eq = false
				break
			}
		}
		if eq {
			return true
		}
	}
	return false
}

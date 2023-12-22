package util

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	logger "github.com/sirupsen/logrus"
)

const (
	RedisTimeout time.Duration = time.Duration(60) * time.Second
)

func generateLockKey(key string) string {
	return fmt.Sprintf("%s_lock", key)
}

func AcquireSpinLock(redisClient *redis.Client, key string, timeout time.Duration, ttl time.Duration) int64 {
	currentTS := time.Now()
	expiredTS := currentTS.Add(timeout)
	lockValue := currentTS.UnixNano()
	lockKey := generateLockKey(key)
	redisCtx, _ := context.WithTimeout(context.Background(), RedisTimeout)
	for {
		ret, err := redisClient.SetNX(redisCtx, lockKey, lockValue, ttl).Result()
		if err != nil || !ret {
			if expiredTS.Before(time.Now()) {
				return 0
			}
			time.Sleep(time.Duration(5) * time.Millisecond)
		} else {
			return lockValue
		}
	}
}

func ReleaseSpinLock(redisClient *redis.Client, key string, lockValue int64) bool {
	lockKey := generateLockKey(key)
	redisCtx, _ := context.WithTimeout(context.Background(), RedisTimeout)
	value, err := redisClient.Get(redisCtx, lockKey).Result()
	if err != nil {
		if err == redis.Nil {
			return true
		} else {
			logger.Warnf("[orm][ReleaseSpinLock]get object %s failed:%s", key, err.Error())
			return false
		}
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false
	}
	if intValue == lockValue {
		value, err := redisClient.Del(redisCtx, lockKey).Result()
		if err != nil || value == 0 {
			return false
		}
		return true
	}
	return false
}

func GetObjectFromRedis(redisClient *redis.Client, key string, data interface{}) interface{} {
	redisCtx, _ := context.WithTimeout(context.Background(), RedisTimeout)
	value, err := redisClient.Get(redisCtx, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.Warnf("[orm][GetObjectFromRedis]get object %s failed:%s", key, err.Error())
		}
		return nil
	}
	if len(value) == 0 {
		return nil
	}
	err = json.Unmarshal([]byte(value), data)
	if err != nil {
		logger.Warnf("[orm][GetObjectFromRedis]get object %s json parse failed:%s", key, err.Error())
		return nil
	}
	return data
}

func SetObjectToRedis(redisClient *redis.Client, key string, data interface{}, ttl time.Duration) bool {
	dataStr, _ := json.Marshal(data)
	redisCtx, _ := context.WithTimeout(context.Background(), RedisTimeout)
	_, err := redisClient.Set(redisCtx, key, dataStr, ttl).Result()
	if err != nil {
		logger.Warnf("[orm][SetObjectToRedis]set object %s failed:%s", key, err.Error())
		return false
	}
	return true
}

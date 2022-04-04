package database_manager

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"inshortsProj/constant"
	"time"
)

var (
	redisClient *redis.Client
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     constant.REDIS_INSTANCE_URI,
		Password: constant.REDIS_AUTH,
		DB:       0,
	})
}

func SetinRedis(ctx *gin.Context, key string, value interface{}) error {
	duration := 30 * time.Minute
	err := redisClient.Set(key, value, duration).Err()
	if err != nil {
		fmt.Println("error in inserting to redis : " + err.Error())
		return err
	}
	return nil
}

func GetFromRedis(ctx *gin.Context, key string) (interface{}, error) {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		fmt.Println("error in getting from redis : " + err.Error())
		return nil, err
	}
	return val, nil
}

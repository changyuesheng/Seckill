package data

import (
	"fmt"
	"seckill/config"

	"github.com/go-redis/redis"
)

// 开启redis连接池
func initRedis(config config.AppConfig) *redis.Client {
	fmt.Println("---Start initializing redis connection---")

	redis_client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassowrd, // It's ok if password is "".
		DB:       config.Databasenum,   // use default DB
	})

	_, err := redis_client.Ping().Result()

	if err != nil {
		fmt.Println("init redis failed: ", err)
		panic(err)
	}

	if _, err := FlushAll(redis_client); err != nil {
		println("Error when flushAll. " + err.Error())
		panic(err)
	}

	fmt.Println("---Redis connection is initialized---")
	return redis_client
}

// 测试前删除缓存中的所有数据
func FlushAll(redis_client *redis.Client) (string, error) {
	return redis_client.FlushAll().Result()
}

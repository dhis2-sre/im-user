package storage

import (
	"fmt"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/go-redis/redis"
)

func ProvideRedis(c config.Config) *redis.Client {
	host := c.Redis.Host
	port := c.Redis.Port

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

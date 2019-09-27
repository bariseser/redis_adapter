package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type RedisAdapter struct {
	Client *redis.Client
}

var db *redis.Client

func ConnectRedis(redisHost string, redisPort int, database int, password string) *RedisAdapter {

	connectionStr := fmt.Sprintf("%s:%d", redisHost, redisPort)

	db := redis.NewClient(
		&redis.Options{
			Addr:     connectionStr,
			Password: password,
			DB:       database,
		})

	_, err := db.Ping().Result()

	if err != nil {
		log.Fatal(err.Error())
	}

	return &RedisAdapter{Client: db}
}

func (c *RedisAdapter) Get(key string) (string, error) {
	key, err := c.Client.Get(key).Result()
	return key, err
}

func (c *RedisAdapter) Set(key string, value interface{}, duration time.Duration) (string, error) {
	key, err := c.Client.Set(key, value, duration).Result()
	return key, err
}

func (c *RedisAdapter) Del(key string) error {
	_, err := c.Client.Del(key).Result()
	return err
}

func (c *RedisAdapter) Ping() (string, error) {
	status, err := c.Client.Ping().Result()

	if err != nil {
		return status, err
	}
	return status, nil
}

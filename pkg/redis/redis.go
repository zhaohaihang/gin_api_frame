package redis

import (
	"gin_api_frame/config"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
)

var Pool *redis.Pool

func NewRedisPool(config *config.Config) (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle:     20,
		MaxActive:   256,
		IdleTimeout: time.Duration(15) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", config.Redis.RedisAddr,
				redis.DialPassword(config.Redis.RedisPw),
				redis.DialConnectTimeout(time.Duration(30)*time.Second),
				redis.DialReadTimeout(time.Duration(30)*time.Second),
				redis.DialWriteTimeout(time.Duration(30)*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	Pool = pool
	return pool, nil
}

var RedisPoolProviderSet = wire.NewSet(NewRedisPool)

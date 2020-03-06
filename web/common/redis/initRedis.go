package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_api_frame/web/common/config"
	"log"
	"strings"
	"time"
)

var MyRedis *redis.Pool

func InitRedis() {
	err := startRedis()
	if err != nil {
		log.Fatal("init redis err:", err)
	}
}

func startRedis() error {
	MyRedis = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp", config.MyConfig.RedisHost)
			if err != nil {
				return nil, err
			}
			if len(strings.Join(strings.Fields(strings.TrimSpace(config.MyConfig.RedisPwd)), "")) != 0 {
				if _, err := c.Do("AUTH", config.MyConfig.RedisPwd); err != nil {
					err := c.Close()
					if err != nil {
						return nil, err
					}
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         config.MyConfig.RedisMaxIde,
		MaxActive:       config.MyConfig.RedisActive,
		IdleTimeout:     config.MyConfig.RedisIdleTimeout,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	return nil
}

func SetCache(key, value string, time int) error {
	conn := MyRedis.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("redis close err:", err)
		}
	}()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func ExistCache(key string) bool {
	conn := MyRedis.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("redis close err:", err)
		}
	}()
	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		fmt.Println("redis exists err:", err)
		return false
	}
	return ok
}

func GetCache(key string) (string, error) {
	conn := MyRedis.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("redis close err:", err)
		}
	}()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get err:", err)
		return "", err
	}
	return value, nil
}

func Delete(key string) (bool, error) {
	conn := MyRedis.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("redis close err:", err)
		}
	}()
	return redis.Bool(conn.Do("DEL", key))
}

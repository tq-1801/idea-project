package util

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var RedisPool *redis.Pool //创建redis连接池

func RedisInit() {
	RedisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", Cfg.Redis.RedisIps)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", Cfg.Redis.RedisPwd); err != nil {
				conn.Close()
				return nil, err
			}
			//if _, err := c.Do("SELECT",1); err != nil {
			// c.Close()
			// return nil, err
			//}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

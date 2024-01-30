package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func initRedisPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数目
		MaxActive:   maxActive,   //最大连接数
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", address)
		},
	}
	fmt.Println("redis初始化完成")
}

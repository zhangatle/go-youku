package redisClient

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 直接连接
func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", beego.AppConfig.String("redisdb"),redis.DialPassword("Itzler.666"))
	return pool
}


// 通过连接池连接
func PoolConnect() redis.Conn {
	// 建立连接池
	pool := &redis.Pool{
		// 最大空闲连接数
		MaxIdle: 5000,
		// 最大连接数
		MaxActive: 10000,
		// 空闲连接超时时间
		IdleTimeout: 180 * time.Second,
		// 超时是否等待
		Wait: true,
		// 建立连接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redisdb"),redis.DialPassword("Itzler.666"))
			if err != nil {
				return nil, err
			}
			// 选择db
			_, _ = c.Do("SELECT", 0)
			return c, nil
		},
	}
	return pool.Get()
}
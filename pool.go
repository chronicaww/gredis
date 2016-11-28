package gredis

import (
	"time"

	// "github.com/astaxie/beego"

	"github.com/garyburd/redigo/redis"
)

var (
	// RedisClient 连接池
	RedisClient *redis.Pool
	// RedisHost 地址
	RedisHost string
	// RedisDB 数据库编号
	RedisDB int
)

// InitPool 初始化连接池
func InitPool(host string, redisDB, maxIdle, maxActive int, idleTimeout time.Duration) {
	// 从配置文件获取redis的ip以及db
	// RedisHost = beego.AppConfig.String("redis.host")
	// RedisDB, _ = beego.AppConfig.Int("redis.db")
	RedisHost = host
	RedisDB = redisDB
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		// MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		// MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDB)
			return c, nil
		},
	}
}

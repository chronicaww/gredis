package gredis

import "github.com/garyburd/redigo/redis"

// HGet ..
func (rc *RedisConn) HGet(key, field string) string {
	s, _ := redis.String(rc.Get().Do("HGET", key, field))
	return s
}

// HSet ..
func (rc *RedisConn) HSet(key, field, value string) error {
	_, e := rc.Get().Do("HSET", key, field, value)
	return e
}

// HMSet ..
func (rc *RedisConn) HMSet(args ...interface{}) error {
	_, e := rc.Get().Do("HMSET", args...)
	return e
}

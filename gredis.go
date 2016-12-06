package gredis

import (
	"errors"
	"time"

	"github.com/chronicaww/gredis/funcAnysis"
	"github.com/garyburd/redigo/redis"
)

// RedisConn 连接
type RedisConn struct {
	conn redis.Conn
}

// Delay 延迟警告时长
var Delay = int64(1)

// Get 获得连接
func (rc *RedisConn) Get() redis.Conn {
	tm := time.Now()
	if rc.conn == nil || rc.conn.Flush() != nil {
		rc.conn = RedisClient.Get()
	}
	funcAnysis.Instance().Once("rc.Get", tm)
	return rc.conn
}

// CheckConn 检测连接
func (rc *RedisConn) CheckConn() error {
	if rc == nil {
		return errors.New("rc is nil")
	}
	if rc.conn == nil {
		return errors.New("conn is nil")
	}
	return rc.conn.Flush()
}

// Destroy 销毁
func (rc *RedisConn) Destroy() {
	if rc != nil {
		if nil != rc.conn {
			rc.conn.Close()
		}
		rc = nil
	}
}

// Close 关闭连接
func (rc *RedisConn) Close() {
	if rc != nil && nil != rc.conn {
		rc.conn.Close()
	}
}

//IsMembers 判定是否全包含
func IsMembers(rc *RedisConn, key string, b [][]byte) bool {
	for _, v := range b {
		if !ChkInSet(rc, key, string(v)) {
			return false
		}
	}
	return true
}

//SetValue 设置一个key值
func SetValue(rc *RedisConn, key string, value []byte) error {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("SET", key, value)
	funcAnysis.Instance().Once("gredis.SetValue", tm)
	return e
}

//GetValue 获取一个key值
func GetValue(rc *RedisConn, key string) ([]byte, error) {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, e := redis.Bytes(rc.Get().Do("GET", key))
	funcAnysis.Instance().Once("gredis.GetValue", tm)
	return result, e
}

//GetValues 模糊获取key值，以*作为模糊标记
func GetValues(rc *RedisConn, key string) ([][]byte, error) {
	tm := time.Now()
	if nil == rc {
		rc = new(RedisConn)
	}
	result, e := redis.ByteSlices(rc.Get().Do("GET", key))
	funcAnysis.Instance().Once("gredis.GetValues", tm)
	return result, e
}

//GetKeys 获取key名
func GetKeys(rc *RedisConn, v string) []string {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Strings(rc.Get().Do("KEYS", v))
	funcAnysis.Instance().Once("gredis.GetKeys", tm)
	return result
}

//DelValue 删除一个key值
func DelValue(rc *RedisConn, key string) error {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("DEL", key)
	funcAnysis.Instance().Once("gredis.DelValue", tm)
	return e
}

//DelValues 模糊删除key值，以*作为模糊标记
func DelValues(rc *RedisConn, key string) error {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	keyArr := GetKeys(rc, key)

	for _, v := range keyArr {
		e := DelValue(rc, v)
		if e != nil {
			return e
		}
	}
	funcAnysis.Instance().Once("gredis.DelValues", tm)
	return nil
}

//AddSet 向指定的set添加值
func AddSet(rc *RedisConn, key, value string) error {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("SADD", key, []byte(value))
	funcAnysis.Instance().Once("gredis.AddSet", tm)
	return e
}

//AddListSet 向指定的set加入一多值
func AddListSet(rc *RedisConn, key string, values [][]byte) error {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	for _, v := range values {
		_, e := rc.Get().Do("SADD", key, v)
		if e != nil {
			return e
		}
	}
	funcAnysis.Instance().Once("gredis.AddListSet", tm)
	return nil
}

//GetSet 获取set
func GetSet(rc *RedisConn, key string) [][]byte {
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.ByteSlices(rc.Get().Do("SMEMBERS", key))
	return result
}

//GetSetString 获取set
func GetSetString(rc *RedisConn, key string) []string {
	tm := time.Now()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Strings(rc.Get().Do("SMEMBERS", key))
	funcAnysis.Instance().Once("gredis.GetSetString", tm)
	return result
}

//GetSetInt 获取set
func GetSetInt(rc *RedisConn, key string) []int {
	tm := time.Now()
	result, _ := redis.Ints(rc.Get().Do("SMEMBERS", key))
	funcAnysis.Instance().Once("gredis.GetSetInt", tm)
	return result
}

//GetSetLen 获取set长度
func GetSetLen(rc *RedisConn, key string) int64 {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetSetLen", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Int64(rc.Get().Do("SCARD", key))
	return result
}

//ChkInSet 判定一个值是否在set中
func ChkInSet(rc *RedisConn, key string, v string) bool {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.ChkInSet", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Bool(rc.Get().Do("SISMEMBER", key, v))
	return result
}

//DelSet 删除set
func DelSet(rc *RedisConn, key string) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.DelSet", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	e := DelValue(rc, key)
	return e

}

//RemSetValue 从set中删除值
func RemSetValue(rc *RedisConn, key string, v string) (bool, error) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.RemSetValue", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.Bool(rc.Get().Do("SREM", key, []byte(v)))
}

//FlushDB 刷新db
func FlushDB(rc *RedisConn) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.FlushDB", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	rc.Get().Do("FLUSHDB")
}

// GetList 获得列表内容
func GetList(rc *RedisConn, key string) [][]byte {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetList", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.ByteSlices(rc.Get().Do("LRANGE", key, 0, -1))
	return result
}

// GetListString 获得列表内容
func GetListString(rc *RedisConn, key string) []string {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetListString", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Strings(rc.Get().Do("LRANGE", key, 0, -1))
	return result
}

// AddList 列表
func AddList(rc *RedisConn, key string, v string, pos int) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.AddList", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	var e error
	if pos == 0 {
		_, e = rc.Get().Do("LPUSH", key, v)
	} else if pos < 0 {
		_, e = rc.Get().Do("RPUSH", key, v)
	} else {
		v0, _ := redis.Int(rc.Get().Do("LINDEX", key, pos))
		_, e = rc.Get().Do("LINSERT", key, "BEFORE", v0, v)
	}
	return e
}

// RemListValue 从列表中移除一个值
func RemListValue(rc *RedisConn, key string, v string) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.RemListValue", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("LREM", key, 0, v)
	return e
}

//FindInList 获得在列表中位置
func FindInList(rc *RedisConn, key, v string) int {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.FindInList", tm)
	list := GetListString(rc, key)
	for i, v0 := range list {
		if v0 == v {
			return i
		}
	}
	return -1
}

// AddScore 添加一个排行榜分数
func AddScore(rc *RedisConn, key string, score float64, v []byte) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.AddScore", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("ZADD", key, score, v)
	return e
}

// RemScore 移除一个排行榜分数
func RemScore(rc *RedisConn, key string, v []byte) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.RemScore", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("ZREM", key, v)
	return e
}

// GetRangeByRank 获取n名排行
func GetRangeByRank(rc *RedisConn, key string, from, to int64, bAsc bool) ([][]byte, error) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetRangeByRank", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	if bAsc { //升序
		return redis.ByteSlices(rc.Get().Do("ZRANGE", key, from, to))
	}
	return redis.ByteSlices(rc.Get().Do("ZREVRANGE", key, from, to))
}

// GetRangeByScore 指定分数区间的排名
func GetRangeByScore(rc *RedisConn, key string, from, to float64) ([][]byte, error) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetRangeByScore", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.ByteSlices(rc.Get().Do("ZRANGEBYSCORE", key, from, to))
}

// GetRangeByScoreN 指定分数区间的排名
func GetRangeByScoreN(rc *RedisConn, key string, from, to float64, asc bool) ([][]byte, error) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetRangeByScreN", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	if asc {
		return redis.ByteSlices(rc.Get().Do("ZRANGEBYSCORE", key, from, to))
	}
	return redis.ByteSlices(rc.Get().Do("ZREVRANGEBYSCORE", key, from, to))
}

// GetScore 获取分数
func GetScore(rc *RedisConn, key string, v []byte) (float64, error) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetScore", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.Float64(rc.Get().Do("ZSCORE", key, v))
}

// GetRank 获取排名
func GetRank(rc *RedisConn, key string, v []byte, bAsc bool) (int64, error) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetRank", tm)
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.Int64(rc.Get().Do("ZREVRANK", key, v))
}

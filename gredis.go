package gredis

import "github.com/garyburd/redigo/redis"

// RedisConn 连接
type RedisConn struct {
	conn redis.Conn
}

// var dbID = 1

// Delay 延迟警告时长
var Delay = int64(1)

// func getClient() redis.Client {
// 	spec := redis.DefaultSpec().Db(dbID).Password("")
// 	client, e := redis.NewSynchClientWithSpec(spec)
//
// 	if e != nil {
// 		fmt.Println("failed to create the client", e)
// 	}
// 	return client
// }

// func connToRedis() {
// if _self.client != nil {
// 	_self.client.Quit()
// }
// spec := redis.DefaultSpec().Db(dbID).Password("")
// client, e := redis.NewSynchClientWithSpec(spec)
// if e != nil {
// fmt.Println("failed to create the client", e)
// }

// _self.client = client
// }

// func connToRedis2() {
// if _self.client2 != nil {
// 	_self.client2.Quit()
// }
// spec := redis.DefaultSpec().Db(dbID).Password("")
// client, e := redis.NewSynchClientWithSpec(spec)
// if e != nil {
// 	fmt.Println("failed to create the client", e)
// }

// 	_self.client2 = client
// }

// Get 获得连接
func (rc *RedisConn) Get() redis.Conn {
	if rc.conn == nil || rc.conn.Flush() != nil {
		rc.conn = RedisClient.Get()
	}
	return rc.conn
}

// Destroy 销毁
func (rc *RedisConn) Destroy() {
	rc.conn.Close()
	rc = nil
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
	// tm := time.Now().Unix()
	//
	// client := getClient()
	// defer client.Quit()
	// e := client.Set(key, value)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("SetValueB Delay:", d, len(value), ToString(value), ToInt32(value))
	// 	tm = time.Now().Unix()
	// 	connToRedis()
	// }
	// if e != nil {
	// 	return e
	// }
	// d = time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("SetValueC Delay:", d)
	// }
	//
	// return nil
	// rc := RedisClient.Get()
	// defer rc.Close()

	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("SET", key, value)
	return e
}

//GetValue 获取一个key值
func GetValue(rc *RedisConn, key string) ([]byte, error) {
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	// result, e := client.Get(key)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetValue Delay:", d)
	// }
	// return result, e
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.Bytes(rc.Get().Do("GET", key))
}

//GetValues 模糊获取key值，以*作为模糊标记
func GetValues(key string) ([][]byte, error) {
	// tm := time.Now().Unix()
	//
	// result := [][]byte{}
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	// keyArr, e := client.Keys(key)
	// if e != nil {
	// 	return result, e
	// }
	// for _, v := range keyArr {
	// 	val, e := client.Get(v)
	// 	if e != nil {
	// 		return result, e
	// 	}
	// 	result = append(result, val)
	// }
	//
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetValues Delay:", d)
	// }
	//
	// return result, nil
	rc := RedisClient.Get()
	defer rc.Close()
	return redis.ByteSlices(rc.Do("GET", key))
}

//GetKeys 获取key名
func GetKeys(rc *RedisConn, v string) []string {
	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	// keyArr, e := client.Keys(v)
	// if e != nil {
	// 	// keyArr, e = c.client2.Keys(v)
	// 	// if e != nil {
	// 	// fmt.Printf("gredis.GetKeys2 Error:%v\n", e)
	// 	// connToRedis2()
	// 	return []string{}
	// 	// }
	// }
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetKeys Delay:", d)
	// }
	// return keyArr
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Strings(rc.Get().Do("KEYS", v))
	return result
}

//DelValue 删除一个key值
func DelValue(rc *RedisConn, key string) error {
	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	//
	// _, e := client.Del(key)
	// // if e != nil {
	// // connToRedis()
	// // _, e = c.client2.Del(key)
	// // if e != nil {
	// // 	connToRedis2()
	// // }
	// // }
	//
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("DelValue Delay:", d)
	// }
	//
	// return e
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("DEL", key)
	return e
}

//DelValues 模糊删除key值，以*作为模糊标记
func DelValues(rc *RedisConn, key string) error {
	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()

	// keyArr, e := client.Keys(key)
	// if e != nil {
	// connToRedis()
	// return e
	// }
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
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("DelValues Delay:", d)
	// }
	return nil
}

//AddSet 向指定的set添加值
func AddSet(rc *RedisConn, key, value string) error {

	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	//
	// _, e := client.Sadd(key, []byte(value))
	// if e != nil {
	// 	// connToRedis()
	// 	// _, e = client2.Sadd(key, []byte(value))
	// 	return e
	// }
	//
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("AddSet Delay:", d)
	// }
	// return nil
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("SADD", key, []byte(value))
	return e
}

//AddListSet 向指定的set加入一多值
func AddListSet(rc *RedisConn, key string, values [][]byte) error {
	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	//
	// for _, v := range values {
	// 	_, e := client.Sadd(key, v)
	// 	if e != nil {
	// 		// 	connToRedis()
	// 		// 	_, e = c.client2.Sadd(key, v)
	// 		return e
	// 	}
	// }
	//
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("AddListSet Delay:", d)
	// }
	// return nil
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	for _, v := range values {
		_, e := rc.Get().Do("SADD", key, v)
		if e != nil {
			return e
		}
	}
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
	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	//
	// result, _ := client.Smembers(key)
	// // if e != nil {
	// // 	connToRedis()
	// // 	result, e = c.client2.Smembers(key)
	// // 	if e != nil {
	// // 		connToRedis2()
	// // 	}
	// // }
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetSet Delay:", d)
	// }
	// return result
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Strings(rc.Get().Do("SMEMBERS", key))
	return result
}

//GetSetInt 获取set
func GetSetInt(rc *RedisConn, key string) []int {
	// rc := RedisClient.Get()
	// defer rc.Close()
	result, _ := redis.Ints(rc.Get().Do("SMEMBERS", key))
	return result
}

//GetSetLen 获取set长度
func GetSetLen(rc *RedisConn, key string) int64 {
	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	//
	// result, _ := client.Scard(key)
	// // if e != nil {
	// // 	connToRedis()
	// // 	result, e = c.client2.Scard(key)
	// // 	if e != nil {
	// // 		connToRedis2()
	// // 	}
	// // }
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetSetLen Delay:", d)
	// }
	// return result
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Int64(rc.Get().Do("SCARD", key))
	return result
}

//ChkInSet 判定一个值是否在set中
func ChkInSet(rc *RedisConn, key string, v string) bool {
	// 	tm := time.Now().Unix()
	// 	// c := Instance()
	// 	client := getClient()
	// 	defer client.Quit()
	//
	// 	bIs, _ := client.Sismember(key, []byte(v))
	//
	// 	// if e != nil {
	// 	// 	connToRedis()
	// 	// 	bIs, e = c.client.Sismember(key, []byte(v))
	// 	// 	if e != nil {
	// 	// 		connToRedis2()
	// 	// 	}
	// 	// }
	// 	d := time.Now().Unix() - tm
	// 	if d > Delay {
	// 		fmt.Println("ChkInSet Delay:", d)
	// 	}
	// 	return bIs
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Bool(rc.Get().Do("SISMEMBER", key, v))
	return result
}

//DelSet 删除set
func DelSet(rc *RedisConn, key string) error {
	if rc == nil {
		rc = new(RedisConn)
	}
	e := DelValue(rc, key)
	return e

}

//RemSetValue 从set中删除值
func RemSetValue(rc *RedisConn, key string, v string) (bool, error) {
	// tm := time.Now().Unix()
	// // c := Instance()
	// client := getClient()
	// defer client.Quit()
	// bRem, e := client.Srem(key, []byte(v))
	// // if e != nil {
	// // 	connToRedis()
	// // 	bRem, e = c.client2.Srem(key, []byte(v))
	// // 	if e != nil {
	// // 		connToRedis2()
	// // 	}
	// // }
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("RemSetValue Delay:", d)
	// }
	// return bRem, e
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.Bool(rc.Get().Do("SREM", key, []byte(v)))
}

//FlushDB 刷新db
func FlushDB(rc *RedisConn) {
	// c := Instance(dbID)
	// c.client.Flushdb()
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	rc.Get().Do("FLUSHDB")
}

// GetList 获得列表内容
func GetList(rc *RedisConn, key string) [][]byte {
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	// len, e := client.Llen(key)
	// if e != nil {
	// 	return [][]byte{}
	// }
	// result, _ := client.Lrange(key, int64(0), len)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetList Delay:", d)
	// }
	// return result
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.ByteSlices(rc.Get().Do("LRANGE", key, 0, -1))
	return result
}

// GetListString 获得列表内容
func GetListString(rc *RedisConn, key string) []string {
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	result, _ := redis.Strings(rc.Get().Do("LRANGE", key, 0, -1))
	return result
}

// AddList 列表
func AddList(rc *RedisConn, key string, v string, pos int) error {
	// rc := RedisClient.Get()
	// defer rc.Close()
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
	// 	tm := time.Now().Unix()
	// 	client := getClient()
	// 	defer client.Quit()
	//
	// 	if pos == 0 {
	// 		return client.Lpush(key, []byte(v))
	// 	}
	// 	if pos < 0 {
	// 		return client.Rpush(key, []byte(v))
	// 	}
	//
	// 	len, e := client.Llen(key)
	// 	if e != nil {
	// 		return e
	// 	}
	// 	tmpLeft, e := client.Lrange(key, int64(0), int64(pos-1))
	// 	if e != nil {
	// 		return e
	// 	}
	// 	// fmt.Println("lenl:", tmpLeft)
	//
	// 	tmpRight, e := client.Lrange(key, int64(pos), len)
	// 	if e != nil {
	// 		return e
	// 	}
	// 	// fmt.Println("lenr:", tmpRight)
	//
	// 	result := make([][]byte, len+1)
	// 	copy(result, tmpLeft)
	// 	copy(result[pos+1:], tmpRight)
	// 	result[pos] = []byte(v)
	// 	DelSet(key)
	//
	// 	for _, v0 := range result {
	// 		client.Rpush(key, v0)
	// 	}
	// 	d := time.Now().Unix() - tm
	// 	if d > Delay {
	// 		fmt.Println("AddList Delay:", d)
	// 	}
	// 	return nil
}

// RemListValue 从列表中移除一个值
func RemListValue(rc *RedisConn, key string, v string) error {
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	// _, e := client.Lrem(key, []byte(v), int64(0))
	//
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("RemListValue Delay:", d)
	// }
	// fmt.Println("nnn:", key, []byte(v), v, n)
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("LREM", key, 0, v)
	return e
}

//FindInList 获得在列表中位置
func FindInList(rc *RedisConn, key, v string) int {
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
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("ZADD", key, score, v)
	return e
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	// _, e := client.Zadd(key, score, v)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("AddScore Delay:", d)
	// }
	// return e
}

// RemScore 移除一个排行榜分数
func RemScore(rc *RedisConn, key string, v []byte) error {
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	_, e := rc.Get().Do("ZREM", key, v)
	return e
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	// _, e := client.Zrem(key, v)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("RemScore Delay:", d)
	// }
	// return e
}

// GetRangeByRank 获取n名排行
func GetRangeByRank(rc *RedisConn, key string, from, to int64, bAsc bool) ([][]byte, error) {
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	if bAsc { //升序
		return redis.ByteSlices(rc.Get().Do("ZRANGE", key, from, to))
	}
	return redis.ByteSlices(rc.Get().Do("ZREVRANGE", key, from, to))
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	//
	// if bAsc {
	// 	result, e := client.Zrange(key, from, to)
	// 	d := time.Now().Unix() - tm
	// 	if d > Delay {
	// 		fmt.Println("GetRangeByRank Delay:", d)
	// 	}
	// 	return result, e
	// }
	// result, e := client.Zrevrange(key, from, to)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetRangeByRank Delay:", d)
	// }
	// return result, e
}

// GetRangeByScore 指定分数区间的排名
func GetRangeByScore(rc *RedisConn, key string, from, to float64) ([][]byte, error) {
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.ByteSlices(rc.Get().Do("ZRANGEBYSCORE", key, from, to))
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	//
	// result, e := client.Zrangebyscore(key, from, to)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetRangeByScore Delay:", d)
	// }
	//
	// return result, e
}

// GetRangeByScoreN 指定分数区间的排名
func GetRangeByScoreN(rc *RedisConn, key string, from, to float64, asc bool) ([][]byte, error) {
	// rc := RedisClient.Get()
	// defer rc.Close()
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
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.Float64(rc.Get().Do("ZSCORE", key, v))
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	//
	// result, e := client.Zscore(key, v)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetScore Delay:", d)
	// }
	// return result, e
}

// GetRank 获取排名
func GetRank(rc *RedisConn, key string, v []byte, bAsc bool) (int64, error) {
	// rc := RedisClient.Get()
	// defer rc.Close()
	if rc == nil {
		rc = new(RedisConn)
	}
	return redis.Int64(rc.Get().Do("ZREVRANK", key, v))
	// tm := time.Now().Unix()
	// client := getClient()
	// defer client.Quit()
	//
	// if bAsc {
	// 	result, e := client.Zrank(key, v)
	// 	return result, e
	// }
	// result, e := client.Zrevrank(key, v)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("GetRank Delay:", d)
	// }
	// return result, e
}

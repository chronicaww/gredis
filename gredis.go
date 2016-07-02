package gredis

import (
	"fmt"
	"github.com/chronicaww/Go-Redis"
	"time"
)

//GRedis 关系数据库结构
type GRedis struct {
	client  redis.Client
	client2 redis.Client
}

var _self *GRedis
var dbID = 1
var Delay = int64(1)

//Instance 获取实例
func Instance(id int) *GRedis {
	if _self == nil {
		dbID = id
		_self = new(GRedis)
		// spec := redis.DefaultSpec().Db(1).Password("")
		// client, e := redis.NewSynchClientWithSpec(spec)
		// if e != nil {
		// 	fmt.Println("failed to create the client", e)
		// 	// return
		// }
		connToRedis()
		connToRedis2()
		// _self.client = client
	}
	return _self
}

func getClient() redis.Client {
	spec := redis.DefaultSpec().Db(dbID).Password("")
	client, e := redis.NewSynchClientWithSpec(spec)

	if e != nil {
		fmt.Println("failed to create the client", e)
	}
	return client
}

func connToRedis() {
	// if _self.client != nil {
	// 	_self.client.Quit()
	// }
	spec := redis.DefaultSpec().Db(dbID).Password("")
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		fmt.Println("failed to create the client", e)
	}

	_self.client = client
}

func connToRedis2() {
	// if _self.client2 != nil {
	// 	_self.client2.Quit()
	// }
	spec := redis.DefaultSpec().Db(dbID).Password("")
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		fmt.Println("failed to create the client", e)
	}

	_self.client2 = client
}

//IsMembers 判定是否全包含
func IsMembers(key string, b [][]byte) bool {
	for _, v := range b {
		if !ChkInSet(key, string(v)) {
			return false
		}
	}
	return true
}

//SetValue 设置一个key值
func SetValue(key string, value []byte) error {
	tm := time.Now().Unix()
	// c := Instance(dbID)
	// d := time.Now().Unix() - tm
	// if d > Delay {
	// 	fmt.Println("SetValueA Delay:", d)
	// 	tm = time.Now().Unix()
	// }

	client := getClient()
	defer client.Quit()
	e := client.Set(key, value)
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("SetValueB Delay:", d, len(value), ToString(value), ToInt32(value))
		tm = time.Now().Unix()
		connToRedis()
	}
	if e != nil {
		return e
		// 	connToRedis()
		// 	d = time.Now().Unix() - tm
		// 	if d > Delay {
		// 		fmt.Println("SetValueD1 Delay:", d)
		// 		tm = time.Now().Unix()
		// 	}
		// 	e := c.client2.Set(key, value)
		// 	if e != nil {
		// 		return e
		// 	}
		// 	d = time.Now().Unix() - tm
		// 	if d > Delay {
		// 		fmt.Println("SetValueD2 Delay:", d)
		// 		tm = time.Now().Unix()
		// 	}
	}
	d = time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("SetValueC Delay:", d)
	}

	return nil
}

//GetValue 获取一个key值
func GetValue(key string) ([]byte, error) {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()
	result, e := client.Get(key)
	// if e != nil {
	// 	fmt.Printf("gredis.GetValue Error:%v\n", e)
	// 	connToRedis()
	// 	result, e = c.client2.Get(key)
	// 	if e != nil {
	// 		fmt.Printf("gredis.GetValue2 Error:%v\n", e)
	// 		connToRedis2()
	// 	}
	// 	// tm := 0
	// 	// for {
	// 	// 	if tm > 3 {
	// 	// 		break
	// 	// 	}
	// 	// 	tm++
	// 	// 	result, e = c.client.Get(key)
	// 	// 	if e != nil {
	// 	// 		fmt.Printf("gredis.GetValue Error:%v\n", e)
	// 	// 		connToRedis()
	// 	// 		continue
	// 	// 	}
	// 	// 	break
	// 	// }
	// }
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("GetValue Delay:", d)
	}
	return result, e
}

//GetValues 模糊获取key值，以*作为模糊标记
func GetValues(key string) ([][]byte, error) {
	tm := time.Now().Unix()

	result := [][]byte{}
	// c := Instance()
	client := getClient()
	defer client.Quit()
	keyArr, e := client.Keys(key)
	if e != nil {
		return result, e
	}
	for _, v := range keyArr {
		val, e := client.Get(v)
		if e != nil {
			return result, e
		}
		result = append(result, val)
	}

	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("GetValues Delay:", d)
	}

	return result, nil
}

//GetKeys 获取key名
func GetKeys(v string) []string {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()
	keyArr, e := client.Keys(v)
	if e != nil {
		// keyArr, e = c.client2.Keys(v)
		// if e != nil {
		// fmt.Printf("gredis.GetKeys2 Error:%v\n", e)
		// connToRedis2()
		return []string{}
		// }
	}
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("GetKeys Delay:", d)
	}
	return keyArr
}

//DelValue 删除一个key值
func DelValue(key string) error {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()

	_, e := client.Del(key)
	// if e != nil {
	// connToRedis()
	// _, e = c.client2.Del(key)
	// if e != nil {
	// 	connToRedis2()
	// }
	// }

	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("DelValue Delay:", d)
	}

	return e
}

//DelValues 模糊删除key值，以*作为模糊标记
func DelValues(key string) error {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()

	keyArr, e := client.Keys(key)
	if e != nil {
		// connToRedis()
		return e
	}
	for _, v := range keyArr {
		e := DelValue(v)
		if e != nil {
			return e
		}
		// _, e := c.client.Del(v)
		// if e != nil {
		// 	connToRedis()
		// 	_,e=c.client2.Del(v)
		// 	return e
		// }
	}
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("DelValues Delay:", d)
	}
	return nil
}

//AddSet 向指定的set添加值
func AddSet(key, value string) error {

	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()

	_, e := client.Sadd(key, []byte(value))
	if e != nil {
		// connToRedis()
		// _, e = client2.Sadd(key, []byte(value))
		return e
	}

	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("AddSet Delay:", d)
	}
	return nil
}

//AddListSet 向指定的set加入一多值
func AddListSet(key string, values [][]byte) error {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()

	for _, v := range values {
		_, e := client.Sadd(key, v)
		if e != nil {
			// 	connToRedis()
			// 	_, e = c.client2.Sadd(key, v)
			return e
		}
	}

	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("AddListSet Delay:", d)
	}
	return nil
}

//GetSet 获取set
func GetSet(key string) [][]byte {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()

	result, _ := client.Smembers(key)
	// if e != nil {
	// 	connToRedis()
	// 	result, e = c.client2.Smembers(key)
	// 	if e != nil {
	// 		connToRedis2()
	// 	}
	// }
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("GetSet Delay:", d)
	}
	return result
}

//GetSetLen 获取set长度
func GetSetLen(key string) int64 {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()

	result, _ := client.Scard(key)
	// if e != nil {
	// 	connToRedis()
	// 	result, e = c.client2.Scard(key)
	// 	if e != nil {
	// 		connToRedis2()
	// 	}
	// }
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("GetSetLen Delay:", d)
	}
	return result
}

//ChkInSet 判定一个值是否在set中
func ChkInSet(key string, v string) bool {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()

	bIs, _ := client.Sismember(key, []byte(v))

	// if e != nil {
	// 	connToRedis()
	// 	bIs, e = c.client.Sismember(key, []byte(v))
	// 	if e != nil {
	// 		connToRedis2()
	// 	}
	// }
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("ChkInSet Delay:", d)
	}
	return bIs
}

//DelSet 删除set
func DelSet(key string) error {
	e := DelValue(key)

	return e

}

//RemSetValue 从set中删除值
func RemSetValue(key string, v string) (bool, error) {
	tm := time.Now().Unix()
	// c := Instance()
	client := getClient()
	defer client.Quit()
	bRem, e := client.Srem(key, []byte(v))
	// if e != nil {
	// 	connToRedis()
	// 	bRem, e = c.client2.Srem(key, []byte(v))
	// 	if e != nil {
	// 		connToRedis2()
	// 	}
	// }
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("RemSetValue Delay:", d)
	}
	return bRem, e
}

//FlushDB 刷新db
func FlushDB() {
	c := Instance(dbID)
	c.client.Flushdb()
}

// GetList 获得列表内容
func GetList(key string) [][]byte {
	tm := time.Now().Unix()
	client := getClient()
	defer client.Quit()
	len, e := client.Llen(key)
	if e != nil {
		return [][]byte{}
	}
	result, _ := client.Lrange(key, int64(0), len)
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("GetList Delay:", d)
	}
	return result
}

// AddList 列表
func AddList(key string, v string, pos int) error {
	tm := time.Now().Unix()
	client := getClient()
	defer client.Quit()

	if pos == 0 {
		return client.Lpush(key, []byte(v))
	}
	if pos < 0 {
		return client.Rpush(key, []byte(v))
	}

	len, e := client.Llen(key)
	if e != nil {
		return e
	}
	tmpLeft, e := client.Lrange(key, int64(0), int64(pos-1))
	if e != nil {
		return e
	}
	// fmt.Println("lenl:", tmpLeft)

	tmpRight, e := client.Lrange(key, int64(pos), len)
	if e != nil {
		return e
	}
	// fmt.Println("lenr:", tmpRight)

	result := make([][]byte, len+1)
	copy(result, tmpLeft)
	copy(result[pos+1:], tmpRight)
	result[pos] = []byte(v)
	DelSet(key)

	for _, v0 := range result {
		client.Rpush(key, v0)
	}
	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("AddList Delay:", d)
	}
	return nil
}

// RemListValue 从列表中移除一个值
func RemListValue(key string, v string) error {
	tm := time.Now().Unix()
	client := getClient()
	defer client.Quit()
	_, e := client.Lrem(key, []byte(v), int64(0))

	d := time.Now().Unix() - tm
	if d > Delay {
		fmt.Println("RemListValue Delay:", d)
	}
	// fmt.Println("nnn:", key, []byte(v), v, n)
	return e
}

//FindInList 获得在列表中位置
func FindInList(key, v string) int {
	client := getClient()
	defer client.Quit()
	list := GetList(key)
	for i, v0 := range list {
		if string(v0) == v {
			return i
		}
	}
	return -1
}

// AddScore 添加一个排行榜分数
func AddScore(key string, score float64, v []byte) error {
	client := getClient()
	defer client.Quit()
	_, e := client.Zadd(key, score, v)
	return e
}

// RemScore 移除一个排行榜分数
func RemScore(key string, v []byte) error {
	client := getClient()
	defer client.Quit()
	_, e := client.Zrem(key, v)
	return e
}

// GetRangeByRank 获取n名排行
func GetRangeByRank(key string, from, to int64) ([][]byte, error) {
	client := getClient()
	defer client.Quit()

	result, e := client.Zrange(key, from, to)
	return result, e
}

// GetRangeByScore 指定分数区间的排名
func GetRangeByScore(key string, from, to float64) ([][]byte, error) {
	client := getClient()
	defer client.Quit()

	result, e := client.Zrangebyscore(key, from, to)
	return result, e
}

// GetScore 获取分数
func GetScore(key string, v []byte) (float64, error) {
	client := getClient()
	defer client.Quit()

	result, e := client.Zscore(key, v)

	return result, e
}

// GetRank 获取排名
func GetRank(key string, v []byte) (int64, error) {
	client := getClient()
	defer client.Quit()

	result, e := client.Zrank(key, v)

	return result, e
}

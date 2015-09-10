package gredis

import (
	"fmt"
	"redis"
)

//GRedis 关系数据库结构
type GRedis struct {
	client  redis.Client
	client2 redis.Client
}

var _self *GRedis

//Instance 获取实例
func Instance() *GRedis {
	if _self == nil {
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
	spec := redis.DefaultSpec().Db(1).Password("")
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
	spec := redis.DefaultSpec().Db(1).Password("")
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
	spec := redis.DefaultSpec().Db(1).Password("")
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
	c := Instance()
	e := c.client.Set(key, value)
	if e != nil {
		connToRedis()
		e := c.client2.Set(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

//GetValue 获取一个key值
func GetValue(key string) ([]byte, error) {
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
	return result, e
}

//GetValues 模糊获取key值，以*作为模糊标记
func GetValues(key string) ([][]byte, error) {
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
	return result, nil
}

//GetKeys 获取key名
func GetKeys(v string) []string {
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
	return keyArr
}

//DelValue 删除一个key值
func DelValue(key string) error {
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
	return e
}

//DelValues 模糊删除key值，以*作为模糊标记
func DelValues(key string) error {
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
	return nil
}

//AddSet 向指定的set添加值
func AddSet(key, value string) error {
	// c := Instance()
	client := getClient()
	defer client.Quit()

	_, e := client.Sadd(key, []byte(value))
	if e != nil {
		// connToRedis()
		// _, e = client2.Sadd(key, []byte(value))
		return e
	}
	return nil
}

//AddListSet 向指定的set加入一多值
func AddListSet(key string, values [][]byte) error {
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
	return nil
}

//GetSet 获取set
func GetSet(key string) [][]byte {
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
	return result
}

//GetSetLen 获取set长度
func GetSetLen(key string) int64 {
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
	return result
}

//ChkInSet 判定一个值是否在set中
func ChkInSet(key string, v string) bool {
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
	return bIs
}

//DelSet 删除set
func DelSet(key string) error {
	e := DelValue(key)
	return e

}

//RemSetValue 从set中删除值
func RemSetValue(key string, v string) (bool, error) {
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
	return bRem, e
}

//FlushDB 刷新db
func FlushDB() {
	c := Instance()
	c.client.Flushdb()
}

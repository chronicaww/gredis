package gredis

import (
	"strconv"
	"time"

	"github.com/chronicaww/gocsv"
	"github.com/chronicaww/gredis/funcAnysis"
)

//GetIdsByIndex 根据索引获得key
func GetIdsByIndex(rc *RedisConn, csvName, indexName, indexID string) []int32 {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetIdsByIndex", tm)
	return ArrByteToArrInt32(GetSet(rc, GetIndexKey(csvName, indexName, indexID)))
}

//AddCsvIndex 增加索引
func AddCsvIndex(rc *RedisConn, csvName, indexName string) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.AddCsvIndex", tm)
	ids := GetCsvIDs(rc, csvName)
	for _, v := range ids {
		indexID := GetCsvValue(rc, csvName, indexName, v)
		key := GetIndexKey(csvName, indexName, string(indexID))
		AddSet(rc, key, strconv.Itoa(int(v)))
	}
}

//GetIndexKey 获取索引key
func GetIndexKey(csvName, indexName, indexID string) string {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetIndexKey", tm)
	key := "csv:" + csvName + ":IndexName:" + indexName + ":IndexID:" + string(indexID)
	return key
}

//ReadCsvWithKey 读取
func ReadCsvWithKey(rc *RedisConn, csvName, fileName, keyName string, vols []string) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.ReadCsvWithKey", tm)
	records, e := csv.Read(fileName)
	if e != nil {
		return e
	}
	titles, _ := csv.ReadTitle(fileName, "#!")
	n := 0
	for i, u := range titles {
		if u == keyName {
			n = i
			break
		}
	}

	for _, v := range records {
		SetCsvSetKeyValue(rc, csvName, v[n])
		for i, u := range titles {
			if i == n {

				continue
			}
			bIn := false
			if len(vols) <= 0 {
				bIn = true
			} else {
				for _, w := range vols {
					if w == u {
						bIn = true
						break
					}
				}
			}

			if bIn {
				SetCsvValue(rc, csvName, u, v[n], v[i])
			}
		}
	}
	return nil

}

//RemCsv 移除旧数据
func RemCsv(rc *RedisConn, csvName string) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.RemCsv", tm)
	keys := GetKeys(rc, csvName+":*")
	for _, v := range keys {
		DelValue(rc, v)
	}

	DelSet(rc, GetCsvSetKey(csvName))

	return nil
}

//ReadCsv 读取
func ReadCsv(rc *RedisConn, csvName, fileName string) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.ReadCsv", tm)
	RemCsv(rc, csvName)
	return ReadCsvWithKey(rc, csvName, fileName, "", []string{})
	// records, e := csv.Read(fileName)
	// if e != nil {
	// 	return e
	// }
	// titles, e := csv.ReadTitle(fileName, "#!")
	// for _, v := range records {
	// 	for i, u := range titles {
	// 		if i < 1 {
	// 			continue
	// 		}
	// 		SetCsvValue(csvName, u, common.SetInt32(v[0]), v[i])
	// 	}
	// 	SetCsvSetKeyValue(csvName, v[0])
	// }
	// return nil
}

//GetCsvValue 取得某个值
func GetCsvValue(rc *RedisConn, csvName, colName string, keyID int32) []byte {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetCsvValue", tm)
	r, _ := GetValue(rc, GetCsvKey(csvName, colName, strconv.Itoa(int(keyID))))
	return r
}

//GetCsvKey 获得key
func GetCsvKey(csvName, colName string, keyID string) string {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetCsvKey", tm)
	return "csv:" + csvName + ":" + keyID + ":" + colName
}

//SetCsvValue 设置值
func SetCsvValue(rc *RedisConn, csvName, colName string, keyID string, value string) {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.SetCsvValue", tm)
	SetValue(rc, GetCsvKey(csvName, colName, keyID), []byte(value))
}

//GetCsvValueByStringKey 取得某个值
func GetCsvValueByStringKey(rc *RedisConn, csvName, colName string, keyID string) []byte {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetCsvValueByStringKey", tm)
	r, _ := GetValue(rc, GetCsvKeyByStringKey(csvName, colName, keyID))
	return r
}

//GetCsvKeyByStringKey 获得key
func GetCsvKeyByStringKey(csvName, colName string, keyID string) string {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetCsvKeyByStringKey", tm)
	return "csv:" + csvName + ":" + keyID + ":" + colName
}

//GetCsvSetKey 获得setkey
func GetCsvSetKey(csvName string) string {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetCsvSetKey", tm)
	return "csv:" + csvName + ":" + "Index"
}

//SetCsvSetKeyValue 设置set值
func SetCsvSetKeyValue(rc *RedisConn, csvName, v string) error {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.SetCsvSetKeyValue", tm)
	e := AddSet(rc, GetCsvSetKey(csvName), v)
	return e
}

//ChkID 判定key是否存在
func ChkID(rc *RedisConn, csvName string, id int32) bool {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.ChkID", tm)
	return ChkInSet(rc, GetCsvSetKey(csvName), strconv.Itoa(int(id)))
}

//GetCsvIDs 获得所有id
func GetCsvIDs(rc *RedisConn, csvName string) []int32 {
	tm := time.Now()
	defer funcAnysis.Instance().Once("gredis.GetCsvIDs", tm)
	idsByte := GetSet(rc, GetCsvSetKey(csvName))
	ids := []int32{}
	for _, v := range idsByte {
		id := ToInt32(v)
		ids = append(ids, id)
	}
	return ids
}

package gredis

import (
	"encoding/binary"
	"strconv"
	"strings"
)

//ToString 字节转string
func ToString(b []byte) string {
	return string(b)
}

//ToFloat64 字节转float64
func ToFloat64(b []byte) float64 {
	result, e := strconv.ParseFloat(string(b), 0)
	if e != nil {
		return float64(0)
	}
	return result
}

//ToInt32 字节转int32
func ToInt32(b []byte) int32 {
	// result, e := strconv.Atoi(string(b))
	// if e != nil {
	// 	return int32(0)
	// }
	// if len(b) < 8 {
	// 	result, e := strconv.Atoi(string(b))
	// 	if e != nil {
	// 		return int32(0)
	// 	}
	// 	return int32(result)
	// }
	// if len(b) != 8 {
	return StringToInt32(ToString(b))
	// }
	// return int32(binary.BigEndian.Uint32(b))
}

//BigEndianToInt32 字节转int32
func BigEndianToInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}

//ToInt64 字节转int64
func ToInt64(b []byte) int64 {
	// result, e := strconv.ParseInt(string(b), 10, 0)
	// if e != nil {
	// 	return int64(0)
	// }
	// if len(b) != 8 {
	return StringToInt64(ToString(b))
	// }
}

//BigEndianToInt64 字节转int64
func BigEndianToInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

//BigEndianIntToByte int转byte
func BigEndianIntToByte(i int) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

//IntToByte int转byte
func IntToByte(i int) []byte {
	return []byte(strconv.Itoa(i))
}

//StringToInt32 字符串转int32
func StringToInt32(s string) int32 {
	tmp, _ := strconv.Atoi(s)
	return int32(tmp)
}

//StringToInt64 字符串转int64
func StringToInt64(s string) int64 {
	tmp, _ := strconv.Atoi(s)
	return int64(tmp)
}

//ConvEnum 转换枚举
func ConvEnum(s string) [][]byte {
	result := [][]byte{}
	arr1 := strings.Split(s, "...")
	is := 0
	ie := 0
	for _, v := range arr1 {
		arr2 := strings.Split(v, ".")
		if len(arr2) <= 0 {
			continue
		}
		if is > 0 {
			ie, _ = strconv.Atoi(arr2[0])
			for i := is + 1; i < ie; i++ {
				result = append(result, IntToByte(i))
			}
		}
		for _, u := range arr2 {
			// tmp, _ := strconv.Atoi(arr2[0])
			result = append(result, []byte(u))
		}
		is, _ = strconv.Atoi(arr2[len(arr2)-1])
	}
	return result
}

//ArrByteToArrInt32 字节数组转换为int32数组
func ArrByteToArrInt32(b [][]byte) []int32 {
	result := []int32{}
	for _, v := range b {
		tmp := ToInt32(v)
		if tmp > int32(0) {
			result = append(result, tmp)
		}
	}
	return result
}

//ArrInt32ToArrByte int32数组转换为字节数组
func ArrInt32ToArrByte(b []int32) [][]byte {
	result := [][]byte{}
	for _, v := range b {
		tmp := IntToByte(int((v)))
		result = append(result, tmp)
	}
	return result
}

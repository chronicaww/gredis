package gredis

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Base(t *testing.T) {
	fmt.Println("testbase")
	InitPool("127.0.0.1:6379", 1, 1, 100, 180)
	rc := new(RedisConn)
	FlushDB(rc)
	e1 := SetValue(rc, "TestKey", []byte("testValue"))
	if e1 != nil {
		t.Error("Error SetValue:", e1.Error())
	}
	s, e2 := GetValue(rc, "TestKey")
	if e2 != nil {
		t.Error("Error GetValue:", e2.Error())
	}
	if string(s) != "testValue" {
		t.Error("Error GetValue Wrong:", string(s))
	}
	e3 := DelValue(rc, "TestKey")
	if e3 != nil {
		t.Error("Error DelValue:", e3.Error())
	}
	s1, _ := GetValue(rc, "TestKey")
	if string(s1) != "" {
		t.Error("Error DelValue Wrong:", string(s))
	}
}

func Test_Set(t *testing.T) {
	fmt.Println("testset")
	InitPool("127.0.0.1:6379", 1, 1, 100, 180)
	rc := new(RedisConn)
	FlushDB(rc)
	e := AddSet(rc, "TestSetKey", "v1")
	if e != nil {
		t.Error("Error AddSet:", e.Error())
	}
	AddSet(rc, "TestSetKey", "v2")
	result := GetSet(rc, "TestSetKey")
	count := 0
	for _, v := range result {
		if string(v) == "v1" || string(v) == "v2" {
			count++
		}
	}
	if count != 2 {
		t.Error("Error GetSet:", result)
	}
	b, e := RemSetValue(rc, "TestSetKey", "v2")
	if e != nil {
		t.Error("Error RemSetValue:", result)
	}
	if b != true {
		t.Error("Error RemSetValue result:", b)
	}
	if GetSetLen(rc, "TestSetKey") != 1 {
		t.Error("Error GetSetLen :", GetSetLen(rc, "TestSetKey"))
	}
	e2 := AddListSet(rc, "TestSetKey", [][]byte{[]byte("v3"), []byte("v4")})
	if e != nil {
		t.Error("Error AddListSet :", e2)
	}
	if GetSetLen(rc, "TestSetKey") != 3 {
		t.Error("Error AddListSet :", ArrByteToArrString(GetSet(rc, "TestSetKey")))
	}
	if !ChkInSet(rc, "TestSetKey", "v3") {
		t.Error("Error ChkInSet :v3")
	}
	if ChkInSet(rc, "TestSetKey", "v2") {
		t.Error("Error ChkInSet :v2")
	}
	e3 := DelSet(rc, "TestSetKey")
	if e3 != nil {
		t.Error("Error DelSet:", e)
	}
	if GetSetLen(rc, "TestSetKey") > 0 {
		t.Error("Error DelSet len")
	}
}

func Test_List(t *testing.T) {
	fmt.Println("testlist")
	InitPool("127.0.0.1:6379", 1, 1, 100, 180)
	rc := new(RedisConn)
	FlushDB(rc)
	e := AddList(rc, "TestListKey", "0", 0)
	if e != nil {
		t.Error("Error AddList:", e.Error())
	}
	AddList(rc, "TestListKey", "2", -1)
	AddList(rc, "TestListKey", "1", 1)
	if (len(GetList(rc, "TestListKey"))) != 3 {
		t.Error("Error AddList len:", (GetList(rc, "TestListKey")))
	}
	for i, v := range GetList(rc, "TestListKey") {
		if strconv.Itoa(i) != string(v) {
			t.Error("Error AddList GetList:", i, string(v))
		}
	}
	e2 := RemListValue(rc, "TestListKey", "1")
	if e2 != nil {
		t.Error("Error RemListValue:", e2.Error())
	}
	if FindInList(rc, "TestListKey", "1") >= 0 {
		t.Error("Error FindInList:", FindInList(rc, "TestListKey", "1"), GetListString(rc, "TestListKey"))
	}
}

func Test_Rank(t *testing.T) {
	InitPool("127.0.0.1:6379", 1, 1, 100, 180)
	rc := new(RedisConn)
	FlushDB(rc)
	e1 := AddScore(rc, "TestScoreRank", float64(100), []byte("No1"))
	if e1 != nil {
		t.Error("Error AddScore:", e1.Error())
	}
	AddScore(rc, "TestScoreRank", float64(90), []byte("No2"))
	AddScore(rc, "TestScoreRank", float64(75), []byte("No3"))
	AddScore(rc, "TestScoreRank", float64(70), []byte("No4"))
	AddScore(rc, "TestScoreRank", float64(65), []byte("No5"))
	AddScore(rc, "TestScoreRank", float64(60), []byte("No6"))
	AddScore(rc, "TestScoreRank", float64(55), []byte("No7"))
	AddScore(rc, "TestScoreRank", float64(44), []byte("No8"))
	rank, e := GetRank(rc, "TestScoreRank", []byte("No2"), false)
	if e != nil {
		t.Error("Error GetRank:", e.Error())
	}
	// fmt.Println("rank:", rank, e)
	if rank != int64(1) {
		t.Error("Error GetRank:", rank)
	}
	s, e2 := GetScore(rc, "TestScoreRank", []byte("No4"))
	if e2 != nil {
		t.Error("Error GetScore:", e2.Error())
	}
	if int32(s) != 70 {
		t.Error("Error GetScore:", s)
	}
	tmp3, e3 := GetRangeByScore(rc, "TestScoreRank", float64(40), float64(70))
	if e3 != nil {
		t.Error("Error GetRangeByScore:", e3.Error())
	}
	if len(tmp3) != 5 {
		fmt.Println("Error GetRangeByScore len:", len(tmp3))
	}
	tmp4, e4 := GetRangeByScoreN(rc, "TestScoreRank", float64(70), float64(40), false)
	if e4 != nil {
		t.Error("Error GetRangeByScore:", e4.Error())
	}
	if len(tmp4) != 5 {
		fmt.Println("Error GetRangeByScore len:", len(tmp4))
	}
	tmp5, e5 := GetRangeByRank(rc, "TestScoreRank", 0, 10, false)
	if e5 != nil {
		t.Error("Error GetRangeByRank:", e5.Error())
	}
	fmt.Println(ArrByteToArrString(tmp5))
	if string(tmp5[0]) != "No1" {
		t.Error("Error GetRangeByRank 1:", string(tmp5[0]))
	}
}

// func Test_List(t *testing.T) {
// 	Instance(1)
// 	FlushDB()
//
// 	list1 := [][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc"), []byte("eee")}
// 	// list2 := [][]byte{[]byte("ccc"), []byte("ddd"), []byte("eee")}
// 	for _, v := range list1 {
// 		AddList("TestingList", string(v), -1)
// 	}
//
// 	if len(GetList("TestingList")) != 4 {
// 		t.Error("Error AddList1:", len(GetList("TestingList")))
// 	}
//
// 	e := AddList("TestingList", "add", 1)
// 	if e != nil {
// 		t.Error("Error AddList0:", e)
// 		return
// 	}
//
// 	pos := FindInList("TestingList", "add")
// 	if pos != 1 {
// 		t.Error("error pos:", pos)
// 		for _, v := range GetList("TestingList") {
// 			t.Error("Error AddList:", string(v))
// 		}
// 		return
// 	}
//
// 	RemListValue("TestingList", "add")
// 	pos = FindInList("TestingList", "bbb")
// 	if pos != 1 {
// 		t.Error("Error RemListValue:", pos)
// 	}
//
// 	list := GetList("TestingList")
// 	if len(list) != 4 {
// 		t.Error("Error GetList:", list)
// 	}
// }

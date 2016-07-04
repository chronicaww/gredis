package gredis

import (
	"fmt"
	"testing"
)

// func Test_convEnum(t *testing.T) {
// 	sNum := "1.2...10.11...13.14.15...20"
// 	result := ConvEnum(sNum)
// 	for _, v := range result {
// 		fmt.Printf("%v ", ToInt32(v))
// 	}
// 	fmt.Println("121")
// 	if len(result) != 20 {
// 		t.Error("Error:", result)
// 	}
// }

func Test_Rank(t *testing.T) {
	Instance(1)
	FlushDB()
	AddScore("TestScoreRank", float64(-100), []byte("No1"))
	AddScore("TestScoreRank", float64(-90), []byte("No2"))
	AddScore("TestScoreRank", float64(-75), []byte("No3"))
	AddScore("TestScoreRank", float64(-70), []byte("No4"))
	AddScore("TestScoreRank", float64(-65), []byte("No5"))
	AddScore("TestScoreRank", float64(-60), []byte("No6"))
	AddScore("TestScoreRank", float64(-55), []byte("No7"))
	AddScore("TestScoreRank", float64(-44), []byte("No8"))
	rank, e := GetRank("TestScoreRank", []byte("No7"), false)

	fmt.Println("rank:", rank, e)
	// if rank != int64(1) {
	// 	t.Error("Error GetRank:", rank)
	// }
	// list, _ := GetRangeByRank("TestScoreRank", 0, 7)
	// for _, v := range list {
	// 	fmt.Println("name", ToString(v))
	// }

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

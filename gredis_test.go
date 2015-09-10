package gredis

import (
	"fmt"
	"testing"
)

func Test_convEnum(t *testing.T) {
	sNum := "1.2...10.11...13.14.15...20"
	result := ConvEnum(sNum)
	for _, v := range result {
		fmt.Printf("%v ", ToInt32(v))
	}
	fmt.Println()
	if len(result) != 20 {
		t.Error("Error:%v", result)
	}
}

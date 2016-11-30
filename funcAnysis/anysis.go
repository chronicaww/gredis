package funcAnysis

import (
	"fmt"
	"time"
)

var _self *FuncAnysis

// FuncAnysis 函数分析
type FuncAnysis struct {
	funcs map[string]*funcRun
}

type funcRun struct {
	name  string
	count int64
	cost  time.Duration
}

// Instance 单例
func Instance() *FuncAnysis {
	if nil == _self {
		_self = new(FuncAnysis)
		_self.funcs = make(map[string]*funcRun)
	}
	return _self
}

// Once 运行一次
func (f *FuncAnysis) Once(name string, tm time.Time) {
	dur := time.Since(tm)
	fr, exists := f.funcs[name]
	if exists && fr != nil { //已存在
		fr.cost += dur
		fr.count++
	} else {
		fr := new(funcRun)
		fr.name = name
		fr.cost += dur
		fr.count = 1
		f.funcs[name] = fr
	}
}

// Print 输出运行时间
func (f *FuncAnysis) Print() {
	for _, v := range f.funcs {
		var avg time.Duration
		if v.count > 0 {
			avg = time.Duration(int64(v.cost) / v.count)
		}
		fmt.Println(v.name, ":", v.count, ",", v.cost, ":", avg)

	}
}

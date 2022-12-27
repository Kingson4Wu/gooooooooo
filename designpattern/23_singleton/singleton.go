package singleton

import "sync"

var once sync.Once

// 不可导出对象
type earth struct {
	desc string
}

func (e *earth) String() string {
	return e.desc
}

// theEarth 地球单实例
var theEarth *earth

// TheEarth 获取地球单实例
func TheEarth() *earth {
	if theEarth == nil {
		once.Do(func() {
			theEarth = &earth{
				desc: "美丽的地球，孕育了生命。",
			}
		})
	}
	return theEarth
}

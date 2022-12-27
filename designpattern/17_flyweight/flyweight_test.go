package flyweight

import (
	"fmt"
	"testing"
)

/**
（一）概念

享元是一种结构型设计模式，它允许你在消耗少量内存的情况下支持大量对象。



模式通过共享多个对象的部分状态来实现上述功能。换句话来说，享元会将不同对象的相同数据进行缓存以节省内存。
*/

func TestFlyweight(t *testing.T) {
	dispatcher := NewTaxiDispatcher("北京市出租车调度系统")
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 10, 20)
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 20, 30)
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 30, 40)
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 40, 50)

	dispatcher.AddTrace("京B.567890", "红色", "一汽大众", "首汽", 20, 40)
	dispatcher.AddTrace("京B.567890", "红色", "一汽大众", "首汽", 50, 50)

	fmt.Println(dispatcher.ShowTraces("京B.123456"))
	fmt.Println(dispatcher.ShowTraces("京B.567890"))
}

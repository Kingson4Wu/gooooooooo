package interface_

import (
	"fmt"
	"testing"
)

// 为什么 Go 不支持 []T 转换为 []interface
// https://mp.weixin.qq.com/s/lG-sswO8i6w_6sMztvyOjw

// 在 Go 中，如果 interface{} 作为函数参数的话，是可以传任意参数的，然后通过类型断言来转换。
func foo(v interface{}) {
	if v1, ok1 := v.(string); ok1 {
		fmt.Println(v1)
	} else if v2, ok2 := v.(int); ok2 {
		fmt.Println(v2)
	}
}

func TestF(t *testing.T) {
	foo(233)
	foo("666")
}

// 是否可以将 []T 转换为 []interface

func foo2([]interface{}) { /* do something */ }

func TestF2(t *testing.T) {
	var a []string = []string{"hello", "world"}
	fmt.Println(a)
	//foo2(a)
	// cannot use a (variable of type []string) as []interface{} value in argument to foo2
	//这段代码是不能编译通过的，如果想直接通过 b := []interface{}(a) 的方式来转换
	//还是会报错：
	//cannot use a (type []string) as type []interface {} in function argument
}

//正确的转换方式需要这样写：

func TestF3(t *testing.T) {
	var a []string = []string{"hello", "world"}
	b := make([]interface{}, len(a))
	for i := range a {
		b[i] = a[i]
	}
}

/**
官方解释
这个问题在官方 Wiki 中是有回答的，我复制出来放在下面：

The first is that a variable with type []interface{} is not an interface! It is a slice whose element type happens to be interface{}. But even given this, one might say that the meaning is clear. Well, is it? A variable with type []interface{} has a specific memory layout, known at compile time. Each interface{} takes up two words (one word for the type of what is contained, the other word for either the contained data or a pointer to it). As a consequence, a slice with length N and with type []interface{} is backed by a chunk of data that is N*2 words long. This is different than the chunk of data backing a slice with type []MyType and the same length. Its chunk of data will be N*sizeof(MyType) words long. The result is that you cannot quickly assign something of type []MyType to something of type []interface{}; the data behind them just look different.

大概意思就是说，主要有两方面原因：

[]interface{} 类型并不是 interface，它是一个切片，只不过碰巧它的元素是 interface；
[]interface{} 是有特殊内存布局的，跟 interface 不一样。
*/

/**

通用方法
通过以上分析，我们知道了不能转换的原因，那有没有一个通用方法呢？因为我实在是不想每次多写那几行代码。

也是有的，用反射 reflect，但是缺点也很明显，效率会差一些，不建议使用。

func InterfaceSlice(slice interface{}) []interface{} {
 s := reflect.ValueOf(slice)
 if s.Kind() != reflect.Slice {
  panic("InterfaceSlice() given a non-slice type")
 }

 // Keep the distinction between nil and empty slice input
 if s.IsNil() {
  return nil
 }

 ret := make([]interface{}, s.Len())

 for i := 0; i < s.Len(); i++ {
  ret[i] = s.Index(i).Interface()
 }

 return ret
}
还有其他方式吗？答案就是 Go 1.18 支持的泛型，这里就不过多介绍了，大家有兴趣的话可以继续研究。!!!!

*/

//go:build goexperiment.arenas

package arena_test

import (
	"arena"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"unsafe"
)

//https://zhuanlan.zhihu.com/p/604686258

/**
基础原理
尽管直观上我们认为arena要把Go变成C++了，实际上arena只是一个内存池的技术——创建一个大的连续内存块，该内存块只需要预先分配一次，然后在此内存上创建对象，使用完后统一释放内存。
如下，相比不使用arena，业务(JSON解析等)存在大量小对象，GC会消耗大量CPU和内存来实现垃圾回收，而使用arena只需要分配一次内存，所有对象都在池中管理，手动选择合适的时机释放。

*/

/**
定义环境变量： export GOEXPERIMENT=arenas
运行程序同时开启： GOEXPERIMENT=arenas go run main.go
指定Build Tag： go run main.go -tags goexperiment.arenas

*/

/**
编写相关代码，可在需要开启arena特性文件增加 //go:build goexperiment.arenas。具体使用很简单，先创建arena池，然后在此池上分配变量，使用完后统一释放arena池。

使用步骤
1.创建arena内存池，不需要的时候释放

NewArena(): 创建一个Arena, 你可以创建多个Arena, 批量创建一批对象，统一手工释放。它不是线程安全的。
Free(): 释放Arena以及它上面创建出来的所有的对象。释放的对象你不应该再使用了，否则可能会导致意想不到的错误。
2.从池中分配需要的空间
当前只支持具体对象和slice，还没有实现MakeMap、MakeChan这样在Arena上创建map和channel的方法，后续可能会加上。

NewT any *T： 创建一个对象
MakeSliceT any []T: 在Arena创建一个Slice。
3.如果希望内存池被释放后还使用，可拷贝到堆分配空间上

CloneT any: 克隆一个Arena上对象，只能是指针、slice或者字符串。如果传入的对象不是在Arena分配的，直接原对象返回，否则脱离Arena创建新的对象。
*/

/**

防止错误
和C++一样，自己管理内存，实际中最容易遇到的问题

忘记释放内存，导致OOM
引用已经释放的arena池上分配的遍历，导致程序Crash
通常用的手段是，程序实际上线前，借助一些地址/内存预检测手段，常用的就是address sanitizer (asan)/memory sanitizer (msan)。


生产建议
不要滥用，和对待unsafe, reflect, or cgo一样，只有必要时用
注意释放Free，需要释放后使用的记得Clone
实际封装，可以全局封装一个多个持有arena池的单实例对象，或者参考鸟窝大佬的做法，类似context，每个函数传递一个全局分配好的arena池

*/

/**
原理区别
同样都是为了解决频繁分配对象和大量对象GC带来的开销

sync.Pool
相同类型的对象，使用完后暂时缓存，不GC，下次再有相同的对象分配时直接用之前的缓存的对象，这样避免频繁创建大量对象。
不承诺这些缓存对象的生命周期，GC时会释放之前的缓存，适合解决频繁创建相同对象带来的压力，短时间(两次GC之间)大量创建可能还是会有较大冲击，使用相对简单，但只能用于相同结构创建，不能创建slice等复杂结构

arena
自己管理内存分配，统一手动释放，对象的生命周期完全自己控制，使用相对复杂，支持slice等复杂结构且可定制性强



*/

func processRequest(req *http.Request) {
	// 开始创建公共arena内存池
	mem := arena.NewArena()
	// 最后统一释放内存池
	defer mem.Free()

	// 分配一系列单对象
	for i := 0; i < 10; i++ {
		obj := arena.New[T](mem)
		obj.Foo = "Hello"

		fmt.Printf("%v\n", obj)
	}

	// 或者分配slice 暂时不支持map
	// 参数 mem, length, capacity
	slice := arena.MakeSlice[T](mem, 100, 200)
	slice[0].Foo = "hello"
	fmt.Printf("%v\n", slice)

	// 不能直接分配string，可借助bytes转换
	src := "source string"

	bs := arena.MakeSlice[byte](mem, len(src), len(src))
	copy(bs, src)
	str := unsafe.String(&bs[0], len(bs))

	fmt.Printf("%v\n", str)
}

type MyObj struct {
	Index int
}

func BenchmarkCreateObj(b *testing.B) {
	b.ReportAllocs()

	var p *MyObj

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			p = new(MyObj)
			p.Index = j
		}
	}
}

var (
	objPool = sync.Pool{
		New: func() interface{} {
			return &MyObj{}
		},
	}
)

func BenchmarkCreateObj_SyncPool(b *testing.B) {
	b.ReportAllocs()

	var p *MyObj

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			p = objPool.Get().(*MyObj)
			p.Index = 23
			objPool.Put(p)
		}
	}
}

func BenchmarkCreateObj_Arena(b *testing.B) {
	b.ReportAllocs()

	var p *MyObj

	a := arena.NewArena()
	defer a.Free()

	for i := 0; i < b.N; i++ {

		for j := 0; j < 1000; j++ {
			p = arena.New[MyObj](a)
			p.Index = 23
		}

	}
}

/**
相比原始的对象分配，sync.Pool和arena都降低每次操作内存分配-0 allocs/op，前者是复用对象，每次操作没内存申请 0 B/op，后者从arena中分配 8067 B/op
整体来看，相比原始的操作，syncPool每次操作耗时基本不变，但是内存分配大大减少，但是arena虽然内存分配减少，但每次操作耗时增加，可见不是每种场合arena都合适
*/

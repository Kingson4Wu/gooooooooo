+ 最全Go select底层原理，一文学透高频用法: https://mp.weixin.qq.com/s/n-2T7Bzj-Yr5Pfz6ns5gHg


为了节省各位开发者时间，本文先给出结论，若您时间不足可以先看完本节并收藏，后续再持续阅读消化：
第一，Go select语句采用的多路复用思想，本质上是为了达到通过一个协程同时处理多个IO请求（Channel读写事件）。

第二，select的基本用法是：通过多个case监听多个Channel的读写操作，任何一个case可以执行则选择该case执行，否则执行default。如果没有default，且所有的case均不能执行，则当前的goroutine阻塞。

第三，编译器会对select有不同的case的情况进行优化以提高性能。首先，编译器对select没有case、有单case和单case+default的情况进行单独处理。这些处理或者直接调用运行时函数，或者直接转成对channel的操作，或者以非阻塞的方式访问channel，多种灵活的处理方式能够提高性能，尤其是避免对channel的加锁。

第四，对最常出现的select有多case的情况，会调用 runtime.selectgo() 函数来获取执行 case 的索引，并生成 if 语句执行该case的代码。

第五，selectgo函数的执行分为四个步骤：首先，随机生成一个遍历case的轮询顺序 pollorder 并根据 channel 地址生成加锁顺序 lockorder，随机顺序能够避免channel饥饿，保证公平性，加锁顺序能够避免死锁；然后，根据 pollorder 的顺序查找 scases 是否有可以立即收发的channel，如果有则获取case索引进行处理；再次，如果pollorder顺序上没有可以直接处理的case，则将当前 goroutine 加入各 case 的 channel 对应的收发队列上并等待其他 goroutine 的唤醒；最后，当调度器唤醒当前 goroutine 时，会再次按照 lockorder 遍历所有的case，从中查找需要被处理的case索引进行读写处理，同时从所有case的发送接收队列中移除掉当前goroutine。


+ select的结构跟switch有些相似，不过仅仅只是形式上相似而已，本质上大为不同。select中的多个case的表达式必须都是Channel的读写操作，不能是其他的数据类型。select通过多个case监听多个Channel的读写操作，任何一个case可以执行则选择该case执行，否则执行default。如果没有default，且所有的case均不能执行，则当前的goroutine阻塞。

+ 对于空的 select 语句，程序会被阻塞，确切的说是当前协程被阻塞，同时 Go 自带死锁检测机制，当发现当前协程再也没有机会被唤醒时，则会发生 panic。所以上述程序会 panic。


+ 总结

综合上面的分析，总结如下：
编译器会对select有不同的case的情况进行优化以提高性能。首先，编译器对select没有case、有单case和单case+default的情况进行单独处理，这些处理或者直接调用运行时函数，或者直接转成对channel的操作，或者以非阻塞的方式访问channel，多种灵活的处理方式能够提高性能，尤其是避免对channel的加锁。

对最常出现的select有多case的情况，会调用runtime.selectgo()函数来获取执行case的索引，并生成 if 语句执行该case的代码。

selectgo函数的执行分为四个步骤：首先，随机生成一个遍历case的轮询顺序 pollorder 并根据 channel 地址生成加锁顺序 lockorder，随机顺序能够避免channel饥饿，保证公平性，加锁顺序能够避免死锁和重复加锁；然后，根据 pollorder 的顺序查找 scases 是否有可以立即收发的channel，如果有则获取case索引进行处理；再次，如果pollorder顺序上没有可以直接处理的case，则将当前 goroutine 加入各 case 的 channel 对应的收发队列上并等待其他 goroutine 的唤醒；最后，当调度器唤醒当前 goroutine 时，会再次按照 lockorder 遍历所有的case，从中查找需要被处理的case索引进行读写处理，同时从所有case的发送接收队列中移除掉当前goroutine。


+ 怎么找到go 的 select关键字对应的实现源码

ChatGPT
要查看select关键字的实现源码，你需要查看 Go 语言的标准库。select是 Go 语言中用于处理多个通道操作的关键字。你可以在 Go 的源码仓库中找到这些实现。

首先，你需要知道 select 的功能由 Go 语言的运行时（runtime）库提供，因此你需要查找 runtime 包中的相关代码。

访问 Go 的源码仓库（https://github.com/golang/go）。

在该仓库中，你可以找到 runtime 包的位置：src/runtime/select.go。

在 select.go 文件中，你会找到 select 关键字的实现代码。

在阅读源码时，你可能会发现涉及调度器、协程和通道操作等底层细节的代码，因为 select 的实现是依赖于这些机制的


+ /usr/local/go/src/cmd/compile/internal/walk/select.go
+ /usr/local/go/src/runtime/select.go
+ 核心函数：selectgo()

+ /usr/local/go/src/go/token/token.go (自己全局搜出来的)
+ cmd/compile/internal/walk/stmt.go
```go
    	case ir.OSELECT:
		n := n.(*ir.SelectStmt)
		walkSelect(n)
		return n
```
+ cmd/compile/internal/ir/fmt.go
```go
OSELECT:           "select",
```
大概率是这句 对应到相应的关键字


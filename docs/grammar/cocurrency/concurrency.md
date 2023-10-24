+ Java 和 Go 在并发实现中的区别:<https://mp.weixin.qq.com/s/FWqSDNNA4tj9Rpf9m97jaQ>

+ Java 使用 OS 线程通过 Java 运行时管理的线程来完成并行进程。Golang 通过 Goroutine 使用线程 os 完成并行进程。在 Java 和 Golang 之间的并行概念中，没有太大区别，几乎相同，只是名称不同。
+ 并发概念不同。在 Java 中，这是通过将 Java 运行时线程映射到 os线 程来完成的。同时，golang 使用此 goroutine 映射到 golang 上的调度程序，将其映射到更深层次。
+ Goroutine 本身不是线程系统或由 os 线程管理的线程。但是，这种想法更多的是将函数协程处理到 os 线程中的多路复用器。这样，当发生进程阻塞时，它将被移至未使用的线程或绿色线程，Go 调度程序的任务是将此绿色线程映射到 OS 线程，并将 goroutine 分配到空闲的绿色线程中。

乍一看，goroutine 概念与 Reactive .io 中以 Reactore 3 或 RxJava 表示的 React Java 的非阻塞概念几乎相同。但是，Java反 应流概念比 goroutines 具有更高级的方法。

+ 一文了解 Go 并发模型：<https://mp.weixin.qq.com/s?__biz=Mzg5Mjc3MjIyMA==&mid=2247544036&idx=1&sn=b34b770f7bc15ea2f8486d2d7d6829a9&source=41#wechat_redirect>

### CSP 并发模型
+ Golang 语言中实现了两种并发模式，一种是我们熟悉的线程与锁并发模型，它主要依赖于共享内存实现。线程与锁模型类似于对底层硬件运行过程的形式化，程序的正确运行很大程度依赖于开发人员的能力和技巧，程序在出错时不易排查。另一种是 Golang 中倡导使用的 CSP（communicating sequential processes）通信顺序进程模型。


---



### Java 并发模型和 Golang
尽管在并发问题中有不同的实现方法，但是模型几乎相同。

#### 异步过程
Java

创建从 Thread 类扩展的类。
实现 Runnable 的接口。
Golang

Goroutine 开始

#### 同步过程
Java

方法上的同步块。
使用 java.util.concurrent 包中的 Lock.ReentrantLock

Golang

使用通道的概念，即术语“不通过共享内存进行通信和通过通信共享内存”的行话的实现。
Sync.Mutex 锁定资源。

#### 进程间通讯
Java

使用 object.wait()，object.Notify() 或 object.NotifyAll() 方法。
在多个线程上共享块队列
使用 PipeReader 和 PipeWriter 方法

Golang

使用 channel
使用 WaitGroup


----

+ ​Golang 并发编程指南:<https://mp.weixin.qq.com/s/V0krCjWrndzz71cVOPBxdg>
### goroutine 并发模型

调度器主要结构
主要调度器结构是 M，P，G

M，内核级别线程，goroutine 基于 M 之上，代表执行者，底层线程，物理线程
P，处理器，用来执行 goroutine，因此维护了一个 goroutine 队列，里面存储了所有要执行的 goroutine，将等待执行的 G 与 M 对接，它的数目也代表了真正的并发度( 即有多少个 goroutine 可以同时进行 )；
G，goroutine 实现的核心结构，相当于轻量级线程，里面包含了 goroutine 需要的栈，程序计数器，以及所在 M 的信息
P 的数量由环境变量中的 GOMAXPROCS 决定，通常来说和核心数对应。

### goroutine 使用
```go
//demo1
go list.Sort()

//demo2
func Announce(message string, delay time.Duration) {
 go func() {
        time.Sleep(delay)
        fmt.println(message)
    }()
}

```

### channel
```go
// 创建 channel
a := make(chan int)
b := make(chan int, 10)
// 单向 channel
c := make(chan<- int)
d := make(<-chan int)

v, ok := <-a  // 检查是否成功关闭(ok = false：已关闭)

```

### channel 使用/技巧
+ 等待一个事件，也可以通过 close 一个 channel 就足够了。
```go
c := make(chan bool)
go func() {
    // close 的 channel 会读到一个零值
    close(c)
}()
<-c
```
+ 阻塞程序
 - 开源项目【是一个支持集群的 im 及实时推送服务】里面的基准测试的案例
    ```go
    var exit chan bool
    <-exit
    ``` 

+ 还有很多，看  Golang 并发编程指南.pdf



### goroutine 泄露预防与排查
一个 goroutine 启动后没有正常退出，而是直到整个服务结束才退出，这种情况下，goroutine 无法释放，内存会飙高，严重可能会导致服务不可用

goroutine 的退出其实只有以下几种方式可以做到
main 函数退出
context 通知退出
goroutine panic 退出
goroutine 正常执行完毕退出

大多数引起 goroutine 泄露的原因基本上都是如下情况
channel 阻塞，导致协程永远没有机会退出
异常的程序逻辑(比如循环没有退出条件)

杜绝:
想要杜绝这种出现泄露的情况，需要清楚的了解 channel 再 goroutine 中的使用，循环是否有正确的跳出逻辑

TODO

---


+ golang chan 最详细原理剖析，全面源码分析！看完不可能不懂的！：<https://zhuanlan.zhihu.com/p/299592156>
+ golang 的 chan 使用非常简单，这些简单的语法糖背后其实都是对应了相应的函数实现，这个翻译由编译器来完成。深入理解这些函数的实现。
+ 多图详解Go中的Channel源码:<https://mp.weixin.qq.com/s/S9zkYIE2U6Xjx9R4JwTJ_w>

channel 使用姿势

我们从宏观的 chan 使用姿势入手，总结来讲，有以下几种姿势：

chan 的创建
chan 入队
chan 出队
select 和 chan 结合
for-range 和 chan 结合

chan 创建
创建一个 channel ，一般用户使用姿势有两种，分别是创建有 buffer 和没有 buffer 的 channel 。

// no buffer 的 channel
c := make(chan int)
// 自带 buffer 的 channel 
c1 := make(chan int , 10)
这个对应了实际函数是 makechan ，位于 runtime/chan.go 文件里。

chan 入队
用户使用姿势：

c <- x
对应函数实现 chansend ，位于 runtime/chan.go 文件。

chan 出队
用户使用姿势：

v := <-c
v, ok := <-c
对应函数分别是 chanrecv1 和 chanrecv2 ，位于 runtime/chan.go 文件。

结合 select 语句
用户使用姿势：

select {
case c <- v:
 //  ... foo
default:
 //  ... bar
}
对应函数实现为 selectnbsend , 位于 runtime/chan.go 文件中。

用户使用姿势：

select {
case v = <-c:
 //  ... foo
default:
 //  ... bar
}
对应函数实现为 selectnbrecv , 位于 runtime/chan.go 文件中。

用户使用姿势：

select {
case v, ok = <-c:
 //  ... foo
default:
 //  ... bar
}
对应函数实现为 selectnbrecv2 , 位于 runtime/chan.go 文件中。

结合 for-range 语句
用户使用姿势：

for m := range c {
    // ...   do something
}
对应使用函数 chanrecv2 ，位于 runtime/chan.go 文件中。

----

+ Go语言中如何检测一个channel已经被关闭了？：<https://www.zhihu.com/question/450188866>

作者：qiya
链接：https://www.zhihu.com/question/450188866/answer/1792300859
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 错误示例
func isChanClose(ch chan int) bool {
    _, ok := <- c
}
上面是个错误示例，因为 _, ok := <-c 编译出来的是 chanrecv2 ，这个函数 block 赋值 true 传入的，所以当 c 是正常的时候，这里是阻塞的，所以这个不能用来作为一个正常的函数调用，因为会卡死协程，怎么解决这个问题？用 select  和 <-chan  来结合可以解决这个问题，select 和 <-chan 结合起来是对应 selectnbrecv  和 selectnbrecv2 这两个函数，这两个函数是非阻塞的（block = false ）。正确示例：func isChanClose(ch chan int) bool {
    select {
    case _, received := <- ch:
        return !received
    default:
    }
    return false
}
网上很多人举了一个 isChanClose  错误的例子，错误示例：func isChanClose(ch chan int) bool {
    select {
    case  <- ch:
        return true
    default:
    }
    return false
}
思考下：为什么第一个例子是对的，第二个例子是错的？因为，第一个例子编译出来对应的函数是 selectnbrecv2 ，第二个例子编译出来对应的是 selectnbrecv1 ，这两个函数的区别在于 selectnbrecv2 多了一个返回参数 received，只有这个函数才能指明是否元素出队成功，而 selected 只是判断是否要进到 select case 分支。我们通过 received 这个返回值（其实是一个入参，只不过是指针类型，函数内可修改）来反向推断 chan 是否 close 了。

小结：
case 的代码必须是 _, received := <- ch 的形式，如果仅仅是 <- ch 来判断，是错的逻辑，因为我们关注的是 received 的值；select 必须要有 default 分支，否则会阻塞函数，我们这个函数要保证一定能正常返回；chan close 原则永远不要尝试在读取端关闭 channel ，写入端无法知道 channel 是否已经关闭，往已关闭的 channel 写数据会 panic ；
一个写入端，在这个写入端可以放心关闭 channel；
多个写入端时，不要在写入端关闭 channel ，其他写入端无法知道 channel 是否已经关闭，关闭已经关闭的 channel 会发生 panic （你要想个办法保证只有一个人调用 close）；
channel 作为函数参数的时候，最好带方向；其实这些原则只有一点：一定要是安全的是否才能去 close channel 。


### 其实并不需要 isChanClose 函数 !!!
作者：qiya
链接：https://www.zhihu.com/question/450188866/answer/1792300859
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

上面实现的 isChanClose 是可以判断出 channel 是否 close，但是适用场景优先，因为可能等你 isChanClose 判断的时候返回值 false，你以为 channel 还是正常的，但是下一刻 channel 被关闭了，这个时候往里面“写”数据就又会 panic ，如下：if isChanClose( c ) {
    // 关闭的场景，exit  
    return
}
// 未关闭的场景，继续执行（可能还是会 panic）
c <- x
因为判断之后还是有时间窗，所以 isChanClose 的适用还是有限，那么是否有更好的办法？我们换一个思路，你其实并不是一定要判断 channel 是否 close，真正的目的是：安全的使用 channel，避免使用到已经关闭的 closed channel，从而导致 panic 。这个问题的本质上是保证一个事件的时序，官方推荐通过 context 来配合使用，我们可以通过一个 ctx 变量来指明 close 事件，而不是直接去判断 channel 的一个状态。举个栗子：select {
case <-ctx.Done():
    // ... exit
    return
case v, ok := <-c:
    // do something....
default:
    // do default ....
}
ctx.Done() 事件发生之后，我们就明确不去读 channel 的数据。或者select {
case <-ctx.Done():
    // ... exit
    return
default:
    // push 
    c <- x
}
ctx.Done() 事件发生之后，我们就明确不写数据到 channel ，或者不从 channel 里读数据，那么保证这个时序即可。就一定不会有问题。我们只需要确保一点：触发时序保证：一定要先触发 ctx.Done() 事件，再去做 close channel 的操作，保证这个时序的才能保证 select 判断的时候没有问题； 只有这个时序，才能保证在获悉到 Done 事件的时候，一切还是安全的；条件判断顺序：select 的 case 先判断 ctx.Done() 事件，这个很重要哦，否则很有可能先执行了 chan 的操作从而导致 panic 问题；

作者：qiya
链接：https://www.zhihu.com/question/450188866/answer/1792300859
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

### 怎么优雅关闭 chan ？
方法一：panic-recover关闭一个 channel 直接调用 close 即可，但是关闭一个已经关闭的 channel 会导致 panic，怎么办？panic-recover 配合使用即可。func SafeClose(ch chan int) (closed bool) {
    defer func() {
        if recover() != nil {
            closed = false
        }
    }()
    // 如果 ch 是一个已经关闭的，会 panic 的，然后被 recover 捕捉到；
    close(ch)
    return true
}
这并不优雅。方法二：sync.Once可以使用 sync.Once 来确保 close 只执行一次。type ChanMgr struct {
    C    chan int
    once sync.Once
}
func NewChanMgr() *ChanMgr {
    return &ChanMgr{C: make(chan int)}
}
func (cm *ChanMgr) SafeClose() {
    cm.once.Do(func() { close(cm.C) })
}
这看着还可以。方法三：事件同步来解决对于关闭 channel 这个我们有两个简要的原则：永远不要尝试在读端关闭 channel ；永远只允许一个 goroutine（比如，只用来执行关闭操作的一个 goroutine ）执行关闭操作；可以使用 sync.WaitGroup 来同步这个关闭事件，遵守以上的原则，举几个例子：第一个例子：一个 senderpackage main

import "sync"

func main() {
    // channel 初始化
    c := make(chan int, 10)
    // 用来 recevivers 同步事件的
    wg := sync.WaitGroup{}

    // sender（写端）
    go func() {
        // 入队
        c <- 1
        // ...
        // 满足某些情况，则 close channel
        close(c)
    }()

    // receivers （读端）
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // ... 处理 channel 里的数据
            for v := range c {
                _ = v
            }
        }()
    }
    // 等待所有的 receivers 完成；
    wg.Wait()
}
这里例子里面，我们在 sender 的 goroutine 关闭 channel，因为只有一个 sender，所以关闭自然是安全的。receiver 使用 WaitGroup 来同步事件，receiver 的 for 循环只有在 channel close 之后才会退出，主协程的 wg.Wait() 语句只有所有的 receivers 都完成才会返回。所以，事件的顺序是：写端入队一个整形元素关闭 channel所有的读端安全退出主协程返回一切都是安全的。第二个例子：多个 senderpackage main

import (
    "context"
    "sync"
    "time"
)

func main() {
    // channel 初始化
    c := make(chan int, 10)
    // 用来 recevivers 同步事件的
    wg := sync.WaitGroup{}
    // 上下文
    ctx, cancel := context.WithCancel(context.TODO())

    // 专门关闭的协程
    go func() {
        time.Sleep(2 * time.Second)
        cancel()
        // ... 某种条件下，关闭 channel
        close(c)
    }()

    // senders（写端）
    for i := 0; i < 10; i++ {
        go func(ctx context.Context, id int) {
            select {
            case <-ctx.Done():
                return
            case c <- id: // 入队
                // ...
            }
        }(ctx, i)
    }

    // receivers（读端）
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // ... 处理 channel 里的数据
            for v := range c {
                _ = v
            }
        }()
    }
    // 等待所有的 receivers 完成；
    wg.Wait()
}
这个例子我们看到有多个 sender 和 receiver ，这种情况我们还是要保证一点：close(ch) 操作的只能有一个人，我们单独抽出来一个 goroutine 来做这个事情，并且使用 context 来做事件同步，事件发生顺序是：10 个写端协程（sender）运行，投递元素；10 个读端协程（receiver）运行，读取元素；2 分钟超时之后，单独协程执行 close(channel) 操作；主协程返回；一切都是安全的。总结channel 并没有直接提供判断是否 close 的接口，官方推荐使用 context 和 select 语法配合使用，事件通知的方式，达到优雅判断 channel 关闭的效果；channel 关闭姿势也有讲究，永远不要尝试在读端关闭，永远保持一个关闭入口处，使用 sync.WaitGroup 和 context 实现事件同步，达到优雅关闭效果；

+ Go 最细节篇 — chan 为啥没有判断 close 的接口 ?:<https://mp.weixin.qq.com/s/8Ks_6Y3wI39llLPxakP1pg>

先说个结论：Go语言并没有「检查channel是否关闭，但又不造成其他影响」的功能。
Go并发编程的基本原则：对于任何channel，任何时候只有一个
goroutine是这个channel的owner。而只有owner才有权对这个协程做出写操作（写操作有两种，分别是将元素投递到channel，以及关闭channel）。
既然只有一个协程负责投递或者close，那么这个协程自然知道有没有做过close操作了。

----

### Go并发控制方式
+ 你知道几种Go并发控制方式？:<https://mp.weixin.qq.com/s/tloEYzrnKNrrAo1YKdeyrw>
+ 在goroutine并发行为控制中，有三种常见的方式，分别是WaitGroup、channel和Context。

#### WaitGroup

WaitGroup位于sync包下，它的使用方法如下。

```go
func main() {
  var wg sync.WaitGroup

  wg.Add(2) //添加需要完成的工作量2

  go func() {
    wg.Done() //完成工作量1
    fmt.Println("goroutine 1 完成工作！")
  }()

  go func() {
    wg.Done() //完成工作量1
    fmt.Println("goroutine 2 完成工作！")
  }()

  wg.Wait() //等待工作量2均完成
  fmt.Println("所有的goroutine均已完成工作！")
}
```

输出:
//goroutine 2 完成工作！
//goroutine 1 完成工作！
//所有的goroutine均已完成工作！

WaitGroup这种并发控制方式尤其适用于：某任务需要多 goroutine 协同工作，每个 goroutine 只能做该任务的一部分，只有全部的 goroutine 都完成，任务才算是完成。因此，WaitGroup同名字的含义一样，是一种等待的方式。

但是，在实际的业务中，有这么一种场景：当满足某个要求时，需主动的通知某一个 goroutine 结束。比如我们开启一个后台监控goroutine，当不再需要监控时，就应该通知这个监控 goroutine 结束，不然它会一直空转，造成泄漏。

#### Channel

对于上述场景，WaitGroup无能为力。那能想到的最简单的方法：定义一个全局变量，在其它地方通过修改这个变量进行通知，后台 goroutine 会不停的检查这个变量，如果发现变量发生了变化，即自行关闭，但是这个方法未免有些笨拙。这种情况，channel+select可派上用场。
```go
func main() {
  exit := make(chan bool)

  go func() {
    for {
      select {
      case <-exit:
        fmt.Println("退出监控")
        return
      default:
        fmt.Println("监控中")
        time.Sleep(2 * time.Second)
      }
    }
  }()

  time.Sleep(5 * time.Second)
  fmt.Println("通知监控退出")
  exit <- true

  //防止main goroutine过早退出
  time.Sleep(5 * time.Second)
}
```

输出：
//监控中
//监控中
//监控中
//通知监控退出
//退出监控


这种 channel+select 的组合，是比较优雅的通知goroutine 结束的方式。

但是，该方案同样存在局限性。试想，如果有多个 goroutine 都需要控制结束怎么办？如果这些 goroutine 又衍生了其它更多的goroutine 呢？当然我们可以定义很多 channel 来解决这个问题，但是 goroutine 的关系链导致这种场景的复杂性。

#### Context

以上场景常见于CS架构模型下。在Go中，常常为每个client开启单独的goroutine（A）来处理它的一系列request，并且往往单个A中也会请求其他服务（启动另一个goroutine B），B也可能会请求另外的goroutine C，C再将request发送给例如Databse的server。设想，当client断开连接，那么与之相关联的A、B、C均需要立即退出，系统才可回收A、B、C所占用的资源。退出A简单，但是，如何通知B、C也退出呢？

这个时候，Context就出场了。
```go
func A(ctx context.Context, name string)  {
  go B(ctx ,name) //A调用了B
  for {
    select {
    case <-ctx.Done():
      fmt.Println(name, "A退出")
      return
    default:
      fmt.Println(name, "A do something")
      time.Sleep(2 * time.Second)
    }
  }
}

func B(ctx context.Context, name string)  {
  for {
    select {
    case <-ctx.Done():
      fmt.Println(name, "B退出")
      return
    default:
      fmt.Println(name, "B do something")
      time.Sleep(2 * time.Second)
    }
  }
}

func main() {
  ctx, cancel := context.WithCancel(context.Background())

  go A(ctx, "【请求1】") //模拟client来了1个连接请求

  time.Sleep(3 * time.Second)
  fmt.Println("client断开连接，通知对应处理client请求的A,B退出")
  cancel() //假设满足某条件client断开了连接，那么就传播取消信号，ctx.Done()中得到取消信号

  time.Sleep(3 * time.Second)
}
```
输出：
//【请求1】 A do something
//【请求1】 B do something
//【请求1】 A do something
//【请求1】 B do something
//client断开连接，通知对应处理client请求的A,B退出
//【请求1】 B退出
//【请求1】 A退出


示例中模拟了客户端来了连接请求，相应开启Goroutine A进行处理，A同时开启了B处理，A和B都使用了 Context 进行跟踪，当我们使用 cancel 函数通知取消时，这 2个 goroutine 都会被结束。

这就是 Context 的控制能力，它就像一个控制器一样，按下开关后，所有基于这个 Context 或者衍生的子 Context 都会收到通知，这时就可以进行清理操作了，最终释放 goroutine，这就优雅的解决了 goroutine 启动后不可控的问题。

关于Context的详细用法，不在本文讨论范围之内。后续会出专门对Context包的讲解文章，敬请关注。

#### 总结

本文列举了三种Golang中并发行为控制模式。模式之间没有好坏之分，只在于不同的场景用恰当的方案。实际项目中，往往多种方式混合使用。

WaitGroup：多个goroutine的任务处理存在依赖或拼接关系。

channel+select：可以主动取消goroutine；多groutine中数据传递；channel可以代替WaitGroup的工作，但会增加代码逻辑复杂性；多channel可以满足Context的功能，同样，也会让代码逻辑变得复杂。

Context：多层级groutine之间的信号传播（包括元数据传播，取消信号传播、超时控制等）。

----


+ 深入了解 Go 语言与并发编程:<https://mp.weixin.qq.com/s/obFUsRnppgEsGkoo08nWeQ> TODO

+ Go：有了 sync 为什么还有 atomic？:<https://mp.weixin.qq.com/s/YIIQODPJmZRrrX4hvGEwXg>
    - 对于高吞吐量系统，性能变得非常重要，因此减少锁争用（即一个进程或线程试图获取另一个进程或线程持有的锁的情况）变得更加重要。执行此操作的最基本方法之一是使用读写锁 ( sync.RWMutex) 而不是标准 sync.Mutex，但是 Go 还提供了一些原子内存原语即 atomic 包
    - atomic 不是灵丹妙药，它显然不能替代互斥锁，但是当涉及到可以使用读取-复制-更新[1]模式管理的共享资源时，它非常出色。在这种技术中，我们通过引用获取当前值，当我们想要更新它时，我们不修改原始值，而是替换指针（因此没有人访问另一个线程可能访问的相同资源）。前面的示例无法使用此模式实现，因为它应该随着时间的推移扩展现有资源而不是完全替换其内容，但在许多情况下，读取-复制-更新是完美的。
    - atomic 在写入时比在读取时慢得多，但仍然比互斥锁快得多。有趣的是，我们可以看到互斥锁读取和写入之间的差异不是很明显（慢 30%）。尽管如此， atomic 仍然表现得更好（比互斥锁快 2-4 倍）。
    -  为什么 atomic 这么快？
    简而言之，原子操作很快，因为它们依赖于原子 CPU 指令而不是依赖外部锁。使用互斥锁时，每次获得锁时，goroutine 都会短暂暂停或中断，这种阻塞占使用互斥锁所花费时间的很大一部分。原子操作可以在没有任何中断的情况下执行。
    - atomic 无法解决所有问题，某些操作只能使用互斥锁来解决。
    - 总结
    竞态条件很糟糕，应该保护对共享资源的访问。互斥体很酷，但由于锁争用而趋于缓慢。对于某些读取-复制-更新模式有意义的情况（这往往是动态配置之类的东西，例如特性标志、日志级别或 map 或结构体，一次填充例如通过 JSON 解析等），尤其是当读取次数比写入次数多时。atomic 通常不应用于其他用例（例如，随时间增长的变量，如缓存），并且该特性的使用需要非常小心。

    可能最重要的方法是将锁保持在最低限度，如果你在在考虑原子等替代方案，请务必在投入生产之前对其进行广泛的测试和试验。

---

+ Golang channel 三大坑，你踩过了嘛？:<https://mp.weixin.qq.com/s/chQAircuP6_dQUBHbk6ESw> TODO!!
    - 死锁
        + 2.1 只有生产者，没有消费者，或者反过来
        + 2.2 生产者和消费者出现在同一个 goroutine 中
        + 2.3 buffered channel 已满，且出现上述情况
    - 内存泄漏
        + 生产者/消费者 所在的 goroutine 已经退出，而其对应的 消费者/生产者 所在的 goroutine 会永远阻塞住，直到进程退出
        + 生产者阻塞导致泄漏
        + 消费者阻塞导致泄漏
        + 如何预防内存泄漏？
            - 创建 goroutine 时就要想清楚它什么时候被回收    
                1. 当 goroutine 退出时，需要考虑它使用的 channel 有没有可能阻塞对应的生产者、消费者的 goroutine
                2. 尽量使用 buffered channel使用 buffered channel 能减少阻塞发生、即使疏忽了一些极端情况，也能降低 goroutine 泄漏的概率
    -  panic
        + 向已经 close 掉的 channel 继续发送数据
        + 多次 close 同一个 channe
    - 如何优雅地 close channel
        + 结论：除非必须关闭 chan，否则不要主动关闭。关闭 chan 最优雅的方式，就是不要关闭 chan~ 
        + chan 关闭的原则:
            1. Don't close a channel from the receiver side 不要在消费者端关闭 chan
            2. Don't close a channel if the channel has multiple concurrent senders  有多个并发写的生产者时也别关               


---

+ 从 bug 中学习：六大开源项目告诉你 go 并发编程的那些坑:<https://mp.weixin.qq.com/s/FDV77dO9nwtPltmx5cB7Lw>  TODO!!
    - unbuffered channel 由于 receiver 退出导致 sender 侧 block
        + 这个 bug 的修复方式也是非常的简单，把 unbuffered channel 修改成 buffered channel。
    - WaitGroup 误用导致阻塞
    - context 误用导致资源泄漏
    - 多个 goroutine 同时读写共享变量导致的 bug
    - channel 被关闭多次引发的 bug
    - timer 误用产生的 bug
    - 读写锁误用引发的 bug    


----


+ Golang 开发需要协程池吗？:<https://www.zhihu.com/question/302981392>

显然不需要，goroutine的初衷就是轻量级的线程，为的就是让你随用随起，结果你又搞个池子来，这不是脱裤子放屁么？你需要的是限制并发，而协程池是一种违背了初衷的方法。池化要解决的问题一个是频繁创建的开销，另一个是在等待时占用的资源。goroutine 和普通线程相比，创建和调度都不需要进入内核，也就是创建的开销已经解决了。同时相比系统线程，内存占用也是轻量的。所以池化技术要解决的问题goroutine 都不存在，为什么要创建 goroutine pool 呢？如果因为 goroutine 持有资源而要去创建goroutine pool，那只能说明代码的耦合度较高，应该为这类资源创建一个goroutine-safe的对象池，而不是把goroutine本身池化。

作者：Angry Bugs
链接：https://www.zhihu.com/question/302981392/answer/561075901
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


---
协程池

绝大部分应用场景，go是不需要协程池的。当然，协程池还是有一些自己的优势：
可以限制goroutine数量，避免无限制的增长。
减少栈扩容的次数。
频繁创建goroutine的场景下，资源复用，节省内存。（需要一定规模。一般场景下，效果不太明显。）

go对goroutine有一定的复用能力。所以要根据场景选择是否使用协程池，不恰当的场景不仅得不到收益，反而增加系统复杂性。

---

只要控制gorountine数量的问题，我觉得就像sync.Pool一样，Pool中的数据在每次GC的时候都会清掉，所以不能用在一些需要保持连接的场景下，但是存在即合理，sync.Pool就是个临时对象池，可以减轻程序频繁创建对象的消耗，以减轻垃圾回收的压力，gorountine是不是也是一样呢，虽然起一个协程的开销很小，但是在极限情况下，合理复用总是没错的，这事儿还是得结合具体场景来选择，简单限制数量，开发起来快速，彻底池化，可以压榨极限性能。

作者：silsuer
链接：https://www.zhihu.com/question/302981392/answer/549889122
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

----

golang协程池模型思考： <https://www.jianshu.com/p/f79424050d6a>

总结：

1、如果程序并发创建协程数据量很大，每个协程处理任务的事件较长，需要维持协程池，毕竟服务器的资源是有上限的，到达一定数量之后，会导致协程数过多而产生CPU负载较高的情况。

2、从自己压测数据看，创建协程的过程确实比较小的系统开销，通过上述压测数据来看也只有11us的差距。

作者：GoSnail
链接：https://www.jianshu.com/p/f79424050d6a
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


----

https://mp.weixin.qq.com/s/rgecHCCgBEpSC3lOQMn9Lg

我们简单讨论下协程池的问题: 协程是如此的轻量级，是否还需要协程池呢？
首先，协程的创建、销毁和上下文切换也是有代价的。如果这个代价在整个CPU的时间片中使用占比过高，仍然有创建协程池的必要。IO密集型的服务正好符合这个条件，比如高并发的Web服务器。相反，CPU密集型的服务，大部分CPU时间花在了计算逻辑上，就不需要协程池。


---

channel池化实例，goroutine利用channel，阻塞等待

+ goroutine复用easy版，尽可能减少不必要的新协程：<https://blog.csdn.net/weixin_45566022/article/details/107900263>

```go
package main

import (
	"fmt"
	"time"
)

// 定义方法接口
type Pool interface {
	Dispatch(task func() int) error
}

// 定义协程池结构体
type pool struct {
	// 待执行任务chan
	work chan func() int
	// 用于限制最大协程数
	numControl chan struct{}
	// 用于停止协程
	stopChan chan struct{}
}

// 实例化pool,设置最大协程数size（也可以根据runtime.NumCPU()来进行限制)
func NewPool(size int) Pool {
	return &pool{
		work:       make(chan func() int),
		numControl: make(chan struct{}, size),
		stopChan:   make(chan struct{}, size),
	}
}

// 任务调度，当work和numControl都处于阻塞时，阻塞直到有一方解除阻塞
func (p *pool) Dispatch(task func() int) error {
	select {
	// 第一次启动阻塞，worker运行后此chan停止阻塞，
	case p.work <- task:
	case p.numControl <- struct{}{}: // 第一次运行时执行此处，且当任务触发，同时其他协程阻塞或未结束时，开启新的协程
		go p.Worker(task)
	}
	return nil
}

func (p *pool) Worker(task func() int) {
	defer func() { <-p.numControl }()
firstFor:
	for {
		// task执行
		i := task()
		if i == 0 {
			p.stopChan <- struct{}{}
		}
		// task执行完毕阻塞等待work chan
		select {
		case task = <-p.work:
		case <-p.stopChan: // 根据条件通知结束协程
			fmt.Println("stop！！！！")
			break firstFor
		}
	}
}

func (p *pool) StopWorker() {
	p.stopChan <- struct{}{}
}

func main() {
	// 初识化pool
	p := NewPool(5)
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		// 当程序第一次运行时,并未开启协程运行Worker,由于work chan无缓冲，且无此chan的消费者，因此p.work <- task处于阻塞状态，
		//继而先执行p.numControl <- struct{}{}，开启协程运行Worker
		p.Dispatch(func() int {
			fmt.Println(i)
			return i % 7
		})
	}

}

```

---



---
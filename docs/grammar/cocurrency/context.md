+ 深入浅出golang——context：<https://zhuanlan.zhihu.com/p/381116234>

+ golang中Context的使用场景有哪些吗
场景一：RPC调用

在主goroutine上有4个RPC，RPC2/3/4是并行请求的，我们这里希望在RPC2请求失败之后，直接返回错误，并且让RPC3/4停止继续计算。这个时候，就使用的到Context。

场景二：PipeLine

场景三：超时请求

我们发送RPC请求的时候，往往希望对这个请求进行一个超时的限制。当一个RPC请求超过10s的请求，自动断开。当然我们使用CancelContext，也能实现这个功能（开启一个新的goroutine，这个goroutine拿着cancel函数，当时间到了，就调用cancel函数）。

场景四：HTTP服务器的request互相传递数据

context还提供了valueCtx的数据结构。

----

+ 走进Golang之Context的使用:<https://cloud.tencent.com/developer/article/1676355>


---

+ Go标准库：context 详解：<https://mp.weixin.qq.com/s/HzODjvg42deE4f8dTp-KZw>

一、为什么使用Context



（一）go的扛把子


要论go最津津乐道的功能莫过于go强大而简洁的并发能力。



func main(){
  go func(){
    fmt.Println("Hello World")
  }()
}


通过简单的go func(){}，go可以快速生成新的协程并运行。




（二）想象一个没有Context的世界


有并发的地方就有江湖。每个编程语言都有各自的并发编程方式，也有不同的并发控制方法，比如java通过join()来做主子线程同步。



go里面常用于协程间通信和管理的有channel和sync包。比如channel可以通知协程做特定操作（退出，阻塞等），sync可以加锁和同步。



假如我要实现一个可以同时关闭所有协程的程序，可以这样实现。



closed := make(chan struct{})

for i := 0; i < 2; i++ {
   // do something

   go func(i int) {
      select {
      case <-closed:
         fmt.Printf("%d Closed\n", i)
      }
   }(i)
}

// 发送指令关闭所有协程
close(closed)

time.Sleep(1 * time.Second)


因为go的协程不支持直接从外部退出，不像C++和Java有个线程ID可以操作。所以只能通过协程自己退出的方式。一般来说通过channel来控制是最方便的。



如果我想加点功能，比如到时间后退出，只要给channel增加关闭条件即可。



closed := make(chan struct{})

for i := 0; i < 2; i++ {
   go func(i int) {
      // do something

      select {
      case <-closed:
         fmt.Printf("%d Timeout\n", i)
      }
   }(i)
}

// 加个时间条件
ta := time.After(5 * time.Second)
select {
case <-ta:
   close(closed)
}

time.Sleep(1 * time.Second)




（三）用Context精简代码


上面的代码已经够简单了，但是还是显得有些复杂。比如每次都要在协程内部增加对channel的判断，也要在外部设置关闭条件。试想一下，如果程序要限制的是总时长，而不是单个操作的时长，这样每个操作要限制多少时间也是个难题。



图片



这个时候就轮到Context登场了。Context顾名思义是协程的上下文，主要用于跟踪协程的状态，可以做一些简单的协程控制，也能记录一些协程信息。



下面试着用Context改造下前面的例子：



// 空的父context
pctx := context.TODO()

// 子context（携带有超时信息），cancel函数（可以主动触发取消）
//ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
ctx, _ := context.WithTimeout(pctx, 5*time.Second)

for i := 0; i < 2; i++ {
   go func(i int) {
      // do something

    // 大部分工具库内置了对ctx的判断，下面的部分几乎可以省略
      select {
      case <-ctx.Done():
         fmt.Printf("%d Done\n", i)
      }
   }(i)
}

// 调用cancel会直接关闭ctx.Done()返回的管道，不用等到超时
//cancel()

time.Sleep(6 * time.Second)


通过Context可以进一步简化控制代码，且更为友好的是，大多数go库，如http、各种db driver、grpc等都内置了对ctx.Done()的判断，我们只需要将ctx传入即可。




二、Context基础用法



接下来介绍Context的基础用法，最为重要的就是3个基础能力，取消、超时、附加值。



（一）新建一个Context


ctx := context.TODO()
ctx := context.Background()


这两个方法返回的内容是一样的，都是返回一个空的context，这个context一般用来做父context。





（二）WithCancel


// 函数声明
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// 用法:返回一个子Context和主动取消函数
ctx, cancel := context.WithCancel(parentCtx)


这个函数相当重要，会根据传入的context生成一个子context和一个取消函数。当父context有相关取消操作，或者直接调用cancel函数的话，子context就会被取消。



举个日常业务中常用的例子：



// 一般操作比较耗时或者涉及远程调用等，都会在输入参数里带上一个ctx，这也是公司代码规范里提倡的
func Do(ctx context.Context, ...) {
  ctx, cancel := context.WithCancel(parentCtx)
  
  // 实现某些业务逻辑
  
  // 当遇到某种条件，比如程序出错，就取消掉子Context，这样子Context绑定的协程也可以跟着退出
  if err != nil {
    cancel()
  }
}




（三）WithTimeout


// 函数声明
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// 用法：返回一个子Context（会在一段时间后自动取消），主动取消函数
ctx := context.WithTimeout(parentCtx, 5*time.Second)


这个函数在日常工作中使用得非常多，简单来说就是给Context附加一个超时控制，当超时ctx.Done()返回的channel就能读取到值，协程可以通过这个方式来判断执行时间是否满足要求。



举个日常业务中常用的例子：



// 一般操作比较耗时或者涉及远程调用等，都会在输入参数里带上一个ctx，这也是公司代码规范里提倡的
func Do(ctx context.Context, ...) {
  ctx, cancel := context.WithTimeout(parentCtx)
  
  // 实现某些业务逻辑

  for {
    select {
     // 轮询检测是否已经超时
      case <-ctx.Done():
        return
      // 有时也会附加一些错误判断
      case <-errCh:
        cancel()
      default:
    }
  }

}


现在大部分go库都实现了超时判断逻辑，我们只需要传入ctx就好。





（四）WithDeadline


// 函数声明
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
// 用法：返回一个子Context（会在指定的时间自动取消），主动取消函数
ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(5*time.Second))


这个函数感觉用得比较少，和WithTimeout相比的话就是使用的是截止时间。





（五）WithValue


// 函数声明
func WithValue(parent Context, key, val interface{}) Context
// 用法: 传入父Context和(key, value)，相当于存一个kv
ctx := context.WithValue(parentCtx, "name", 123)
// 用法：将key对应的值取出
v := ctx.Value("name")


这个函数常用来保存一些链路追踪信息，比如API服务里会有来保存一些来源ip、请求参数等。



因为这个方法实在是太常用了，比如grpc-go里的metadata就使用这个方法将结构体存储在ctx里。



func NewOutgoingContext(ctx context.Context, md MD) context.Context {
    return context.WithValue(ctx, mdOutgoingKey{}, rawMD{md: md})
}




三、Context源码实现



（一）理解Context

Context是一个接口

虽然我们平时写代码时直接context.Context拿来就用，但实际上context.Context是一个接口，源码里是有多种不同的实现的，借此实现不同的功能。


type Context interface {
  // 返回这个ctx预期的结束时间
  Deadline() (deadline time.Time, ok bool)
  // 返回一个channel，当执行结束或者取消时被close，我们平时可以用这个来判断ctx绑定的协程是否该退出。实现里用的懒汉模式，所以一开始可能会返回nil
  Done() <-chan struct{}
  // 如果未完成，返回nil。已完成源码里目前就两种错误，已被取消或者已超时
  Err() error
  // 返回ctx绑定的key对应的value值
  Value(key interface{}) interface{}
}




Context们是一棵树


图片



context整体是一个树形结构，不同的ctx间可能是兄弟节点或者是父子节点的关系。



同时由于Context接口有多种不同的实现，所以树的节点可能也是多种不同的ctx实现。总的来说我觉得Context的特点是：



树形结构，每次调用WithCancel, WithValue, WithTimeout, WithDeadline实际是为当前节点在追加子节点。



继承性，某个节点被取消，其对应的子树也会全部被取消。



多样性，节点存在不同的实现，故每个节点会附带不同的功能。





Context的果子们


在源码里实际只有4种实现，要弄懂context的源码其实把这4种对应的实现学习一下就行，他们分别是：



emptyCtx：一个空的ctx，一般用于做根节点。



cancelCtx：核心，用来处理取消相关的操作。



timerCtx：用来处理超时相关操作。



valueCtx：附加值的实现方法。



现在先简单对这几个实现有个概念，后面会对其中核心关键的部分讲解下。





（二）Context类图


图片



从类图中可以看出，源码里有4种结构和3种接口，相对于其他go库源码来说是比较简单的。



核心的接口是Context，里面包含了最常用的判断是否处理完成的Done()方法 。其他所有结构都通过①实现方法或②组合的方式来实现该接口。



核心的结构是cancelCtx，被timerCtx包含。cancelCtx和timerCtx可以说代表了Context库最核心的取消和超时相关的实现，也最为复杂些。





（三）Context源码


因为篇幅关系，不会把每一行源码都拎出来，会挑比较重点的方法讲下。由于平时我们使用都是通过几个固定的方法入口，所以会围绕这几个方法讲下


emptyCtx

对外体现


var (
   background = new(emptyCtx)
   todo       = new(emptyCtx)
)

func Background() Context {
   return background
}

func TODO() Context {
   return todo
}


TODO()，Background()其实都是返回一个emptyCtx。


实现


type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
   return
}
func (*emptyCtx) Done() <-chan struct{} {
   return nil
}
func (*emptyCtx) Err() error {
   return nil
}
func (*emptyCtx) Value(key interface{}) interface{} {
   return nil
}
func (e *emptyCtx) String() string {
   switch e {
   case background:
      return "context.Background"
   case todo:
      return "context.TODO"
   }
   return "unknown empty Context"
}


这个结构非常简单，都是返回nil。emptyCtx主要用于新建一个独立的树。比方说，我想在协程里做些异步操作，但是又想脱离主协程的ctx控制如使用独立的超时限制，就可以使用这种方式。但是在整个go程序里只有todo和background两个大根节点，所以TODO()和Background()其实是新建第二层级的子树。



func demo(ctx context.Context){
  nctx := context.TODO()
  nctx := context.WithTimeout(nctx, 5*time.Second)
  ...
}


图片





valueCtx

对外体现


// 设置key, value值
func WithValue(parent Context, key, val interface{}) Context {
   if key == nil {
      panic("nil key")
   }
   if !reflectlite.TypeOf(key).Comparable() {
      panic("key is not comparable")
   }
   // 在当前节点下生成新的子节点
   return &valueCtx{parent, key, val}
}
// 根据key读取value
func (c *valueCtx) Value(key interface{}) interface{} {
   if c.key == key {
      return c.val
   }
   return c.Context.Value(key)
}


通过公共方法设置值，再通过valueCtx的内部方法获取值。后面再仔细讲下Value的实现方式。



实现


type valueCtx struct {
   Context
   key, val interface{}
}
// 根据key读取value
func (c *valueCtx) Value(key interface{}) interface{} {
  // 每个ctx只绑定一个key，匹配则返回。否则向上追溯到匹配为止
   if c.key == key {
      return c.val
   }
   return c.Context.Value(key)
}


从实现上可以看出，每当我们往ctx里调WithValue塞值时，都会生成一个新的子节点。调用的次数多了，生成的子树就很庞大。 






若当前节点的key和传入的key不匹配会沿着继承关系向上递归查找。递归到根就变成nil，表示当前key在该子树序列里没存。




cancelCtx


介绍完上面两种比较简单的结构后，终于要来到复杂的cancelCtx。cancelCtx和timerCtx关联性很强，基本上弄懂一个，另外一个也差不多了。



对外方法


func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
   // 新建一个cancelCtx
   c := newCancelCtx(parent)
   // 将父节点的取消函数和子节点关联，做到父节点取消，子节点也跟着取消
   propagateCancel(parent, &c)
   // 返回当前节点和主动取消函数（调用会将自身从父节点移除，并返回一个已取消错误）
   return &c, func() { c.cancel(true, Canceled) }
}


对外的方法里包含的几个方法都是重点的方法，后面主要讲下



结构


type cancelCtx struct {
   Context

   mu       sync.Mutex            // protects following fields
   done     chan struct{}         // created lazily, closed by first cancel call
   children map[canceler]struct{} // set to nil by the first cancel call
   err      error                 // set to non-nil by the first cancel call
}


done：用于判断是否完成。



cancel：存子取消节点。



err：取消时的错误，超时或主动取消。



type canceler interface {
   cancel(removeFromParent bool, err error)
   Done() <-chan struct{}
}


这个接口约定了可以取消的context，比如cancelCtx和timerCtx是可以取消的，emptyCtx和valueCtx是不可以取消的。



初始化


// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) cancelCtx {
   return cancelCtx{Context: parent}
}


初始化就是将父节点设置了一下，其他不设置。



cancelCtx的取消实现


// cancel closes c.done, cancels each of c's children, and, if
// removeFromParent is true, removes c from its parent's children.
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
  // 取消无论是通过父节点还是自身主动取消，err都不为空
   if err == nil {
      panic("context: internal error: missing cancel error")
   }
   c.mu.Lock()
   if c.err != nil {
     // c.err 不为空表示已经被取消过，比如父节点取消时子节点可能已经主动调用过取消函数
      c.mu.Unlock()
      return // already canceled
   }
   c.err = err
   if c.done == nil {
     // closedchan 是一个已经关闭的channel，要特殊处理是因为c.done是懒加载的方式。只有调用c.Done()时才会实际创建
      c.done = closedchan
   } else {
      close(c.done)
   }
   // 递归取消子节点
   for child := range c.children {
      // NOTE: acquiring the child's lock while holding parent's lock.
      child.cancel(false, err)
   }
   c.children = nil
   c.mu.Unlock()

  // 从父节点中移除当前节点
   if removeFromParent {
      removeChild(c.Context, c)
   }
}


整个过程可以总结为：



前置判断，看是否为异常情况。



关闭c.done，这样外部调用cancelCtx.Done()就会有返回结果。



递归调用子节点的cancel方法。



视情况从父节点中移除子节点。



这里child.cancel(false，err)不从父节点移除子节点是因为当前节点操作已取过锁，移除操作会再取锁造成冲突，故先全部cancel后再将children置为nil一次性移除。



propagateCancel 绑定父子节点的取消关系


// propagateCancel arranges for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
   done := parent.Done()
   if done == nil {
     // 若当前节点追溯到根没有cancelCtx或者timerCtx的话，表示当前节点的祖先没有可以取消的结构，后面的父子绑定的操作就可以不用做了，可参考下图
      return // parent is never canceled
   }

   select {
   case <-done:
     // 父节点已取消就直接取消子节点，无需移除是因为父子关系还没加到parent.children
      // parent is already canceled
      child.cancel(false, parent.Err())
      return
   default:
   }

  // 获取最近的可取消的祖先
   if p, ok := parentCancelCtx(parent); ok {
      p.mu.Lock()
      if p.err != nil {
      // 和前面一样，如果祖先节点已经取消过了，后面就没必要绑定，直接取消就好
         // parent has already been canceled
         child.cancel(false, p.err)
      } else {
        // 绑定父子关系
         if p.children == nil {
            p.children = make(map[canceler]struct{})
         }
         p.children[child] = struct{}{}
      }
      p.mu.Unlock()
   } else {
     // 当ctx是开发者自定义的并继承context.Context接口会进入这个分支，另起一个协程来监听取消动作，因为开发者自定义的习惯可能和源码中用c.done和c.err的判断方式有所不同
      atomic.AddInt32(&goroutines, +1)
      go func() {
         select {
         case <-parent.Done():
            child.cancel(false, parent.Err())
         case <-child.Done():
         }
      }()
   }
}


图片



①当祖先继承链里没有cancelCtx或timerCtx等实现时，Done()方法总是返回nil，可以作为前置判断。



②parentCancelCtx取的是可以取消的最近祖先节点。


总结


总结一下，cancelCtx的作用其实就两个：



绑定父子节点，同步取消信号，父节点取消子节点也跟着取消。



提供主动取消函数。





timerCtx


结构体


type timerCtx struct {
   cancelCtx
   timer *time.Timer // Under cancelCtx.mu.

   deadline time.Time
}


相比cancelCtx多了一个计时器和截止时间。



取消方法


func (c *timerCtx) cancel(removeFromParent bool, err error) {
   c.cancelCtx.cancel(false, err)
   if removeFromParent {
      // Remove this timerCtx from its parent cancelCtx's children.
      removeChild(c.cancelCtx.Context, c)
   }
   c.mu.Lock()
   if c.timer != nil {
      c.timer.Stop()
      c.timer = nil
   }
   c.mu.Unlock()
}


取消方法就是直接调用cancelCtx的取消外加计时器停止。


对外方法


func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
   if cur, ok := parent.Deadline(); ok && cur.Before(d) {
     // 传入的截止时间在父节点截止时间之后，则父节点取消时会同步取消当前子节点，不需要额外再设置计费器了，可以当普通的cancelCtx对待。
      // The current deadline is already sooner than the new one.
      return WithCancel(parent)
   }
   c := &timerCtx{
      cancelCtx: newCancelCtx(parent),
      deadline:  d,
   }
   propagateCancel(parent, c)
   dur := time.Until(d)
   if dur <= 0 {
     // 已超时直接取消
      c.cancel(true, DeadlineExceeded) // deadline has already passed
      return c, func() { c.cancel(false, Canceled) }
   }
   c.mu.Lock()
   defer c.mu.Unlock()
   if c.err == nil {
     // 间隔时间到后主动触发取消
      c.timer = time.AfterFunc(dur, func() {
         c.cancel(true, DeadlineExceeded)
      })
   }
   return c, func() { c.cancel(true, Canceled) }
}




四、总结



综上所述，Context的主要功能就是用于控制协程退出和附加链路信息。核心实现的结构体有4个，最复杂的是cancelCtx，最常用的是cancelCtx和valueCtx。整体呈树状结构，父子节点间同步取消信号。


----

轻松上手！手把手带你掌握从Context到go设计理念:<https://mp.weixin.qq.com/s/Qmh_aZLWbMSRh5k7M1Gk8A>


context包比较小，是阅读源码比较理想的一个入手，并且里面也涵盖了许多go设计理念可以学习。

go的Context作为go并发方式的一种，无论是在源码net/http中，开源框架例如gin中，还是内部框架trpc-go中都是一个比较重要的存在，而整个 context 的实现也就不到600行，所以也想借着这次机会来学习学习，本文基于go 1.18.4。话不多说，例：

为了使可能对context不太熟悉的同学有个熟悉，先来个example ，摘自源码：


我们利用WithCancel创建一个可取消的Context，并且遍历频道输出，当 n==5时，主动调用cancel来取消。

而在gen func中有个协程来监听ctx当监听到ctx.Done（）即被取消后就退出协程。

func main(){
gen := func(ctx context.Context) <-chan int {
    dst := make(chan int)
    n := 1
    go func() {
      for {
        select {
        case <-ctx.Done():
                                        close(dst)
          return // returning not to leak the goroutine
        case dst <- n:
          n++
        }
      }
    }()
    return dst
  }

  ctx, cancel := context.WithCancel(context.Background())
  // defer cancel() // 实际使用中应该在这里调用 cancel

  for n := range gen(ctx) {
    fmt.Println(n)
    if n == 5 {
                       cancel() // 这里为了使不熟悉 go 的更能明白在这里调用了 cancel()
      break
    }
  }
  // Output:
  // 1
  // 2
  // 3
  // 4
  // 5
}

这是最基本的使用方法。


图片

概览


对于context包先上一张图，便于大家有个初步了解（内部函数并未全列举，后续会逐一讲解）：



图片


最重要的就是右边的接口部分，可以看到有几个比较重要的接口，下面逐一来说下：


type Context interface{

        Deadline() (deadline time.Time, ok bool)

        Done() <-chan struct{}
        Err() error

        Value(key any) any

}

首先就是Context接口，这是整个context包的核心接口，就包含了四个 method，分别是：


Deadline() (deadline time.Time, ok bool) // 获取 deadline 时间，如果没有的话 ok 会返回 false

Done() <-chan struct{} // 返回的是一个 channel ，用来应用监听任务是否已经完成

Err() error // 返回取消原因 例如：Canceled\DeadlineExceeded
Value(key any) any // 根据指定的 key 获取是否存在其 value 有则返回


可以看到这个接口非常清晰简单明了，并且没有过多的Method，这也是go 设计理念，接口尽量简单、小巧，通过组合来实现丰富的功能，后面会看到如何组合的。



再来看另一个接口canceler，这是一个取消接口，其中一个非导出 method cancel，接收一个bool和一个error，bool用来决定是否将其从父Context中移除，err用来标明被取消的原因。还有个Done（）和Context接口一样，这个接口为何这么设计，后面再揭晓。


type canceler interface{
  cancel(removeFromParent bool, err error)
  Done() <-chan struct{}
}

接下来看这两个接口的实现者都有谁，首先Context直接实现者有 *valueCtx（比较简单放最后讲）和*emptyCtx



而canceler直接实现者有*cancelCtx和*timerCtx ，并且这两个同时也实现了Context接口（记住我前面说得另外两个是直接实现，这俩是嵌套接口实现松耦合，后面再说具体好处），下面逐一讲解每个实现。



图片

空的


见名知义，这是一个空实现，事实也的确如此，可以看到啥啥都没有，就是个空实现，为何要写呢？


type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
  return
}

func (*emptyCtx) Done() <-chan struct{} {
  return nil
}

func (*emptyCtx) Err() error {
  return nil
}

func (*emptyCtx) Value(key any) any {
  return nil
}

func (e *emptyCtx) String() string {
  switch e {
  case background:
    return "context.Background"
  case todo:
    return "context.TODO"
  }
  return "unknown empty Context"
}


再往下读源码会发现两个有意思的变量，底层一模一样，一个取名叫 background，一个取名叫todo，为何呢？耐心的可以看看解释，其实是为了方便大家区分使用，背景 是在入口处来传递最初始的context，而todo 则是当你不知道用啥，或者你的函数虽然接收ctontext参数，但是并没有做任何实现时，那么就使用todo即可。后续如果有具体实现再传入具体的上下文。所以上面才定义了一个空实现，就为了给这俩使用呢，这俩也是我们最常在入口处使用的。



var (
  background = new(emptyCtx)
  todo       = new(emptyCtx)
)

// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
  return background
}

// TODO returns a non-nil, empty Context. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context {
  return todo
}


下面再看看具体的定义吧。



cancelCtx与timerCtx、valueCtx

type cancelCtx struct{
  Context
  mu       sync.Mutex  // 锁住下面字段的操作
  // 存放的是 chan struct{}, 懒创建,
  //  只有第一次被 cancel 时才会关闭
  done     atomic.Value  
  // children 存放的是子 Context canceler ，并且当第一次被 cancel 时会被
  // 设为 nil
  children map[canceler]struct{} 
  //  第一次被调用 cancel 时，会被设置
  err      error                 
}

type timerCtx struct{
 cancelCtx
 timer *time.Timer // 定时器，用来监听是否超时该取消
 deadline time.Time // 终止时间
}

type valueCtx struct {
 Context
 key, val any
}

这里就看出来为何cancelCtx为非导出了，因为它通过内嵌Context接口也也是实现了Context的。并且通过这种方式实现了松耦合，可以通过 WithCancel（父Context） （ctx Context，cancel CancelFunc） 来传递任何自定义的Context实现。



而timerCtx是嵌套的cancelCtx，同样他也可以同时调用Context接口所有 method与cancelCtx所有method ，并且还可以重写部分方法。而 valueCtx和上面两个比较独立，所以直接嵌套的Context。



这里应该也看明白了为何canceler为何一个可导出Done一个不可导出 cancel，Done是重写Context的method会由上层调用，所以要可导出， cancel则是由return func（）{c.cancel（false，DeadlineExeceed） 类似的封装导出，所以不应该导出。



这是go中推崇的通过组合而非继承来编写代码。其中字段解释我已在后面注明，后面也会讲到。看懂了大的一个设计理念，下面我们就逐一击破，通过上面可以看到timerCtx其实是复用了cancelCtx能力，所以cancelCtx最为重要，下面我们就先将cancelCtx实现。



图片

取消


它非导出，是通过一个方法来直接返回Context类型的，这也是go理念之一，不暴露实现者，只暴露接口（前提是实现者中的可导出method不包含接口之外的method， 否则导出的method外面也无法调用）。



先看看外部构造函数WithCancel，



先判断parent是否为nil，如果为nil就panic，这是为了避免到处判断是否为nil。所以永远不要使用nil来作为一个Context传递。



接着将父Context封装到cancelCtx并返回，这没啥说得，虽然只有一行代码，但是多处使用，所以做了封装，并且后续如果要更改行为调用者也无需更改。很方便。


调用propagateCancel，这个函数作用就是当parent是可以被取消的时候就会对子Context也进行取消的取消或者准备取消动作。



返回Context与CancelFunc type >CancelFunc func（）就是一个 type func别名，底层封装的是c.cancel方法，为何这么做呢？这是为了给上层应用一个统一的调用，cancelCtx与timerCtx以及其他可以实现不同的cancel但是对上层是透明并且一致的行为就可。这个func应该是协程安全并且多次调用只有第一次调用才有效果。


func WithCancel(parent Context) (ctx Context, cancel CancelFunc){
     if parent == nil {
    panic("cannot create context from nil parent")
  }
  c := newCancelCtx(parent)
  propagateCancel(parent, &c)
  return&c, func() { c.cancel(true, Canceled) }
}

func newCancelCtx(parent Context) cancelCtx {
 return cancelCtx{Context: parent}
}


接下来就来到比较重要的func  propagateCancel，我们看看它做了啥，

首先是判断父context的Done（）方法返回的channel是否为nil，如果是则直接返回啥也不做了。这是因为父Context从来不会被取消的话，那就没必要进行下面动作。这也表名我们使用.与猫（上下文。Background（）） 这个函数是不会做任何动作的。


 done := parent.Done()
  if done == nil {
    return // parent is never canceled
  }

接下里就是一个select ，如果父Context已经被取消了的话，那就直接取消子Context就好了，这个也理所应当，父亲都被取消了，儿子当然也应该取消，没有存在必要了。


select {
  case <-done:
    // parent is already canceled
    child.cancel(false, parent.Err())
    return
  default:
  }


如果父 Context 没有被取消，这里就会做个判断，



看看parent是否是一个*cancelCtx，如果是的话就返回其p，再次检查 p.err是否为nil，如果不为nil就说明parent被取消，接着取消 子 Context，如果没被取消的话，就将其加入到p.children中，看到这里的 map是个canceler，可以接收任何实现取消器 的类型。这里为何要加锁呢？因为要对p.err以及p.children进行读取与写入操作，要确保协程安全所以才加的锁。



如果不是*cancelCtx类型就说明parent是个被封装的其他实现 Context 接口的类型，则会将goroutines是个int加1这是为了测试使用的，可以不管它。并且会启动个协程，监听父Context ，如果父Context被取消，则取消子Context，如果监听到子Context已经结束（可能是上层主动调用CancelFunc）则就啥也不用做了。


  if p, ok := parentCancelCtx(parent); ok {
    p.mu.Lock()
    if p.err != nil {
      // parent has already been canceled
      child.cancel(false, p.err)
    } else {
      if p.children == nil {
        p.children = make(map[canceler]struct{})
      }
      p.children[child] = struct{}{}
    }
    p.mu.Unlock()
  } else {
    atomic.AddInt32(&goroutines, +1)
    go func() {
      select {
      case <-parent.Done():
        child.cancel(false, parent.Err())
      case <-child.Done():
      }
    }()
  }


接下来看看parentCancelCtx的实现：它是为了找寻parent底下的 *cancelCtx，



它首先检查parent.Done（）如果是一个closedchan这个频道 在初始化时已经是个一个被关闭的通道或者未nil的话（emptyCtx）那就直接返回 nil，false。


func parentCancelCtx(parent Context) (*cancelCtx, bool) {
  done := parent.Done()
  if done == closedchan || done == nil {
    return nil, false
}

var closedchan = make(chan struct{})

func init() {
  close(closedchan)
}

p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
if !ok {
  return nil, false
}

接着判断是否parent是*cancelCtx类型，如果不是则返回nil，false，这里调用了parent.Value方法，并最终可能会落到value方法：



func value(c Context, key any) any {
  for {
    switch ctx := c.(type) {
    case *valueCtx:
      if key == ctx.key {
        return ctx.val
      }
      c = ctx.Context
    case *cancelCtx:
      if key == &cancelCtxKey {
        return c
      }
      c = ctx.Context
    case *timerCtx:
      if key == &cancelCtxKey {
        return &ctx.cancelCtx
      }
      c = ctx.Context
    case *emptyCtx:
      return nil
    default:
      return c.Value(key)
    }
  }
}

如果是*valueCtx，并且key==ctx.key则返回，否则会将c赋值为 ctx.Context，继续下一个循环



如果是*cancelCtx并且key==&cancelCtxKey则说明找到了，直接返回，否则c= ctx.上下文继续



如果是*timerCtx，并且key== &cancelCtxKey则会返回内部的*cancelCtx



如果是*emptyCtx 则直接返回nil，



默认即如果是用户自定义实现则调用对应的Value找寻



可以发现如果嵌套实现过多的话这个方法其实是一个递归调用。



如果是则要继续判断p.done与parent.Done（）是否相等，如果没有则说明：*cancelCtx已经被包装在一个自定义实现中，提供了一个不同的包装，在这种情况下就返回nil，false：


pdone, _ := p.done.Load().(chan struct{})
if pdone != done {
  return nil, false
}
return p, true


构造算是结束了，接下来看看如何取消的：



检查err是否为nil


   if err == nil {
    panic("context: internal error: missing cancel error")
  }

由于要对err、cancelCtx.done以及children进行操作，所以要加锁



如果c.err不为nil则说明已经取消过了，直接返回。否则将c.err=err赋值，这里看到只有第一次调用才会赋值，多次调用由于已经有 ！= nil+锁的检查，所以会直接返回，不会重复赋值

c.mu.Lock()
  if c.err != nil {
    c.mu.Unlock()
    return // already canceled
  }
       c.err = err


会尝试从c.done获取，如果为nil，则保存一个closedchan，否则就关闭d，这样当你context.Done（）方法返回的channel才会返回。


d, _ := c.done.Load().(chan struct{})
  if d == nil {
    c.done.Store(closedchan)
  } else {
    close(d)
  }

循环遍历c.children去关闭子Context，可以看到释放子context时会获取 子Context的锁，同时也会获取父Context的锁。所以才是线程安全的。结束后释放锁


     for child := range c.children {
    // NOTE: acquiring the child's lock while holding parent's lock.
    child.cancel(false, err)
  }
  c.children = nil
  c.mu.Unlock()


如果要将其从父Context删除为true，则将其从父上下文删除


if removeFromParent {
    removeChild(c.Context, c)
  }


removeChild也比较简单，当为*cancelCtx就将其从Children内删除，为了保证线程安全也是加锁的。


func removeChild(parent Context, child canceler) {
  p, ok := parentCancelCtx(parent)
  if !ok {
    return
  }
  p.mu.Lock()
  if p.children != nil {
    delete(p.children, child)
  }
  p.mu.Unlock()
}

Done就是返回一个channel用于告知应用程序任务已经终止：这一步是只读没有加锁，如果没有读取到则尝试加锁，再读一次，还没读到则创建一个chan，可以看到这是一个懒创建的过程。所以当用户主动调用CancelFunc时，其实根本就是将c.done内存储的chan close掉，这其中可能牵扯到父关闭，也要循环关闭子Context过程。


func (c *cancelCtx) Done() <-chan struct{} {
  d := c.done.Load()
  if d != nil {
    return d.(chan struct{})
  }
  c.mu.Lock()
  defer c.mu.Unlock()
  d = c.done.Load()
  if d == nil {
    d = make(chan struct{})
    c.done.Store(d)
  }
  return d.(chan struct{})
}

cancelCtx主要内容就这么多，接下里就是timerCtx了



图片

计时器


回顾下timerCtx定义，就是内嵌了一个cancelCtx另外多了两个字段timer和deadline，这也是组合的体现。


type timerCtx struct {
  cancelCtx
  timer *time.Timer // Under cancelCtx.mu.

  deadline time.Time
}

下面就看看两个构造函数，WithDeadline与WithTimeout，WithTimeout就是对WithDealine的一层简单封装。



检查不多说了， 第二个检查如果父context的截止时间比传递进来的早的话，这个时间就无用了，那么就退化成cancelCtx了。


func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
  if parent == nil {
    panic("cannot create context from nil parent")
  }
  if cur, ok := parent.Deadline(); ok && cur.Before(d) {
    return WithCancel(parent)
  }

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
  return WithDeadline(parent, time.Now().Add(timeout))
}

构造timerCtx并调用propagateCancel，这个已经在上面介绍过了。


 c := &timerCtx{
    cancelCtx: newCancelCtx(parent),
    deadline:  d,
  }
  propagateCancel(parent, c)


接着会看，会先利用time.直到（d.分时。Now（）） 来判断传入的 deadlineTime与当前时间差值，如果在当前时间之前的话说明已经该取消了，所以会直接调用cancel函数进行取消，并且将其从父Context中删除。否则就创建一个定时器，当时间到达会调用取消函数，这里是定时调用，也可能用户主动调用。


dur := time.Until(d)
  if dur <= 0 {
    c.cancel(true, DeadlineExceeded) 
    return c, func() { c.cancel(false, Canceled) }
  }
  c.mu.Lock()
  defer c.mu.Unlock()
  if c.err == nil {
    c.timer = time.AfterFunc(dur, func() {
      c.cancel(true, DeadlineExceeded)
    })
  }
  return c, func() { c.cancel(true, Canceled) }


下面看看cancel实现吧，相比较cancelCtx就比较简单了，先取消 cancelCtx，也要加锁，将c.timer停止并赋值nil，这里也是第一次调用才会赋值nil，因为外层还有个c.timer ！=nil的判断，所以多次调用只有一次赋值。


func (c *timerCtx) cancel(removeFromParent bool, err error) {
  c.cancelCtx.cancel(false, err)
  if removeFromParent {
    // Remove this timerCtx from its parent cancelCtx's children.
    removeChild(c.cancelCtx.Context, c)
  }
  c.mu.Lock()
  if c.timer != nil {
    c.timer.Stop()
    c.timer = nil
  }
  c.mu.Unlock()
}


相比较于cancelCtx还覆盖实现了一个Deadline（），就是返回当前 Context的终止时间。


func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
  return c.deadline, true
}


下面就到了最后一个内置的valueCtx了。



图片

值


结构器就更加加单，就多了key，val


type valueCtx struct {
 Context
 key, val any
}


也就有个Value method不同，可以看到底层使用的就是我们上面介绍的value函数，重复复用


func (c *valueCtx) Value(key any) any {
  if c.key == key {
    return c.val
  }
  return value(c.Context, key)
}


几个主要的讲解完了，可以看到不到600行代码，就实现了这么多功能，其中蕴含了组合、封装、结构体嵌套接口等许多理念，值得好好琢磨。下面我们再看看其中有些有意思的地方。我们一般打印字符串都是使用 fmt 包，那么不使用fmt包该如何打印呢？context包里就有相应实现，也很简单，就是 switch case来判断v类型并返回，它这么做的原因也有说：



“因为我们不希望上下文依赖于unicode表”，这句话我还没理解，有知道的小伙伴可以在底下评论，或者等我有时间看看fmt包实现。


func stringify(v any) string {
  switch s := v.(type) {
  case stringer:
    return s.String()
  case string:
    return s
  }
  return "<not Stringer>"
}

func (c *valueCtx) String() string {
  return contextName(c.Context) + ".WithValue(type " +
    reflectlite.TypeOf(c.key).String() +
    ", val " + stringify(c.val) + ")"
}



使用Context的几个原则


直接在函数参数传递，不要在struct传递，要明确传递，并且作为第一个参数，因为这样可以由调用方来传递不同的上下文在不同的方法上，如果你在 struct内使用context则一个实例是公用一个context也就导致了协程不安全，这也是为何net包Request要拷贝一个新的Request WithRequest（context go 1.7 才被引入），net包牵扯过多，要做到兼容才嵌入到 struct内。



不要使用nil而当你不知道使用什么时则使用TODO，如果你用了nil则会 panic。避免到处判断是否为nil。



WithValue不应该传递业务信息，只应该传递类似request-id之类的请求信息。



无论用哪个类型的Context，在构建后，一定要加上：defer cancel（），因为这个函数是可以多次调用的，但是如果没有调用则可能导致Context没有被取消继而其关联的上下文资源也得不到释放。



在使用WithValue时，包应该将键定义为未导出的类型以避免发生碰撞，这里贴个官网的例子：



// package user 这里为了演示直接在 main 包定义
// User 是存储在 Context 值
type User struct {
  Name string
  Age  int
}

// key 是非导出的，可以防止碰撞
type key int

// userKey 是存储 User 类型的键值，也是非导出的。
var userKey key

// NewContext 创建一个新的 Context，携带 *User
func NewContext(ctx context.Context, u *User) context.Context {
  return context.WithValue(ctx, userKey, u)
}

// FromContext 返回存储在 ctx 中的 *User
func FromContext(ctx context.Context) (*User, bool) {
  u, ok := ctx.Value(userKey).(*User)
  return u, ok
}

那怎么能够防止碰撞呢？可以做个示例：看最后输出，我们在第一行就用 userKey的值0，存储了一个值“a”。



然后再利用NewContext存储了&User，底层实际用的是 context.WithValue（ctx，userKey，u）



读取时用的是FromContext，两次存储即使底层的key值都为0， 但是互不影响，这是为什么呢？



还记得WithValue怎么实现的么？你每调用一次都会包一层，并且一层一层解析，而且它会比较c.key==key，这里记住go的==比较是比较值和类型的，二者都相等才为true，而我们使用type key int所以userKey与0底层值虽然一样，但是类型已经不一样了（这里就是main.userKey与0），所以外部无论定义何种类型都无法影响包内的类型。这也是容易令人迷惑的地方



package main

import (
  "context"
  "fmt"
)

func main() {
  ctx := context.WithValue(context.Background(), , "a")
  ctx = NewContext(ctx, &User{})
  v, _ := FromContext(ctx)
  fmt.Println(ctx.Value(0), v) // Output: a, &{ 0}
}


----



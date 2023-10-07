+ https://100go.co/
+ https://github.com/teivah/100-go-mistakes


+ 100 Go Mistakes and How to Avoid Them-速览(一): <https://zhuanlan.zhihu.com/p/611451902>
+ 100 Go Mistakes and How to Avoid Them-速览(二): <https://zhuanlan.zhihu.com/p/609548795>

+ 深度阅读之《100 Go Mistakes and How to Avoid Them》:<https://qcrao.com/post/100-go-mistakes-reading-notes/>

---


2.3 #3 Misusing init functions（错误使用和理解 init 函数）

2.7 #7 Returning interfaces (返回接口)
这条规则要配合上一条一起食用，如果在实现端，创建一个返回接口的函数，会因为引用循环的原因导致无法实现。

本文在坚持2.6观点的前提下，提出了2点最佳实践：

返回结构体，而不是接口
函数或者结构体尽量接受接口
当然这2点提出也是引经据典引申出来的。

当然，规则不是100%准确的，比如 Error 类型经常被函数返回，标准库 io 包里也经常有返回接口的函数。

func LimitReader(r Reader, n int64) Reader {
    return &LimitedReader{r, n}
}
如果我们知道(不是预见)一个抽象将对调用方有帮助，我们可以考虑返回一个接口


2.11 #11: Not using the functional options pattern（编写构造函数的时候，使用option模式）

解决方案一：Config struct

第一种解决方案是，使用一个Config对象存储配置。

这里遇到第一个点，配置字段为指针类型，判断用户是否配置了该值

解决方案二：Builder pattern
解决方案三：Functional options pattern

2.12 #12: Project misorganization（推荐的项目组织结构）
project-layout/README_zh-CN.md at master · golang-standards/project-layout
​github.com/golang-standards/project-layout/blob/master/README_zh-CN.md


2.13 #13: Creating utility packages（不推荐使用utility包）
util、common、shared这类包名会抹去很多信息，不如直接使用有意义的包名来替代一些工具包

2.16 #16: Not using linters（不使用linters）
https://github.com/golangci/golangci-lint
​github.com/golangci/golangci-lint


3.1 #17: Creating confusion with octal literals（混淆八进制）
在GO中以0或者0o开头的数字被认为是8进制数

file, err := os.OpenFile("foo", os.O_RDONLY, 0o644)
二进制-使用0b或者0B作为前缀（0b100代表4）
16进制-使用0x或者0X前缀（0xF代表15）
虚数-使用i作为后缀（3i）
除此之外我们可以在数字中穿插下划线提升可读性：1_000_000_000


3.2 #18 Neglecting integer overflows（忽略整数溢出）
GO有专门处理大数的包：math/big

3.3 #19: Not understanding floating points（对浮点数理解欠缺）

3.4 #20: Not understanding slice length and capacity（不理解切片len和cap）

3.5 #21: Inefficient slice initialization（低效的切片初始化）!!!
func convertGivenCapacity(foos []Foo) []Bar {
    n := len(foos)
  // 预先指定 cap
    bars := make([]Bar, 0, n)

    for _, foo := range foos {
        bars = append(bars, fooToBar(foo))
    }
    return bars
}

func convertGivenLength(foos []Foo) []Bar {
    n := len(foos)
    // 预先指定len，通过下标赋值，比使用append内置函数效率更高一些
    bars := make([]Bar, n)

    for i, foo := range foos {
        bars[i] = fooToBar(foo)
    }
    return bars
}

3.6 #22: Being confused about nil vs. empty slices（区分nil和empty切片）
nil 切片和 empty 切片的区别是 nil切片不占用空间，可以依次来判断是否使用nil 切片, 在json解析场景下，nil和empty的切片也不同

3.8 #24: Not making slice copies correctly（没有正确的复制切片）
copy函数会复制元素的个数为2个切片长度最小值

func bad() {
    src := []int{0, 1, 2}
    var dst []int
    copy(dst, src)
    fmt.Println(dst)

    _ = src
    _ = dst
}

func correct() {
    src := []int{0, 1, 2}
    dst := make([]int, len(src))
    copy(dst, src)
    fmt.Println(dst)

    _ = src
    _ = dst
}

3.9 #25: Unexpected side effects using slice append（使用切片append时产生的副作用-s[low:high:max]用法）
s2和s1共享底层数组，当s2的cap大于len，对s2 append会影响s1的结果

// 当调用f函数的时候，对s入参做的修改都会反应到上层函数中的s
func listing1() {
    s := []int{1, 2, 3}

    f(s[:2])
    fmt.Println(s)
}

// 解决方案一：对s进行深拷贝
func listing2() {
    s := []int{1, 2, 3}
    sCopy := make([]int, 2)
    copy(sCopy, s)

    f(sCopy)
    result := append(sCopy, s[2])
    fmt.Println(result)
}

// 解决方案二：使用 full slice expression: s[low:high:max]，这样就限制了切片的cap
// cap = max-low = 2-0 = 2
// 这样f函数对切片进行append操作就不用影响其他共享底层存储的切片了
func listing3() {
    s := []int{1, 2, 3}
    f(s[:2:2])
    fmt.Println(s)
}

func f(s []int) {
    _ = append(s, 10)
}

3.10 #26: Slices and memory leaks（切片的内存泄露）！！！

func main() {
    foos := make([]Foo, 1_000)
    printAlloc()

    for i := 0; i < len(foos); i++ {
        foos[i] = Foo{
            v: make([]byte, 1024*1024),
        }
    }
    printAlloc()

  // 和上面类似不会回收剩下998个元素
    two := keepFirstTwoElementsOnly(foos)
    runtime.GC()
    printAlloc()
    runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
    return foos[:2]
}

// keepFirstTwoElementsOnlyCopy 和 keepFirstTwoElementsOnlyMarkNil
// 都能保证剩下的资源被回收，但是如果想留下的元素不是2，我们想保留500个元素，
// 使用 keepFirstTwoElementsOnlyMarkNil 会有更高的效率，不用额外拷贝资源
func keepFirstTwoElementsOnlyCopy(foos []Foo) []Foo {
    res := make([]Foo, 2)
    copy(res, foos)
    return res
}

// keepFirstTwoElementsOnlyMarkNil 给元素中的指针成员置为nil，能够让GC自动回收
func keepFirstTwoElementsOnlyMarkNil(foos []Foo) []Foo {
    for i := 2; i < len(foos); i++ {
        foos[i].v = nil
    }
    return foos[:2]
}

func printAlloc() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("%d KB\n", m.Alloc/1024)
}

3.11 #27: Inefficient map initialization（低效的map初始化）

就像切片一样，如果我们预先知道map将包含的元素数量，我们应该通过提供初始大小来创建它。这样做可以避免潜在的map增长，这是相当繁重的计算，因为它需要重新分配足够的空间并重新平衡所有元素。


3.12 #28: Maps and memory leaks（map类型的内存泄露，map中的值建议使用指针类型）！！！

GO的缩容操作就是不增长内存，并没有实际上的缩容，所以创建一个较大的map之后，需要自己手动处理缩容的问题

GO的map没有真正意义上的缩容，删除了所有元素之后，降低的内存应该是bmap中overflow指向的bucket，map[int][128]byte中的数组[128]byte空值就占用128byte。

解决方案一是不断的拷贝新的map，但是在下个gc周期前，内存会翻倍；
解决方案二修改map中value的类型为指针，减少value空值占用的空间；
对于方案二,可以看到扩缩容场景的内存占用更小：

3.13 #29: Comparing values incorrectly（错误的对比值）
对与一些不可比较的类型：slice不能直接使用 == 来比较


使用 reflect.DeepEqual 方法可以根据类型不同进行不同的比较，但是运行时间是 == 的100倍，因此在对性能有需要的场景，建议自行编写对比函数。


4.2 #31: Ignoring how arguments are evaluated in range loops（忽略rang的对象是何时被计算）
slice

// range loop 的对象是一个表达式，这个表达式可以是 string, slice等，当执行
// 循环的时候，exp只会被计算一次，对原始迭代值进行拷贝生成一个副本
for i, v := range exp

// 迭代对象只计算一次，迭代3次就结束，不会无限循环
s := []int{0, 1, 2}
for range s {
    s = append(s, 10)
}

// 但是下面这种写法就会无限循环
s := []int{0, 1, 2}
for i := 0; i < len(s); i++ {
    s = append(s, 10)
}

4.5 #34: Ignoring how the break statement works（不明确break语句是如何工作-break label的使用）
在switch和select语句中使用break只会跳出switch和select的作用域，如果这2个语句被for循环包裹是不会跳出for循环的,可以在break的时候指定label标签来跳出循环


4.6 #35: Using defer inside a loop（在loop循环中使用defer语句） ！！！
在每个for循环的迭代作用域中调用defer是不会在迭代结束的时候执行，只有最内层函数调用返回的时候才会执行defer，如果每次迭代都打开了一个文件描述符，想要及时关闭：

每次迭代之后就主动关闭，不借助于defer的能力
放在一个函数闭包内，当闭包执行完之后执行defer !!!!

func readFiles(ch <-chan string) error {
    for path := range ch {

        err := func() error {
            file, err := os.Open(path)
            if err != nil {
                return err
            }

            defer file.Close()

            // Do something with file
            return nil
        }() // 放在闭包内并及时调用
        if err != nil {
            return err
        }
    }
    return nil
}


5.1 #36: Not understanding the concept of a rune（不理解rune的概念）
Go源代码使用UTF-8进行编码。因此，所有字符串都是UTF-8字符串。但是，由于字符串可以包含任意字节，如果它来自其他地方（而不是源代码），则不能保证基于UTF-8编码
rune对应着Unicode中的码点的概念，一个rune代表Unicode中的一个字符
使用UTF-8编码，Unicode码点可以编码为1到4个字节
len函数返回的是字节长度，而不是rune的长度

5.3 #38: Misusing trim functions（错误使用 trim 函数）
func main() {
    // output: 123
    // TrimRight removes all the trailing runes contained in a given set
    fmt.Println(strings.TrimRight("123oxo", "xo"))

    // output: 123o
    // TrimSuffix returns a string without a provided trailing suffix
    fmt.Println(strings.TrimSuffix("123oxo", "xo"))

    // output: 123
    // TrimLeft like TrimRight
    fmt.Println(strings.TrimLeft("oxo123", "ox"))

    // output: o123
    // TrimPrefix like TrimSuffix
    fmt.Println(strings.TrimPrefix("oxo123", "ox"))

    // output: 123
    fmt.Println(strings.Trim("oxo123oxo", "ox"))
}

5.4 #39: Under-optimized string concatenation（如何高效的字符串拼接-strings.Builder）
！！！

5.5 #40: Useless string conversions（无用的字符串转换）
bytes 包支持很多和strings包类似的方法，很多时候没有必要一定要把[]byte转化为string：比如bytes.TrimSpace 方法

5.6 #41: Substrings and memory leaks（子字符串和内存泄露-Consistent with #26）
和 #26 条类似，子字符串引用了一个较长父字符串的部分内容，导致父字符串不能被GC，导致内存暴增

6.1 #42: Not knowing which type of receiver to use（不知道receiver应该使用什么样的类型）
！！！
提供几个场景供我们决定选择receiver的类型：

一定要用指针类型的场景

方法需要修改receiver的值
如果receiver的成员不能被拷贝
应该使用指针类型的场景

如果receiver是一个大的对象，使用指针类型可以更高效的处理程序，这里大的具体数值不好界定，需要看实际的场景，这里可以使用benchmark来估计。
一定要使用值类型的场景

如果强制要求receiver是不可变的
如果receiver是一个map、function、channel类型，否则会有编译错误
应该使用值类型的场景

receiver是一个slice，并且一定需要修改
receiver是一个array或者是一个没有可变字段的struct，如time.Time
receiver是基本类型，例如int、float64或string。

6.4 ⚠️ #45: Returning a nil receiver（返回了一个nil的值，该值的指针类型实现了某一接口）
！！！

这类属于常见的错误，和for循环中的临时变量的错误差不多，当返回值指针实现了一个接口，即使这个指针为nil，函数返回的时候转化为接口也会被当成非nil的值。

6.5 #46: Using a filename as a function input（没有考虑好函数参数抽象）
这一节主要是提醒我们定义函数的时候，要考虑可扩展性，使用一个统一的接口类型作为参数不仅能方便代码测试，也可以提升代码的抽象能力，把所有的数据源(file, http, string) 使用 io.Reader 类型代替，会是一个更好的解决方案。

6.6 ⚠️ #47: Ignoring how defer arguments and receivers are evaluated（认清defer是如何以及何时计算函数参数）
defer调用函数时候常犯的错误
defer调用函数的时候函数参数使用指针类型
defer调用函数的时候，把函数放在闭包内，利用闭包引用环境变量的能力。


7.2 #49: Ignoring when to wrap an error（不明确何时包装一个error，fmt.Errorf用法）
!!!

在如下场景你可以选择包装一下error：

你需要在错误上增加一些上下文
你需要把错误封装在另一个错误内

GO支持如下几种包装error的方式：

自定义Error结构体
type BarError struct {
	Err error
}

func (b BarError) Error() string {
	return "bar failed:" + b.Err.Error()
}

func test() {
	err := bar()
	if err != nil {
		return BarError{Err: err}
	}
}
 
使用 fmt.Errorf 方法，⚠️ %w 和 %v 有不同的执行效果
if err != nil {
  // 使用 %w 指令，会返回一个包装了err的错误，接收方可以从父error中获取到源error
  // 并根据判断错误是否为某一个错误类型
  // sourceErr --wrap--> WrapErr(sourceErr)
    return fmt.Errorf("bar failed: %w", err)
}

if err != nil {
    // 使用 %v 指令，不会包装错误，会直接转化为另一个错误，源错误不再可用
    // sourceErr --transform--> otherErr
    return fmt.Errorf("bar failed: %v", err)
}


如果保留源错误的信息，会有潜在的耦合信息，因为接收方需要直接函数实现的细节，需要知道被包装的错误是什么类型，所以没有特殊需求在使用fmt.Errorf的时候使用%v指令。


7.3 #50: Checking an error type inaccurately（如何正确的判断Error类型-error.As用法）

当使用%w指令或者结构体封装的方式包装一个错误的时候，可以使用 errors.As 递归判断错误类型，errors.As函数需要传递一个目标错误类型的指针。

7.4 #51: Checking an error value inaccurately（如何正确的处理Error的值-error.Is用法）
sentinel error 是指定义为全局变量的错误类型。命名规约是以Err开头加上类型。

import "errors"
var ErrFoo = errors.New("foo")
这种类型的错误，有时候是程序允许存在的错误，比如查询数据库返回没有查询到结果的sql.ErrNoRows和io.Reader读取返回io.EOF.他们传递的信息客户端可以接受，并认为是正常的情况。

在查询数据的场景，我们想判断错误的值是否等于sql.ErrNoRows这种sentinel error，使用 == 可能会出现问题，因为ErrNoRows可能已经被包装过。对于这种问题，可以直接使用 errors.Is 来判断，他可以帮你递归判断出正确的错误。

7.7 #54: Not handling defer errors（没有处理defer语句中的error） !!!
defer语句经常会做一些收尾工作，关闭socket，释放锁等..,有些语句会返回错误，如果不处理这些错误就会导致错误信息丢失，造成资源泄露等问题。本节提出了一个通过给错误返回值命名，在defer的时候把错误值赋值给返回结果，向上传递错误。


// 一种解决方式是使用命名结果，把错误信息通过命名结果返回
func getBalance(db *sql.DB, clientID string) (balance float32, err error) {
    rows, err := db.Query(query, clientID)
    if err != nil {
        return 0, err
    }
    defer func() {
        closeErr := rows.Close()
        if err != nil {
            if closeErr != nil {
        // 如果db.Query语句也出现了执行错误，把close错误的信息打印出来
                log.Printf("failed to close rows: %v", err)
            }
            return
        }
        // 当只要close的时候有错误，把closeErr赋值给err传递到上一层
        err = closeErr
    }()

    // Use rows
    return 0, nil
}

8.3 #57: Being puzzled about when to use channels or mutexes（不清楚什么时候用锁什么时候用通道） ！！！！
锁用来保护临界区，通道就用来消息传递和事件通知


8.4 #58: Not understanding race problems（不理解race问题）
当两个或多个goroutine同时访问同一个内存位置并且至少一个正在写入时，就会发生数据竞争。

下面的例子就会出现数据竞争的问题

func listing1() {
    i := 0

    go func() {
        i++
    }()

    go func() {
        i++
    }()
}
解决方案一：调用 atomic.AddInt64 做加法

func listing2() {
    var i int64

    go func() {
        atomic.AddInt64(&i, 1)
    }()

    go func() {
        atomic.AddInt64(&i, 1)
    }()
}
解决方案二：用mutex来保护临界区

func listing3() {
    i := 0
    mutex := sync.Mutex{}

    go func() {
        mutex.Lock()
        i++
        mutex.Unlock()
    }()

    go func() {
        mutex.Lock()
        i++
        mutex.Unlock()
    }()
}
解决方案三：用channel传递数据

func listing4() {
    i := 0
    ch := make(chan int)

    go func() {
        ch <- 1
    }()

    go func() {
        ch <- 1
    }()

    i += <-ch
    i += <-ch
}


9.1 #61: Propagating an inappropriate context（错误传递context）
！！！！ 没怎么看懂

为了能够异步执行写入消息队列的动作，又能继承来自request context中的值，我们可以编写自定义的context，只继承父context的值。

9.3 #63: Not being careful with goroutines and loop variables（在range loop中执行协程，注意不要引用range生成的变量）
参考 #30

9.4 #64: Expecting deterministic behavior using select and channels（错误的以为 select 的执行结果是确定的）
select 中的 case 执行时机是随机的

解决方案可以是使用for-select嵌套：


9.5 #65: Not using notification channels（chan struct{} 可以用于通知事件） 
！！！
9.6 #66: Not using nil channels（利用nil的channel读写都会block的特性）
！！！

GO对于close的channel仍然能够读取到0值，所以需要使用 v, open := <-ch 语法中的open来判断channel是否被close

解决方案：通过 v, open := <-ch 语法判断 chan 是否被close，避免读取0值，close之后的channel置为nil，避免for循环空转

9.7 #67: Being puzzled about channel size（不清楚channel的大小如何设置）
无缓存的channel可以用于同步场景
带有缓存的channel经常用于传递消息，控制协程数量等等

9.9 #69: Creating data races with append（使用append函数的时候出现data race）
注意append操作不是线程安全的

9.12 #72: Forgetting about sync.Cond（别忘记条件变量sync.Cond）


func main() {
    type Donation struct {
        cond    *sync.Cond
        balance int
    }

    donation := &Donation{
        cond: sync.NewCond(&sync.Mutex{}),
    }

    // Listener goroutines
    f := func(goal int) {
        donation.cond.L.Lock()
        for donation.balance < goal {
            donation.cond.Wait()
        }
        fmt.Printf("%d$ goal reached\n", donation.balance)
        donation.cond.L.Unlock()
    }

    go f(10)
    go f(15)

    // Updater goroutine
    for {
        time.Sleep(time.Second)
        donation.cond.L.Lock()
        donation.balance++
        donation.cond.L.Unlock()
        donation.cond.Broadcast()
    }
}
使用Cond可以避免CPU空转的情况。!!!!!

9.13 #73: Not using errgroup（使用 errgroup 采集多个协程执行结果）

使用 golang.org/x/sync/errgroup 采集多个协程的错误, 但是 errorgroup限制比较多，需要函数签名符合func() error {}
其中一个goroutine错误会通知其他的goroutine停止 ！！！
 
k8s 中提供了集合多种错误的功能，可以参考 http://k8s.io/apimachinery/pkg

9.14 #74: Copying a sync type（同步类型的值不能被拷贝）
以下类型均不能被拷贝

sync.Cond
sync.Map
sync.Mutex
sync.RWMutex
sync.Once
sync.Pool
sync.WaitGroup

10.2 #76: time.After and memory leaks （使用 time.After 可能会导致内存溢出）
每次调用 time.After 时使用大约200字节的内存。但是只有在指定的时间到达的时候才会GC，如果在1小时内，频繁的调用 time.After 会导致内存爆炸。！！！

```go
func consumer1(ch <-chan Event) {
    for {
        select {
        case event := <-ch:
            handle(event)
        case <-time.After(time.Hour):
            log.Println("warning: no messages received")
        }
    }
}

// 解决方案一
func consumer2(ch <-chan Event) {
    for {
        ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
        select {
        case event := <-ch:
            cancel()
            handle(event)
        case <-ctx.Done():
            log.Println("warning: no messages received")
        }
    }
}

// 解决方案二，创建一次 Timer 每次循环都会重置时间
func consumer3(ch <-chan Event) {
    timerDuration := 1 * time.Hour
    timer := time.NewTimer(timerDuration)

    for {
        timer.Reset(timerDuration)
        select {
        case event := <-ch:
            handle(event)
        case <-timer.C:
            log.Println("warning: no messages received")
        }
    }
}
```

10.3 #77: Common JSON-handling mistakes（常见的 JSON 处理错误）
Unexpected behavior due to type embedding
如果类型实现了 Marshaler 接口，在调用 json.Marshal 的时候，会直接调用 MarshalJSON 方法。

JSON and the monotonic clock
可以使用 Time.Equal() 对比时间是否相关，这里判断相等不会包含单调时间
使用 Time.Truncate() 函数排除单调时钟

Map of any
当把json解析到 map[string]any 类型，会出现数字类型解析错误的情况
解决方案: 使用 json.Decoder 来代替 json.Unmarshal 方法
```go
decoder := json.NewDecoder(bytes.NewReader(getMessage()))
decoder.UseNumber()
var m map[string]any
decoder.Decode(&personFromJSON)
```


10.4 #78: Common SQL mistakes（常见的 SQL 错误）

Forgetting that sql.Open doesn’t necessarily establish connections to a database
sql.Open 可能只是验证其参数而不创建与数据库的连接

如果我们要确保使用 sql.Open 的函数也保证底层数据库可访问，我们应该使用Ping方法

Forgetting about connections pooling
sql.Open 返回一个sql.DB结构。这个结构不表示单个数据库连接，而是表示连接池。

Not using prepared statements
简单聊聊 SQL 中的 Prepared Statements
​manjusaka.itscoder.com/posts/2020/01/05/simple-introdution-about-sql-prepared/
使用 Prepare 方法创建 prepared statements ，提升查询性能，避免重复解析 SQL 带来的开销


10.5 #79: Not closing transient resources（没有及时关闭临时资源）
10.5.1 HTTP body

需要注意的点：

如果你没有读取Respose.Body的内容，那么默认的 http transport 会直接关闭连接
如果你读取了Body的内容，下次连接可以直接复用
在高并发的场景下，建议你使用长连接，可以调用 io.Copy(io.Discard, resp.Body) 读取Body的内容。
!!!!

10.5.2 sql.Rows
标准库只有close才会将连接归还给连接池

10.5.3 os.File
写入操作是异步的，所以对写入的文件进行close操作，可能会遇到在buffer内的数据没有写到磁盘的错误，所以在close的时候如果遇到错误要及时上报。
但是如果使用了Sync调用可以同步的把数据写入磁盘，所以调用Close方法的时候也可以不用在意错误，因为数据已经正常写入。


11.1 #82: Not categorizing tests（没有对测试类型分类）
不同类型的测试:单元测试、集成测试、E2E测试各自的执行时间和个数都有很大差距，如下的测试金字塔显示了测试的占比，各种测试的执行时间也呈反比。这一节介绍了一些方法提醒开发者，不同类型的测试，需要明确区分，并且最好独立执行，可以提升开发效率。

使用 go:build tag来给测试归类,用法可参考：
Separate Test Cases in Golang With Build Tags | Clivern
https://clivern.com/separate-test-cases-in-golang-with-build-tags/

```go
//go:build integration
// +build integration

package db

import (
	"os"
	"testing"
)

func TestInsert1(t *testing.T) {
	// ...
}
 
//使用环境变量分类
func TestInsert2(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test")
	}

	// ...
 
//Short mode，执行go test的时候加上 -short 选项，可以选择性的执行耗时短的测试
func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test")
	}
	// ...
}

// % go test -short -v .
// === RUN   TestLongRunning
//     foo_test.go:9: skipping long-running test
// --- SKIP: TestLongRunning (0.00s)
// PASS
// ok      foo  0.174s
 
```

11.2 #83: Not enabling the -race flag（没开启 -race flag来检验并发冲突）
开启 -race 之后性能和内存占用都有影响，所以生产环境不建议开启（废话。。）

go test -race ./...

11.3 #84: Not using test execution modes（没有使用test执行模式：并行 or shuffle ）
The parallel flag 开启并行模式执行测试
// 并行执行测试，并行最大个数 16
go test -parallel 16 .
The -shuffle flag 开启 shuffle 模式执行测试
go test -shuffle=on -v

开启 shuffle模式可以检测出更多未知错误，但是因为每次执行都是乱序，如果想复现错误，并Debug，就需要在
shuffle执行的时候，指定 shuffle id， 这样下次执行的时候就会按照特定的顺序执行


11.4 #85: Not using table-driven tests（使用 table-driven 的模式编写测试）
https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

1.5 #86: Sleeping in unit tests（在测试代码中包含sleep逻辑，导致出现flaky测试）
依赖等待一段时间直到特定的逻辑执行完成的方式会因为时间设置问题导致出现不稳定的测试，这种情况，可以通过指定多次重试或者使用同步的方法来避免直接调用 time.sleep 来等待。

可以使用 testify 或者 Gomega 中的 Eventually 方法

Testing for asynchronous results without sleep in Go

11.7 #88: Not using testing utility packages（没有使用内置的工具包，比如httptest、iotest）
略，此处就不讲解包的使用了，知道在mock http 和 io 操作的时候可以使用这2个内置的工具包

httptest package - net/http/httptest - Go Packages

iotest package - testing/iotest - Go Packages

11.8 #89: Writing inaccurate benchmarks（编写了不正确的BencMark）
Not resetting or pausing the timer

在具体测量某一段函数性能时，一些SetUp操作可能比较耗时会影响测量结果。调用 ResetTimer 函数可以重置一些Bench数据。

每次迭代都会有一些耗时动作，可以调用 b.StopTimer() 和 b.StartTimer() 暂停和启动BenchMark计量


Making wrong assumptions about micro-benchmarks
在一些做一些 micro-benchmark 的时候，如果不多次进行基准测试很容易就得出错误的结论

解决方案：对于 micro-benchmark 需要进行多次 BenchMark，可以利用 benchstat 统计BenchMark的结果进行均值计算。!!!

go test -bench=. -count=10 | tee stats.txt
利用 benchstat汇总结果：benchstat stats.txt

Not being careful about compiler optimizations
golang的内联优化会导致我们的测试函数被优化，执行结果不符合我们的预期（效果更好）

https://github.com/golang/go/issues/14813

How to write benchmarks in Go
https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

给`go test`增加`-gcflags="-m"`参数，`-m`表示打印编译器做出的优化决定。可以看到是否做了内联优化，如果有 `inlining call to xxx 函数` 的字样就是做了优化

执行go test的时候，增加-gcfloags="-l"参数，-l表示禁用编译器的内联优化。
使用//go:noinline 编译器指令(compiler directive)，编译器在编译时会识别到这个指令，不做内联优化。
 //go:noinline 
func add(a int, b int) int {  return a + b }

Being fooled by the observer effect
在物理学中，观测者效应是观测行为对被观测系统的扰动。这种影响也可以在基准测试中看到，并可能导致对结果的错误假设。
因为BenchMark事先创建了一个矩阵并重复计算，再加上缓存命中的影响更加加重了测试差距，解决方案就是每次迭代都创建新的矩阵

11.9 #90: Not exploring all the Go testing features（没有利用 Go testing的特性）
使用 coverprofile flag 查看代码覆盖率
// 获得测试覆盖率文件
go test -coverprofile=coverage.out ./...

// 可视化展示覆盖率
go tool cover -html=coverage.out


12.1 #91: Not understanding CPU caches（利用缓存加速代码执行速度）
关于CPU Cache部分不属于GO专属的知识，这里就不细讲了

Sinaean Dean：细说Cache-L1/L2/L3/TLB
763 赞同 · 62 评论文章

cache line 是固定大小的连续内存段，通常为64字节（8个int64变量）。

12.2 #92: Writing concurrent code that leads to false sharing（编写并发代码的时候，注意避免伪共享）


12.3 #93: Not taking into account instruction-level parallelism（没有考虑到CPU指令并行优化）

12.4 #94: Not being aware of data alignment（没有数据对齐的意识）
Go struct 内存对齐 | Go 语言高性能编程 | 极客兔兔
​geektutu.com/post/hpg-struct-alignment.html

12.5 #95: Not understanding stack vs. heap（不理解堆栈和逃逸分析）

```go
// -gcflags "-m=2" 可以帮助查看是否有逃逸
$ go build -gcflags "-m=2"
...
./main.go:12:2: z escapes to heap:
```

12.6 #96: Not knowing how to reduce allocations（减少内存分配）
本书已经讲解了很多种优化内存的方式：

39 使用strings.Builder拼接字符串
40 避免不必要的string和[]byte类型转换
21和#27给slice和map预先分配内存
94 用更好的内存分配方式减少内存占用

sync.Pool

Go 语言从 1.3 版本开始提供了对象重用的机制，即 sync.Pool。sync.Pool 是可伸缩的，同时也是并发安全的，其大小仅受限于内存的大小。sync.Pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值。这样就可以不用再次经过内存分配，可直接复用已有对象，减轻 GC 的压力，从而提升系统的性能。

sync.Pool 的大小是可伸缩的，高负载时会动态扩容，存放在池中的对象如果不活跃了会被自动清理。

12.7 #97: Not relying on inlining（忘记依赖 inlining 编译优化）
内联有两个主要好处。首先，它消除了函数调用的开销（即使自Go 1.17和基于寄存器的调用约定以来开销已经减轻）。其次，它允许编译器进行进一步的优化。例如，在内联函数后，编译器可以决定把一些逃逸的变量放在堆上。

12.9 #99: Not understanding how the GC works（不理解GC是如何工作的）
Go 语言垃圾收集器的实现原理
​draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/
垃圾回收的优化问题
​golang.design/go-questions/memgc/optimize/
比较重要的问题是，每次GC是何时发生？与Java等其他语言相比，Go 配置仍然相当简单。它依赖于一个环境变量：GOGC。该变量定义自上次GC触发另一个GC之前的堆增长百分比：默认值为100%。

假设一次GC刚刚被触发，当前堆大小为128 MB。如果 GOGC=100，则当堆大小达到256 MB时触发下一次GC。

默认情况下，每次堆大小翻倍时都会执行GC。此外，如果在过去2分钟内没有执行GC，Go将强制运行一次。

在大多数操作系统上，分配这个 min 变量不会使我们的应用程序消耗1 GB内存。调用make会导致系统调用mmap()，而mmap调用会懒分配内存。

12.10 #100: Not understanding the impacts of running Go in Docker and Kubernetes（GO应用在k8s环境下运行会遇到错误分配GOMAXPROCS的场景）
如果k8s运行的环境不是安全容器，进程读到的全局CPU核数和宿主机一致，会导致错误配置了GOMAXPROCS的值和宿主机一样，GO默认开启的协程个数就会远超容器实际运行环境提供的CPU个数，导致协程频繁的调度切换程序运行时间被拖慢。

解决方案：使用 automaxprocs 包来配置GOMAXPROCS

https://github.com/uber-go/automaxprocs
​github.com/uber-go/automaxprocs
但是现在大部分生产环境的容器运行时都是安全容器，隔离性更强不会出现错误配置GOMAXPROCS的情况













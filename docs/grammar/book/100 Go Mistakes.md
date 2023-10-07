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




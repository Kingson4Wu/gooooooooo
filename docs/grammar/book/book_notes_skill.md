# 《Go语言圣经》

+ map的遍历是随机的
+ Printf打印对照
+ 这个程序里的io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
+  `secs := time.Since(start).Seconds()`
+ 用if和ParseForm结合可以让代码更加简单，并且可以限制err这个变量的作用域，这么做是很不错的。
```go
if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
```
+ 无tag switch(tagless switch)；这和switch true是等价的。
+ Go语言程序员推荐使用 驼峰式 命名
+ 由于new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型

+ f函数里的x变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到，虽然它是在函数内部定义的；用Go语言的术语说，这个x局部变量从函数f中逃逸了。相反，当g函数返回时，变量*y将是不可达的，也就是说可以马上被回收的。因此，*y并没有从函数g中逃逸，编译器可以选择在栈上分配*y的存储空间（译注：也可以选择在堆上分配，然后由Go语言的GC回收这个变量的内存空间），虽然这里用的是new方式。其实在任何时候，你并不需为了编写正确的代码而要考虑变量的逃逸行为，要记住的是，逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。

+ 我们在这个包声明了两种类型：Celsius和Fahrenheit分别对应不同的温度单位。它们虽然有着相同的底层类型float64，但是它们是不同的数据类型，因此它们不可以被相互比较或混在一个表达式运算。刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误

+ init
+ 每个包在解决依赖的前提下，以导入声明的顺序初始化，每个包只会被初始化一次。因此，如果一个p包导入了q包，那么在p包初始化的时候可以认为q包必然已经初始化过了。初始化工作是自下而上进行的，main包最后被初始化。以这种方式，可以确保在main函数执行之前，所有依赖的包都已经完成初始化工作了。
不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时的属性。一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时的概念。

## 第5章　函数
+ golang.org/x/... 目录下存储了一些由Go团队设计、维护，对网络编程、国际化文件处理、移动平台、图像处理、加密解密、开发者工具提供支持的扩展包。未将这些扩展包加入到标准库原因有二，一是部分包仍在开发中，二是对大多数Go语言的开发者而言，扩展包提供的功能很少被使用。

+ 对于那些将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递错误信息。如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok。比如，cache.Lookup失败的唯一原因是key不存在，那么代码可以按照下面的方式组织：

```go
value, ok := cache.Lookup(key)
if !ok {
    // ...cache[key] does not exist…
}
```

+ 调用者只需通过简单的比较，就可以检测出这个错误。下面的例子展示了如何从标准输入中读取字符，以及判断文件结束。（4.3的chartcount程序展示了更加复杂的代码）
```go
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed:%v", err)
    }
    // ...use r…
}
```

+ 函数值之间是不可比较的，也不能用函数值作为map的key。

+ strings.Map对字符串中的每个字符调用add1函数，并将每个add1函数的返回值组成一个新的字符串返回给调用者。
```go
    func add1(r rune) rune { return r + 1 }

    fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
    fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
    fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
```

+ 更为重要的是，通过这种方式定义的函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以引用该函数的变量
```go
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}
```
+ squares的例子证明，函数值不仅仅是一串代码，还记录了状态。在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。这就是函数值属于引用类型和函数值不可比较的原因。Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包

+ 5.6.1. 警告：捕获迭代变量 !!!
+ 如果你使用go语句（第八章）或者defer语句（5.8节）会经常遇到此类问题。这不是go或defer本身导致的，而是因为它们都会等待循环结束后，再执行函数值。!!!

+ 你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。

+ 基于以上原因，安全的做法是有选择性的recover。换句话说，只恢复应该被恢复的panic异常，此外，这些异常所占的比例应该尽可能的低。为了标识某个panic是否应该被恢复，我们可以将panic value设置成特殊类型。在recover时对panic value进行检查，如果发现panic value是特殊类型，就将这个panic作为error处理，如果不是，则按照正常的panic进行处理（在下面的例子中，我们会看到这种方式）。

+ deferred函数调用recover，并检查panic value。当panic value是bailout{}类型时，deferred函数生成一个error返回给调用者。当panic value是其他non-nil值时，表示发生了未知的panic异常，deferred函数将调用panic函数并将当前的panic value作为参数传入；此时，等同于recover没有做任何操作。（请注意：在例子中，对可预期的错误采用了panic，这违反了之前的建议，我们在此只是想向读者演示这种机制。）
有些情况下，我们无法恢复。某些致命错误会导致Go在运行时终止程序，如内存不足。

## 第7章　接口

+ 7.12. 通过类型断言询问行为
+ 下面这段逻辑和net/http包中web服务器负责写入HTTP头字段（例如："Content-type:text/html"）的部分相似。io.Writer接口类型的变量w代表HTTP响应；写入它的字节最终被发送到某个人的web浏览器上。
```go
func writeHeader(w io.Writer, contentType string) error {
    if _, err := w.Write([]byte("Content-Type: ")); err != nil {
        return err
    }
    if _, err := w.Write([]byte(contentType)); err != nil {
        return err
    }
    // ...
}
```

+ 因为Write方法需要传入一个byte切片而我们希望写入的值是一个字符串，所以我们需要使用[]byte(...)进行转换。这个转换分配内存并且做一个拷贝，但是这个拷贝在转换后几乎立马就被丢弃掉。让我们假装这是一个web服务器的核心部分并且我们的性能分析表示这个内存分配使服务器的速度变慢。这里我们可以避免掉内存分配么？
这个io.Writer接口告诉我们关于w持有的具体类型的唯一东西：就是可以向它写入字节切片。如果我们回顾net/http包中的内幕，我们知道在这个程序中的w变量持有的动态类型也有一个允许字符串高效写入的WriteString方法；这个方法会避免去分配一个临时的拷贝。（这可能像在黑夜中射击一样，但是许多满足io.Writer接口的重要类型同时也有WriteString方法，包括*bytes.Buffer，*os.File和*bufio.Writer。）
我们不能对任意io.Writer类型的变量w，假设它也拥有WriteString方法。但是我们可以定义一个只有这个方法的新接口并且使用类型断言来检测是否w的动态类型满足这个新接口。

```go
// writeString writes s to w.
// If w has a WriteString method, it is invoked instead of w.Write.
func writeString(w io.Writer, s string) (n int, err error) {
    type stringWriter interface {
        WriteString(string) (n int, err error)
    }
    if sw, ok := w.(stringWriter); ok {
        return sw.WriteString(s) // avoid a copy
    }
    return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
    if _, err := writeString(w, "Content-Type: "); err != nil {
        return err
    }
    if _, err := writeString(w, contentType); err != nil {
        return err
    }
    // ...
}
```

+ 为了避免重复定义，我们将这个检查移入到一个实用工具函数writeString中，但是它太有用了以致于标准库将它作为io.WriteString函数提供。这是向一个io.Writer接口写入字符串的推荐方法。
这个例子的神奇之处在于，没有定义了WriteString方法的标准接口，也没有指定它是一个所需行为的标准接口。一个具体类型只会通过它的方法决定它是否满足stringWriter接口，而不是任何它和这个接口类型所表达的关系。它的意思就是上面的技术依赖于一个假设，这个假设就是：如果一个类型满足下面的这个接口，然后WriteString(s)方法就必须和Write([]byte(s))有相同的效果。
有一个将单个操作对象转换成一个字符串的步骤，像下面这样：

```go
package fmt

func formatOneValue(x interface{}) string {
    if err, ok := x.(error); ok {
        return err.Error()
    }
    if str, ok := x.(Stringer); ok {
        return str.String()
    }
    // ...all other types...
}
```
+ 如果x满足这两个接口类型中的一个，具体满足的接口决定对值的格式化方式。如果都不满足，默认的case或多或少会统一地使用反射来处理所有的其它类型；我们可以在第12章知道具体是怎么实现的。
再一次的，它假设任何有String方法的类型都满足fmt.Stringer中约定的行为，这个行为会返回一个适合打印的字符串。
接口被以两种不同的方式使用。在第一个方式中，以io.Reader，io.Writer，fmt.Stringer，sort.Interface，http.Handler和error为典型，一个接口的方法表达了实现这个接口的具体类型间的相似性，但是隐藏了代码的细节和这些具体类型本身的操作。重点在于方法上，而不是具体的类型上。

+ 第二个方式是利用一个接口值可以持有各种具体类型值的能力，将这个接口认为是这些类型的联合。类型断言用来动态地区别这些类型，使得对每一种情况都不一样。在这个方式中，重点在于具体的类型满足这个接口，而不在于接口的方法（如果它确实有一些的话），并且没有任何的信息隐藏。我们将以这种方式使用的接口描述为discriminated unions（可辨识联合）。
如果你熟悉面向对象编程，你可能会将这两种方式当作是subtype polymorphism（子类型多态）和 ad hoc polymorphism（非参数多态）

## 第8章　Goroutines和Channels

+  Go语言中的并发程序可以用两种手段来实现。本章讲解goroutine和channel，其支持“顺序通信进程”（communicating sequential processes）或被简称为CSP。CSP是一种现代的并发编程模型，在这种编程模型中值会在不同的运行实例（goroutine）中传递，尽管大多数情况下仍然是被限制在单一实例中

```go
func main() {
    // ...create abort channel...

    fmt.Println("Commencing countdown.  Press return to abort.")
    select {
    case <-time.After(10 * time.Second):
        // Do nothing.
    case <-abort:
        fmt.Println("Launch aborted!")
        return
    }
    launch()
}
```

## ## 第9章　基于共享变量的并发

+ 我们将使用下面的httpGetBody函数作为我们需要缓存的函数的一个样例。这个函数会去进行HTTP GET请求并且获取http响应body。对这个函数的调用本身开销是比较大的，所以我们尽量避免在不必要的时候反复调用。 
```go
func httpGetBody(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
```

+ Goroutine没有ID号

## 第10章　包和工具

+ 第一点，所有导入的包必须在每个文件的开头显式声明，这样的话编译器就没有必要读取和分析整个源文件来判断包的依赖关系。第二点，禁止包的环状依赖，因为没有循环依赖，包的依赖关系形成一个有向无环图，每个包可以被独立编译，而且很可能是被并发编译。第三点，编译后包的目标文件不仅仅记录包本身的导出信息，目标文件同时还记录了包的依赖关系。因此，在编译一个包的时候，编译器只需要读取每个直接导入包的目标文件，而不需要遍历所有依赖的的文件

+ 10.7.5. 内部包
+ Go语言的构建工具对包含internal名字的路径段的包导入路径做了特殊处理。这种包叫internal包，一个internal包只能被和internal目录有同一个父目录的包所导入。例如，net/http/internal/chunked内部包只能被net/http/httputil或net/http包导入，但是不能被net/url包导入。不过net/url包却可以导入net/http/httputil包。

## 第11章　测试

+ 在*_test.go文件中，有三种类型的函数：测试函数、基准测试（benchmark）函数、示例函数。

+ go test -v -run=Coverage
+ go test -run=Coverage -coverprofile=c.out
+ go test -cover

```shell
$ go test -cpuprofile=cpu.out
$ go test -blockprofile=block.out
$ go test -memprofile=mem.out

$ go test -run=NONE -bench=ClientServerParallelTLS64 \
    -cpuprofile=cpu.log net/http
 PASS
 BenchmarkClientServerParallelTLS64-8  1000
    3141325 ns/op  143010 B/op  1747 allocs/op
ok       net/http       3.395s

$ go tool pprof -text -nodecount=10 ./http.test cpu.log
2570ms of 3590ms total (71.59%)
```

## 第12章　反射
+ 12.2. reflect.Type 和 reflect.Value
+ 反射是由 reflect 包提供的。它定义了两个重要的类型，Type 和 Value。一个 Type 表示一个Go类型。它是一个接口，有许多方法来区分类型以及检查它们的组成部分，例如一个结构体的成员或一个函数的参数等。唯一能反映 reflect.Type 实现的是接口的类型描述信息（§7.5），也正是这个实体标识了接口值的动态类型。

## 第13章　底层编程
+ 13.1. unsafe.Sizeof, Alignof 和 Offsetof

+ 13.2. unsafe.Pointer

+ 13.4. 通过cgo调用C代码

----

# 《Go语言高级编程》

+ 语言基因族谱
+ 首先看基因图谱的左边一支。可以明确看出 Go 语言的并发特性是由贝尔实验室的 Hoare 于 1978 年发布的 CSP 理论演化而来。其后，CSP 并发模型在 Squeak/NewSqueak 和 Alef 等编程语言中逐步完善并走向实际应用，最终这些设计经验被消化并吸收到了 Go 语言中。业界比较熟悉的 Erlang 编程语言的并发编程模型也是 CSP 理论的另一种实现。

```go
func main() {
    done := make(chan int)

    go func(){
        println("你好, 世界")
        done <- 1
    }()

    <-done
}
```

+ 严谨的并发程序的正确性不应该是依赖于 CPU 的执行速度和休眠时间等不靠谱的因素的。严谨的并发也应该是可以静态推导出结果的：根据线程内顺序一致性，结合 Channel 或 sync 同步事件的可排序性来推导，最终完成各个线程各段代码的偏序关系排序。如果两个事件无法根据此规则来排序，那么它们就是并发的，也就是执行先后顺序不可靠的。
解决同步问题的思路是相同的：使用显式的同步。

### 1.6 常见的并发模式
+ Go 语言最吸引人的地方是它内建的并发支持。Go 语言并发体系的理论是 C.A.R Hoare 在 1978 年提出的 CSP（Communicating Sequential Process，通讯顺序进程）。CSP 有着精确的数学模型，并实际应用在了 Hoare 参与设计的 T9000 通用计算机上。从 NewSqueak、Alef、Limbo 到现在的 Go 语言，对于对 CSP 有着 20 多年实战经验的 Rob Pike 来说，他更关注的是将 CSP 应用在通用编程语言上产生的潜力。作为 Go 并发编程核心的 CSP 理论的核心概念只有一个：同步通信。关于同步通信的话题我们在前面一节已经讲过，本节我们将简单介绍下 Go 语言中常见的并发模式。

+ Do not communicate by sharing memory; instead, share memory by communicating.
不要通过共享内存来通信，而应通过通信来共享内存。
+ 这是更高层次的并发编程哲学(通过管道来传值是 Go 语言推荐的做法)。虽然像引用计数这类简单的并发问题通过原子操作或互斥锁就能很好地实现，但是通过 Channel 来控制访问能够让你写出更简洁正确的程序。

## 第 2 章 CGO 编程
## 第 3 章 Go 汇编语言
## 第 4 章 RPC 和 Protobuf
## 第 5 章 go 和 Web
## 第 6 章 分布式系统

## 附录A：Go语言常见坑

+ 数组是值传递
在函数调用参数中，数组是值传递，无法通过修改数组类型的参数返回结果。
```go
func main() {
    x := [3]int{1, 2, 3}

    func(arr [3]int) {
        arr[0] = 7
        fmt.Println(arr)
    }(x)

    fmt.Println(x)
}
```

+ recover必须在defer函数中运行
recover捕获的是祖父级调用时的异常，直接调用时无效：
```go
func main() {
    recover()
    panic(1)
}
```
直接defer调用也是无效：
```go
func main() {
    defer recover()
    panic(1)
}
```
defer调用时多层嵌套依然无效：
```go
func main() {
    defer func() {
        func() { recover() }()
    }()
    panic(1)
}
```
必须在defer函数中直接调用才有效：
```go
func main() {
    defer func() {
        recover()
    }()
    panic(1)
}
```

+ 独占CPU导致其它Goroutine饿死
Goroutine 是协作式抢占调度（Go1.14版本之前），Goroutine本身不会主动放弃CPU：
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()

    for {} // 占用CPU
}
```
解决的方法是在for循环加入runtime.Gosched()调度函数：
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()

    for {
        runtime.Gosched()
    }
}
```
或者是通过阻塞的方式避免CPU占用：
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        os.Exit(0)
    }()

    select{}
}
```

+ 闭包错误引用同一个变量
```go
func main() {
    for i := 0; i < 5; i++ {
        defer func() {
            println(i)
        }()
    }
}
```
改进的方法是在每轮迭代中生成一个局部变量：
```go
func main() {
    for i := 0; i < 5; i++ {
        i := i
        defer func() {
            println(i)
        }()
    }
}
```
或者是通过函数参数传入：
```go
func main() {
    for i := 0; i < 5; i++ {
        defer func(i int) {
            println(i)
        }(i)
    }
}
```

+ 在循环内部执行defer语句
defer在函数退出时才能执行，在for执行defer会导致资源延迟释放：
```go
func main() {
    for i := 0; i < 5; i++ {
        f, err := os.Open("/path/to/file")
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
    }
}
```
解决的方法可以在for中构造一个局部函数，在局部函数内部执行defer：
```go
func main() {
    for i := 0; i < 5; i++ {
        func() {
            f, err := os.Open("/path/to/file")
            if err != nil {
                log.Fatal(err)
            }
            defer f.Close()
        }()
    }
}
```
+ 全部看一遍吧,...


## 附录B：有趣的代码片段
这里收集一些比较有意思的Go程序片段。

+ 禁止 main 函数退出的方法
```go 
func main() {
    defer func() { for {} }()
}

func main() {
    defer func() { select {} }()
}

func main() {
    defer func() { <-make(chan bool) }()
}
```

+ Assert测试断言
```go
type testing_TBHelper interface {
    Helper()
}

func Assert(tb testing.TB, condition bool, args ...interface{}) {
    if x, ok := tb.(testing_TBHelper); ok {
        x.Helper() // Go1.9+
    }
    if !condition {
        if msg := fmt.Sprint(args...); msg != "" {
            tb.Fatalf("Assert failed, %s", msg)
        } else {
            tb.Fatalf("Assert failed")
        }
    }
}

func Assertf(tb testing.TB, condition bool, format string, a ...interface{}) {
    if x, ok := tb.(testing_TBHelper); ok {
        x.Helper() // Go1.9+
    }
    if !condition {
        if msg := fmt.Sprintf(format, a...); msg != "" {
            tb.Fatalf("Assertf failed, %s", msg)
        } else {
            tb.Fatalf("Assertf failed")
        }
    }
}

func AssertFunc(tb testing.TB, fn func() error) {
    if x, ok := tb.(testing_TBHelper); ok {
        x.Helper() // Go1.9+
    }
    if err := fn(); err != nil {
        tb.Fatalf("AssertFunc failed, %v", err)
    }
}

```

----

# 《Go语言精进之路》

# 一部分 熟知Go语言的一切

## 3条 理解Go语言的设计哲学
### 3.1　追求简单，少即是多
### 3.2　偏好组合，正交解耦
Go采用了组合的方式，也是唯一的方式。

+ 垂直扩展

```go
// $GOROOT/src/sync/pool.go
type poolLocal struct {
    private interface{}
    shared  []interface{}
    Mutex
    pad     [128]byte
}
```
我们在poolLocal这个结构体类型中嵌入了类型Mutex

+ 水平组合
通过interface将程序各个部分组合在一起的方法，笔者称之为“水平组合”。水平组合的模式有很多，一种常见的方法是通过接受interface类型参数的普通函数进行组合，例如下面的代码。

```go
// $GOROOT/src/io/ioutil/ioutil.go
func ReadAll(r io.Reader)([]byte, error)

// $GOROOT/src/io/io.go
func Copy(dst Writer, src Reader)(written int64, err error)
```

+ Go语言内置的并发能力也可以通过组合的方式实现对计算能力的串联，比如通过goroutine+channel的组合实现类似Unix Pipe的能力。

+ 类型嵌入为类型提供垂直扩展能力，interface是水平组合的关键，它好比程序肌体上的“关节”，给予连接“关节”的两个部分各自“自由活动”的能力，而整体上又实现了某种功能。

### 3.3　原生并发，轻量高效

（1）Go语言采用轻量级协程并发模型，使得Go应用在面向多核硬件时更具可扩展性
（2）Go语言为开发者提供的支持并发的语法元素和机制
 (3）并发原则对Go开发者在程序结构设计层面的影响

### 3.4　面向工程，“自带电池”

## 第4条　使用Go语言原生编程思维来写Go代码

### 4.1　语言与思维——来自大师的观点

### 4.2　现实中的“投影”
（3）Go语言版本

```go
// chapter1/sources/sieve.go

func Generate(ch chan<- int) {
    for i := 2; ; i++ {
        ch <- i
    }
}

func Filter(in <-chan int, out chan<- int, prime int) {
    for {
        i := <-in
        if i%prime != 0 {
            out <- i
        }
    }
}

func main() {
    ch := make(chan int)
    go Generate(ch)
    for i := 0; i < 10; i++ {
        prime := <-ch
        print(prime, "\n")
        ch1 := make(chan int)
        go Filter(ch, ch1, prime)
        ch = ch1
    }
}
```
### 4.3　Go语言原生编程思维

# 第二部分　项目结构、代码风格与标识符命名

## 第5条　使用得到公认且广泛使用的项目结构

### 1. Go项目结构的最小标准布局

非官方标准的建议结构布局

### 2. 以构建二进制可执行文件为目的的Go项目结构

一个支持（在cmd下）构建二进制可执行文件的典型Go项目的结构，我们分别来看一下各个重要目录的用途。
### 3. 以只构建库为目的的Go项目结构

库类型项目结构与Go项目的最小标准布局也是兼容的，但比以构建二进制可执行文件为目的的Go项目要简单一些。

去除了cmd和pkg两个子目录：由于仅构建库，没必要保留存放二进制文件main包源文件的cmd目录；由于Go库项目的初衷一般都是对外部（开源或组织内部公开）暴露API，因此也没有必要将其单独聚合到pkg目录下面了。

### 4. 关于internal目录
对于不想暴露给外部引用，仅限项目内部使用的包，在项目结构上可以通过Go 1.4版本中引入的internal包机制来实现。以库项目为例，最简单的方式就是在顶层加入一个internal目录，将不想暴露到外部的包都放在该目录下，比如下面项目结构中的ilib1、ilib2

<pre>
// 带internal的Go库项目结构

$tree -F ./chapter2/sources/GoLibProj
GoLibProj
├── LICENSE
├── Makefile
├── README.md
├── go.mod
├── internal/
│  ├── ilib1/
│  └── ilib2/
├── lib.go
├── lib1/
│  └── lib1.go
└── lib2/
      └── lib2.go
</pre>

## 第6条　提交前使用gofmt格式化源码

### 6.1　gofmt：Go语言在解决规模化问题上的最佳实践

### 6.2　使用gofmt
### 6.3　使用goimports
### 6.4　将gofmt/goimports与IDE或编辑器工具集成

## 第7条　使用Go命名惯例对标识符进行命名

要想做好Go标识符的命名（包括对包的命名），至少要遵循两个原则：简单且一致；利用上下文辅助命名。

### 7.1　简单且一致
1. 包对于Go中的包（package），一般建议以小写形式的单个单词命名。
2. 变量、类型、函数和方法
Go语言官方要求标识符命名采用驼峰命名法（CamelCase）

3. 常量
在Go语言中，常量在命名方式上与变量并无较大差别，并不要求全部大写。只是考虑其含义的准确传递，常量多使用多单词组合的方式命名。
4. 接口
在Go语言中，对于接口类型优先以单个单词命名。对于拥有唯一方法（method）或通过多个拥有唯一方法的接口组合而成的接口，Go语言的惯例是用“方法名+er”命名。

Go语言推荐尽量定义小接口，并通过接口组合的方式构建程序

### 7.2　利用上下文环境，让最短的名字携带足够多的信息

### 小结
Go语言命名惯例深受C语言的影响，这与Go语言之父有着深厚的C语言背景不无关系。Go语言追求简单一致且利用上下文辅助名字信息传达的命名惯例，如果你刚从其他语言转向Go，这可能会让你感到不适应，但这就是Go语言文化的一部分，也许等你编写的Go代码达到一定的量，你就能理解这种命名惯例的好处了。

# 第三部分　声明、类型、语句与控制结构

## 第8条　使用一致的变量声明形式

### 8.1　包级变量的声明形式
1. 声明并同时显式初始化
从声明一致性的角度出发，Go语言官方更推荐后者，这样就统一了接受默认类型和显式指定类型两种声明形式。尤其是在将这些变量放在一个var块中声明时，我们更青睐这样的形式：

```go
var (
    a = 17
    f = float32(3.14)
)
```
2. 声明但延迟初始化
虽然没有显式初始化，但Go语言会让这些变量拥有初始的“零值”。如果是自定义的类型，保证其零值可用是非常必要的，这一点将在后文中详细说明。
3. 声明聚类与就近原则

### 8.2　局部变量的声明形式
1. 对于延迟初始化的局部变量声明，采用带有var关键字的声明形式
另一种常见的采用带var关键字声明形式的变量是error类型的变量err（将error类型变量实例命名为err也是Go的一个惯用法），尤其是当defer后接的闭包函数需要使用err判断函数/方法退出状态时。

```go
func Foo() {
    var err error
    defer func() {
        if err != nil {
            ...
        }
    }()

    err = Bar()
    ...
}
```
2. 对于声明且显式初始化的局部变量，建议使用短变量声明形式
3. 尽量在分支控制时应用短变量声明形式
由于良好的函数/方法设计讲究的是“单一职责”，因此每个函数/方法规模都不大，很少需要应用var块来聚类声明局部变量。当然，如果你在声明局部变量时遇到适合聚类的应用场景，你也应该毫不犹豫地使用var块来声明多个局部变量。

```go
// $GOROOT/src/net/dial.go
func (r *Resolver) resolveAddrList(ctx context.Context, op, network,
                            addr string, hint Addr) (addrList, error) {
    ...
    var (
        tcp      *TCPAddr
        udp      *UDPAddr
        ip       *IPAddr
        wildcard bool
    )
    ...
}
```

## 第9条　使用无类型常量简化代码
```go
// $GOROOT/src/io/io.go
const (
    SeekStart   = 0
    SeekCurrent = 1
    SeekEnd     = 2
)
```

###  9.3 无类型常量消除烦恼，简化代码
+ 无类型常量也拥有自己的默认类型：无类型的布尔型常量、整数常量、字符常量、浮点数常量、复数常量、字符串常量对应的默认类型分别为bool、int、int32(rune)、float64、complex128和string。

### 小结


+ 所有常量表达式的求值计算都可以在编译期而不是在运行期完成，这样既可以减少运行时的工作，也能方便编译器进行编译优化。

## 第10条 使用iota实现枚举常量

```go
 const (    
    _ = iota      // 0    
    Pin1    
    Pin2    
    Pin3    
    _   // 相当于_ = iota，略过了4这个枚举值    
    Pin5    // 5
    )
```
```go
const (   
     _ = iota    
     Blue    
     Black    
     Red    
     Yellow
     )
```
+ 枚举常量多数是无类型常量，如果要严格考虑类型安全，也可以定义有类型枚举常量
```go
// $GOROOT/src/time/time.go
type Weekday intconst (    
    Sunday 
    Weekday = iota    
    Monday    
    Tuesday    
    Wednesday    
    Thursday    
    Friday    
    Saturday
    )
```

### 11.1 Go类型的零值

+ 当通过声明或调用new为变量分配存储空间，或者通过复合文字字面量或调用make创建新值，且不提供显式初始化时，Go会为变量或值提供默认值

+ Go语言中的每个原生类型都有其默认值，这个默认值就是这个类型的零值。下面是Go规范定义的内置原生类型的默认值（零值）。所有整型类型：0浮点类型：0.0布尔类型：false字符串类型：""指针、interface、切片（slice）、channel、map、function：nil

+ Go的零值初始是递归的，即数组、结构体等类型的零值初始化就是对其组成元素逐一进行零值初始化。

### 11.2 零值可用

+ 在Go标准库和运行时代码中还有很多践行“零值可用”理念的好例子，最典型的莫过于sync.Mutex和bytes.Buffer了。

+ 无须对bytes.Buffer类型的变量b进行任何显式初始化，即可直接通过b调用Buffer类型的方法进行写入操作。这是因为bytes.Buffer结构体用于存储数据的字段buf支持零值可用策略的切片类型

### 小结

+ Go语言零值可用的理念给内置类型、标准库的使用者带来很多便利。不过Go并非所有类型都是零值可用的，并且零值可用也有一定的限制，比如：在append场景下，零值可用的切片类型不能通过下标形式操作数据：

+ 像map这样的原生类型也没有提供对零值可用的支持

+ 另外零值可用的类型要注意尽量避免值复制：

+ 我们可以通过指针方式传递类似Mutex这样的类型：
```go
var mu 
sync.Mutexfoo(&mu) // 正确
```

## 第12条 使用复合字面值作为初值构造器

+ Go语言中的复合类型包括结构体、数组、切片和map。

+ Go提供的复合字面值（composite literal）语法可以作为复合类型变量的初值构造器。

```go
s := myStruct{"tony", 23}
a := [5]int{13, 14, 15, 16, 17}
sl := []int{23, 24, 25, 26, 27}
m := map[int]string {1:"hello", 2:"gopher", 3:"!"}
```

### 小结
+ 对于零值不适用的场景，我们要为变量赋予一定的初值。对于复合类型，我们应该首选Go提供的复合字面值作为初值构造器。对于不同复合类型，我们要记住下面几点：使用field:value形式的复合字面值为结构体类型的变量赋初值；在为稀疏元素赋值或让编译器推导数组大小的时候，多使用index:value的形式为数组/切片类型变量赋初值；使用key:value形式的复合字面值为map类型的变量赋初值。（Go 1.5版本后，复合字面值中的key和value类型均可以省略不写。）

### 13.1 切片究竟是什么

+ Go数组是值语义的，这意味着一个数组变量表示的是整个数组，这点与C语言完全不同。在C语言中，数组变量可视为指向数组第一个元素的指针。而在Go语言中传递数组是纯粹的值拷贝，对于元素类型长度较大或元素个数较多的数组，如果直接以数组类型参数传递到函数中会有不小的性能损耗。这时很多人会使用数组指针类型来定义函数参数，然后将数组地址传进函数，这样做的确可以避免性能损耗，但这是C语言的惯用法，在Go语言中，更地道的方式是使用切片。

+ 切片之所以能在函数参数传递时避免较大性能损耗，是因为它是“描述符”的特性，切片这个描述符是固定大小的，无论底层的数组元素类型有多大，切片打开的窗口有多长。

+ 还可以通过语法s[low: high]基于已有切片创建新的切片，这被称为切片的reslicing，

+ 新创建的切片与原切片同样是共享底层数组的，并且通过新切片对数组的修改也会反映到原切片中。

+ 当切片作为函数参数传递给函数时，实际传递的是切片的内部表示，也就是上面的runtime.slice结构体实例，因此无论切片描述的底层数组有多大，切片作为参数传递带来的性能损耗都是很小且恒定的，甚至小到可以忽略不计，这就是函数在参数中多使用切片而不用数组指针的原因之一。

+ 而另一个原因就是切片可以提供比指针更为强大的功能，比如下标访问、边界溢出校验、动态扩容等。

### 13.2 切片的高级特性：动态扩容


+ 零值切片也可以通过append预定义函数进行元素赋值操作：
```go
var s []byte // s被赋予零值nil
s = append(s, 1)
```
由于初值为零值，s这个描述符并没有绑定对应的底层数组。而经过append操作后，s显然已经绑定了属于它的底层数组。

+ append会根据切片对底层数组容量的需求对底层数组进行动态调整。

+ 通过语法u[low: high]形式进行数组切片化而创建的切片，一旦切片cap触碰到数组的上界，再对切片进行append操作，切片就会和原数组解除绑定

+ 当切片和数组作为参数在函数（func）中传递时，数组传递的是值，而切片传递的是指针!!!

+ 数组（Array）和切片（Slice）的区别:<https://www.jianshu.com/p/10d23e9ffc36>


### 13.3 尽量使用cap参数创建切片

+ 使用带cap参数创建的切片进行append操作的平均性能（9250ns）是不带cap参数的切片（36 484ns）的4倍左右，并且每操作平均仅需一次内存分配。

## 第14条 了解map实现原理并高效使用

### 14.1 什么是map

+ map对value的类型没有限制，但是对key的类型有严格要求：key的类型应该严格定义了作为“==”和“!=”两个操作符的操作数时的行为，因此函数、map、切片不能作为map的key类型。

+ 和切片一样，map也是引用类型，将map类型变量作为函数参数传入不会有很大的性能损耗，并且在函数内部对map变量的修改在函数外部也是可见的

### 14.2 map的基本操作

+ 所谓查找就是判断某个key是否存在于某个map中。我们可以使用“comma ok”惯用法来进行查找：
```go
_, ok := m["key"] if !ok {    
    // "key"不在map中
}
```  
+ Go语言的一个最佳实践是总是使用“comma ok”惯用法读取map中的值。

+ 如果你需要一个稳定的遍历次序，那么一个比较通用的做法是使用另一种数据结构来按需要的次序保存key，比如切片
```go
func main() {    
    var sl []int    
    m := map[int]int{        
        1: 11,        
        2: 12,        
        3: 13,    
    }    
    for k, _ := range m {        
        sl = append(sl, k) // 将元素按初始次序保存在切片中    
    }    
    for i := 0; i < 3; i++ {        
        doIteration(sl, m)    
    }
}
```
### 14.3 map的内部实现
### 14.4 尽量使用cap参数创建map

+ 使用cap参数的map实例的平均写性能是不使用cap参数的2倍。

### 小结
和切片一样，map是Go语言提供的重要数据类型，也是Gopher日常编码中最常使用的类型之一。通过本条的学习我们掌握了map的基本操作和运行时实现原理，并且我们在日常使用map的场合要把握住下面几个要点：不要依赖map的元素遍历顺序；map不是线程安全的，不支持并发写；不要尝试获取map中元素（value）的地址；尽量使用cap参数创建map，以提升map平均访问性能，减少频繁扩容带来的不必要损耗。

## 第15条 了解string实现原理并高效使用

### 15.2 字符串的内部表示

### 15.3 字符串的高效构造
+  做了预初始化的strings.Builder连接构建字符串效率最高；
•  带有预初始化的bytes.Buffer和strings.Join这两种方法效率十分接近，分列二三位；
•  未做预初始化的strings.Builder、bytes.Buffer和操作符连接在第三档次；
•  fmt.Sprintf性能最差，排在末尾。

+ 可以得出一些结论：
•  在能预估出最终字符串长度的情况下，使用预初始化的strings.Builder连接构建字符串效率最高；
•  strings.Join连接构建字符串的平均性能最稳定，如果输入的多个字符串是以[]string承载的，那么strings.Join也是不错的选择；
•  使用操作符连接的方式最直观、最自然，在编译器知晓欲连接的字符串个数的情况下，使用此种方式可以得到编译器的优化处理；
•  fmt.Sprintf虽然效率不高，但也不是一无是处，如果是由多种不同类型变量来构建特定格式的字符串，那么这种方式还是最适合的。

+ 在Go运行时层面，字符串与rune slice、byte slice相互转换对应的函数如下：
```go
// $GOROOT/src/runtime/string.go 
slicebytetostring: 
[]byte -> string
slicerunetostring: 
[]rune -> string
stringtoslicebyte: 
string -> []byte
stringtoslicerune: 
string -> []rune
```

+ 以byte slice为例，看看slicebytetostring和stringtoslicebyte的实现：
```go
// $GOROOT/src/runtime/string.go 
const tmpStringBufSize = 32 
type tmpBuf [tmpStringBufSize]byte

func stringtoslicebyte(buf *tmpBuf, s string) []byte {
	var b []byte
	if buf != nil && len(s) <= len(buf) {
        //重置为0值
		*buf = tmpBuf{}
		b = buf[:len(s)]
	} else {
		b = rawbyteslice(len(s))
	}
	copy(b, s)
	return b
}
```
+ 想要更高效地进行转换，唯一的方法就是减少甚至避免额外的内存分配操作

+ slice类型是不可比较的，而string类型是可比较的，因此在日常Go编码中，我们会经常遇到将slice临时转换为string的情况。

+ Go编译器为这样的场景提供了优化。在运行时中有一个名为slicebytetostringtmp的函数就是协助实现这一优化的：
```go
// $GOROOT/src/runtime/string.go 
func slicebytetostringtmp(b []byte) 
```

+ 该函数的“秘诀”就在于不为string新开辟一块内存，而是直接使用slice的底层存储。当然使用这个函数的前提是：在原slice被修改后，这个string不能再被使用了。因此这样的优化是针对以下几个特定场景的。

+ Go语言还在标准库中提供了strings和strconv包，可以辅助Gopher对string类型数据进行更多高级操作。

## 第16条 理解Go语言的包导入

### 16.1 Go程序构建过程

+ Go程序的构建简单来讲也是由编译（compile）和链接（link）两个阶段组成的。
+ go build命令传入-x -v命令行选项来输出详细的构建日志信息

+ 所谓的使用第三方包源码，实际上是链接了以该最新包源码编译的、存放在临时目录下的包的.a文件而已。
+ 默认情况下对于标准库中的包，编译器直接链接的是$GOROOT/pkg/darwin_amd64下的.a文件。

### 16.2 究竟是路径名还是包名

+ 编译器在编译过程中必然要使用的是编译单元（一个包）所依赖的包的源码。而编译器要找到依赖包的源码文件，就需要知道依赖包的源码路径。这个路径由两部分组成：基础搜索路径和包导入路径。

+ 关于包导入，Go语言还有一个惯用法：当包名与包导入路径中的最后一个目录名不同时，最好用下面的语法将包名显式放入包导入语句。

### 16.3 包名冲突问题

+ 用为包导入路径下的包显式指定包名的方法

### 小结
在本条中，我们通过实验进一步理解了Go语言的包导入，Gopher应牢记以下几个结论：
•  Go编译器在编译过程中必然要使用的是编译单元（一个包）所依赖的包的源码；
•  Go源码文件头部的包导入语句中import后面的部分是一个路径，路径的最后一个分段是目录名，而不是包名；
•  Go编译器的包源码搜索路径由基本搜索路径和包导入路径组成，两者结合在一起后，编译器便可确定一个包的所有依赖包的源码路径的集合，这个集合构成了Go编译器的源码搜索路径空间；
•  同一源码文件的依赖包在同一源码搜索路径空间下的包名冲突问题可以由显式指定包名的方式解决。

## 第17条 理解Go语言表达式的求值顺序

### 17.1 包级别变量声明语句中的表达式求值顺序

+ 在一个Go包内部，包级别变量声明语句的表达式求值顺序是由初始化依赖（initialization dependencies）规则决定的。

### 17.2 普通求值顺序

+ Go规定表达式操作数中的所有函数、方法以及channel操作按照从左到右的次序进行求值。

### 17.3 赋值语句的求值

### 小结

•  包级别变量声明语句中的表达式求值顺序由变量的声明顺序和初始化依赖关系决定，并且包级变量表达式求值顺序优先级最高。
•  表达式操作数中的函数、方法及channel操作按普通求值顺序，即从左到右的次序进行求值。
•  赋值语句求值分为两个阶段：先按照普通求值规则对等号左边的下标表达式、指针解引用表达式和等号右边的表达式中的操作数进行求值，然后按从左到右的顺序对变量进行赋值。
•  重点关注switch-case和select-case语句中的表达式“惰性求值”规则。

## 第18条 理解Go语言代码块与作用域


### 18.1 Go代码块与作用域简介
### 18.1 Go代码块与作用域简介

+ 代码块是代码执行流流转的基本单元，代码执行流总是从一个代码块跳到另一个代码块。

+ Go语言中有两类代码块，一类是我们在代码中直观可见的由一堆大括号包裹的显式代码块，比如函数的函数体、for循环的循环体、if语句的某个分支等：

+ 另一类则是没有大括号包裹的隐式代码块。Go规范定义了如下几种隐式代码块。
•  宇宙（Universe）代码块：所有Go源码都在该隐式代码块中，就相当于所有Go代码的最外层都存在一对大括号。
•  包代码块：每个包都有一个包代码块，其中放置着该包的所有Go源码。
•  文件代码块：每个文件都有一个文件代码块，其中包含着该文件中的所有Go源码。
•  每个if、for和switch语句均被视为位于其自己的隐式代码块中。
•  switch或select语句中的每个子句都被视为一个隐式代码块。


## 第19条 了解Go语言控制语句惯用法及使用注意事项
+  switch的case语句执行完毕后，默认不会像C语言那样继续执行下一个case中的语句，除非显式使用fallthrough关键字，这“填补”了C语言中每个case语句都要以break收尾的“坑”；

+  增加针对channel通信的switch-case语句——select-case。

### 19.1 使用if控制语句时应遵循“快乐路径”原则

+  当出现错误时，快速返回

+  尝试将“正常逻辑”提取出来，放到“快乐路径”中；

### 19.2 for range的避“坑”指南

1. 小心迭代变量的重用
2. 注意参与迭代的是range表达式的副本 !!!
for range语句中，range后面接受的表达式的类型可以是数组、指向数组的指针、切片、字符串、map和channel（至少需具有读权限）。

```go
 // chapter3/sources/control_structure_idiom_2.go... 
 func arrayRangeExpression() {    
    var a = [5]int{1, 2, 3, 4, 5}    
    var r [5]int    
    fmt.Println("arrayRangeExpression result:")    
    fmt.Println("a = ", a)    
    for i, v := range a {        
        if i == 0 {            
            a[1] = 12            
            a[2] = 13        
        }        
        r[i] = v    
    }    
    fmt.Println("r = ", r)    
    fmt.Println("a = ", a)
}
```
用数组指针作为range表达式
```go
 for i, v := range &a {        
    if i == 0 {            
        a[1] = 12            
        a[2] = 13        
    }        
    r[i] = v    
}
```

在Go中，大多数应用数组的场景都可以用切片替代
``` go
 for i, v := range a[:] {        
    if i == 0 {            
        a[1] = 12            
        a[2] = 13        
 }
```
+ 用切片也能满足预期要求

+ 切片副本的结构体中的*T依旧指向原切片对应的底层数组，因此对切片副本的修改也都会反映到底层数组a上

+ 切片与数组还有一个不同点，就是其len在运行时可以被改变，而数组的长度可认为是一个常量，不可改变。

+ range表达式的复制行为还会带来一些性能上的消耗，尤其是当range表达式的类型为数组时，range需要复制整个数组；而当range表达式类型为数组指针或切片时，这个消耗将小得多，因为仅仅需要复制一个指针或一个切片的内部表示（一个结构体）即可。

3. 其他range表达式类型的使用注意事项

for range对于string来说，每次循环的单位是一个rune，而不是一个byte
for range对map副本的操作即对源map的操作。
channel的指针副本也指向原channel。
当channel作为range表达式类型时，for range最终以阻塞读的方式阻塞在channel表达式上，即便是带缓冲的channel亦是如此：当channel中无数据时，for range也会阻塞在channel上，直到channel关闭

如果使用一个nil channel作为range表达式
for range将永远阻塞在这个nil channel上，直到Go运行时发现程序陷入deadlock状态，并抛出panic

### 19.3 break跳到哪里去了
+ Go break语法的一个“小坑”。和大家习惯的C家族语言中的break不同，Go语言规范中明确规定break语句（不接label的情况下）结束执行并跳出的是同一函数内break语句所在的最内层的for、switch或select的执行。

+ 要修正这一问题，可以利用Go语言为for提供的一项高级能力：break [label]。
```go
 // chapter3/sources/control_structure_idiom_7.go 
 func main() {    
    exit := make(chan interface{})    
    go func() {    
        loop:        
        for {            
            select {            
                case <-time.After(time.Second):                
                    fmt.Println("tick")            
                case <-exit:                
                    fmt.Println("exiting...")                
                    break loop            
            }        
         }        
        fmt.Println("exit!")    
    }()    
    time.Sleep(3 * time.Second)    
    exit <- struct{}{}    // 等待子goroutine退出    
    time.Sleep(3 * time.Second)
}

``` 

```go
 outerLoop:    for i := 0; i < n; i++ {        
    // ...        
    for j := 0; j < m; j++ {            
        // 当不满足某些条件时，直接终止最外层循环的执行           
        break outerLoop            // 当满足某些条件时，直接跳出内层循环，回到外层循环继续执行            
        continue outerLoop        
        }    
  }
```

### 19.4 尽量用case表达式列表替代fallthrough

+ 通过case接表达式列表的方式要比使用fallthrough更加简洁和易读。


### 小结

•  使用if语句时遵循“快乐路径”原则；
•  小心for range的循环变量重用，明确真实参与循环的是range表达式的副本；
•  明确break和continue执行后的真实“目的地”；
•  使用fallthrough关键字前，考虑能否用更简洁、清晰的case表达式列表替代。

# 第四部分 函数与方法

## 第20条 在init函数中检查包级变量的初始状态

### 20.1 认识init函数

+ 一个Go包可以拥有多个init函数，每个组成Go包的Go源文件中可以定义多个init函数。在初始化Go包时，Go运行时会按照一定的次序逐一调用该包的init函数。Go运行时不会并发调用init函数，它会等待一个init函数执行完毕并返回后再执行下一个init函数，且每个init函数在整个Go程序生命周期内仅会被执行一次。因此，init函数极其适合做一些包级数据的初始化及初始状态的检查工作。

+ 一般来说，先被传递给Go编译器的源文件中的init函数先被执行，同一个源文件中的多个init函数按声明顺序依次执行。但Go语言的惯例告诉我们：不要依赖init函数的执行次序。

### 20.2 程序初始化顺序

+ init函数为何适合做包级数据的初始化及初始状态检查工作呢？除了init函数是顺序执行并仅被执行一次之外，Go程序初始化顺序也给init函数提供了胜任该工作的前提条件。

### 20.3 使用init函数检查包级变量的初始状态

1. 重置包级变量值
2. 对包级变量进行初始化，保证其后续可用
3. init函数中的注册模式

```go
// github.com/lib/pq/conn.go...
func init() {    
    sql.Register("postgres", &Driver{})
}
```
4. init函数中检查失败的处理方法
快速失败是最佳选择。我们一般建议直接调用panic或者通过log.Fatal等函数记录异常日志，然后让程序快速退出。

### 小结
要深入理解init函数，记住本条介绍的几个要点即可。
•  init函数的几个特点：运行时调用、顺序、仅执行一次。
•  Go程序的初始化顺序。
•  init函数是包出厂前的唯一“质检员”。

## 第21条 让自己习惯于函数是“一等公民”

+ Go语言中没有那些典型的面向对象语言的语法，比如类、继承、对象等。Go语言中的方法（method）本质上是函数的一个变种。因此，在Go语言中，函数是唯一一种基于特定输入、实现特定任务并可反馈任务执行结果的代码块。本质上，我们可以说Go程序就是一组函数的集合。

### 21.1 什么是“一等公民”

+ 如果一门编程语言对某种语言元素的创建和使用没有限制，我们可以像对待值（value）一样对待这种语法元素，那么我们就称这种语法元素是这门编程语言的“一等公民”。拥有“一等公民”待遇的语法元素可以存储在变量中，可以作为参数传递给函数，可以在函数内部创建并可以作为返回值从函数返回。在动态类型语言中，语言运行时还支持对“一等公民”类型的检查。

### 21.2 函数作为“一等公民”的特殊运用

+ 为了充分理解这种显式类型转换的技巧，我们再来看一个简化后的例子：
```go
// chapter4/sources/function_as_first_class_citizen_3.go 
type BinaryAdder interface {    
    Add(int, int) int
}
type MyAdderFunc func(int, int) int
func (f MyAdderFunc) Add(x, y int) int {    
    return f(x, y)
}
func MyAdd(x, y int) int {    
    return x + y
}
func main() {    
    var i BinaryAdder = MyAdderFunc(MyAdd)    
    fmt.Println(i.Add(5, 6))
}
```

2. 函数式编程
（1）柯里化函数
```go
func times(x, y int) int {    
    return x * y
}
func partialTimes(x int) func(int) int {    
    return func(y int) int {        
        return times(x, y)    
    }
}
func main() {    
    timesTwo := partialTimes(2)    
    timesThree := partialTimes(3)    
    timesFour := partialTimes(4)    
    fmt.Println(timesTwo(5))    
    fmt.Println(timesThree(5))    
    fmt.Println(timesFour(5))
}
```
闭包是将函数内部和函数外部连接起来的桥梁

（2）函子
+ 什么是函子呢？具体来说，函子需要满足两个条件：
•  函子本身是一个容器类型，以Go语言为例，这个容器可以是切片、map甚至channel；
•  该容器类型需要实现一个方法，该方法接受一个函数类型参数，并在容器的每个元素上应用那个函数，得到一个新函子，原函子容器内部的元素值不受影响。

```go
// chapter4/sources/function_as_first_class_citizen_5.go 
type IntSliceFunctor interface {    
    Fmap(fn func(int) int) IntSliceFunctor
}
type intSliceFunctorImpl struct {    
    ints []int
}
func (isf intSliceFunctorImpl) Fmap(fn func(int) int) IntSliceFunctor {    
    newInts := make([]int, len(isf.ints))    
    for i, elt := range isf.ints {        
        retInt := fn(elt)        
        newInts[i] = retInt    
    }    
    return intSliceFunctorImpl{ints: newInts}
}

func NewIntSliceFunctor(slice []int) IntSliceFunctor {    
    return intSliceFunctorImpl{ints: slice}
}
func main() {    
    // 原切片    
    intSlice := []int{1, 2, 3, 4}    
    fmt.Printf("init a functor from int slice: %#v\n", intSlice)    
    f := NewIntSliceFunctor(intSlice)    
    fmt.Printf("original functor: %+v\n", f)    
    mapperFunc1 := func(i int) int {        
        return i + 10    
    }    
    mapped1 := f.Fmap(mapperFunc1)    
    fmt.Printf("mapped functor1: %+v\n", mapped1)    
    mapperFunc2 := func(i int) int {        
        return i * 3    
    }    
    mapped2 := mapped1.Fmap(mapperFunc2)    
    fmt.Printf("mapped functor2: %+v\n", mapped2)    
    fmt.Printf("original functor: %+v\n", f) 
    // 原函子没有改变    
    fmt.Printf("composite functor: %+v\n", f.Fmap(mapperFunc1).Fmap(mapperFunc2))
}
```

（3）延续传递式

函数式编程有一种被称为延续传递式（Continuation-passing Style，CPS）的编程风格可以充分运用函数作为“一等公民”的特质。

```go
// chapter4/sources/function_as_first_class_citizen_8.go 
func Max(n int, m int, f func(y int)) {    
    if n > m {        
        f(n)    
    } else {        
        f(m)    
    }
}
func main() {    
    Max(5, 6, func(y int) { 
        fmt.Printf("%d\n", y) 
        }
    )
 }
```

这种CPS风格虽然利用了函数作为“一等公民”的特质，但是其代码理解起来颇为困难，这种风格真的好吗？朋友们的担心是有道理的。这里对CPS风格的讲解其实是一个反例，

目的就是告诉大家，尽管作为“一等公民”的函数给Go带来了强大的表达能力，但是如果选择了不适合的风格或者为了函数式而进行函数式编程，那么就会出现代码难于理解且代码执行效率不高的情况（CPS需要语言支持尾递归优化，但Go目前并不支持）。

## 第22条 使用defer让函数更简洁、更健壮

### 22.1 defer的运作机制
### 22.2 defer的常见用法

1. 拦截panic
```go 
func foo() {    
    defer func() {        
        if e := recover(); e != nil {            
            fmt.Println("recovered from a panic")        
        }    
    }()    
    bar()
}
func main() {    
    foo()    
    fmt.Println("main exit normally")
}
```
+ deferred函数虽然可以拦截绝大部分的panic，但无法拦截并恢复一些运行时之外的致命问题

2. 修改函数的具名返回值
```go
// chapter4/sources/deferred_func_5.go 
func foo(a, b int) (x, y int) {    
    defer func() {        
        x = x * 5        
        y = y * 10    
    }()    
    x = a + 5    
    y = b + 6    
    return
}
func main() {    
    x, y := foo(1, 2)    
    fmt.Println("x=", x, "y=", y)
}
```

3. 输出调试信息
4. 还原变量旧值

### 22.3 关于defer的几个关键问题

1. 明确哪些函数可以作为deferred函数

+ 对于自定义的函数或方法，defer可以给予无条件的支持，但是对于有返回值的自定义函数或方法，返回值会在deferred函数被调度执行的时候被自动丢弃。

+ Go语言中除了有自定义的函数或方法，还有内置函数。下面是Go语言内置函数的完整列表：
append cap close complex copy delete imag lenmake new panic print println real recover

+ Go编译器给出了一组错误提示！从中我们看到，append、cap、len、make、new等内置函数是不可以直接作为deferred函数的，而close、copy、delete、print、recover等可以。

+ 对于那些不能直接作为deferred函数的内置函数，我们可以使用一个包裹它的匿名函数来间接满足要求。以append为例：
`defer func() {    _ = append(sl, 11)}()`
但这么做有什么实际意义需要开发者自己把握。

+ defer关键字后面的表达式是在将deferred函数注册到deferred函数栈的时候进行求值的。
```go
 // chapter4/sources/deferred_func_7.go 
 func foo1() {    
    for i := 0; i <= 3; i++ {        
        defer fmt.Println(i)    
    }
}
// 3 2 0 1
func foo2() {    
    for i := 0; i <= 3; i++ {        
        defer func(n int) {               
            fmt.Println(n)        
        }(i)    
    }
}
//3 2 1 0
func foo3() {    
    for i := 0; i <= 3; i++ {        
        defer func() {            
            fmt.Println(i)        
        }()    
    }
}
//4 4 4 4
func main() {    
    fmt.Println("foo1 result:")    
    foo1()    
    fmt.Println("\nfoo2 result:")    
    foo2()    
    fmt.Println("\nfoo3 result:")    
    foo3()
}
```
我们逐一分析foo1、foo2和foo3中defer关键字后的表达式的求值时机：
在foo1中，defer后面直接接的是fmt.Println函数，每当defer将fmt.Println注册到deferred函数栈的时候，都会对Println后面的参数进行求值。根据上述代码逻辑，依次压入deferred函数栈的函数是：
fmt.Println(0)fmt.Println(1)fmt.Println(2)fmt.Println(3)
因此，在foo1返回后，deferred函数被调度执行时，上述压入栈的deferred函数将以LIFO次序出栈执行，因此输出的结果为：
3210
在foo2中，defer后面接的是一个带有一个参数的匿名函数。每当defer将匿名函数注册到deferred函数栈的时候，都会对该匿名函数的参数进行求值。根据上述代码逻辑，依次压入deferred函数栈的函数是：
func(0)func(1)func(2)func(3)

>> 因此，在foo2返回后，deferred函数被调度执行时，上述压入栈的deferred函数将以LIFO次序出栈执行，因此输出的结果为：
3210
在foo3中，defer后面接的是一个不带参数的匿名函数。根据上述代码逻辑，依次压入deferred函数栈的函数是：
func()func()func()func()
因此，在foo3返回后，deferred函数被调度执行时，上述压入栈的deferred函数将以LIFO次序出栈执行。匿名函数以闭包的方式访问外围函数的变量i，并通过Println输出i的值，此时i的值为4，因此foo3的输出结果为：
4444

## 第23条 理解方法的本质以选择正确的receiver类型

+ 方法是类型T的一个方法，我们可以通过类型T或*T的实例调用该方法

+ 方法定义要与类型定义放在同一个包内。由此我们可以推出：不能为原生类型（如int、float64、map等）添加方法，只能为自定义类型定义方法

+ 同理，可以推出：不能横跨Go包为其他包内的自定义类型定义方法。

+ receiver参数的基类型本身不能是指针类型或接口类型

### 23.1 方法的本质

```go
type T struct {    
    a int
}
func (t T) Get() int {    
    return t.a
}
func (t *T) Set(a int) int {    
    t.a = a    
    return t.a
}
```
>> 转换为下面的普通函数：
```go
func Get(t T) int {    
    return t.a
}
func Set(t *T, a int) int {    
    t.a = a    
    return t.a
}
```

Go方法的一般使用方式如下：
```go
var t T
t.Get()
t.Set(1)
```
我们可以用如下方式等价替换上面的方法调用：
```go
var t T
T.Get(t)
(*T).Set(&t, 1)
```

+ Go方法的本质：一个以方法所绑定类型实例为第一个参数的普通函数。

+ Go方法自身的类型就是一个普通函数，我们甚至可以将其作为右值赋值给函数类型的变量：
```go
var t T
f1 := (*T).Set 
// f1的类型，也是T类型Set方法的原型：func (t *T, int)int
f2 := T.Get    
// f2的类型，也是T类型Get方法的原型：func (t T)int
f1(&t, 3)
fmt.Println(f2(t))
```

### 23.2 选择正确的receiver类型


+ 方法和函数的等价变换公式：
func (t T) M1() <=> M1(t T)
func (t *T) M2() <=> M2(t *T)

```go
// chapter4/sources/method_nature_1.go 
type T struct {    
    a int
}
func (t T) M1() {    
    t.a = 10
}
func (t *T) M2() {    
    t.a = 11
}
func main() {    
    var t T 
    // t.a = 0    
    println(t.a)    
    t.M1()    
    println(t.a)    
    t.M2()    
    println(t.a)
}
```
运行该程序：
`$ go run method_nature_1.go  0 0 11`

+ 改实例值一定要用指针的方法!!!


•  如果要对类型实例进行修改，那么为receiver选择*T类型。
•  如果没有对类型实例修改的需求，那么为receiver选择T类型或*T类型均可；但考虑到Go方法调用时，receiver是以值复制的形式传入方法中的，如果类型的size较大，以值形式传入会导致较大损耗，这时选择*T作为receiver类型会更好些。

### 小结

+ Go语言未提供对经典面向对象机制的语法支持，但实现了类型的方法，方法与类型间通过方法名左侧的receiver建立关联。为类型的方法选择合适的receiver类型是Gopher为类型定义方法的重要环节。

•  Go方法的本质：一个以方法所绑定类型实例为第一个参数的普通函数。
•  Go语法糖使得我们在通过类型实例调用类型方法时无须考虑实例类型与receiver参数类型是否一致，编译器会为我们做自动转换。
•  在选择receiver参数类型时要看是否要对类型实例进行修改。如有修改需求，则选择*T；如无修改需求，T类型receiver传值的性能损耗也是考量因素之一。

## 第24条 方法集合决定接口实现

### 24.1　方法集合

+ Go语言的一个创新是，自定义类型与接口之间的实现关系是松耦合的：如果某个自定义类型T的方法集合是某个接口类型的方法集合的超集，那么就说类型T实现了该接口，并且类型T的变量可以被赋值给该接口类型的变量，即我们说的方法集合决定接口实现。

+ 这里我们实现了一个工具函数，它可以方便地输出一个自定义类型或接口类型的方法集合。
```go
// chapter4/sources/method_set_utils.go

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()
	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}
	fmt.Printf("%s's method set:\n", elemTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
```

+ 这符合Go语言规范：对于非接口类型的自定义类型T，其方法集合由所有receiver为T类型的方法组成；而类型`*T`的方法集合则包含所有receiver为T和`*T`类型的方法。

+ 到这里，我们完全明确了为receiver选择类型时需要考虑的第三点因素：是否支持将T类型实例赋值给某个接口类型变量。如果需要支持，我们就要实现receiver为T类型的接口类型方法集合中的所有方法。

### 24.2 类型嵌入与方法集合

+ Go的设计哲学之一是偏好组合，Go支持用组合的思想来实现一些面向对象领域经典的机制，比如继承。而具体的方式就是利用类型嵌入（type embedding）。

+ 与接口类型和结构体类型相关的类型嵌入有三种组合：在接口类型中嵌入接口类型、在结构体类型中嵌入接口类型及在结构体类型中嵌入结构体类型。

1. 在接口类型中嵌入接口类型

>> io包中的ReadWriter、ReadWriteCloser等接口类型就是通过嵌入Reader、Writer或Closer三个基本接口类型形成的

2. 在结构体类型中嵌入接口类型
结构体类型在嵌入某接口类型的同时，也实现了这个接口。这一特性在单元测试中尤为有用，尤其是在应对下面这样的场景时：
```go
// chapter4/sources/method_set_9.go 
package employee
type Result struct {    
    Count int
}
func (r Result) Int() int { 
    return r.Count 
}
type Rows []struct{}
type Stmt interface {    
    Close() error    
    NumInput() int    
    Exec(stmt string, args ...string) (Result, error)    
    Query(args []string) (Rows, error)
}
// 返回男性员工总数
func MaleCount(s Stmt) (int, error) {    
    result, err := s.Exec("select count(*) from employee_tab where gender=?", "1")    
    if err != nil {        
        return 0, err    
    }    
    return result.Int(), nil
}
```
在这个例子中有一个employee包，该包中的MaleCount方法通过传入的Stmt接口的实现从数据库中获取男性员工的数量。

>> 现在我们要对MaleCount方法编写单元测试代码。对于这种依赖外部数据库操作的方法，惯例是使用伪对象（fake object）来冒充真实的Stmt接口实现。不过现在有一个问题是，Stmt接口类型的方法集合中有4个方法，如果针对每个测试用例所用的伪对象都实现这4个方法，那么这个工作量有点大，而我们需要的仅仅是Exec这一个方法。如何快速建立伪对象呢？在结构体类型中嵌入接口类型便可以帮助我们：
```go
// chapter4/sources/method_set_9_test.go 
package employee
import "testing"
type fakeStmtForMaleCount struct {    
    Stmt
}
func (fakeStmtForMaleCount) Exec(stmt string, args ...string) (Result, error) {    
    return Result{Count: 5}, nil
}
func TestEmployeeMaleCount(t *testing.T) {    
    f := fakeStmtForMaleCount{}    
    c, _ := MaleCount(f)    
    if c != 5 {        
        t.Errorf("want: %d, actual: %d", 5, c)        
        return    
    }
}
```
我们为TestEmployeeMaleCount测试用例建立了一个fakeStmtForMaleCount的伪对象，在该结构体类型中嵌入Stmt接口类型，这样fakeStmtForMaleCount就实现了Stmt接口，我们达到了快速建立伪对象的目的。之后，我们仅需为fakeStmtForMaleCount实现MaleCount所需的Exec方法即可。

3. 在结构体类型中嵌入结构体类型

+ 在结构体类型中嵌入结构体类型为Gopher提供了一种实现“继承”的手段，外部的结构体类型T可以“继承”嵌入的结构体类型的所有方法的实现，并且无论是T类型的变量实例还是*T类型变量实例，都可以调用所有“继承”的方法。

### 24.3 defined类型的方法集合

+ Go语言支持基于已有的类型创建新类型，比如：
```go
type MyInterface I
type Mystruct T
```
已有的类型（比如上面的I、T）被称为underlying类型，而

+ Go对于分别基于接口类型和自定义非接口类型创建的defined类型给出了不一致的结果：
•  基于接口类型创建的defined类型与原接口类型的方法集合是一致的，如上面的Interface和Interface1；
•  而基于自定义非接口类型创建的defined类型则并没有“继承”原类型的方法集合，新的defined类型的方法集合是空的。!!!!!

+ 方法集合决定接口实现。基于自定义非接口类型的defined类型的方法集合为空，这决定了即便原类型实现了某些接口，基于其创建的defined类型也没有“继承”这一隐式关联。新defined类型要想实现那些接口，仍需重新实现接口的所有方法。

### 24.4 类型别名的方法集合

+ 类型别名与原类型几乎是等价的。Go预定义标识符rune、byte就是通过类型别名语法定义的：

```go
// $GOROOT/src/builtin/builtin.go 
type byte = uint8 
type rune = int32
```
+ 类型别名与原类型拥有完全相同的方法集合，无论原类型是接口类型还是非接口类型。

## 第25条 了解变长参数函数的妙用
### 25.1 什么是变长参数函数
### 25.2 模拟函数重载
+ Go语言不允许在同一个作用域下定义名字相同但函数原型不同的函数，如果定义这样的函数

+ 但Go语言并不支持函数重载，Go语言官方常见问答[1]中给出的不支持的理由如下：
其他语言的经验告诉我们，使用具有相同名称但函数签名不同的多种方法有时会很有用，但在实践中也可能会造成混淆和脆弱性。在Go的类型系统中，仅按名称进行匹配并要求类型一致是一个主要的简化决策。

### 25.3 模拟实现函数的可选参数与默认参数

### 25.4 实现功能选项模式

+ 版本3：使用功能选项模式

+ Go语言之父Rob Pike早在2014年就在其博文“自引用函数与选项设计”[2]中论述了一种被后人称为“功能选项”（functional option）的模式，这种模式应该是目前进行功能选项设计的最佳实践。

```go
// chapter4/sources/variadic_function_9.go

type FinishedHouse struct {    
	style                  int    // 0: Chinese; 1: American; 2: European   
	 centralAirConditioning bool   // true或false    
	 floorMaterial          string  // "ground-tile"或"wood"    
	 wallMaterial           string // "latex"或"paper"或"diatom-mud"
	 }
type Option func(*FinishedHouse)
func NewFinishedHouse(options ...Option) *FinishedHouse {   
	 h := &FinishedHouse{        // default options        
		style:                  0,        
		centralAirConditioning: true,       
		 floorMaterial:          "wood",        
		 wallMaterial:           "paper",    
		 }    
		 for _, option := range options {        
			option(h)    
			}    
	return h
}
func WithStyle(style int) Option {    
	return func(h *FinishedHouse) {        
		h.style = style    
		}
	}
func WithFloorMaterial(material string) Option {    
	return func(h *FinishedHouse) {        
		h.floorMaterial = material    
	}
}
func WithWallMaterial(material string) Option {    
	return func(h *FinishedHouse) {        
		h.wallMaterial = material    
		}
	}
func WithCentralAirConditioning(centralAirConditioning bool) Option {    
	return func(h *FinishedHouse) {        
		h.centralAirConditioning = centralAirConditioning    
	}
}
func main() {    
	fmt.Printf("%+v\n", NewFinishedHouse()) // 使用默认选项    
	fmt.Printf("%+v\n", NewFinishedHouse(WithStyle(1),        
	WithFloorMaterial("ground-tile"),        
	WithCentralAirConditioning(false)))
}
```

+ 在设计和实现类似NewFinishedHouse这样带有配置选项的函数或方法时，功能选项模式让我们可以收获如下好处：
•  更漂亮的、不随时间变化的公共API；
•  参数可读性更好；
•  配置选项高度可扩展；
•  提供使用默认选项的最简单方式；
•  使用更安全（不会像版本2那样在创建函数被调用后，调用者仍然可以修改options）。

# 第五部分　接口

## 第26条 了解接口类型变量的内部表示
+ 接口类型变量在程序运行时可以被赋值为不同的动态类型变量，从而支持运行时多态。

### 26.1　nil error值!= nil
```go
 // chapter5/sources/interface-internal-1.go 
type MyError struct {    
    error
}
var ErrBad = MyError{    
    error: errors.New("bad error"),
}
func bad() bool {    
    return false
}
func returnsError() error {    
    var p *MyError = nil    
    if bad() {        
        p = &ErrBad    
    }    
    return p
}
func main() {    
    e := returnsError()    
    if e != nil {        
        fmt.Printf("error: %+v\n", e)        
        return    
    }    
    fmt.Println("ok")
}
``` 

+ 初学者的思路大致是这样的：p为nil，returnsError返回p，那么main函数中的e就等于nil，于是程序输出ok后退出。但真实的运行结果是什么样的呢？我们来看一下：
`$go run interface-internal-1.go error: <nil>`

+ 疑惑出现了：明明returnsError函数返回的p值为nil，为何却满足了if e != nil的条件进入错误处理分支呢？要想弄清楚这个问题，非了解接口类型变量的内部表示不可。


### 26.2 接口类型变量的内部表示

+ 接口类型“动静兼备”的特性决定了它的变量的内部表示绝不像静态类型（如int、float64）变量那样简单。我们可以在$GOROOT/src/runtime/runtime2.go中找到接口类型变量在运行时的表示：
```go
// $GOROOT/src/runtime/runtime2.go 
type iface struct {    
    tab  *itab    
    data unsafe.Pointer
}
type eface struct {    
    _type *_type    
    data  unsafe.Pointer
}
```

+ 我们看到在运行时层面，接口类型变量有两种内部表示——eface和iface，这两种表示分别用于不同接口类型的变量。
•  eface：用于表示没有方法的空接口（empty interface）类型变量，即interface{}类型的变量。
•  iface：用于表示其余拥有方法的接口（interface）类型变量。
这两种结构的共同点是都有两个指针字段，并且第二个指针字段的功用相同，都指向当前赋值给该接口类型变量的动态类型变量的值。

### 26.3 输出接口类型变量内部表示的详细信息
###  26.4 接口类型的装箱原理

+ 装箱是一个有性能损耗的操作，因此Go在不断对装箱操作进行优化，包括对常见类型（如整型、字符串、切片等）提供一系列快速转换函数：
```go
// $GOROOT/src/cmd/compile/internal/gc/builtin/runtime.go
// 实现在 $GOROOT/src/runtime/iface.go中
func convT16(val any) unsafe.Pointer     // val必须是一个uint-16相关类型的参数
func convT32(val any) unsafe.Pointer     // val必须是一个unit-32相关类型的参数
func convT64(val any) unsafe.Pointer     // val必须是一个unit-64相关类型的参数
func convTstring(val any) unsafe.Pointer // val必须是一个字符串类型的参数
func convTslice(val any) unsafe.Pointer  // val必须是一个切片类型的参数
```

### 小结
本条从Go FAQ中的一个例子出发，解释了nil接口变量不等于nil的原因，并和大家一起深入探究了Go接口类型的两种内部表示，了解了接口类型变量的装箱过程。
本条要点：
•  接口类型变量在运行时表示为eface和iface，eface用于表示空接口类型变量，iface用于表示非空接口类型变量；
•  当且仅当两个接口类型变量的类型信息（eface._type/iface.tab._type）相同，且数据指针（eface.data/iface.data）所指数据相同时，两个接口类型才是相等的；
•  通过println可以输出接口类型变量的两部分指针变量的值；
•  可通过复制runtime包eface和iface相关类型源码，自定义输出eface/iface详尽信息的函数；
•  接口类型变量的装箱操作由Go编译器和运行时共同完成。

## 第27条 尽量定义小接口

### 27.1 Go推荐定义小接口

### 27.2 小接口的优势

1. 接口越小，抽象程度越高，被接纳度越高
2. 易于实现和测试
3. 契约职责单一，易于复用组合

### 27.3 定义小接口可以遵循的一些点

1. 抽象出接口

初期不要在意接口的大小，因为对问题域的理解是循序渐进的，期望在第一版代码中直接定义出小接口可能并不现实。标准库中的io.Reader和io.Writer也不是在Go刚诞生时就有的，而是在发现对网络、文件、其他字节数据处理的实现十分相似之后才抽象出来的。此外，越偏向业务层，抽象难度越高

2. 将大接口拆分为小接口

3. 接口的单一契约职责


### 第28条 尽量避免使用空接口作为函数参数类型
+ 最后，总结一下，本条的主要内容如下：
•  仅在处理未知类型数据时使用空接口类型；
•  在其他情况下，尽可能将你需要的行为抽象成带有方法的接口，并使用这样的非空接口类型作为函数或方法的参数。

### 第29条 使用接口作为程序水平组合的连接点

+ “偏好组合，正交解耦”是Go语言的重要设计哲学之一。如果说“追求简单”聚焦的是为Go程序提供各种小而精的零件，那么组合关注的就是如何将这些零件关联到一起，搭建出程序的静态骨架。

### 29.1 一切皆组合

+ Go语言中主要有两种组合方式。
•  垂直组合（类型组合）：Go语言主要通过类型嵌入机制实现垂直组合，进而实现方法实现的复用、接口定义重用等。
•  水平组合：通常Go程序以接口类型变量作为程序水平组合的连接点。接口是水平组合的关键，它就好比程序肌体上的关节，给予连接关节的两个部分或多个部分各自自由活动的能力，而整体又实现了某种功能。

### 29.2 垂直组合回顾

+ Go语言通过类型的垂直组合而不是继承让单一类型承载更多的功能。由于不是继承，所以也就没有“父子类型”的概念，也没有向上、向下转型（type casting）；被嵌入的类型也不知道将其嵌入的外部类型的存在。调用方法时，方法的匹配取决于方法名称，而不是类型。

+ Go语言通过类型嵌入实现垂直组合。组合方式莫过于以下3种。
（1）通过嵌入接口构建接口
（2）通过嵌入接口构建结构体
（3）通过嵌入结构体构建新结构体

### 29.3 以接口为连接点的水平组合
1. 基本形式
水平组合的基本形式是接受接口类型参数的函数或方法，代码如下。
func YourFuncName(param YourInterfaceType)
2. 包裹函数
包裹函数（wrapper function）的形式是这样的：它接受接口类型参数，并返回与其参数类型相同的返回值。其代码如下：
```go
func YourWrapperFunc(param YourInterfaceType) YourInterfaceType
```
通过包裹函数可以实现对输入数据的过滤、装饰、变换等操作，并将结果再次返回给调用者。
下面是Go标准库中一个典型的包裹函数io.LimitReader：
```go
// $GOROOT/src/io/io.go 
func LimitReader(r Reader, n int64) Reader { 
    return &LimitedReader{r, n} 
}
type LimitedReader struct {    
    R Reader    N int64
}
func (l *LimitedReader) Read(p []byte) (n int, err error) {    ...}
```
我们看到LimitReader的一个输入参数为io.Reader接口类型，返回值类型依然为io.Reader。

+ 由于包裹函数的返回值类型与参数类型相同，因此我们可以将多个接受同一接口类型参数的包裹函数组合成一条链来调用，其形式如下：
YourWrapperFunc1(YourWrapperFunc2(YourWrapperFunc3(...)))

3. 适配器函数类型

>> 适配器函数类型（adapter function type）是一个辅助水平组合实现的“工具”类型。强调一下，它是一个类型。它可以将一个满足特定函数签名的普通函数显式转换成自身类型的实例，转换后的实例同时也是某个单方法接口类型的实现者。最典型的适配器函数类型莫过于第21条提到过的http.HandlerFunc了。
```go
// $GOROOT/src/net/http/server.go
type Handler interface {    
	ServeHTTP(ResponseWriter, *Request)
	}
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {    
	f(w, r)
}// chapter5/sources/horizontal-composition-3.go
func greetings(w http.ResponseWriter, r *http.Request) {    
	fmt.Fprintf(w, "Welcome!")
}
func main() {    
	http.ListenAndServe(":8080", http.HandlerFunc(greetings))
}
```
可以看到，在上述例子中通过http.HandlerFunc这个适配器函数类型，可以将普通函数greetings快速转换为实现了http.Handler接口的类型。转换后，我们便可以将其实例用作实参，实现基于接口的组合了。

4. 中间件
“中间件”（middleware）这个词的含义可大可小，在Go Web编程中，它常常指的是一个实现了http.Handler接口的http.HandlerFunc类型实例。实质上，这里的中间件就是包裹函数和适配器函数类型结合的产物。
```go
// chapter5/sources/horizontal-composition-4.go 
func validateAuth(s string) error {    
    if s != "123456" {        
        return fmt.Errorf("%s", "bad auth token")    
    }    
    return nil
}
func greetings(w http.ResponseWriter, r *http.Request) {    
    fmt.Fprintf(w, "Welcome!")
}
func logHandler(h http.Handler) http.Handler {    
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {        
            t := time.Now()        
            log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)        
            h.ServeHTTP(w, r)    
        })
}
func authHandler(h http.Handler) http.Handler {    
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {        
            err := validateAuth(r.Header.Get("auth"))        
            if err != nil {            
                http.Error(w, "bad auth param", http.StatusUnauthorized)            
                return        
            }        
            h.ServeHTTP(w, r)    
        })
}
func main() {    
    http.ListenAndServe(":8080", logHandler(
        authHandler(
            http.HandlerFunc(greetings))))
}
```
运行这个示例，并用curl工具命令对其进行测试：
`$ go run horizontal-composition-4.go$curl http://localhost:8080bad auth param$curl -H "auth:123456" localhost:8080/Welcome!`

+ 我们看到所谓中间件（如logHandler、authHandler）本质上就是一个包裹函数（支持链式调用），但其内部利用了适配器函数类型（http.HandlerFunc）将一个普通函数（如例子中的几个匿名函数）转换为实现了http.Handler的类型的实例，并将其作为返回值返回。

## 第30条 使用接口提高代码的可测试性


### 30.2 使用接口来降低耦合

>> 接口本是契约，天然具有降低耦合的作用。下面我们就用接口对v1版SendMailWithDisclaimer实现进行改造，将其对github.com/jordan-wright/email的依赖去除，将发送邮件的行为抽象成接口MailSender，并暴露给SendMailWithDisclaimer的用户。
```go
// chapter5/sources/send_mail_with_disclaimer/v2/mail.go
// 考虑到篇幅，这里省略一些代码...
type MailSender interface {    
    Send(subject, from string, to []string, content string, mailserver string, a smtp.Auth) error
}
func SendMailWithDisclaimer(sender MailSender, subject, from string,    to []string, content string, mailserver string, a smtp.Auth) error {    
    return sender.Send(subject, from, to, attachDisclaimer(content), mailserver, a)
}
```
现在如果要对SendMailWithDisclaimer进行测试，我们完全可以构造出一个或多个fake MailSender（根据不同单元测试用例的需求定制），

>> 下面是一个例子：
```go
// chapter5/sources/send_mail_with_disclaimer/v2/mail_test.go
package mail_test
import (    
    "net/smtp"    
    "testing"    
    mail "github.com/bigwhite/mail"
)
type FakeEmailSender struct {    
    subject string    
    from    string    
    to      []string    
    content string
}
func (s *FakeEmailSender) Send(subject, from string,    to []string, content string, mailserver string, a smtp.Auth) error {    
    s.subject = subject    
    s.from = from    
    s.to = to    
    s.content = content    
    return nil
}
func TestSendMailWithDisclaimer(t *testing.T) {    
    s := &FakeEmailSender{}    
    err := mail.SendMailWithDisclaimer(
        s, "gopher mail test v2",        "YOUR_MAILBOX",        []string{"DEST_MAILBOX"},        "hello, gopher",        "smtp.163.com:25",        smtp.PlainAuth("", "YOUR_EMAIL_ACCOUNT", "YOUR_EMAIL_PASSWD!", "smtp.163.com"))    
        if err != nil {        
            t.Fatalf("want: nil, actual: %s\n", err)        
            return    
        }    
        want := "hello, gopher" + "\n\n" + mail.DISCLAIMER    
        if s.content != want {        
            t.Fatalf("want: %s, actual: %s\n", want, s.content)    
        }
}
```
和v1版中的测试用例不同，v2版的测试用例不再对外部有任何依赖，是具备跨环境可重复性的。在这个用例中，我们对经过mail.SendMailWithDisclaimer处理后的content字段进行了验证，验证其是否包含免责声明，这也是在v1版中无法进行测试验证的。

>> 如果依然要使用github.com/jordan-wright/email包中Email实例作为邮件发送者，那么由于Email类型并不是上面MailSender接口的实现者，我们需要在业务代码中做一些适配工作，比如下面的代码：
```go
// chapter5/sources/send_mail_with_disclaimer/v2/example_test.go 
package mail_test
import (    
    "fmt"    
    "net/smtp"    
    mail "github.com/bigwhite/mail"    
    email "github.com/jordan-wright/email"
)
type EmailSenderAdapter struct {    
    e *email.Email
}
func (adapter *EmailSenderAdapter) Send(subject, from string,    to []string, content string, mailserver string, a smtp.Auth) error {    
    adapter.e.Subject = subject    
    adapter.e.From = from    
    adapter.e.To = to    
    adapter.e.Text = []byte(content)    
    return adapter.e.Send(mailserver, a)
}
func ExampleSendMailWithDisclaimer() {    
    adapter := &EmailSenderAdapter{        
        e: email.NewEmail(),    
    }    
    err := mail.SendMailWithDisclaimer(adapter, "gopher mail test v2",        "YOUR_MAILBOX",        []string{"DEST_MAILBOX"},        "hello, gopher",        "smtp.163.com:25",        smtp.PlainAuth("", "YOUR_EMAIL_ACCOUNT", "YOUR_EMAIL_PASSWD!", "smtp.163.com"))    
    if err != nil {        
        fmt.Printf("SendMail error: %s\n", err)        
        return    
    }    
    fmt.Println("SendMail ok")    // OutPut:    // SendMail ok
}
```
我们使用一个适配器对github.com/jordan-wright/email包中的Email实例进行了包装，使其成为接口MailSender的实现者，从而顺利传递给SendMailWithDisclaimer承担发送邮件的责任。

+ 接口MailSender将SendMailWithDisclaimer与具体的Email发送实现之间的耦合解开。通过上述例子我们也可以看出接口在测试过程中成为fake对象或mock对象的注入点。通过这种方式，我们可以通过灵活定制接口实现者以控制实现行为，继而实现对被测代码的代码逻辑的测试覆盖。小结代码的可测试性已经成为判定Go代码是否优秀的一条重要标准。适当抽取接口，让接口成为好代码与单元测试之间的桥梁是Go语言的一种最佳实践。


# 第六部分 并发编程

## 第31条 优先考虑并发设计

+ 优先考虑并发设计并发不是并行，并发关乎结构，并行关乎执行。——Rob Pike，Go语言之父

Go语言的设计哲学之一是“原生并发，轻量高效”

### 31.2 Go并发设计实例

写得很好!!!!

3. 第三版：并发方案

>> 假设机场鉴于现有建设规模，最大只能开通三条安检通道。机场旅客量依旧在增多，即便使用了并行方案，旅客的安检时长也无法再缩短。因为原安检程序采用的是顺序设计，即便机场目前有充足的人手（计算资源）可用，每个安检通道也只能用到一名工作人员。也就是说，原安检程序无法很好地适应工作人员（计算资源）的增加，是时候调整应用的结构了。
原先的安检程序（顺序设计）弊端很明显：当工作人员（计算资源）处于某一个检查环节（如人身检查），其他两个环节便处于“等待”状态。一条很显然的改进思路是让这些环节“同时”运行起来，就像流水线一样，这就是并发（见图31-5）。

## 第32条 了解goroutine的调度原理
### 32.1 goroutine调度器
### 32.2 goroutine调度模型与演进过程

1. G-M模型
2. G-P-M模型
### 32.3 对goroutine调度器原理的进一步理解
### 32.4 调度器状态的查看方法
### 32.4 调度器状态的查看方法


+ Go提供了调度器当前状态的查看方法：使用Go运行时环境变量GODEBUG。

+ 给GODEBUG传入了"schedtrace=1000"，其含义就是每1000ms打印输出一次goroutine调度器的状态，每次一行。

` SCHED 6016ms: gomaxprocs=4 idleprocs=0 threads=26 spinningthreads=0 idlethreads=20 runqueue=1 [3 4 0 10]`

•  SCHED：调试信息输出标志字符串，代表本行是goroutine调度器相关信息的输出。
•  6016ms：从程序启动到输出这行日志经过的时间。
•  gomaxprocs：P的数量。
•  idleprocs：处于空闲状态的P的数量。通过gomaxprocs和idleprocs的差值，我们就可以知道当前正在执行Go代码的P的数量。
•  threads：操作系统线程的数量，包含调度器使用的M数量，加上运行时自用的类似sysmon这样的线程的数量。
•  spinningthreads：处于自旋（spin）状态的操作系统数量。
•  idlethread：处于空闲状态的操作系统线程的数量。
•  runqueue=1：Go调度器全局运行队列中G的数量。
•  [3 4 0 10]：分别为4个P的本地运行队列中的G的数量。

+ 还可以输出每个goroutine、M和P的详细调度信息（对于Gopher来说，在大多数情况下这是不必要的）：
`$ GODEBUG=schedtrace=1000,scheddetail=1 godoc -http=:6060`

+  关于Go调度器调试信息输出的详细信息，可以参考Dmitry Vyukov的文章 “ Debugging Performance Issues in Go Programs”[1]，这也应该是每个Gopher必读的经典文章。更详尽的信息可参考$GOROOT/src/runtime/proc.go中schedtrace函数的实现。

`[1]https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs`



### 32.5 goroutine调度实例简要分析


+ 3）反转：如何在GOMAXPROCS=1的情况下让main goroutine得到调度？

```go
>> // chapter6/sources/go-scheduler-model-case3.go 
func add(a, b int) int {    
	return a + b
}
func deadloop() {    
	for {        
		add(3, 5)    
	}
}
func main() {    
	runtime.GOMAXPROCS(1)    
	go deadloop()    
	for {        
		time.Sleep(time.Second * 1)        
		fmt.Println("I got scheduled!")    
	}
}
```

+ 查看Go程序的汇编代码有多种方法。
•  使用objdump工具：objdump -S go二进制文件。
•  使用gdb disassemble。
•  使用go tool工具生成汇编代码文件：go build -gcflags '-S ' xx.go > xx.s 2>&1。

+ •  将Go代码编译成汇编代码：go tool compile -S xx.go > xx.s。
•  使用go tool工具反编译Go程序：go tool objdump -S go-binary > xx.s。

+ 这里使用最后一种方法：利用go tool objdump反编译（并结合其他输出的汇编形式）。
$go build -o go-scheduler-model-case3 go-scheduler-model-case3.go$go tool objdump -S go-scheduler-model-case3 > go-scheduler-model-case3.s

+ deadloop中对add函数的调用并未出现。这显然是Go编译器在生成代码时执行了内联（inline）优化的结果，因为add的调用对deadloop的行为结果没有任何影响。



## 第33条 掌握Go并发模型和常见并发模式

+ 不要通过共享内存来通信，而应该通过通信来共享内存。
——Rob Pike，Go语言之父

### 33.1 Go并发模型

+ 传统语言的并发模型是基于共享内存的模型

+ 在新并发模型设计中借鉴了著名计算机科学家Tony Hoare提出的CSP（Communicating Sequential Process，通信顺序进程）模型

+ 虽然CSP模型已经成为Go语言支持的主流并发模型，但Go也支持传统的基于共享内存的并发模型，并提供基本的低级同步原语（主要是sync包中的互斥锁、条件变量、读写锁、原子操作等）。

+ Go始终推荐以CSP模型风格构建并发程序，尤其是在复杂的业务层面。这将提升程序逻辑的清晰度，大大降低并发设计的复杂性，并让程序更具可读性和可维护性；对于局部情况，比如涉及性能敏感的区域或需要保护的结构体数据，可以使用更为高效的低级同步原语（如sync.Mutex），以保证goroutine对数据的同步访问。

### 33.2 Go常见的并发模式 !!!!!

+ 在语言层面，Go针对CSP模型提供了三种并发原语。
•  goroutine：对应CSP模型中的P，封装了数据的处理逻辑，是Go运行时调度的基本执行单元。
•  channel：对应CSP模型中的输入/输出原语，用于goroutine之间的通信和同步。
•  select：用于应对多路输入/输出，可以让goroutine同时协调处理多个channel操作。

1. 创建模式
Go语言使用go关键字+函数/方法创建goroutine

+ 但在稍复杂一些的并发程序中，需要考虑通过CSP模型输入/输出原语的承载体channel在goroutine之间建立联系。为了满足这一需求，我们通常使用下面的方式来创建goroutine：
```go
type T struct {...}
func spawn(f func()) chan T {    
	c := make(chan T)    
	go func() {        
		// 使用channel变量c(通过闭包方式)与调用spawn的goroutine通信        ...        
		f()        ...    
	}()        
	return c
}
func main() {    
	c := spawn(func(){})    // 使用channel变量c与新创建的goroutine通信
}
```

2. 退出模式
（1）分离模式
这里借鉴了一些线程模型中的术语，比如分离（detached）模式。分离模式是使用最为广泛的goroutine退出模式。对于分离的goroutine，创建它的goroutine不需要关心它的退出，这类goroutine在启动后即与其创建者彻底分离，其生命周期与其执行的主函数相关，函数返回即goroutine退出。这类goroutine有两个常见用途。
1）一次性任务：顾名思义，新创建的goroutine用来执行一个简单的任务，执行后即退出。比如下面标准库中的代码
2）常驻后台执行一些特定任务，如监视（monitor）、观察（watch）等。其实现通常采用for {...}或for { select{...} }代码段形式，并多以定时器（timer）或事件（event）驱动执行。
Go为每个goroutine调度模型中的P内置的GC goroutine就是这种类型的

(2）join模式
在Go中，我们有时候也有类似的需求：goroutine的创建者需要等待新goroutine结束。笔者为这样的goroutine退出模式起名为“join模式”
① 等待一个goroutine退出
② 获取goroutine的退出状态
③ 等待多个goroutine退出
通过Go语言提供的sync.WaitGroup实现等待多个goroutine退出的模式
④ 支持超时机制的等待

(3）notify-and-wait模式
① 通知并等待一个goroutine退出
② 通知并等待多个goroutine退出

（4）退出模式的应用
聚焦在实现一个“超时等待退出”框架，以统一解决各种运行形态goroutine的优雅退出问题。

3. 管道模式

+ Go中管道模式被实现成了由channel连接的一条“数据流水线”。在该流水线中，每个数据处理环节都由一组功能相同的goroutine完成。在每个数据处理环节，goroutine都要从数据输入channel获取前一个环节生产的数据，然后对这些数据进行处理，并将处理后的结果数据通过数据输出channel发往下一个环节。

+ 两种基于管道模式的扩展模式。
（1）扇出模式

>> 在某个处理环节，多个功能相同的goroutine从同一个channel读取数据并处理，直到该channel关闭，这种情况被称为“扇出”（fan-out）。使用扇出模式可以在一组goroutine中均衡分配工作量，从而更均衡地利用CPU。

>> （2）扇入模式
在某个处理环节，处理程序面对不止一个输入channel。我们把所有输入channel的数据汇聚到一个统一的输入channel，然后处理程序再从这个channel中读取数据并处理，直到该channel因所有输入channel关闭而关闭。这种情况被称为“扇入”（fan-in）。

4. 超时与取消模式

编写一个从气象数据服务中心获取气象信息的客户端。该客户端每次会并发向三个气象数据服务中心发起数据查询请求，并以最快返回的那个响应信息作为此次请求的应答返回值。

```go
// chapter6/sources/go-concurrency-pattern-12.go 
type result struct {    
	value string
}
func first(servers ...*httptest.Server) (result, error) {    
	c := make(chan result)    
	ctx, cancel := context.WithCancel(context.Background())    
	defer cancel()    
	queryFunc := func(i int, server *httptest.Server) {        
		url := server.URL        
		req, err := http.NewRequest("GET", url, nil)        
		if err != nil {            
			log.Printf("query goroutine-%d: http NewRequest error: %s\n", i, err)            
			return        
		}        
		req = req.WithContext(ctx)                
		log.Printf("query goroutine-%d: send request...\n", i)        
		resp, err := http.DefaultClient.Do(req)        
		if err != nil {            
			log.Printf("query goroutine-%d: get return error: %s\n", i, err)            
			return        
		}        
		log.Printf("query goroutine-%d: get response\n", i)        
		defer resp.Body.Close()        
		body, _ := ioutil.ReadAll(resp.Body)                
		c <- result{            
			value: string(body),        
		}        
		return    
	}        
	for i, serv := range servers {        
		go queryFunc(i, serv)    
	}        
	select {    
		case r := <-c:        return r, nil    
		case <-time.After(500 * time.Millisecond):        return result{}, errors.New("timeout")    
	}
}
func fakeWeatherServer(name string, interval int) *httptest.Server {    
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,         r *http.Request) {        
		log.Printf("%s receive a http request\n", name)        
		time.Sleep(time.Duration(interval) * time.Millisecond)        
		w.Write([]byte(name + ":ok"))    
	}))
}
func main() {    
	result, err := first(fakeWeatherServer("open-weather-1", 200),        fakeWeatherServer("open-weather-2", 1000),        
	fakeWeatherServer("open-weather-3", 600))    
	if err != nil {        
		log.Println("invoke first error:", err)        
		return    
	}        
	fmt.Println(result)    
	time.Sleep(10 * time.Second)
}
```
+ 在这版实现中，我们利用context.WithCancel创建了一个可以取消的context.Context变量，在每个发起查询请求的goroutine中，我们用该变量更新了request中的ctx变量，使其支持被取消。这样在first函数中，无论是成功得到某个查询goroutine的返回结果，还是超时失败返回，通过defer cancel()设定cancel函数在first函数返回前被执行，那些尚未返回的在途（on-flight）查询的goroutine都将收到cancel事件并退出（http包支持利用context.Context的超时和cancel机制）。

### 本条要点：
•  了解基于CSP的并发模型与传统基于共享内存的并发模型的区别；
•  了解Go为实现CSP模型而提供的并发原语及功能；
•  掌握常见的并发模式，包括创建模式、多种退出模式、管道模式、超时和取消模式等。

## 第34条 了解channel的妙用 !!!!!
+ channel是Go语言提供的一种重要的并发原语。从前文中我们了解到，它在Go语言的CSP模型中扮演着重要的角色：既可以实现goroutine间的通信，又可以实现goroutine间的同步。

### 34.1 无缓冲channel

+ 无缓冲channel兼具通信和同步特性，在并发程序中应用颇为广泛。

+ 对于无缓冲channel而言，我们得到以下结论：
•  发送动作一定发生在接收动作完成之前；
•  接收动作一定发生在发送动作完成之前。

1. 用作信号传递

（1）一对一通知信号
（2）一对多通知信号

2. 用于替代锁机制

+ 无缓冲channel具有同步特性，这让它在某些场合可以替代锁，从而使得程序更加清晰，可读性更好。

### 34.2 带缓冲channel

1. 用作消息队列

（2）多收多发性能基准测试

2. 用作计数信号量
3. len(channel)的应用

+ 如果s是chan T类型，那么len(s)针对channel的类型不同，有如下两种语义：
◦  当s为无缓冲channel时，len(s)总是返回0；
◦  当s为带缓冲channel时，len(s)返回当前channel s中尚未被读取的元素个数。

+ 为了不阻塞在channel上，常见的方法是将判空与读取放在一个事务中，将判满与写入放在一个事务中，而这类事务我们可以通过select实现。

### 34.3 nil channel的妙用
+  对没有初始化的channel（nil channel）进行读写操作将会发生阻塞
+ 但nil channel并非一无是处，有些时候妙用nil channel可以达到事半功倍的效果。

```go
// chapter6/sources/go-channel-case-10.go 
func main() {    
	c1, c2 := make(chan int), make(chan int)    
	go func() {        
		time.Sleep(time.Second * 5)        
		c1 <- 5        
		close(c1)    
	}()        
	go func() {        
		time.Sleep(time.Second * 7)        
		c2 <- 7        
		close(c2)    
	}()        
	for {        
		select {        
			case x, ok := <-c1:            
			if !ok {                
				c1 = nil            
			} else {                
				fmt.Println(x)            
			}        
			case x, ok := <-c2:                
			if !ok {                    
				c2 = nil                
			} else {                    
				fmt.Println(x)                
			}        
		}        
		if c1 == nil && c2 == nil {            
			break        
		}    
	}    
	fmt.Println("program end")
}
```
改进后的示例程序的最关键变化是在判断c1或c2被关闭后，显式地将c1或c2置为nil。我们知道，对一个nil channel执行获取操作，该操作将被阻塞，因此已经被置为nil的c1或c2的分支将再也不会被select选中执行。

### 34.4 与select结合使用的一些惯用法

1. 利用default分支避免阻塞select语句的default分支的语义是在其他分支均因通信未就绪而无法被选择的时候执行，这就为default分支赋予了一种“避免阻塞”的特性
2. 实现超时机制

+ 带超时机制的select是Go语言中一种常见的select和channel的组合用法，通过超时事件，我们既可以避免长期陷入某种操作的等待中，也可以做一些异常处理工作。下面的示例代码实现了一次具有30s超时的select：
```go
func worker() {    
	select {    
		case <-c:        // ...    
		case <-time.After(30 *time.Second):        
		return    
	}
}
```
3. 实现心跳机制结合time包的Ticker，我们可以实现带有心跳机制的select。这种机制使我们可以在监听channel的同时，执行一些周期性的任务

### 小结
Go channel就像Go并发模型中的“胶水”，它将诸多并发执行单元连接起来，或者正是因为有channel的存在，Go并发模型才能迸发出强大的表达能力。

### 本条要点：
了解Go并发原语channel和select的基本语义；掌握无缓冲channel在信号传递、替代锁同步场景下的应用模式；掌握带缓冲channel在消息队列、计数信号量场景下的应用模式，了解在特定场景下利用len函数侦测带缓冲channel的状态；了解nil channel在特定场景下的用途；掌握select与channel结合使用的一些惯用法及注意事项。

## 第35条 了解sync包的正确用法

+ Go语言在提供CSP并发模型原语的同时，还通过标准库的sync包提供了针对传统基于共享内存并发模型的基本同步原语，包括互斥锁（sync.Mutex）、读写锁（sync.RWMutex）、条件变量（sync.Cond）等。

### 35.1 sync包还是channel

+ Go语言提倡“不要通过共享内存来通信，而应该通过通信来共享内存”。

（1）需要高性能的临界区同步机制场景

（2）不想转移结构体对象所有权，但又要保证结构体内部状态数据的同步访问的场景基于channel的并发设计的一个特点是，在goroutine间通过channel转移数据对象的所有权。只有拥有数据对象所有权（从channel接收到该数据）的goroutine才可以对该数据对象进行状态变更。如果你的设计中没有转移结构体对象所有权，但又要保证结构体内部状态数据能在多个goroutine之间同步访问，那么你可以使用sync包提供的低级同步原语来实现，比如最常用的sync.Mutex。

### 35.2 使用sync包的注意事项

+ 在$GOROOT/src/sync/mutex.go文件中，我们看到这样一行关于使用sync包的注意事项：// Values containing the types defined in this package should not be copied.// 不应复制那些包含了此包中类型的值在sync包的其他源文件中，我们还会看到如下的一些注释：// $GOROOT/src/sync/mutex.go// A Mutex must not be copied after first use. (禁止复制首次使用后的Mutex)// $GOROOT/src/sync/rwmutex.go// A RWMutex must not be copied after first use.(禁止复制首次使用后的RWMutex)// $GOROOT/src/sync/cond.go// A Cond must not be copied after first use.(禁止复制首次使用后的Cond)...

+ 通过上述示例我们直观地看到，那些sync包中类型的实例在首次使用后被复制得到的副本一旦再被使用将导致不可预期的结果，为此在使用sync包中类型时，推荐通过闭包方式或传递类型实例（或包裹该类型的类型实例）的地址或指针的方式进行，这是使用sync包最值得注意的事项。

### 35.3 互斥锁还是读写锁

+ sync包提供了两种用于临界区同步的原语：互斥锁（Mutex）和读写锁（RWMutex）。互斥锁是临界区同步原语的首选，它常被用来对结构体对象的内部状态、缓存等进行保护，是使用最为广泛的临界区同步原语。

•  在并发量较小的情况下，互斥锁性能更好；随着并发量增大，互斥锁的竞争激烈，导致加锁和解锁性能下降。
•  读写锁的读锁性能并未随并发量的增大而发生较大变化，性能始终恒定在40ns左右。
•  在并发量较大的情况下，读写锁的写锁性能比互斥锁、读写锁的读锁都差，并且随着并发量增大，其写锁性能有继续下降的趋势。

+ 读写锁适合应用在具有一定并发量且读多写少的场合

### 35.4 条件变量

+ sync.Cond是传统的条件变量原语概念在Go语言中的实现。一个条件变量可以理解为一个容器，这个容器中存放着一个或一组等待着某个条件成立的goroutine。当条件成立时，这些处于等待状态的goroutine将得到通知并被唤醒以继续后续的工作。这与百米飞人大战赛场上各位运动员等待裁判员的发令枪声十分类似。

### 35.5 使用sync.Once实现单例模式

+ 程序运行期间只被执行一次且goroutine安全的函数只有每个包的init函数。sync包提供了另一种更为灵活的机制，可以保证任意一个函数在程序运行期间只被执行一次，这就是sync.Once。

+ 在Go标准库中，sync.Once的“仅执行一次”语义被一些包用于初始化和资源清理的过程中，以避免重复执行初始化或资源关闭操作。

## 35.6 使用sync.Pool减轻垃圾回收压力 !!!

•  它是goroutine并发安全的，可以被多个goroutine同时使用；
•  放入该缓存池中的数据对象的生命是暂时的，随时都可能被垃圾回收掉；
•  缓存池中的数据对象是可以重复利用的，这样可以在一定程度上降低数据对象重新分配的频度，减轻GC的压力；
•  sync.Pool为每个P（goroutine调度模型中的P）单独建立一个local缓存池，进一步降低高并发下对锁的争抢。

（1）限制要放回缓存池中的数据对象大小

（2）建立多级缓存池

### 本条要点：
•  明确sync包中原语应用的适用场景；
•  sync包内定义的结构体或包含这些类型的结构体在首次使用后禁止复制；
•  明确sync.RWMutex的适用场景；
•  掌握条件变量的应用场景和使用方法；
•  实现单例模式时优先考虑sync.Once；
•  了解sync.Pool的优点、使用中可能遇到的问题及解决方法。

## 第36条 使用atomic包实现伸缩性更好的并发读取


### 36.1 atomic包与原子操作

+ atomic包封装了CPU实现的部分原子操作指令，为用户层提供体验良好的原子操作函数，因此atomic包中提供的原语更接近硬件底层，也更为低级，它常被用于实现更为高级的并发同步技术（比如channel和sync包中的同步原语）。

## 36.2 对共享整型变量的无锁读写

+ atomic包提供了两大类原子操作接口：一类是针对整型变量的，包括有符号整型、无符号整型以及对应的指针类型；另一类是针对自定义类型的。第一类原子操作接口的存在让atomic包天然适合于实现某一个共享整型变量的并发同步。

•  读写锁的性能随着并发量增大的变化情况与前面讲解sync.RWMutex时的一致；
•  利用原子操作的无锁并发写的性能随着并发量增大几乎保持恒定；
•  利用原子操作的无锁并发读的性能随着并发量增大有持续提升的趋势，并且性能约为读锁的200倍。

### 36.3 对共享自定义类型变量的无锁读写

+ 我们再来看atomic包另一类函数的应用。如图36-2所示，atomic通过Value类型的装拆箱操作实现了对任意自定义类型的原子操作（Load和Store），从而实现对共享自定义类型变量无锁读写的支持。

•  利用原子操作的无锁并发写的性能随着并发量的增大而小幅下降；
•  利用原子操作的无锁并发读的性能随着并发量增大有持续提升的趋势，并且性能约为读锁的100倍。


### 小结
是时候揭晓答案了。由上面两类atomic包应用的例子可知，随着并发量提升，使用atomic实现的共享变量的并发读写性能表现更为稳定，尤其是原子读操作，这让atomic与sync包中的原语比起来表现出更好的伸缩性和更高的性能。由此可以看出atomic包更适合一些对性能十分敏感、并发量较大且读多写少的场合。
但atomic原子操作可用来同步的范围有较大限制，仅是一个整型变量或自定义类型变量。如果要对一个复杂的临界区数据进行同步，那么首选依旧是sync包中的原语。

# 第七部分 错误处理

## 第37条 了解错误处理的4种策略 !!!

### 37.1 构造错误值
### 37.2 透明错误处理策略
### 37.3 “哨兵”错误处理策略
### 37.4 错误值类型检视策略
### 37.5 错误行为特征检视策略

### 小结
Go社区中关于如何进行错误处理的讨论有很多，但唯一正确的结论是没有哪一种错误处理策略适用于所有项目或场合。综合上述的构造错误值方法及错误处理策略，请记住如下几点：
•  尽量使用透明错误处理策略降低错误处理方与错误值构造方之间的耦合；
•  如果可以通过错误值类型的特征进行错误检视，那么尽量使用错误行为特征检视策略；
•  在上述两种策略无法实施的情况下，再用“哨兵”策略和错误值类型检视策略；
•  在Go 1.13及后续版本中，尽量用errors.Is和errors.As方法替换原先的错误检视比较语句。


## 第38条 尽量优化反复出现的if err != nil  !!!

### 38.1 两种观点

### 38.2 尽量优化

+ Lohuizen也对if err != nil的重复出现情况进行了研究。如图38-2所示，他发现代码所在栈帧越低（越接近于main函数栈帧），if err != nil就越不常见；反之，代码在栈中的位置越高（更接近于网络I/O操作或操作系统API调用），if err != nil就越常见

### 38.3 优化思路
1）改善代码的视觉呈现。
1. 视觉扁平化
2. 重构：减少if err != nil的重复次数
3. check/handle风格化
+ panic和recover让函数调用的性能降低了约90%。因此，我们在使用这种方案优化重复代码前，需要全面了解这些约束。
4. 封装：内置error状态

+ 这显然是消除if err != nil代码片段重复出现的理想方法。我们还是以CopyFile为例，看看使用这种“内置error状态”的新封装方法后，能得到什么样的代码：
```go
// chapter7/sources/go-if-error-check-optimize-3.go
type FileCopier struct {    
	w   *os.File    
	r   *os.File    
	err error
}
func (f *FileCopier) open(path string) (*os.File, error) {    
	if f.err != nil {        
		return nil, f.err    
	}        
	h, err := os.Open(path)    
	if err != nil {        
		f.err = err        
		return nil, err    
	}    
	return h, nil
}
func (f *FileCopier) openSrc(path string) {    
	if f.err != nil {        
		return    
	}        
	f.r, f.err = f.open(path)    
	return
}
func (f *FileCopier) createDst(path string) {    
	if f.err != nil {        
		return    
	}        
	f.w, f.err = os.Create(path)    
	return
}
func (f *FileCopier) copy() {    
	if f.err != nil {        
		return    
	}        
	if _, err := io.Copy(f.w, f.r); err != nil {        
		f.err = err    
	}
}

func (f *FileCopier) CopyFile(src, dst string) error {    
	if f.err != nil {        
		return f.err    
	}        
	defer func() {        
		if f.r != nil {            
			f.r.Close()        
		}        
		if f.w != nil {            
			f.w.Close()        
		}        
		if f.err != nil {            
			if f.w != nil {                
				os.Remove(dst)            
			}        
		}    
	}()        
	f.openSrc(src)    
	f.createDst(dst)    
	f.copy()    
	return f.err
}
func main() {    
	var fc FileCopier    
	err := fc.CopyFile("foo.txt", "bar.txt")    
	if err != nil {        
		fmt.Println("copy file error:", err)        
		return    
	}    
	fmt.Println("copy file ok")
}
```
这次的重构很彻底。我们将原CopyFile函数彻底抛弃，而重新将其逻辑封装到一个名为FileCopier结构的CopyFile方法中。FileCopier结构内置了一个err字段用于保存内部的错误状态，这样在其CopyFile方法中，我们只需按照正常业务逻辑，顺序执行openSrc、createDst和copy即可，正常业务逻辑的视觉连续性就这样被很好地实现了。同时该CopyFile方法的复杂度因if检查的“大量缺席”而变得很低。

### 小结
Go显式错误处理的设计既有其优势，也有其编写冗长的不足，至今针对Go错误处理尚未形成一致的改进意见。我们能做的就是尽可能对反复出现的if err != nil进行优化，本条给出了若干优化思路。
本条要点：
•  使用显式错误结果和显式的错误检查是Go语言成功的重要因素，也是if err != nil反复出现的根本原因；
•  了解关于改善Go错误处理的两种观点；
•  了解减少甚至消除if err != nil代码片段的两个优化方向，即改善视觉呈现与降低复杂度；
•  掌握错误处理代码优化的四种常见方法（位于三个不同象限中），并根据所处场景与约束灵活使用。

## 第39条 不要使用panic进行正常的错误处理

+ Go的正常错误处理与异常处理之间是泾渭分明的，这与其他主流编程语言使用结构化错误处理统一处理错误与异常是两种不同的理念。Go提供了panic专门用于处理异常，而我们建议不要使用panic进行正常的错误处理。

1. checked exception实质是错误，而panic是异常

+ 和Java中checked exception的“有意而为之”相反，在Go中，panic则是“不得已而为之”，即所有引发panic的情形，无论是显式的（我们主动调用panic函数引发的）还是隐式的（Go运行时检测到违法情况而引发的），都是我们不期望看到的。对这些引发的panic，我们很少有预案应对，更多的是让程序快速崩溃掉。因此一旦发生panic，就意味着我们的代码很大可能出现了bug。因此，Go中的panic更接近于Java的RuntimeException+Error，而不是checked exception。

2. API调用者没有义务处理panic

3. 未被捕获的panic意味着“游戏结束”

+ 综上，Go panic不应被当作Java的checked exception来进行正常的错误处理。使用错误 （error）和多返回值的显式错误处理方式才符合Go的错误处理哲学。

>> [1]https://blog.golang.org/defer-panic-and-recover

### 39.2 panic的典型应用

1. 充当断言角色，提示潜在bug
2. 用于简化错误处理控制结构
3. 使用recover捕获panic，防止goroutine意外退出

### 39.3 理解panic的输出信息 !!!


+ 关于发生panic后输出的栈跟踪信息（stack trace）的识别，总体可遵循以下几个要点。
•  栈跟踪信息中每个函数/方法后面的“参数数值”个数与函数/方法原型的参数个数不是一一对应的。
•  栈跟踪信息中每个函数/方法后面的“参数数值”是按照函数/方法原型参数列表中从左到右的参数类型的内存布局逐一展开的，每个数值占用一个字（word，64位平台下为8字节）。
•  如果是方法，则第一个参数是receiver自身。如果receiver是指针类型，则第一个参数数值就是一个指针地址；如果是非指针的实例，则栈跟踪信息会按照其内存布局输出。
•  函数/方法返回值放在栈跟踪信息的“参数数值”列表的后面；如果有多个返回值，则同样按从左到右的顺序，按照返回值类型的内存布局输出。
•  指针类型参数：占用栈跟踪信息的“参数数值”列表的一个位置；数值表示指针值，也是指针指向的对象的地址。
•  string类型参数：由于string在内存中由两个字表示（第一个字是数据指针，第二个字是string的长度），因此在栈跟踪信息的“参数数值”列表中将占用两个位置。
•  slice类型参数：由于slice类型在内存中由三个字表示（第一个字是数据指针，第二个字是len，第三个字是cap），因此在栈跟踪信息的“参数数值”列表中将占用三个位置。
•  内建整型（int、rune、byte）：由于按字逐个输出，对于类型长度不足一个字的参数，会进行合并处理。比如，一个函数有5个int16类型的参数，那么在栈跟踪信息中这5个参数将占用“参数数值”列表中的两个位置：第一个位置是前4个参数的“合体”，第二个位置则是最后那个int16类型的参数值。
•  struct类型参数：会按照struct中字段的内存布局顺序在栈跟踪信息中展开。
•  interface类型参数：由于interface类型在内存中由两部分组成（一部分是接口类型的参数指针，另一部分是接口值的参数指针），因此interface类型参数将使用“参数数值”列表中的两个位置。
•  栈跟踪输出的信息是在函数调用过程中的“快照”信息，因此一些输出数值虽然看似不合理，但由于其并不是最终值，问题也不一定发生在它们身上，比如返回值参数

+ 在Go 1.11及以后版本中，Go编译器得到更深入的优化，很多简单的函数或方法会被自动内联（inline）。函数一旦内联化，我们就无法在栈跟踪信息中看到栈帧信息了，栈帧信息都变成了省略号

+ 要想看到栈跟踪信息中的栈帧数据，我们需要使用-gcflags="-l"来告诉编译器不要执行内联优化
!!!!


# 第八部分 测试、性能剖析与调试

## 第40条理解包内测试与包外测试的差别
### 40.1　官方文档的“自相矛盾”
### 40.2　包内测试与包外测试
1. Go标准库中包内测试和包外测试的使用情况Go标准库是Go代码风格和惯用法一贯的风向标。
2. 包内测试的优势与不足
3. 包外测试（仅针对导出API的测试）
4. 优先使用包外测试


### 本条要点：
go test执行测试的原理；理解包内测试的优点与不足；理解包外测试的优点与不足；掌握通过export_test.go为包外测试添加“后门”的惯用法；优先使用包外测试；当运用包外测试与包内测试共存的方式时，可考虑让包外测试和包内测试聚焦于不同的测试类别。

## 第41条 有层次地组织测试代码

### 41.1　经典模式——平铺

+ 平铺模式的测试代码组织方式的优点是显而易见的。简单：没有额外的抽象，上手容易。独立：每个测试函数都是独立的，互不关联，避免相互干扰。

### 41.2　xUnit家族模式

+ go test的输出也更有层次感了，我们可以一眼看出对哪些函数/方法进行了测试、这些被测对象对应的测试套件以及套件中的每个测试用例。

### 41.3　测试固件

+ 我们一般使用setUp和tearDown来代表测试固件的创建/设置与拆除/销毁的动作。

### 41.3　测试固件

+ 我们一般使用setUp和tearDown来代表测试固件的创建/设置与拆除/销毁的动作。

+ Go 1.14版本testing包增加了testing.Cleanup方法，为测试固件的销毁提供了包级原生的支持：
```go
func setUp() func(){    ...    
return func() {    }
}
func TestXxx(t *testing.T) {    
    t.Cleanup(setUp())    ...
}
```
+ Go 1.4版本引入了TestMain，使得包级别测试固件的创建和销毁终于有了正式的施展舞台。

+ 平铺模式下的测试执行流有些时候，一些测试函数所需的测试固件是相同的，在平铺模式下为每个测试函数都单独创建/销毁一次测试固件就显得有些重复和冗余。在这样的情况下，我们可以尝试采用测试套件来减少测试固件的重复创建。

+ xUnit实践模式下的测试执行流小结在确定了将测试代码放入包内测试还是包外测试之后，我们在编写测试前，还要做好测试包内部测试代码的组织规划，建立起适合自己项目规模的测试代码层次体系。简单的测试可采用平铺模式，复杂的测试可借鉴xUnit的最佳实践，利用subtest建立包、测试套件、测试用例三级的测试代码组织形式，并利用TestMain和testing.Cleanup方法为各层次的测试代码建立测试固件。


## 第42条 优先编写表驱动的测试

### 42.1　Go测试代码的一般逻辑

+ 对测试失败与否的判断在于测试代码逻辑是否进入了包含Error/Errorf、Fatal/Fatalf等方法调用的代码分支。一旦进入这些分支，即代表该测试失败。不同的是Error/Errorf并不会立刻终止当前goroutine的执行，还会继续执行该goroutine后续的测试，而Fatal/Fatalf则会立刻停止当前goroutine的测试执行。
### 42.2　表驱动的测试实践

+ 上面仅有三组预置输入数据的示例的测试代码已显得十分冗长，如果为测试预置的数据组数增多，测试函数本身就将变得十分庞大。并且，我们看到上述示例的测试逻辑中存在很多重复的代码，显得十分烦琐。我们来尝试对上述示例做一些改进：
```go
// chapter8/sources/table_driven_strings_test.go
func TestCompare(t *testing.T) {    
    compareTests := []struct {        
        a, b string        
        i    int    
    }
    {        
        {"", "", 0},        
        {"a", "", 1},        
        {"", "a", -1},    
    }    
    for _, tt := range compareTests {        
        cmp := strings.Compare(tt.a, tt.b)        
        if cmp != tt.i {            
            t.Errorf(`want %v, but Compare(%q, %q) = %v`, tt.i, tt.a, tt.b, cmp)        
        }    
    }
```
在上面这个改进的示例中，我们将之前示例中重复的测试逻辑合并为一个，并将预置的输入数据放入一个自定义结构体类型的切片中。这个示例的长度看似并没有比之前的实例缩减多少，但它却是一个可扩展的测试设计。

+  无须改动后面的测试逻辑，只需在切片中增加数据条目即可。在这种测试设计中，这个自定义结构体类型的切片（上述示例中的compareTests）就是一个表（自定义结构体类型的字段就是列），而基于这个数据表的测试设计和实现则被称为“表驱动的测试”。

## 42.3　表驱动测试的优点

+ 表驱动测试本身是编程语言无关的。

表驱动测试有着诸多优点。
（1）简单和紧凑
（2）数据即测试
（3）结合子测试后，可单独运行某个数据项的测试

+ 我们将表驱动测试与子测试（subtest）结合来改造一下上面的strings_test示例：
```go
// chapter8/sources/table_driven_strings_with_subtest_test.go
func TestCompare(t *testing.T) {    
    compareTests := []struct {        
        name, a, b string        
        i          int    
    }
    {        
        {`compareTwoEmptyString`, "", "", 0},        
        {`compareSecondParamIsEmpty`, "a", "", 1},        
        {`compareFirstParamIsEmpty`, "", "a", -1},    
    }    
    for _, tt := range compareTests {        
        t.Run(tt.name, func(t *testing.T) {            
            cmp := strings.Compare(tt.a, tt.b)            
            if cmp != tt.i {                
                t.Errorf(`want %v, but Compare(%q, %q) = %v`, tt.i, tt.a, tt.b, cmp)            
            }        
        })    
    }
}
```

+ 我们将表驱动测试与子测试（subtest）结合来改造一下上面的strings_test示例：
```go
// chapter8/sources/table_driven_strings_with_subtest_test.go
func TestCompare(t *testing.T) {    
    compareTests := []struct {        
        name, a, b string        
        i          int    
    }
    {        
        {`compareTwoEmptyString`, "", "", 0},        
        {`compareSecondParamIsEmpty`, "a", "", 1},        
        {`compareFirstParamIsEmpty`, "", "a", -1},    
    }    
    for _, tt := range compareTests {        
        t.Run(tt.name, func(t *testing.T) {            
            cmp := strings.Compare(tt.a, tt.b)            
            if cmp != tt.i {                
                t.Errorf(`want %v, but Compare(%q, %q) = %v`, tt.i, tt.a, tt.b, cmp)            
            }        
        })    
    }
}
```

### 42.4　表驱动测试实践中的注意事项

1. 表的实现方式

+ 在上面的示例中，测试中使用的表是用自定义结构体的切片实现的，表也可以使用基于自定义结构体的其他集合类型（如map）来实现。我们将上面的例子改造为采用map来实现测试数据表

+ 为了在表测试驱动的测试中快速从输出的结果中定位导致测试失败的表项，我们需要在测试失败的输出结果中输出数据表项的唯一标识。
最简单的方法是通过输出数据表项在数据表中的偏移量来辅助定位“元凶”

+ 另一个更直观的方式是使用名字来区分不同的数据项

```go
if cmp != tt.i {            
    t.Errorf(`[%s] want %v, but Compare(%q, %q) = %v`, tt.name, tt.i, tt.a, tt.b, cmp)        
}
```

3. Errorf还是Fatalf

+ 一般而言，如果一个数据项导致的测试失败不会对后续数据项的测试结果造成影响，那么推荐Errorf，这样可以通过执行一次测试看到所有导致测试失败的数据项；否则，如果数据项导致的测试失败会直接影响到后续数据项的测试结果，那么可以使用Fatalf让测试尽快结束，因为继续执行的测试的意义已经不大了。

## 第43条使用testdata管理测试依赖的外部数据文件

+ 测试固件是Go测试执行所需的上下文环境，其中测试依赖的外部数据文件就是一种常见的测试固件（可以理解为静态测试固件，因为无须在测试代码中为其单独编写固件的创建和清理辅助函数）。

### 43.1　testdata目录

+ Go语言规定：Go工具链将忽略名为testdata的目录。这样开发者在编写测试时，就可以在名为testdata的目录下存放和管理测试代码依赖的数据文件。而go test命令在执行时会将被测试程序包源码所在目录设置为其工作目录，这样如果要使用testdata目录下的某个数据文件，我们无须再处理各种恼人的路径问题，而可以直接在测试代码中像下面这样定位到充当测试固件的数据文件：
`f, err := os.Open("testdata/data-001.txt")`

+ 差别（Windows下使用反斜线“\”，Linux/macOS下使用斜线“/”），使用下面的方式可以使测试代码更具可移植性：
`f, err := os.Open(filepath.Join("testdata", "data-001.txt"))`

+ 我们还经常将预期结果数据保存在文件中并放置在testdata下，然后在测试代码中将被测对象输出的数据与这些预置在文件中的数据进行比较，一致则测试通过；反之，测试失败。

### 43.2　golden文件惯用法

+ Go标准库为我们提供了一种惯用法：golden文件

### 小结
+ 在这一条中，我们了解到面向工程的Go语言对测试依赖的外部数据文件的存放位置进行了规范，统一使用testdata目录，开发人员可以采用将预期数据文件放在testdata下的方式为测试提供静态测试固件。而Go golden文件的惯用法实现了testdata目录下测试依赖的预期结果数据文件的数据采集与测试代码的融合。

## 第44条正确运用fake、stub和mock等辅助单元测试 !!!!

### 44.1　fake：真实组件或服务的简化实现版替身

+ fake测试就是指采用真实组件或服务的简化版实现作为替身，以满足被测代码的外部依赖需求。
+ 使用fake替身进行测试的最常见理由是在测试环境无法构造被测代码所依赖的外部组件或服务，或者这些组件/服务有副作用。fake替身的实现也有两个极端：要么像标准库fakedb_test.go那样实现一个全功能的简化版内存数据库driver，要么像faketest1例子中那样针对被测代码的调用请求仅返回硬编码的成功或失败。这两种极端实现有一个共同点：并不具备在测试前对返回结果进行预设置的能力。这也是上面例子中我们针对成功和失败两个用例分别实现了一个替身的原因。（如果非要说成功和失败也是预设置的，那么fake替身的预设置能力也仅限于设置单一的返回值，即无论调用多少次，传入什么参数，返回值都是一个。）

### 44.2　stub：对返回结果有一定预设控制能力的替身

+ stub也是一种替身概念，和fake替身相比，stub替身增强了对替身返回结果的间接控制能力，这种控制可以通过测试前对调用结果预设置来实现。不过，stub替身通常仅针对计划之内的结果进行设置，对计划之外的请求也无能为力。

+ 使用Go标准库net/http/httptest实现的用于测试的Web服务就可以作为一些被测对象所依赖外部服务的stub替身。
+ 在GitHub上有一个名为gostub（https://github.com/prashantv/gostub）的第三方包可以用于简化stub替身的管理和编写。

### 44.3　mock：专用于行为观察和验证的替身

+ 和fake、stub替身相比，mock替身更为强大：它除了能提供测试前的预设置返回结果能力之外，还可以对mock替身对象在测试过程中的行为进行观察和验证。不过相比于前两种替身形式，mock存在应用局限（尤指在Go中）。

+  和前两种替身相比，mock的应用范围要窄很多，只用于实现某接口的实现类型的替身。
◦  一般需要通过第三方框架实现mock替身。Go官方维护了一个mock框架——gomock（https://github.com/golang/mock），该框架通过代码生成的方式生成实现某接口的替身类型。

+ gomock是一个通用的mock框架，社区还有一些专用的mock框架可用于快速创建mock替身，比如：go-sqlmock（https://github.com/DATA-DOG/go-sqlmock）专门用于创建sql/driver包中的Driver接口实现的mock替身，可以帮助Gopher简单、快速地建立起对数据库操作相关方法的单元测试。

### 小结
本条介绍了当被测代码对外部组件或服务有强依赖时可以采用的测试方案，这些方案采用了相同的思路：为这些被依赖的外部组件或服务建立替身。这里介绍了三类替身以及它们的适用场合与注意事项。
本条要点如下。
◦  fake、stub、mock等替身概念之间并非泾渭分明的，对这些概念的理解容易混淆。比如标准库net/http/transfer_test.go文件中的mockTransferWriter类型，虽然其名字中带有mock，但实质上它更像是一个fake替身。
◦  我们更多在包内测试应用上述替身概念辅助测试，这就意味着此类测试与被测代码是实现级别耦合的，这样的测试健壮性较差，一旦被测代码内部逻辑有变化，测试极容易失败。
◦  通过fake、stub、mock等概念实现的替身参与的测试毕竟是在一个虚拟的“沙箱”环境中，不能代替与真实依赖连接的测试，因此，在集成测试或系统测试等使用真实外部组件或服务的测试阶段，务必包含与真实依赖的联测用例。
◦  fake替身主要用于被测代码依赖组件或服务的简化实现。
◦  stub替身具有有限范围的、在测试前预置返回结果的控制能力。
◦  mock替身则专用于对替身的行为进行观察和验证的测试，一般用作Go接口类型的实现的替身。


## 第45条使用模糊测试让潜在bug无处遁形

+ 模糊测试就是指半自动或自动地为程序提供非法的、非预期、随机的数据，并监控程序在这些输入数据下是否会出现崩溃、内置断言失败、内存泄露、安全漏洞等情况（见图45-1）。

### 45.1　模糊测试在挖掘Go代码的潜在bug中的作用

### 45.2　go-fuzz的初步工作原理

### 45.3　go-fuzz使用方法

### 小结
+ 通过这一条，我们认识到模糊测试对于提升Go代码质量、挖掘潜在bug的重要作用。但模糊测试不是“银弹”，它有其适用的范围。模糊测试最适合那些处理复杂输入数据的程序，比如文件格式解析、网络协议解析、人机交互界面入口等。模糊测试是软件测试技术的一个重要分支，与单元测试等互为补充，相辅相成。
目前，并非所有编程语言都有对模糊测试工具的支持，Gopher和Go社区很幸运，Dmitry Vyukov为我们带来了go-fuzz模糊测试工具。如果你是追求高质量Go代码的开发者，请为你的Go代码建立起模糊测试。


## 第46条为被测对象建立性能基准

+ 是否优化、何时优化实质上是一个决策问题，但决策不能靠直觉，要靠数据说话。借用上面名言中的句型：没有数据支撑的过早决策是万恶之源。

### 46.1　性能基准测试在Go语言中是“一等公民”

+ 性能基准测试还可以通过传入-benchmem命令行参数输出内存分配信息（与基准测试代码中显式调用b.ReportAllocs的效果是等价的）

` $go test -bench=Join ./benchmark_intro_test.go -benchmemgoos: darwingoarch: amd64BenchmarkConcatStringByJoin-8     23004709   48.8 ns/op   48 B/op     1 allocs/opPASSok         command-line-arguments 1.183s`

这里输出的内存分配信息告诉我们，每执行一次concatStringByJoin平均进行一次内存分配，每次平均分配48字节的数据。

### 46.2　顺序执行和并行执行的性能基准测试

+ 根据是否并行执行，Go的性能基准测试可以分为两类：顺序执行的性能基准测试和并行执行的性能基准测试。

1. 顺序执行的性能基准测试
2. 并行执行的性能基准测试

```go
func BenchmarkXxx(b *testing.B) {    
    // ...    
    b.RunParallel(func(pb *testing.PB) {        
        for pb.Next() {            
            // 被测对象的执行代码        
        }    
    }
}
```

### 46.3　使用性能基准比较工具
2. benchstat
为了提高对性能基准数据比较的科学性，Go核心团队又开发了benchstat这款工具以替代benchcmp。

### 46.4　排除额外干扰，让基准测试更精确

### 小结
无论你是否认为性能很重要，都请你为被测代码（尤其是位于系统关键业务路径上的代码）建立性能基准。如果你编写的是供其他人使用的软件包，则更应如此。只有这样，我们才能至少保证后续对代码的修改不会带来性能回退。已经建立的性能基准可以为后续是否进一步优化的决策提供数据支撑，而不是靠程序员的直觉。
本条要点：
◦  性能基准测试在Go语言中是“一等公民”，在Go中我们可以很容易为被测代码建立性能基准；
◦  了解Go的两种性能基准测试的执行原理；
◦  使用性能比较工具协助解读测试结果数据，优先使用benchstat工具；
◦  使用testing.B提供的定时器操作方法排除额外干扰，让基准测试更精确，但不要在Run-Parallel中使用ResetTimer、StartTimer和StopTimer，因为它们具有全局副作用。


## 第47条使用pprof对程序进行性能剖析 !!!!
TODO

## 第48条使用expvar输出度量数据，辅助定位性能瓶颈点

### 小结
在本条中，我们学习了如何使用Go标准库的expvar包输出应用程序内省数据来辅助定位应用性能瓶颈点。expvar包不仅可用于辅助缩小定位性能瓶颈的范围，还可以用来输出度量数据以对应用的运行状态进行监控，这样当程序出现问题时，我们可以快速发现问题并利用输出的度量数据对程序进行诊断并快速定位问题。

### 本条要点：
◦  将应用内部状态以度量指标数据的形式输出，可以帮助我们最大限度地缩小性能瓶颈的搜索范围并快速定位瓶颈点；
◦  了解expvar包的工作原理；
◦  使用expvar包提供的内置类型实现应用要输出的度量指标；
◦  通过expvarmon等第三方工具快速展示应用内部状态信息。

## 第49条使用Delve调试Go代码

# 第九部分 标准库、反射与cgo

## 第50条理解Go TCP Socket网络编程模型
TODO

### 小结
在这一条中，我们学习了常见的网络I/O模型，了解了Go基于非阻塞Socket+I/O多路复用模型的网络编程模型的优点，包括降低通信复杂性，大幅减轻开发者的心智负担等，最后通过实例说明了在Go网络编程模型下，建立TCP连接、Socket读写（包括并发读写）、Socket属性设置及关闭连接的行为特点和注意事项。


## 第51条使用net/http包实现安全通信

### 小结
在本条中，我们了解了如何利用Go标准库提供的net/http、crypto/tls及crypto/x509等包建立一条安全的HTTPS协议通信通道。

### 本条要点：
◦  了解HTTP协议的优点与不足；
◦  了解HTTPS协议安全传输层的建立过程；
◦  理解非对称加密体系以及数字证书的组成与功用；
◦  数字证书就是使用CA私钥对证书申请者的公钥和证书相关信息进行签名后的满足标准证书格式的信息；
◦  了解如何使用Go实现对服务端和客户端证书的双向校验。


## 第52条掌握字符集的原理和字符编码方案间的转换

### 小结
在本条中，我们学习了Go默认字符集Unicode以及采用的编码方案UTF-8，深入理解了字符、字符集的属性——码点和内存编码表示（位模式）以及它们之间的关系，并通过实例讲解了如何利用Go标准库及扩展包实现不同字符编码方案间的转换。

## 第53条掌握使用time包的正确方式
### 53.3　定时器的使用 !!!!

## 第54条不要忽略对系统信号的处理

### 54.3　使用系统信号实现程序的优雅退出

+ 与优雅退出对立的是强制退出，也就是我们常说的使用kill -9，即kill -s SIGKILL pid。

```go
// chapter9/sources/go-signal/go-program-exit-gracefully-with-notify.go...
func main() {    
    var wg sync.WaitGroup    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {        
        fmt.Fprintf(w, "Hello, Signal!\n")    
    })    
    var srv = http.Server{        
        Addr: "localhost:8080",    
    }    
    srv.RegisterOnShutdown(func() {        
        // 在一个单独的goroutine中执行        
        fmt.Println("clean resources on shutdown...")        
        time.Sleep(2 * time.Second)        
        fmt.Println("clean resources ok")        
        wg.Done()    
    })    
    wg.Add(2)    
    go func() {        
        quit := make(chan os.Signal, 1)        
        signal.Notify(quit, syscall.SIGINT,           syscall.SIGTERM,           syscall.SIGQUIT,           syscall.SIGHUP)        
        <-quit        
        timeoutCtx, cf := context.WithTimeout(context.Background(),            time.Second*5)        
        defer cf()        
        var done = make(chan struct{}, 1)        
        go func() {            
            if err := srv.Shutdown(timeoutCtx); err != nil {                
                fmt.Printf("web server shutdown error: %v", err)            
            } else {                
                fmt.Println("web server shutdown ok")            
            }            
            done <- struct{}{}            
            wg.Done()        
        }()        
        select {        
            case <-timeoutCtx.Done():            
            fmt.Println("web server shutdown timeout")        
            case <-done:        
        }    
    }()    
    err := srv.ListenAndServe()    
    if err != nil {        
        if err != http.ErrServerClosed {            
            fmt.Printf("web server start failed: %v\n", err)            return        
        }    
    }    
    wg.Wait()    
    fmt.Println("program exit ok")
}
```


## 第55条使用crypto下的密码学包构建安全应用

### 小结
本条讲解了密码学中常用的几种工具以及在Go中对应的实现包与使用方法。
◦  对称密码（分组密码）：解决数据机密性的问题。
◦  公钥密码：解决密钥分发的问题。
◦  单向散列函数：解决消息完整性检查问题。
◦  消息认证码：可以识别伪装者。
◦  数字签名：解决消息究竟是谁所发的问题，防止否认。
◦  随机数：密码学建构的基础。

## 第56条掌握bytes包和strings包的基本操作

TODO

## 第57条理解标准库的读写模型

### 57.1　直接读写字节序列

+ 使用os.OpenFile创建并打开文件，传入的os.O_APPEND表示采用追加模式打开文件并写入数据

### 57.2　直接读写抽象数据类型实例

1. 利用fmt.Fscan和fmt.Fprint系列函数进行读写

2. 利用binary.Read和binary.Write函数进行读写

3. 利用gob包的Decode和Encode方法进行读写

### 57.3　通过包裹类型读写数据

### 小结
抽象数据类型实例与字节序列间的编解码方案除了gob外，还可以使用标准库提供的json和xml等。鉴于篇幅有限，这里就不详细展开了。
本条要点：
◦  Go标准库的读写模型以io.Reader和io.Writer接口为中心；
◦  模型既可以直接读写字节序列数据，也可以直接读写抽象数据类型实例；
◦  本质上，抽象数据类型实例的读写也会被转换为字节序列，只不过这种转换由Go标准库的包代劳了；
◦  通过包裹函数返回的包裹类型，我们可以轻松实现对读取或写入数据的缓冲、变换等处理。这种模式在标准库中有广泛应用。

## 第58条掌握unsafe包的安全使用模式

### 小结
作为最初以系统编程语言为目标的语言，Go为了兼顾性能以及低级代码操作，在其安全类型的保护盾下开了一个“后门”。在大多数情况下，这是Go核心团队自用的机制。我们要想使用unsafe包，就必须遵循unsafe包，尤其是unsafe.Pointer的安全使用规则。
本条要点：
◦  Go语言在常规操作下是类型安全的，但使用unsafe包可以“刺透”Go的类型安全保护层；
◦  Go兼容性并不包含对unsafe包的任何承诺，因此除非必要，尽量不要使用unsafe包，尤其是unsafe.Pointer；
◦  uintptr仅仅是一个整型值，即便它存储的是内存对象的地址值，它对内存对象也起不到引用的作用；
◦  使用unsafe包前，请先牢记并理解unsafe.Pointer的六条安全使用模式；
◦  如果使用了unsafe包，请使用go vet等工具对代码进行unsafe包使用合规性的检查。

## 第59条谨慎使用reflect包提供的反射能力

TODO

### 小结
reflect包所提供的Go反射能力是一把“双刃剑”，它既可以被用于优雅地解决一类特定的问题，但也会带来逻辑不清晰、性能问题以及难于发现问题和调试等困惑。因此，我们应谨慎使用这种能力，在做出使用的决定之前，认真评估反射是不是问题的唯一解决方案；在确定要使用反射能力后，也要遵循上述三个反射法则的要求。

## 第60条了解cgo的原理和使用开销

# 第十部分 工具链与工程实践

## 第61条使用module管理包依赖

TODO

## 第62条 构建最小Go程序容器镜像



TODO

## 第63条 自定义Go包的导入路径
### 63.1　govanityurls
### 63.2　使用govanityurls1. 安装govanityurls
### 小结
在这一条中，我们了解到自定义包导入路径具有诸多优点（如通过权威导入路径减少对包用户的影响、便于管理、路径简短等），并学习了一种基于govanityurls实现的自定义包导入路径的可行方案。该方案支持通过HTTPS访问并支持获取私有module。

使用 https://github.com/goproxyio/goproxy.git 应该更好


## 第64条 熟练掌握Go常用工具 !!!

TODO

### 64.3　构建

1. -x -v：让构建过程一目了然
-v用于输出当前正在编译的包，而-x则用于输出go build执行的每一个命令。

go build执行命令的顺序大致如下：
1）创建用于构建的临时目录；
2）下载构建module s依赖的module t和u；
3）分别编译module t和u，将编译后的结果存储到临时目录及GOCACHE目录下；
4）编译module s；
5）定位和汇总module s的各个依赖包构建后的目标文件（.a文件）的位置，形成importcfg.link文件，供后续链接器使用；
6）链接成可执行文件；
7）清理临时构建环境。

## 第65条使用go generate驱动代码生成

### 65.1　go generate：Go原生的代码生成“驱动器”

### 小结
go generate这个工具通常是由Go包的作者使用和执行的，其生成的Go源码一般会提交到代码仓库中，这个过程对生成的包的使用者来说是透明的。为了提醒使用者这是一个代码自动生成的源文件，我们通常会在源文件的开始处以注释的形式写入类似下面的文字：
// Code generated by XXX. DO NOT EDIT.

### 本条要点：
◦  尽量使用Go原生的go generate驱动代码生成；
◦  明确go generate应在go build、go run或go test等命令之前执行；
◦  go generate不会按照Go语法解析源文件，它只是将Go源码文件当成普通文本读取并识别其中的go generate指示符；
◦  go generate多用于生成枚举常量类型的String方法、protobuf文件对应的Go源文件，以及将静态资源文件数据嵌入二进制可执行文件中等场景；
◦  go generate多数情况仅被Go包的作者使用，对Go包的使用者透明。

## 第66条牢记Go的常见“陷阱” !!!
1. 短变量声明相关的“坑”
（1）短变量声明不总是会声明一个新变量
（2）短变量声明会导致难于发现的变量遮蔽

2. nil相关的“坑”

（1）不是所有以nil作为零值的类型都是零值可用的
这句话读起来有些拗口，我们可以将其分成两部分来理解。
◦  以nil为零值的类型：根据Go语言规范，诸如切片（slice）、map、接口类型和指针类型的零值均为nil。
◦  零值可用的类型：在第11条中，我们学习过什么是零值可用的类型，常见的有sync.Mutex和bytes.Buffer等。Go原生的切片类型只在特定使用方式下才可以被划到零值可用的范畴。

（2）值为nil的接口类型变量并不总等于nil

3. for range相关的“坑”

（1）你得到的是序号值而不是元素值
（2）针对string类型的for range迭代不是逐字节迭代
（3）对map类型内元素的迭代顺序是随机的
（4）在“复制品”上进行迭代
（5）迭代变量是重用的

4. 切片相关的“坑”

（1）对内存的过多占用
（2）隐匿数据的暴露与切片数据篡改
（3）新切片与原切片底层存储可能会“分家”

 5. string相关的“坑”
 6. switch语句相关的“坑”
 7. goroutine相关的“坑”
 （1）无法得到goroutine退出状态
 （2）程序随着main goroutine退出而退出，不等待其他goroutine
（3）任何一个goroutine出现panic，如果没有及时捕获，那么整个程序都将退出

+ 那么如何避免呢？没有好办法，只能采用防御型代码，即在每个goroutine的启动函数中加上对panic的捕获逻辑。对上面的示例的改造如下：
```go
// chapter10/sources/go-trap/goroutine_5.go

func safeRun(g func()) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("caught a panic:", e)
		}
	}()
	g()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	println("main goroutine: start to work...")
	go safeRun(func() {
		defer wg.Done()
		println("goroutine1: start to work...")
		time.Sleep(5 * time.Second)
		println("goroutine1: work done!")
	})
	go safeRun(func() {
		defer wg.Done()
		println("goroutine2: start to work...")
		time.Sleep(1 * time.Second)
		panic("division by zero")
		println("goroutine2: work done!")
	})
	wg.Wait()
	println("main goroutine: work done!")
}
```

8. channel相关的“坑”

+ 日常进行Go开发时，我们一般面对的都是有效状态（已初始化，尚未关闭）下的channel实例，但channel还有另外两种特殊状态：
◦  零值channel（nil channel）；
◦  已关闭的channel（closed channel）。
Go新手面对这两种特殊状态下的channel极易掉入“坑”中。为了避免掉“坑”，建议牢记这两种状态下的channel行为特征，见表66-1。
表66-1　两种特殊状态下的channel行为特征

9. 方法相关的“坑”

（1）使用值类型receiver的方法无法改变类型实例的状态
Go语言的方法（method）很独特，除了参数和返回值，它还拥有一个代表着类型实例的receiver。receiver有两类：值类型receiver和指针类型receiver。而采用值类型receiver的方法无法改变类型实例的状态。

（2）值类型实例可以调用采用指针类型receiver的方法，指针类型实例也可以调用采用值类型receiver的方法

10. break语句相关的“坑”

### 66.2　标准库类
1. time包相关的“坑”
2. encoding/json包相关的“坑”
（1）未导出的结构体字段不会被编码到JSON文本中
（2）nil切片和空切片可能被编码为不同文本
（3）字节切片可能被编码为base64编码的文本
（4）当JSON文本中的整型数值被解码为interface{}类型时，其底层真实类型为float64

3. net/http包相关的“坑”
（1）http包需要我们手动关闭Response.Body
（2）HTTP客户端默认不会及时关闭已经用完的HTTP连接

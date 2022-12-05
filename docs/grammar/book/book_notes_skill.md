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

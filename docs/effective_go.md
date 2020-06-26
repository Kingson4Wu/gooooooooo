+ https://learnku.com/docs/effective-go/2020

+ Go 包的源码 不仅是核心库，同时也是学习如何使用 Go 语言的示例源码。 此外，其中的一些包还包含了可独立的可执行示例

+  Go 官方翻译：Go 常见问题（Go FAQ）：<https://learnku.com/go/wikis/38175> (可反复看！！！)

+ 社区 Wiki:<https://learnku.com/go/wikis>
    - 编程规范和设计原则：<https://learnku.com/go/wikis/38426> (后续看)


### 格式化
+ go fmt命令只会格式化被直接保存在指定代码包对应目录下的Go语言源码文件
+ 使用gofmt格式化代码:<https://zhuanlan.zhihu.com/p/39928126>   
### 代码注释
+ godoc 既是一个程序，又是一个 Web 服务器，它对 Go 的源码进行处理，并提取包中的文档内容。 出现在顶级声明之前，且与该声明之间没有空行的注释，
将与该声明一起被提取出来，作为该条目的说明文档。 这些注释的类型和风格决定了 godoc 生成的文档质量。
### 命名规则
+ 某个名称在包外是否可见，就取决于其首个字符是否为大写字母
+ 按照约定，只包含一个方法的接口应当以该方法的名称加上 - er 后缀来命名，
+ 驼峰 命名: Go 中的约定是使用 MixedCaps 或 mixedCaps 而不是下划线来编写多个单词组成的命名  
### 分号
+ 无论如何，你都不应将一个控制结构（if、for、switch 或 select）的左大括号放在下一行。如果这样做，就会在大括号前面插入一个分号，这可能引起不需要的效果。
### 控制结构
+ <https://learnku.com/docs/effective-go/2020/control-structure/6241>
+  Go 不再使用 do 或 while 循环，只有一个更通用的 for；switch 要更灵活一点；if 和 switch 像 for 一样可接受可选的初始化语句； 
此外，还有一个包含类型选择和多路通信复用器的新控制结构：select。 其语法也有些许不同：没有圆括号，而其主体必须始终使用大括号括住。
+ 由于 if 和 switch 可接受初始化语句， 因此用它们来设置局部变量十分常见。
```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```
+ 以 return (break、continue、goto) 结束， 之后的代码也就无需 else 了  
<pre>
在满足下列条件时，已被声明的变量 v 可出现在:= 声明中：

本次声明与已声明的 v 处于同一作用域中（若 v 已在外层作用域中声明过，则此次声明会创建一个新的变量 §），
在初始化中与其类型相应的值才能赋予 v，且
在此次声明中至少另有一个变量是新声明的。
这个特性简直就是纯粹的实用主义体现，它使得我们可以很方便地只使用一个 err 值，例如，在一个相当长的 if-else 语句链中， 你会发现它用得很频繁。

§ 值得一提的是，即便 Go 中的函数形参和返回值在词法上处于大括号之外， 但它们的作用域和该函数体仍然相同。

</pre>
<pre>
For

// 类似 C 语言中的 for 用法
for init; condition; post { }

// 类似 C 语言中的 while 用法
for condition { }

// 类似 C 语言中的 for(;;) 用法
for { }
</pre>
<pre>
switch 并不会自动下溯，但 case 可通过逗号分隔来列举相同的处理条件。

func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}

</pre>
+ 在 Go 中，我们只需将标签放置到循环外，然后 “蹦” 到那里即可。
<pre>
类型选择
switch 也可用于判断接口变量的动态类型。如 类型选择 通过圆括号中的关键字 type 使用类型断言语法。若 switch 在表达式中声明了一个变量，那么该变量的每个子句中都将有该变量对应的类型。

var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T 打印任何类型的 t
case bool:
    fmt.Printf("boolean %t\n", t)             // t 是 bool 类型
case int:
    fmt.Printf("integer %d\n", t)             // t 是 int 类型
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
}

</pre>
### 函数
+ 多返回值
+ 命名结果参数！！！！
    - Go 函数的返回值或结果 “形参” 可被命名，并作为常规变量使用，就像传入的形参一样。 命名后，一旦该函数开始执行，它们就会被初始化为与其类型相应的零值； 
    若该函数执行了一条不带实参的 return 语句，则结果形参的当前值将被返回。
```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}

```    
+ 延迟 defer
    - defer 语句用于预设一个函数调用（即推迟执行函数）， 该函数会在执行 defer 的函数返回之前立即执行。它显得非比寻常， 
    但却是处理一些事情的有效方式，例如无论以何种路径返回，都必须释放资源的函数。 典型的例子就是解锁互斥和关闭文件。

### 数据 !!! (可反复看！)
+ 分配内存 （Go 提供了两种分配原语，即内建函数 new 和 make
+ https://learnku.com/docs/effective-go/2020/data/6243

<pre>
下面的例子阐明了 new 和 make 之间的区别：

var p *[]int = new([]int)       // 分配切片结构；*p == nil；很少用到
var v  []int = make([]int, 100) // 切片 v 现在引用了一个具有 100 个 int 元素的新数组

// 没必要的复杂用法:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 常规用法:
v := make([]int, 100)
请记住，make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用 new 分配内存。

</pre>

#### 数组和切片

<pre>
以下为数组在 Go 和 C 中的主要区别。在 Go 中，

数组是值。将一个数组赋予另一个数组会复制其所有元素。
特别地，若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针。
数组的大小是其类型的一部分。类型 [10]int 和 [20]int 是不同的。
数组为值的属性很有用，但代价高昂；若你想要 C 那样的行为和效率，你可以传递一个指向该数组的指针。

</pre>       
+ 数组其实在 Go 语言中没有那么常用，更加常见的数据结构其实是切片，切片其实就是动态数组，它的长度并不固定，可以追加元素并会在切片容量不足时进行扩容。
+ (切片)扩容其实就是需要为切片分配一块新的内存空间，分配内存空间之前需要先确定新的切片容量，Go 语言根据切片的当前容量选择不同的策略进行扩容：
    1.如果期望容量大于当前容量的两倍就会使用期望容量；
    2. 如果当前切片容量小于 1024 就会将容量翻倍；
    3. 如果当前切片容量大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；
+ 数组的大多数操作在 编译期间 都会转换成对内存的直接读写；而切片的很多功能就都是在运行时实现的了，无论是初始化切片，还是对切片进行追加或扩容都需要运行时的支持，
需要注意的是在遇到大切片扩容或者复制时可能会发生大规模的内存拷贝，一定要在使用时减少这种情况的发生避免对程序的性能造成影响。
+ Go语言数组和切片的原理：<cnblogs.com/itbsl/p/10599948.html>

+ 切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。
 若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 这可以理解为传递了底层数组的指针。(!!!!!!!!!!)

#### 二维切片
<pre>

Go 的数组和切片都是一维的。要创建等价的二维数组或切片，就必须定义一个数组的数组， 或切片的切片，就像这样：

type Transform [3][3]float64  // 一个 3x3 的数组，其实是包含多个数组的一个数组。
type LinesOfText [][]byte     // 包含多个字节切片的一个切片。

</pre> 
+ 是切面还是数组，取决于声明方式！
<pre>
 以下是这两种方法的大概代码，仅供参考。首先是一次一行的：

// 分配底层切片.
picture := make([][]uint8, YSize) // y每一行的大小
//循环遍历每一行
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
现在是一次分配，对行进行切片：

// 分配底层切片
picture := make([][]uint8, YSize) //  每 y 个单元一行。
// 分配一个大一些的切片以容纳所有的元素
pixels := make([]uint8, XSize*YSize) // 指定类型[]uint8, 即便图片是 [][]uint8.
//循环遍历图片所有行，从剩余像素切片的前面对每一行进行切片。
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}

</pre>

#### 映射
+ 映射是方便而强大的内建数据结构，它可以关联不同类型的值。其键可以是任何相等性操作符支持的类型， 如整数、浮点数、复数、字符串、指针、接口（只要其动态类型支持相等性判断）、
结构以及数组。 切片不能用作映射键，因为它们的相等性还未定义。与切片一样，映射也是引用类型。 若将映射传入函数中，并更改了该映射的内容，则此修改对调用者同样可见。

+ 与切片一样，映射也是引用类型。!!!

#### 打印
<pre>
以下示例中各行产生的输出都是一样的。

fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))

</pre>  
+ 映射中的键可能按任意顺序输出。当打印结构体时，改进的格式 %+v 会为结构体的每个字段添上字段名，而另一种格式 %#v 将完全按照 Go 的语法打印值
+ （请注意其中的 & 符号）当遇到 string 或 []byte 值时， 可使用 %q 产生带引号的字符串；而格式 %#q 会尽可能使用反引号。 （%q 格式也可用于整数和符文，它会产生一个带单引号的符文常量。） 此外，%x 还可用于字符串、字节数组以及整数，并生成一个很长的十六进制字符串， 而带空格的格式（% x）还会在字节之间插入空格。
+ 另一种实用的格式是 %T，它会打印某个值的类型。

#### 追加

### 初始化
#### 常量
```go
type ByteSize float64

const (
    _           = iota // 通过赋予空白标识符来忽略第一个值
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

```
#### 变量
+ 能像常量一样初始化，而且可以初始化为一个可在运行时得出结果的普通表达式。
```go
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```
#### init 函数
+ 每个源文件都可以通过定义自己的无参数 init 函数来设置一些必要的状态。

### 方法
+ 若我们将函数修改为与标准 Write 类似的方法，就像这样，
```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // 同上。
    *p = slice
    return len(data), nil
}
```
+ 以指针或值为接收者的区别在于：值方法可通过指针和值调用， 而指针方法只能通过指针来调用。

之所以会有这条规则是因为指针方法可以修改接收者；通过值调用它们会导致方法接收到该值的副本， 因此任何修改都将被丢弃，因此该语言不允许这种错误。不过有个方便的例外：若该值是可寻址的， 那么该语言就会自动插入取址操作符来对付一般的通过值调用的指针方法。在我们的例子中，变量 b 是可寻址的，因此我们只需通过 b.Write 来调用它的 Write 方法，编译器会将它重写为 (&b).Write。

顺便一提，在字节切片上使用 Write 的想法已被 bytes.Buffer 所实现

### 接口与其它类型
#### 类型转换
+ 通过类型转换技术，在 String 方法中安全调用 Sprintf 的另个一例子。若我们忽略类型名的话，这两种类型（Sequence 和 []int）其实是相同的，因此在二者之间进行转换是合法的。
```go
type Sequence []int

func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}

```
#### 接口转换与类型断言
```go
type Stringer interface {
    String() string
}

var value interface{} // Value 由调用者提供
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```
等价于

```go
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}

```
+ 若我们只关心一种类型呢？若我们知道该值拥有一个 string 而想要提取它呢？ 只需一种情况的类型选择就行，但它需要类型断言。类型断言接受一个接口值， 并从中提取指定的明确类型的值。其语法借鉴自类型选择开头的子句，但它需要一个明确的类型， 而非 type 关键字：
但若它所转换的值中不包含字符串，该程序就会以运行时错误崩溃。为避免这种情况， 需使用 “逗号，ok” 惯用测试它能安全地判断该值是否为字符串
```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

#### 通用性 !!!
+ 在本节中，我们通过一个结构体，一个整数，一个信道和一个函数，建立了一个 HTTP 服务器， 这一切都是因为接口只是方法的集和，而几乎任何类型都能定义方法。

### 空白标识符
#### 多个参数赋值中的空白标识符
#### 未使用的导入和变量
+ 若导入某个包或声明某个变量而不使用它就会产生错误。未使用的包会让程序膨胀并拖慢编译速度， 而已初始化但未使用的变量不仅会浪费计算能力，还有可能暗藏着更大的 Bug。 然而在程序开发过程中，经常会产生未使用的导入和变量。虽然以后会用到它们， 但为了完成编译又不得不删除它们才行，这很让人烦恼。空白标识符就能提供一个工作空间。
要让编译器停止关于未使用导入的包，需要空白标识符来引用已导入包中的符号。 同样，将未使用的变量 fd 赋予空白标识符也能关闭未使用变量错误。 该程序的以下版本可以编译。
```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf  // 用于调试，结束时删除。
var _ io.Reader    // 用于调试，结束时删除。

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}

```
#### 为辅助作用而导入
+ 只为了其辅助作用来导入该包， 只需将包重命名为空白标识符：
`import _ "net/http/pprof"`
这种导入格式能明确表示该包是为其辅助作用而导入的，因为没有其它使用该包的可能： 在此文件中，它没有名字。（若它有名字而我们没有使用，编译器就会拒绝该程序。）
#### 接口检查 !!
<pre>
尽管有些接口检查会在运行时进行。encoding/json 包中就有个实例它定义了一个 Marshaler 接口。当 JSON 编码器接收到一个实现了该接口的值，那么该编码器就会调用该值的编组方法， 将其转换为 JSON，而非进行标准的类型转换。 编码器在运行时通过类型断言检查其属性，就像这样：

m, ok := val.(json.Marshaler)
若只需要判断某个类型是否是实现了某个接口，而不需要实际使用接口本身 （可能是错误检查部分），就使用空白标识符来忽略类型断言的值：

if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
当需要确保某个包中实现的类型一定满足该接口时，就会遇到这种情况。 若某个类型（例如 json.RawMessage） 需要一种自定义的 JSON 表现时，它应当实现 json.Marshaler， 不过现在没有静态转换可以让编译器去自动验证它。若该类型通过忽略转换失败来满足该接口， 那么 JSON 编码器仍可工作，但它却不会使用自定义的实现。为确保其实现正确， 可在该包中用空白标识符声明一个全局变量：

var _ json.Marshaler = (*RawMessage)(nil)
在此声明中，我们调用了一个 *RawMessage 转换并将其赋予了 Marshaler，以此来要求 *RawMessage 实现 Marshaler，这时其属性就会在编译时被检测。 若 json.Marshaler 接口被更改，此包将无法通过编译， 而我们则会注意到它需要更新。

在这种结构中出现空白标识符，即表示该声明的存在只是为了类型检查。 不过请不要为满足接口就将它用于任何类型。作为约定， 只有当代码中不存在静态类型转换时才能使用这种声明，毕竟这是种非常罕见的情况。

</pre>

###  内嵌
+ 有种区分内嵌与子类的重要手段。当内嵌一个类型时，该类型的方法会成为外部类型的方法， 但当它们被调用时，该方法的接收者是内部类型，而非外部的。
<pre>
内嵌同样可以提供便利。这个例子展示了一个内嵌字段和一个常规的命名字段。

type Job struct {
    Command string
    *log.Logger
}
Job 类型现在有了 Log、Logf 和 *log.Logger 的其它方法。我们当然可以为 Logger 提供一个字段名，但完全不必这么做。现在，一旦初始化后，我们就能记录 Job 了：

job.Println("starting now...")
Logger 是 Job 结构体的常规字段， 因此我们可在 Job 的构造函数中，通过一般的方式来初始化它，就像这样：

func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
或通过复合字面：

job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}

</pre>
+ 内嵌类型会引入命名冲突的问题，但解决规则却很简单。首先，字段或方法 X 会隐藏该类型中更深层嵌套的其它项 X。若 log.Logger 包含一个名为 Command 的字段或方法，Job 的 Command 字段会覆盖它。


### 并发
#### 通过通信共享内存
+  Go 语言另辟蹊径，它将共享的值通过信道传递，实际上，多个独立执行的线程从不会主动共享。 在任意给定的时间点，只有一个 Go 协程能够访问该值。
数据竞争从设计上就被杜绝了。 为了提倡这种思考方式，我们将它简化为一句口号：
不要通过共享内存来通信，而应通过通信来共享内存。
+ Unix 管道就与这种模型完美契合。 尽管 Go 的并发处理方式来源于 Hoare 的通信顺序处理（CSP）， 它依然可以看做是类型安全的 Unix 管道的实现。
#### 协程（goroutine）
+ 在 Go 中，字面函数都是闭包：其实现在保证了函数内引用变量的生命周期与函数的活动时间相同。
#### 信道
+ 若信道是不带缓冲的，那么在接收者收到值前， 发送者会一直阻塞；若信道是带缓冲的，则发送者仅在值被复制到缓冲区前阻塞； 若缓冲区已满，发送者会一直等待直到某个接收者取出一个值为止。
<pre>
它的写法看起来有点奇怪

req := req
但在 Go 中这样做是合法且常见的。你用相同的名字获得了该变量的一个新的版本， 以此来局部地刻意屏蔽循环变量，使它对每个 Go 协程保持唯一。
</pre>
+  runtime.GOMAXPROCS，会返回用户设置可用 CPU 数量。默认情况下使用  runtime.NumCPU 的值，但是可以被命令行环境变量，或者调用此函数并传参正整数。传参 0 的话会返回值，假如说我们尊重用户对资源的分配，就应该这么写：

var numCPU = runtime.GOMAXPROCS(0)
+ 尽管 Go 的并发特性能够让某些问题更易构造成并行计算， 但 Go 仍然是种并发而非并行的语言，且 Go 的模型并不适合所有的并行问题。

###  错误
#### Panic
+ 向调用者报告错误的一般方式就是将 error 作为额外的值返回。 标准的 Read 方法就是个众所周知的实例，它返回一个字节计数和一个 error。但如果错误时不可恢复的呢？有时程序就是不能继续运行。

为此，我们提供了内建的 panic 函数，它会产生一个运行时错误并终止程序 （但请继续看下一节）。该函数接受一个任意类型的实参（一般为字符串），并在程序终止时打印。 它还能表明发生了意料之外的事情，比如从无限循环中退出了。
#### recover
+ 当 panic 被调用后（包括不明确的运行时错误，例如切片越界访问或类型断言失败）， 程序将立刻终止当前函数的执行，并开始回溯 Go 协程的栈，运行任何被推迟的函数。 若回溯到达 Go 协程栈的顶端，程序就会终止。不过我们可以用内建的 recover 函数来重新或来取回 Go 协程的控制权限并使其恢复正常执行。


----
### 指针和引用
+ &取得变量的地址
+ *取得指针变量指向的内存地址的值
+ Go语言中的引用类型有：映射（map），数组切片（slice），通道（channel），方法与函数。
+ new函数与&操作符
  Go语言中提供两种创建变量的方式，同时可以获得指向它们的指针：new函数与&操作符
<pre>
type Person struct {
   name string
   sex  string
   age int
}
func main() {
   person1 := Person{"zhangsan","man",25} //创建一个person1对象
   person2 := new(Person)//使用new创建一个person2对象，同时获得person的指针
   person2.name,person2.sex,person2.age = "wangwu","man",25
   person3 := &Person{"lisi","man",25}//使用&创建一个person3对象，同时获得person的指针
   fmt.Printf("person1:%v, person2:%v, person3:%v\n",person1,person2,person3)
}
-----output-----
person1:{zhangsan man 25}, person2:&{wangwu man 25}, person3:&{lisi man 25}
</pre>    
+ <https://www.php.cn/be/go/439923.html>
+ 指针变量存储的是另一个变量的地址。
  引用变量指向另外一个变量。
+ 指针和引用之间的第二个关键区别。指针可以重分配，而引用不能。换句话说，指针可以被分配另一个不同的地址。  



       
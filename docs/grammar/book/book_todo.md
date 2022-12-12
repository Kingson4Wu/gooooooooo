### go语言圣经

+ 5.6.1. 警告：捕获迭代变量
本节，将介绍Go词法作用域的一个陷阱。请务必仔细的阅读，弄清楚发生问题的原因。即使是经验丰富的程序员也会在这个问题上犯错误。
考虑这样一个问题：你被要求首先创建一些目录，再将目录删除。在下面的例子中我们用函数值来完成删除操作。下面的示例代码需要引入os包。为了使代码简单，我们忽略了所有的异常处理。

+ 7.12. 通过类型断言询问行为
+ 下面这段逻辑和net/http包中web服务器负责写入HTTP头字段（例如："Content-type:text/html"）的部分相似。io.Writer接口类型的变量w代表HTTP响应；写入它的字节最终被发送到某个人的web浏览器上。


### Go语言精进之路

看到运行时实现转换的函数中已经加入了一些避免每种情况都要分配新内存操作的优化（如tmpBuf的复用）。

// chapter4/sources/method_nature_3.gotype field struct {    name string}func (p *field) print() {    fmt.Println(p.name)}func main() {    data1 := []*field{{"one"}, {"two"}, {"three"}}    for _, v := range data1 {        go v.print()    }    data2 := []field{{"four"}, {"five"}, {"six"}}    for _, v := range data2 {        go v.print()    }    time.Sleep(3 * time.Second)}


// chapter4/sources/method_set_utils.gofunc DumpMethodSet(i interface{}) {    v := reflect.TypeOf(i)    elemTyp := v.Elem()    n := elemTyp.NumMethod()    if n == 0 {        fmt.Printf("%s's method set is empty!\n", elemTyp)        return    }    fmt.Printf("%s's method set:\n", elemTyp)    for j := 0; j < n; j++ {        fmt.Println("-", elemTyp.Method(j).Name)    }    fmt.Printf("\n")}
接下来，我们就用该工


// chapter4/sources/variadic_function_9.gotype FinishedHouse struct {    style                  int    // 0: Chinese; 1: American; 2: European    centralAirConditioning bool   // true或false    floorMaterial          string  // "ground-tile"或"wood"    wallMaterial           string // "latex"或"paper"或"diatom-mud"}type Option func(*FinishedHouse)func NewFinishedHouse(options ...Option) *FinishedHouse {    h := &FinishedHouse{        // default options        style:                  0,        centralAirConditioning: true,        floorMaterial:          "wood",        wallMaterial:           "paper",    }    for _, option := range options {        option(h)    }    return h}func WithStyle(style int) Option {    return func(h *FinishedHouse) {        h.style = style    }}func WithFloorMaterial(material string) Option {    return func(h *FinishedHouse) {        h.floorMaterial = material    }}func WithWallMaterial(material string) Option {    return func(h *FinishedHouse) {        h.wallMaterial = material    }}func WithCentralAirConditioning(centralAirConditioning bool) Option {    return func(h *FinishedHouse) {        h.centralAirConditioning = centralAirConditioning    }}func main() {    fmt.Printf("%+v\n", NewFinishedHouse()) // 使用默认选项    fmt.Printf("%+v\n", NewFinishedHouse(WithStyle(1),        WithFloorMaterial("ground-tile"),        WithCentralAirConditioning(false)))}


3. 适配器函数类型
适配器函数类型（adapter function type）是一个辅助水平组合实现的“工具”类型。强调一下，它是一个类型。它可以将一个满足特定函数签名的普通函数显式转换成自身类型的实例，转换后的实例同时也是某个单方法接口类型的实现者。最典型的适配器函数类型莫过于第21条提到过的http.HandlerFunc了。
// $GOROOT/src/net/http/server.gotype Handler interface {    ServeHTTP(ResponseWriter, *Request)}type HandlerFunc func(ResponseWriter, *Request)func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {    f(w, r)}// chapter5/sources/horizontal-composition-3.gofunc greetings(w http.ResponseWriter, r *http.Request) {    fmt.Fprintf(w, "Welcome!")}func main() {    http.ListenAndServe(":8080", http.HandlerFunc(greetings))}
可以看到，在上述例子中通过http.HandlerFunc这个适配器函数类型，可以将普通函数greetings快速转换为实现了http.Handler接口的类型。转换后，我们便可以将其实例用作实参，实现基于接口的组合了。


我们还可以通过测试代码的文件名来区分所属测试类别，比如：net/http包就使用transport_internal_test.go这个名字来明确该测试文件采用包内测试的方法，而对应的transport_test.go则是一个采用包外测试的源文件。

这个地方想了一下，defer func1()是在结束当前函数调用后会运行func1，但是func1其实是setUp()的返回值，所以，确实是会先打印Setup，然后Execute,最后打印tearDown


f, err := os.Open(filepath.Join("testdata", "data-001.txt"))


小结
通过这一条，我们认识到模糊测试对于提升Go代码质量、挖掘潜在bug的重要作用。但模糊测试不是“银弹”，它有其适用的范围。模糊测试最适合那些处理复杂输入数据的程序，比如文件格式解析、网络协议解析、人机交互界面入口等。模糊测试是软件测试技术的一个重要分支，与单元测试等互为补充，相辅相成。
目前，并非所有编程语言都有对模糊测试工具的支持，Gopher和Go社区很幸运，Dmitry Vyukov为我们带来了go-fuzz模糊测试工具。如果你是追求高质量Go代码的开发者，请为你的Go代码建立起模糊测试。


通过Go原生提供的性能基准测试为被测对象建立性能基准了。但被测代码更新前后的性能基准比较依然要靠人工计算和肉眼比对，十分不方便。为此，Go核心团队先后开发了两款性能基准比较工具：benchcmp（https://github.com/golang/tools/tree/master/cmd/benchcmp）和benchstat（https://github.com/golang/perf/tree/master/benchstat）。

Go核心团队已经给benchcmp工具打上了“deprecation”（不建议使用）的标签，因此建议大家使用benchstat来进行性能基准数据的比较。


小结
无论你是否认为性能很重要，都请你为被测代码（尤其是位于系统关键业务路径上的代码）建立性能基准。如果你编写的是供其他人使用的软件包，则更应如此。只有这样，我们才能至少保证后续对代码的修改不会带来性能回退。已经建立的性能基准可以为后续是否进一步优化的决策提供数据支撑，而不是靠程序员的直觉。
本条要点：
◦  性能基准测试在Go语言中是“一等公民”，在Go中我们可以很容易为被测代码建立性能基准；
◦  了解Go的两种性能基准测试的执行原理；
◦  使用性能比较工具协助解读测试结果数据，优先使用benchstat工具；
◦  使用testing.B提供的定时器操作方法排除额外干扰，让基准测试更精确，但不要在Run-Parallel中使用ResetTimer、StartTimer和StopTimer，因为它们具有全局副作用。

通过top -cum，我们看到handleHi累积消耗CPU最多（用户层代码范畴）。通过list命令进一步展开handleHi函数


这里进行内存分配的优化方法如下：
◦  删除w.Header().Set这行调用；
◦  使用fmt.Fprintf替代w.Write。




要实现零内存分配，可以像下面这样优化代码：
// chapter8/sources/go-pprof-optimization-demo/step4/demo.go...
var visitors int64 // 必须被自动访问
var rxOptionalID = regexp.MustCompile(`^\d*$`)
var bufPool = sync.Pool{    
	New: func() interface{} {        
		return bytes.NewBuffer(make([]byte, 128))    
		},
	}
func handleHi(w http.ResponseWriter, r *http.Request) {    
	if !rxOptionalID.MatchString(r.FormValue("color")) {        
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)        
		return    
	}    
	visitNum := atomic.AddInt64(&visitors, 1)    
	buf := bufPool.Get().(*bytes.Buffer)    
	defer bufPool.Put(buf)    
	buf.Reset()    
	buf.WriteString("<h1 style='color: ")    
	buf.WriteString(r.FormValue("color"))    
	buf.WriteString("'>Welcome!</h1>You are visitor number ")    
	b := strconv.AppendInt(buf.Bytes(), visitNum, 10)    
	b = append(b, '!')    
	w.Write(b)
}
这里有几点主要优化：
◦  使用sync.Pool减少重新分配bytes.Buffer的次数；
◦  采用预分配底层存储的bytes.Buffer拼接输出；
◦  使用strconv.AppendInt将整型数拼接到bytes.Buffer中，



（1）充分的代码检查
充分利用你手头上的工具对你编写的代码进行严格的检查，这些工具包括编译器（尽可能将警告级别提升到你可以接受的最高级别）、静态代码检查工具（linter，如go vet）等。


// chapter9/sources/go-bytes-and-strings/string_and_bytes_reader.go...
func main() {
	var buf bytes.Buffer
	var s = "I love Go!!"
	_, err := io.Copy(&buf, strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf.String()) // "I love Go!!"
	buf.Reset()
	var b = []byte("I love Go!!")
	_, err = io.Copy(&buf, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf.String()) // "I love Go!!"
}
通过创建的strings.Reader或bytes.Reader新实例，我们就可以读取作为数据源的字符串或字节切片中的数据。



通过包裹函数返回的包裹类型，我们还可以实现对读出或写入数据的变换，比如压缩等。Go标准库中的compress/gzip包就提供了这样的包裹函数与包裹类型。我们看一个压缩数据并形成压缩文件的例子：
// chapter9/sources/go-read-and-write/gzip_compress_data.go...
func main() {    
	file := "./hello_gopher.gz"    
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)    
	...    
	defer f.Close()    
	zw := gzip.NewWriter(f)    
	defer zw.Close()    
	_, err = zw.Write([]byte("hello, gopher! I love golang!!"))    
	if err != nil {        
		fmt.Println("write compressed data error:", err)    
	}    
	fmt.Println("write compressed data ok")
}


（2）高效类型转换
使用unsafe包，Gopher可以绕开Go类型系统的安全检查，因此可以通过unsafe包实现性能更好的类型转换。最常见的类型转换是string与[]byte类型间的相互转换：
func Bytes2String(b []byte) string {    return *(*string)(unsafe.Pointer(&b))}func String2Bytes(s string) []byte {    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))    bh := reflect.SliceHeader{        Data: sh.Data,        Len:  sh.Len,        Cap:  sh.Len,    }    return *(*[]byte)(unsafe.Pointer(&bh))}
在Go中，string类型变量是不可变的（immutable），通过常规方法将一个string类型变量转换为[]byte类型，Go会为[]byte类型变量分配一块新内存，并将string类型变量的值复制到这块新内存中。而通过上面基于unsafe包实现的String2Bytes函数，这种转换并不需要额外的内存复制：转换后的[]byte变量与输入参数中的string类型变量共享底层存储（但注意，我们依旧无法通过对返回的切片的修改来改变原字符串）。而将[]byte变量转换为string类型则更为简单，因为[]byte的内部表示是一个三元组(ptr, len, cap)，string的内部表示为一个二元组(ptr, len)，通过unsafe.Pointer将[]byte的内部表示重新解释为string的内部表示，这就是Bytes2String的原理。


务必通过-gcflags '-l'关闭内联优化，这样才能得到公平的测试结果：
$go test -bench . -gcflags '-l'

$cat /proc/15829/status|grep -i threadThreads:    3

当Go编译器执行跨平台编译时，它会将CGO_ENABLED置为0，即关闭cgo

即便显式开启cgo，cgo调用的macOS上的外部链接器clang也会因无法识别目标平台的目标文件格式而报错，macOS上的clang默认并不具备跨平台编译Linux应用的能力。


3. -race：让并发bug无处遁形
-race命令行选项会在构建的结果中加入竞态检测的代码。在程序运行过程中，如果发现对数据的并发竞态访问，这些代码会给出警告，这些警告信息可以用来辅助后续查找和解决竞态问题。不过由于插入竞态检测的代码这个动作，带有-race的构建过程会比标准构建略慢一些。

Go社区的一个最佳实践是在正式发布到生产环境之前的调试、测试环节使用带有-race构建选项构建出的程序，以便于在正式发布到生产环境之前尽可能多地发现程序中潜在的并发竞态问题并快速将其解决。

以-m为例，我们可以通过下面的命令输出更为详尽的逃逸分析过程信息：
go build -gcflags='-m'go build -gcflags='-m -m'      // 输出比上一条命令更为详尽的逃逸分析过程信息go build -gcflags='-m=2'       // 与上一条命令等价go build -gcflags='-m -m -m'   // 输出最为详尽的逃逸分析过程信息go build -gcflags='-m=3'       // 与上一条命令等价

1）-X：设定包中string类型变量的值（仅支持string类型变量）。
通过-X选项，我们可以在编译链接期间动态地为程序中的字符串变量进行赋值，这个选项的一个典型应用就是在构建脚本中设定程序的版本值。我们通常会为应用程序添加version命令行标志选项，用来输出当前程序的版本信息

可以通过设置环境变量GODEBUG='gctrace=1'让位于Go程序中的运行时在每次GC执行时输出此次GC相关的跟踪信息。

2. 提交代码前请使用go vet对代码进行静态检查
如果编

//go:generate protoc -I ./IDL msg.proto --gofast_out=./msg
这就是预先“埋”在代码中的可以被go generate命令识别的指示符（directive）。当我们在示例的目录下执行go generate命令时，上面这行指示符中的命令将被go generate识别并被驱动执行，执行的结果就是protoc基于IDL目录下的msg.proto生成了main包所需要的msg包源码。


2. go generate驱动从静态资源文件数据到Go源码的转换
Go语言的优点之一是可以将源码编译成一个对外部没有任何依赖或只有较少依赖的二进制可执行文件，这大大降低了Gopher在部署阶段的心智负担。而为了将这一优势发挥到极致，Go社区甚至开始着手将静态资源文件也嵌入可执行文件中，尤其是在Web开发领域，Gopher希望将一些静态资源文件（比如CSS文件等）嵌入最终的二进制文件中一起发布和部署。而go generate结合go-bindata工具（https://github.com/go-bindata/go-bindata）常被用于实现这一功能。


那么如何避免呢？没有好办法，只能采用防御型代码，即在每个goroutine的启动函数中加上对panic的捕获逻辑。对上面的示例的改造如下：
// chapter10/sources/go-trap/goroutine_5.gofunc safeRun(g func()) {    defer func() {        if e := recover(); e != nil {            fmt.Println("caught a panic:", e)        }    }()    g()}func main() {    var wg sync.WaitGroup    wg.Add(2)    println("main goroutine: start to work...")    go safeRun(func() {        defer wg.Done()        println("goroutine1: start to work...")        time.Sleep(5 * time.Second)        println("goroutine1: work done!")    })    go safeRun(func() {        defer wg.Done()        println("goroutine2: start to work...")        time.Sleep(1 * time.Second)        panic("division by zero")        println("goroutine2: work done!")    })    wg.Wait()    println("main goroutine: work done!")}


（1）http包需要我们手动关闭Response.Body
通过http包我们很容易实现一个HTTP客户端，比如：
// chapter10/sources/go-trap/http_1.gofunc main() {    resp, err := http.Get("https://tip.golang.org")    if err != nil {        fmt.Println(err)        return    }    body, err := ioutil.ReadAll(resp.Body)    if err != nil {        fmt.Println(err)        return    }    fmt.Println(string(body))}
这个示例通过http.Get获取某个网站的页面内容，然后读取应答Body字段中的数据并输出到命令行控制台上。但仅仅这么做还不够，因为http包需要我们配合完成一项任务：务必关闭resp.Body。
// chapter10/sources/go-trap/http_1.goresp, err := http.Get("https://tip.golang.org")if err != nil {    fmt.Println(err)    return}defer resp.Body.Close()
目前http包的实现逻辑是只有当应答的Body中的内容被全部读取完毕且调用了Body.Close()，默认的HTTP客户端才会重用带有keep-alive标志的HTTP连接，否则每次HTTP客户端发起请求都会单独向服务端建立一条新的TCP连接，这样做的消耗要比重用连接大得多。
注：仅在作为客户端时，http包才需要我们手动关闭Response.Body；如果是作为服务端，http包会自动处理Request.Body。



```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())

    ch := func(ctx context.Context) <-chan int {
        ch := make(chan int)
        go func() {
            for i := 0; ; i++ {
                select {
                case <- ctx.Done():
                    return
                case ch <- i:
                }
            }
        } ()
        return ch
    }(ctx)

    for v := range ch {
        fmt.Println(v)
        if v == 5 {
            cancel()
            break
        }
    }
}
```

+ Go语言内置的并发能力也可以通过组合的方式实现对计算能力的串联，比如通过goroutine+channel的组合实现类似Unix Pipe的能力。


+ 在Go中，大多数应用数组的场景都可以用切片替代
``` go
 for i, v := range a[:] {        
    if i == 0 {            
        a[1] = 12            
        a[2] = 13        
 }
```


```go
// chapter3/sources/control_structure_idiom_5.go 
func recvFromUnbufferedChannel() {    
    var c = make(chan int)    
    go func() {        
        time.Sleep(time.Second * 3)        
        c <- 1        
        c <- 2        
        c <- 3        
        close(c)    
    }()    
    for v := range c {        
        fmt.Println(v)   
    }
}
```



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


 ④ 支持超时机制的等待

```go
// chapter6/sources/go-concurrency-pattern-4.go 
func main() {    
	done := spawnGroup(5, worker, 30)    
	println("spawn a group of workers")        
	timer := time.NewTimer(time.Second * 5)    
	defer timer.Stop()    
	select {    
		case <-timer.C:        
		println("wait group workers exit timeout!")    
		case <-done:        println("group workers done")    
	}
}
```


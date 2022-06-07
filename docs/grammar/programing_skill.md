+ 这是一篇实践者对 Go 语言的微吐槽:<https://mp.weixin.qq.com/s/UP_Rf-x1HekU-nl7BneVyg>
    - 零初始化：不注意可能导致0值参数导致不合理的配置
    - 过度 linting：go1.8已经不会强制错误
    - 返回错误：发生错误时，它们无需为非错误返回值提供一个值。？？
    - nil 切片和 JSON：
    - Go 模块和 Gitlab：1.第一个问题是 Gitlab 允许用户拥有递归项目组；2.由于我们的 Gitlab 实例是私有的，并且 Go 尝试通过 https 下载 git 存储库，因此当我们尝试下载未经任何身份验证的 Go 模块时，会收到 401 错误。使用我们的 Gitlab 密码进行身份验证是不切实际的选择，尤其是在涉及 CI/CD 的情况下。我们找到的解决方案是在使用这个.gitconfig 发出 https 请求时，强制 git 使用 ssh。
    - 日期格式 API：须使用“2006-01-02”格式的字符串（我也觉得特别恶心。。）


   + Go语言真的在设计上是一个糟糕的语言吗？：<https://www.zhihu.com/question/310386573/answer/852786886>
    - Go差一步就可以有一个相当不错的范型系统的，可惜最后一步走错了路，结果做出了interface这个令人迷惑而且不伦不类的东西。

 
---

+ Golang中常用的代码优化点:<https://mp.weixin.qq.com/s/QONfbKioFf6VqJE2OwP7Kw>

+ 在初始化slice的时候尽量补全cap
growslice的作用就是扩充slice的容量大小
growsslice的操作是一个比较复杂的操作，它的表现和复杂度会高于最基本的初始化make方法。对追求性能的程序来说，应该能避免尽量避免。
具体对growsslice函数具体实现同学有兴趣的可以参考源码src的 runtime/slice.go
我们并不是每次都能在slice初始化的时候就能准确预估到最终的使用容量的。所以这里使用了一个“尽量”。明白是否设置slice容量的区别，我们在能预估容量的时候，请尽量使用方法2那种预估容量后的slice初始化方式

+ 初始化一个类的时候，如果类的构造参数较多，尽量使用Option写法 
```go
type Foo struct {
 name string
 id int
 age int

 db interface{}
}

// FooOption 代表可选参数
type FooOption func(foo *Foo)

// WithName 代表Name为可选参数
func WithName(name string) FooOption {
   return func(foo *Foo) {
      foo.name = name
   }
}

// WithAge 代表age为可选参数
func WithAge(age int) FooOption {
   return func(foo *Foo) {
      foo.age = age
   }
}

// WithDB 代表db为可选参数
func WithDB(db interface{}) FooOption {
   return func(foo *Foo) {
      foo.db = db
   }
}

// NewFoo 代表初始化
func NewFoo(id int, options ...FooOption) *Foo {
   foo := &Foo{
      name: "default",
      id:   id,
      age:  10,
      db:   nil,
   }
   for _, option := range options {
      option(foo)
   }
   return foo
}

// 具体使用NewFoo的函数
func Bar() {
   foo := NewFoo(1, WithAge(15), WithName("foo"))
   fmt.Println(foo)
}

```


+ 巧用大括号控制变量作用域
    - 在golang写的过程中，你一定有过为 := 和 = 烦恼的时刻。一个变量，到写的时候，我还要记得前面是否已经定义过了，如果没有定义过，使用 := ，如果已经定义过，使用 =。
```go
var name string
var folder string
var mod string
...
{
   prompt := &survey.Input{
      Message: "请输入目录名称：",
   }
   err := survey.AskOne(prompt, &name)
   if err != nil {
      return err
   }

   ...
}
{
   prompt := &survey.Input{
      Message: "请输入模块名称(go.mod中的module, 默认为文件夹名称)：",
   }
   err := survey.AskOne(prompt, &mod)
   if err != nil {
      return err
   }
   ...
}
{
   // 获取hade的版本
   client := github.NewClient(nil)
   prompt := &survey.Input{
      Message: "请输入版本名称(参考 https://github.com/gohade/hade/releases，默认为最新版本)：",
   }
   err := survey.AskOne(prompt, &version)
   if err != nil {
      return err
   }
   ...
}
```
+ 首先我将最终解析出来的最终变量在最开始做定义，然后使用三个大括号，分别将 name, mod, version 三个变量的解析逻辑封装在里面。而在每个大括号里面，err变量的作用域就完全局限在括号中了，每次都可以直接使用 := 来创建一个新的 err并处理它，不需要额外思考这个err 变量是否前面已经创建过了。
+ 如果你自己观察，大括号在代码语义上还有一个好处，就是归类和展示。归类的意思是，这个大括号里面的变量和逻辑是一个完整的部分，他们内部创建的变量不会泄漏到外部。这个等于等于告诉后续的阅读者，你在阅读的时候，如果对这个逻辑不感兴趣，不阅读里面的内容，而如果你感兴趣的话，可以进入里面进行阅读。基本上所有IDE都支持对大括号封装的内容进行压缩，我使用Goland，压缩后，我的命令行的主体逻辑就更清晰了。
+ 所以使用大括号，结合IDE，你的代码的可读性能得到很大的提升。



---

+ 一定记住，Go 中不要犯这 5 个错误:<https://mp.weixin.qq.com/s/ZJvGqPYbudzjd8KcAozA_A>
+ 1、循环内部
    - 1.1、循环迭代器变量中使用引用
    - 1.2、在循环中调用 WaitGroup.Wait
    - 1.3、循环内使用 defer
+ 2、channel 堵塞
    - 解决办法是将 ch 从无缓冲的通道改为有缓冲的通道，因此子goroutine 即使在父 goroutine 退出后也始终可以发送结果。
    - ????? 缓存区为1，发完1个没接收还不是会继续阻塞？？？？？
+ 3、不使用接口
    - 接口可以使代码更灵活。这是在代码中引入多态的一种方法。接口允许你定义一组行为而不是特定类型。不使用接口可能不会导致任何错误，但是会导致代码简单性，灵活性和扩展性降低。
    - 你应该知道的重要注意事项是，始终关注行为。在上面的示例中，虽然 io.ReadWriteCloser 也可以使用，但你只需要 Write 方法。接口越大，抽象性越弱。在 Go 中，通常提倡小接口。
+ 4、不注意结构体字段顺序
    - 这个问题不会导致程序错误，但是可能会占用更多内存。
    - 看起来这两个类型都占用的空间都是 21字节，但是结果却不是这样。我们使用 GOARCH=amd64 编译代码，发现 BadOrderedPerson 类型占用 32 个字节，而  OrderedPerson 类型只占用 24 个字节。为什么？原因是数据结构对齐[1]。在 64 位体系结构中，内存分配连续的 8 字节数据。
    - 当你使用大型常用类型时，可能会导致性能问题。但是不用担心，你不必手动处理所有结构。这工具可以轻松的解决此类问题：https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment。 ！！！！！
- 5、测试中不使用 race 探测器 !!!!
    - 数据争用会导致莫名的故障，通常是在代码已部署到线上很久之后才出现。因此，它们是并发系统中最常见且最难调试的错误类型。为了帮助区分此类错误，Go 1.1 引入了内置的数据争用检测器（race detector）。可以简单地添加 -race flag 来使用。
    ```shell
    $ go test -race pkg    # to test the package
        $ go run -race pkg.go  # to run the source file
        $ go build -race       # to build the package
        $ go install -race pkg # to install the package
    ```         
    - 启用数据争用检测器后，编译器将记录在代码中何时以及如何访问内存，而  runtime 监控对共享变量的非同步访问。
    找到数据竞争后，竞争检测器将打印一份报告，其中包含用于冲突访问的堆栈跟踪。      



1.1、循环迭代器变量中使用引用
出于效率考虑，经常使用单个变量来循环迭代器。由于在每次循环迭代中会有不同的值，有些时候这会导致未知的行为。例如：

in := []int{1, 2, 3}

var out []*int
for  _, v := range in {
 out = append(out, &v)
}

fmt.Println("Values:", *out[0], *out[1], *out[2])
fmt.Println("Addresses:", out[0], out[1], out[2])
输出结果：

Values: 3 3 3
Addresses: 0xc000014188 0xc000014188 0xc000014188
是不是很惊讶？在 out 这个 slice 中的元素都是 3。实际上很容易解释为什么会这样：在每次迭代中，我们都将 v append 到 out 切片中。因为 v 是单个变量（内存地址不变），每次迭代都采用新值。在输出的第二行证明了地址是相同的，并且它们都指向相同的值。

简单的解决方法是将循环迭代器变量复制到新变量中：

in := []int{1, 2, 3}

var out []*int
for  _, v := range in {
 v := v
 out = append(out, &v)
}

fmt.Println("Values:", *out[0], *out[1], *out[2])
fmt.Println("Addresses:", out[0], out[1], out[2])
新的输出：

Values: 1 2 3
Addresses: 0xc0000b6010 0xc0000b6018 0xc0000b6020
在 goroutine 中使用循环迭代变量会有相同的问题。

list := []int{1, 2, 3}

for _, v := range list {
 go func() {
  fmt.Printf("%d ", v)
 }()
}
输出将是：

3 3 3
可以使用上述完全相同的解决方案进行修复。请注意，如果不使用 goroutine 运行该函数，则代码将按预期运行。

这个错误犯错率是很高的，要特别注意！！

所以，使用匿名函数的时候go func的时候要时刻注意循环变量的Scope, 该传参传参，该重新定义重新定义。好在 Goland 最新版本已经会提示i存在Scope问题了。但是好像没几个人会注意IDE警告，所以，习惯很重要，不要写出IDE警告的代码也是一个不错的编程理念。

----

+ 哦，原来是这么回事：Golang 中的一些常识:<https://mp.weixin.qq.com/s/-l9R_QblXr1_JHGtjldoQw>

for _, i := range ss， ss 中的元素是 copy 到 变量i 的

现象

for range 的时候 slice 中的元素是copy给 变量i的，并且下次for循环，变量i会被直接覆盖。并不是把 n号元素的地址给了i，i 是第 n 号元素的 copy。


要更改生效也很简单，主要有两种方案，一种是使用切片指针 []*User，这样对于i的修改会被自动寻址到数字元素上。另一种是使用下标 主动寻址如 users[idx].Uid = 2 。


这个问题看似简单，如果将其使用go关键字并发将会发生巨大威力，造成血淋淋的事故。
其实用go的公司经常听到这样的事故：
* 某公司发运营push全部发给了同一个uid
* 某研发发运营消息发短信发给了同一个uid (如果通道商不限制，我相信用户哭了，哄不好的那种)
* 批量发优惠券，给同一个uid发了几百张
* ....
闭包问题一点都不新鲜，就是由于在go func里边使用for了循环的变量i了，然后因为函数体并没在go的时候立即执行需要申请资源挂载然后由M进行运行需要一些时间，所以一般for循环执行一段时间之后go func才会执行，这时候 内部函数取到的值就得听天命了。


+ []T 还是 []*T

现象

一般来说[]T 会比较高效一些，但是如果T比较大，在For循环时存在Copy开销，个人觉得[]*T也是可以的。


+ `[]interface{}`并不能接收[]T类型

现象

很多时候我们都以为interface可以传递任意类型，凡事总有例外，他就不能接收 []T 类型, 如果你需要进行赋值，那你要将T转成interface{}
理解

因为一个[]interface{}的空间是一定的，但是 []T 不是，因为占用空间不一致，编译器觉得有些代价，并没有进行转换.


+ Send on closed chan 会Panic，但是 Receive from closed chan 不会

现象

往已经关闭的channel 再send数据会触发runtime panic，但是receive从已经关闭的channel中消费不会触发.
理解

很多人有误区，认为chan关闭了就不能再操作了，但是send进chan的数据总归要消费完的，不然就丢了，你品。

+ Goroutine 之间不能 Recover painc

现象

goroutine没有父子关系（创建应该不算父子吧），不能在一个go中 recover 另一个 go 的 panic
理解

GPM模型在go的调度时没有上下级关系, 也没有跨goroutine的异常捕获机制。


----

+ 腾讯发布了Go语言代码安全指南 - 知乎:<https://zhuanlan.zhihu.com/p/400078436?utm_source=wechat_session&utm_medium=social&utm_oi=35352332992512&utm_campaign=shareopn>
+  https://github.com/Tencent/secguide !!!!!
    - 1.1.1切片长度校验
    - 1.1.2nil指针判断
    - 1.1.3整数安全
    - 1.1.4make分配长度验证 : `ifsize>64*1024*1024{returnnil,errors.New("value too large")`
    - 1.1.5禁止SetFinalizer和指针循环引用同时使用
    - 1.1.6禁止重复释放channel
    - 1.1.7确保每个协程都能退出
    - 1.1.8不使用unsafe包
    - 1.1.9不使用slice作为函数入参: slice是引用类型，在作为函数入参时采用的是地址传递，对slice的修改也会影响原始数据
    - 1.2.1 路径穿越检查
    - 1.2.2 文件访问权限
    - 1.3.1命令执行检查
    - 1.4.1网络通信采用TLS方式
    - 1.4.2TLS启用证书验证
    - 1.9.1禁止在闭包中直接调用循环变量: 在循环中启动协程，当协程中使用到了循环的索引值，由于多个协程同时使用同一个变量会产生数据竞争，造成执行结果异常。
    - 1.9.2禁止并发写map


---

+ Golang 编程思维和工程实战：<https://mp.weixin.qq.com/s/llmE9QpnrvA02AtvfHtqJQ> TODO !!!


### 一 Golang 编程思维
要理解 Golang 编程思维，首先要理解 Golang 这门语言的创始初衷，初衷就是为了解决好 Google 内部大规模高并发服务的问题，主要核心就是围绕高并发来开展；并且同时又不想引入面向对象那种很复杂的继承关系。首先，就是可以方便的解决好并发问题（包括高并发），那么就需要有并发思维，能够并发处理就通过并发来进行任务分配

这个就是涉及到了 context、 goroutine、channel（select）；
创建大量 goroutine， 但是需要能通过 context、 channel 建立 "父子"关系，保证子任务可以能够被回收、被主动控制（如 杀死）。
再者，面向对象编程思想，利用好 interface、 struct 来实现继承、多态的用法：

struct 匿名组合来实现继承；
terface 和 struct 来实现多态；
interface 定义接口，尽可能的保持里面的方法定义简单，然后多个 interface 进行组合。
然后，理解 Golang 语言本身的一些特性: - 强类型，语法上要注意处理；- GC，实际中要观察 GC 日志并分析；- 注意语法语义尽可能的简单、保持各种类型定义尽可能精简。

最后，从 Golang 社区的一些最佳实践来看，Golang 的各种组件需要尽可能的精简。

Golang 中用好的一些开源组件库，都是比较轻量级的，然后可以各自随意组合来达到最佳实践。
我们自己进行组件封装、模块封装的时候，也是保持这个原则，尽可能的精简，然后使用方进行组合。

### 二、Golang 高级编码技巧
+ 1 优雅的实现构造函数编程思想
+ 2 优雅的实现继承编程思想
+ 3 优雅的实现虚多态编程思想
+ 4 Golang 的 model service 模型【类 MVC 模型】
+ 5 Golang 单例模式
+ 6 Golang layout
<pre>
Golang 工程 Layout 规范，网上有较多探讨，每个人的理解也会不一致，但是有些基础的理解是可以保持统一的：

cmd

main 函数文件目录，这个目录下面，每个文件在编译之后都会生成一个可执行的文件。如果只有一个 app 文件，那就是 main.go。这里面的代码尽可能简单。
conf

配置文件，如 toml、yaml 等文件
config

配置文件的解析
docs

文档
pkg

底层各种实现，每一种实现封装一个文件夹
业界知名开源项目如 Kubernetes、Istio 都是这样的姿势
build

编译脚本
CI 脚本
上下线脚本
vendor

依赖库

$ tree  -d  -L 2
├── build
├── cmd
│   ├── apply
│   └── check
├── conf
├── config
├── docs
├── pkg
│   ├── apply
│   ├── check
│   ├── files
│   ├── k8s
│   └── options
└── vendor
</pre>

+ 7 cmd & command & flag


-----

+ Go 高性能编程技法:<https://mp.weixin.qq.com/s/Lv2XTD-SPnxT2vnPNeREbg> !!! 写得很好

### 常用数据结构
+ 1.1 优先使用 strconv 而不是 fmt
基本数据类型与字符串之间的转换，优先使用 strconv 而不是 fmt，因为前者性能更佳。
+ 1.2 少量的重复不比反射差
`go test -bench=. -benchmem main/reflect `
+ 1.3 慎用 binary.Read 和 binary.Write
binary.Read 和 binary.Write 使用反射并且很慢。如果有需要用到这两个函数的地方，我们应该手动实现这两个函数的相关功能，而不是直接去使用它们。

+ 2.避免重复的字符串到字节切片的转换
不要反复从固定字符串创建字节 slice，因为重复的切片初始化会带来性能损耗。相反，请执行一次转换并捕获结果。
+ 3.指定容器容量
+ 3.1 指定 map 容量提示
+ 3.2 指定切片容量
+ 4.字符串拼接方式的选择
+ 4.1 行内拼接字符串推荐使用运算符+
行内字符串的拼接，主要追求的是代码的简洁可读。fmt.Sprintf() 能够接收不同类型的入参，通过格式化输出完成字符串的拼接，使用非常方便。但因其底层实现使用了反射，性能上会有所损耗。

运算符 + 只能简单地完成字符串之间的拼接，非字符串类型的变量需要单独做类型转换。行内拼接字符串不会产生内存分配，也不涉及类型地动态转换，所以性能上优于fmt.Sprintf()。

从性能出发，兼顾易用可读，如果待拼接的变量不涉及类型转换且数量较少（<=5），行内拼接字符串推荐使用运算符 +，反之使用 fmt.Sprintf()。

+ 4.2 非行内拼接字符串推荐使用 strings.Builder
字符串拼接还有其他的方式，比如strings.Join()、strings.Builder、bytes.Buffer和byte[]，这几种不适合行内使用。当待拼接字符串数量较多时可考虑使用。
综合易用性和性能，一般推荐使用strings.Builder来拼接字符串。
+ 5.遍历 []struct{} 使用下标而不是 range
两种通过 index 遍历 []struct 性能没有差别，但是 range 遍历 []struct 中元素时，性能非常差。

range 只遍历 []struct 下标时，性能比 range 遍历  []struct 值好很多。从这里我们应该能够知道二者性能差别之大的原因。

Item 是一个结构体类型 ，Item 由两个字段构成，一个类型是 int，一个是类型是 [1024]byte，如果每次遍历 []Item，都会进行一次值拷贝，所以带来了性能损耗。

此外，因为 range 时获取的是值拷贝的副本，所以对副本的修改，是不会影响到原切片。

切片元素从结构体 Item 替换为指针 *Item 后，for 和 range 的性能几乎是一样的。而且使用指针还有另一个好处，可以直接修改指针对应的结构体的值。
+ 5.4 小结
range 在迭代过程中返回的是元素的拷贝，index 则不存在拷贝。

如果 range 迭代的元素较小，那么 index 和 range 的性能几乎一样，如基本类型的切片 []int。但如果迭代的元素较大，如一个包含很多属性的 struct 结构体，那么 index 的性能将显著地高于 range，有时候甚至会有上千倍的性能差异。对于这种场景，建议使用 index。如果使用 range，建议只迭代下标，通过下标访问元素，这种使用方式和 index 就没有区别了。如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，才能不影响性能。

### 内存管理
+ 1.使用空结构体节省内存
+ 1.1 不占内存空间
在 Go 中，我们可以使用 unsafe.Sizeof 计算出一个数据类型实例需要占用的字节数。
Go 中空结构体 struct{} 是不占用内存空间，不像 C/C++ 中空结构体仍占用 1 字节。
+ 1.2 用法
因为空结构体不占据内存空间，因此被广泛作为各种场景下的占位符使用。一是节省资源，二是空结构体本身就具备很强的语义，即这里不需要任何值，仅作为占位符，达到的代码即注释的效果。
+ 1.2.1 实现集合（Set）
Go 语言标准库没有提供 Set 的实现，通常使用 map 来代替。事实上，对于集合来说，只需要 map 的键，而不需要值。即使是将值设置为 bool 类型，也会多占据 1 个字节，那假设 map 中有一百万条数据，就会浪费 1MB 的空间。

因此呢，将 map 作为集合（Set）使用时，可以将值类型定义为空结构体，仅作为占位符使用即可。
+ 1.2.2 不发送数据的信道
有时候使用 channel 不需要发送任何的数据，只用来通知子协程（goroutine）执行任务，或只用来控制协程的并发。这种情况下，使用空结构体作为占位符就非常合适了。
+ 1.2.3 仅包含方法的结构体
在部分场景下，结构体只包含方法，不包含任何的字段。例如上面例子中的 Door，在这种情况下，Door 事实上可以用任何的数据结构替代。
+ 2.struct 布局要考虑内存对齐
+ 2.1 为什么需要内存对齐
简言之：合理的内存对齐可以提高内存读写的性能，并且便于实现变量操作的原子性。
+ 2.2 Go 内存对齐规则
编译器一般为了减少 CPU 访存指令周期，提高内存的访问效率，会对变量进行内存对齐。Go 作为一门追求高性能的后台编程语言，当然也不例外。
Go Language Specification 中 Size and alignment guarantees 描述了内存对齐的规则。
+ 2.3 合理的 struct 布局
因为内存对齐的存在，合理的 struct 布局可以减少内存占用，提高程序性能。
因此，在对内存特别敏感的结构体的设计上，我们可以通过调整字段的顺序，将字段宽度从小到大由上到下排列，来减少内存的占用。
+ 2.4 空结构与空数组对内存对齐的影响
空结构与空数组在 Go 中比较特殊。没有任何字段的空 struct{} 和没有任何元素的 array 占据的内存空间大小为 0。
+ 3.减少逃逸，将变量限制在栈上
变量逃逸一般发生在如下几种情况：

变量较大
变量大小不确定
变量类型不确定
返回指针
返回引用
闭包
知道变量逃逸的原因后，我们可以有意识的控制变量不发生逃逸，将其控制在栈上，减少堆变量的分配，降低 GC 成本，提高程序性能。
+ 3.1 小的拷贝好过引用
我们都知道 Go 里面的 Array 以 pass-by-value 方式传递后，再加上其长度不可扩展，考虑到性能我们一般很少使用它。实际上，凡事无绝对。有时使用数组进行拷贝传递，比使用切片要好。
+ 3.2 返回值 VS 返回指针
值传递会拷贝整个对象，而指针传递只会拷贝地址，指向的对象是同一个。返回指针可以减少值的拷贝，但是会导致内存分配逃逸到堆中，增加垃圾回收(GC)的负担。在对象频繁创建和删除的场景下，传递指针导致的 GC 开销可能会严重影响性能。

一般情况下，对于需要修改原对象值，或占用内存比较大的结构体，选择返回指针。对于只读的占用内存较小的结构体，直接返回值能够获得更好的性能。
+ 3.3 返回值使用确定的类型
如果变量类型不确定，那么将会逃逸到堆上。所以，函数返回值如果能确定的类型，就不要使用 interface{}。
+ 4.sync.Pool 复用对象
+ 4.1 简介
sync.Pool 是 sync 包下的一个组件，可以作为保存临时取还对象的一个“池子”。个人觉得它的名字有一定的误导性，因为 Pool 里装的对象可以被无通知地被回收，可能 sync.Cache 是一个更合适的名字。
+ sync.Pool 是可伸缩的，同时也是并发安全的，其容量仅受限于内存的大小。存放在池中的对象如果不活跃了会被自动清理。
+ 4.2 作用
对于很多需要重复分配、回收内存的地方，sync.Pool 是一个很好的选择。频繁地分配、回收内存会给 GC 带来一定的负担，严重的时候会引起 CPU 的毛刺，而 sync.Pool 可以将暂时不用的对象缓存起来，待下次需要的时候直接使用，不用再次经过内存分配，复用对象的内存，减轻 GC 的压力，提升系统的性能。
一句话总结：用来保存和复用临时对象，减少内存分配，降低 GC 压力。
+ 4.3 如何使用
sync.Pool 的使用方式非常简单，只需要实现 New 函数即可。对象池中没有对象时，将会调用 New 函数创建。
+ 4.4 性能差异
我们以 bytes.Buffer 字节缓冲器为例，利用 sync.Pool 复用 bytes.Buffer 对象，避免重复创建与回收内存，来看看对性能的提升效果。
从测试结果也可以看出，使用了 Pool 复用对象，每次操作不再有内存分配。
+ 4.5 在标准库中的应用
Go 标准库也大量使用了 sync.Pool，例如 fmt 和 encoding/json。以 fmt 包为例，我们看下其是如何使用 sync.Pool 的。
fmt.Printf() 的调用是非常频繁的，利用 sync.Pool 复用 pp 对象能够极大地提升性能，减少内存占用，同时降低 GC 压力。

### 并发编程
+ 1.关于锁
+ 1.1 无锁化
无锁化主要有两种实现，无锁数据结构和串行无锁。
+ 1.1.1 无锁数据结构
+ 1.1.2 串行无锁
串行无锁是一种思想，就是避免对共享资源的并发访问，改为每个并发操作访问自己独占的资源，达到串行访问资源的效果，来避免使用锁。不同的场景有不同的实现方式。比如网络 I/O 场景下将单 Reactor 多线程模型改为主从 Reactor 多线程模型，避免对同一个消息队列锁读取。
+ 1.2 减少锁竞争
+ 1.3 优先使用共享锁而非互斥锁
如果并发无法做到无锁化，优先使用共享锁而非互斥锁。

所谓互斥锁，指锁只能被一个 Goroutine 获得。共享锁指可以同时被多个 Goroutine 获得的锁。

Go 标准库 sync 提供了两种锁，互斥锁（sync.Mutex）和读写锁（sync.RWMutex），读写锁便是共享锁的一种具体实现。
+ 1.3.1 sync.Mutex
+ 1.3.2 sync.RWMutex
+ 1.3.3 性能对比
大部分业务场景是读多写少，所以使用读写锁可有效提高对共享数据的访问效率。最坏的情况，只有写请求，那么读写锁顶多退化成互斥锁。所以优先使用读写锁而非互斥锁，可以提高程序的并发性能。
可见读多写少的场景，使用读写锁并发性能会更优。可以预见的是如果写占比更低，那么读写锁带的并发效果会更优。
+ 2.限制协程数量
+ 2.1 协程数过多的问题
+ 2.1.1 程序崩溃
Go 程（goroutine）是由 Go 运行时管理的轻量级线程。通过它我们可以轻松实现并发编程。但是当我们无限开辟协程时，将会遇到致命的问题。
每个协程至少需要消耗 2KB 的空间，那么假设计算机的内存是 4GB，那么至多允许 4GB/2KB = 1M 个协程同时存在。那如果协程中还存在着其他需要分配内存的操作，那么允许并发执行的协程将会数量级地减少。
+ 2.1.2 协程的代价
Go 的开销主要是三个方面：创建（占用内存）、调度（增加调度器负担）和删除（增加 GC 压力）。

内存开销
空间上，一个 Go 程占用约 2K 的内存，在源码 src/runtime/runtime2.go里面，我们可以找到 Go 程的结构定义type g struct。

调度开销
时间上，协程调度也会有 CPU 开销。我们可以利用runntime.Gosched()让当前协程主动让出 CPU 去执行另外一个协程，下面看一下协程之间切换的耗时。
可见一次协程的切换，耗时大概在 100ns，相对于线程的微秒级耗时切换，性能表现非常优秀，但是仍有开销。
GC 开销 创建 Go 程到运行结束，占用的内存资源是需要由 GC 来回收，如果无休止地创建大量 Go 程后，势必会造成对 GC 的压力。

+ 2.2 限制协程数量
可以利用信道 channel 的缓冲区大小来实现。
```go
func main() {
 var wg sync.WaitGroup
 ch := make(chan struct{}, 3)
 for i := 0; i < 10; i++ {
  ch <- struct{}{}
  wg.Add(1)
  go func(i int) {
   defer wg.Done()
   log.Println(i)
   time.Sleep(time.Second)
   <-ch
  }(i)
 }
 wg.Wait()
}
```
+ 2.3 协程池化
上面的例子只是简单地限制了协程开辟的数量。在此基础之上，基于对象复用的思想，我们可以重复利用已开辟的协程，避免协程的重复创建销毁，达到池化的效果。

协程池化，我们可以自己写一个协程池，但不推荐这么做。因为已经有成熟的开源库可供使用，无需再重复造轮子。目前有很多第三方库实现了协程池，可以很方便地用来控制协程的并发数量，比较受欢迎的有：

Jeffail/tunny
panjf2000/ants
下面以 panjf2000/ants 为例，简单介绍其使用。

ants 是一个简单易用的高性能 Goroutine 池，实现了对大规模 Goroutine 的调度管理和复用，允许使用者在开发并发程序的时候限制 Goroutine 数量，复用协程，达到更高效执行任务的效果。

+ 2.4 小结
Golang 为并发而生。Goroutine 是由 Go 运行时管理的轻量级线程，通过它我们可以轻松实现并发编程。Go 虽然轻量，但天下没有免费的午餐，无休止地开辟大量 Go 程势必会带来性能影响，甚至程序崩溃。所以，我们应尽可能的控制协程数量，如果有需要，请复用它。

+ 3.使用 sync.Once 避免重复执行
+ 3.1 简介
sync.Once 是 Go 标准库提供的使函数只执行一次的实现，常应用于单例模式，例如初始化配置、保持数据库连接等。作用与 init 函数类似，但有区别。
init 函数是当所在的 package 首次被加载时执行，若迟迟未被使用，则既浪费了内存，又延长了程序加载时间。
sync.Once 可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下是线程安全的。
在多数情况下，sync.Once 被用于控制变量的初始化，这个变量的读写满足如下三个条件：

当且仅当第一次访问某个变量时，进行初始化（写）；
变量初始化过程中，所有读都被阻塞，直到初始化完成；
变量仅初始化一次，初始化完成后驻留在内存里。
+ 3.2 原理
sync.Once 用来保证函数只执行一次。要达到这个效果，需要做到两点：

计数器，统计函数执行次数；
线程安全，保障在多 Go 程的情况下，函数仍然只执行一次，比如锁。
+ 3.2.1 源码
下面看一下 sync.Once 结构，其有两个变量。使用 done 统计函数执行次数，使用锁 m 实现线程安全。
+ 3.2.2  done 为什么是第一个字段
+ 3.3 性能差异
使用 sync.Once 保证函数只会被执行一次和多次执行，二者的性能差异。
sync.Once 中保证了 Config 初始化函数仅执行了一次，避免了多次重复初始化，在并发环境下很有用。
+ 4.使用 sync.Cond 通知协程 ！！！！
+ 4.1 简介
sync.Cond 是基于互斥锁/读写锁实现的条件变量，用来协调想要访问共享资源的那些 Goroutine，当共享资源的状态发生变化的时候，sync.Cond 可以用来通知等待条件发生而阻塞的 Goroutine。

sync.Cond 基于互斥锁/读写锁，它和互斥锁的区别是什么呢？

互斥锁 sync.Mutex 通常用来保护共享的临界资源，条件变量 sync.Cond 用来协调想要访问共享资源的 Goroutine。当共享资源的状态发生变化时，sync.Cond 可以用来通知被阻塞的 Goroutine。
+ 4.2 使用场景
sync.Cond 经常用在多个 Goroutine 等待，一个 Goroutine 通知（事件发生）的场景。如果是一个通知，一个等待，使用互斥锁或 channel 就能搞定了。


我们想象一个非常简单的场景：

有一个协程在异步地接收数据，剩下的多个协程必须等待这个协程接收完数据，才能读取到正确的数据。在这种情况下，如果单纯使用 chan 或互斥锁，那么只能有一个协程可以等待，并读取到数据，没办法通知其他的协程也读取数据。

这个时候，就需要有个全局的变量来标志第一个协程数据是否接受完毕，剩下的协程，反复检查该变量的值，直到满足要求。或者创建多个 channel，每个协程阻塞在一个 channel 上，由接收数据的协程在数据接收完毕后，逐个通知。总之，需要额外的复杂度来完成这件事。

Go 语言在标准库 sync 中内置一个 sync.Cond 用来解决这类问题。

channel+select 的组合，是比较优雅的通知？？？

+ 4.3 原理
sync.Cond 内部维护了一个等待队列，队列中存放的是所有在等待这个 sync.Cond 的 Go 程，即保存了一个通知列表。sync.Cond 可以用来唤醒一个或所有因等待条件变量而阻塞的 Go 程，以此来实现多个 Go 程间的同步。
+ 4.4 使用示例
我们实现一个简单的例子，三个协程调用 Wait() 等待，另一个协程调用 Broadcast() 唤醒所有等待的协程。
```go
var done = false

func read(name string, c *sync.Cond) {
 c.L.Lock()
 for !done {
  c.Wait()
 }
 log.Println(name, "starts reading")
 c.L.Unlock()
}

func write(name string, c *sync.Cond) {
 log.Println(name, "starts writing")
 time.Sleep(time.Second)
 done = true
 log.Println(name, "wakes all")
 c.Broadcast()
}

func main() {
 cond := sync.NewCond(&sync.Mutex{})

 go read("reader1", cond)
 go read("reader2", cond)
 go read("reader3", cond)
 write("writer", cond)

 time.Sleep(time.Second * 3)
}
```
+ 4.5 注意事项
sync.Cond 不能被复制
唤醒顺序
从等待队列中按照顺序唤醒，先进入等待队列，先被唤醒。
调用 Wait() 前要加锁
调用 Wait() 函数前，需要先获得条件变量的成员锁，原因是需要互斥地变更条件变量的等待队列。在 Wait() 返回前，会重新上锁。

---



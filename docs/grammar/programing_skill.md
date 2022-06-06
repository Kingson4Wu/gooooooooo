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

+ 声明类型： type ByKey []mr.KeyValue
+ 交换变量：func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
+ 获取入参：os.Args ，第一个代表执行文件路径；pg*.txt 传入的是匹配的字符数组
+ 打印：Go语言的%d,%p,%v等占位符的使用：<https://www.jianshu.com/p/66aaf908045e>
    - `fmt.Printf("%v,%v,%v,%v\n", os.Args[0], os.Args[1], os.Args[2], os.Args[3])`
+ 判断字符是否字母：`unicode.IsLetter(r)`
+ 数组 append 会自动扩容 
+ 字符串分割：`strings.FieldsFunc(contents, ff)`
+ strconv.Itoa()函数: 将数字转换成对应的字符串类型的数字
+ 切片循环：`for _, filename := range os.Args[2:]`
+ 读文件全部内容：`ioutil.ReadAll(file)`
+ 创建文件：`ofile, _ := os.Create(oname)`
    - 按格式写文件：`fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)`
+ Plugin symbol as function return : `mapf := xmapf.(func(string, string) []mr.KeyValue)` 入参类型，返回类型
```go
    doctor, ok := doc.(GoodDoctor)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
```
+ 排序
```go

type KeyValue struct {
	Key   string
	Value string
}

type ByKey []mr.KeyValue

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

intermediate := []mr.KeyValue{}
sort.Sort(ByKey(intermediate))

```
+ golang中判断字符串是否为空的方法：
if len(str) == 0{
}
或使用下面的方法判断：
if str == "" {
}




### 进阶
+ Go进阶25:Go插件plugin教程:<https://mojotv.cn/go/golang-plugin-tutorial>
+ golang中的init函数以及main函数:<https://www.cnblogs.com/TimLiuDream/p/9929934.html>
    - go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。
+ go语言net包rpc远程调用的使用
    - go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
    - go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码
    - net/rpc: <https://cloud.tencent.com/developer/section/1143675>
        + net/rpc 包允许 RPC 客户端程序通过网络或是其他 I/O 连接调用一个远端对象的公开方法（必须是大写字母开头、可外部调用的）
    - RPC / JSON-RPC in Golang:<https://www.jianshu.com/p/74ac2439afb2>
    - golang与java间的json-rpc跨语言调用:<https://www.cnblogs.com/geomantic/p/4751859.html>
+ go语言之行--golang核武器goroutine调度原理、channel详解 :<https://www.cnblogs.com/wdliu/p/9272220.html>   
    - https://www.runoob.com/w3cnote/go-channel-intro.html
    - Channel可以作为一个先入先出(FIFO)的队列，接收的数据和发送的数据的顺序是一致的。
+ 深入理解 Go 语言 defer:<https://zhuanlan.zhihu.com/p/63354092>
+ Go的异常处理 defer, panic, recover:<https://www.cnblogs.com/ghj1976/archive/2013/02/11/2910114.html>
+ golang error与panic处理:<https://www.jianshu.com/p/aee7e1bc4841>
    - error：可预见的错误, panic：不可预见的异常
    - error （方法返回值有）
        + `if err != nil {}`
    -  panic (使用recover处理)
        + `if err := recover(); err != nil {}`
+ Golang: 深入理解panic and recover:<https://ieevee.com/tech/2017/11/23/go-panic.html>        
    - 遇到处理不了的错误，找panic
    - panic有操守，退出前会执行本goroutine的defer，方式是原路返回(reverse order)
    - panic不多管，不是本goroutine的defer，不执行 
    ```go
    defer func() {
            if r := recover(); r != nil {
                if _, ok := r.(runtime.Error); ok {
                    panic(r)
                }
                err = r.(error)
            }
        }()
    ```      
+ error、panic 和 recover、defer
    - 使用 defer 语句进行延迟调用，用来关闭或释放资源。
    - 使用 panic 和 recover 来抛出错误和恢复。
    - 使用 panic 一般有两种情况：
        1. 程序遇到无法执行的错误时，主动调用 panic 结束运行；
        2. 在调试程序时，主动调用 panic 结束运行，根据抛出的错误信息来定位问题。
    - 为了程序的健壮性，可以使用 recover 捕获错误，恢复程序运行。
    - recover 恢复，recover 只能放到 defer 函数里面，不能放到子函数。

+ golang 中 sync.Mutex 和 sync.RWMutex:<https://www.jianshu.com/p/679041bdaa39>
+ 在WaitGroup
在WaitGroup 对象实现中，内部有一个计数器，最初从0开始，它有三个方法：
Add()：计数器加一
Done()：计数器减一
Wait()：等待计数器清零
    1. 计数器不能为负值 (报错)
    2. WaitGroup对象不是一个引用类型
+ GO 条件锁sync.Cond
  - Cond是一个条件锁，就是当满足某些条件下才起作用的锁，有的地方也叫定期唤醒锁，有的地方叫条件变量conditional variable。
  - 三种操作方法
    1. 等待通知: wait
    阻塞当前线程，直到收到该条件变量发来的通知
    2. 单发通知: signal
    让该条件变量向至少一个正在等待它的通知的线程发送通知，表示共享数据的状态已经改变。
    3. 广播通知: broadcast
    让条件变量给正在等待它的通知的所有线程都发送通知。
    - wait前必须先 lock
    
+ flag
在 Golang 程序中有很多种方法来处理命令行参数。简单的情况下可以不使用任何库，直接处理 os.Args；其实 Golang 的标准库提供了 flag 包来处理命令行参数

+ golang中的字符串底层实现是通过byte数组的，中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8
byte 等同于int8，常用来处理ascii字符
rune 等同于int32,常用来处理unicode或utf-8字符

+ 空结构体应用场景
https://mp.weixin.qq.com/s/Jy2wxqZYNMpQe7s1jtIR1g
空结构体主要有以下几个特点：
零内存占用
地址相同
无状态

空结构体的使用场景
1. 实现 Set 集合类型： 空结构体作为 value
2. 用于通道信号
3. 作为方法接收器

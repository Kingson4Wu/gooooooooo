+ golang.org/x/sys timeout 解决方式
https://www.cxymm.net/article/qq_29719097/102406092

解决方式：
 export GOPROXY=https://goproxy.io


 ---

+ 内存泄漏的定位与排查：Heap Profiling 原理解析:<https://mp.weixin.qq.com/s/acp0VqDI3SurOPp-xlHLmw>
+ golang：快来抓住内存泄漏的“真凶”！:<https://mp.weixin.qq.com/s/HosxXlz9e1jOmIY60RRkaQ>
 
<pre>
 Go runtime 内置了方便的 profiler，heap 是其中一种类型。我们可以通过如下方式开启一个 debug 端口：
import _ "net/http/pprof"

go func() {
   log.Print(http.ListenAndServe("0.0.0.0:9999", nil))
}()
然后在程序运行期间使用命令行拿到当前的 Heap Profiling 快照：

$ go tool pprof http://127.0.0.1:9999/debug/pprof/heap
或者也可以在应用程序代码的特定位置直接获取一次 Heap Profiling 快照：
import "runtime/pprof"

pprof.WriteHeapProfile(writer)
</pre>

+ 滴滴实战分享：通过 profiling 定位 golang 性能问题 - 内存篇:<http://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651439020&idx=1&sn=c2094f4dccb53385dc207958e7f42f9e&chksm=80bb615eb7cce8481eb7a8f09d4a13e2974b3785c241dd31245647cd7540dde414d64f2b3719&mpshare=1&scene=1&srcid=&sharer_sharetime=1585753889438&sharer_shareid=dcfe0eae58d1da3d4cc1d60a98c3905c#rd>
   - cpu 占用 99%
      ```go
      import _ "net/http/pprof"
      func main() {
      go func() {
      log.Println(http.ListenAndServe(
      "0.0.0.0:8005"
      , 
      nil
      ))
      }()
      // ..... 下面业务代码不用动
      }
      ```
    - 等到故障再次复现时，我们首先对 cpu 性能进行采样分析：
      + `brew install graphviz # 安装graphviz，只需要安装一次就行了 `
      + `go tool pprof -http=:1234 http://your-prd-addr:8005/debug/pprof/profile?seconds=30`



+ 腾讯 Go 性能优化实战:<https://mp.weixin.qq.com/s/Z9DoVGwdAtpbjealQLEMkw> !!!!
<pre>
* 性能查看工具 pprof,trace 及压测工具 wrk 或其他压测工具的使用要比较了解。
* 代码逻辑层面的走读非常重要，要尽量避免无效逻辑。
* 对于 golang 自身库存在缺陷的，可以寻找第三方库或自己改造。
* golang 版本尽量更新，这次的测试是在 golang1.12 下进行的。而 go1.13 甚至 go1.14 在很多地方进行了改进。比如 fmt.Sprintf，sync.Pool 等。替换成新版本应该能进一步提升性能。
* 本地 benchmark 结果不等于线上运行结果。尤其是在使用缓存来提高处理速度时，要考虑 GC 的影响。
* 传参数或返回值时，尽量按 golang 的设计哲学，少用指针，多用值对象，避免引起过多的变量逃逸，导致 GC 耗时暴涨。struct 的大小一般在 2K 以下的拷贝传值，比使用指针要快（可针对不同的机器压测，判断各自的阈值)。
* 值类型在满足需要的情况下，越小越好。能用 int8，就不要用 int64。
* 资源尽量复用,在 golang1.13 以上，可以考虑使用 sync.Pool 缓存会重复申请的内存或对象。或者自己使用并管理大块内存，用来存储小对象，避免 GC 影响（如本地缓存的场景)。
</pre>

+ 字节跳动踩坑记：Goroutine 泄漏:<https://mp.weixin.qq.com/s/5q5eIMDHz35ycTBTkB33JQ>

+ 快速定位线上性能问题：Profiling 在微服务应用下的落地实践:<https://mp.weixin.qq.com/s/KqNHNs75CimBMX9cF2zwZw>
+ Golang pprof 解析
<pre>
* cpu：CPU 分析，按照一定的频率采集所监听的应用程序的 CPU 使用情况。
* heap：内存分析，记录内存分配情况。用于监控当前和历史内存使用情况，辅助检查内存泄漏。
* threadcreate：反映系统线程的创建情况。
* goroutine：当前所有 goroutine 的堆栈跟踪。
* block：阻塞分析，记录 goroutine 阻塞等待同步的位置，可以用来分析和查找死锁等性能瓶颈。
* mutex：互斥锁分析，记录互斥锁的竞争情况。
</pre>
+ 进程 Profiling
<pre>
* 非常驻进程
可以通过如下方式引入 runtime/pprof 库，在进程退出后，就可以获得 Profiling 数据：
import"runtime/pprof"
funcmain() {
 f, err := os.Create(
"path/to/cpu.out"
)
if
 err != 
nil
 {
 log.Fatal(err)
 }
 pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
 ...

 * 常驻进程
可以引入 net/http/pprof 来通过特定的 http 接口获得 Profiling 数据，这个库会注册如下的路由：
http.HandleFunc(
"/debug/pprof/"
, Index)
http.HandleFunc(
"/debug/pprof/cmdline"
, Cmdline)
http.HandleFunc(
"/debug/pprof/profile"
, Profile)
http.HandleFunc(
"/debug/pprof/symbol"
, 
Symbol
)
http.HandleFunc(
"/debug/pprof/trace"
, Trace)
只需要在代码里启动一个 http server，就可以对外暴露出 pprof 信息，然后使用 go tool pprof 命令就可以通过这些路由获得数据：
go tool pprof 
http:
/
/localhost:6060/debug
/pprof/profile
# 30-second CPU profile
go tool pprof 
http:
/
/localhost:6060/debug
/pprof/heap
# heap profile
go tool pprof 
http:
/
/localhost:6060/debug
/pprof/block
# goroutine blocking profile

</pre>

+  函数 Profiling
<pre>
Golang 的 go test -bench 命令已经集成了 pprof 功能，只要针对特定函数编写 Benchmark 测试函数：
// in source file
funcfoo(){}
// in test file
funcBenchmark_foo(b *testing.B) {
for
 i := 
0
; i < b.N; i++ {
 foo()
 }
}
使用如下指令可以在不侵入原有代码的情况下获得 foo 函数 Profiling 数据：
go 
test
 -benchmem -cpuprofile=path/to/cpu.out -bench 
'^(Benchmark_foo)$'
 .

</pre>

---

+ Go 问题排查实战：一个死锁bug的始末:<https://mp.weixin.qq.com/s/KqNHNs75CimBMX9cF2zwZw>
+ perf
+ 火焰图
+ dlv
+ core file debug
+ 汇编、源码

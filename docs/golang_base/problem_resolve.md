+ golang.org/x/sys timeout 解决方式
https://www.cxymm.net/article/qq_29719097/102406092

解决方式：
 export GOPROXY=https://goproxy.io


 ---

 + 内存泄漏的定位与排查：Heap Profiling 原理解析:<https://mp.weixin.qq.com/s/acp0VqDI3SurOPp-xlHLmw>
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


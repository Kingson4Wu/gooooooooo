+ 通过例子学 Go pprof:<https://mp.weixin.qq.com/s/mHhwohY4HySi82FVa4Asag>
+ 查询性能瓶颈在哪里？查询内存泄漏在哪里？好在 pprof 是处理此类问题的利器，共有两套标准库，分别适用于不同的场景：
    - runtime/pprof[1]：采集工具型应用运行数据进行分析
    - net/http/pprof[2]：采集服务型应用运行时数据进行分析

+ 通过 go tool pprof 查看 /debug/pprof/profile：
`go tool pprof -http :8080 http://localhost:6060/debug/pprof/profile`


---

import _ "net/http/pprof"

// cpu
go tool pprof -http :8080 http://localhost:30116/debug/pprof/profile

go tool pprof -http :8080 http://localhost:30116/debug/fgprof
fgprof 不仅能检测到 onCPU（也就是 cpuIntensiveTask）部分，还能检测到 offCPU （也就是 slowNetworkRequest）部分,不过需要注意的是 fgprof 对性能的影响较大

//memory
go tool pprof -http :8080 http://localhost:30116/debug/pprof/heap

curl -s http://localhost:30116/debug/pprof/heap > base.out
curl -s http://localhost:30116/debug/pprof/heap > current.out
go tool pprof -base base.out current.out

Could not execute dot; may need to install graphviz.

brew install graphviz

<pre>
==> nginx
Docroot is: /usr/local/var/www

The default port has been set in /usr/local/etc/nginx/nginx.conf to 8080 so that
nginx can run without sudo.

nginx will load all files in /usr/local/etc/nginx/servers/.

To restart nginx after an upgrade:
  brew services restart nginx
Or, if you don't want/need a background service you can just run:
  /usr/local/opt/nginx/bin/nginx -g daemon off;
</pre>


http://www.graphviz.org/download/

(/usr/local/Cellar/wrk/4.1.0/bin)
    + `./wrk -d 60 -c 10000  -t 32 'http://127.0.0.1:3242/svdvsdv'`


   

---

https://zhuanlan.zhihu.com/p/51559344

PProf 关注的模块
CPU profile：报告程序的 CPU 使用情况，按照一定频率去采集应用程序在 CPU 和寄存器上面的数据
Memory Profile（Heap Profile）：报告程序的内存使用情况
Block Profiling：报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈
Goroutine Profiling：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎样的

两种引入方式
PProf 可以从以下两个包中引入：

import "net/http/pprof"
import "runtime/pprof"
其中 net/http/pprof 使用 runtime/pprof 包来进行封装，并在 http 端口上暴露出来。runtime/pprof 可以用来产生 dump 文件，再使用 Go Tool PProf 来分析这运行日志。

使用 net/http/pprof 可以做到直接看到当前 web 服务的状态，包括 CPU 占用情况和内存使用情况等。

这个路径下还有几个子页面：

/debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载
/debug/pprof/heap： Memory Profiling 的路径，访问这个链接会得到一个内存 Profiling 结果的文件
/debug/pprof/block：block Profiling 的路径
/debug/pprof/goroutines：运行的 goroutines 列表，以及调用关系

----
### 生成火焰图
+ https://github.com/uber/go-torch
+ 有两种方式：go-torch（golang version < 1.10）和golang原生的pprof（golang version < 1.10+的pprof集成了火焰图功能）

+ go tool pprof -http=:6061 http://localhost:6060/debug/pprof/block
+ http://localhost:6061/ui/flamegraph即可查看生成的火焰图。
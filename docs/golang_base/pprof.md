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
(这种采样是当前的快照)

go tool pprof -http :8080 http://localhost:30116/debug/fgprof
fgprof 不仅能检测到 onCPU（也就是 cpuIntensiveTask）部分，还能检测到 offCPU （也就是 slowNetworkRequest）部分,不过需要注意的是 fgprof 对性能的影响较大

//memory
go tool pprof -http :8080 http://localhost:30116/debug/pprof/heap

curl -s http://localhost:30116/debug/pprof/heap > base.out
curl -s http://localhost:30116/debug/pprof/heap > current.out
go tool pprof -base base.out current.out

Could not execute dot; may need to install graphviz.

brew install graphviz

go tool pprof -alloc_space -base ~/Downloads/profile.pb.gz ~/Downloads/profile\ \(2\).pb.gz

go tool pprof -alloc_space  ~/Downloads/profile\ \(2\).pb.gz

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


+ lua.post
```lua
wrk.method = "POST"
wrk.body = 'xxxx'
wrk.headers["protocol"] = "json"
wrk.headers["Content-Type"] = "application/x-thrift"
wrk.headers["xxx"] = "cxxx"
wrk.headers["Accept"]= "*/*"
response = function(status, headers, body)
--print(body)
end
```
+ ./wrk -d 60 -c 100  -t 32 -s post.lua http://127.0.0.1:30116/xxxxx
   

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


flat、flat% 表示函数在 CPU 上运行的时间以及百分比
sum% 表示当前函数累加使用 CPU 的比例
cum、cum%表示该函数以及子函数运行所占用的时间和比例，应该大于等于前两列的值

inuse_space — 已分配但尚未释放的内存空间
inuse_objects——已分配但尚未释放的对象数量
alloc_space — 分配的内存总量（已释放的也会统计）
alloc_objects — 分配的对象总数（无论是否释放）

+ golang 内存分析/内存泄漏:<https://cloud.tencent.com/developer/article/1703742>
<https://zhuanlan.zhihu.com/p/265080950>

----
### 生成火焰图
+ https://github.com/uber/go-torch
+ 有两种方式：go-torch（golang version < 1.10）和golang原生的pprof（golang version < 1.10+的pprof集成了火焰图功能）

+ go tool pprof -http=:6061 http://localhost:6060/debug/pprof/block
+ http://localhost:6061/ui/flamegraph即可查看生成的火焰图。


----

+ <https://mp.weixin.qq.com/s/3wTuWy-wzMmHCSrwfVARcA>

😀1、编写并执行单元测试，只有保证功能测试通过之后，才可以进行性能测试，否则可能会产生无效的性能测试结果。


😆2、执行性能测试并查看指标，这块一般通过工具法、比如常见vmstat、iostat、top、perf等工具查看性能指标。Golang相对是一门面向性能设计的编程语言，可以直接在单元测试中执行benchmark，并提供了pprof，生成性能CPU和内存热点数据。


😅3、分析指标是否满足要求，这块主要分析时延和资源占用，一般情况下，时延的重要性要高于资源占用，当然这个根据自己的实际情况分析性能指标数据。


🤣4、生成 CPU 和 内存 pprof 文件并关注热点事件。


👹5、分析火焰图中的热点函数并优化，性能优化本身就是一个反反复复，持续不断的过程，非常考察技术人员的耐心和综合实力。



----

执行 go tool pprof -http :8081 http://localhost:8033/debug/pprof/profile 命令后，生成的 pprof 文件会被加载到内存中，并启动一个 HTTP 服务器，监听本地的 8081 端口，以便您可以通过浏览器访问该服务器上的 Web 界面。

当您在浏览器中访问该 Web 界面时，您会看到一个交互式的 CPU 图表，显示了程序在执行时的 CPU 使用情况。这个图表是实时的吗？不完全是。实际上，该图表只显示了取样数据的聚合结果，而不是实时数据。这是因为，go tool pprof 命令默认使用 30 秒的时间窗口来聚合取样数据，并生成图表。

换句话说，您在浏览器中看到的图表是过去 30 秒内程序的 CPU 使用情况的聚合结果，而不是实时的 CPU 使用情况。如果您想看到更实时的结果，可以尝试减少时间窗口的大小，例如：

go tool pprof -http :8081 http://localhost:8033/debug/pprof/profile?seconds=5
这个命令会使用 5 秒的时间窗口来聚合取样数据，并生成图表，因此您会看到更实时的 CPU 使用情况。需要注意的是，更小的时间窗口可能会导致图表的噪声增加，因为它只显示了更短的时间段内的 CPU 使用情况。因此，您需要根据具体情况选择合适的时间窗口大小。

---

不完全正确。当您访问 go tool pprof 生成的 Web 界面时，会显示在一段时间窗口内（默认是 30 秒）的 CPU 使用情况的聚合结果，这个结果并不是实时更新的。也就是说，每次刷新页面时，您看到的数据都是相同的，直到时间窗口结束并重新开始。

如果您想要实时更新数据，可以使用 go tool pprof 提供的 -seconds 选项来指定时间窗口的大小，例如：

go tool pprof -http :8081 http://localhost:8033/debug/pprof/profile?seconds=5
这个命令会使用 5 秒的时间窗口来聚合数据，并在 Web 界面上实时显示结果，每 5 秒更新一次。您可以根据自己的需要调整 -seconds 选项的值，以获得更合适的实时更新效果。


----
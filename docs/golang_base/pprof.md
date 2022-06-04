+ 通过例子学 Go pprof:<https://mp.weixin.qq.com/s/mHhwohY4HySi82FVa4Asag>
+ 查询性能瓶颈在哪里？查询内存泄漏在哪里？好在 pprof 是处理此类问题的利器，共有两套标准库，分别适用于不同的场景：
    - runtime/pprof[1]：采集工具型应用运行数据进行分析
    - net/http/pprof[2]：采集服务型应用运行时数据进行分析

+ 通过 go tool pprof 查看 /debug/pprof/profile：
`go tool pprof -http :8080 http://localhost:6060/debug/pprof/profile`

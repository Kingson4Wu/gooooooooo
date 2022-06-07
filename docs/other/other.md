+ 深入Go代码覆盖率使用、场景与原理:<https://mp.weixin.qq.com/s/D9gHNAyGdXNLloetr7bR3g>

总结
代码覆盖率是判断代码书写的质量，识别无效代码的重要工具。在go生态中，go1.2提供了测试代码覆盖率的cover工具。
* 静态代码
对于静态的代码，要识别代码没有被使用，可以使用golangci-lint工具
golangci-lint run --disable-all --enable unused
* 对于线下的单元测试
可以使用go test -cover工具
* 测试环境下的代码
对于测试环境下有请求或长时间运行程序的单元覆盖率，可以借助cover工具使用文中巧妙的方式来测试。如果测试环境不充分完备，没有办法测试出来。线上统计函数是否被调用有两种方式倾入性和非倾入性。
* 对于线上倾入性的方式
主要是在关键位置注入一行函数进行统计。可以是函数级别的，甚至可以借助go test -cover 在每一个逻辑分支注入了函数。
如果只是考虑函数级别的，可以考虑直接shel脚本注入函数，这样做时间成本最低。也可以考虑借助AST抽象语法树的方式，成本略高。
如果是逻辑分支级别的，可以考虑借鉴开源库https://github.com/qiniu/goc 的手法来生成打桩后的go文件。它是对go test -cover 代码的封装，借助AST抽象语法树，在特定位置插入一行。
* 对于线上非倾入性的方式
通过pprof，信号中断处理的手法，抽样获取堆栈信息。一些处理时间很短的函数，将很难被检测到。这种方式对于检测到经常使用的函数有用，但是不适合推断没有使用过的函数。
参考资料
官方文档：https://go.dev/blog/cover
合并profile：https://github.com/rakyll/pprof-merge
七牛云goc：https://github.com/qiniu/goc
利用 go/ast 语法树做代码生成: https://segmentfault.com/a/1190000039215176
Golang AST语法树使用教程及示例：https://juejin.cn/post/6844903982683389960
hook思路：https://github.com/brahma-adshonor/gohook


+ Go能实现AOP吗？:<https://mp.weixin.qq.com/s/ACtjNgUKx-GQf52ujUes_w>
+ 从Golang调度器的作者视角探究其设计之道！:<https://mp.weixin.qq.com/s/mH23ola6B_n8N9PRc1kpPw>
+ 1.6万字长文：Go 协程的实现原理:<https://mp.weixin.qq.com/s/nyTF3IgPf1qkBWCJZQuTuA> TODO
+ Go 中 goroutine 是如何协作与抢占:<https://mp.weixin.qq.com/s/59eBnrnoigz9A_J5uxIqRg>
+ Linux、K8S、Go等是如何设计调度系统的？调度系统设计精要:<https://mp.weixin.qq.com/s/BSZ-mf6YAeHMlc_pc8CTPw>
+ 一套可供参考的Go微服务开发方法论:<https://mp.weixin.qq.com/s/RQxqi7jXag-LVWUCR4N9Gg>
+ 深入理解Go的GC回收机制:<https://mp.weixin.qq.com/s/cZVSx_L3tIhDG1dyO2klhA> TODO
+ 图解 Go GC 是怎样监听你的应用的？:<https://mp.weixin.qq.com/s/Yc3-RXfDo1aFqA1vHdqfxg>
+ Go语言如何实现stop the world？:<http://mp.weixin.qq.com/s?__biz=MzAwMDU1MTE1OQ==&mid=2653551973&idx=1&sn=63e7dc5c6ba48c34218d06280d647050&chksm=813a6cfdb64de5eb85af961f436fe92bf7c1f8f07048607a00a440fce932272dd8e7d836a406&mpshare=1&scene=1&srcid=&sharer_sharetime=1584423192957&sharer_shareid=dcfe0eae58d1da3d4cc1d60a98c3905c#rd>
+ 深入理解Go中的内存分配:<https://mp.weixin.qq.com/s/bH1j7fpvR4MhbEq1CWRJNw>
+ Go 为什么这么“快”？:<https://mp.weixin.qq.com/s/CCU5b-9RFwsjdZQ4hugIgA>
    - 本文主要从 Go 调度器架构层面上介绍了 G-P-M 模型，通过该模型怎样实现少量内核线程支撑大量 Goroutine 的并发运行。以及通过 NetPoller、sysmon 等帮助 Go 程序减少线程阻塞，充分利用已有的计算资源，从而最大限度提高 Go 程序的运行效率。


+ 硬核，图解bufio包系列之读取原理:<https://mp.weixin.qq.com/s/rjQ9_8TxfHXpZF4B1gj32w>
+ Go 通过 Map/Filter/ForEach 等流式 API 高效处理数据:<https://mp.weixin.qq.com/s/t3INtSfFSmv-nsJqLmdPew>
    - 至此 Stream 组件就全部实现完了，核心逻辑是利用 channel 当做管道，数据当做水流，不断的用协程接收/写入数据到 channel 中达到异步非阻塞的效果。
    - 回到开篇提到的问题，未动手前想要实现一个 stream 难度似乎非常大，很难想象在 go 中 300 多行的代码就能实现如此强大的组件。
    - 实现高效的基础来源三个语言特性：
        - channel   
        - 协程
        - 函数式编程
    - https://github.com/zeromicro/go-zero

+ 后台自动化测试与持续部署实践:<https://mp.weixin.qq.com/s/lqwGUCKZM0AvEw_xh-7BDA>
+ Golang 从零到一开发实现 RPC 框架:<https://mp.weixin.qq.com/s/cx3O0wrc7cVy7TygHiWnvw>
+ 手把手带你从0搭建一个Golang ORM框架:<https://mp.weixin.qq.com/s/AoFb1UNvlWGhsvwUIdMt7w>
+ 如何使用 WebAsemmbly 在浏览器中编译 Go 代码:<https://mp.weixin.qq.com/s/AqEaYIRPSKjzZ1AE4cDtHA>
+ Go开源项目推荐：HTTP 请求时间花在哪:<https://mp.weixin.qq.com/s/73IU0WxCsaNIJZ-1tda0xQ>

+ BPF 和 Go: Linux 中的现代内省形式:<https://mp.weixin.qq.com/s/XjthFHWG6GMbrh1zyGpteg>
+ 通过 eBPF 深入探究 Go GC:<https://mp.weixin.qq.com/s/Nx5yLo80nCtTuSD_34HN4w>

+ Go：用 kqueue 实现一个简单的 TCP Server:<https://mp.weixin.qq.com/s/LbaWZCZzlHsxcpc8FHrKGQ>
+ 知乎社区核心业务 Golang 化实践:<https://mp.weixin.qq.com/s/10BBSbuk1mn3QC5AgI5sWA>

+ Go+ 下个里程碑：超越 cgo，无缝对接 C 语言:<https://mp.weixin.qq.com/s/CJT_MyZ4gQxB0elUBSj7wg>
+ https://github.com/goplus/c2go
    - 简单一句话，cgo 太鸡肋，与 C 语言的兼容上，Go 也就是做到了聊胜于无而已。
    - 最终，我没有选择优化 cgo，而是选择了：c2go。简单说，就是把 C 代码转换为 Go 代码，然后重新用 Go 编译器进行编译。
    - 在介绍 Go+ 编译器实现原理（请移步 bilibili 搜索 "Go+ 公开课 · 第1期｜Go+ v1.x 的设计与实现" 进行查看）的时候，我提到过 github.com/goplus/gox 这个项目，它是用来辅助生成 Go AST 的模块。在 C AST 转为 Go AST 中，我们也会借助它来大幅减少开发工作量。
    - 生成了 Go AST，剩下来的工作就和我们在 Go+ 一样了，通过 Go AST 调用 go/format 这个标准库来生成 Go 源文件，最后用 Go 编译器编译它。

+ Golang 最细节篇之 — Reader 和 ReaderAt 的区别:<https://mp.weixin.qq.com/s/BcSFBpHUbb_mLIaSYfoqrw> TODO

---

### golang面试
+ gc
+ 调度器、并发、channel
+ 内存分配
+ pprof使用: pprof分析，竞态分析，逃逸分析，这些基础的手段是必须要学会的

进阶要去学习怎么定位 Go 在线上系统的问题，成为一个 Go 的高级工程师。这部分需要大家了解一些 Go 的底层知识，学习基于 goroutine 和 channel 的各种并发编程模式，以及常用的工具链：比如 pprof 怎么用，怎么用 --base 去找内存泄露，出了性能问题怎么做优化等等。要达到的目标是：线上的 Go 系统出了问题，能够通过自己的知识储备快速定位。Go 的底层知识现在国内比 java 圈还卷，文章很泛滥，可以随意搜搜，择优阅读。

https://mp.weixin.qq.com/s/2wXNMd9fD3q5v9QNEUWa1A

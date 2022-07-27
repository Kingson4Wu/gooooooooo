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

+ 从源代码角度看 epoll 在 Go 中的使用（一）：<https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651438660&idx=2&sn=9eefcfdc79e1ed307515b58badb1d1cd&chksm=80bb60b6b7cce9a03d06a4aa307a203093b5ccf604cd7078a49cf3aef6c0a6be2cb3cc034084&mpshare=1&scene=1&srcid=&sharer_sharetime=1582693740247&sharer_shareid=dcfe0eae58d1da3d4cc1d60a98c3905c#rd>



---

### golang面试
+ gc
+ 调度器、并发、channel
+ 内存分配
+ pprof使用: pprof分析，竞态分析，逃逸分析，这些基础的手段是必须要学会的

进阶要去学习怎么定位 Go 在线上系统的问题，成为一个 Go 的高级工程师。这部分需要大家了解一些 Go 的底层知识，学习基于 goroutine 和 channel 的各种并发编程模式，以及常用的工具链：比如 pprof 怎么用，怎么用 --base 去找内存泄露，出了性能问题怎么做优化等等。要达到的目标是：线上的 Go 系统出了问题，能够通过自己的知识储备快速定位。Go 的底层知识现在国内比 java 圈还卷，文章很泛滥，可以随意搜搜，择优阅读。

https://mp.weixin.qq.com/s/2wXNMd9fD3q5v9QNEUWa1A


+ Go语言面试问得最多的面试题：<https://zhuanlan.zhihu.com/p/360306642>

https://github.com/chrislusf/seaweedfs !!!


<pre>
TinyRPC

TinyRPC 是基于Go语言标准库 net/rpc 扩展的远程过程调用框架，它具有以下特性：
* 基于TCP传输层协议
* 支持多种压缩格式：gzip、snappy、zlib；
* 基于二进制的 Protocol Buffer 序列化协议：具有协议编码小及高扩展性和跨平台性；
* 支持生成工具：TinyRPC提供的 protoc-gen-tinyrpc 插件可以帮助开发者快速定义自己的服务；
TinyRPC 的源代码仅有一千行左右，通过学习 TinyRPC ，开发者可以得到以下收获：
* 代码简洁规范
* 涵盖大多数 Go 语言基础用法和高级特性
* 单元测试编写技巧
* TCP流中处理数据包的技巧
* RPC框架的设计理念
TinyBalancer

TinyBalancer 是基于Go语言标准库 net/http/httputil 扩展的反向代理负载均衡器，它支持以下特性：
* 支持http以及https协议
* 支持七种负载均衡算法，分别是：round-robin、random、power of 2 random choice、consistent hash、consistent hash with bounded、ip-hash、least-load。
* 支持心跳检测，故障恢复
TinyBalancer 的源代码仅有一千行左右，通过学习 TinyBalancer ，开发者可以得到以下收获：
* 深入理解负载均衡算法
* 代码简洁规范
* 用Go语言设计反向代理的技巧
* 单元测试编写技巧
* 工厂设计模式在go语言中的应用
</pre>


---

https://www.zhihu.com/question/64178718

大部分能想到用到的东西，除了手机app，剩下基本都可以做。
1、命令行程序。不分windows、linux、macos，扔进去就能用（当然，需要交叉编译，具体不展开，下同），读写数据库、小爬虫、定时任务等等等等，想怎么玩怎么玩。个人目前主要是用来同步数据，定时备份和清理垃圾。
2、图形化工具。还是全平台通用，官方虽然没有GUI库，但是第三方有的是。选个合适的库，c++能搞的，golang差不多都可以搞（GUI库的选择请参考https://www.zhihu.com/question/268536384/answer/1215107185）。个人目前是给一些命令行工具配界面用，用golang是因为实在喜欢协程的写法，再就是工作电脑和家庭电脑操作系统不一样，懒得分开写。
3、私人服务器。这回不光不挑平台了，还能直接扔在路由器或者旧手机上了（旧安卓手机废物利用请参考https://zhuanlan.zhihu.com/p/92664273）。golang天生支持arm，私人服务，无非就是个网盘、博客之类的，完全没必要买云空间（家庭网络穿透定位部分不展开）。
4、各种正牌服务。这个不多说，吃饭用的。反正用不了多少资源，5000qps的服务随便跑跑，完全没必要装什么jvm，python、php之类的运行环境，想跑就跑，想搬就搬。目前已经用caddy代替了nginx（证书自动展期方便），gin代替了tomcat（服务器内存太贵了）。剩下普通的API服务，自己写就完事了。目前最长的高吞吐量服务，已经跑了700多天还活蹦乱跳，上次停机还是停电的时候。
golang最大的特点，就是什么都能沾点边，而且学习起来曲线比较平滑（rust，说的就是你）。一个擅长其他编程语言的程序员，兼修golang，基本就是一两个星期的事儿。但是不建议没有编程经验的人直接上，容易被带偏。
当然，这东西坑也很多，几乎是唯一一个避坑指南比入门教程还长的编程语言（可以参考https://zhuanlan.zhihu.com/p/27518650，只是一部分，实际上更多）。

---
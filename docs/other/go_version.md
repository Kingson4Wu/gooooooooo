### Go 1.21
+ 2023 年 8月
+ Go 1.21新特性前瞻：https://mp.weixin.qq.com/s/PwSJQM7WmJeLnDDbCr78hg
+ 支持WASI （不就是Java的虚拟机， Java字节码，多个平台运行？？）

+ Go从1.11版本[13]就开始支持将Go源码编译为wasm二进制文件，并在支持wasm的浏览器环境中运行。

不过WebAssembly绝不仅仅被设计为仅限于在Web浏览器中运行，核心的WebAssembly语言是独立于其周围环境的，WebAssembly完全可以通过API与外部世界互动。在Web上，它自然使用浏览器提供的现有Web API。然而，在浏览器之外，之前还没有一套标准的API可以让WebAssembly程序使用。这使得创建真正可移植的非Web WebAssembly程序变得困难。WebAssembly System Interface(WASI)[14]是一个填补这一空白的倡议，它有一套干净的API，可以由多个引擎在多个平台上实现，并且不依赖于浏览器的功能（尽管它们仍然可以在浏览器中运行）。

### Go 1.20
+ 2023 年 2月
+ https://segmentfault.com/a/1190000043400204
+ 大部分更改都在工具链、运行时和库的实现中。
+ Go 1.17 添加了从切片到数组指针的转换。Go 1.20 扩展了它以允许从切片到数组的转换：给定一个切片x,[4]byte(x)现在可以写成*(*[4]byte)(x).


#### Arena
+ https://zhuanlan.zhihu.com/p/583572024
+ Arena 指的是一种从一个连续的内存区域分配一组内存对象的方式。优点比一般的内存分配更有效率，也可以一次性释放。当然了，它的重点是要手动管理内存。
+ Go1.20 将会支持 arena 特性，通过 GOEXPERIMENT=arena 来打开
+ Go 1.20新特性Arena手动内存管理:<https://juejin.cn/post/7195889966756380730>


### 2022
+ https://mp.weixin.qq.com/s/TVMHl8CPutpvjp0ry3OzjQ

### Go 1.19
+ 2022 年 8月
+ 相对于Go 1.18，Go 1.19的确是一个“小版本”。但Go 1.19对memory model的更新、SetMemoryLimit的加入、go doc comment的修订以及对go runtime的持续打磨依然可以让gopher们产生一丝丝“小兴奋”，尤其是SetMemoryLimit的加入，是否能改善Go应用因GC不及时被kill的情况呢，让我们拭目以待。
+ https://zhuanlan.zhihu.com/p/527810013

### Go 1.18
+ 2022 年 3月
+ 泛型 Generics: 引入了对使用参数化类型的泛型代码的新支持, 达到了算法可复用的目的
+ 模糊测试Fuzzing: 提供了一种自动化测试的选择, Go 是第一个将模糊测试完全集成到其标准工具链中的主要语言
+ Workspaces: 解决go mod遗留下来的本地多模块开发依赖问题
+ 其次还包括CPU性能提升20%, 但是由于支持了泛型，对比1.17版本Go1.18 编译时间可能会慢 15-18%。

### Go 1.17
+ 2021 年 8月
+ 改进了编译器，具体来说是采用了一种新的函数参数和结果传递方式
+ 支持从 slice （切片）到数组指针的转换
+ unsafe包增加了两个函数：Add与Slice
+ go module同样有几处显著变化，其中最最重要的一个变化就是pruned module graph

### Go 1.16
+ 2021 年 2月
+ 核心库增加新成员 embed - 支持静态资源嵌入
+ 添加了对 macOS ARM64 的支持（也称为 Apple 芯片）
+ 默认开启Go modules
+ 开始禁止 import 导入的模块以 . 开头，模块路径中也不允许出现任何非 ASCII 字符

### Go 1.15
+ 2020 年 8月
+ 完全重写的链接器
+ 编译器改进，包括略微小了些的二进制文件
+ 内嵌 tzdata（时区数据）
+ 增加 testing.TB.TempDir


### Go 1.14
+ 2020 年 2 月
+ Go Module 已经可以用于生产环境，鼓励所有用户迁移到 Module。该版本支持嵌入具有重叠方法集的接口。
+ 性能方面做了较大的改进，包括：进一步提升 defer 性能、页分配器更高效，同时 timer 也更高效。
+ Goroutine 支持异步抢占。

### Go 1.13 
+ 2019 年 9 月
+ 改进了 sync 包中的 Pool，在 gc 运行时不会清除 pool。它引进了一个缓存来清理两次 gc 运行时都没有被引用的 pool 中的实例。
+ 重写了逃逸分析，减少了 Go 程序中堆上的内存申请的空间

### Go 1.12 
+ 2019 年 2 月
+ 基于 analysis 包重写了 go vet 命令，为开发者写自己的检查器提供了更大的灵活性

### Go 1.11
+ 2018 年 8 月
+ 一个重要的新功能：Go modules。去年的调查显示，Go modules 是 Go 社区遭遇重大挑战后的产物
+ 实验性的 WebAssembly，为开发者提供了把 Go 程序编译成一个可兼容四大主流 Web 浏览器的二进制格式的能力

### Go 1.10 
+ 2018 年 2 月
+ test 包引进了一个新的智能 cache，运行会测试后会缓存测试结果。如果运行完一次后没有做任何修改，那么开发者就不需要重复运行测试，节省时间
+ 为了加快构建速度，go build 命令现在也维持了一份最近构建包的缓存


### Go 1.9  
+ 2017 年 8 月
+ 支持别名声明
+ sync 包新增了一个 Map 类型，是并发写安全的

### Go 1.8
+ 2017 年 2 月
+ 把 gc 的停顿时间减少到了 1 毫秒以下
+ 改进了 defer 函数

### Go 1.7
+ 2016 年 8 月
+ 发布了 context 包，为用户提供了处理超时和任务取消的方法
+ 对编译工具链也作了优化，编译速度更快，生成的二进制文件更小，有时甚至可以减小 20% 到 30%

### Go 1.6
+ 2016 年 2 月
+ 使用 HTTPS 时默认支持 HTTP/2

### Go 1.5
+ 2015 年 8 月
+ 对 gc 进行了重新设计。归功于并发的回收，在回收期间的等待时间大大减少。一个 Twitter 生产环境的服务器的例子，等待时间由 300ms 降到 30ms
+ 发布了运行时追踪，用命令 go tool trace 可以查看。测试过程或运行时生成的追踪信息可以用浏览器窗口展示

### Go 1.4 
+ 2014 年 12 月
+ 官方对 Android 的支持，golang.org/x/mobile ) 让我们可以只用 Go 代码就能写出简单的 Android 程序
+ 更高效的 gc，之前用 C 和汇编写的运行时代码被翻译成 Go 后，堆的大小降低了 10% 到 30%
+ 提供了 go generate 命令通过扫描用 //go:generate 指示的代码来简化代码生成过程

### Go 1.3
+ 2014 年 6 月
+ 对栈管理做了重要的改进。栈可以申请连续的内存片段，提高了分配的效率，使下一个版本的栈空间降到 2KB。
+ 在 sync 包中发布了 Pool。 这个元素允许我们复用结构体，减少了申请的内存的次数，同时也是很多 Go 生态获得改进的根源，如标准库或包里的 encoding/json 或 net/http，还有 Go 社区里的 zap。

### Go 1.2
+ 2013 年 12 月
+ test 命令支持测试代码覆盖范围并提供了一个新命令 go tool cover ，此命令能测试代码覆盖率

### Go 1.1
+ 2013 年 5 月
+ 专注于优化语言（编译器，gc，map，go 调度器）和提升它的性能

### Go 1.0
+ 2012 年 3 月



----


作者：CatchZeng
链接：https://www.jianshu.com/p/f27c4f561544
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

Go 各版本回顾:<https://studygolang.com/articles/28435>


--- 
+ Go 1.18新特性解读（万字长文）:<https://mp.weixin.qq.com/s/8CyoGLuepuCI4Hj1Ev0GcQ>

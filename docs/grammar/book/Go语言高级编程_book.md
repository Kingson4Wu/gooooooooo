## 语言基因族谱

+ 首先看基因图谱的左边一支。可以明确看出 Go 语言的并发特性是由贝尔实验室的 Hoare 于 1978 年发布的 CSP 理论演化而来。其后，CSP 并发模型在 Squeak/NewSqueak 和 Alef 等编程语言中逐步完善并走向实际应用，最终这些设计经验被消化并吸收到了 Go 语言中。业界比较熟悉的 Erlang 编程语言的并发编程模型也是 CSP 理论的另一种实现。
再看基因图谱的中间一支。中间一支主要包含了 Go 语言中面向对象和包特性的演化历程。Go 语言中包和接口以及面向对象等特性则继承自 Niklaus Wirth 所设计的 Pascal 语言以及其后所衍生的相关编程语言。其中包的概念、包的导入和声明等语法主要来自于 Modula-2 编程语言，面向对象特性所提供的方法的声明语法等则来自于 Oberon 编程语言。最终 Go 语言演化出了自己特有的支持鸭子面向对象模型的隐式接口等诸多特性。
最后是基因图谱的右边一支，这是对 C 语言的致敬。Go 语言是对 C 语言最彻底的一次扬弃，不仅仅是语法和 C 语言有着很多差异，最重要的是舍弃了 C 语言中灵活但是危险的指针运算。而且，Go 语言还重新设计了 C 语言中部分不太合理运算符的优先级，并在很多细微的地方都做了必要的打磨和改变。当然，C 语言中少即是多、简单直接的暴力编程哲学则被 Go 语言更彻底地发扬光大了（Go 语言居然只有 25 个关键字，sepc 语言规范还不到 50 页）。
Go 语言其它的一些特性零散地来自于其他一些编程语言；比如 iota 语法是从 APL 语言借鉴，词法作用域与嵌套函数等特性来自于 Scheme 语言（和其他很多编程语言）。Go 语言中也有很多自己发明创新的设计。比如 Go 语言的切片为轻量级动态数组提供了有效的随机存取的性能，这可能会让人联想到链表的底层的共享机制。还有 Go 语言新发明的 defer 语句（Ken 发明）也是神来之笔。

+ 1.1.1 来自贝尔实验室特有基因
作为 Go 语言标志性的并发编程特性则来自于贝尔实验室的 Tony Hoare 于 1978 年发表的鲜为外界所知的关于并发研究的基础文献：顺序通信进程（communicating sequential processes ，缩写为 CSP）。在最初的 CSP 论文中，程序只是一组没有中间共享状态的平行运行的处理过程，它们之间使用管道进行通信和控制同步。Tony Hoare 的 CSP 并发模型只是一个用于描述并发性基本概念的描述语言，它并不是一个可以编写可执行程序的通用编程语言。
CSP 并发模型最经典的实际应用是来自爱立信发明的 Erlang 编程语言。不过在 Erlang 将 CSP 理论作为并发编程模型的同时，同样来自贝尔实验室的 Rob Pike 以及其同事也在不断尝试将 CSP 并发模型引入当时的新发明的编程语言中。他们第一次尝试引入 CSP 并发特性的编程语言叫 Squeak（老鼠的叫声），是一个用于提供鼠标和键盘事件处理的编程语言，在这个语言中管道是静态创建的。然后是改进版的 Newsqueak 语言（新版老鼠的叫声），新提供了类似 C 语言语句和表达式的语法，还有类似 Pascal 语言的推导语法。Newsqueak 是一个带垃圾回收的纯函数式语言，它再次针对键盘、鼠标和窗口事件管理。但是在 Newsqueak 语言中管道已经是动态创建的，管道属于第一类值、可以保存到变量中。然后是 Alef 编程语言（Alef 也是 C 语言之父 Ritchie 比较喜爱的编程语言），Alef 语言试图将 Newsqueak 语言改造为系统编程语言，但是因为缺少垃圾回收机制而导致并发编程很痛苦（这也是继承 C 语言手工管理内存的代价）。在 Aelf 语言之后还有一个叫 Limbo 的编程语言（地狱的意思），这是一个运行在虚拟机中的脚本语言。Limbo 语言是 Go 语言最接近的祖先，它和 Go 语言有着最接近的语法。到设计 Go 语言时，Rob Pike 在 CSP 并发编程模型的实践道路上已经积累了几十年的经验，关于 Go 语言并发编程的特性完全是信手拈来，新编程语言的到来也是水到渠成了。
图1-2展示了 Go 语言库早期代码库日志可以看出最直接的演化历程（Git 用 git log --before={2008-03-03} --reverse 命令查看）。
纵观整个贝尔实验室的编程语言的发展进程，从 B 语言、C 语言、Newsqueak、Alef、Limbo 语言一路走来，Go 语言继承了来着贝尔实验室的半个世纪的软件设计基因，终于完成了 C 语言革新的使命。纵观这几年来的发展趋势，Go 语言已经成为云计算、云存储时代最重要的基础编程语言。
1.2.6 Go 语言 - 2007~2009
贝尔实验室后来经历了多次动荡，包括 Ken Thompson 在内的 Plan9 项目原班人马最终加入了 Google 公司。在发明 Limbo 等前辈语言诞生十多年之后，在 2007 年底，Go 语言三个最初的作者因为偶然的因素聚集到一起批斗 C++（传说是 C++ 语言的布道师在 Google 公司到处鼓吹的 C++11 各种牛逼特性彻底惹恼了他们），他们终于抽出了 20% 的自由时间创造了 Go 语言。最初的 Go 语言规范从 2008 年 3 月开始编写，最初的 Go 程序也是直接编译到 C 语言然后再二次编译为机器码。到了 2008 年 5 月，Google 公司的领导们终于发现了 Go 语言的巨大潜力，从而开始全力支持这个项目（Google 的创始人甚至还贡献了func关键字），让他们可以将全部工作时间投入到 Go 语言的设计和开发中。在 Go 语言规范初版完成之后，Go 语言的编译器终于可以直接生成机器码了。
入口函数 main 已经去掉了返回值，程序默认通过隐式调用 exit(0) 来返回。Go 语言朝着简单的方向逐步进化。

### 1.5 面向并发的内存模型
常见的并行编程有多种模型，主要有多线程、消息传递等。从理论上来看，多线程和基于消息的并发编程是等价的。由于多线程并发模型可以自然对应到多核的处理器，主流的操作系统因此也都提供了系统级的多线程支持，同时从概念上讲多线程似乎也更直观，因此多线程编程模型逐步被吸纳到主流的编程语言特性或语言扩展库中。而主流编程语言对基于消息的并发编程模型支持则相比较少，Erlang 语言是支持基于消息传递并发编程模型的代表者，它的并发体之间不共享内存。Go 语言是基于消息并发模型的集大成者，它将基于 CSP 模型的并发编程内置到了语言中，通过一个 go 关键字就可以轻易地启动一个 Goroutine，与 Erlang 不同的是 Go 语言的 Goroutine 之间是共享内存的。

+ 1.5.1 Goroutine和系统线程
Goroutine是 Go 语言特有的并发体，是一种轻量级的线程，由 go 关键字启动。在真实的 Go 语言的实现中，goroutine 和系统线程也不是等价的。尽管两者的区别实际上只是一个量的区别，但正是这个量变引发了 Go 语言并发编程质的飞跃。
首先，每个系统级线程都会有一个固定大小的栈（一般默认可能是 2MB），这个栈主要用来保存函数递归调用时参数和局部变量。固定了栈的大小导致了两个问题：一是对于很多只需要很小的栈空间的线程来说是一个巨大的浪费，二是对于少数需要巨大栈空间的线程来说又面临栈溢出的风险。针对这两个问题的解决方案是：要么降低固定的栈大小，提升空间的利用率；要么增大栈的大小以允许更深的函数递归调用，但这两者是没法同时兼得的。相反，一个 Goroutine 会以一个很小的栈启动（可能是 2KB 或 4KB），当遇到深度递归导致当前栈空间不足时，Goroutine 会根据需要动态地伸缩栈的大小（主流实现中栈的最大值可达到1GB）。因为启动的代价很小，所以我们可以轻易地启动成千上万个 Goroutine。
Go的运行时还包含了其自己的调度器，这个调度器使用了一些技术手段，可以在 n 个操作系统线程上多工调度 m 个 Goroutine。Go 调度器的工作和内核的调度是相似的，但是这个调度器只关注单独的 Go 程序中的 Goroutine。Goroutine 采用的是半抢占式的协作调度，只有在当前 Goroutine 发生阻塞时才会导致调度；同时发生在用户态，调度器会根据具体函数只保存必要的寄存器，切换的代价要比系统线程低得多。运行时有一个 runtime.GOMAXPROCS 变量，用于控制当前运行正常非阻塞 Goroutine 的系统线程数目。
在 Go 语言中启动一个 Goroutine 不仅和调用函数一样简单，而且 Goroutine 之间调度代价也很低，这些因素极大地促进了并发编程的流行和发展。
```go
func main() {
    go println("你好, 世界")
}
```
根据 Go 语言规范，main函数退出时程序结束，不会等待任何后台线程。因为 Goroutine 的执行和 main 函数的返回事件是并发的，谁都有可能先发生，所以什么时候打印，能否打印都是未知的。
用前面的原子操作并不能解决问题，因为我们无法确定两个原子操作之间的顺序。解决问题的办法就是通过同步原语来给两个事件明确排序：
```go
func main() {
    done := make(chan int)

    go func(){
        println("你好, 世界")
        done <- 1
    }()

    <-done
}
```
当 <-done 执行时，必然要求 done <- 1 也已经执行。根据同一个 Goroutine 依然满足顺序一致性规则，我们可以判断当 done <- 1 执行时，println("你好, 世界") 语句必然已经执行完成了。因此，现在的程序确保可以正常打印结果。
当然，通过 sync.Mutex 互斥量也是可以实现同步的：
```go
func main() {
    var mu sync.Mutex

    mu.Lock()
    go func(){
        println("你好, 世界")
        mu.Unlock()
    }()

    mu.Lock()
}
```
可以确定后台线程的 mu.Unlock() 必然在 println("你好, 世界") 完成后发生（同一个线程满足顺序一致性），main 函数的第二个 mu.Lock() 必然在后台线程的 mu.Unlock() 之后发生（sync.Mutex 保证），此时后台线程的打印工作已经顺利完成了。

+ 严谨的并发程序的正确性不应该是依赖于 CPU 的执行速度和休眠时间等不靠谱的因素的。严谨的并发也应该是可以静态推导出结果的：根据线程内顺序一致性，结合 Channel 或 sync 同步事件的可排序性来推导，最终完成各个线程各段代码的偏序关系排序。如果两个事件无法根据此规则来排序，那么它们就是并发的，也就是执行先后顺序不可靠的。
解决同步问题的思路是相同的：使用显式的同步。

### 1.6 常见的并发模式
+ Go 语言最吸引人的地方是它内建的并发支持。Go 语言并发体系的理论是 C.A.R Hoare 在 1978 年提出的 CSP（Communicating Sequential Process，通讯顺序进程）。CSP 有着精确的数学模型，并实际应用在了 Hoare 参与设计的 T9000 通用计算机上。从 NewSqueak、Alef、Limbo 到现在的 Go 语言，对于对 CSP 有着 20 多年实战经验的 Rob Pike 来说，他更关注的是将 CSP 应用在通用编程语言上产生的潜力。作为 Go 并发编程核心的 CSP 理论的核心概念只有一个：同步通信。关于同步通信的话题我们在前面一节已经讲过，本节我们将简单介绍下 Go 语言中常见的并发模式。
首先要明确一个概念：并发不是并行。并发更关注的是程序的设计层面，并发的程序完全是可以顺序执行的，只有在真正的多核 CPU 上才可能真正地同时运行。并行更关注的是程序的运行层面，并行一般是简单的大量重复，例如 GPU 中对图像处理都会有大量的并行运算。为更好的编写并发程序，从设计之初 Go 语言就注重如何在编程语言层级上设计一个简洁安全高效的抽象模型，让程序员专注于分解问题和组合方案，而且不用被线程管理和信号互斥这些繁琐的操作分散精力。
在并发编程中，对共享资源的正确访问需要精确的控制，在目前的绝大多数语言中，都是通过加锁等线程同步方案来解决这一困难问题，而 Go 语言却另辟蹊径，它将共享的值通过 Channel 传递(实际上多个独立执行的线程很少主动共享资源)。在任意给定的时刻，最好只有一个 Goroutine 能够拥有该资源。数据竞争从设计层面上就被杜绝了。为了提倡这种思考方式，Go 语言将其并发编程哲学化为一句口号：

+ Do not communicate by sharing memory; instead, share memory by communicating.
不要通过共享内存来通信，而应通过通信来共享内存。
+ 这是更高层次的并发编程哲学(通过管道来传值是 Go 语言推荐的做法)。虽然像引用计数这类简单的并发问题通过原子操作或互斥锁就能很好地实现，但是通过 Channel 来控制访问能够让你写出更简洁正确的程序。

+ 使用 sync.Mutex 互斥锁同步是比较低级的做法。我们现在改用无缓存的管道来实现同步：
```go
func main() {
    done := make(chan int)

    go func(){
        fmt.Println("你好, 世界")
        <-done
    }()

    done <- 1
}
```
上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无法保证 main 退出之前后台线程能正常打印了。更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步事件受管道缓存大小的影响：
```go
func main() {
    done := make(chan int, 1) // 带缓存的管道

    go func(){
        fmt.Println("你好, 世界")
        done <- 1
    }()

    <-done
}
```
对于带缓冲的 Channel，对于 Channel 的第 K 个接收完成操作发生在第 K+C 个发送操作完成之前，其中 C 是 Channel 的缓存大小。虽然管道是带缓存的，main 线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的。
基于带缓存的管道，我们可以很容易将打印线程扩展到 N 个。下面的例子是开启 10 个后台线程分别打印：
```go
func main() {
    done := make(chan int, 10) // 带 10 个缓存

    // 开 N 个后台打印线程
    for i := 0; i < cap(done); i++ {
        go func(){
            fmt.Println("你好, 世界")
            done <- 1
        }()
    }

    // 等待 N 个后台线程完成
    for i := 0; i < cap(done); i++ {
        <-done
    }
}
```
对于这种要等待 N 个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用 sync.WaitGroup 来等待一组事件：
```go
func main() {
    var wg sync.WaitGroup

    // 开 N 个后台打印线程
    for i := 0; i < 10; i++ {
        wg.Add(1)

        go func() {
            fmt.Println("你好, 世界")
            wg.Done()
        }()
    }

    // 等待 N 个后台线程完成
    wg.Wait()
}
```
其中 wg.Add(1) 用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行则不能保证被正常执行到）。当后台线程完成打印工作之后，调用 wg.Done() 表示完成一个事件。main 函数的 wg.Wait() 是等待全部的事件完成。

## 第 2 章 CGO 编程
过去的经验往往是走向未来的枷锁，因为在过气技术中投入的沉没成本会阻碍人们拥抱新技术。——chai2010
曾经一度因未能习得 C++ 令人眼花缭乱的新标准而痛苦不已；Go 语言 “少既是多” 大道至简的理念让我重拾信心，寻回了久违的编程乐趣。——Ending
C/C++ 经过几十年的发展，已经积累了庞大的软件资产，它们很多久经考验而且性能已经足够优化。Go 语言必须能够站在 C/C++ 这个巨人的肩膀之上，有了海量的 C/C++ 软件资产兜底之后，我们才可以放心愉快地用 Go 语言编程。C 语言作为一个通用语言，很多库会选择提供一个 C 兼容的 API，然后用其他不同的编程语言实现。Go 语言通过自带的一个叫 CGO 的工具来支持 C 语言函数调用，同时我们可以用 Go 语言导出 C 动态库接口给其它语言使用。本章主要讨论 CGO 编程中涉及的一些问题。
其实 CGO 不仅仅用于 Go 语言中调用 C 语言函数，还可以用于导出 Go 语言函数给 C 语言函数调用。
要使用 CGO 特性，需要安装 C/C++ 构建工具链，在 macOS 和 Linux 下是要安装 GCC，在 windows 下是需要安装 MinGW 工具。同时需要保证环境变量 CGO_ENABLED 被设置为 1，这表示 CGO 是被启用的状态。在本地构建时 CGO_ENABLED 默认是启用的，当交叉构建时 CGO 默认是禁止的。比如要交叉构建 ARM 环境运行的 Go 程序，需要手工设置好 C/C++ 交叉构建的工具链，同时开启 CGO_ENABLED 环境变量。然后通过 import "C" 语句启用 CGO 特性。
CGO 是架接 Go 语言和 C 语言的桥梁，它使二者在二进制接口层面实现了互通，但是我们要注意因两种语言的内存模型的差异而可能引起的问题。如果在 CGO 处理的跨语言函数调用时涉及到了指针的传递，则可能会出现 Go 语言和 C 语言共享某一段内存的场景。我们知道 C 语言的内存在分配之后就是稳定的，但是 Go 语言因为函数栈的动态伸缩可能导致栈中内存地址的移动 (这是 Go 和 C 内存模型的最大差异)。如果 C 语言持有的是移动之前的 Go 指针，那么以旧指针访问 Go 对象时会导致程序崩溃。
CGO 在使用 C/C++ 资源的时候一般有三种形式：直接使用源码；链接静态库；链接动态库。直接使用源码就是在 import "C" 之前的注释部分包含 C 代码，或者在当前包中包含 C/C++ 源文件。链接静态库和动态库的方式比较类似，都是通过在 LDFLAGS 选项指定要链接的库方式链接。
动态库出现的初衷是对于相同的库，多个进程可以共享同一个，以节省内存和磁盘资源。但是在磁盘和内存已经白菜价的今天，这两个作用已经显得微不足道了，那么除此之外动态库还有哪些存在的价值呢？从库开发角度来说，动态库可以隔离不同动态库之间的关系，减少链接时出现符号冲突的风险。而且对于 windows 等平台，动态库是跨越 VC 和 GCC 不同编译器平台的唯一的可行方式。
对于 CGO 来说，使用动态库和静态库是一样的，因为动态库也必须要有一个小的静态导出库用于链接动态库（Linux 下可以直接链接 so 文件，但是在 Windows 下必须为 dll 创建一个 .a 文件用于链接）。我们还是以前面的 number 库为例来说明如何以动态库方式使用。
2.11 补充说明
CGO 是 C 语言和 Go 语言混合编程的技术，因此要想熟练地使用 CGO 需要了解这两门语言。C 语言推荐两本书：第一本是 C 语言之父编写的《C 程序设计语言》；第二本是讲述 C 语言模块化编程的《C 语言接口与实现: 创建可重用软件的技术》。Go 语言推荐官方出版的《The Go Programming Language》和 Go 语言自带的全部文档和全部代码。
为何要花费巨大的精力学习 CGO 是一个问题。任何技术和语言都有它自身的优点和不足，Go 语言不是银弹，它无法解决全部问题。而通过 CGO 可以继承 C/C++ 将近半个世纪的软件遗产，通过 CGO 可以用 Go 给其它系统写 C 接口的共享库，通过 CGO 技术可以让 Go 语言编写的代码可以很好地融入现有的软件生态——而现在的软件正式建立在 C/C++ 语言之上的。因此说 CGO 是一个保底的后备技术，它是 Go 的一个重量级的替补技术，值得任何一个严肃的 Go 语言开发人员学习。

## 第 3 章 Go 汇编语言
Go 语言中很多设计思想和工具都是传承自 Plan9 操作系统，Go 汇编语言也是基于 Plan9 汇编演化而来。根据 Rob Pike 的介绍，大神 Ken Thompson 在 1986 年为 Plan9 系统编写的 C 语言编译器输出的汇编伪代码就是 Plan9 汇编的前身。所谓的 Plan9 汇编语言只是便于以手工方式书写该 C 语言编译器输出的汇编伪代码而已。
无论高级语言如何发展，作为最接近 CPU 的汇编语言的地位依然是无法彻底被替代的。只有通过汇编语言才能彻底挖掘 CPU 芯片的全部功能，因此操作系统的引导过程必须要依赖汇编语言的帮助。只有通过汇编语言才能彻底榨干 CPU 芯片的性能，因此很多底层的加密解密等对性能敏感的算法会考虑通过汇编语言进行性能优化。
对于每一个严肃的 Gopher，Go 汇编语言都是一个不可忽视的技术。因为哪怕只懂一点点汇编，也便于更好地理解计算机原理，也更容易理解 Go 语言中动态栈、接口等高级特性的实现原理。而且掌握了 Go 汇编语言之后，你将重新站在编程语言鄙视链的顶端，不用担心再被任何其它所谓的高级编程语言用户鄙视。
本章我们将以 AMD64 为主要开发环境，简单地探讨 Go 汇编语言的基础用法。
为了简单，我们先用 Go 语言定义并赋值一个整数变量，然后查看生成的汇编代码。
首先创建一个 pkg.go 文件，内容如下：

package pkg

var Id = 9527
代码中只定义了一个 int 类型的包级变量，并进行了初始化。然后用以下命令查看的 Go 语言程序对应的伪汇编代码：
```shell
$ go tool compile -S pkg.go
"".Id SNOPTRDATA size=8
  0x0000 37 25 00 00 00 00 00 00                          '.......
```
+ 其中 go tool compile 命令用于调用 Go 语言提供的底层命令工具，其中 -S 参数表示输出汇编格式。输出的汇编比较简单，其中 "".Id 对应 Id 变量符号，变量的内存大小为 8 个字节。变量的初始化内容为 37 25 00 00 00 00 00 00，对应十六进制格式的 0x2537，对应十进制为 9527。SNOPTRDATA 是相关的标志，其中 NOPTR 表示数据中不包含指针数据。
以上的内容只是目标文件对应的汇编，和 Go 汇编语言虽然相似当并不完全等价。Go 语言官网自带了一个 Go 汇编语言的入门教程，地址在：https://golang.org/doc/asm 。
汇编语言的真正威力来自两个维度：一是突破框架限制，实现看似不可能的任务；二是突破指令限制，通过高级指令挖掘极致的性能。对于第一个问题，我们将演示如何通过 Go 汇编语言直接访问系统调用，和直接调用 C 语言函数。对于第二个问题，我们将演示 X64 指令中 AVX 等高级指令的简单用法
在操作系统中，每个进程都会有一个唯一的进程编号，每个线程也有自己唯一的线程编号。同样在 Go 语言中，每个 Goroutine 也有自己唯一的 Go 程编号，这个编号在 panic 等场景下经常遇到。虽然 Goroutine 有内在的编号，但是 Go 语言却刻意没有提供获取该编号的接口。本节我们尝试通过 Go 汇编语言获取 Goroutine ID。
目前 Go 语言支持 GDB、LLDB 和 Delve 几种调试器。其中 GDB 是最早支持的调试工具，LLDB 是 macOS 系统推荐的标准调试工具。但是 GDB 和 LLDB 对 Go 语言的专有特性都缺乏很大支持，而只有 Delve 是专门为 Go 语言设计开发的调试工具。而且 Delve 本身也是采用 Go 语言开发，对 Windows 平台也提供了一样的支持。本节我们基于 Delve 简单解释如何调试 Go 汇编程序。
3.10 补充说明
如果是纯粹学习汇编语言，则可以从《深入理解程序设计：使用 Linux 汇编语言》开始，该书讲述了如何以 C 语言的思维实现汇编程序。如果是学习 X86 汇编，则可以从《汇编语言：基于 x86 处理器》开始，然后再结合《现代 x86 汇编语言程序设计》学习 AVX 等高级汇编指令的使用。
Go 汇编语言的官方文档非常匮乏。其中 “A Quick Guide to Go's Assembler” 是唯一的一篇系统讲述 Go 汇编语言的官方文章，该文章中又引入了另外两篇 Plan9 的文档：A Manual for the Plan 9 assembler 和 Plan 9 C Compilers。Plan9 的两篇文档分别讲述了汇编语言以及和汇编有关联的 C 语言编译器的细节。看过这几篇文档之后会对 Go 汇编语言有了一些模糊的概念，剩下的就是在实战中通过代码学习了。
Go 语言的编译器和汇编器都带了一个 -S 参数，可以查看生成的最终目标代码。通过对比目标代码和原始的 Go 语言或 Go 汇编语言代码的差异可以加深对底层实现的理解。同时 Go 语言连接器的实现代码也包含了很多相关的信息。Go 汇编语言是依托 Go 语言的语言，因此理解 Go 语言的工作原理是也是必要的。比较重要的部分是 Go 语言 runtime 和 reflect 包的实现原理。如果读者了解 CGO 技术，那么对 Go 汇编语言的学习也是一个巨大的帮助。最后是要了解 syscall 包是如何实现系统调用的。
得益于 Go 语言的设计，Go 汇编语言的优势也非常明显：跨操作系统、不同 CPU 之间的用法也非常相似、支持 C 语言预处理器、支持模块。同时 Go 汇编语言也存在很多不足：它不是一个独立的语言，底层需要依赖 Go 语言甚至操作系统；很多高级特性很难通过手工汇编完成。虽然 Go 语言官方尽量保持 Go 汇编语言简单，但是汇编语言是一个比较大的话题，大到足以写一本 Go 汇编语言的教程。本章的目的是让大家对 Go 汇编语言简单入门，在看到底层汇编代码的时候不会一头雾水，在某些遇到性能受限制的场合能够通过 Go 汇编突破限制。

## 第 4 章 RPC 和 Protobuf
Protobuf 的 protoc 编译器是通过插件机制实现对不同语言的支持。比如 protoc 命令出现 --xxx_out 格式的参数，那么 protoc 将首先查询是否有内置的 xxx 插件，如果没有内置的 xxx 插件那么将继续查询当前系统中是否存在 protoc-gen-xxx 命名的可执行程序，最终通过查询到的插件生成代码。对于 Go 语言的 protoc-gen-go 插件来说，里面又实现了一层静态插件系统。比如 protoc-gen-go 内置了一个 gRPC 插件，用户可以通过 --go_out=plugins=grpc 参数来生成 gRPC 相关代码，否则只会针对 message 生成相关代码。
参考 gRPC 插件的代码，可以发现 generator.RegisterPlugin 函数可以用来注册插件。插件是一个 generator.Plugin 接口：
为了便于维护，我们基于 Go 语言的模板来生成服务代码，其中 tmplService 是服务的模板。
当 Protobuf 的插件定制工作完成后，每次 hello.proto 文件中 RPC 服务的变化都可以自动生成代码。也可以通过更新插件的模板，调整或增加生成代码的内容。在掌握了定制 Protobuf 插件技术后，你将彻底拥有这个技术。
gRPC 是 Google 公司基于 Protobuf 开发的跨语言的开源 RPC 框架。gRPC 基于 HTTP/2 协议设计，可以基于一个 HTTP/2 连接提供多个服务，对于移动设备更加友好。
gRPC 建立在 HTTP/2 协议之上，对 TLS 提供了很好的支持。我们前面章节中 gRPC 的服务都没有提供证书支持，因此客户端在连接服务器中通过 grpc.WithInsecure() 选项跳过了对服务器证书的验证。没有启用证书的 gRPC 服务在和客户端进行的是明文通讯，信息面临被任何第三方监听的风险。为了保障 gRPC 通信不被第三方监听篡改或伪造，我们可以对服务器启动 TLS 加密特性。
前面讲述的基于证书的认证是针对每个 gRPC 连接的认证。gRPC 还为每个 gRPC 方法调用提供了认证支持，这样就基于用户 Token 对不同的方法访问进行权限管理。
要实现对每个 gRPC 方法进行认证，需要实现 grpc.PerRPCCredentials 接口
gRPC 中的 grpc.UnaryInterceptor 和 grpc.StreamInterceptor 分别对普通方法和流方法提供了截取器的支持
gRPC 服务一般用于集群内部通信，如果需要对外暴露服务一般会提供等价的 REST 接口。通过 REST 接口比较方便前端 JavaScript 和后端交互。开源社区中的 grpc-gateway 项目就实现了将 gRPC 服务转为 REST 服务的能力。
最新的 Nginx 对 gRPC 提供了深度支持。可以通过 Nginx 将后端多个 gRPC 服务聚合到一个 Nginx 服务。同时 Nginx 也提供了为同一种 gRPC 服务注册多个后端的功能，这样可以轻松实现 gRPC 负载均衡的支持。Nginx 的 gRPC 扩展是一个较大的主题，感兴趣的读者可以自行参考相关文档。
Protobuf 本身具有反射功能，可以在运行时获取对象的 Proto 文件。gRPC 同样也提供了一个名为 reflection 的反射包，用于为 gRPC 服务提供查询。gRPC 官方提供了一个 C++ 实现的 grpc_cli 工具，可以用于查询 gRPC 列表或调用 gRPC 方法。但是 C++ 版本的 grpc_cli 安装比较复杂，我们推荐用纯 Go 语言实现的 grpcurl 工具。本节将简要介绍 grpcurl 工具的用法。
grpcurl 是 Go 语言开源社区开发的工具，需要手工安装：

$ go get github.com/fullstorydev/grpcurl
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
grpcurl 中最常使用的是 list 命令，用于获取服务或服务方法的列表。比如 grpcurl localhost:1234 list 命令将获取本地 1234 端口上的 grpc 服务的列表。在使用 grpcurl 时，需要通过 -cert 和 -key 参数设置公钥和私钥文件，连接启用了 tls 协议的服务。对于没有没用 tls 协议的 grpc 服务，通过 -plaintext 参数忽略 tls 证书的验证过程。如果是 Unix Socket 协议，则需要指定 -unix 参数。
如果没有配置好公钥和私钥文件，也没有忽略证书的验证过程，那么将会遇到类似以下的错误：

$ grpcurl localhost:1234 list
Failed to dial target host "localhost:1234": tls: first record does not \
look like a TLS handshake
如果 grpc 服务正常，但是服务没有启动 reflection 反射服务，将会遇到以下错误：

$ grpcurl -plaintext localhost:1234 list
Failed to list services: server does not support the reflection API

## 第 5 章 go 和 Web

5.3.4 哪些事情适合在中间件中做
以较流行的开源 Go 语言框架 chi 为例：

compress.go
  => 对 http 的响应体进行压缩处理
heartbeat.go
  => 设置一个特殊的路由，例如 / ping，/healthcheck，用来给负载均衡一类的前置服务进行探活
logger.go
  => 打印请求处理处理日志，例如请求处理时间，请求路由
profiler.go
  => 挂载 pprof 需要的路由，如 `/pprof`、`/pprof/trace` 到系统中
realip.go
  => 从请求头中读取 X-Forwarded-For 和 X-Real-IP，将 http.Request 中的 RemoteAddr 修改为得到的 RealIP
requestid.go
  => 为本次请求生成单独的 requestid，可一路透传，用来生成分布式调用链路，也可用于在日志中串连单次请求的所有逻辑
timeout.go
  => 用 context.Timeout 设置超时时间，并将其通过 http.Request 一路透传下去
throttler.go
  => 通过定长大小的 channel 存储 token，并通过这些 token 对接口进行限流
5.5.1 从 database/sql 讲起
Go 官方提供了 database/sql 包来给用户进行和数据库打交道的工作，database/sql 库实际只提供了一套操作数据库的接口和规范，例如抽象好的 SQL 预处理（prepare），连接池管理，数据绑定，事务，错误处理等等。官方并没有提供具体某种数据库实现的协议支持。
和具体的数据库，例如 MySQL 打交道，还需要再引入 MySQL 的驱动，像下面这样：

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")

import _ "github.com/go-sql-driver/mysql"
这条 import 语句会调用了 mysql 包的 init 函数，做的事情也很简单：

func init() {
    sql.Register("mysql", &MySQLDriver{})
}
我们需要衡量一下这个 Web 服务的吞吐量，再具体一些，就是接口的 QPS。借助 wrk，在家用电脑 Macbook Pro 上对这个 hello world 服务进行基准测试，Mac 的硬件情况如下：

CPU: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
Core: 2
Threads: 4

Graphics/Displays:
      Chipset Model: Intel Iris Graphics 6100
          Resolution: 2560 x 1600 Retina
    Memory Slots:
          Size: 4 GB
          Speed: 1867 MHz
          Size: 4 GB
          Speed: 1867 MHz
Storage:
          Size: 250.14 GB (250,140,319,744 bytes)
          Media Name: APPLE SSD SM0256G Media
          Size: 250.14 GB (250,140,319,744 bytes)
          Medium Type: SSD
测试结果：

~ ❯❯❯ wrk -c 10 -d 10s -t10 http://localhost:9090
Running 10s test @ http://localhost:9090
  10 threads and 10 connections
  Thread Stats   Avg	  Stdev	 Max   +/- Stdev
    Latency   339.99us	1.28ms  44.43ms   98.29%
    Req/Sec	 4.49k   656.81	 7.47k	73.36%
  449588 requests in 10.10s, 54.88MB read
Requests/sec:  44513.22
Transfer/sec:	  5.43MB

~ ❯❯❯ wrk -c 10 -d 10s -t10 http://localhost:9090
Running 10s test @ http://localhost:9090
  10 threads and 10 connections
  Thread Stats   Avg	  Stdev	 Max   +/- Stdev
    Latency   334.76us	1.21ms  45.47ms   98.27%
    Req/Sec	 4.42k   633.62	 6.90k	71.16%
  443582 requests in 10.10s, 54.15MB read
Requests/sec:  43911.68
Transfer/sec:	  5.36MB

~ ❯❯❯ wrk -c 10 -d 10s -t10 http://localhost:9090
Running 10s test @ http://localhost:9090
  10 threads and 10 connections
  Thread Stats   Avg	  Stdev	 Max   +/- Stdev
    Latency   379.26us	1.34ms  44.28ms   97.62%
    Req/Sec	 4.55k   591.64	 8.20k	76.37%
  455710 requests in 10.10s, 55.63MB read
Requests/sec:  45118.57
Transfer/sec:	  5.51MB
多次测试的结果在 4 万左右的 QPS 浮动，响应时间最多也就是 40ms 左右，对于一个 Web 程序来说，这已经是很不错的成绩了，我们只是照抄了别人的示例代码，就完成了一个高性能的 hello world 服务器，是不是很有成就感？
这还只是家用 PC，线上服务器大多都是 24 核心起，32G 内存 +，CPU 基本都是 Intel i7。所以同样的程序在服务器上运行会得到更好的结果。
这里的 hello world 服务没有任何业务逻辑。真实环境的程序要复杂得多，有些程序偏网络 IO 瓶颈，例如一些 CDN 服务、Proxy 服务；有些程序偏 CPU/GPU 瓶颈，例如登陆校验服务、图像处理服务；有些程序瓶颈偏磁盘，例如专门的存储系统，数据库。不同的程序瓶颈会体现在不同的地方，这里提到的这些功能单一的服务相对来说还算容易分析。如果碰到业务逻辑复杂代码量巨大的模块，其瓶颈并不是三下五除二可以推测出来的，还是需要从压力测试中得到更为精确的结论。
对于 IO/Network 瓶颈类的程序，其表现是网卡 / 磁盘 IO 会先于 CPU 打满，这种情况即使优化 CPU 的使用也不能提高整个系统的吞吐量，只能提高磁盘的读写速度，增加内存大小，提升网卡的带宽来提升整体性能。而 CPU 瓶颈类的程序，则是在存储和网卡未打满之前 CPU 占用率先到达 100%，CPU 忙于各种计算任务，IO 设备相对则较闲。
无论哪种类型的服务，在资源使用到极限的时候都会导致请求堆积，超时，系统 hang 死，最终伤害到终端用户。对于分布式的 Web 服务来说，瓶颈还不一定总在系统内部，也有可能在外部。非计算密集型的系统往往会在关系型数据库环节失守，而这时候 Web 模块本身还远远未达到瓶颈。
不管我们的服务瓶颈在哪里，最终要做的事情都是一样的，那就是流量限制。
5.6.3 服务瓶颈和 QoS
前面我们说了很多 CPU 瓶颈、IO 瓶颈之类的概念，这种性能瓶颈从大多数公司都有的监控系统中可以比较快速地定位出来，如果一个系统遇到了性能问题，那监控图的反应一般都是最快的。
虽然性能指标很重要，但对用户提供服务时还应考虑服务整体的 QoS。QoS 全称是 Quality of Service，顾名思义是服务质量。QoS 包含有可用性、吞吐量、时延、时延变化和丢失等指标。一般来讲我们可以通过优化系统，来提高 Web 服务的 CPU 利用率，从而提高整个系统的吞吐量。但吞吐量提高的同时，用户体验是有可能变差的。用户角度比较敏感的除了可用性之外，还有时延。虽然你的系统吞吐量高，但半天刷不开页面，想必会造成大量的用户流失。所以在大公司的 Web 服务性能指标中，除了平均响应时延之外，还会把响应时间的 95 分位，99 分位也拿出来作为性能标准。平均响应在提高 CPU 利用率没受到太大影响时，可能 95 分位、99 分位的响应时间大幅度攀升了，那么这时候就要考虑提高这些 CPU 利用率所付出的代价是否值得了。
在线系统的机器一般都会保持 CPU 有一定的余裕。
在 MyType 定义的地方，不需要 import "io" 就可以直接实现 io.Writer 接口，我们还可以随意地组合很多函数，以实现各种类型的接口，同时接口实现方和接口定义方都不用建立 import 产生的依赖关系。因此很多人认为 Go 的这种正交是一种很优秀的设计。
但这种 “正交” 性也会给我们带来一些麻烦。当我们接手了一个几十万行的系统时，如果看到定义了很多接口，例如订单流程的接口，我们希望能直接找到这些接口都被哪些对象实现了。但直到现在，这个简单的需求也就只有 Goland 实现了，并且体验尚可。Visual Studio Code 则需要对项目进行全局扫描，来看到底有哪些结构体实现了该接口的全部函数。那些显式实现接口的语言，对于 IDE 的接口查找来说就友好多了。另一方面，我们看到一个结构体，也希望能够立刻知道这个结构体实现了哪些接口，但也有着和前面提到的相同的问题。
虽有不便，接口带给我们的好处也是不言而喻的：一是依赖反转，这是接口在大多数语言中对软件项目所能产生的影响，在 Go 的正交接口的设计场景下甚至可以去除依赖；二是由编译器来帮助我们在编译期就能检查到类似 “未完全实现接口” 这样的错误，如果业务未实现某个流程，但又将其实例作为接口强行来使用的话：
5.9.2 通过业务规则进行灰度发布
常见的灰度策略有多种，较为简单的需求，例如我们的策略是要按照千分比来发布，那么我们可以用用户 id、手机号、用户设备信息，等等，来生成一个简单的哈希值，然后再求模，用伪代码表示一下：

// pass 3/1000
func passed() bool {
    key := hashFunctions(userID) % 1000
    if key <= 2 {
        return true
    }

    return false
}
5.9.2.1 可选规则
常见的灰度发布系统会有下列规则提供选择：
	1	按城市发布
	2	按概率发布
	3	按百分比发布
	4	按白名单发布
	5	按业务线发布
	6	按 UA 发布 (APP、Web、PC)
	7	按分发渠道发布

## 第 6 章 分布式系统

## 附录A：Go语言常见坑
这里列举的Go语言常见坑都是符合Go语言语法的，可以正常的编译，但是可能是运行结果错误，或者是有资源泄漏的风险。
可变参数是空接口类型
当参数的可变参数是空接口类型时，传入空接口的切片时需要注意参数展开的问题。
```go
func main() {
    var a = []interface{}{1, 2, 3}

    fmt.Println(a)
    fmt.Println(a...)
}
```
不管是否展开，编译器都无法发现错误，但是输出是不同的：

[1 2 3]
1 2 3

+ 数组是值传递
在函数调用参数中，数组是值传递，无法通过修改数组类型的参数返回结果。
```go
func main() {
    x := [3]int{1, 2, 3}

    func(arr [3]int) {
        arr[0] = 7
        fmt.Println(arr)
    }(x)

    fmt.Println(x)
}
```

必要时需要使用切片。
map遍历是顺序不固定
map是一种hash表实现，每次遍历的顺序都可能不一样。
```go
func main() {
    m := map[string]string{
        "1": "1",
        "2": "2",
        "3": "3",
    }

    for k, v := range m {
        println(k, v)
    }
}
```


+ 返回值被屏蔽
在局部作用域中，命名的返回值内同名的局部变量屏蔽：
```go
func Foo() (err error) {
    if err := Bar(); err != nil {
        return
    }
    return
}
```

+ recover必须在defer函数中运行
recover捕获的是祖父级调用时的异常，直接调用时无效：
```go
func main() {
    recover()
    panic(1)
}
```
直接defer调用也是无效：
```go
func main() {
    defer recover()
    panic(1)
}
```
defer调用时多层嵌套依然无效：
```go
func main() {
    defer func() {
        func() { recover() }()
    }()
    panic(1)
}
```
必须在defer函数中直接调用才有效：
```go
func main() {
    defer func() {
        recover()
    }()
    panic(1)
}
```

+ main函数提前退出
后台Goroutine无法保证完成任务。
```go
func main() {
    go println("hello")
}
```
通过Sleep来回避并发中的问题
休眠并不能保证输出完整的字符串：
```go
func main() {
    go println("hello")
    time.Sleep(time.Second)
}
```
类似的还有通过插入调度语句：
```go
func main() {
    go println("hello")
    runtime.Gosched()
}
```

+ 独占CPU导致其它Goroutine饿死
Goroutine 是协作式抢占调度（Go1.14版本之前），Goroutine本身不会主动放弃CPU：
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()

    for {} // 占用CPU
}
```
解决的方法是在for循环加入runtime.Gosched()调度函数：
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()

    for {
        runtime.Gosched()
    }
}
```
或者是通过阻塞的方式避免CPU占用：
```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        os.Exit(0)
    }()

    select{}
}
```

+ Go1.14 版本引入基于系统信号的异步抢占调度，可以避免 Goroutine 饿死的情况。
不同Goroutine之间不满足顺序一致性内存模型
因为在不同的Goroutine，main函数中无法保证能打印出hello, world:
```go
var msg string
var done bool

func setup() {
    msg = "hello, world"
    done = true
}

func main() {
    go setup()
    for !done {
    }
    println(msg)
}
```
解决的办法是用显式同步：
```go
var msg string
var done = make(chan bool)

func setup() {
    msg = "hello, world"
    done <- true
}

func main() {
    go setup()
    <-done
    println(msg)
}
```
msg的写入是在channel发送之前，所以能保证打印hello, world

+ 闭包错误引用同一个变量
```go
func main() {
    for i := 0; i < 5; i++ {
        defer func() {
            println(i)
        }()
    }
}
```
改进的方法是在每轮迭代中生成一个局部变量：
```go
func main() {
    for i := 0; i < 5; i++ {
        i := i
        defer func() {
            println(i)
        }()
    }
}
```
或者是通过函数参数传入：
```go
func main() {
    for i := 0; i < 5; i++ {
        defer func(i int) {
            println(i)
        }(i)
    }
}
```

+ 在循环内部执行defer语句
defer在函数退出时才能执行，在for执行defer会导致资源延迟释放：
```go
func main() {
    for i := 0; i < 5; i++ {
        f, err := os.Open("/path/to/file")
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
    }
}
```
解决的方法可以在for中构造一个局部函数，在局部函数内部执行defer：
```go
func main() {
    for i := 0; i < 5; i++ {
        func() {
            f, err := os.Open("/path/to/file")
            if err != nil {
                log.Fatal(err)
            }
            defer f.Close()
        }()
    }
}
```

+ 切片会导致整个底层数组被锁定
切片会导致整个底层数组被锁定，底层数组无法释放内存。如果底层数组较大会对内存产生很大的压力。
```go
func main() {
    headerMap := make(map[string][]byte)

    for i := 0; i < 5; i++ {
        name := "/path/to/file"
        data, err := ioutil.ReadFile(name)
        if err != nil {
            log.Fatal(err)
        }
        headerMap[name] = data[:1]
    }

    // do some thing
}
```
解决的方法是将结果克隆一份，这样可以释放底层的数组：
```go
func main() {
    headerMap := make(map[string][]byte)

    for i := 0; i < 5; i++ {
        name := "/path/to/file"
        data, err := ioutil.ReadFile(name)
        if err != nil {
            log.Fatal(err)
        }
        headerMap[name] = append([]byte{}, data[:1]...)
    }

    // do some thing
}
```

+ 空指针和空接口不等价
比如返回了一个错误指针，但是并不是空的error接口：
```go
func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = ErrBad
    }
    return p // Will always return a non-nil error.
}
```

+ 内存地址会变化
Go语言中对象的地址可能发生变化，因此指针不能从其它非指针类型的值生成：
```go
func main() {
    var x int = 42
    var p uintptr = uintptr(unsafe.Pointer(&x))

    runtime.GC()
    var px *int = (*int)(unsafe.Pointer(p))
    println(*px)
}
```
当内存发生变化的时候，相关的指针会同步更新，但是非指针类型的uintptr不会做同步更新。
同理CGO中也不能保存Go对象地址。

+ Goroutine泄露
Go语言是带内存自动回收的特性，因此内存一般不会泄漏。但是Goroutine确存在泄漏的情况，同时泄漏的Goroutine引用的内存同样无法被回收。
```go
func main() {
    ch := func() <-chan int {
        ch := make(chan int)
        go func() {
            for i := 0; ; i++ {
                ch <- i
            }
        } ()
        return ch
    }()

    for v := range ch {
        fmt.Println(v)
        if v == 5 {
            break
        }
    }
}
```
上面的程序中后台Goroutine向管道输入自然数序列，main函数中输出序列。但是当break跳出for循环的时候，后台Goroutine就处于无法被回收的状态了。
我们可以通过context包来避免这个问题：
```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())

    ch := func(ctx context.Context) <-chan int {
        ch := make(chan int)
        go func() {
            for i := 0; ; i++ {
                select {
                case <- ctx.Done():
                    return
                case ch <- i:
                }
            }
        } ()
        return ch
    }(ctx)

    for v := range ch {
        fmt.Println(v)
        if v == 5 {
            cancel()
            break
        }
    }
}
```
当main函数在break跳出循环时，通过调用cancel()来通知后台Goroutine退出，这样就避免了Goroutine的泄漏。

	•	KusonStack一站式可编程配置技术栈(Go): https://github.com/KusionStack/kusion
	•	KCL 配置编程语言(Rust): https://github.com/KusionStack/KCLVM
	•	凹语言™: https://github.com/wa-lang/wa

---    

## 附录B：有趣的代码片段
这里收集一些比较有意思的Go程序片段。

+ 自重写程序
UNIX/Go语言之父 Ken Thompson 在1983年的图灵奖演讲 Reflections on Trusting Trust 就给出了一个C语言的自重写程序。
最短的C语言自重写程序是 Vlad Taeerov 和 Rashit Fakhreyev 的版本：

`main(a){printf(a="main(a){printf(a=%c%s%c,34,a,34);}",34,a,34);}`
下面的Go语言版本自重写程序是 rsc 提供的：
```go
/* Go quine */
package main

import "fmt"

func main() {
    fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `/* Go quine */
```

在 golang-nuts 中还有很多版本：

package main;func main(){c:="package main;func main(){c:=%q;print(c,c)}";print(c,c)}

package main;func main(){print(c+"\x60"+c+"\x60")};var c=`package main;func main(){print(c+"\x60"+c+"\x60")};var c=`
如果有更短的版本欢迎告诉我们。


+ 三元表达式
```go
func If(condition bool, trueVal, falseVal interface{}) interface{} {
    if condition {
        return trueVal
    }
    return falseVal
}

a, b := 2, 3
max := If(a > b, a, b).(int)
println(max)

+ 禁止 main 函数退出的方法

func main() {
    defer func() { for {} }()
}

func main() {
    defer func() { select {} }()
}

func main() {
    defer func() { <-make(chan bool) }()
}
```

+ 基于管道的随机数生成器
随机数的一个特点是不好预测。如果一个随机数的输出是可以简单预测的，那么一般会称为伪随机数。
```go
func main() {
    for i := range random(100) {
        fmt.Println(i)
    }
}

func random(n int) <-chan int {
    c := make(chan int)
    go func() {
        defer close(c)
        for i := 0; i < n; i++ {
            select {
            case c <- 0:
            case c <- 1:
            }
        }
    }()
    return c
}
```

+ 基于select语言特性构造的随机数生成器。

+ Assert测试断言
```go
type testing_TBHelper interface {
    Helper()
}

func Assert(tb testing.TB, condition bool, args ...interface{}) {
    if x, ok := tb.(testing_TBHelper); ok {
        x.Helper() // Go1.9+
    }
    if !condition {
        if msg := fmt.Sprint(args...); msg != "" {
            tb.Fatalf("Assert failed, %s", msg)
        } else {
            tb.Fatalf("Assert failed")
        }
    }
}

func Assertf(tb testing.TB, condition bool, format string, a ...interface{}) {
    if x, ok := tb.(testing_TBHelper); ok {
        x.Helper() // Go1.9+
    }
    if !condition {
        if msg := fmt.Sprintf(format, a...); msg != "" {
            tb.Fatalf("Assertf failed, %s", msg)
        } else {
            tb.Fatalf("Assertf failed")
        }
    }
}

func AssertFunc(tb testing.TB, fn func() error) {
    if x, ok := tb.(testing_TBHelper); ok {
        x.Helper() // Go1.9+
    }
    if err := fn(); err != nil {
        tb.Fatalf("AssertFunc failed, %v", err)
    }
}

```
+ Go语言进阶之路（五）：通道和goroutine、GPM:<https://blog.csdn.net/c315838651/article/details/105008305>

Go语言使用了用户级线程的实现方式，Go语言中的goroutine可以理解为用户态的线程，调度切换goroutine直接在用户态进行，不用切换到内核态。

Go 调度器模型我们通常叫做GPM 模型，包括 4 个重要结构：

G:Goroutine，每个 Goroutine 对应一个 G 结构体，我们使用go关键字创建goroutine，并非就一定创建了G结构体的实例，只有当没有可用的G时，才会创建G来装载我们创建的goroutine，否则，会复用现有可用的G来装载goroutine。G 存储 Goroutine 的运行堆栈、状态以及任务函数，可重用。G 并非执行体，每个 G 需要绑定到 P 才能被调度执行。

P: Processor，表示逻辑处理器，对 G 来说，P 相当于 CPU 核心，G 只有绑定到 P 才能被调度。对 M 来说，P 提供了相关的执行环境(Context)，如内存分配状态(mcache)，任务队列(G)等。P 的数量决定了系统内最大可并行的 G 的数量（前提：物理 CPU 核数 >= P 的数量）。P 的数量由用户设置的 GOMAXPROCS 决定，但是不论 GOMAXPROCS 设置为多大，P 的数量最大为 256。

M: Machine，OS 内核线程抽象，代表着真正执行计算的资源，在绑定有效的 P 后，进入 schedule 循环；而 schedule 循环的机制大致是从 Global 队列、P 的 Local 队列以及 wait 队列中获取。M 的数量是不定的，由 Go Runtime 调整，为了防止创建过多 OS 线程导致系统调度不过来，目前默认最大限制为 10000 个。M 并不保留 G 状态，这是 G 可以跨 M 调度的基础。

Sched：Go 调度器，它维护有存储 M 和 G 的队列以及调度器的一些状态信息等。
————————————————
版权声明：本文为CSDN博主「程序猿架构」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/c315838651/article/details/105008305


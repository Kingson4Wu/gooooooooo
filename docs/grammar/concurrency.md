+ Java 和 Go 在并发实现中的区别:<https://mp.weixin.qq.com/s/FWqSDNNA4tj9Rpf9m97jaQ>

+ Java 使用 OS 线程通过 Java 运行时管理的线程来完成并行进程。Golang 通过 Goroutine 使用线程 os 完成并行进程。在 Java 和 Golang 之间的并行概念中，没有太大区别，几乎相同，只是名称不同。
+ 并发概念不同。在 Java 中，这是通过将 Java 运行时线程映射到 os线 程来完成的。同时，golang 使用此 goroutine 映射到 golang 上的调度程序，将其映射到更深层次。
+ Goroutine 本身不是线程系统或由 os 线程管理的线程。但是，这种想法更多的是将函数协程处理到 os 线程中的多路复用器。这样，当发生进程阻塞时，它将被移至未使用的线程或绿色线程，Go 调度程序的任务是将此绿色线程映射到 OS 线程，并将 goroutine 分配到空闲的绿色线程中。

乍一看，goroutine 概念与 Reactive .io 中以 Reactore 3 或 RxJava 表示的 React Java 的非阻塞概念几乎相同。但是，Java反 应流概念比 goroutines 具有更高级的方法。


### Java 并发模型和 Golang
尽管在并发问题中有不同的实现方法，但是模型几乎相同。

#### 异步过程
Java

创建从 Thread 类扩展的类。
实现 Runnable 的接口。
Golang

Goroutine 开始

#### 同步过程
Java

方法上的同步块。
使用 java.util.concurrent 包中的 Lock.ReentrantLock

Golang

使用通道的概念，即术语“不通过共享内存进行通信和通过通信共享内存”的行话的实现。
Sync.Mutex 锁定资源。

#### 进程间通讯
Java

使用 object.wait()，object.Notify() 或 object.NotifyAll() 方法。
在多个线程上共享块队列
使用 PipeReader 和 PipeWriter 方法

Golang

使用 channel
使用 WaitGroup
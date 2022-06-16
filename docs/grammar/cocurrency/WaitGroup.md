+ sync.WaitGroup 用于阻塞等待一组 Go 程的结束。主 Go 程调用 Add() 来设置等待的 Go 程数，然后该组中的每个 Go 程都需要在运行结束时调用 Done()， 递减 WaitGroup 的 Go 程计数器 counter。当 counter 变为 0 时，主 Go 程被唤醒继续执行。

+ 如果使用过程中通过 Add()添加的 Go 程数与调用 Done() 的次数不符，即 sync.WaitGroup 的 Go 程计数器等所有子 Go 程结束后不为 0，则会引发 panic。

+ Done()过多
当 Go 程计数器变为负数时，将引发 panic。

+ Done() 过少
在最后一个活动线程 foo2 退出的时候，Go 检测到当前没有还在运行的 Go 程，但主 Go 程仍在等待，发生了死锁现象，于是引发 panic，这是 Go 的一种自我保护机制。


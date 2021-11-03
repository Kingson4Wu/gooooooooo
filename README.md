
GOROOT 为go 安装目录，此处brew安装，进入go目录之后可看到bin目录是在libexec下面，所以goroot为上述

如果是源码，可以按照习惯直接复制到/usr/local/go下面，那么此时GOROOT 为/usr/local/go

GOPATH 为go的工作目录，即我们的code目录

 ---

 Go 语言的未来是不可限量的。当然，我个人觉得，Go 可能会吞食很多 C、C++、Java 的项目。不过，Go 语言所吞食的项目应该主要是中间层的项目，既不是非常底层也不会是业务层。

也就是说，Go 语言不会吞食底层到 C 和 C++ 那个级别的，也不会吞食到上层如 Java 业务层的项目。Go 语言能吞食的一定是 PaaS 上的项目，比如一些消息缓存中间件、服务发现、服务代理、控制系统、Agent、日志收集等等，他们没有复杂的业务场景，也到不了特别底层（如操作系统）的软件项目或工具。而 C 和 C++ 会被打到更底层，Java 会被打到更上层的业务层。这是我的一个判断。

(摘自陈皓)

---

Go项目组织实践
https://zhuanlan.zhihu.com/p/124198314
https://blog.csdn.net/jmilk/article/details/107285314
https://juejin.cn/post/6944649692319842340

---

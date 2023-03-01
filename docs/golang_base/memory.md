+ golang内存泄漏，服务怎么限制最大内存呢:<https://www.zhihu.com/answer/2869462136>
    - 内存泄露应该优先在部署环境、代码逻辑上分析问题，限制最大内存并不是好的方案。
对于 golang 服务，通常是以容器方式部署在服务器上，你可以通过其命令行工具中提供的参数来限制 Cpu、Mem 等资源的使用。
如果你没有使用容器技术，比如是部署在 Linux 操作系统中，也可以通过在你的 systemd 启动文件中设置 MemoryMax 来限制内存使用。（和容器方式类似，内部都是通过 cgroup 实现）
如果你已经升级到 golang 1.19 版本，启动程序时可以加上以下两个变量。
GOMEMLIMIT=xxx GOGC=xxx
具体可参考: https://pkg.go.dev/runtime


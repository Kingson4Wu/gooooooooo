https://github.com/agiledragon/gomonkey/issues/70

Line 19: - permission denied
goroutine 4 [running]:

发生在

err := syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
if err != nil {
    panic(err)
}
环境

MacOS 11.6.1 (ARM64 M1 以及 X86都如此)
运行GoMonkey的最新Demo
使用的gomonkey 2.2.0
描述:

使用GOARCH=amd64可以解决, 但是无法debug很不幸.
在Linux/Windows都是正常
macos-golink-wrapper 方案可以解决x86上此问题, 但是arm64依然不行
如果使用单元测试框架结果导致无法断点调试也是太影响了


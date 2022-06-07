+ `go get <package name>`
    1. 从远程下载需要用到的包
    2. 执行go install
+ go install
    - go install 会生成可执行文件直接放到bin目录下;如果是一个普通的包，会被编译生成到pkg目录下该文件是.a结尾  

+ go build
    - go build -o helloWorld.exe helloWorld.go

+ go run
    - 直接运行命令go run hello.go来执行程序

<pre>
go build、go install 和 go run 的区别
详细的可以查看参考资料4，这里简单说一下：

go build 编译包，如果是main包则在当前目录生成可执行文件，其他包不会生成.a文件；
go install 编译包，同时复制结果到$GOPATH/bin，$GOPATH/pkg等对应目录下；
go run gofiles... 编译列出的文件，并生成可执行文件然后执行。注意只能用于main包，否则会出现go run: cannot run non-main package的错误。
此外，go run是不需要设置$GOPATH的，但go build和go install必须设置。go run常用来测试一些功能，这些代码一般不包含在最终的项目中。
</pre>    

+ go env -w GOPROXY=https://goproxy.cn,direct

+ 一个新项目，你会执行如下命令：
`$ go mod init <project name>`
这是初始化项目，会生成 go.mod 文件

+ 移除了一些依赖或增加一些依赖
`$ go mod tidy`
记住：你不需要手动修改 go.mod 文件。

+ go vet
该命令检查 Go 源代码并报告可疑的情况，例如 Printf 调用，其参数与格式字符串不对齐。Vet 使用的启发式方法不能保证所有报告都是真实的问题，但它可以发现编译器没有捕获到的错误。

这里是该命令的官方文档：https://pkg.go.dev/cmd/vet。



+ go get
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

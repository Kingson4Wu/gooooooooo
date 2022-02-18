+ GOROOT：go的安装路径
+ GOPATH: go的工作路径
+ Go 1.8 版本之前，GOPATH 环境变量默认是空的；1.8版本之后，默认路径是：$HOME/go
<pre>
src: 存放源代码（比如：.go .c .h .s等）
pkg: 编译后生成的文件（比如：.a）
bin: 编译后生成的可执行文件, 为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录
</pre>

go mod init github.com/kingson4wu/gooooooooo
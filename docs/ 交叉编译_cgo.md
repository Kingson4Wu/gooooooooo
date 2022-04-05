+ Go语言涉及CGO的交叉编译(跨平台编译)解决办法:<https://zhuanlan.zhihu.com/p/338891206>

+ CGO_ENABLED 这个参数默认为1，开启CGO。需要指定为0来关闭，因为CGO不支持交叉编译。

1. 交叉编译：跨平台编译
2. CGO：golang调用C/C++库


+ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static -fpic"' main.go
CGO_ENABLED 这个参数默认为1，开启CGO。需要指定为0来关闭，因为CGO不支持交叉编译。

GOOS 和 GOARCH 用来指定要构建的平台为Linux

可选参数-ldflags 是编译选项：

-s -w 去掉调试信息，可以减小构建后文件体积，
--extldflags "-static -fpic" 完全静态编译[2]，这样编译生成的文件就可以任意放到指定平台下运行，而不需要运行环境配置。
显然对于带CGO的交叉编译，CGO_ENABLED必须开启。这也就需要辅助编译器来帮我们实现交叉编译了。

---

如果你是mac平台，可以用这个工具 FiloSottile/musl-cross/musl-cross 直接通过brew安装就可以使用

brew install FiloSottile/musl-cross/musl-cross
安装成功后，有多种编译器可以帮我们实现交叉编译，使用时只需在编译对应参数下指定就可以了。

----
https://zhuanlan.zhihu.com/p/343551786

docker方式，未尝试！！！

----

https://www.cnblogs.com/yubs/p/14929638.html
https://github.com/messense/homebrew-macos-cross-toolchains
 
brew tap messense/macos-cross-toolchains
 
brew install x86_64-unknown-linux-gnu
 
 brew install aarch64-unknown-linux-gnu
　　

编译

1
CGO_ENABLED=1 GOOS=linux CC=x86_64-unknown-linux-gnu-gcc CXX=x86_64-unknown-linux-gnu-g++ go build -a -installsuffix cgo -o app .
　　
不支持arm64！！


---

https://juejin.cn/post/6844904166851084296

go version
go version go1.17.7 darwin/amd64
docker build -t gobuilder:1.17.7-stretch .

`docker run --rm -it -v ~/Personal/go-src/weixin-app/:/go/src/app  -v ~/Downloads/:/go/output gobuilder:1.17.7-stretch`

env GOOS=linux GOARCH=arm64 go build  -o ~/Downloads/weixinapp ./cmd/main.go


`docker run --rm -it -v ~/Personal/go-src/weixin-app/:/go/src/app  -v ~/Downloads/:/go/output gobuilder:1.17.7-stretch`
# runtime/cgo
gcc_linux_arm64.c: In function 'x_cgo_init':
gcc_linux_arm64.c:89:15: error: cast from pointer to integer of different size [-Werror=pointer-to-int-cast]
  g->stacklo = (uintptr)&size - size + 4096;
               ^
cc1: all warnings being treated as errors


如果用到CGO, 就不建议使用交叉编译. 自己用虚拟机搭一个编译环境. 在windows下即便交叉编译通过, copy到linux上,我也遇到过不能正常运行的情况.

https://www.golangtc.com/t/568f7b02b09ecc72d5000013
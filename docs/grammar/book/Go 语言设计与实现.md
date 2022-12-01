这本书适合深入,进阶、学习底层原理的时候看,

TODO

----




+ git clone  --depth=1  https://github.com/golang/go.git

### 1.1 调试源代码 

+ `vi src/fmt/print.go`
+ 修改
+ `./src/make.bash`
make.bash must be run from $GOROOT/src

+ 将 Go 语言的源代码编译成汇编语言: `go build -gcflags -S main.go`

+ 下面的命令获取汇编指令的优化过程：

`$ GOSSAFUNC=main go build main.go`

+ https://draveness.me/golang/docs/part1-prerequisite/ch01-prepare/golang-debug/

## 第二章 编译原理

+  https://draveness.me/golang/docs/part1-prerequisite/ch02-compile/golang-compile-intro/

# 第二部分 基础知识

## 第三章 数据结构


# 第三部分 运行时

## 第六章 并发编程

### 6.1 上下文 Context

https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

### 6.6 网络轮询器

https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-netpoller/

# 第四部分 进阶内容

## 第八章 元编程

### 8.1 插件系统


https://draveness.me/golang/docs/part4-advanced/ch08-metaprogramming/golang-plugin/

#### 8.1.1 设计原理 #
Go 语言的插件系统基于 C 语言动态库实现的，所以它也继承了 C 语言动态库的优点和缺点，我们在本节中会对比 Linux 中的静态库和动态库，分析它们各自的特点和优势。

静态库或者静态链接库是由编译期决定的程序、外部函数和变量构成的，编译器或者链接器会将程序和变量等内容拷贝到目标的应用并生成一个独立的可执行对象文件1；
动态库或者共享对象可以在多个可执行文件之间共享，程序使用的模块会在运行时从共享对象中加载，而不是在编译程序时打包成独立的可执行文件2；
由于特性不同，静态库和动态库的优缺点也比较明显；只依赖静态库并且通过静态链接生成的二进制文件因为包含了全部的依赖，所以能够独立执行，但是编译的结果也比较大；而动态库可以在多个可执行文件之间共享，可以减少内存的占用，其链接的过程往往也都是在装载或者运行期间触发的，所以可以包含一些可以热插拔的模块并降低内存的占用。

使用静态链接编译二进制文件在部署上有非常明显的优势，最终的编译产物也可以直接运行在大多数的机器上，静态链接带来的部署优势远比更低的内存占用显得重要，所以很多编程语言包括 Go 都将静态链接作为默认的链接方式。

插件系统 #
在今天，动态链接带来的低内存占用优势虽然已经没有太多作用，但是动态链接的机制却可以为我们提供更多的灵活性，主程序可以在编译后动态加载共享库实现热插拔的插件系统。





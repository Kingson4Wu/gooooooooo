+ 如何阅读Golang的源码:<https://www.zhihu.com/answer/756625130>

+ 安装包的 src/ 

+ go env
+ GOROOT="/usr/local/Cellar/go/1.18/libexec"
+ /usr/local/Cellar/go/1.18/libexec/src


----


+ Go内建函数源码在哪里： https://zhuanlan.zhihu.com/p/337575452
    - go doc builtin.int也可以查看预定义标识符int的文档
    - 内建函数的源码在哪里？作为预声明标识符子集的内建函数们在builtin.go中也都有自己的位置
    - 内建函数仅仅是一个标识符，在Go源码编译期间，Go编译器遇到内建函数标识符时会将其替换为若干runtime的调用，我们还以append函数为例，我们输出下面代码的汇编代码(Go 1.14)：
    ```go
    // append.go
    package main
    
    import "fmt"

    func main() {
    var s = []int{5, 6}
    s = append(s, 7, 8)
    fmt.Println(s)
    }
    ```
    `$go tool compile -S append.go > append.s`


+ Go内建关键字的源码，以select为例
    + /usr/local/go/src/cmd/compile/internal/walk/select.go
    + /usr/local/go/src/runtime/select.go
    + 核心函数：selectgo()

    + /usr/local/go/src/go/token/token.go (自己全局搜出来的)
    + cmd/compile/internal/walk/stmt.go
    ```go
            case ir.OSELECT:
            n := n.(*ir.SelectStmt)
            walkSelect(n)
            return n
    ```
    + cmd/compile/internal/ir/fmt.go
    ```go
    var OpNames = []string{
        OSELECT:           "select",
    }
    ```
    大概率是这句 对应到相应的关键字



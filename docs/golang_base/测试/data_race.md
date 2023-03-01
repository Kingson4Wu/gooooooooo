+ Go：从一个data race问题学到的:<https://mp.weixin.qq.com/s/P_bPVzfZZhpokoLXllGxdw>

```go
package main

import "time"

func main() {
 running := true
 go func() {
  println("start thread1")
  count := 1
  for running {
   count++
  }
  println("end thread1: count =", count)
 }()
 go func() {
  println("start thread2")
  for {
   running = false
  }
 }()
 time.Sleep(time.Hour)
}
```


+ 问题代码中的循环之所以不会结束，和所谓的「CPU 缓存一致性中的线程可见性问题」并没有任何关系，只是因为编译器把部分代码看成死代码，直接优化掉了，这个过程称之为「Dead code elimination[5]」，不过当激活 race 检测的时候，编译器并没有执行死代码的优化，所以程序看上去又正常了。

+ plan9 assembly 完全解析 :<https://segmentfault.com/a/1190000039978109>
+ 工具 SSA


---

https://cloud.tencent.com/developer/article/1489437
谈谈 Golang 中的 Data Race（续）
发布于2019-08-19 14:50:06阅读 6140
我在上一篇文章中曾指出：在 Go 的内存模型中，有 race 的 Go 程序的行为是未定义行为，理论上出现什么情况都是正常的。并尝试通过一段有 data race 的代码来说明问题：

package main

import (
    "fmt"
    "runtime"
    "time"
)

var i = 0

func main() {
    runtime.GOMAXPROCS(2)

    go func() {
        for {
            fmt.Println("i is", i)
            time.Sleep(time.Second)
        }
    }()

    for {
        i += 1
    }
}
复制
当通过 go run cmd.go 执行时，大概率会得到下面这样的输出：

i is: 0
i is: 0
i is: 0
i is: 0
复制
然而有些同学提到：之所以输出 0 是因为 i+=1 所在的 goroutine 没有新的栈帧创建，因此没有被调度器调度到。解释似乎也合理，但是事实却不是这样的。真实的原因是：编译器把那段自增的 for 循环给全部优化掉了。

要验证这一点，我们要先从编译器优化说起。传统的编译器通常分为三个部分，前端(frontEnd)，优化器(Optimizer)和后端(backEnd)。在编译过程中，前端主要负责词法和语法分析，将源代码转化为抽象语法树；优化器则是在前端的基础上，对得到的中间代码进行优化，使代码更加高效；后端则是将已经优化的中间代码转化为针对各自平台的机器代码。

go 的编译器也一样，在生成目标代码的时候会做很多优化，重要的有：

指令重排
逃逸分析
函数内联
死码消除
当我们通过:

go build cmd.go
go tool objdump -s main.main cmd
复制
查看编译出的二进制可执行文件的汇编代码：

cmd.go:11        0x4858c0        64488b0c25f8ffffff  MOVQ FS:0xfffffff8, CX
  cmd.go:11        0x4858c9        483b6110        CMPQ 0x10(CX), SP
  cmd.go:11        0x4858cd        7635            JBE 0x485904
  cmd.go:11        0x4858cf        4883ec18        SUBQ $0x18, SP
  cmd.go:11        0x4858d3        48896c2410      MOVQ BP, 0x10(SP)
  cmd.go:11        0x4858d8        488d6c2410      LEAQ 0x10(SP), BP
  cmd.go:12        0x4858dd        48c7042402000000    MOVQ $0x2, 0(SP)
  cmd.go:12        0x4858e5        e83605f8ff      CALL runtime.GOMAXPROCS(SB)
  cmd.go:14        0x4858ea        c7042400000000      MOVL $0x0, 0(SP)
  cmd.go:14        0x4858f1        488d05a8640300      LEAQ 0x364a8(IP), AX
  cmd.go:14        0x4858f8        4889442408      MOVQ AX, 0x8(SP)
  cmd.go:14        0x4858fd        e89eaafaff      CALL runtime.newproc(SB)
  cmd.go:22        0x485902        ebfe            JMP 0x485902
  cmd.go:11        0x485904        e8e79afcff      CALL runtime.morestack_noctxt(SB)
  cmd.go:11        0x485909        ebb5            JMP main.main(SB)
复制
显然，下面这一段直接被优化没了：

for {
    i += 1
}
复制
why? 因为这段代码是有竞态的，没有任何同步机制。go 编译器认为这一段是 dead code，索性直接优化掉了。

而当我们通过 go build-race cmd.go 编译后：

cmd.go:22        0x4d3430        488d05c9211100      LEAQ main.i(SB), AX
  cmd.go:22        0x4d3437        48890424        MOVQ AX, 0(SP)
  cmd.go:22        0x4d343b        e8d096faff      CALL runtime.raceread(SB)
  cmd.go:22        0x4d3440        488b05b9211100      MOVQ main.i(SB), AX
  cmd.go:22        0x4d3447        4889442410      MOVQ AX, 0x10(SP)
  cmd.go:22        0x4d344c        488d0dad211100      LEAQ main.i(SB), CX
  cmd.go:22        0x4d3453        48890c24        MOVQ CX, 0(SP)
  cmd.go:22        0x4d3457        e8f496faff      CALL runtime.racewrite(SB)
  cmd.go:22        0x4d345c        488b442410      MOVQ 0x10(SP), AX
  cmd.go:22        0x4d3461        48ffc0          INCQ AX
  cmd.go:22        0x4d3464        48890595211100      MOVQ AX, main.i(SB)
复制
可以明显看到有 INCQ 指令了，这是因为 -race 选项打开了 data race detector 用来检查这个错误而关闭了相关的编译器优化：

==================
WARNING: DATA RACE
Read at 0x0000005e5600 by goroutine 6:
  main.main.func1()
      /root/gofourge/src/lab/cmd.go:16 +0x63

Previous write at 0x0000005e5600 by main goroutine:
  main.main()
      /root/gofourge/src/lab/cmd.go:22 +0x7b

Goroutine 6 (running) created at:
  main.main()
      /root/gofourge/src/lab/cmd.go:14 +0x4f
==================
i is: 4085
i is: 56001323
i is: 112465799
i is: 168640611
复制
如此，运行结果就“看似正确”了。

最后再引用一句 golang-nuts 上的评论：

Any race is a bug. When there is a race, the compiler is free to do whatever it wants.

参考资料
Go compiler - Loop transformations
Would this race condition be considered a bug?


+ 因为这段代码是有竞态的，没有任何同步机制。go 编译器认为这一段是 dead code，索性直接优化掉了。 !!!!


----


go test -race ./...
syscall.ptrace: nosplit stack over 792 byte limit
syscall.ptrace<1>
    grows 120 bytes, calls syscall.ptrace1<1>

这个错误是因为在运行 Go 的测试时，您使用了 -race 标志来检测数据竞争问题。在检测过程中，运行时系统会使用一个更小的栈大小来运行每个协程，并使用更严格的栈大小限制来检测竞争条件。这可能会导致您的程序在运行时出现栈溢出的问题，尤其是在使用系统调用（例如 ptrace）时。

为了解决这个问题，您可以尝试以下几种方法：

增加栈大小

您可以通过设置 GOMAXPROCS 和 GODEBUG 环境变量，来增加每个协程的栈大小。例如：

bash
Copy code
export GOMAXPROCS=4
export GODEBUG="asyncpreemptoff=1,asyncpreemptoff=0,GOROOT_FINAL=1,stacksize=16384"
在这个示例中，我们将每个协程的栈大小增加到 16KB。

使用 cgo 调用系统调用

如果您的程序使用了系统调用（例如 ptrace），则可能需要使用 cgo 调用这些调用，而不是直接在 Go 代码中使用它们。cgo 允许您编写 C 或 C++ 代码来调用系统调用，并在 Go 代码中使用它们。这样可以避免栈溢出问题，因为 C 或 C++ 代码使用的栈比 Go 代码使用的栈要小得多。

在使用 cgo 时，请确保遵循官方文档中的最佳实践，以避免其他问题。

减少每个协程使用的栈大小

您还可以尝试减少每个协程使用的栈大小。为此，可以使用 -gcflags 标志并将 -l 标志与其结合使用。例如：

bash
Copy code
go test -race -gcflags "-l" ./...
-l 标志可以让 Go 编译器在生成代码时使用更少的栈空间。

请注意，这些解决方案并不能保证完全解决问题。如果您的程序依赖于大量的系统调用，并且它们使用的栈空间很大，那么即使采用了上述解决方案，仍然可能会出现栈溢出的问题。


+ 深入Go的底层，带你走近一群有追求的人:<https://mp.weixin.qq.com/s/obnnVkO2EiFnuXk_AIDHWw>


---

+ Go 使用的 plan9 汇编语言初探:<https://kcode.icu/posts/go/2021-03-20-go-%E4%BD%BF%E7%94%A8%E7%9A%84-plan9-%E6%B1%87%E7%BC%96%E8%AF%AD%E8%A8%80%E5%88%9D%E6%8E%A2/>
    - 在命令行中执行 `go tool compile -S -N -l main.go` 查看汇编代码
    - 总结
plan9 是一个操作系统，对应一套汇编语法，被称为 plan9 汇编，与 AT&T 规范非常相似，也不完全一样。与传统 x86 汇编差异较大。

plan9 重新定义了 4 个伪寄存器来方便编程，函数嵌套调用的核心就是依靠栈的先进先出特性，在调用前来保护 Caller 的执行现场，并在 Callee 执行完返回后恢复现场，从而继续执行。

没有 push，pop 等栈操作，栈的增缩是通过对 sp 栈顶寄存器的加减实现的，例如 +8（sp）
操作数的大小不是通过区分操作码来实现的，而是通过操作码的后缀实现，例如复制 ax 寄存器低 1 字节到第二字节，在 x86 汇编中是通过 mov ah,al，而在 plan9 中则是通过 movb al,ah
b w d q分别代表1 2 4 8 Byte，含义是byte，word，double word，quatury。
操作数在第一个参数，结果放在第二个参数，与intel汇编相反，具体看第二条
我们学习 plan9 汇编其实也并没有太大困难，前期在学习时也未必要看懂全部的汇编代码，可先学习大概思路，再深入了解即可。

但是为什么 plan9 不与 x86 系兼容呢，而要单独开发一套自己的语法呢，我看了不少人的解释，貌似是因为 unix 团队是学院派，特立独行，也不屑于商业化（就是玩


---

A Quick Guide to Go's Assembler: https://go.dev/doc/asm

`GOOS=linux GOARCH=amd64 go tool compile -S x.go        # or: go build -gcflags -S x.go`

To see what gets put in the binary after linking, use go tool objdump:

$ go build -o x.exe x.go
$ go tool objdump -s main.main x.exe
TEXT main.main(SB) /tmp/x.go
  x.go:3		0x10501c0		65488b0c2530000000	MOVQ GS:0x30, CX
  x.go:3		0x10501c9		483b6110		CMPQ 0x10(CX), SP
  x.go:3		0x10501cd		7634			JBE 0x1050203
  x.go:3		0x10501cf		4883ec10		SUBQ $0x10, SP
  x.go:3		0x10501d3		48896c2408		MOVQ BP, 0x8(SP)
  x.go:3		0x10501d8		488d6c2408		LEAQ 0x8(SP), BP
  x.go:4		0x10501dd		e86e45fdff		CALL runtime.printlock(SB)
  x.go:4		0x10501e2		48c7042403000000	MOVQ $0x3, 0(SP)
  x.go:4		0x10501ea		e8e14cfdff		CALL runtime.printint(SB)
  x.go:4		0x10501ef		e8ec47fdff		CALL runtime.printnl(SB)
  x.go:4		0x10501f4		e8d745fdff		CALL runtime.printunlock(SB)
  x.go:5		0x10501f9		488b6c2408		MOVQ 0x8(SP), BP
  x.go:5		0x10501fe		4883c410		ADDQ $0x10, SP
  x.go:5		0x1050202		c3			RET
  x.go:3		0x1050203		e83882ffff		CALL runtime.morestack_noctxt(SB)
  x.go:3		0x1050208		ebb6			JMP main.main(SB)


-----

了解Go第一步：Go与Plan 9汇编语言: <https://bioitblog.com/blog/2022/03/17/go-and-asm/>

阅读Go汇编常用的命令为go tool compile -N -l -S 。-N代表不优化，不然Go汇编和我们想象的可能大不一样，-l为不内联，-S为打印汇编信息。

----

golang内核系列--深入理解plan9汇编&实践： https://zhuanlan.zhihu.com/p/56750445

编译/反编译工具
实践出真知，很多时候我们无法确定一块代码是如何执行的，需要通过生成汇编、反汇编来研究golang。这里给一些工具来帮助我们了解golang

// 编译
go build -gcflags="-S"
go tool compile -S hello.go
go tool compile -N -S hello.go // 禁止优化
// 反编译
go tool objdump <binary>

----

book： https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-01-basic.html


用 Delve 调试 Go 汇编程序的过程比调试 Go 语言程序更加简单。调试汇编程序时，我们需要时刻关注寄存器的状态，如果涉及函数调用或局部变量或参数还需要重点关注栈寄存器 SP 的状态。

-----

https://lrita.github.io/images/posts/go/GoFunctionsInAssembly.pdf


---

A Quick Guide to Go's Assembler: https://golang.org/doc/asm
Rob Pike, How to Use the Plan 9 C Compiler:http://doc.cat-v.org/plan_9/2nd_edition/papers/comp
Rob Pike, A Manual for the Plan 9 assembler:https://9p.io/sys/doc/asm.html
Debugging Go Code with GDB:https://golang.org/doc/gdb

----






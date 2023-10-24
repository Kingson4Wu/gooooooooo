+ https://mp.weixin.qq.com/s/zbV2UwHYvgsuwnXR2Vr04Q
+ 
“Organizing a Go module”的文档: https://go.dev/doc/modules/layout

[7] 
官方指南: https://go.dev/doc/modules/layout

1. Go项目的类型
我们知道Go项目(project)一般有两类：library和executable。library是以构建库为目的的Go项目，而executable则是以构建二进制可执行文件为目的的Go项目。

“Organizing a Go module”这篇文档也是按照Go项目类型为Gopher提供项目布局建议的。这篇文档将library类的项目叫作package类，executable类的项目叫作command。

2. 官方版Go项目目录布局指南
2.1 basic package

project-root-directory/
├── go.mod
├── modname.go
└── modname_test.go

或

project-root-directory/
├── go.mod
├── modname.go
├── modname_test.go
├── auth.go
├── auth_test.go
├── hash.go
└── hash_test.go

2.2 basic command
project-root-directory/
├── go.mod
└── main.go

或

project-root-directory/
├── go.mod
├── main.go
├── auth.go
├── auth_test.go
├── hash.go
└── hash_test.go

如果这个repo的url为github.com/someuser/modname，那么我们可以通过下面命令安装这个command的可执行程序：

$go install github.com/someuser/modname@latest

2.3 package with supporting packages

这样Go官方建议将这些supporting packages放入internal目录[8]。

project-root-directory/
├── go.mod
├── modname.go
├── modname_test.go
└── internal/
    ├── auth/
    │   ├── auth.go
    │   └── auth_test.go
    └── hash/
        ├── hash.go
        └── hash_test.go

2.4 command with supporting packages

project-root-directory/
├── go.mod
├── main.go
└── internal/
    ├── auth/
    │   ├── auth.go
    │   └── auth_test.go
    └── hash/
        ├── hash.go
        └── hash_test.go

和package with supporting packages不同的是，main.go使用的包名为main，这样Go编译器才能将其构建为command。

2.5 multiple packages

project-root-directory/
├── go.mod
├── modname.go
├── modname_test.go
├── auth/
│   ├── auth.go
│   ├── auth_test.go
│   └── token/
│       ├── token.go
│       └── token_test.go
├── hash/
│   ├── hash.go
│   └── hash_test.go
└── internal/
    └── trace/
        ├── trace.go
        └── trace_test.go

2.6 multiple commands

project-root-directory/
├── go.mod
├── prog1/
│   └── main.go
├── prog2/
│   └── main.go
└── internal/
    └── trace/
        ├── trace.go
        └── trace_test.go


可以通过下面步骤来编译command：

$go build github.com/someuser/modname/prog1
$go build github.com/someuser/modname/prog2
command的用户通过下面步骤可以安装这些命令：

$go install github.com/someuser/modname/prog1@latest
$go install github.com/someuser/modname/prog2@latest

2.7 multiple packages and commands

project-root-directory/
├── go.mod
├── modname.go
├── modname_test.go
├── auth/
│   ├── auth.go
│   ├── auth_test.go
│   └── token/
│       ├── token.go
│       └── token_test.go
├── hash/
│   ├── hash.go
│   └── hash_test.go
├── internal/
│       └── trace/
│           ├── trace.go
│           └── trace_test.go
└── cmd/
    ├── prog1/
    │   └── main.go
    └── prog2/
        └── main.go

为了区分导出package和command，这个示例增加了一个专门用来存放command的cmd目录，prog1和prog2两个command都放在这个目录下。这也是Go语言的一个惯例。

这个示例项目既导出了下面的包：

github.com/user/modname
github.com/user/modname/auth
github.com/user/modname/hash
又包含了两个可安装使用的command，用户按下面步骤安装即可：

$go install github.com/someuser/modname/cmd/prog1@latest
$go install github.com/someuser/modname/cmd/prog2@latest

3. 小结
经过对“Organizing a Go module”的文档[10]这篇Go官方项目目录布局指南的学习，我发现指南中的建议与我个人在以往文章、书和专栏中对Go项目目录布局的建议非常相近，几乎一致，唯独不同的是在pkg目录的使用上。

在multiple packages类型项目中，如果要导出的package非常多，那么项目顶层目下会有大量的目录，这让项目顶层目录显得很“臃肿”，我个人建议将这些导出包统一放置到project-root-directory/pkg下面，这样项目顶层目录就会显得很简洁。





+ GOROOT、GOPATH、GOBIN、project目录:<https://www.jianshu.com/p/bd19a607b5a8>
+ GOBIN
go install编译存放路径。不允许设置多个路径。可以为空。为空时则遵循“约定优于配置”原则，可执行文件放在各自GOPATH目录的bin文件夹中（前提是：package main的main函数文件不能直接放到GOPATH的src下面。

#### 关于go的整体一个开发目录
<pre>
go_project     // go_project为GOPATH目录
  -- bin
     -- myApp1  // 编译生成
     -- myApp2  // 编译生成
     -- myApp3  // 编译生成
  -- pkg
  -- src
     -- myApp1     // project1
        -- models
        -- controllers
        -- others
        -- main.go 
     -- myApp2     // project2
        -- models
        -- controllers
        -- others
        -- main.go 
     -- myApp3     // project3
        -- models
        -- controllers
        -- others
        -- main.go 
</pre>

+ Go项目的目录结构详解:<https://www.jb51.net/article/56746.htm>

+ 现在有go module 了， 不用这样子了！！！！

#### go package
+ Go 语言的源码复用建立在包（package）基础之上。Go 语言的入口 main() 函数所在的包（package）叫 main，main 包想要引用别的代码，必须同样以包的方式进行引用，本章内容将详细讲解如何导出包的内容及如何导入其他包。
+ Go 语言的包与文件夹一一对应，所有与包相关的操作，必须依赖于工作目录（GOPATH）。
+ 关于golang中包（package）的二三事儿:<https://www.cnblogs.com/dajianshi/p/3596492.html>
+ 在Go语言中，和java的main是有所区别的，具体区别如下：
在java中，任何一个java文件都可以有唯一一个main方法当做启动函数
在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件
也就是说，一个package下，只能有一个main方法，不管是在那个文件中，但是只能有一个，这个package是按照文件夹区分的

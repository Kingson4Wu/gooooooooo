+ GOROOT：go的安装路径
+ GOPATH: go的工作路径
+ Go 1.8 版本之前，GOPATH 环境变量默认是空的；1.8版本之后，默认路径是：$HOME/go
<pre>
src: 存放源代码（比如：.go .c .h .s等）
pkg: 编译后生成的文件（比如：.a）
bin: 编译后生成的可执行文件, 为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录
</pre>

go mod init github.com/kingson4wu/gooooooooo

---

go mod tidy

go mod vendor

---

+ Go mod 七宗罪：<https://jishuin.proginn.com/p/763bfbd56b9d>


---

### go.sum
+ <https://www.topgoer.cn/docs/gozhuanjia/chapter123.10-module-go-sum>

+ go.sum 是使用 Go Modules 时自动生成的，是否需要把它放到 .gitignore 文件里排除？

官方的建议是要将 go.sum 和 go.mod 两个文件一起提交到代码库中，这样才能保证项目依赖包版本的一致，同时保证 Build 的一致性。

————————————————
原文作者：Summer
转自链接：https://learnku.com/go/t/39186
版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。

https://github.com/golang/go/wiki/Modules#should-i-commit-my-gosum-file-as-well-as-my-gomod-file


https://stackoverflow.com/questions/53837919/should-go-sum-file-be-checked-in-to-the-git-repository
---

+ Go语言重新开始，Go Modules的前世今生与基本使用:<https://mp.weixin.qq.com/s/0FU8YrPhb2ezc5Is7wVpNQ>

+ Go项目组织：在单一repo中管理多个Go module指南:<https://mp.weixin.qq.com/s/IEsgLu5PFwcKppUWABA6-Q>

----

在单个仓库中支持多个 go mod 模块： https://zhuanlan.zhihu.com/p/134184461

版本发布
对模块进行发版时，只需打上 [模块名]/版本号 即可 以我们的示例为例，对模块 a 进行发版时我们只需要打上 tag a/v1.0.0 ，同理对模块 b 进行发版时，需要打上 tag b/v1.0.0 即可





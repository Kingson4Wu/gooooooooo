+ 关于visual studio code无法安装golang插件的解决办法:<https://blog.csdn.net/dmt742055597/article/details/85865186>

<pre>
创建github.com插件目录及下载插件
cd $GOPATH/src
mkdir github.com
cd $GOPATH/src/github.com
mkdir acroca cweill derekparker go-delve josharian karrick mdempsky pkg ramya-rao-a rogpeppe sqs uudashr 
cd $GOPATH/src/github.com/acroca
git clone https://github.com/acroca/go-symbols.git
cd $GOPATH/src/github.com/cweill
git clone https://github.com/cweill/gotests.git
cd $GOPATH/src/github.com/derekparker
git clone https://github.com/derekparker/delve.git
cd $GOPATH/src/github.com/go-delve
git clone https://github.com/go-delve\delve.git
cd $GOPATH/src/github.com/josharian
git clone https://github.com/josharian/impl.git
cd $GOPATH/src/github.com/karrick
git clone https://github.com/karrick/godirwalk.git
cd $GOPATH/src/github.com/mdempsky
git clone https://github.com/mdempsky/gocode.git
cd $GOPATH/src/github.com/pkg
git clone https://github.com/pkg/errors.git
cd $GOPATH/src/github.com/ramya-rao-a
git clone https://github.com/ramya-rao-a/go-outline.git
cd $GOPATH/src/github.com/rogpeppe
git clone https://github.com/rogpeppe/godef.git
cd $GOPATH/src/github.com/sqs
git clone https://github.com/sqs/goreturns.git
cd $GOPATH/src/github.com/uudashr
git clone https://github.com/uudashr/gopkgs.git

创建golang.org插件目录及下载插件
cd $GOPATH/src
mkdir -p golang.org/x
cd golang.org/x
git clone https://github.com/golang/tools.git
git clone https://github.com/golang/lint.git

手动安装插件
cd $GOPATH/src
go install github.com/mdempsky/gocode
go install github.com/uudashr/gopkgs/cmd/gopkgs
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns
go install github.com/derekparker/delve/cmd/dlv
go install github.com/cweill/gotests
go install github.com/josharian/impl
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename
go install golang.org/x/lint/golint

go install github.com/stamblerre/gocode

多重启几下
</pre>



### 快捷键
 ＋ F9 断点
 ＋ 代码格式化可通过以下快捷方式：
Windows Shift + Alt + F
在Mac上Shift + Option + F

### Code Runner
+ 代码一键运行
<pre>
键盘快捷键 Ctrl+Alt+N
快捷键 F1 调出 命令面板, 然后输入 Run Code
在编辑区，右键选择 Run Code
在左侧的文件管理器，右键选择 Run Code
右上角的运行小三角按钮
</pre>
+ 停止代码运行
<pre>
键盘快捷键 Ctrl+Alt+M
快捷键 F1 调出 命令面板, 然后输入 Stop Code Run
在Output Channel，右键选择 Stop Code Run
</pre>
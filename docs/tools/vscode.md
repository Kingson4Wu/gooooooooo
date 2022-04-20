+ 关于visual studio code无法安装golang插件的解决办法:<https://blog.csdn.net/dmt742055597/article/details/85865186>

+ VSCode Remote SSH
Remote SSH 只支持 Linux 作为服务器，且必须是 64 位版本。

在进行嵌入式Linux开发的时候，为了方便，通常在Windows上使用代码编辑器编辑代码，交叉编译工具在Linux虚拟机或者服务器上，在开发期间需要不停的进行如下的循环操作：

编辑好代码，使用基于SSH的SCP将文件上传到服务器；
使用SSH远程终端，在服务器上编译出可执行文件；
编译完成后使用基于SSH的SCP将文件传回到本地；

https://cloud.tencent.com/developer/article/1726694


+ VSCode 开发 Go 程序也可以和 GoLand 一样强大:<https://mp.weixin.qq.com/s/J01LY7s6xMB8Lk10sxTFhg>

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

+ GitLens 可以增强 VSCode 内置 Git 的功能。例如 commits 搜索，历史记录和显示的代码作者身份


### 快捷键
+ 查看快捷键方式：prefereneces -> KeyboardShortcuts (cmd k cmd s)
+ F9 断点
+ 代码格式化可通过以下快捷方式：
Windows Shift + Alt + F
在Mac上Shift + Option + F
+ 条件断点（点击 add function breakpoint）
+ 同时打开多个项目，cmd＋shift＋n
+ 上一步 Ctrl + -
+ 下一步 (冲突，自定义) crtl ＋＝
+ teminal : ctrl + `
+ cmd+P 找文件等操作


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

+ 注释代码
1、单行注释，使用"Ctrl + /"，或者先按"CTRL+K"，再按"CTRL+U"；
2、块注释，使用"Alt+Shift+A"

### Plugins
+ Bracket Pair Colorizer / Bracket Pair Colorizer 2
+ Highlight Matching Tag这也是一个找对象的插件，找的是标签的对象
+ Material Theme Icons /VSCode Great Icons
+ Code Spell Checker
+ Git History / GitLens — Git supercharged
+ Local History
+ Partial Diff
+ Project Manager
+ json2ts
+ Settings Sync


+ Go 语言环境安装：如何使用 Go 多版本（Mac）:<https://learnku.com/go/wikis/61549>

+ 使用以下方法安装最新版本：
`$ brew install go`

+ 使用以下方法安装指定版本：
`$ brew install go@1.15`

+ 首先 unlink：

`$ brew unlink go`
Unlinking /usr/local/Cellar/go/1.17.1... 0 symlinks removed.
link 指定版本：

`$ brew link go@1.15`



---

+ Go卸载重装和版本升级:<https://www.jianshu.com/p/79f84695cec7>
+ `export GOROOT=/usr/local/Cellar/go/1.18/libexec`


---

+ 深入理解 go build 和 go install:<https://juejin.cn/post/6844903938060222471>
+ Go 1.18新特性解读（万字长文）:<https://mp.weixin.qq.com/s/8CyoGLuepuCI4Hj1Ev0GcQ>

+ Go 1.18明确了能修改go.mod、go.sum的命令只有三个：go get、go mod tidy和go mod download。这样开发人员就可以放心的在项目根目录下执行Go工具链提供的其他命令了。

+ 引入Go workspace（工作区）
+ Go fuzzing
+ 内置函数append对切片的扩容算法发生变化
Go 1.18中的算法有一定变化，目的是使得在一个门槛值前后的变化更丝滑。具体算法大家看下面$GOROOT/src/runtime/slice.go中的growslice函数中的部分逻辑
+ 新增net/netip包
+ strings包和bytes包新增Cut函数
+ sync包新增Mutex.TryLock、RWMutex.TryLock和RWMutex.TryRLock

---

+ Go 1.18 Release Note 中文版:<https://mp.weixin.qq.com/s/idXVZC3Qwfkmh30YYtXcWQ>

+ Go1.18 快讯：Module 工作区模式，太棒了:<https://mp.weixin.qq.com/s/S3I919YZb-bgaEnHgKq7fg>

我们当然可以提交 mypkg 到 github，但我们没修改一次 mypkg，就需要提交，否则 example 中就没法使用上最新的。

针对这种情况，目前是建议通过 replace 来解决，即在 example 中的 go.mod 增加如下 replace：（v1.0.0 根据具体情况修改，还未提交，可以使用 v1.0.0）

`replace github.com/polaris1119/mypkg => ../mypkg`

当都开发完成时，我们需要手动删除 replace，并执行 go mod tidy 后提交，否则别人使用就报错了。

+ 工作区模式

go.work 看看长什么样：

go 1.18

directory (
    ./example
    ./mypkg
)

go.work 不需要提交到 Git 中，因为它只是你本地开发使用的



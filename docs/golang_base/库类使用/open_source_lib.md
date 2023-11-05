+ https://golangrepo.com/
+ https://go.libhunt.com/
+ https://libs.garden/go

---

+ 有哪有优秀的golang库？ - Go语言入门到精通 的回答 - 知乎:<https://www.zhihu.com/answer/2481047615> !!!! TODO
+ 有哪些值得学习的 Go 语言开源项目？:<https://www.zhihu.com/question/20801814>
+ https://github.com/avelino/awesome-go

+ 有没有推荐的golang的练手项目？ - 终端研发部 的回答:<https://www.zhihu.com/answer/2232959861>

---

+ Go官方的限流器 time/rate
`golang.org/x/time/rate`
<https://zhuanlan.zhihu.com/p/89820414>

+ 每日一库之一个兼具单机和集群模式的轻量级限流器 throttled
https://github.com/throttled/throttled

+ 用 Cobra 构建命令行工具
`go get -u github.com/spf13/cobra/cobra`

+ 依赖注入 wire
<https://darjun.github.io/2020/03/02/godailylib/wire/>

+ https://github.com/go-redis/redis
+ https://github.com/go-gorm/gorm

+ goutil/dump
这是一个 golang 数据打印工具包，可以打印漂亮易读的 go slice、map、struct 数据
`dump.P(vars…)` 使用简单，直接调用即可

+ testing是 Go 语言标准库自带的测试库
+ testify可以说是最流行的（从 GitHub star 数来看）Go 语言测试库了。testify提供了很多方便的函数帮助我们做assert和错误信息输出。使用标准库testing，我们需要自己编写各种条件判断，根据判断结果决定输出对应的信息。

+ github.com/golang/mock/gomock
+ github.com/DATA-DOG/go-sqlmock 结合 gorm 来进行 mock

+ Mockery (https://github.com/vektra/mockery) 是一个补充工具(Golang 的模拟代码自动生成器)，可以和Testify的Mock包一起使用，利用它的功能来创建模拟对象，并为项目源代码中的Golang接口自动生成模拟对象。它有助于减少很多模板式的编码任务。

+ Gofakeit（https://github.com/brianvoe/gofakeit）是一个随机数据生成器。目前，它提供了160多个函数，涵盖少数不同的主题/类别，如Person、 Animals、 Address、 Games、 Cars、 Beers等等。这里的关键词是随机。随机性是计划和编写测试时需要考虑的一个重要方面。与其在测试用例中使用常量值，不如使用随机生成的值，这样我们会更有信心，即使面对未知的（随机的）输入，我们的代码仍然会以我们期望的方式表现出来。

+ 一种命令行解析的新思路（Go 语言描述）:<https://mp.weixin.qq.com/s/RxpcqBGhUT-5z4N5kRXvBg>
Go 语言中使用最广泛功能最强大的命令行解析库是 cobra，但丰富的功能让 cobra 相比标准库的 flag 而言，变得异常复杂，为了减少使用的复杂度，cobra 甚至提供了代码生成的功能，可以自动生成命令行的骨架。然而，自动生成在节省了开发时间的同时，也让代码变得不够直观。
+ Go 实战：构建漂亮的 CLI 应用程序:<https://mp.weixin.qq.com/s/JZR62ulwFPxYfRZJr-yrvA>
    - go-pretty

+ Gock（https://github.com/h2non/gock）是一个用于Golang的HTTP服务器模拟和期望库，它在很大程度上受到了NodeJs的流行和较早的同类库的启发，称为Nock[2]。与Pact-go（https://github.com/pact-foundation/pact-go）不同，它是一个轻量级的解决方案，通过http.DefaultTransport或任何http.Client使用的http.Transport拦截出站请求。

+ zap是uber开源的日志库，选择zap他有两个优势：
它非常的快
它同时提供了结构化日志记录和printf风格的日志记录

+ jsoniter

做业务开发离不开json的序列化与反序列化，标准库虽然提供了encoding/json，但是它主要是通过反射来实现的，所以性能消耗比较大。jsoniter可以解决这个痛点，其是一款快且灵活的 JSON 解析器，具有良好的性能并能100%兼容标准库

+ robfig/cron

github地址：https://github.com/robfig/cron
业务开发更离不开定时器的使用了，cron就是一个用于管理定时任务的库，用 Go 实现 Linux 中crontab这个命令的效果，与Linux 中crontab命令相似，cron库支持用 5 个空格分隔的域来表示时间。

+ ants

某些业务场景还会使用到goroutine池，ants就是一个广泛使用的goroute池，可以有效控制协程数量，防止协程过多影响程序性能。ants也是国人开发的，设计博文写的也很详细的，目前很多大厂也都在使用ants，经历过线上业务检验的，所以可以放心使用。

https://github.com/panjf2000/ants

+ gorilla/mux是 gorilla Web 开发工具包中的路由管理库。

+ gopsutil是 Python 工具库psutil 的 Golang 移植版，可以帮助我们方便地获取各种系统和硬件信息。

+ go-streams : 用于 Go 的轻量级流处理库，提供了一种简单简洁的 DSL 来构建数据管道。

+ redis官网推荐的go版本的分布式锁：redsync
+ go get github.com/go-redsync/redsync/v4

+ 日期管理
8、Carbon：这是一个轻量级的、易于使用的、语义智能的日期时间库，适用于Go开发者。
地址：https://github.com/golang-module/carbon

+ web爬虫
12、Colly：这是一个很棒的Go网页爬虫框架，特别适合存档(我经常用它)和数据挖掘。
地址：https://github.com/gocolly/colly

+ 13、Retry
用于重试逻辑和回退。它是高度可扩展的，可以完全控制重试发生的方式和时间。还可以通过实现backoff接口编写自己的自定义后退函数。
地址：https://github.com/sethvargo/go-retry

+ 14、endless
golang HTTP和HTTPS服务器的零停机重启。
地址：https://github.com/fvbock/endless



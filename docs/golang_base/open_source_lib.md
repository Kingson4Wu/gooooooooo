+ https://golangrepo.com/
+ https://go.libhunt.com/
+ https://libs.garden/go

---

+ Go官方的限流器 time/rate
`golang.org/x/time/rate`
<https://zhuanlan.zhihu.com/p/89820414>

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


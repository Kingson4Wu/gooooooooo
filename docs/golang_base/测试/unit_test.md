+ Golang 单元测试详尽指引:<https://mp.weixin.qq.com/s/eAptnygPQcQ5Ex8-6l0byA>

+ GoConvey 和其他 Stub/Mock 框架的兼容性相比 Testify 更好。
Testify 自带 Mock 框架，但是用这个框架 Mock 类需要自己写。像这样重复有规律的部分在 GoMock 中是一键自动生成的。

### Stub/Mock 框架
Golang 有以下 Stub/Mock 框架：

GoStub
GoMock
Monkey
一般来说，GoConvey 可以和 GoStub、GoMock、Monkey 中的一个或多个搭配使用。

Testify 本身有自己的 Mock 框架，可以用自己的也可以和这里列出来的 Stub/Mock 框架搭配使用。

### 项目总结
在编写桩模块时会发现，模块之间的调用关系在工程规模并不大的本案例中，也依然比较复杂，需要开发相应桩函数，代码量会增加许多，也会消耗一些开发人员的时间，因此反推到之前流程的开发实践中，可以得出结论就是提高模块的内聚度可简化单元测试，如果每个模块只能完成一个，所需测试用例数目将显著减少，模块中的错误也更容易发现。

Go 单元测试框架是相当易用的，其它的第三方库基本都是建立在 testing 原生框架的基础上进行的增补扩充，在日常开发中，原生包可以满足基本需求，但同样也有缺陷，原生包不提供断言的语法使得代码中的这种片段非常多：

if err != nil{
 //...
}
所以引入了 convey、assert 包的断言语句，用于简化判断逻辑，使得程序更加易读。

在完成项目单测时，遇到了不少问题，比较重要的比如由于架构分层不够清晰，还是有部分耦合代码，导致单测时需要屏蔽的模块太多，代码写起来不便。因此还是需要倒推到开发模块之前，就要设计更好的结构，在开发的过程中遵循相应的规则，通过测试先行的思想，使开发的工程具有更好的可测试性。

开发过程中遇到的场景肯定不局限于本文所讨论的范围，有关更丰富的最佳实践案例可以参照：

go-sqlmock
go-mock

----

+ ​手把手教你如何进行 Golang 单元测试:<https://mp.weixin.qq.com/s/N5wby-aWWEPc7mHN_lN3lQ>

由于我们把单测简单的分为了两种：

对于无第三方依赖的纯逻辑代码，我们只需要验证相关逻辑即可，这里只需要使用 assert （断言），通过控制输入输出比对结果即可。

对于有第三方依赖的代码，在验证相关代码逻辑之前，我们需要将相关的依赖 mock （模拟），之后才能通过断言验证逻辑。这里需要借助第三方工具库来处理。

因此，对于 assert （断言）工具，可以选择 testify 或 convery，笔者这里选择了 testify。对于 mock （模拟）工具，笔者这里选择了 gomock 和 gomonkey。关于 mock 工具同时使用 gomock 和 gomonkey，这里跟 Golang 的语言特性有关，下面会详细的说明。


---


## testify
+ https://cloud.tencent.com/developer/article/1869961
+ test/test/testify_test.go

## gomock

+ https://geektutu.com/post/quick-gomock.html
+ test/test/gomock_test.go

+ gomock 是官方提供的 mock 框架，同时还提供了 mockgen 工具用来辅助生成测试代码。

+ go install github.com/golang/mock/mockgen@latest
+ mockgen -source=db.go -destination=db_mock.go -package=test

## gomonkey

+ https://www.cnblogs.com/lanyangsh/p/14587921.html

+ 很强!!!

+ gomonkey支持任何场景的打桩，而gomock仅支持对interface的打桩

+ gomonkey支持为private method打桩了:<https://www.jianshu.com/p/7546e788613b>

+ https://www.jianshu.com/u/1381dc29fed9

## goconvey
+ https://www.jianshu.com/p/1bd1ece2fa38
+ 要使用GoConvey的web界面，这时需要提前安装GoConvey的二进制，命令为go install github.com/smartystreets/goconvey@latest
+ go get github.com/smartystreets/goconvey/convey@v1.7.2

## sqlmock
+ https://cloud.tencent.com/developer/article/1881962
+ https://blog.csdn.net/lanyang123456/article/details/123303324

## miniredis
+ https://github.com/alicebob/miniredis
+ go get github.com/alicebob/miniredis/v2
+ https://golang-tech-stack.com/post/4329




---


GoConvey(断言、嵌套) 
Gomonkey(全局变量、函数、方法、接口等等打桩) 
sqlmock(DB模拟)
miniredis(redis模拟)

---
+ Gin 优化那些事:<https://mp.weixin.qq.com/s/dIf3g9CT6Ih9_VfL1N0rFA>

+ 原生 http 实现和 gin 对比:<https://blog.csdn.net/baidu_32452525/article/details/117138171>

net/http 和 gin 的关系
net/http
net/http 是 Go 语言中原生的 http 实现，可以提供 http 服务器的功能，其中默认的 DefaultServeMux 提供了基础的路有功能。
net/http 提供了良好的抽象：Server，Listener，Conn，HandlerFunc，Handler 定义了一整套http 请求的处理流程。
net/http 也存在一些问题
请求响应编解码繁琐
默认的mutex性能问题
时间复杂度： o(n) + 正则匹配
没有中间件、监控支持
不太好的内存管理
request/response无法复用（请求级别）
无条件的解析请求头
gin
实现了http.Handler接口的轻量级框架
提供了高性能的路由：Radix Tree 实现
提供工具简化了输入输出处理：binding 处理
提供了中间件的支持
提供web服务的常用工具函数，如panic捕获，json格式校验等
使用context池，减少runtime的GC工作量。
强大的工具包: gin.Context
gin.Context 提供了一系列解析、校验请求的方法，其中内置了 validator 参数校验
————————————————
版权声明：本文为CSDN博主「CoLiuRs」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/baidu_32452525/article/details/117138171


----

https://gin-gonic.com/zh-cn/docs/
https://gin-gonic.com/docs/

https://gin-gonic.com/zh-cn/docs/examples/

https://gin-gonic.com/zh-cn/docs/examples/goroutines-inside-a-middleware/


```go
// 创建在 goroutine 中使用的副本
		cCp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
```

https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/

https://gin-gonic.com/zh-cn/docs/examples/custom-validators/

https://gin-gonic.com/zh-cn/docs/examples/run-multiple-service/

https://gin-gonic.com/zh-cn/docs/examples/redirects/



---


Slots []Slot `json:"slots" binding:"required,min=1,dive"` 解释一下这段代码

这段代码是一个Go语言的结构体定义。让我们逐步解释它：

1. `Slots`: 这是结构体的名称，表示一个槽（slot）的集合。
2. `[]Slot`: 这表示一个`Slot`类型的切片（slice），即一个包含多个`Slot`对象的列表。
3. ``json:"slots" binding:"required,min=1,dive"``: 这是结构体字段的标签（tag），用于提供额外的元数据信息。在这个例子中，标签使用了两个键值对，用于定义结构体字段的JSON序列化和验证行为。

   - `json:"slots"`: 这个键值对指示在JSON序列化过程中，将该字段命名为"slots"。这是有用的，例如，当通过网络传输数据时，可以将结构体字段映射为特定的JSON键。
   - `binding:"required,min=1,dive"`: 这个键值对指定了对该字段进行验证的规则。在这里，使用了`binding`包提供的规则。具体来说：
     - `required` 表示该字段是必需的，即不能为零值或空值。
     - `min=1` 表示该字段必须至少包含一个元素。
     - `dive` 表示验证应该递归地应用于该字段的每个元素。

因此，这段代码定义了一个名为`Slots`的结构体，该结构体包含一个`Slot`类型的切片字段，并通过标签定义了该字段的JSON序列化和验证行为。验证规则要求该字段在序列化和验证过程中必须至少包含一个非零值的元素。


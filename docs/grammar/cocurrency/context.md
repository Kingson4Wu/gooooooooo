+ 深入浅出golang——context：<https://zhuanlan.zhihu.com/p/381116234>

+ golang中Context的使用场景有哪些吗
场景一：RPC调用

在主goroutine上有4个RPC，RPC2/3/4是并行请求的，我们这里希望在RPC2请求失败之后，直接返回错误，并且让RPC3/4停止继续计算。这个时候，就使用的到Context。

场景二：PipeLine

场景三：超时请求

我们发送RPC请求的时候，往往希望对这个请求进行一个超时的限制。当一个RPC请求超过10s的请求，自动断开。当然我们使用CancelContext，也能实现这个功能（开启一个新的goroutine，这个goroutine拿着cancel函数，当时间到了，就调用cancel函数）。

场景四：HTTP服务器的request互相传递数据

context还提供了valueCtx的数据结构。

----

+ 走进Golang之Context的使用:<https://cloud.tencent.com/developer/article/1676355>


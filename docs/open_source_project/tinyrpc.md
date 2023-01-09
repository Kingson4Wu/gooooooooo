+ https://github.com/zehuamama/tinyrpc
+ TinyRPC 是基于Go语言标准库 net/rpc 扩展的远程过程调用框架

基于TCP传输层协议支持多种压缩格式：gzip、snappy、zlib；
基于二进制的 Protocol Buffer 序列化协议：
具有协议编码小及高扩展性和跨平台性；
支持生成工具：TinyRPC提供的 protoc-gen-tinyrpc 插件可以帮助开发者快速定义自己的服务；

TinyRPC 的源代码仅有一千行左右，通过学习 TinyRPC ，开发者可以得到以下收获：
代码简洁规范
涵盖大多数 Go 语言基础用法和高级特性
单元测试编写技巧
TCP流中处理数据包的技巧
RPC框架的设计理念

作者：马丸子
链接：https://www.zhihu.com/question/369863905/answer/2444168149
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

---

+ net.rpc包的使用
https://studygolang.com/articles/34486


+ proto gen 插件自定义

+ gRPC服务的响应设计:<https://mp.weixin.qq.com/s/iN1mnp0tdEjwpbEzRy3ApA>
+ gRPC的错误处理实践:<https://mp.weixin.qq.com/s?__biz=Mzg5Mjc3MjIyMA==&mid=2247544582&idx=1&sn=59cf489d5406d062916a63d3ee63aa98&source=41#wechat_redirect>


+ gRPC客户端的那些事儿:<https://mp.weixin.qq.com/s/9eFtjNMSVxxq2qPPN7aGBA>
    - 在本文中我们了解了gRPC的四种通信模式。我们重点关注了在最常用的simple RPC(unary RPC)模式下gRPC Client侧需要考虑的事情，包括：
        1. 如何实现一个helloworld的一对一的通信
        2. 如何实现一个自定义的Resolver以实现一个client到一个静态服务实例集合的通信
        3. 如何实现一个自定义的Resolver以实现一个client到一个动态服务实例集合的通信
        4. 如何自定义客户端Balancer
+ gRPC入门指南 — 客户端流式RPC（三）:<https://mp.weixin.qq.com/s/BTe8-_6oP5DOw3D5j4DQ9Q>

---

+ Go - 使用工具生成易读的 Protocol 文档:<https://mp.weixin.qq.com/s/iHMUoGg-d_NBkq9majYQTw>
    - protoc-gen-doc

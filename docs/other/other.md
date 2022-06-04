
+ Go 通过 Map/Filter/ForEach 等流式 API 高效处理数据:<https://mp.weixin.qq.com/s/t3INtSfFSmv-nsJqLmdPew>
    - 至此 Stream 组件就全部实现完了，核心逻辑是利用 channel 当做管道，数据当做水流，不断的用协程接收/写入数据到 channel 中达到异步非阻塞的效果。
    - 回到开篇提到的问题，未动手前想要实现一个 stream 难度似乎非常大，很难想象在 go 中 300 多行的代码就能实现如此强大的组件。
    - 实现高效的基础来源三个语言特性：
        - channel   
        - 协程
        - 函数式编程
    - https://github.com/zeromicro/go-zero

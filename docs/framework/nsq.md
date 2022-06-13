+ NSQ 消息队列实现消息落地使用的是 FIFO 队列。
实现为 diskqueue , 使用包 github.com/nsqio/go-diskqueue

+ 特性总结
消息投放是不保序的
原因是内存队列、持久化队列、以及重新消费的数据混合在一起消费导致的
多个consumer 订阅同一个channel，消息将随机发送到不同的consumer 上
消息是可靠的
当消息发送出去之后，会进入in_flight_queue 队列
当恢复FIN 之后，才会从队列中将消费成功的消息清除
如果客户端发送REQ，消息将会重发
消息发送采用的是推模式，减少延迟
支持延迟消费的模式: DPUB, 或者 RRQ (消费未成功，延时消费) 命令


nsq不丢消息？？？TODO


nsq_to_nsq
nsq 作为消息队列，有个优势是nsqd 各节点之间是不关联的，如果一个节点出了问题，仅仅影响该节点下的topic，channel，以及相关的生产者、消费者。 也就是官方说明的特性第一条：no SPOF ( single point of failure 单点故障)。好处不言而喻，坏处也是有的，如果节点出问题，没有备份数据无法恢复。

所以，在官方提供了 nsq_to_nsq 作为 nsqd 节点复制的工具，用于做 nsqd 节点数据的备份, 或者也可以用于数据的分发。


总结
由于nsqd 本身是不保序的，因此nsq_to_nsq 也是此特性，在复制数据和分发的时候，如果有多个接收的nsqd，并不能保证消息分发到相同的nsqd，因此无法保序。

nsq_to_file
除了使用nsq_to_nsq做节点备份外，也可以通过数据落地的方式，做消息的物理备份。
nsq_to_file 可以将nsq接收到的数据，落地到硬盘。如需数据恢复，可以通过读取文件数据，重新生产即可。



单机模式部署
NSQD 是可以脱离 nsqlookup 做单机部署的。
由于 nsqd 足够轻量，可以把服务部署在消息发布的服务器上，加快 pub 消息的速度，也能兼顾消费端消息的分发

集群模式
NSQD 是一个SPOF的系统，每个服务可以独立部署。当采用集群模式时，建议开启nsqlookup服务，用于管理多个 nsqd 的服务

一般的消息队列都会提供rebalance 的功能，nsqd 是没有的。
不过可以通过nsq_to_nsq 做消息的复制，做服务的主备，当服务挂机后，可以切换到另外的服务器做消费。（中间channel 不会切换，因此可能会重复消费，或者丢一定消息）
nsqd 正常情况下，如果配置合理，消息是不会落地的。如果需要落地，可以使用nsq_to_file, 新建一个channel订阅 相关topic, 把消息落地到硬盘。

在集群模式下，可以部署多个 nsqlookupd 服务, 这些服务之间是互相没有依赖的，nsqd 在做消息广播的时候，会对每一个nsqlookupd的服务遍历一次，更新服务上的信息

消费模式！！！

client通过lookup连接nsqd消费数据！！！！！

nsqd
    
    每个channel的消息都会进行排队，直到一个worker把他们消费，如果此队列超出了内存限制，消息将会被写入到磁盘中。Nsqd节点首先会向nsqlookup广播他们的位置信息，一旦它们注册成功，worker将会从nsqlookup服务器节点上发现所有包含事件topic的nsqd节点。

作者：yikejiucai
链接：https://juejin.cn/post/6844903502360084494
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



----

### nsq client sdk java版
#### producer
+ 使用netty和nsqd建立tcp连接
+ client连接后保持一定数量的连接对象Connection
+ com.trendrr.nsq.Connection#commandAndWait
+ 每个Connection包含两个LinkedBlockingQueue（request和response），容量都是1（同时只接受一个请求），通过offer和poll入队和出队阻塞（带超时时间）；发送topic往request写对象，然后写netty channel通道对象，channel通过接受到响应时候，offer到response，主线程通过response.poll判断是否发送完成，最后通过request.poll清空对象，以便接收新的发送请求。
+ 通过LinkedBlockingQueue巧妙达到锁的效果。通过netty异步化发送接收过程。
+ 每个nsqd地址创建三个连接对象Connection
+ 延时队列（ deferredPQ ） - 通过 DPUB 发布消息时带有 timeout 属性值实现，表示从当前时间戳多久后可以取出来消费；
#### consumer
+ 通过lookup API 获取注册上去的所有nsqd
+ 每个nsqd地址创建三个连接对象Connection
+ SUB - 消费者订阅 Topic/Channel
+ RDY - 客户端连接就绪



+ <https://xiaomi-info.github.io/2019/12/06/nsq-src/> !!!
+ <https://juejin.cn/post/6979256037970804766>
+ <https://zhuanlan.zhihu.com/p/323679681>

### nsq client sdk go版
+ https://github.com/nsqio/go-nsq


```go

func (t *Topic) put(m *Message) error {
    select {
    case t.memoryMsgChan <- m:  //内存队列
    default:
        b := bufferPoolGet()
        err := writeMessageToBackend(b, m, t.backend) //磁盘队列
        bufferPoolPut(b)
        t.ctx.nsqd.SetHealth(err)
        if err != nil {
            t.ctx.nsqd.logf(
                "TOPIC(%s) ERROR: failed to write message to backend - %s",
                t.name, err)
            return err
        }
    }
    return nil
}

```

---

+ NSQ源码剖析——简单高性能的MQ实现：<https://zhuanlan.zhihu.com/p/152243465> !!!

writeMessageToBackend是一个多态函数，它根据初始化backend的类型选择采用持久化还是临时对象来执行调用Put方法。临时对象在缓冲区满的情况下会被丢弃。


消息的持久化
默认的情况下，只有内存队列不足时MemQueueSize:10000时，才会把数据保存到文件内进行持久到硬盘
如果将 --mem-queue-size 设置为 0，所有的消息将会存储到磁盘。我们不用担心消息会丢失，nsq 内部机制保证在程序关闭时将队列中的数据持久化到硬盘，重启后就会恢复。

+ NSQ 消息队列实现消息落地使用的是 FIFO 队列。
实现为 diskqueue , 使用包 github.com/nsqio/go-diskqueue ,



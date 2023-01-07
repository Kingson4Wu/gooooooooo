+ https://cloud.tencent.com/developer/article/1967978
+ golang中内置的可以直接用来做本地缓存的无非就是map和sync.Map。而这两者中，map是非并发安全的数据结构，在使用时需要加锁；而sync.Map虽然是线程安全的。但是需要在并发读写时加锁。此外二者均无法支持数据的过期和淘汰，同时在存储大量数据时，又会产生比较频繁的GC问题，更严重的情况下导致线上服务无法稳定运行。

+ 本地缓存组件中，实现零GC的方案主要就两种：
a.无GC：分配堆外内存(Mmap)
b.避免GC：map非指针优化(map[uint64]uint32)或者采用slice实现一套无指针的map
c.避免GC：数据存入[]byte slice(可考虑底层采用环形队列封装循环使用空间)



----

使用map就会GC？？？看一下GC原理TODO！！！

----

+ go: 实现简单的内存级别缓存:<https://cloud.tencent.com/developer/article/1595117>

```go
package cache

import (
    "sync"
    "sync/atomic"
    "time"
)

var globalMap sync.Map
var len int64

func Set(key string, data interface{}, timeout int) {
    globalMap.Store(key, data)
    atomic.AddInt64(&len, 1)
    time.AfterFunc(time.Second*time.Duration(timeout), func() {
        atomic.AddInt64(&len, -1)
        globalMap.Delete(key)
    })
}

func Get(key string) (interface{}, bool) {
    return globalMap.Load(key)
}

func Len() int {
    return int(len)
}
```


----

+ 一文了解Golang中的缓存库freecache：<https://www.php.cn/be/go/489034.html>
+ 从是否对GC影响角度来看缓存框架大致分为2类：
    - 零GC开销：比如freecache或bigcache这种，底层基于ringbuf，减小指针个数；
    - 有GC开销：直接基于Map来实现的缓存框架。
+ 对于map而言，gc时会扫描所有key/value键值对，如果其都是基本类型，那么gc便不会再扫描。    


+ https://github.com/coocood/freecache
+ https://github.com/allegro/bigcache
+ https://github.com/VictoriaMetrics/fastcache
+ https://github.com/bluele/gcache
+ https://github.com/patrickmn/go-cache

---

https://zhuanlan.zhihu.com/p/398693305

比较知名的包：

github.com/patrickmn/go-cache 但是2019年之后就不维护了

另外性能比较好的两个包:

github.com/allegro/bigcache github.com/coocood/freecache 但是只能缓存值为[]byte类型的对象，适用性比较窄。

---

## go-cache
+ Golang package轻量级KV数据缓存——go-cache源码分析 :<https://www.cnblogs.com/Moon-Light-Dream/p/12494683.html>


<pre>
go-cache的源码代码里很小，代码结构和处理逻辑都比较简单，可以作为golang新手阅读的很好的素材。
对于单机轻量级的内存缓存如果仅从功能实现角度考虑，go-cache是一个不错的选择，使用简单。
但在实际使用中需要注意：
go-cache没有对内存使用大小或存储数量进行限制，可能会造成内存峰值较高；
go-cache中存储的value尽量使用指针类型，相比于存储对象，不仅在性能上会提高，在内存占用上也会有优势。由于golang的gc机制，map在扩容后原来占用的内存不会立刻释放，因此如果value存储的是对象会造成占用大量内存无法释放。
</pre>

+ 从结构体上, go-cache 主要还是由 map + RWMutex 组成.

+ 一次错误使用 go-cache 导致出现的线上问题:<https://www.cnblogs.com/457220157-FTD/p/14707868.html>
<pre>
我使用的问题
背景: 某接口 QPS 有点高

当时考虑到用户购买状态(这个状态可能随时变化)如果能够在本地缓存中缓存 10s, 那么用户再次点进来的时候能从本地取了, 就造成大量的数据都写入了 map 中
由于接口 QPS 比较高, 设置用户购买状态时就可能造成竞争, 造成接口响应超时
go-cache 使用注意点
尽量存放那些相对不怎么变化的数据, 适用于所有的 local cache(包括 map, sync.map)
go-cache 的过期检查时间要设置相对较小, 也不能过小
那些高 QPS 的接口尽量不要去直接 Set 数据, 如果必须 Set 可以采用异步操作
监控 go-cache 里面 key 的数量, 如果过多时, 需要及时调整参数
</pre>

### 源码分析
+ mu                sync.RWMutex               // map本身非线程安全，操作时需要加锁
+ stop     chan bool     // 用来接收结束信息
```go
type janitor struct {
	Interval time.Duration // 定时轮询周期
	stop     chan bool     // 用来接收结束信息
}

func (j *janitor) Run(c *cache) {
	ticker := time.NewTicker(j.Interval) // 创建一个timeTicker定时触发
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()            // 调用DeleteExpired接口处理删除过期记录
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}
```
+ **对于janitor的处理，这里使用的技巧值得学习 **，下面这段代码是在New() cache对象时，会同时开启一个goroutine跑janitor，在run之后可以看到做了runtime.SetFinalizer的处理，这样处理了可能存在的内存泄漏问题。 !!!!!!
    - 可能的泄漏场景如下，使用者创建了一个cache对象，在使用后置为nil，在使用者看来在gc的时候会被回收，但是因为有goroutine在引用，在gc的时候不会被回收，因此导致了内存泄漏。
    - 解决方案可以增加Close接口，在使用后调用Close接口，通过channel传递信息结束goroutine，但如果使用者在使用后忘了调用Close接口，还是会造成内存泄漏。
    - 另外一种解决方法是使用runtime.SetFinalizer，不需要用户显式关闭， gc在检查C这个对象没有引用之后， gc会执行关联的SetFinalizer函数，主动终止goroutine，并取消对象C与SetFinalizer函数的关联关系。这样下次gc时，对象C没有任何引用，就可以被gc回收了。（回收C时执行一个回调方法，回调方法是自己的定义的stopJanitor，stopJanitor传true给stop这个channel，从而能停止goroutine）
+ 通过`runtime.SetFinalizer` 回调方法避免内存泄漏！
  - 用goroutine开启定时任务，会有这个问题，所以控制结束时机
  - java使用线程开启定时任务，同样会有这个问题！！
+ runtime.SetFinalizer 是Go提供对象被GC回收时的一个注册函数，可以在对象被回收的时候回调函数


----

## Freecache

+ 深入理解Freecache:<https://blog.csdn.net/chizhenlian/article/details/108435024>

使用freecache的注意事项
缓存的数据如果可以的话，大小尽量均匀一点，可以减少RingBuf容量不足时的置换工作开销。
缓存的数据不易过大，这样子才能缓存更多的key，提高缓存命中率。

+ 深入理解golang内存缓存利器-FreeCache:<https://zhuanlan.zhihu.com/p/402841754> --- 这篇写得更好!!!


## Bigcache

避免高额的GC开销
在bigCache中，map中没有使用指针，在 Golang(>1.4) 中，如果map中不包含指针的话，GC 便会忽略这个 map。

在bigCache中，bigCache将数据存储在BytesQueue中，BytesQueue的底层结构是[]byte ，这样只给GC增加了一个额外对象，

由于字节切片除了自身对象并不包含其他指针数据，所以GC对于整个对象的标记时间是O(1)的。
————————————————
版权声明：本文为CSDN博主「CoLiuRs」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/baidu_32452525/article/details/118199442

+ <https://zhuanlan.zhihu.com/p/285133613>

FreeCache
FreeCache 将缓存分成了 256 段，每段包括 256 个槽和一个 ring buffer 存储数据。当一个新的元素被添加进来的时候，使用 hash 值下 8 位作为标识 id，通过使用 LSB 9-16 的值作为槽 ID。将数据分配到多个槽里面，有助于优化查询的时间 ( 分治策略 )。

数据被存储在 ring buffer 中，位置被保存在一个排序的数组里面。如果 ring buffer 内存不足，则会利用 LRU 的策略在 ring buffer 逐个扫描，如果缓存的最后访问时间小于平均访问的时间，就会被删掉。要找到一个缓存内容，在槽中是通过二分查找法对一个已经排好的数据进行查询。

GroupCache
GroupCache 使用链表和 Map 实现了一个精准的 LRU 删除策略的缓存。为了进行公平的比较，我们在 GroupCache 的基础上，实现了一个包括 256 个分片的切片结构。

+ 读:由于读锁是无消耗的，所以 BigCache 的伸缩性更好。FreeCache 和 GroupCache 读锁是有消耗的，并且在并发数达到 20 的时候，伸缩性下降了。(Y 轴越大越好 )
+ 写:在只写的情况下，三者的性能表现比较接近，FreeCache 比另两个的情况，稍微好一点。
+ 读写情况 (25% 写，75% 读 ): 两者混合的情况下，BigCache 看起来是唯一一个在伸缩性上表现完美的

### 源码分析
认真学完 bigcache 的代码，我们至少有以下几点收获：

可以通过 sharding 来降低资源竞争
可以用位运算来取余数做 sharding （需要是 2 的整数幂 - 1）
避免 map 中出现指针、使用 go 基础类型可以显著降低 GC 压力、提升性能
bigcache 底层存储是 bytes queue，初始化时设置合理的配置项可以减少 queue 扩容的次数，提升性能

作者：翔叔架构笔记
链接：https://juejin.cn/post/7107635176263385118
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

### todo
+ Golang 高性能 LocalCache：BigCache 设计与分析:<https://pandaychen.github.io/2020/03/03/BIGCACHE-ANALYSIS/>
+ Go开源项目BigCache如何加速并发访问以及避免高额的GC开销:<https://juejin.cn/post/6844903993114624008>
+ 妙到颠毫: bigcache优化技巧:<https://colobu.com/2019/11/18/how-is-the-bigcache-is-fast/>


## ristretto

+ Ristretto简介：高性能Go缓存:<https://www.zhihuclub.com/87009.shtml>
+ Ristretto: 高性能内存绑定Go缓存:<https://www.5axxw.com/wiki/content/fzxtu6>

+ Introducing Ristretto: A High-Performance Go Cache:<https://phenix3443.github.io/notebook/golang/Introducing_Ristretto_A_High-Performance_Go_Cache.html>


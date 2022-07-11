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
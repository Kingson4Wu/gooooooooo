+ Golang sync.Pool 简介与用法:<https://cloud.tencent.com/developer/article/1478061>

+ Pool 是可伸缩、并发安全的临时对象池，用来存放已经分配但暂时不用的临时对象，通过对象重用机制，缓解 GC 压力，提高程序性能。

+ 注意，sync.Pool 是一个临时的对象池，适用于储存一些会在 goroutine 间共享的临时对象，其中保存的任何项都可能随时不做通知地释放掉，所以不适合用于存放诸如 socket 长连接或数据库连接的对象。

```go
package main

import (
    "bytes"
    "io"
    "os"
    "sync"
    "time"
)

var bufPool = sync.Pool {
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func PoolTest(w io.Writer, key, val string) {
    b, _ := bufPool.Get().(*bytes.Buffer)
    b.Reset()
    b.WriteString(time.Now().UTC().Format("2006-01-02 15:04:05"))
    b.WriteByte('|')
    b.WriteString(key)
    b.WriteByte('=')
    b.WriteString(val)
    w.Write(b.Bytes())
	w.Write([]byte("\n"))
    bufPool.Put(b)
}

func main() {
    PoolTest(os.Stdout, "dablelv", "monkey")
}
```


---

+ 深入Golang之sync.Pool详解: <https://www.cnblogs.com/sunsky303/p/9706210.html>

+ Get方法并不会对获取到的对象值做任何的保证，因为放入本地池中的值有可能会在任何时候被删除，但是不通知调用者。放入共享池中的值有可能被其他的goroutine偷走。 所以对象池比较适合用来存储一些临时切状态无关的数据，但是不适合用来存储数据库连接的实例，因为存入对象池重的值有可能会在垃圾回收时被删除掉，这违反了数据库连接池建立的初衷。

根据上面的说法，Golang的对象池严格意义上来说是一个临时的对象池，适用于储存一些会在goroutine间分享的临时对象。主要作用是减少GC，提高性能。在Golang中最常见的使用场景是fmt包中的输出缓冲区。

在Golang中如果要实现连接池的效果，可以用container/list来实现，开源界也有一些现成的实现，比如go-commons-pool，具体的读者可以去自行了解。

---

+ 深度分析 Golang sync.Pool 底层原理:<https://www.cyhone.com/articles/think-in-sync-pool/>
+ sync包中主要有：Locker, Cond, Map, Mutex, Once, Pool,、RWMutex, WaitGroup



## sync-cond

+ https://geektutu.com/post/hpg-sync-cond.html

```go

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}
```

<pre>
done 即互斥锁需要保护的条件变量。
read() 调用 Wait() 等待通知，直到 done 为 true。
write() 接收数据，接收完成后，将 done 置为 true，调用 Broadcast() 通知所有等待的协程。
write() 中的暂停了 1s，一方面是模拟耗时，另一方面是确保前面的 3 个 read 协程都执行到 Wait()，处于等待状态。main 函数最后暂停了 3s，确保所有操作执行完毕。
</pre>
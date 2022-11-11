+ 当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞


## TODO
+ 通俗易懂！图解Go协程原理及实战:<https://mp.weixin.qq.com/s/POZGQXsYu5aWXvu29vSM8g>
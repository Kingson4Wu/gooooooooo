### runtime.SetFinalizer
+ runtime.SetFinalizer 是Go提供对象被GC回收时的一个注册函数，可以在对象被回收的时候回调函数
+ Finalizer 函数返回结果是系统会忽略，所以你返回error也无所谓，但是切记不可以panic，程序是无法recover的

+ 指针构成的 "循环引⽤" 加上 runtime.SetFinalizer 会导致内存泄露。
https://cloud.tencent.com/developer/article/1410230

+ https://www.cnblogs.com/binHome/p/12901392.html

上面是官方文档对SetFinalizer的一些解释，主要含义是对象可以关联一个SetFinalizer函数， 当gc检测到unreachable对象有关联的SetFinalizer函数时，会执行关联的SetFinalizer函数， 同时取消关联。 这样当下一次gc的时候，对象重新处于unreachable状态并且没有SetFinalizer关联， 就会被回收。

<pre>
仔细看文档，还有几个需要注意的点：

* 即使程序正常结束或者发生错误， 但是在对象被 gc 选中并被回收之前，SetFinalizer 都不会执行， 所以不要在SetFinalizer中执行将内存中的内容flush到磁盘这种操作

* SetFinalizer 最大的问题是延长了对象生命周期。在第一次回收时执行 Finalizer 函数，且目标对象重新变成可达状态，直到第二次才真正 “销毁”。这对于有大量对象分配的高并发算法，可能会造成很大麻烦

* 指针构成的 "循环引⽤" 加上 runtime.SetFinalizer 会导致内存泄露
</pre>

+ 一个程序会有单独一个go程顺序执行所有的终止器。如果一个终止器必须运行较长时间，它应该在内部另开go程执行该任务。


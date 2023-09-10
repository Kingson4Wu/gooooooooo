+ Go 每日一库之 zap 高性能设计与实现:<https://mp.weixin.qq.com/s/6dLtjHbtDRekVHC8G_Rv9w>

<pre>
zap 高性能实现细节
通过内建的数据类型 zapcore.Field 和内建的日志编码器 (Encoder 接口)，避免标准库的序列化方法使用反射带来的性能损耗
通过内建的数据类型 zapcore.Field, 避免使用 interface{} 带来的开销 (拆装箱、对象逃逸到堆上)
通过内建的 []byte 缓冲池配合 zapcore.Field 进一步提升日志数据的写入性能
获取调用堆栈方法优化 (使用 runtime.Callers 而非 runtime.Stack)
写时复制机制 (多个日志共享一个 Logger 对象，在属性变更时复制一个新的对象，详情见 Logger.clone 方法及其调用方)
按需分配机制 (Check 方法检查可写后，再通过 CheckedEntry.Write 方法写入日志数据，可以保证 zapcore.Field 日志对象内存按需分配)
对象复用避免 GC (位于 hot path 上面的对象全部使用对象池管理模式进行复用)
避免数据竞态，虽然有对象池管理复用，但是对象的获取都需要经过各种条件过滤，有效缓解了底层 sync.Pool 内部的数据竞态问题
重复检测，每个日志保证只写入一次，提升性能并且避免应用层的错误使用导致的 Bug
</pre>
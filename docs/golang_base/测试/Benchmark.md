+ 基准测试的函数名须以 Benchmark 开头， 参数须为 *testing.B；循环中的 b.N， go 会根据系统情况生成，不用用户设定。

+ go test -bench=. -run=none


+ allocs/op 表示每个操作(单次迭代)发生了多少不同的内存分配。
B/op 是每个操作分配的字节数。
ns/op表示每次需要花费纳秒。
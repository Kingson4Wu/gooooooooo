+ 基准测试的函数名须以 Benchmark 开头， 参数须为 *testing.B；循环中的 b.N， go 会根据系统情况生成，不用用户设定。

+ go test -bench=. -run=none
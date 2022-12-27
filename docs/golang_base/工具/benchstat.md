+ go install golang.org/x/perf/cmd/benchstat@latest
+ https://segmentfault.com/a/1190000016354758

对比标准 benchmarks 和 benchstat
确定两组基准测试结果之间的差异可能是单调乏味且容易出错的。  Benchstat 可以帮助我们解决这个问题。

//Benchstat 可以获取一组基准测试数据，并告诉你它的稳定性如何。
//go test -bench=Fib20 -count=10 | tee old.txt
/**
goos: darwin
goarch: amd64
pkg: git.kugou.net/cupid/gooooooooo/command/cmd/hello/benchstat
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
BenchmarkFib20-4           24994             47238 ns/op
BenchmarkFib20-4           25598             46672 ns/op
BenchmarkFib20-4           25603             48440 ns/op
BenchmarkFib20-4           22058             78869 ns/op
BenchmarkFib20-4           17175             75345 ns/op
BenchmarkFib20-4           21458             50045 ns/op
BenchmarkFib20-4           24278             48631 ns/op
BenchmarkFib20-4           25544             46383 ns/op
BenchmarkFib20-4           23900             66975 ns/op
BenchmarkFib20-4           22447             48052 ns/op
PASS
ok      git.kugou.net/cupid/gooooooooo/command/cmd/hello/benchstat      19.332s

benchstat old.txt
name     time/op
Fib20-4  55.7µs ±42%


*/

可以保存生成它的二进制文件。 为此，请使用-c标志来保存测试二进制文件

---


分析两组样本
同上，把性能优化后的结果输出到名为BenchmarkReadGoSum.after的文件，然后使用benchstat分析优化的效果：

$ benchstat BenchmarkReadGoSum.before BenchmarkReadGoSum.after 
name         old time/op  new time/op  delta
ReadGoSum-4   531µs ± 3%   518µs ± 7%  -2.41%  (p=0.033 n=13+15)
当只有两组样本时，benchstat还会额外计算出差值，比如本例中，平均花费时间下降了2.41%。

另外，p=0.033表示结果的可信程度，p 值越大可信程度越低，统计学中通常把p=0.05作为临界值，超过此值说明结果不可信，可能是样本过少等原因。

https://www.cnblogs.com/failymao/p/15033094.html


---


benchstat 命令
benchstat computes and compares statistics about benchmarks.

benchstat 计算并比较关于基准测试的统计数据。

benchstat 这个工具可以将多次测试的结果汇总，生成概要信息。


https://maiyang.me/post/2018-04-20-benchstat-guide/

---

https://chenlujjj.github.io/go/benchstat/

benchstat -h

当输入是一个文件时：打印出该文件中的 benchmark 统计结果
当输入是一对文件时：分别打印出这两个文件的 benchmark 统计结果，以及它们的比较信息
当输入是三个及以上数量的文件时：分别打印出这多个文件的 benchmark 统计结果

go test -run=NONE -bench=Fib40 -count=5 | tee -a old.txt
go test -run=NONE -bench=Fib40 -count=5 | tee -a new.txt
查看单次 benchmark 统计结果
benchstat old.txt
比较两次的统计结果
benchstat old.txt new.txt


+ 深入Go代码覆盖率使用、场景与原理:<https://mp.weixin.qq.com/s/xAPEI7q5FaztxoZNOH-afA>

+ https://mp.weixin.qq.com/s/yycu10nLvpC0XiRemSy3lA

目前xUnit Patterns[3]至少定义了五种类型的Test Doubles：

Test stubs
Mock objects
Test spies
Fake objects
Dummy objects
这其中最为常用的是Fake objects、stub和mock objects。

2.1 fake object
Go标准库中还有net/dnsclient_unix_test.go中的fakeDNSServer等。此外，Go标准库中一些以mock做前缀命名的变量、类型等其实质上是fake object

2.2 stub
stub显然也是一个在测试阶段专用的、用来替代真实外部协作者与SUT进行交互的对象。与fake object稍有不同的是，stub是一个内置了预期值/响应值且可以在多个测试间复用的替身object。

stub可以理解为一种fake object的特例。

fake object与stub的优缺点基本一样。多数情况下，大家也无需将这二者划分的很清晰。

2.3 mock object
和fake/stub一样，mock object也是一个测试替身。通过上面的例子我们看到fake建立困难(比如创建一个近2千行代码的fakeDriver)，但使用简单。而mock object则是一种建立简单，使用简单程度因被测目标与外部协作者交互复杂程度而异的test double


----

依赖Kafka的Go单元测试例解 : <https://mp.weixin.qq.com/s/M_lV7FaIMxZGd6pZ9dDaAw>
    - testcontainers-go是一个Go语言开源项目，专门用于简化创建和清理基于容器的依赖项，常用于Go项目的单元测试、自动化集成或冒烟测试中。
    

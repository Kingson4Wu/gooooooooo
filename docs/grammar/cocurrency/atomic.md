+ Go标准库学习笔记-原子操作 (sync/atomic):<https://blog.csdn.net/preyta/article/details/80192026>

+ AddInt32
AddInt32可以实现对元素的原子增加或减少，其函数定义如下，函数接收两个参数，分别是需要修改的变量的地址和修改的差值，函数会直接在传递的地址上进行修改操作，此外函数会返回修改之后的新值

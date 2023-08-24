// 导出原子加法函数给C代码调用
.global AtomicAdd
.type AtomicAdd, %function
AtomicAdd:
    ldxr    x1, [x0]
    add     x1, x1, x2
retry:
    stxr    w3, x1, [x0]
    cbnz    w3, retry
    ret

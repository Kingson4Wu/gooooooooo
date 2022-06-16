+ func(interface{}) 使用这个作为形参，那么func参数类型就不会被限制，使用时转化到相应类型即可，比如：`n := i.(int32)`
+ 如果类型不匹配，会产生panic错误，如果想不产生panic错误

str, ok := param.(string)
如果ok为false，则str为空，不报错


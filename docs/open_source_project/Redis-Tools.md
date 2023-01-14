
+ https://github.com/zehuamama/redis-tools

Redis-Toolsredis-tools聚合了后端小伙伴常见的工具集，它有以下特性：
分布式锁、自旋锁原子命令 
compare and swap原子命令 
compare and delete
redis-tools 代码量极少，学习它，开发者可以得到以下收获：
代码简洁规范
redis lua脚本实现
redis 分布式锁的正确实现

作者：马丸子
链接：https://www.zhihu.com/question/20801814/answer/2461184379
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


+ load script  复用lua脚本

```go
/*scr := rdb.ScriptLoad(ctx, compareAndDeleteScript)
	hash, err := scr.Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash)*/
	hash := "867ddd622a45c70d280b9f3bcaceb191fdd7e452"
	//867ddd622a45c70d280b9f3bcaceb191fdd7e452
	fmt.Println("-------")
	rdb.Set(ctx, "kkkggg", "5553333", 0)
	res, err = rdb.EvalSha(ctx, hash, []string{"kkkggg"}, "5553333").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
```

+ 借鉴Redisson 实现Go的分布式锁: <https://juejin.cn/post/6867004134894403592>
+ 在Go中如何用Redlock实现分布式锁: <https://www.aisoutu.com/a/435487>
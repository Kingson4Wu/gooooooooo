package dp

func climbStairs(n int) int {

	dp := make([]int, n)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n-1]
}

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

/**
自己递归做完了。。。

递归超出时间限制。。。。


执行结果：
超出时间限制

func climbStairs(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n-2) + climbStairs(n-1)
}

自己改成dp，完成

执行结果：
通过
显示详情
添加备注

执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
1.8 MB
, 在所有 Go 提交中击败了
35.46%
的用户
通过测试用例：
45 / 45

*/

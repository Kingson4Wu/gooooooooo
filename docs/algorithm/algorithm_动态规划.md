+ 凑零钱问题
+ 带“备忘录”的递归
+ 自顶向下、自底向上

+ 状态转移方程

+ 动态规划的特性，穷举加备忘录/DP table 优化
+ 只要遇到求最值的算法问题，首先要思考的就是：如何穷举出所有可能的结果？
+ 穷举主要有两种算法，就是回溯算法和动态规划，前者就是暴力穷举，而后者是根据状态转移方程推导「状态」


+ 第一步要明确两点，「状态」和「选择」。
+ 第二步要明确 dp 数组的定义。

+ 二维动态规划压缩成一维动态规划吗？这就是状态压缩


---

+ leetcode/dp/70. 爬楼梯.go
 climbStairs(n) = climbStairs(n-2) + climbStairs(n-1)

+ leetcode/dp/119. 杨辉三角 II.go
算好规律

+ leetcode/dp/5. 最长回文子串.go  !!!!


+ leetcode/dp/64. 最小路径和.go
dp[i,j]=min(dp[i-1,j],dp[i,j-1]) + val(i,j)


+ leetcode/dp/300. 最长递增子序列.go !!!!

+ leetcode/dp/55. 跳跃游戏.go !!!
贪心算法！！ 遍历更新最大可达位置！！！


+ leetcode/dp/45. 跳跃游戏 II.go
遍历计算到达每一步的最小步数！
需考虑数组长度为1 的情况 ！！！！

+ leetcode/dp/139. 单词拆分.go !!!!

```go
func wordBreak(s string, wordDict []string) bool {
    wordDictSet := make(map[string]bool)
    for _, w := range wordDict {
        wordDictSet[w] = true
    }
    dp := make([]bool, len(s) + 1)
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordDictSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}
```

+ leetcode/dp/221. 最大正方形.go ！！！



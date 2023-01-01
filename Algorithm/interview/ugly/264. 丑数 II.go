package ugly

/**
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。



示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。


提示：

1 <= n <= 1690
*/

func nthUglyNumber(n int) int {

	return 1
}

/**
方法一：最小堆
要得到从小到大的第 nnn 个丑数，可以使用最小堆实现。

初始时堆为空。首先将最小的丑数 111 加入堆。

每次取出堆顶元素 xxx，则 xxx 是堆中最小的丑数，由于 2x,3x,5x2x, 3x, 5x2x,3x,5x 也是丑数，因此将 2x,3x,5x2x, 3x, 5x2x,3x,5x 加入堆。

上述做法会导致堆中出现重复元素的情况。为了避免重复元素，可以使用哈希集合去重，避免相同元素多次加入堆。

在排除重复元素的情况下，第 nnn 次从最小堆中取出的元素即为第 nnn 个丑数。

var factors = []int{2, 3, 5}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func nthUglyNumber(n int) int {
    h := &hp{sort.IntSlice{1}}
    seen := map[int]struct{}{1: {}}
    for i := 1; ; i++ {
        x := heap.Pop(h).(int)
        if i == n {
            return x
        }
        for _, f := range factors {
            next := x * f
            if _, has := seen[next]; !has {
                heap.Push(h, next)
                seen[next] = struct{}{}
            }
        }
    }
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/ugly-number-ii/solutions/712102/chou-shu-ii-by-leetcode-solution-uoqd/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


方法二：动态规划
方法一使用最小堆，会预先存储较多的丑数，维护最小堆的过程也导致时间复杂度较高。可以使用动态规划的方法进行优化。

定义数组 dp，其中 dp[i]表示第 i个丑数，第 n个丑数即为 dp[n]。

由于最小的丑数是 111，因此 dp[1]=1。

如何得到其余的丑数呢？定义三个指针 p2,p3,p5，表示下一个丑数是当前指针指向的丑数乘以对应的质因数。初始时，三个指针的值都是 1。

当 2≤i≤n 时，令 dp[i]=min⁡(dp[p2]×2,dp[p3]×3,dp[p5]×5)，然后分别比较 dp[i]和 dp[p2]×2,dp[p3]×3,dp[p5]×5 是否相等，如果相等则将对应的指针加 1。


直接看代码,更容易理解!!..

func nthUglyNumber(n int) int {
    dp := make([]int, n+1)
    dp[1] = 1
    p2, p3, p5 := 1, 1, 1
    for i := 2; i <= n; i++ {
        x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
        dp[i] = min(min(x2, x3), x5)
        if dp[i] == x2 {
            p2++
        }
        if dp[i] == x3 {
            p3++
        }
        if dp[i] == x5 {
            p5++
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

复杂度分析

时间复杂度：O(n)。需要计算数组 dp 中的 n 个元素，每个元素的计算都可以在 O(1)的时间内完成。

空间复杂度：O(n)。空间复杂度主要取决于数组 dp\textit{dp}dp 的大小。

作者：力扣官方题解
链接：https://leetcode.cn/problems/ugly-number-ii/solutions/712102/chou-shu-ii-by-leetcode-solution-uoqd/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

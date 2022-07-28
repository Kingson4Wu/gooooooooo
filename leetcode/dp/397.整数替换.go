package dp

/**
不会做！

1、贪心算法，总结规律

*/
/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换

 给定一个正整数 n ，你可以做如下操作：

如果 n 是偶数，则用 n / 2替换 n 。
如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
返回 n 变为 1 所需的 最小替换次数 。



示例 1：

输入：n = 8
输出：3
解释：8 -> 4 -> 2 -> 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

方法一：枚举所有的情况
思路与算法

我们可以使用递归的方法枚举所有将 nn 变为 11 的替换序列：

当 nn 为偶数时，我们只有唯一的方法将 nn 替换为 \dfrac{n}{2}
2
n
 。

当 nn 为奇数时，我们可以选择将 nn 增加 11 或减少 11。由于这两种方法都会将 nn 变为偶数，那么下一步一定是除以 22，因此这里我们可以看成使用两次操作，将 nn 变为 \dfrac{n+1}{2}
2
n+1
​
  或 \dfrac{n-1}{2}
2
n−1


方法二：记忆化搜索
思路与算法

我们给方法一的递归加上记忆化，这样递归树的每一层最多只会计算两个 nn 值，时间复杂度降低为 O(\log n)O(logn)。


方法三：贪心
思路与算法

实际上，方法一和方法二中的递归枚举中的「最优解」是固定的：

当 nn 为偶数时，我们只有唯一的方法将 nn 替换为 \dfrac{n}{2}
2
n
​
 ；

当 nn 为奇数时，nn 除以 44 的余数要么为 11，要么为 33。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/integer-replacement/solution/zheng-shu-ti-huan-by-leetcode-solution-swef/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// @lc code=start
func integerReplacement(n int) int {
	ans := 0
	for n != 1 {
		switch {
		case n%2 == 0:
			ans++
			n /= 2
		case n%4 == 1:
			ans += 2
			n /= 2
		case n == 3:
			ans += 2
			n = 1
		default:
			ans += 2
			n = n/2 + 1
		}
	}
	return ans
}

// @lc code=end

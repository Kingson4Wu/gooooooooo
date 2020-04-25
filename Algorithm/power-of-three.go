package algorithm

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

/**
使用 32 位来表示数字，所以范围的一半用于负数，0 是正数的一部分。
知道了 n 的限制，我们现在可以推断出 n 的最大值，也就是 3 的幂，是 1162261467

作者：LeetCode
链接：https://leetcode-cn.com/problems/power-of-three/solution/3de-mi-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

执行用时 :
48 ms
, 在所有 Go 提交中击败了
17.10%
的用户
内存消耗 :
6.1 MB
, 在所有 Go 提交中击败了
80.00%
的用户
性能不行
*/

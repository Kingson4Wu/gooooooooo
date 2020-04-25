package algorithm

/**
给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
https://leetcode-cn.com/problems/power-of-four/
*/
func isPowerOfFour(num int) bool {

	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
			if count > 1 {
				return false
			}
		}
		if num>>1&1 == 1 {
			return false
		}

		num = num >> 2
	}
	return count == 1

}

/**
自己做的!!
1. 与特殊的值与，特别是题目上有提示多少位的整数
2. 善用数学公式推导
*/
/**

方法三：位操作
算法：

我们首先检查 num 是否为 2 的幂：x > 0 and x & (x - 1) == 0。

现在的问题是区分 2 的偶数幂（当 xx 是 4 的幂时）和 2 的奇数幂（当 xx 不是 4 的幂时）。在二进制表示中，这两种情况都只有一位为 1，其余为 0。

因此 4 的幂与数字 (101010...10)相与会得到 0。
​
  用十六进制表示为 ：(aaaaaaaa)_{16}(aaaaaaaa)

class Solution {
  public boolean isPowerOfFour(int num) {
    return (num > 0) && ((num & (num - 1)) == 0) && ((num & 0xaaaaaaaa) == 0);
  }
}

方法四：位运算 + 数学运算
算法：

我们首先检车 xx 是否为 2 的幂：x > 0 and x & (x - 1) == 0。然后可以确定 x = 2^ax=2
a
 ，若 xx 为 4 的幂则 aa 为偶数。
下一步是考虑 a=2ka=2k 和 a=2k+1a=2k+1 两种情况，对 xx 对 3 进行取模：
(2^{2k} \mod 3) = (4^k \mod 3) = ((3 + 1)^k \mod 3) = 1(2
2k
 mod3)=(4
k
 mod3)=((3+1)
k
 mod3)=1

((2^{2k + 1}) \mod 3) = ((2 \times 4^k) \mod 3) = ((2 \times(3 + 1)^k) \mod 3) = 2((2
2k+1
 )mod3)=((2×4
k
 )mod3)=((2×(3+1)
k
 )mod3)=2

若 xx 为 2 的幂且 x%3 == 1，则 xx 为 4 的幂。

class Solution {
  public boolean isPowerOfFour(int num) {
    return (num > 0) && ((num & (num - 1)) == 0) && (num % 3 == 1);
  }
}

作者：LeetCode
链接：https://leetcode-cn.com/problems/power-of-four/solution/4de-mi-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

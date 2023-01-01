package ugly

/**
丑数 就是只包含质因数 2、3 和 5 的正整数。

给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。



示例 1：

输入：n = 6
输出：true
解释：6 = 2 × 3
示例 2：

输入：n = 1
输出：true
解释：1 没有质因数，因此它的全部质因数是 {2, 3, 5} 的空集。习惯上将其视作第一个丑数。
示例 3：

输入：n = 14
输出：false
解释：14 不是丑数，因为它包含了另外一个质因数 7 。


提示：

-231 <= n <= 231 - 1
*/

func isUgly(n int) bool {

	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	var m int
	m = n / 2
	if m*2 == n {
		return isUgly(m)
	}
	m = n / 3
	if m*3 == n {
		return isUgly(m)
	}
	m = n / 5
	if m*5 == n {
		return isUgly(m)
	}

	return false
}

// 自己做的,第一次提交错误,没判断负数,负数不算丑数
/**
时间
0 ms
击败
100%
内存
1.9 MB
击败
99.59%
*/

/**

 官方的简洁,我的写法太蠢...

var factors = []int{2, 3, 5}

func isUgly(n int) bool {
    if n <= 0 {
        return false
    }
    for _, f := range factors {
        for n%f == 0 {
            n /= f
        }
    }
    return n == 1
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/ugly-number/solutions/712106/chou-shu-by-leetcode-solution-fazd/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

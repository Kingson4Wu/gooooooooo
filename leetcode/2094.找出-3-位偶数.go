package leetcode

/*
 * @lc app=leetcode.cn id=2094 lang=golang
 *
 * [2094] 找出 3 位偶数

 不会。。。


 ----

 方法一：枚举数组中的元素组合
思路与算法

我们可以从数组中枚举目标整数的三个整数位，判断组成的整数是否满足以下条件：

整数为偶数；

整数不包含前导零（即整数不小于 100100）；

三个整数位对应的数组下标不能重复。

为了避免重复，我们用一个哈希集合来维护符合要求的 33 位偶数，如果枚举产生的整数满足上述三个条件，则我们将该整数加入哈希集合。

最终，我们将该哈希集合内的元素放入数组中，按照递增顺序排序并返回。

方法二：遍历所有可能的 33 位偶数
思路与算法

我们也可以从小到大遍历所有 33 位偶数（即 [100, 999][100,999] 闭区间内的所有偶数），并判断对应的三个整数位是否为 \textit{digits}digits 数组中三个不同元素。如果是，则该偶数为目标偶数；反之亦然。

具体地，我们首先用哈希表 \textit{freq}freq 维护 \textit{digits}digits 数组中每个数出现的次数。在遍历偶数时，我们同样用哈希表 \textit{freq}_1freq
1
​
  维护每个偶数中每个数位出现的次数。此时，该偶数能够被数组中不重复元素表示的充要条件即为：

\textit{freq}_1freq
1
​
  中每个元素的出现次数都不大于它在 \textit{freq}freq 中的出现次数。

我们按照上述条件判断每个偶数是否为目标偶数，并按顺序统计这些偶数。最终，我们返回目标偶数的数组作为答案。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/finding-3-digit-even-numbers/solution/zhao-chu-3-wei-ou-shu-by-leetcode-soluti-hptf/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/**
digits := []int{2, 1, 3, 0}
	leetcode.Exexute(digits)
*/
/**
func Exexute(digits []int) []int {
	return findEvenNumbers(digits)
}*/

// @lc code=start
func findEvenNumbers(digits []int) []int {

	digitsArr := [10]int{}
	for _, digit := range digits {
		digitsArr[digit]++
	}

	result := []int{}
	for i := 100; i <= 998; i++ {
		if i%2 != 0 {
			continue
		}
		hitMap := make(map[int]int)
		ss := i
		for {
			s := ss % 10
			ss = ss / 10
			hitMap[s]++
			if ss == 0 {
				break
			}
		}
		hit := true
		for k, v := range hitMap {
			if v > digitsArr[k] {
				hit = false
				break
			}
		}
		if hit {
			result = append(result, i)
		}

	}
	return result
}

// @lc code=end

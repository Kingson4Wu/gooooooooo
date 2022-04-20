package leetcode

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III

这里使用的是
方法一：使用额外空间
复杂度分析

时间复杂度：O(N)，其中 NN 为字符串的长度。原字符串中的每个字符都会在 O(1) 的时间内放入新字符串中。

空间复杂度：O(N)。我们开辟了与原字符串等大的空间。

---
方法二：原地解法
思路与算法

此题也可以直接在原字符串上进行操作，避免额外的空间开销。当找到一个单词的时候，我们交换字符串第一个字符与倒数第一个字符，随后交换第二个字符与倒数第二个字符……如此反复，就可以在原空间上翻转单词。

需要注意的是，原地解法在某些语言（比如 Java，JavaScript）中不适用，因为在这些语言中 String 类型是一个不可变的类型。

时间复杂度：O(N)。字符串中的每个字符要么在 O(1) 的时间内被交换到相应的位置，要么因为是空格而保持不动。

空间复杂度：O(1)。因为不需要开辟额外的数组。


作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/solution/fan-zhuan-zi-fu-chuan-zhong-de-dan-ci-iii-by-lee-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

---
总结：其实可以不用额外开辟空间！！
被平常的开发习惯所局限了！平常开发业务需求，这样是很骚的！但是做算法题，这样是很巧的！

*/

// @lc code=start
func reverseWords(s string) string {
	len := len(s)
	ss := []rune(s)
	sn := make([]rune, len)
	start := 0
	end := 0

	for i, c := range ss {
		if c == ' ' {
			end = i - 1
			for j := start; j <= end; j++ {
				sn[j] = ss[(end)-(j-start)]
			}
			sn[i] = ' '

			start = i + 1
		}

		if i == (len - 1) {
			end = i
			for j := start; j <= end; j++ {
				sn[j] = ss[(end)-(j-start)]
			}
		}

	}
	return string(sn)
}

// @lc code=end

package algorithm

/**
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func reverseStr(s string, k int) string {

	kk := make([]rune, k)
	r := make([]rune, len(s))

	l := 2 * k
	index := 0
	bb := false
	ii := 0
	for i, c := range s {
		index = i % l
		if index < k {
			bb = true
			kk[index] = c
			ii = i
		} else {
			bb = false
			if index == k {
				for j := 0; j < k; j++ {
					r[i-k+j] = kk[k-1-j]
				}
			}
			r[i] = c
		}
	}

	if bb {
		for i := ii - ii%k; i < len(s); i++ {
			r[i] = kk[ii%k-i%k]
		}
	}

	return string(r)

}

/**
想复杂了。。。
1.可以通过下标直接算，奇数k需要寻找反转的下标，偶数不用

func reverseStr(s string, k int) string {

	r := make([]rune, len(s))
	ss := []rune(s)

	remain := len(s) % k
	lastIndex := len(s)
	if remain > 0 {
		lastIndex = len(s) - remain
	}

	for i := 0; i < len(s); i++ {
		if i%(2*k) < k {

			l := k
			if i >= lastIndex {
				l = remain
			}
			index := i % k
			r[i] = ss[i-index+l-index-1]
		} else {
			r[i] = ss[i]
		}
	}

	return string(r)
}
*/

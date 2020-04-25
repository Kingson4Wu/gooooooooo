package algorithm

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
https://leetcode-cn.com/problems/power-of-two/
*/
func isPowerOfTwo(n int) bool {

	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
			if count > 1 {
				return false
			}

		}
		n = n >> 1
	}
	return count == 1

}

/**
自己做的
*/

package algorithm

func lengthOfLongestSubstring(s string) int {

	m := make(map[byte]int)

	len := len(s)

	start, end, maxLen := 0, 0, 0

	for i := 0; i < len; i++ {

		//存在key
		if _, ok := m[s[i]]; ok {
			maxLen = max(maxLen, end-start)
			delete(m, s[i])
		}

	}

	return 0

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

package algorithm

/**
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestPalindrome(s string) int {
	var arr [58]int
	for _, c := range s {
		arr[c-65]++
	}

	count := 0
	b := false

	for _, c := range arr {
		if c%2 == 0 {
			count += c
		} else {
			count += c
			count--
			b = true
		}
	}

	if b {
		count++
	}
	return count
}

/**
do myself

1. 字母ascii码
2. 数组代替hash

other:
class Solution {
    public int longestPalindrome(String s) {
        int[] cnts = new int[58];
         for (int i = 0; i < s.length(); i++) {
             cnts[s.charAt(i) - 'A']++;
         }
         int palindrome = 0;
         for (int cnt : cnts) {
             palindrome += (cnt / 2) * 2;
         }
         if (palindrome < s.length()) {
             palindrome++;
         }
         return palindrome;
     }
}
*/

package leetcode

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合

 全局变量！！！+ 递归！！！


 方法一：回溯
首先使用哈希表存储每个数字对应的所有可能的字母，然后进行回溯操作。

回溯过程中维护一个字符串，表示已有的字母排列（如果未遍历完电话号码的所有数字，则已有的字母排列是不完整的）。该字符串初始为空。每次取电话号码的一位数字，从哈希表中获得该数字对应的所有可能的字母，并将其中的一个字母插入到已有的字母排列后面，然后继续处理电话号码的后一位数字，直到处理完电话号码中的所有数字，即得到一个完整的字母排列。然后进行回退操作，遍历其余的字母排列。

回溯算法用于寻找所有的可行解，如果发现一个解不可行，则会舍弃不可行的解。在这道题中，由于每个数字对应的每个字母都可能进入字母组合，因此不存在不可行的解，直接穷举所有的解即可。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/solution/dian-hua-hao-ma-de-zi-mu-zu-he-by-leetcode-solutio/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

// @lc code=start
/* func letterCombinations(digits string) []string {

	result := []string{}

	if len(digits) == 0 {
		return result
	}

	backtrack(digits, 0, "")

	index := 0
	count := getCount(digit)

	for i := 0; i < count; i++ {
		startIndex := getStartIndex(digit)
		s:=[len(digits)]rune
		s[index]=toChar(startIndex+i)
		index++
		for index<len(digits){

		}

	}

	for index < len(digits) {
		digit := digits[index]
		count := getCount(digit)
		startIndex := getStartIndex(digit)
		for i := 0; i < count; i++ {

		}

	}

	for _, digit := range digits {

		result = append(result)
	}

}

func toChar(i int) rune {
	return rune('a' - 1 + i)
}

func getCount(i int) int {
	if i < 2 {
		return 0
	}
	if i == 7 || i == 9 {
		return 4
	}
	return 3
}

func getStartIndex(i int) int {
	if i == 7 {
		return 6*3 + 1
	}
	if i == 8 {
		return 6*3 + 4 + 1
	}
	if i == 9 {
		return 6*3 + 4 + 3 + 1
	}
	return (i-1)*3 + 1
}*/

// @lc code=end

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

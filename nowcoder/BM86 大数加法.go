package nowcoder

import "strconv"

func Solve1(s string, t string) string {
	// write code here

	sTotal := len(s)
	tTotal := len(t)

	if sTotal == 0 {
		return t
	}

	if tTotal == 0 {
		return s
	}

	maxTotal := 0
	if sTotal > tTotal {
		maxTotal = sTotal
	} else {
		maxTotal = tTotal
	}
	maxTotal++

	result := make([]rune, maxTotal)

	index := 0
	up := 0

	for {

		ss := 0
		sEnd := false
		if !sEnd && index < sTotal {
			ss, _ = strconv.Atoi(s[sTotal-index-1 : sTotal-index])
		} else {
			sEnd = true
		}
		tt := 0
		tEnd := false
		if !tEnd && index < tTotal {
			tt, _ = strconv.Atoi(t[tTotal-index-1 : tTotal-index])
		} else {
			tEnd = true
		}

		if sEnd && tEnd {
			break
		}

		sum := ss + tt + up
		up = sum / 10
		result[maxTotal-index-1] = []rune(strconv.Itoa(sum % 10))[0]

		index++
	}

	if up > 0 {
		result[0] = []rune(strconv.Itoa(up))[0]
		return string(result)

	}

	return string(result[1:maxTotal])
}

/**

自己做的，感觉写得有点复杂

运行时间：11ms
超过49.22% 用Go提交的代码
占用内存：2056KB
超过69.29%用Go提交的代码

描述
以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。

数据范围：s.length,t.length \le 100000s.length,t.length≤100000，字符串仅由'0'~‘9’构成
要求：时间复杂度 O(n)
示例1
输入：
"1","99"
复制
返回值：
"100"
复制
说明：
1+99=100
示例2
输入：
"114514",""
复制
返回值：
"114514"

*/

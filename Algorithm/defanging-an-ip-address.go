package algorithm

/**
给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。

所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。



示例 1：

输入：address = "1.1.1.1"
输出："1[.]1[.]1[.]1"
示例 2：

输入：address = "255.100.50.0"
输出："255[.]100[.]50[.]0"


提示：

给出的 address 是一个有效的 IPv4 地址

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/defanging-an-ip-address
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func defangIPaddr(address string) string {

	length := len(address) + 6
	result := make([]rune, length, length)
	index := 0
	for _, c := range address {
		if c == '.' {
			result[index] = '['
			index++
			result[index] = c
			index++
			result[index] = ']'
			index++
		} else {
			result[index] = c
			index++
		}
	}
	return string(result)

}

/**
复习关键点：
1. 遍历字符串是[]rune
2. string([]rune) 转字符串
*/

package nowcoder

func twoSum(numbers []int, target int) []int {
	// write code here

	context := make(map[int]int)

	for i, num := range numbers {

		remain := target - num

		if value, ok := context[remain]; ok {
			return []int{value + 1, i + 1}
		} else {
			if _, ok := context[num]; ok {
				continue
			} else {
				context[num] = i
			}
		}
	}

	return []int{}
}

/**
自己做的，hash表的使用
运行时间：60ms
超过14.94% 用Go提交的代码
占用内存：12252KB
超过23.69%用Go提交的代码
*/

/**
给出一个整型数组 numbers 和一个目标值 target，请在数组中找出两个加起来等于目标值的数的下标，返回的下标按升序排列。
（注：返回的数组下标从1开始算起，保证target一定可以由数组里面2个数字相加得到）


要求：空间复杂度 O(n)，时间复杂度 O(nlogn)
示例1
输入：
[3,2,4],6
复制
返回值：
[2,3]
复制
说明：
因为 2+4=6 ，而 2的下标为2 ， 4的下标为3 ，又因为 下标2 < 下标3 ，所以返回[2,3]
示例2
输入：
[20,70,110,150],90
复制
返回值：
[1,2]
复制
说明：
20+70=90
*/

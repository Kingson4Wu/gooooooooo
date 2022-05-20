package nowcoder

func minNumberDisappeared(nums []int) int {
	// write code here

	min := 1

	for _, v := range nums {

		if v <= 0 {
			continue
		}

	}

	return min
}

/**
个人原想法：堆排序

官方答案：骚操作
1.哈希表
2.原地hash


方法一：哈希表（推荐使用）
知识点：哈希表

哈希表是一种根据关键码（key）直接访问值（value）的一种数据结构。而这种直接访问意味着只要知道key就能在O(1)O(1)O(1)时间内得到value，因此哈希表常用来统计频率、快速检验某个元素是否出现过等。

方法二：原地哈希（扩展思路）
思路：

前面提到了数组要么缺失1～n1～n1～n中的某个数字，要么缺失n+1n+1n+1，而数组正好有下标0～n−10 ～ n-10～n−1可以对应数字1～n1～n1～n，因此只要数字1～n1～n1～n中某个数字出现，我们就可以将对应下标的值做一个标记，最后没有被标记的下标就是缺失的值。
*/

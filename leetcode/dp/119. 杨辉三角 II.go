package dp

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}
	length := rowIndex + 1
	result := make([]int, length)
	result[0] = 1

	for i := 1; i <= rowIndex; i++ {

		if i%2 == 0 {
			length := (i + 2) / 2
			result[length-1] = result[length-2] * 2

			for j := length - 2; j > 0; j-- {
				result[j] = result[j] + result[j-1]
			}
		} else {
			length := (i + 1) / 2
			for j := length - 1; j > 0; j-- {
				result[j] = result[j] + result[j-1]
			}
		}
	}

	for j := length - 1; j >= length/2; j-- {
		result[j] = result[length-1-j]
	}

	return result
}

/**
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。


你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
*/

/**

1
12
13
146
1510
1620

自己做的
算好规律。。。。。

执行结果：
通过
显示详情
添加备注

执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
1.8 MB
, 在所有 Go 提交中击败了
90.08%
的用户
通过测试用例：
34 / 34

*/

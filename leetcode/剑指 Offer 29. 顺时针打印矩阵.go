package leetcode

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func spiralOrder(matrix [][]int) []int {

	length := len(matrix)

	if length == 0 {
		return []int{}
	}

	width := len(matrix[0])
	total := width * length
	result := make([]int, total)
	index := 0

	leftIndex := 0
	rightIndex := width - 1
	topIndex := 0
	buttomIndex := length - 1

	for {

		/** 上 */
		for i := leftIndex; i <= rightIndex; i++ {
			result[index] = matrix[topIndex][i]
			index++
		}
		topIndex++
		if topIndex > buttomIndex {
			break
		}
		/** 右 */
		for i := topIndex; i <= buttomIndex; i++ {
			result[index] = matrix[i][rightIndex]
			index++
		}
		rightIndex--
		if rightIndex < leftIndex {
			break
		}
		/** 下 */
		for i := rightIndex; i >= leftIndex; i-- {
			result[index] = matrix[buttomIndex][i]
			index++
		}
		buttomIndex--
		if buttomIndex < topIndex {
			break
		}

		/** 左 */
		for i := buttomIndex; i >= topIndex; i-- {
			result[index] = matrix[i][leftIndex]
			index++
		}
		leftIndex++
		if leftIndex > rightIndex {
			break
		}

	}

	return result

}

/**

自己写的，不过想得有点久
定义上下左右下标标量，分四个方向遍历，直到左大于右 或 上大于下

执行结果：
通过
显示详情
添加备注

执行用时：
12 ms
, 在所有 Go 提交中击败了
21.21%
的用户
内存消耗：
6 MB
, 在所有 Go 提交中击败了
86.37%
的用户
通过测试用例：
27 / 27
*/

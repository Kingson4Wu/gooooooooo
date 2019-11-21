package algorithm

/**
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func generate(numRows int) [][]int {
	d := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		len := i + 1
		x := make([]int, len)
		d[i] = x

		pre := i - 1
		for j := 0; j < len; j++ {

			if j == 0 || j == len-1 {
				x[j] = 1
			} else {
				x[j] = d[pre][j-1] + d[pre][j]
			}

		}
	}
	return d
}

/**
do by myself
*/

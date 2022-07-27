package dp

func minPathSum(grid [][]int) int {

	height := len(grid)

	if height == 0 {
		return 0
	}

	dp := make([][]int, len(grid))

	width := len(grid[0])

	for i := 0; i < height; i++ {
		dp[i] = make([]int, width)

		for j := 0; j < width; j++ {

			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}

			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}

			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}

			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]

		}

	}

	return dp[height-1][width-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a

}

/**

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/

/**

dp[i,j]=min(dp[i-1,j],dp[i,j-1]) + val(i,j)

*/

/**

第一次写对？？

执行结果：
通过
显示详情
添加备注

执行用时：
4 ms
, 在所有 Go 提交中击败了
97.92%
的用户
内存消耗：
4.2 MB
, 在所有 Go 提交中击败了
43.23%
的用户
通过测试用例：
61 / 61
*/

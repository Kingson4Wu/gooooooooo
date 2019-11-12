package algorithm

/**
给你一个 n 行 m 列的矩阵，最开始的时候，每个单元格中的值都是 0。

另有一个索引数组 indices，indices[i] = [ri, ci] 中的 ri 和 ci 分别表示指定的行和列（从 0 开始编号）。

你需要将每对 [ri, ci] 指定的行和列上的所有单元格的值加 1。

请你在执行完所有 indices 指定的增量操作后，返回矩阵中 「奇数值单元格」 的数目。



示例 1：



输入：n = 2, m = 3, indices = [[0,1],[1,1]]
输出：6
解释：最开始的矩阵是 [[0,0,0],[0,0,0]]。
第一次增量操作后得到 [[1,2,1],[0,1,0]]。
最后的矩阵是 [[1,3,1],[1,3,1]]，里面有 6 个奇数。
示例 2：



输入：n = 2, m = 2, indices = [[1,1],[0,0]]
输出：0
解释：最后的矩阵是 [[2,2],[2,2]]，里面没有奇数。


提示：

1 <= n <= 50
1 <= m <= 50
1 <= indices.length <= 100
0 <= indices[i][0] < n
0 <= indices[i][1] < m

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cells-with-odd-values-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
do by myself ！！
*/
func oddCells(n int, m int, indices [][]int) int {

	x := make([]int, n, n)
	y := make([]int, m, m)

	for _, index := range indices {
		x[index[0]]++
		y[index[1]]++
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (x[i]+y[j])%2 == 1 {
				count++
			}
		}
	}

	return count
}

/**
执行用时 :
0 ms
, 在所有 golang 提交中击败了
100.00%
的用户
内存消耗 :
2.5 MB
, 在所有 golang 提交中击败了
100.00%
的用户

1. 使用数组代替hashmap的思想
2. 空间换取时间
*/

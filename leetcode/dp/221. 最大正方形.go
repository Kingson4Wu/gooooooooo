package dp

func maximalSquare(matrix [][]byte) int {

	height := len(matrix)

	if height == 0 {
		return 0
	}

	maxLen := 0
	width := len(matrix[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			if matrix[i][j] == '0' {
				continue
			}

			start := j
			end := j

			for end+1 < width && matrix[i][end+1] == 1 {
				end++
			}

			total := end - start + 1
			total = minLength(total, height-i)

			/** 判断正方形的长度 */
			for m := 0; m < total; m++ {
				for k := start; k <= end; k++ {

				}

			}

		}
	}

	return maxLen
}

func minLength(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
*/

/**

个人想法：
遍历每一行，遇到连续是1的计算下面的行是否能构成正方形

算了，不写了，感觉越写越复杂

我的想法就是暴力法，看了答案，也写的很复杂


动态规划：

dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1

func maximalSquare(matrix [][]byte) int {
    dp := make([][]int, len(matrix))
    maxSide := 0
    for i := 0; i < len(matrix); i++ {
        dp[i] = make([]int, len(matrix[i]))
        for j := 0; j < len(matrix[i]); j++ {
            dp[i][j] = int(matrix[i][j] - '0')
            if dp[i][j] == 1 {
                maxSide = 1
            }
        }
    }

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[i]); j++ {
            if dp[i][j] == 1 {
                dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
                if dp[i][j] > maxSide {
                    maxSide = dp[i][j]
                }
            }
        }
    }
    return maxSide * maxSide
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/maximal-square/solution/zui-da-zheng-fang-xing-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

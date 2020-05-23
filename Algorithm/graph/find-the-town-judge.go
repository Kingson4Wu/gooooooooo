package algorithm

func findJudge(N int, trust [][]int) int {

	/** 入度表 */
	indegress := make([]int, N)
	/** 出度表*/
	outdegress := make([]int, N)

	/** 存在入度表为N-1, 出度表为0 */

	for _, cp := range trust {
		indegress[cp[1]-1]++
		outdegress[cp[0]-1]++
	}

	for i := 0; i < N; i++ {
		if indegress[i] == N-1 && outdegress[i] == 0 {
			return i + 1
		}
	}
	return -1
}

/**
自己做的！！
*/

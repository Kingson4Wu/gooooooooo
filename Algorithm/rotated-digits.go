package algorithm

func rotatedDigits(N int) int {
	num := 0
	count := 0
	for v := N; v > 0; v /= 10 {
		num++
	}
	g := make([]int, num)
	e := make([]int, num)

	temp := N
	for i := 0; i < num; i++ {
		var s int
		if i < num-1 {
			s = temp / ((num - 1 - i) * i)
		} else {
			s = temp
		}

		if s >= 9 {
			g[i] = 7
			e[i] = 7
		} else if s >= 8 {

		}
	}

	return count

}

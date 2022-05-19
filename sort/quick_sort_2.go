package sort

func QuickSort_1(array []int, start int, end int) {

	if start >= end {
		return
	}

	middle := partition1(array, start, end)

	QuickSort_1(array, start, middle)
	QuickSort_1(array, middle+1, end)

}

//6, 4, 5, 7, 9, 2, 3, 5, 6, 8, 6, 1, 8, 2, 3, 9, 32, 5, 7

// 写错了！！！
func partition1(array []int, start int, end int) int {

	base := array[start]
	start++
	low := start
	high := end

	for low < high {
		for array[low] <= base {
			low++
		}
		for array[high] > base {
			high--
		}

		if low >= high {
			return low
		}
		array[low], array[high] = array[high], array[low]
		low++
		high--
	}
	return low

}

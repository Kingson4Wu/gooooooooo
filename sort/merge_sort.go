package sort

//分治法（Divide and Conquer）
//即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。
//递归

func MergeSort(n []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	MergeSort(n, start, mid)
	MergeSort(n, mid+1, end)
	Merge(n, start, mid, end)
}

func Merge(n []int, start, mid, end int) {
	var temp []int
	i := start
	k := mid + 1
	j := 0

	for ; i <= mid && k <= end; j++ {
		if n[i] < n[k] {
			temp = append(temp, n[i])
			i++
		} else {
			temp = append(temp, n[k])
			k++
		}
	}

	if i > mid {
		temp = append(temp, n[k:end+1]...)
	} else {
		temp = append(temp, n[i:mid+1]...)
	}

	copy(n[start:end+1], temp)
}

//原文链接：https://blog.csdn.net/guidao13/article/details/86430483

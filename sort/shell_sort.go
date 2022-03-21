package sort

//简单插入排序的改进版。
//它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序

//比较大小，直接交换

func ShellSort(n []int, len int) {
	step := len / 2
	for ; step > 0; step = step / 2 {
		for i := step; i < len; i++ {
			j := i - step
			temp := n[i]
			for j >= 0 && temp < n[j] {
				n[j+step] = n[j]
				j = j - step
			}
			n[j+step] = temp
		}
	}
}

//原文链接：https://blog.csdn.net/guidao13/article/details/86430483

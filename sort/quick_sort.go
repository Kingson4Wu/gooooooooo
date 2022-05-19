package sort

//快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

//选择一个元素作为“基准”(pivot)

func QuickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := partition(nums, start, end)
	QuickSort(nums, start, mid)
	QuickSort(nums, mid+1, end)
}

func partition(nums []int, start, end int) int {
	temp := nums[start]
	low := start
	high := end

	for low < high {
		for low < high && temp <= nums[high] {
			high--
		}
		if low < high {
			nums[low] = nums[high]
		}

		for low < high && temp > nums[low] {
			low++
		}

		if low < high {
			nums[high] = nums[low]
		}
	}

	nums[low] = temp

	return low
}

//原文链接：https://blog.csdn.net/guidao13/article/details/86430483

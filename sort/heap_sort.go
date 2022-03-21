package sort

func initHead(nums []int, parent, len int) {
	temp := nums[parent]
	child := 2*parent + 1

	for child < len {
		if child+1 < len && nums[child] < nums[child+1] {
			child++
		}

		if child < len && nums[child] <= temp {
			break
		}

		nums[parent] = nums[child]

		parent = child
		child = child*2 + 1
	}

	nums[parent] = temp
}

func HeadSort(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		initHead(nums, i, len(nums))
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]

		initHead(nums, 0, i)
	}
}

//原文链接：https://blog.csdn.net/guidao13/article/details/86430483

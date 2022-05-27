package nowcoder

func InversePairs(data []int) int {
	// write code here

	if len(data) <= 1 {
		return 0
	}

	total := 0
	result := make([]int, len(data))

	for i, v := range data {
		if i == 0 {
			result[0] = 0
			continue
		}
		prefix := i - 1

		increase := 0
		first := prefix
		for j := prefix; j >= 0; j-- {

			first = j
			if v < data[j] {
				increase++

			} else if v == data[j] {
				break
			}

		}
		result[i] = result[first] + increase
		increase = 0

		total += result[i]

	}

	return total
}

/***
做不出来！！！
结果还超时！！
*/

/**
方法二：归并排序思想
A1： 首先回答一下第一个问题，为什么归并排序需要额外空间？
显然我们知道，归并排序的过程就是，递归划分整个区间为基本相等的左右区间，之间左右区间各只有一个数字，然后就合并两个有序区间。
问题就出在了合并两个有序区间上，需要额外的空间。
为什么呢？
这里我举个例子，比如需要合并的两个有序区间为[3 4] 和 [1 2]
我们需要得到最后的结果为[1 2 3 4]， 如果不需要额外的空间的话，是做不到的，
当比较1 和 3 的时候， 1 比 3 小，就会覆盖原来的位置。

A2：回答第二个问题之前，先了解一下归并排序的过程，主要有以下两个操作：

递归划分整个区间为基本相等的左右两个区间
合并两个有序区间


明白了归并排序的过程，那么回答问题2.
如果两个区间为[4, 3] 和[1, 2]
那么逆序数为(4,1),(4,2),(3,1),(3,2)，同样的如果区间变为有序，比如[3,4] 和 [1,2]的结果是一样的，也就是说区间有序和无序结果是一样的。
但是如果区间有序会有什么好处吗？当然，如果区间有序，比如[3,4] 和 [1,2]
如果3 > 1, 显然3后面的所有数都是大于1， 这里为 4 > 1, 明白其中的奥秘了吧。所以我们可以在合并的时候利用这个规则。


//TODO
*/

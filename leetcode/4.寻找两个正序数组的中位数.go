package leetcode

/**
竟然是困难。。。
自己做的。。
基本思路，记录数组的最大最小下标，最小的之间进行对比，最大的之间进行对比，然后开始下标++，对终止下标--
*/
/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数

 4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

执行结果：
通过
显示详情
添加备注

执行用时：
16 ms
, 在所有 Go 提交中击败了
44.88%
的用户
内存消耗：
4.9 MB
, 在所有 Go 提交中击败了
67.06%
的用户
通过测试用例：
2094 / 2094
*/

/**
digits := []int{2, 1, 3, 0}
	leetcode.Exexute(digits)
*/

func Exexute(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1_length := len(nums1)
	nums2_length := len(nums2)

	total_length := nums1_length + nums2_length

	nums1_start_index, nums2_start_index := 0, 0
	nums1_end_index, nums2_end_index := nums1_length-1, nums2_length-1

	for total_length > 2 {

		if nums1_length == 0 {
			nums2_start_index++
			nums2_end_index--
			nums2_length -= 2
			total_length -= 2
			continue
		}
		if nums2_length == 0 {
			nums1_start_index++
			nums1_end_index--
			nums1_length -= 2
			total_length -= 2
			continue
		}

		//去掉最小的数
		if nums1[nums1_start_index] <= nums2[nums2_start_index] {
			nums1_start_index++
			nums1_length--
			total_length--
		} else {
			nums2_start_index++
			nums2_length--
			total_length--
		}

		//去掉最大的数
		if nums1_length == 0 {
			nums2_end_index--
			nums2_length--
			total_length--
			continue
		}
		if nums2_length == 0 {
			nums1_end_index--
			nums1_length--
			total_length--
			continue
		}

		if nums1[nums1_end_index] >= nums2[nums2_end_index] {
			nums1_end_index--
			nums1_length--
			total_length--
		} else {
			nums2_end_index--
			nums2_length--
			total_length--
		}

	}

	if total_length == 2 {
		if nums1_length == 0 {
			return float64((nums2[nums2_start_index] + nums2[nums2_end_index])) / 2
		}
		if nums2_length == 0 {
			return float64((nums1[nums1_start_index] + nums1[nums1_end_index])) / 2
		}
		return float64((nums1[nums1_start_index] + nums2[nums2_start_index])) / 2
	}

	if nums1_length == 1 {
		return float64(nums1[nums1_start_index])
	}
	return float64(nums2[nums2_start_index])

}

// @lc code=end

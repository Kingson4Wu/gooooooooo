package main

import (
	//"github.com/kingson4wu/gooooooooo/leetcode"
	//"github.com/kingson4wu/gooooooooo/nowcoder"
	"github.com/kingson4wu/gooooooooo/sort"
)

func main() {

	//nums1 := []int{0, 0, 2, 0, 4, 2, 4, 0, 7, 3, 2, 10, 7, 9, 10, 4, 1, 4, 2, 10, 3, 16, 8, 22, 18, 10, 6, 16, 8}
	//nums1 := []int{0, 0, 2, 0, 2, 2, 4, 0, 2, 3, 3, 10, 6, 8, 10, 4, 1, 7, 4, 10, 4, 16, 8, 22, 18, 10, 7, 16, 9, 2}

	//nowcoder.AdjustV2(len(nums1)-1, nums1)
	//nowcoder.GetLeastNumbers_Solution(nums1, 4)

	nums1 := []int{1}
	result := sort.QuickSortTopk_1(nums1, 0)

	for _, v := range result {
		println(v)
	}

}

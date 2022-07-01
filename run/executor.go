package main

import "github.com/kingson4wu/gooooooooo/leetcode/tree"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//nums1 := []int{0, 0, 2, 0, 4, 2, 4, 0, 7, 3, 2, 10, 7, 9, 10, 4, 1, 4, 2, 10, 3, 16, 8, 22, 18, 10, 6, 16, 8}
	//nums1 := []int{0, 0, 2, 0, 2, 2, 4, 0, 2, 3, 3, 10, 6, 8, 10, 4, 1, 7, 4, 10, 4, 16, 8, 22, 18, 10, 7, 16, 9, 2}

	//nowcoder.AdjustV2(len(nums1)-1, nums1)
	//nowcoder.GetLeastNumbers_Solution(nums1, 4)

	/* nums1 := []int{364, 637, 341, 406, 747, 995, 234, 971}
	result :=
		nowcoder.InversePairs(nums1)
	println(result) */

	/**nums1 := []int{1}
	result := sort.QuickSortTopk_1(nums1, 0)

	for _, v := range result {
		println(v)
	}*/

	/* root := &nowcoder.TreeNode{1, nil, nil}
	left := &nowcoder.TreeNode{2, nil, nil}
	right := &nowcoder.TreeNode{3, nil, nil}
	root.Left = left
	root.Right = right
	nowcoder.IsValidBST(root) */

	/* node1 := &ListNode{3, nil}
	node2 := &ListNode{5, nil}
	node1.Next = node2
	reverseBetween(node1, 1, 2) */

	//nowcoder.LongestCommonPrefix([]string{"abca", "abc", "abca", "abc", "abcc"})

	//result := nowcoder.Solve1("1", "99")
	//println(result)

	//[5,3,6,2,4,null,null,1]
	/*root := &tree.TreeNode{5, nil, nil}
	left := &tree.TreeNode{3, nil, nil}
	right := &tree.TreeNode{6, nil, nil}
	root.Left = left
	root.Right = right

	leftLeft := &tree.TreeNode{2, nil, nil}
	leftRight := &tree.TreeNode{4, nil, nil}
	left.Left = leftLeft
	left.Right = leftRight

	leftLeftLeft := &tree.TreeNode{1, nil, nil}
	leftLeft.Left = leftLeftLeft

	val := tree.KthSmallest(root, 3)
	println(val)*/

	root := &tree.TreeNode{1, nil, nil}
	left := &tree.TreeNode{2, nil, nil}
	right := &tree.TreeNode{3, nil, nil}
	root.Left = left
	root.Right = right

	leftLeft := &tree.TreeNode{4, nil, nil}
	leftRight := &tree.TreeNode{5, nil, nil}
	left.Left = leftLeft
	left.Right = leftRight

	leftLeftLeft := &tree.TreeNode{6, nil, nil}
	right.Left = leftLeftLeft
	val := tree.CountNodes(root)
	println(val)

}

/**
数据范围：数组长度 2\le n \le 10002≤n≤1000，数组中每个数的大小 0 < val \le 10000000<val≤1000000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
*/
func FindNumsAppearOnce(array []int) []int {
	// write code here

	for _, v := range array {
		array[v] = -array[v]
	}

	return []int{}
}

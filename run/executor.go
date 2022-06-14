package main

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

	node1 := &ListNode{3, nil}
	node2 := &ListNode{5, nil}
	node1.Next = node2
	reverseBetween(node1, 1, 2)

}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here

	if m == n {
		return head
	}

	index := 0
	indexHead := head

	var previous *ListNode = nil
	var startPrevious *ListNode = nil
	var start *ListNode = nil
	reverse := false

	for indexHead != nil {
		index++

		if !reverse {
			if index == m {
				start = indexHead
				startPrevious = previous
				reverse = true
			}
			previous = indexHead
			indexHead = indexHead.Next
		} else {
			temp := indexHead.Next
			indexHead.Next = previous
			previous = indexHead
			indexHead = temp

			if index == n {

				if startPrevious != nil {
					startPrevious.Next = previous
				} else {
					head = previous
				}

				start.Next = indexHead
				break
			}
		}

	}

	return head
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

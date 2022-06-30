package tree

func sortedArrayToBST(nums []int) *TreeNode {

	len := len(nums)
	rootIndex := len / 2
	rootVal := nums[rootIndex]

	root := &TreeNode{Val: rootVal}

	if rootIndex > 0 {
		root.Left = sortedArrayToBST(nums[:rootIndex])
	}
	if rootIndex < len-1 {
		root.Right = sortedArrayToBST(nums[rootIndex+1:])
	}

	return root
}

/**
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

自己做的，终于会递归了。。。。

执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
3.3 MB
, 在所有 Go 提交中击败了
99.92%
的用户
通过测试用例：
31 / 31
*/

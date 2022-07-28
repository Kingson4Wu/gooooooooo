package tree

/**
104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/
func maxDepth2(root *TreeNode) int {

	if root == nil {
		return 0
	}

	temp := []*TreeNode{}
	temp = append(temp, root)
	count := 0

	for len(temp) > 0 {

		count++
		temp2 := []*TreeNode{}

		for _, item := range temp {
			if item.Left != nil {
				temp2 = append(temp2, item.Left)
			}
			if item.Right != nil {
				temp2 = append(temp2, item.Right)
			}
		}
		temp = temp2
	}
	return count
}

/**
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	large := left
	if left < right {
		large = right
	}

	return 1 + large
}

1.递归做法，自己完成的
2.迭代法，自己完成的（广度优先）- 第二层遍历使用i<len(temp), 限定遍历的次数，可以少用一个临时数组
*/

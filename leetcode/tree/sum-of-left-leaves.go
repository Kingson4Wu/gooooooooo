package tree

/**
计算给定二叉树的所有左叶子之和。
*/

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		} else {
			sum += sumOfLeftLeaves(root.Left)
		}
	}
	if root.Right != nil {
		sum += sumOfLeftLeaves(root.Right)
	}
	return sum
}

/**
简单，自己写的
1. 少了root非空判断
2. left和right都为空才是叶子节点，少了right的判断
*/

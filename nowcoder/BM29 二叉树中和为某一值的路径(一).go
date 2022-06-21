package nowcoder

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)

}

/**

我又会递归了。。。。

运行时间：5ms
超过42.86% 用Go提交的代码
占用内存：1948KB
超过57.14%用Go提交的代码
*/

/**
给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n
*/

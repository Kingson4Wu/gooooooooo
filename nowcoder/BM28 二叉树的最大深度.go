package nowcoder

func maxDepth(root *TreeNode) int {
	// write code here

	maxDepthV := 0

	search(root, 0, &maxDepthV)
	return maxDepthV
}

func search(root *TreeNode, depth int, maxDepthV *int) {

	if root == nil {

		if depth > *maxDepthV {
			*maxDepthV = depth
		}
		return
	}

	depth++
	search(root.Left, depth, maxDepthV)
	search(root.Right, depth, maxDepthV)

}

/**
自己做的，跟中序遍历一模一样即可

运行时间：5ms
超过34.19% 用Go提交的代码
占用内存：1976KB
超过70.01%用Go提交的代码
*/

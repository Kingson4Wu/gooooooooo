package nowcoder

var maxDepthV = 0

func maxDepth(root *TreeNode) int {
	// write code here

	search(root, 0)
	return maxDepthV
}

func search(root *TreeNode, depth int) {

	if root == nil {

		if depth > maxDepthV {
			maxDepthV = depth
		}
		return
	}

	depth++
	search(root.Left, depth)
	search(root.Right, depth)

}

/**
自己做的，跟中序遍历一模一样即可

运行时间：5ms
超过34.19% 用Go提交的代码
占用内存：1976KB
超过70.01%用Go提交的代码
*/

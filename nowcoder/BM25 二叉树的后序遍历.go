package nowcoder

var sffdfgfgdgdsds = []int{}

func postorderTraversal(root *TreeNode) []int {
	// write code here

	inorder(root)

	return sffdfgfgdgdsds
}

func inorder(root *TreeNode) {

	if root == nil {
		return
	}
	inorder(root.Left)
	inorder(root.Right)
	sffdfgfgdgdsds = append(sffdfgfgdgdsds, root.Val)
}

/**

自己做的，通过中序的递归，举一反三

运行时间：4ms
超过47.31% 用Go提交的代码
占用内存：1096KB
超过48.03%用Go提交的代码
*/

package nowcoder

func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here

	if pRoot != nil {
		pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
		Mirror(pRoot.Left)
		Mirror(pRoot.Right)
	}

	return pRoot
}

/**
操作给定的二叉树，将其变换为源二叉树的镜像。
*/

/**

自己做的，想通了，就是左右递归调换

运行时间：7ms
超过25.00% 用Go提交的代码
占用内存：2088KB
超过50.00%用Go提交的代码

*/

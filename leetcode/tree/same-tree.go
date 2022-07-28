package tree

/**
100. 相同的树
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/same-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if q.Left == nil && q.Right == nil && p.Left == nil && p.Right == nil {
		return true
	}

	return isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)
}

/**
自己实现的，good！
*/

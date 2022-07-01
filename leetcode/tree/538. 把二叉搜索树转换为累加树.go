package tree

var prevSum int

func convertBST(root *TreeNode) *TreeNode {

	prevSum = 0
	convertBST2(root)
	return root
}

func convertBST2(root *TreeNode) {

	if root == nil {
		return
	}
	if root.Right != nil {
		convertBST2(root.Right)
	}

	root.Val = prevSum + root.Val
	prevSum = root.Val

	if root.Left != nil {
		convertBST2(root.Left)
	}

}

/**
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/convert-bst-to-greater-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同

*/

/**
自己做的
老套路，从右子树开始遍历即可

执行用时：
16 ms
, 在所有 Go 提交中击败了
11.31%
的用户
内存消耗：
6.6 MB
, 在所有 Go 提交中击败了
64.57%
的用户
通过测试用例：
215 / 215
*/

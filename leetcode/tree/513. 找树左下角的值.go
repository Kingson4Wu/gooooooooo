package tree

func findBottomLeftValue(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return 0
}

/**
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。
*/

/**
层次遍历

另外一种解法

深度优先搜索
使用 height 记录遍历到的节点的高度，curVal 记录高度在curHeight 的最左节点的值。在深度优先搜索时，我们先搜索当前节点的左子节点，再搜索当前节点的右子节点，然后判断当前节点的高度 height 是否大于 curHeight，如果是，那么将 curVal 设置为当前结点的值，curHeight 设置为 height。

著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func findBottomLeftValue(root *TreeNode) (curVal int) {
    curHeight := 0
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, height int) {
        if node == nil {
            return
        }
        height++
        dfs(node.Left, height)
        dfs(node.Right, height)
        if height > curHeight {
            curHeight = height
            curVal = node.Val
        }
    }
    dfs(root, 0)
    return
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/find-bottom-left-tree-value/solution/zhao-shu-zuo-xia-jiao-de-zhi-by-leetcode-weeh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

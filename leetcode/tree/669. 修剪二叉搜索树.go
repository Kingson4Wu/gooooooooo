package tree

func trimBST(root *TreeNode, low int, high int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val < low {
		trimBST(root.Left, low, high)
	}

	if root.Val > high {
		trimBST(root.Right, low, high)
	}

	return nil
}

func trimBST2(root *TreeNode, low int, high int, pre *TreeNode) {

	if root.Val < low {
		trimBST(root.Left, low, high)
	}

	if root.Val > high {
		trimBST(root.Right, low, high)
	}

}

/**
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/trim-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

方法：递归
思路

令 trim(node) 作为该节点上的子树的理想答案。我们可以递归地构建该答案。

算法

当 node.val > R，那么修剪后的二叉树必定出现在节点的左边。

类似地，当 node.val < L，那么修剪后的二叉树出现在节点的右边。否则，我们将会修剪树的两边。

class Solution {
    public TreeNode trimBST(TreeNode root, int L, int R) {
        if (root == null) return root;
        if (root.val > R) return trimBST(root.left, L, R);
        if (root.val < L) return trimBST(root.right, L, R);

        root.left = trimBST(root.left, L, R);
        root.right = trimBST(root.right, L, R);
        return root;
    }
}

作者：LeetCode
链接：https://leetcode.cn/problems/trim-a-binary-search-tree/solution/xiu-jian-er-cha-sou-suo-shu-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

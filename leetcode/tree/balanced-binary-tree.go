package tree

/**
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
https://leetcode-cn.com/problems/balanced-binary-tree/
*/
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	distance := height(root.Left) - height(root.Right)

	return distance >= -1 && distance <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	height := leftHeight
	if rightHeight > leftHeight {
		height = rightHeight
	}

	return 1 + height
}

//[1,2,2,3,3,3,3,4,4,4,4,4,4,null,null,5,5] 不过
//[1,null,2,null,3]
// height自己写的太复杂了。。。

/**
方法一：自顶向下的递归
如上 时间复杂度：O(nlogn)
空间复杂度：O(n)。如果树完全倾斜，递归栈可能包含所有节点。


方法二：自底向上的递归
方法一计算 height 存在大量冗余。每次调用 height 时，要同时计算其子树高度。但是自底向上计算，每个子树的高度只会计算一次。可以递归先计算当前节点的子节点高度，然后再通过子节点高度判断当前节点是否平衡，从而消除冗余。

时间复杂度 O(N)： N为树的节点数；最差情况下，需要递归遍历树的所有节点。
空间复杂度 O(N)： 最差情况下（树退化为链表时），系统递归需要使用 O(N) 的栈空间。

class Solution {
    public boolean isBalanced(TreeNode root) {
        return recur(root) != -1;
    }

    private int recur(TreeNode root) {
        if (root == null) return 0;
        int left = recur(root.left);
        if(left == -1) return -1;
        int right = recur(root.right);
        if(right == -1) return -1;
        return Math.abs(left - right) < 2 ? Math.max(left, right) + 1 : -1;
    }
}

//  叶子节点由 if (root == null) return 0; 和 Math.max(left, right) + 1 算的 ＝ 1

作者：jyd
链接：https://leetcode-cn.com/problems/balanced-binary-tree/solution/balanced-binary-tree-di-gui-fang-fa-by-jin40789108/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


自底向上的递归！！！！！
*/

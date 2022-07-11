package tree

func longestUnivaluePath(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxPath := 0

	leftPath := longestUnivaluePath(root.Left)
	//rightPath := longestUnivaluePath(root.Right)

	if root.Left != nil && root.Left.Val == root.Val {

		maxPath += leftPath
	}

	return 0
}

/**
给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。

两个节点之间的路径长度 由它们之间的边数表示。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-univalue-path
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

不会？？？


方法：递归
思路

我们可以将任何路径（具有相同值的节点）看作是最多两个从其根延伸出的箭头。

具体地说，路径的根将是唯一节点，因此该节点的父节点不会出现在该路径中，而箭头将是根在该路径中只有一个子节点的路径。

然后，对于每个节点，我们想知道向左延伸的最长箭头和向右延伸的最长箭头是什么？我们可以用递归来解决这个问题。

class Solution {
    int ans;
    public int longestUnivaluePath(TreeNode root) {
        ans = 0;
        arrowLength(root);
        return ans;
    }
    public int arrowLength(TreeNode node) {
        if (node == null) return 0;
        int left = arrowLength(node.left);
        int right = arrowLength(node.right);
        int arrowLeft = 0, arrowRight = 0;
        if (node.left != null && node.left.val == node.val) {
            arrowLeft += left + 1;
        }
        if (node.right != null && node.right.val == node.val) {
            arrowRight += right + 1;
        }
        ans = Math.max(ans, arrowLeft + arrowRight);
        return Math.max(arrowLeft, arrowRight);
    }
}

作者：LeetCode
链接：https://leetcode.cn/problems/longest-univalue-path/solution/zui-chang-tong-zhi-lu-jing-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

妈的答案都没看懂
todo！！！
*/

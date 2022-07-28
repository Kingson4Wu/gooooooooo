package tree

/**
二叉搜索树的第k大节点
给定一棵二叉搜索树，请找出其中第k大的节点。
https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
*/

var result11 int

func kthLargest(root *TreeNode, k int) int {

	//if root == nil {
	//	return nil
	//}

	trace(root, k)
	return result11
}

func trace(root *TreeNode, k int) int {

	count := k
	if root.Right != nil {
		count = trace(root.Right, count)
	}

	count--

	if count == 0 {
		result11 = root.Val
	}

	if root.Left != nil {
		count = trace(root.Left, count)
	}

	return count
}

/**
1. 遍历方式，自己做的，关键要使用全局变量！！！，或者使用引用
2. 迭代法，想了很久，还没想出来。。。。
其实已经快要想出来了，
右子树要全部入栈，直到右子数为空
处理完当前结点，在处理左结点，重复左结点的右子树的入栈操作
（哎，自己赶紧太复杂，没坚持写下去）

迭代
class Solution {
    public int kthLargest(TreeNode root, int k) {
        int count = 1;
        Stack<TreeNode> stack = new Stack<>();
        while (Objects.nonNull(root) || !stack.empty()) {
            while (Objects.nonNull(root)) {
                stack.push(root);
                root = root.right;
            }
            TreeNode pop = stack.pop();
            if (count == k) {
                return pop.val;
            }
            count++;
            root = pop.left;
        }
        return 0;
    }
}

作者：yang_hang
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/solution/javaji-bai-100de-xiang-xi-jie-da-di-gui-he-die-dai/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

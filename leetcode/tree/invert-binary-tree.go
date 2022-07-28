package tree

// TreeNode is ...
//https://stackoverflow.com/questions/53004291/exported-type-should-have-comment-or-be-unexported-golang-vs-code
//翻转一棵二叉树(左右翻转): https://leetcode-cn.com/problems/invert-binary-tree/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	Temp := root.Left
	root.Left = root.Right
	root.Right = Temp

	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}

/*func main() {
	root := new(TreeNode)
	root.Val = 3

	invertTree(root)
}*/

/**
自己写的，但是少了第一个root判空
*/

/**
利用前序遍历
class Solution {
        // 先序遍历--从顶向下交换
        public TreeNode invertTree(TreeNode root) {
            if (root == null) return null;
            // 保存右子树
            TreeNode rightTree = root.right;
            // 交换左右子树的位置
            root.right = invertTree(root.left);
            root.left = invertTree(rightTree);
            return root;
        }
    }
*/
//这种写法可以不用temp
//判空nil
//n&(n-1)==0  !!!!!!!!!

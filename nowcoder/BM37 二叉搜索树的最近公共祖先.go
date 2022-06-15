package nowcoder

func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	// write code here
	//这么多描述不想看。。。TODO
	return 0
}

/**

给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
1.对于该题的最近的公共祖先定义:对于有根树T的两个节点p、q，最近公共祖先LCA(T,p,q)表示一个节点x，满足x是p和q的祖先且x的深度尽可能大。在这里，一个节点也可以是它自己的祖先.
2.二叉搜索树是若它的左子树不空，则左子树上所有节点的值均小于它的根节点的值； 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值
3.所有节点的值都是唯一的。
4.p、q 为不同节点且均存在于给定的二叉搜索树中。
数据范围:
3<=节点总数<=10000
0<=节点值<=10000

方法一：搜索路径比较（推荐使用）
step 1：根据二叉搜索树的性质，从根节点开始查找目标节点，当前节点比目标小则进入右子树，当前节点比目标大则进入左子树，直到找到目标节点。这个过程成用数组记录遇到的元素。
step 2：分别在搜索二叉树中找到p和q两个点，并记录各自的路径为数组。
step 3：同时遍历两个数组，比较元素值，最后一个相等的元素就是最近的公共祖先。


方法二：一次遍历（扩展思路）
step 1：首先检查空节点，空树没有公共祖先。
step 2：对于某个节点，比较与p、q的大小，若p、q在该节点两边说明这就是最近公共祖先。
step 3：如果p、q都在该节点的左边，则递归进入左子树。
step 4：如果p、q都在该节点的右边，则递归进入右子树。

import java.util.*;
public class Solution {
    public int lowestCommonAncestor (TreeNode root, int p, int q) {
        //空树找不到公共祖先
        if(root == null)
            return -1;
        //pq在该节点两边说明这就是最近公共祖先(因为从根节点遍历的，那么最先遇到符合条件的就是最近公共祖先！！！！)
        if((p >= root.val && q <= root.val) || (p <= root.val && q >= root.val))
            return root.val;
        //pq都在该节点的左边
        else if(p <= root.val && q <= root.val)
            //进入左子树
            return lowestCommonAncestor(root.left, p, q);
        //pq都在该节点的右边
        else
            //进入右子树
            return lowestCommonAncestor(root.right, p, q);
    }
}

*/

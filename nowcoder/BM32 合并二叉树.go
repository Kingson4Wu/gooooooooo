package nowcoder

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// write code here

	return nil
}

/**
已知两颗二叉树，将它们合并成一颗二叉树。合并规则是：都存在的结点，就将结点值加起来，否则空的位置就由另一个树的结点来代替。
*/

/**

知道递归，但还是做不出来，，！！

方法一：递归前序遍历（推荐使用）

具体做法：

step 1：首先判断t1与t2是否为空，若为则用另一个代替，若都为空，返回的值也是空。
step 2：然后依据前序遍历的特点，优先访问根节点，将两个根点的值相加创建到新树中。
step 3：两棵树再依次同步进入左子树和右子树。

import java.util.*;
public class Solution {
    public TreeNode mergeTrees (TreeNode t1, TreeNode t2) {
        //若只有一个节点返回另一个，两个都为null自然返回null
        if (t1 == null)
            return t2;
        if (t2 == null)
            return t1;
        //根左右的方式递归
        TreeNode head = new TreeNode(t1.val + t2.val);
        head.left = mergeTrees(t1.left, t2.left);
        head.right = mergeTrees(t1.right, t2.right);
        return head;
    }
}

*/

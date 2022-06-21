package nowcoder

func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here

	Convert(pRootOfTree.Left)

	//TODO

	Convert(pRootOfTree.Right)

	return nil
}

/**
知道要中序递归，却没做出来，注意辅助的全局变量！！！！
*/

/**
方法一：递归中序遍历（推荐使用）
具体做法：

step 1：创建两个指针，一个指向题目中要求的链表头（head），一个指向当前遍历的前一节点（pre)。
step 2：首先递归到最左，初始化head与pre。
step 3：然后处理中间根节点，依次连接pre与当前节点，连接后更新pre为当前节点。
step 4：最后递归进入右子树，继续处理。
step 5：递归出口即是节点为空则返回。

public class Solution {
    //返回的第一个指针，即为最小值，先定为null
    public TreeNode head = null;
    //中序遍历当前值的上一位，初值为最小值，先定为null
    public TreeNode pre = null;
    public TreeNode Convert(TreeNode pRootOfTree) {
        if(pRootOfTree == null)
            //中序递归，叶子为空则返回
            return null;
        //首先递归到最左最小值
        Convert(pRootOfTree.left);
        //找到最小值，初始化head与pre
        if(pre == null){
            head = pRootOfTree;
            pre = pRootOfTree;
        }
        //当前节点与上一节点建立连接，将pre设置为当前值
        else{
            pre.right = pRootOfTree;
            pRootOfTree.left = pre;
            pre = pRootOfTree;
        }
        Convert(pRootOfTree.right);
        return head;
    }
}

*/

/**
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表
*/

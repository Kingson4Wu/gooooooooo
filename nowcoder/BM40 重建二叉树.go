package nowcoder

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here

	//TODO 哎，题解都没看懂。。。。
	return nil
}

/***
具体做法：

step 1：先根据前序遍历第一个点建立根节点。
step 2：然后遍历中序遍历找到根节点在数组中的位置。
step 3：再按照子树的节点数将两个遍历的序列分割成子数组，将子数组送入函数建立子树。
step 4：直到子树的序列长度为0，结束递归。


import java.util.*;
public class Solution {
    public TreeNode reConstructBinaryTree(int [] pre,int [] vin) {
        int n = pre.length;
        int m = vin.length;
        //每个遍历都不能为0
        if(n == 0 || m == 0)
            return null;
        //构建根节点
        TreeNode root = new TreeNode(pre[0]);
        for(int i = 0; i < vin.length; i++){
            //找到中序遍历中的前序第一个元素
            if(pre[0] == vin[i]){
                //构建左子树
                root.left = reConstructBinaryTree(Arrays.copyOfRange(pre, 1, i + 1), Arrays.copyOfRange(vin, 0, i));
                //构建右子树
                root.right = reConstructBinaryTree(Arrays.copyOfRange(pre, i + 1, pre.length), Arrays.copyOfRange(vin, i + 1, vin.length));
                break;
            }
        }
        return root;
    }
}

*/

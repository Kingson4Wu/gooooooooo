package nowcoder

func inorderTraversal(root *TreeNode) []int {
	// write code here

	result := []int{}
	stack := []*TreeNode{}

	head := root

	for head != nil {
		// 入栈
		stack = append(stack, head)
		head = head.Left
	}

	for len(stack) > 0 {
		head = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		result = append(result, head.Val)

		head = head.Right

		for head != nil {
			stack = append(stack, head)
			head = head.Left
		}
	}

	return result
}

/**
自己做的，先画图理清思路
中序遍历使用栈

运行时间：9ms
超过34.51% 用Go提交的代码
占用内存：1164KB
超过65.33%用Go提交的代码
*/

/**
递归解法：

import java.util.*;
public class Solution {
    public void inorder(List<Integer> list, TreeNode root){
        //遇到空节点则返回
        if(root == null)
            return;
        //先去左子树
        inorder(list, root.left);
        //再访问根节点
        list.add(root.val);
        //最后去右子树
        inorder(list, root.right);
    }

    public int[] inorderTraversal (TreeNode root) {
        //添加遍历结果的数组
        List<Integer> list = new ArrayList();
        //递归中序遍历
        inorder(list, root);
        //返回的结果
        int[] res = new int[list.size()];
        for(int i = 0; i < list.size(); i++)
            res[i] = list.get(i);
        return res;
    }
}


*/

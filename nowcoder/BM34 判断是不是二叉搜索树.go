package nowcoder

var max int = 0
var cccccc bool
var fffffff bool

func isValidBST(root *TreeNode) bool {
	// write code here
	if root == nil {
		return true
	}
	isValidBST2(root)
	return !cccccc
}

func isValidBST2(root *TreeNode) {

	if cccccc {
		return
	}

	if root == nil {
		return
	}
	isValidBST2(root.Left)

	if !fffffff {
		max = root.Val
		fffffff = true
	}

	if root.Val >= max {
		max = root.Val
	} else {
		cccccc = true
	}

	if cccccc {
		return
	}

	isValidBST2(root.Right)

}

/**

自己写的，中序遍历，然而代码写得很乱。。。。

运行时间：9ms
超过88.26% 用Go提交的代码
占用内存：3840KB
超过97.31%用Go提交的代码
*/

/**

import java.util.*;
public class Solution {
    int pre = Integer.MIN_VALUE;
    //中序遍历
    public boolean isValidBST (TreeNode root) {
        if (root == null)
            return true;
        //先进入左子树
        if(!isValidBST(root.left))
            return false;
        if(root.val < pre)
            return false;
        //更新最值
        pre = root.val;
        //再进入右子树
        return isValidBST(root.right);
    }
}

*/

package nowcoder

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	if pRoot.Left == nil && pRoot.Right == nil {
		return true
	}

	if pRoot.Left != nil && pRoot.Right == nil {
		return false
	}

	if pRoot.Left == nil && pRoot.Right != nil {
		return false
	}

	if pRoot.Left.Val != pRoot.Right.Val {
		return false
	}

	return isSymmetrical2(pRoot.Left, pRoot.Right)
}

func isSymmetrical2(leftRoot *TreeNode, rightRoot *TreeNode) bool {

	leftStack := []*TreeNode{}
	rightStack := []*TreeNode{}

	for leftRoot != nil && rightRoot != nil {
		if leftRoot.Val != rightRoot.Val {
			return false
		}
		leftStack = append(leftStack, leftRoot)
		leftRoot = leftRoot.Left
		rightStack = append(rightStack, rightRoot)
		rightRoot = rightRoot.Right
	}

	if len(leftStack) != len(rightStack) {
		return false
	}

	for len(leftStack) > 0 && len(rightStack) > 0 {
		leftRoot = leftStack[len(leftStack)-1]
		leftStack = leftStack[0 : len(leftStack)-1]
		rightRoot = rightStack[len(rightStack)-1]
		rightStack = rightStack[0 : len(rightStack)-1]

		if leftRoot.Val != rightRoot.Val {
			return false
		}
		if leftRoot.Right != nil && rightRoot.Left == nil {
			return false
		}
		if leftRoot.Right == nil && rightRoot.Left != nil {
			return false
		}

		leftRoot = leftRoot.Right
		rightRoot = rightRoot.Left

		for leftRoot != nil {
			leftStack = append(leftStack, leftRoot)
			leftRoot = leftRoot.Left
		}

		for rightRoot != nil {
			rightStack = append(rightStack, rightRoot)
			rightRoot = rightRoot.Right
		}

	}

	for len(leftStack) != len(rightStack) {
		return false
	}

	return true
}

/**

自己做的，基于中序遍历（栈的实现），代码量上看，感觉实现复杂了

运行时间：7ms
超过6.46% 用Go提交的代码
占用内存：1084KB
超过70.94%用Go提交的代码
*/

/**

方法一：递归（推荐使用）
知识点：二叉树递归

递归是一个过程或函数在其定义或说明中有直接或间接调用自身的一种方法，它通常把一个大型复杂的问题层层转化为一个与原问题相似的规模较小的问题来求解。因此递归过程，最重要的就是查看能不能讲原本的问题分解为更小的子问题，这是使用递归的关键。

而二叉树的递归，则是将某个节点的左子树、右子树看成一颗完整的树，那么对于子树的访问或者操作就是对于原树的访问或者操作的子问题，因此可以自我调用函数不断进入子树。



public class Solution {
    boolean recursion(TreeNode root1, TreeNode root2){
        //可以两个都为空
        if(root1 == null && root2 == null)
            return true;
        //只有一个为空或者节点值不同，必定不对称
        if(root1 == null || root2 == null || root1.val != root2.val)
            return false;
        //每层对应的节点进入递归比较
        return recursion(root1.left, root2.right) && recursion(root1.right, root2.left);
    }
    boolean isSymmetrical(TreeNode pRoot) {
        return recursion(pRoot, pRoot);
    }
}

*/

package tree

//错的
/* func isValidBST(root *TreeNode) bool {


	if root == nil{
		return true
	}
	if root.Left != nil{
		if root.Val <= root.Left.Val {
			return false
		}
	}
	if root.Right != nil{
		if root.Val >= root.Right.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
} */

var pre int
var preInit bool
var isValidBSTResult bool
func isValidBST(root *TreeNode) bool {

	preInit = false
	isValidBSTResult = true

	isValidBST2(root)
	
	return isValidBSTResult
	
}

func isValidBST2(root *TreeNode) {

	if !isValidBSTResult{
		return
	}

	if root == nil{
		return 
	}

	 isValidBST2(root.Left)

	if preInit{
		if root.Val <= pre{
			isValidBSTResult =  false
			return
		} else{
			pre = root.Val
		}
	} else{
		preInit = true
		pre = root.Val
	}

	isValidBST2(root.Right)

}

/**
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
自己做的，做错了日 (写第二遍才对，加了很多全局变量感觉很恶心，有更好的写法？？？)
（使用递归，不用栈来遍历，确实需要全局变量！）

搜索树，中序遍历
递归yyds

执行用时：
4 ms
, 在所有 Go 提交中击败了
91.12%
的用户
内存消耗：
5 MB
, 在所有 Go 提交中击败了
18.07%
的用户
通过测试用例：
80 / 80



输入：
[5,4,6,null,null,3,7]
输出：
true
预期结果：
false

更好的写法：

func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
    if root == nil {
        return true
    }
    if root.Val <= lower || root.Val >= upper {
        return false
    }
    return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/validate-binary-search-tree/solution/yan-zheng-er-cha-sou-suo-shu-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：中序遍历

func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    inorder := math.MinInt64
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    return true
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/validate-binary-search-tree/solution/yan-zheng-er-cha-sou-suo-shu-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
package tree

var result = [][]int{}
var stack = []*TreeNode{}

func pathSum(root *TreeNode, targetSum int) [][]int {

	result = [][]int{}
	stack = []*TreeNode{}
	if root == nil {
		return result
	}

	pathSum2(root, targetSum)

	return result
}

func pathSum2(root *TreeNode, targetSum int) {

	stack = append(stack, root)

	//叶子结点
	if root.Left == nil && root.Right == nil {

		if targetSum == root.Val {
			size := len(stack)
			path := make([]int, size)
			for i, v := range stack {
				path[i] = v.Val
			}
			//path[size-1] = root.Val
			result = append(result, path)

		}

		stack = stack[0 : len(stack)-1]
		return

	}

	if root.Left != nil {
		pathSum2(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		pathSum2(root.Right, targetSum-root.Val)
	}
	stack = stack[0 : len(stack)-1]

}

/**
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

自己做的

执行用时：
4 ms
, 在所有 Go 提交中击败了
85.05%
的用户
内存消耗：
4.1 MB
, 在所有 Go 提交中击败了
99.65%
的用户
通过测试用例：
115 / 115

*/

package tree

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return []string{}
	}

	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)

	result := []string{}
	if len(left) > 0 {
		for _, s := range left {
			result = append(result, strconv.Itoa(root.Val)+"->"+s)
		}
	}
	if len(right) > 0 {
		for _, s := range right {
			result = append(result, strconv.Itoa(root.Val)+"->"+s)
		}
	}
	if len(left) == 0 && len(right) == 0 {
		result = append(result, strconv.Itoa(root.Val))
	}

	return result
}

/**
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]
示例 2：

输入：root = [1]
输出：["1"]
*/

/**
深度优先搜索

自己做的，递归yyds

执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
2.2 MB
, 在所有 Go 提交中击败了
36.52%
的用户
通过测试用例：
208 / 208


*/

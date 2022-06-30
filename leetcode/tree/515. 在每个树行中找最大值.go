package tree

var max map[int]int
var maxDepth int

func largestValues(root *TreeNode) []int {

	max = make(map[int]int)
	maxDepth = 0

	largestValues2(root, 1)

	result := []int{}
	for i := 0; i < maxDepth; i++ {
		result = append(result, max[i+1])
	}

	return result
}

func largestValues2(root *TreeNode, depth int) {

	if root == nil {
		return
	}

	if depth > maxDepth {
		maxDepth = depth
	}

	if val, ok := max[depth]; ok {
		if val < root.Val {
			max[depth] = root.Val
		}
	} else {
		max[depth] = root.Val
	}

	if root.Left != nil {
		largestValues2(root.Left, depth+1)
	}
	if root.Right != nil {
		largestValues2(root.Right, depth+1)
	}

}

/**
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
*/

/**
简单搞定！
执行用时：
4 ms
, 在所有 Go 提交中击败了
94.10%
的用户
内存消耗：
5.6 MB
, 在所有 Go 提交中击败了
32.96%
的用户
通过测试用例：
78 / 78
*/

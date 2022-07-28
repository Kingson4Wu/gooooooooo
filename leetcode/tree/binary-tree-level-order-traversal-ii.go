package tree

/**
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {

		length := len(queue)
		column := []int{}
		for i, item := range queue {
			if i < length {
				column = append(column, item.Val)
				if item.Left != nil {
					queue = append(queue, item.Left)
				}
				if item.Right != nil {
					queue = append(queue, item.Right)
				}
			}
		}
		result = append(result, nil)
		copy(result[1:], result[0:])
		result[0] = column

		queue = queue[length:len(queue)]
	}
	return result
}

/**
自己完成的
*/

package tree

func widthOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxWidth := 1
	indexMap := make(map[*TreeNode]int)

	arr := []*TreeNode{}
	arr = append(arr, root)
	indexMap[root] = 1

	for len(arr) > 0 {
		newArr := []*TreeNode{}
		for i := 0; i < len(arr); i++ {
			node := arr[i]
			if node.Left != nil {
				newArr = append(newArr, node.Left)
				indexMap[node.Left] = indexMap[node]*2 - 1
			}
			if node.Right != nil {
				newArr = append(newArr, node.Right)
				indexMap[node.Right] = indexMap[node] * 2
			}
		}
		arr = newArr
		if len(arr) > 0 {
			width := indexMap[arr[len(arr)-1]] - indexMap[arr[0]] + 1
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	return maxWidth
}

/**
给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-width-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
自己做的，算清楚下标关系即可


执行用时：
4 ms
, 在所有 Go 提交中击败了
85.26%
的用户
内存消耗：
5 MB
, 在所有 Go 提交中击败了
31.79%
的用户
通过测试用例：
113 / 113
*/

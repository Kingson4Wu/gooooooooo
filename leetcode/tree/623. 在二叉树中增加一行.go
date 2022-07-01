package tree

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {

	if depth == 1 {
		node := &TreeNode{Val: val}
		node.Left = root
		return node
	}

	arr := []*TreeNode{}
	arr = append(arr, root)
	currDepth := 2

	for currDepth < depth {
		newArr := []*TreeNode{}
		for _, node := range arr {
			if node.Left != nil {
				newArr = append(newArr, node.Left)
			}
			if node.Right != nil {
				newArr = append(newArr, node.Right)
			}
		}
		arr = newArr
		currDepth++
	}

	for _, node := range arr {
		temp := node.Left
		newNode := &TreeNode{Val: val}
		node.Left = newNode
		newNode.Left = temp

		temp = node.Right
		newNode = &TreeNode{Val: val}
		node.Right = newNode
		newNode.Right = temp

	}

	return root
}

/**
给定一个二叉树的根 root 和两个整数 val 和 depth ，在给定的深度 depth 处添加一个值为 val 的节点行。

注意，根节点 root 位于深度 1 。

加法规则如下:

给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，创建两个值为 val 的树节点作为 cur 的左子树根和右子树根。
cur 原来的左子树应该是新的左子树根的左子树。
cur 原来的右子树应该是新的右子树根的右子树。
如果 depth == 1 意味着 depth - 1 根本没有深度，那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/add-one-row-to-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
自己做的

广度优先搜索即可

输入：
[1,2,3,4]
5
4
输出：
[1,2,3,4]
预期结果：
[1,2,3,4,null,null,null,5,5]

执行用时：
4 ms
, 在所有 Go 提交中击败了
84.52%
的用户
内存消耗：
5.5 MB
, 在所有 Go 提交中击败了
58.33%
的用户
通过测试用例：
109 / 109
*/

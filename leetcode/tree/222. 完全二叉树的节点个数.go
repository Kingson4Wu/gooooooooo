package tree

var nums int
var maxDepth1 int

func CountNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxDepth1 = 0

	node := root
	for node != nil {
		maxDepth1++
		node = node.Left
	}

	nums = 1 << (maxDepth1 - 1)
	calculate(root, 1)

	return nums - 1
}

func calculate(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		if depth == maxDepth1 {
			//depth 不会判断。。。
			nums++
		}
		return
	}
	if root.Left != nil {
		calculate(root.Left, depth+1)
	}

	if root.Right != nil {
		calculate(root.Right, depth+1)
	}

}

/* func CountNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	stack := []*TreeNode{}

	maxDepth := 0

	for root != nil {
		maxDepth++
		stack = append(stack, root)
		root = root.Left
	}

	nums := 1 << (maxDepth - 1)
	depth := maxDepth

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			//叶子结点
			if root.Left == nil && root.Right == nil {
				if depth == maxDepth {
					//depth 不会判断。。。
					nums++
				} else {
					break
				}

			}
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}

	return nums - 1
} */

/**

给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/count-complete-tree-nodes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

广度优先搜索也行，不过相对低效

深度优先搜索
基于中序遍历非递归改造！！
（不会，还是先算深度，再用递归的方式完成了。。。)


执行用时：
12 ms
, 在所有 Go 提交中击败了
93.16%
的用户
内存消耗：
7.1 MB
, 在所有 Go 提交中击败了
54.59%
的用户
通过测试用例：
18 / 18

*/

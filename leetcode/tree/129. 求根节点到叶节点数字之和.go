package tree

func sumNumbers(root *TreeNode) int {

	sum := 0
	stack := []*TreeNode{}

	stack = append(stack, root)

	for len(stack) > 0 {

		root = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		/** 叶子结点 */
		if root.Left == nil && root.Right == nil {

			num := 0
			for i := 0; i < len(stack); i++ {
				num = num * 10
				num += stack[i].Val
			}
			num = num * 10
			num += root.Val
			sum += num
			continue
		}

		if root.Left != nil {

		}

	}

	return 0
}

/**
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sum-root-to-leaf-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
突然间不会写深度优先搜索！！！！
*/

/**
func dfs(root *TreeNode, prevSum int) int {
    if root == nil {
        return 0
    }
    sum := prevSum*10 + root.Val
    if root.Left == nil && root.Right == nil {
        return sum
    }
    return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
    return dfs(root, 0)
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/sum-root-to-leaf-numbers/solution/qiu-gen-dao-xie-zi-jie-dian-shu-zi-zhi-he-by-leetc/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


方法二：广度优先搜索

*/

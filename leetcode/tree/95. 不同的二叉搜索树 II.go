package tree

func generateTrees(n int) []*TreeNode {

	return nil
}

/* func buildTree(n int, k int) []*TreeNode {

	left_size := k - 1
	right_size := n - k

	rootArr := []*TreeNode{}
	for i := 1; i <= left_size; i++ {
		for j := k + 1; j <= right_size; j++ {
			root := &TreeNode{Val: k}
			root.Left = buildTree(left_size, i)
			root.Right = buildTree(right_size, j)

			rootArr = append(rootArr, root)
		}
	}

	return root

} */

/**

给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
*/

/**

明知道是递归，却做不出来？？？

方法一：回溯

二叉搜索树关键的性质是根节点的值大于左子树所有节点的值，小于右子树所有节点的值，且左子树和右子树也同样为二叉搜索树。
因此在生成所有可行的二叉搜索树的时候，假设当前序列长度为 n，如果我们枚举根节点的值为 i，
那么根据二叉搜索树的性质我们可以知道左子树的节点值的集合为 [1…i−1]，右子树的节点值的集合为 [i+1…n]。
而左子树和右子树的生成相较于原问题是一个序列长度缩小的子问题，因此我们可以想到用回溯的方法来解决这道题目。


func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return nil
    }
    return helper(1, n)
}

func helper(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    allTrees := []*TreeNode{}
    // 枚举可行根节点
    for i := start; i <= end; i++ {
        // 获得所有可行的左子树集合
        leftTrees := helper(start, i - 1)
        // 获得所有可行的右子树集合
        rightTrees := helper(i + 1, end)
        // 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
        for _, left := range leftTrees {
            for _, right := range rightTrees {
                currTree := &TreeNode{i, nil, nil}
                currTree.Left = left
                currTree.Right = right
                allTrees = append(allTrees, currTree)
            }
        }
    }
    return allTrees
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/unique-binary-search-trees-ii/solution/bu-tong-de-er-cha-sou-suo-shu-ii-by-leetcode-solut/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



*/

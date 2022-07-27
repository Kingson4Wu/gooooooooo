package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	if root == nil {
		return []int{}
	}

	//path := []*TreeNode{}

	result := []int{}

	return result
}

func seekPath(path *[]*TreeNode, root *TreeNode, target *TreeNode) {

}

func calculateDepth(root *TreeNode, depth int, result *[]int) {

	if root == nil {
		return
	}

	if depth == 0 {
		return
	}
	if depth == 1 {
		*result = append(*result, root.Val)
		return
	}
	calculateDepth(root.Left, depth-1, result)
	calculateDepth(root.Right, depth-1, result)

}

/**
给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。

返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
深度优先搜索

1.递归找出路径
2.判断是否相等得到往哪边回溯
3.计算路径的函数

知道路径怎么走，但就是写不出代码。。。。。
*/

/**
方法一：深度优先搜索 + 哈希表
若将 target 当作树的根结点，我们就能从 target 出发，使用深度优先搜索去寻找与 target 距离为 kk 的所有结点，即深度为 kk 的所有结点。

由于输入的二叉树没有记录父结点，为此，我们从根结点 root 出发，使用深度优先搜索遍历整棵树，同时用一个哈希表记录每个结点的父结点。

然后从 target 出发，使用深度优先搜索遍历整棵树，除了搜索左右儿子外，还可以顺着父结点向上搜索。

代码实现时，由于每个结点值都是唯一的，哈希表的键可以用结点值代替。此外，为避免在深度优先搜索时重复访问结点，递归时额外传入来源结点 from，在递归前比较目标结点是否与来源结点相同，不同的情况下才进行递归。
(左右子树是往下的，父节点是往上的，所以可能相互重复)

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/solution/er-cha-shu-zhong-suo-you-ju-chi-wei-k-de-qbla/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**

func distanceK(root, target *TreeNode, k int) (ans []int) {
    // 从 root 出发 DFS，记录每个结点的父结点
    parents := map[int]*TreeNode{}
    var findParents func(*TreeNode)
    findParents = func(node *TreeNode) {
        if node.Left != nil {
            parents[node.Left.Val] = node
            findParents(node.Left)
        }
        if node.Right != nil {
            parents[node.Right.Val] = node
            findParents(node.Right)
        }
    }
    findParents(root)

    // 从 target 出发 DFS，寻找所有深度为 k 的结点
    var findAns func(*TreeNode, *TreeNode, int)
    findAns = func(node, from *TreeNode, depth int) {
        if node == nil {
            return
        }
        if depth == k { // 将所有深度为 k 的结点的值计入结果
            ans = append(ans, node.Val)
            return
        }
        if node.Left != from {
            findAns(node.Left, node, depth+1)
        }
        if node.Right != from {
            findAns(node.Right, node, depth+1)
        }
        if parents[node.Val] != from {
            findAns(parents[node.Val], node, depth+1)
        }
    }
    findAns(target, nil, 0)
    return
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/solution/er-cha-shu-zhong-suo-you-ju-chi-wei-k-de-qbla/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

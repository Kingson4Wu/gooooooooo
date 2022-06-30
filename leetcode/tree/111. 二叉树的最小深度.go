package tree

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := -1
	if root.Left != nil {
		leftDepth = minDepth(root.Left)
	}
	rightDepth := -1
	if root.Right != nil {
		rightDepth = minDepth(root.Right)
	}
	if leftDepth == -1 {
		return rightDepth + 1
	}
	if rightDepth == -1 {
		return leftDepth + 1
	}

	if leftDepth < rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1

}

/**

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。
*/

/**
果然是简单

执行用时：
168 ms
, 在所有 Go 提交中击败了
39.32%
的用户
内存消耗：
18.8 MB
, 在所有 Go 提交中击败了
45.43%
的用户
通过测试用例：
52 / 52

[2,null,3,null,4,null,5,null,6]
*/

/**
写得真丑

官方答案也差不多丑，用minD := math.MaxInt32代替我的-1

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    minD := math.MaxInt32
    if root.Left != nil {
        minD = min(minDepth(root.Left), minD)
    }
    if root.Right != nil {
        minD = min(minDepth(root.Right), minD)
    }
    return minD + 1
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/minimum-depth-of-binary-tree/solution/er-cha-shu-de-zui-xiao-shen-du-by-leetcode-solutio/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

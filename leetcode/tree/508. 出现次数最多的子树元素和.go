package tree

func findFrequentTreeSum(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	sumArr := []int{}

	var leftSum []int
	var rightSum []int
	if root.Left != nil {

		leftSum = findFrequentTreeSum(root.Left)

		for _, sum := range leftSum {
			sumArr = append(sumArr, sum+root.Val)
		}

	}

	if root.Right != nil {

		rightSum = findFrequentTreeSum(root.Right)

		for _, sum := range rightSum {
			sumArr = append(sumArr, sum+root.Val)
		}

	}

	if len(leftSum) > 0 && len(rightSum) > 0 {

	}

	return []int{}
}

/**

给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。


输入: root = [5,2,-3]
输出: [2,-3,4]


输入: root = [5,2,-5]
输出: [2]

二叉树所有子树（包括包含根节点的树）的元素和当中，出现次数最多的，返回

所有的子树和为： [2, -3, 4] 都是出现一次，所以返回所有
所有的子树和为： [2, -5, 2] 2出现两次 ，为出现的次数最多，返回出现次数最多的元素 ：[2]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/most-frequent-subtree-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

不会！！！

方法一：深度优先搜索
我们可以从根结点出发，深度优先搜索这棵二叉树。对于每棵子树，其子树元素和等于子树根结点的元素值，加上左子树的元素和，以及右子树的元素和。

用哈希表统计每棵子树的元素和的出现次数，计算出现次数的最大值 \textit{maxCnt}maxCnt，最后将出现次数等于 \textit{maxCnt}maxCnt 的所有元素和返回。

func findFrequentTreeSum(root *TreeNode) (ans []int) {
    cnt := map[int]int{}
    maxCnt := 0
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        sum := node.Val + dfs(node.Left) + dfs(node.Right)
        cnt[sum]++
        if cnt[sum] > maxCnt {
            maxCnt = cnt[sum]
        }
        return sum
    }
    dfs(root)

    for s, c := range cnt {
        if c == maxCnt {
            ans = append(ans, s)
        }
    }
    return
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/most-frequent-subtree-sum/solution/chu-xian-ci-shu-zui-duo-de-zi-shu-yuan-s-kdjc/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


*/

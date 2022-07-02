package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	maxVal := nums[0]
	maxIndex := 0

	for i, num := range nums {
		if num > maxVal {
			maxVal = num
			maxIndex = i
		}
	}

	root := &TreeNode{Val: nums[maxIndex]}

	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

/**
给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:

创建一个根节点，其值为 nums 中的最大值。
递归地在最大值 左边 的 子数组前缀上 构建左子树。
递归地在最大值 右边 的 子数组后缀上 构建右子树。
返回 nums 构建的 最大二叉树 。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
自己做的，递归yyds

执行用时：
12 ms
, 在所有 Go 提交中击败了
93.96%
的用户
内存消耗：
6.9 MB
, 在所有 Go 提交中击败了
65.81%
的用户
通过测试用例：
107 / 107
*/

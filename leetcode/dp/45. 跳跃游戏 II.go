package dp

func jump(nums []int) int {

	length := len(nums)

	if length == 1 {
		return 0
	}

	leastSteps := make([]int, length)

	for i := 0; i < length; i++ {
		newStep := leastSteps[i] + 1
		for j := 0; j <= nums[i]; j++ {

			if leastSteps[i+j] == 0 || newStep < leastSteps[i+j] {
				leastSteps[i+j] = newStep
			}
			if leastSteps[length-1] > 0 {
				return leastSteps[length-1]
			}
		}
	}
	return 0
}

/**
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。

示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/jump-game-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

好像竟然被我想出来了，计算到达每一步的最小步数！


执行结果：
解答错误
显示详情
添加备注

通过测试用例：
107 / 109
输入：
[0]
输出：
1
预期结果：
0

执行结果：
解答错误
显示详情
添加备注

通过测试用例：
108 / 109
输入：
[1]
输出：
1
预期结果：
0

需考虑数组长度为1 的情况 ！！！！

执行结果：
通过
显示详情
添加备注

执行用时：
40 ms
, 在所有 Go 提交中击败了
22.64%
的用户
内存消耗：
6.1 MB
, 在所有 Go 提交中击败了
17.67%
的用户
通过测试用例：
109 / 109


还有贪心算法的写法更高效！！todo



*/

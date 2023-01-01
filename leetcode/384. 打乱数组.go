package leetcode

import (
	"math/rand"
)

/**
给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。

实现 Solution class:

Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果
*/

type Solution3 struct {
	nums   []int
	origin []int
}

func Constructor3(nums []int) Solution3 {

	n := make([]int, len(nums))
	copy(n[:], nums[:])
	s := Solution3{nums: n, origin: nums}

	return s
}

func (this *Solution3) Reset() []int {

	copy(this.nums[:], this.origin[:])
	return this.nums
}

func (this *Solution3) Shuffle() []int {

	l := len(this.nums)
	for i := 0; i < l; i++ {
		index := rand.Intn(l-i) + i
		this.nums[index], this.nums[i] = this.nums[i], this.nums[index]
	}
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

/**
我的思路,保存两个数组,一个原始的,一个打乱后的
时间
36 ms
击败
10.85%
内存
7.8 MB
击败
29.25%

原来我自己想的好像就是洗牌算法....

*/

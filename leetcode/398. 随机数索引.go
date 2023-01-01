package leetcode

import "math/rand"

type Solution2 struct {
	data map[int][]int
}

func Constructor2(nums []int) Solution2 {

	data := make(map[int][]int)
	for i, v := range nums {
		if arr, ok := data[v]; ok {
			arr = append(arr, i)
			data[v] = arr
		} else {
			data[v] = []int{i}
		}
	}

	return Solution2{data: data}

}

func (this *Solution2) Pick(target int) int {

	arr := this.data[target]

	return arr[rand.Intn(len(arr))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

/**
我的思路, hash, key是数值, value是下标数组
时间
112 ms
击败
94.55%
内存
9.7 MB
击败
54.55%


*/

/**
方法一：哈希表
如果不考虑数组的大小，我们可以在构造函数中，用一个哈希表 pos 记录 nums中相同元素的下标。

对于 pick 操作，我们可以从 pos 中取出 target 对应的下标列表，然后随机选择其中一个下标并返回。

方法二：水塘抽样
如果数组以文件形式存储（读者可假设构造函数传入的是个文件路径），且文件大小远超内存大小，我们是无法通过读文件的方式，将所有下标保存在内存中的，因此需要找到一种空间复杂度更低的算法。

type Solution []int

func Constructor(nums []int) Solution {
    return nums
}

func (nums Solution) Pick(target int) (ans int) {
    cnt := 0
    for i, num := range nums {
        if num == target {
            cnt++ // 第 cnt 次遇到 target
            if rand.Intn(cnt) == 0 {
                ans = i
            }
        }
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/random-pick-index/solutions/1444589/sui-ji-shu-suo-yin-by-leetcode-solution-ofsq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

空间换时间，时间换空间
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
}

func Constructor(head *ListNode) Solution {

	return Solution{}
}

func (this *Solution) GetRandom() int {

	return 0
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

/**

 方法一：记录所有链表元素
我们可以在初始化时，用一个数组记录链表中的所有元素，这样随机选择链表的一个节点，就变成在数组中随机选择一个元素。

type Solution []int

func Constructor(head *ListNode) (s Solution) {
    for node := head; node != nil; node = node.Next {
        s = append(s, node.Val)
    }
    return s
}

func (s Solution) GetRandom() int {
    return s[rand.Intn(len(s))]
}


方法二：水塘抽样
方法一需要花费 O(n)O(n)O(n) 的空间存储链表中的所有元素，那么能否做到 O(1)O(1)O(1) 的空间复杂度呢？

我们可以设计如下算法：

从链表头开始，遍历整个链表，对遍历到的第 iii 个节点，随机选择区间 [0,i)[0,i)[0,i) 内的一个整数，如果其等于 000，则将答案置为该节点值，否则答案不变。

type Solution struct {
    head *ListNode
}

func Constructor(head *ListNode) Solution {
    return Solution{head}
}

func (s *Solution) GetRandom() (ans int) {
    for node, i := s.head, 1; node != nil; node = node.Next {
        if rand.Intn(i) == 0 { // 1/i 的概率选中（替换为答案）
            ans = node.Val
        }
        i++
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/linked-list-random-node/solutions/1210211/lian-biao-sui-ji-jie-dian-by-leetcode-so-x6it/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



*/

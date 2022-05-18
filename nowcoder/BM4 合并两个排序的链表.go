package nowcoder

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here

	pHead1_index := pHead1
	pHead2_index := pHead2

	var resultHead *ListNode
	var resultTail *ListNode

	for pHead1_index != nil || pHead2_index != nil {
		if pHead1_index == nil && pHead2_index != nil {
			if resultTail == nil {
				resultTail = pHead2_index
			} else {
				resultTail.Next = pHead2_index
				resultTail = resultTail.Next
			}
			if resultHead == nil {
				resultHead = pHead2_index
			}
			pHead2_index = pHead2_index.Next
			continue
		}

		if pHead2_index == nil && pHead1_index != nil {
			if resultTail == nil {
				resultTail = pHead1_index
			} else {
				resultTail.Next = pHead1_index
				resultTail = resultTail.Next
			}
			if resultHead == nil {
				resultHead = pHead1_index
			}
			pHead1_index = pHead1_index.Next
			continue
		}

		if pHead1_index.Val <= pHead2_index.Val {
			if resultTail == nil {
				resultTail = pHead1_index
			} else {
				resultTail.Next = pHead1_index
				resultTail = resultTail.Next
			}
			if resultHead == nil {
				resultHead = pHead1_index
			}
			pHead1_index = pHead1_index.Next
			continue
		} else {
			if resultTail == nil {
				resultTail = pHead2_index
			} else {
				resultTail.Next = pHead2_index
				resultTail = resultTail.Next
			}
			if resultHead == nil {
				resultHead = pHead2_index
			}
			pHead2_index = pHead2_index.Next
		}
	}

	return resultHead
}

/**

运行时间：8ms
超过31.48% 用Go提交的代码
占用内存：984KB
超过81.27%用Go提交的代码

自己写的，不过感觉代码写得有点丑
*/

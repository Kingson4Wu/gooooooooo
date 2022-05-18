package nowcoder

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here

	if pHead == nil {
		return nil
	}

	var previous *ListNode = nil

	for pHead.Next != nil {

		current := pHead

		pHead = pHead.Next

		current.Next = previous

		previous = current

	}

	pHead.Next = previous

	return pHead
}

/**
自己做的，不过搞得挺久，要画画图才能理清逻辑

运行时间：5ms
超过25.75% 用Go提交的代码
占用内存：892KB
超过54.58%用Go提交的代码
*/

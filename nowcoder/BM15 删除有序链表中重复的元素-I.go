package nowcoder

/**
 *
 * @param head ListNode类
 * @return ListNode类
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// write code here

	if head == nil || head.Next == nil {
		return head
	}

	start := head
	end := head

	for end.Next != nil {

		end = end.Next

		if start.Val == end.Val {
			continue
		}

		start.Next = end
		start = end

	}
	start.Next = nil

	return head
}

/**
自己做的

运行时间：4ms
超过37.21% 用Go提交的代码
占用内存：1048KB
超过52.36%用Go提交的代码
*/

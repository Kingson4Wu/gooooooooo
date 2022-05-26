package nowcoder

/**
自己做的，但是感觉写的有点乱

链表的问题基本都挺容易想，但是写起来比价麻烦，需要许多中间变量，所以做之前要先画图整理清楚，否则写的时候很容易乱

*/

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here

	if head == nil {
		return head
	}

	newHead := head

	var tail *ListNode
	start := head
	len := 0
	has := false

	for head != nil {

		len++
		if len == k {

			previous := tail
			oldTail := tail
			tail = start

			for i := 0; i < k; i++ {

				if i == k-1 {
					if oldTail != nil {
						oldTail.Next = start
					}
				}

				temp := start.Next
				start.Next, previous = previous, start
				start = temp

			}
			tail.Next = start
			if !has {
				newHead = previous
				has = true

			}
			len = 0
			head = start

		} else {
			head = head.Next
		}

	}

	////if len > 0 {
	//previous.Next = tail

	//}

	return newHead
}

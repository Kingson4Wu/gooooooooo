package linked_list

// ListNode is ...

/*
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return head
	}

	temp := []*ListNode{}

	temp = append(temp, head)

	next := head

	for next.Next != nil {
		next = next.Next
		temp = append(temp, next)
	}

	len := len(temp)

	if n == 1 {
		if len == 1 {
			return nil
		}
		temp[len-n-1].Next = nil

	} else {
		if n == len {
			head = head.Next
		} else {
			temp[len-n-1].Next = temp[len-n+1]

		}
	}
	return head
}

/**
使用数组索引遍历一遍，再根据下标删除（自己的想法）

[1]
1
长度为1的情况忽略了。。。

[1,2]
2

忽略了删除的是头或者尾的场景！！！

方法二：一次遍历算法
算法

上述算法可以优化为只使用一次遍历。我们可以使用两个指针而不是一个指针。第一个指针从列表的开头向前移动 n+1n+1 步，而第二个指针将从列表的开头出发。现在，这两个指针被 nn 个结点分开。我们通过同时移动两个指针向前来保持这个恒定的间隔，直到第一个指针到达最后一个结点。此时第二个指针将指向从最后一个结点数起的第 nn 个结点。我们重新链接第二个指针所引用的结点的 next 指针指向该结点的下下个结点。

作者：LeetCode
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-by-l/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

两个指针，快慢指针！！！

*/

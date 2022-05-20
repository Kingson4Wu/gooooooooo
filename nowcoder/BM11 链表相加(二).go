package nowcoder

/**
 *
 * @param head1 ListNode类
 * @param head2 ListNode类
 * @return ListNode类
 */
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here

	//个人想法
	//1 反转相加再反转

	//官方题解2
	/**
	方法二：使用辅助栈
	上一个方式是直接对原来的两个链表进行了反转，这个方法则是借助了栈的先进后出的特性来充当链表的反转，因为我们其实是想从两个链表的尾部进行开始操作，所以我们干脆直接将两条链表的结点放进栈中，然后依次出栈操作即可，然后相加完后采用头插法即可得到最终的链表。
	*/

	return head1
}

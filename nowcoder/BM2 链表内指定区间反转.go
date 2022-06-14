package nowcoder

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here

	if m == n {
		return head
	}

	index := 0
	indexHead := head

	var previous *ListNode = nil
	var startPrevious *ListNode = nil
	var start *ListNode = nil
	reverse := false

	for indexHead != nil {
		index++

		if !reverse {
			if index == m {
				start = indexHead
				startPrevious = previous
				reverse = true
			}
			previous = indexHead
			indexHead = indexHead.Next
		} else {
			temp := indexHead.Next
			indexHead.Next = previous
			previous = indexHead
			indexHead = temp

			if index == n {

				if startPrevious != nil {
					startPrevious.Next = previous
				} else {
					head = previous
				}

				start.Next = indexHead
				break
			}
		}

	}

	return head
}

/**
将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)，空间复杂度 O(1)。
*/

/**
运行时间：4ms
超过54.04% 用Go提交的代码
占用内存：1040KB
超过68.70%用Go提交的代码

/**
自己完成的，不过写得比较久
链表的题目比较简单，但是要提前画好图，理清楚需要哪些变量，这样写起来思路会清晰很多！！！
*/

/**
import java.util.*;
public class Solution {
    public ListNode reverseBetween (ListNode head, int m, int n) {
        //加个表头
        ListNode res = new ListNode(-1);
        res.next = head;
        //前序节点
        ListNode pre = res;
        //当前节点
        ListNode cur = head;
        //找到m
        for(int i = 1; i < m; i++){
            pre = cur;
            cur = cur.next;
        }
        //从m反转到n
        for(int i = m; i < n; i++){
            ListNode temp = cur.next;
            cur.next = temp.next;
            temp.next = pre.next;
            pre.next = temp;
        }
        //返回去掉表头
        return res.next;
    }
}

*/

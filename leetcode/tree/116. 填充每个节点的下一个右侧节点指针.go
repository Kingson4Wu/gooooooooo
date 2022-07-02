package tree

func connect3(root *Node) *Node {

	if root == nil {
		return nil
	}

	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) > 0 {
		newQueue := []*Node{}

		pre := queue[0]

		for i := 0; i < len(queue); i++ {

			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}

			if i > 0 {
				pre.Next = queue[i]
				pre = queue[i]
			}

		}
		queue = newQueue

	}

	return root
}

/**

给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

/**
层级遍历

自己做的


执行用时：
4 ms
, 在所有 Go 提交中击败了
90.09%
的用户
内存消耗：
6.4 MB
, 在所有 Go 提交中击败了
45.23%
的用户
通过测试用例：
59 / 59

方法二：使用已建立的 next 指针 ！！！！
空间复杂度：O(1)，不需要存储额外的节点。


*/

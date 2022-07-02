package tree

func flatten(root *TreeNode) {

	if root == nil {
		return
	}
	help(root, nil)

}

func help(root *TreeNode, pre *TreeNode) {

	leftTemp := root.Left
	rightTemp := root.Right
	if pre != nil {
		pre.Right = root
	}
	root.Left = nil

	if leftTemp != nil {
		help(leftTemp, root)
	}

	if rightTemp != nil {
		if leftTemp != nil {
			help(rightTemp, leftTemp)
		} else {
			help(rightTemp, root)
		}
	}

}

/**

做错了。。。。

输入
[1,2,5,3,4,null,6]
输出
[1,null,2,null,5,null,6]
预期结果
[1,null,2,null,3,null,4,null,5,null,6]
*/

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

方法三：寻找前驱节点
前两种方法都借助前序遍历，前序遍历过程中需要使用栈存储节点。有没有空间复杂度是 O(1)O(1) 的做法呢？

注意到前序遍历访问各节点的顺序是根节点、左子树、右子树。如果一个节点的左子节点为空，则该节点不需要进行展开操作。
如果一个节点的左子节点不为空，则该节点的左子树中的最后一个节点被访问之后，该节点的右子节点被访问。
该节点的左子树中最后一个被访问的节点是左子树中的最右边的节点，也是该节点的前驱节点。
因此，问题转化成寻找当前节点的前驱节点。

具体做法是，对于当前节点，如果其左子节点不为空，则在其左子树中找到最右边的节点，作为前驱节点，
将当前节点的右子节点赋给前驱节点的右子节点，然后将当前节点的左子节点赋给当前节点的右子节点，
并将当前节点的左子节点设为空。对当前节点处理结束后，继续处理链表中的下一个节点，直到所有节点都处理结束。

func flatten(root *TreeNode)  {
    curr := root
    for curr != nil {
        if curr.Left != nil {
            next := curr.Left
            predecessor := next
            for predecessor.Right != nil {
                predecessor = predecessor.Right
            }
            predecessor.Right = curr.Right
            curr.Left, curr.Right = nil, next
        }
        curr = curr.Right
    }
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/solution/er-cha-shu-zhan-kai-wei-lian-biao-by-leetcode-solu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

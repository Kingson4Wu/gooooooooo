package tree

func sortedListToBST(head *ListNode) *TreeNode {

	return nil
}

/**

给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var globalHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
    globalHead = head
    length := getLength(head)
    return buildTree(0, length - 1)
}

func getLength(head *ListNode) int {
    ret := 0
    for ; head != nil; head = head.Next {
        ret++
    }
    return ret
}

func buildTree(left, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := (left + right + 1) / 2
    root := &TreeNode{}
    root.Left = buildTree(left, mid - 1)
    root.Val = globalHead.Val
    globalHead = globalHead.Next
    root.Right = buildTree(mid + 1, right)
    return root
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/solution/you-xu-lian-biao-zhuan-huan-er-cha-sou-suo-shu-1-3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

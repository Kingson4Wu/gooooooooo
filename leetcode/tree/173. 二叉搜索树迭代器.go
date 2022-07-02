package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
	root  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {

	return BSTIterator{stack: []*TreeNode{}, root: root}
}

func (this *BSTIterator) Next() int {

	val := 0

	for this.root != nil || len(this.stack) > 0 {
		if this.root != nil {
			this.stack = append(this.stack, this.root)
			this.root = this.root.Left
		} else {
			//TODO
			this.root = this.stack[len(this.stack)-1]
			this.stack = this.stack[0 : len(this.stack)-1]

			val = this.root.Val
			this.root = this.root.Right
			return val
		}
	}
	return val
}

func (this *BSTIterator) HasNext() bool {

	return this.root != nil || len(this.stack) > 0
}

func call(root *TreeNode) {

	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			//TODO
			root = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			root = root.Right
		}

	}

}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

/**

 自己完成的，其实就是中序遍历的栈实现，改成迭代器的形式

执行用时：
24 ms
, 在所有 Go 提交中击败了
49.48%
的用户
内存消耗：
9.5 MB
, 在所有 Go 提交中击败了
81.67%
的用户
通过测试用例：
61 / 61
*/

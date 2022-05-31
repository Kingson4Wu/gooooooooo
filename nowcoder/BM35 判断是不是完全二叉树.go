package nowcoder

func isCompleteTree(root *TreeNode) bool {
	// write code here

	queue := []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	previous := true

	for len(queue) > 0 {

		ele := queue[0]
		queue = queue[1:]

		if ele.Left != nil {
			if !previous {
				return false
			}
			queue = append(queue, ele.Left)
		} else {
			previous = false
		}
		if ele.Right != nil {
			if !previous {
				return false
			}
			queue = append(queue, ele.Right)
		} else {
			previous = false
		}

	}

	return true
}

/**

自己做的，广度优先搜索，队列，判断是不是连续两个不存在

运行时间：3ms
超过100.00% 用Go提交的代码
占用内存：1192KB
超过22.70%用Go提交的代码
*/

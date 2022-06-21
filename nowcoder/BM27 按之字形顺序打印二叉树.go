package nowcoder

func Print(pRoot *TreeNode) [][]int {
	// write code here
	result := [][]int{}

	if pRoot == nil {
		return result
	}

	queue := []*TreeNode{}
	queue = append(queue, pRoot)
	result = append(result, []int{pRoot.Val})
	depth := 1

	for len(queue) > 0 {

		line := []int{}
		newQueue := []*TreeNode{}

		for _, node := range queue {
			if node.Left != nil {
				if depth%2 == 0 {
					line = append(line, node.Left.Val)
				} else {
					line = append(line, 0)
					copy(line[1:], line[0:])
					line[0] = node.Left.Val
				}

				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				if depth%2 == 0 {
					line = append(line, node.Right.Val)
				} else {
					line = append(line, 0)
					copy(line[1:], line[0:])
					line[0] = node.Right.Val
				}
				newQueue = append(newQueue, node.Right)
			}
		}
		if len(line) > 0 {
			result = append(result, line)
		}

		queue = newQueue
		depth++
	}

	return result
}

/**
跟这个一模一样
nowcoder/BM26 求二叉树的层序遍历.go

运行时间：6ms
超过14.29% 用Go提交的代码
占用内存：1096KB
超过71.43%用Go提交的代码


给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）
*/

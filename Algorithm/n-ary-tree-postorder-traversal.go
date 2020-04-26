package algorithm

// Node is ...
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {

	result := []int{}
	stack := []*Node{}

	if root == nil {
		return result
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		result = append(result, node.Val)

		for _, item := range node.Children {
			stack = append(stack, item)
		}
	}

	s := make([]int, len(result), len(result))
	for i, item := range result {
		s[len(result)-1-i] = item
	}
	return s
}

/**
1. 递归改成迭代，使用栈
2. 结果反转

golang 队列和栈的实现:https://www.jianshu.com/p/43ef9a4c458b
*/

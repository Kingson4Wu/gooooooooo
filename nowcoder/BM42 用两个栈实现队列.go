package nowcoder

var stack1 []int
var stack2 []int

func Push(node int) {

	initialize()
	stack1 = append(stack1, node)

}

func Pop() int {
	initialize()

	if len(stack2) == 0 && len(stack1) > 0 {

		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = nil
	}
	if len(stack2) > 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[0 : len(stack2)-1]

		return node
	}
	return 0
}

func initialize() {

	if stack1 == nil {
		stack1 = []int{}
	}
	if stack2 == nil {
		stack2 = []int{}
	}
}

/**

自己做的

运行时间：4ms
超过14.78% 用Go提交的代码
占用内存：1052KB
超过44.03%用Go提交的代码

*/

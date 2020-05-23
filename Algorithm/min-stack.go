package algorithm

//MinStack is
type MinStack struct {
	array  []int
	minVal int
}

/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Constructor ...
/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{
		array:  []int{},
		minVal: 0,
	}
}

//Push ...
func (stack *MinStack) Push(x int) {

	if len(stack.array) == 0 {
		stack.minVal = x
	}

	stack.array = append(stack.array, x)
	if x < stack.minVal {
		stack.minVal = x
	}
}

//Pop ...
func (stack *MinStack) Pop() {

	if len(stack.array) == 0 {
		return
	}

	node := stack.array[len(stack.array)-1]
	stack.array = stack.array[0 : len(stack.array)-1]

	if node == stack.minVal {
		if len(stack.array) == 0 {
			stack.minVal = 0
			return
		}

		stack.minVal = stack.array[0]
		for _, item := range stack.array {
			if item < stack.minVal {
				stack.minVal = item
			}
		}
	}
}

// Top ...
func (stack *MinStack) Top() int {

	if len(stack.array) == 0 {
		return 0
	}

	node := stack.array[len(stack.array)-1]
	return node
}

//GetMin ...
func (stack *MinStack) GetMin() int {

	if len(stack.array) == 0 {
		return 0
	}

	return stack.minVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func test() {
	obj := Constructor()
	obj.Push(3)
	obj.Pop()
	//param_3 := obj.Top()
	//param_4 := obj.GetMin()
}

/**
自己实现的
*/
/**
不是常数的操作实现，竟然能过，
空间换时间
用一个结构体保存每个结点对应当时的最小值，或使用另外一个栈，原栈结点指向对应的第二个栈的最小值

*/
/**
type MinStack struct {
    stack []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minStack: []int{math.MaxInt64},
    }
}

func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    top := this.minStack[len(this.minStack)-1]
    this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
    this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/min-stack/solution/zui-xiao-zhan-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

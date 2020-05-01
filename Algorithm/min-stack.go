package algorithm

//MinStack is
type MinStack struct {
	array  []int
	minVal int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{
		array:  []int{},
		minVal: 0,
	}
}

func (this *MinStack) Push(x int) {

	if len(this.array) == 0 {
		this.minVal = x
	}

	this.array = append(this.array, x)
	if x < this.minVal {
		this.minVal = x
	}
}

func (this *MinStack) Pop() {

	if len(this.array) == 0 {
		return
	}

	node := this.array[len(this.array)-1]
	this.array = this.array[0 : len(this.array)-1]

	if node == this.minVal {
		if len(this.array) == 0 {
			this.minVal = 0
			return
		}

		this.minVal = this.array[0]
		for _, item := range this.array {
			if item < this.minVal {
				this.minVal = item
			}
		}
	}
}

func (this *MinStack) Top() int {

	if len(this.array) == 0 {
		return 0
	}

	node := this.array[len(this.array)-1]
	return node
}

func (this *MinStack) GetMin() int {

	if len(this.array) == 0 {
		return 0
	}

	return this.minVal
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

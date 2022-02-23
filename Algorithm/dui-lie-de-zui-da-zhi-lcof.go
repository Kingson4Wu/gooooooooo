package algorithm

//MaxQueue is ...
type MaxQueue struct {
	queue []int
	deque []int
}

//Constructor is ...
func Constructor1() MaxQueue {
	return MaxQueue{
		queue: []int{},
		deque: []int{},
	}
}

//Max_value is ...
func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]

}

//Push_back is ...
func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)

	if len(this.deque) == 0 {
		this.deque = append(this.deque, value)
	} else {
		index := -1
		for i := len(this.deque) - 1; i >= 0; i-- {
			if this.deque[i] < value {
				index = i
			} else {
				break
			}
		}
		if index != -1 {
			//rear := append([]int{}, this.deque[index:]...)
			this.deque = append(this.deque[:index], value)
			//this.deque = append(this.deque, rear...)
		} else {
			this.deque = append(this.deque, value)
		}
	}

}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}

	value := this.queue[0]
	this.queue = this.queue[1:]

	if this.deque[0] == value {
		this.deque = this.deque[1:]
	}

	return value
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

//[null,-1,-1,-1,null,46,46,-1,-1,null,868,-1,-1,null,525,-1,-1,-1,null,null,646,null,646,646,646,null,123,871,null,871,871,871,646,null,null,null,null,229,871,837,285,45,837,null,null,140,null,null,null,null,837,806,null,806,806,545,806,806,806,null,561,null,237,806,806,806,null,633,null,null,null,98,866,806,866,866,866,717,null,186,null,null,268,null,29,null,866,239,null,3,850,310,null,null,806,null,674,null,null,770]
//[null,-1,-1,-1,null,46,46,-1,-1,null,868,-1,-1,null,525,-1,-1,-1,null,null,646,null,646,646,646,null,123,871,null,871,871,871,646,null,null,null,null,229,871,837,285,45,837,null,null,140,null,null,null,null,837,806,null,806,806,545,806,806,806,null,561,null,237,806,806,806,null,633,null,null,null,98,866,806,866,866,866,717,null,186,null,null,268,null,29,null,866,239,null,3,850,310,null,null,770,null,674,null,null,770]

/**
本题中 max_value、push_back、pop_front 就是一些 API 函数，我们需要来设计这些函数以供他人直接调用，并且调用每个函数时，时间复杂度均为 \mathcal{O}(1)O(1)。

解题思路
为了解决上述问题，我们只需记住当前最大值出队后，队列里的下一个最大值即可。

具体方法是使用一个双端队列 dequedeque，在每次入队时，如果 dequedeque 队尾元素小于即将入队的元素 valuevalue，则将小于 valuevalue 的元素全部出队后，再将 valuevalue 入队；否则直接入队。

这时，辅助队列 dequedeque 队首元素就是队列的最大值。

作者：z1m
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/solution/ru-he-jie-jue-o1-fu-za-du-de-api-she-ji-ti-by-z1m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

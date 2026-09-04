type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	var stack MinStack
	stack.stack = make([]int, 0)

	return stack
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}

func (this *MinStack) GetMin() int {
	min := this.stack[0]

	for _, num := range this.stack {
		if num < min {
			min = num
		}
	}

	return min
}

package stack

type MinStack struct {
	dataStack IntListStack
	minStack  IntListStack
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.dataStack.Empty() {
		this.dataStack.Push(val)
		this.minStack.Push(val)
	} else {
		this.dataStack.Push(val)
		curMin := this.minStack.Peek()
		if curMin > val {
			this.minStack.Push(val)
		} else {
			this.minStack.Push(curMin)
		}
	}
}

func (this *MinStack) Pop() {
	if this.dataStack.Empty() {
		return
	}
	this.dataStack.Pop()
	this.minStack.Pop()
}

func (this *MinStack) Top() int {
	if this.dataStack.Empty() {
		panic("empty")
	}
	return this.dataStack.Peek()
}

func (this *MinStack) GetMin() int {
	if this.minStack.Empty() {
		panic("empty")
	}
	return this.minStack.Peek()
}

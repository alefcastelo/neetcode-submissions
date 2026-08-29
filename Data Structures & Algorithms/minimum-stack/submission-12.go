type MinStack struct {
	vals []int
	top int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	this.top = val
}

func (this *MinStack) Pop() {
	this.vals = this.vals[:len(this.vals)-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	min := this.vals[0]

	for _, val := range this.vals {
		if val < min {
			min = val
		}
	}
	
	return min
}

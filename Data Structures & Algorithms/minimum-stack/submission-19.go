type MinStack struct {
	vals []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}

	this.vals = this.vals[:len(this.vals) - 1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.vals) == 0 {
		return 0
	}

	min := this.vals[0]

	for _, val := range this.vals {
		if val < min {
			min = val
		}
	}
	
	return min
}
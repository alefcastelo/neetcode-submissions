type MinStack struct {
	vals []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
		return 
	}

	lastMin := this.mins[len(this.mins) - 1]

	if val <= lastMin {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}

	lastVal := this.vals[len(this.vals) - 1]
	this.vals = this.vals[:len(this.vals) - 1]

	lastMin := this.mins[len(this.mins) - 1]

	if lastVal == lastMin {
		this.mins = this.mins[:len(this.mins) - 1]
	}
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
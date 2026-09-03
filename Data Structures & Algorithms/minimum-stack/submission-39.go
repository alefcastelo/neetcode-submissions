type MinStack struct {
	nextPos int
	vals []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		nextPos: 0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.vals) == this.nextPos {
		this.vals = append(this.vals, val)
	} else {
		this.vals[this.nextPos] = val	
	}

	this.nextPos++

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
	if this.nextPos == 0 {
		return
	}

	lastVal := this.vals[this.nextPos - 1]
	this.nextPos--
	
	lastMin := this.mins[len(this.mins) - 1]

	if lastVal == lastMin {
		this.mins = this.mins[:len(this.mins) - 1]
	}
}

func (this *MinStack) Top() int {
	return this.vals[this.nextPos - 1]
}

func (this *MinStack) GetMin() int {
	if this.nextPos == 0 {
		return 0
	}

	return this.mins[len(this.mins) - 1]
}
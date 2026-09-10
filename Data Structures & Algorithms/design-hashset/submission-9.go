type Node struct {
	Key int
	Next *Node
}

func (this *Node) Print(){
	fmt.Printf(" -> %d", this.Key)

	if this.Next != nil {
		this.Next.Print()
	}
}

type MyHashSet struct {
	DummyNode *Node
}

func Constructor() MyHashSet {
	return MyHashSet{
		DummyNode: &Node{},
	}
}

func (this *MyHashSet) Add(key int) {
	curr := this.DummyNode

	for curr.Next != nil {
		if curr.Next.Key == key {
			return
		}

		curr = curr.Next
	}

	curr.Next = &Node{
		Key: key,
	}
}

func (this *MyHashSet) Remove(key int) {
	curr := this.DummyNode

	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			return
		}

		curr = curr.Next
	}
}

func (this *MyHashSet) Contains(key int) bool {
    curr := this.DummyNode

	for curr.Next != nil {
		if curr.Next.Key == key {
			return true
		}

		curr = curr.Next
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 
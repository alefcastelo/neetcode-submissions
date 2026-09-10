type Node struct {
	Key int
	Val int
	Next *Node
}

type MyHashMap struct {
	DummyNode *Node
}

func Constructor() MyHashMap {
	return MyHashMap{
		DummyNode: &Node{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
	curr := this.DummyNode

	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next.Val = value
			return
		}

		curr = curr.Next
	}

	curr.Next = &Node{
		Key: key,
		Val: value,
	}
}

func (this *MyHashMap) Get(key int) int {
	curr := this.DummyNode

	for curr.Next != nil {
		if curr.Next.Key == key {
			return curr.Next.Val
		}

		curr = curr.Next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	curr := this.DummyNode

	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			return
		}

		curr = curr.Next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
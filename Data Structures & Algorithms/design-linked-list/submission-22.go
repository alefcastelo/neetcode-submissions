type Node struct {
	prev *Node
	val int
	next *Node
}

type MyLinkedList struct {
	dummyHead *Node
	dummyTail *Node
}

func Constructor() MyLinkedList {
	dummyHead := &Node{}
	dummyTail := &Node{}
	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead

	return MyLinkedList{
		dummyHead: dummyHead,
		dummyTail: dummyTail,
	}
}

func (this *MyLinkedList) Get(index int) int {
    node := this.dummyHead.next

	for index > 0 {
		node = node.next
		index--

		if node == this.dummyTail {
			return -1
		}
	}

	return node.val
}

func (this *MyLinkedList) AddAtHead(val int)  {
	dummyHead, currHead := this.dummyHead, this.dummyHead.next

	newHead := &Node{
		val: val,
	}

	dummyHead.next = newHead
	currHead.prev = newHead
	newHead.prev = dummyHead
	newHead.next = currHead
}


func (this *MyLinkedList) AddAtTail(val int)  {
	currTail, dummyTail := this.dummyTail.prev, this.dummyTail

	newTail := &Node{
		val: val,
	}

	currTail.next = newTail
	dummyTail.prev = newTail
	newTail.prev = currTail
	newTail.next = dummyTail
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    node := this.dummyHead.next

	for index > 0 {
		node = node.next
		index--

		if node == this.dummyTail {
			break
		}
	}

	prev, next := node.prev, node

	newNode := &Node{
		prev: prev,
		val: val,
		next: next,
	}

	prev.next = newNode
	next.prev = newNode
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	 node := this.dummyHead.next

	for index > 0 {
		node = node.next
		index--

		if node == this.dummyTail {
			return
		}
	}

	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

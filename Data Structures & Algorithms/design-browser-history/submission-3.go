type Node struct {
	prev *Node
	url string
	next *Node
}

type BrowserHistory struct {
	position *Node
	dummyHead *Node
	dummyTail *Node
}


func Constructor(homepage string) BrowserHistory {
	dummyHead := &Node{}
	dummyTail := &Node{}

	curr := &Node{
		prev: dummyHead, 
		url: homepage,
		next: dummyTail,
	}
	
	dummyHead.next = curr
	dummyTail.prev = curr

	return BrowserHistory{
		position: curr,
		dummyHead: dummyHead,
		dummyTail: dummyTail,
	}
}


func (this *BrowserHistory) Visit(url string)  {
	newNode := &Node{url: url}

	this.position.next = newNode
	this.dummyTail.prev = newNode

	newNode.next = this.dummyTail
	newNode.prev = this.position

	this.position = newNode
}


func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 {
		if this.position.prev == this.dummyHead {
			break
		}

		this.position = this.position.prev
		steps--
	}

	return this.position.url
}


func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 {
		if this.position.next == this.dummyTail {
			break
		}

		this.position = this.position.next
		steps--
	}

	return this.position.url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
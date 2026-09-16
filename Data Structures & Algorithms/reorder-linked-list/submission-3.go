/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}
	
	firstHalf, fast, secondHalf := &ListNode{Next:head}, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		secondHalf = secondHalf.Next
		firstHalf = firstHalf.Next
	}

	firstHalf.Next = nil

	reverse := &ListNode{}
	for secondHalf != nil {
		tmp := secondHalf.Next
		secondHalf.Next = reverse.Next
		reverse.Next = secondHalf
		secondHalf = tmp
	}

	firstHalf = head
	secondHalf = reverse.Next

	for firstHalf.Next != nil && secondHalf != nil {
		rest1 := firstHalf.Next
		rest2 := secondHalf.Next

		firstHalf.Next = nil
		secondHalf.Next = nil

		secondHalf.Next = rest1
		firstHalf.Next = secondHalf

		firstHalf = rest1
		secondHalf = rest2
	}

	firstHalf.Next = secondHalf
}
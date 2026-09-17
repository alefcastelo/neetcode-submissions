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

	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	first := head
	second := slow.Next
	slow.Next = nil

	reversed := &ListNode{}
	for second != nil {
		rest := second.Next
		second.Next = reversed.Next
		reversed.Next = second
		second = rest
	}

	second = reversed.Next

	for second != nil {
		firstRest :=  first.Next
		secondRest := second.Next

		second.Next = nil
		first.Next = nil

		second.Next = firstRest
		first.Next = second

		second = secondRest
		first = first.Next.Next
	}
}
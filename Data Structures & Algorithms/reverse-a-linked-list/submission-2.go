/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	dummyNode := &ListNode{}

	curr := head

	for curr != nil {
		newNode := &ListNode{
			Val: curr.Val,
			Next: dummyNode.Next,
		}

		dummyNode.Next = newNode
		curr = curr.Next
	}
    

	return dummyNode.Next
}

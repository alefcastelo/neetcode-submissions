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
		dummyNode.Next = &ListNode{
			Val: curr.Val,
			Next: dummyNode.Next,
		}
		
		curr = curr.Next
	}
    

	return dummyNode.Next
}

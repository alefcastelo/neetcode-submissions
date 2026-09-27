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
		temp := curr
		curr = curr.Next
		temp.Next = dummyNode.Next
		dummyNode.Next = temp
	}
    
	return dummyNode.Next
}

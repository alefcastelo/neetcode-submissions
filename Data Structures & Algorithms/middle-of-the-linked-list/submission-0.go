/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	fastDummyNode := head
	slowDummyNode := head

	for fastDummyNode != nil && fastDummyNode.Next != nil {
		fastDummyNode = fastDummyNode.Next.Next
		slowDummyNode = slowDummyNode.Next
	}

	return slowDummyNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}

	curr := dummyNode

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			continue
		}

		curr = curr.Next
	}

	return dummyNode.Next
}

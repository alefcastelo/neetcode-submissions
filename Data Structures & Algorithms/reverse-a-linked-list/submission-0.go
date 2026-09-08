/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	stack, node := []int{}, head

	dummyHead := &ListNode{}

	for node != nil {
		stack = append(stack, node.Val)
		node = node.Next
	}

	for len(stack) > 0 {
		top := stack[0]
		stack = stack[1:]

		newHead := &ListNode{
			Val: top,
			Next: dummyHead.Next,
		}

		dummyHead.Next = newHead
	}

	return dummyHead.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyLeft := &ListNode{
		Next: list1,
	}

	dummyRight := &ListNode{
		Next: list2,
	}

	left := dummyLeft
	right := dummyRight
	dummyNode := &ListNode{}
	currNode := dummyNode

	for left.Next != nil || right.Next != nil {
		if right.Next == nil || (left.Next != nil && left.Next.Val <= right.Next.Val ) {
			newNode := &ListNode{
				Val: left.Next.Val,
			}

			currNode.Next = newNode
			left = left.Next
		} else {
			newNode := &ListNode{
				Val: right.Next.Val,
			}

			currNode.Next = newNode
			right = right.Next
		}
		
		currNode = currNode.Next
	}

	return dummyNode.Next
}

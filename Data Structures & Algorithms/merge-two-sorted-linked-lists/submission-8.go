/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead

	for list1 != nil || list2 != nil {

		if list1 != nil && (list2 == nil || list1.Val <= list2.Val) {
			newNode := &ListNode{
				Val: list1.Val,
			}

			curr.Next = newNode
			curr = curr.Next
			list1 = list1.Next
			continue
		}

		if list2 != nil && (list1 == nil || list2.Val <= list1.Val) {
			newNode := &ListNode{
				Val: list2.Val,
			}

			curr.Next = newNode
			curr = curr.Next
			list2 = list2.Next
		}
	}

	return dummyHead.Next
}

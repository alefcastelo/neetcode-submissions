/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA, currB := headA, headB

	for currA != currB {
		currA = swap(currA, headB)
		currB = swap(currB, headA)
	}

	return currA
}

func swap(node, otherHead *ListNode) *ListNode {
	if node == nil {
		return otherHead
	}

	return node.Next
}

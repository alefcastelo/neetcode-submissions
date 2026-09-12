/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNodes(head *ListNode, m int, n int) *ListNode {

    dummyNode := &ListNode{
        Next: head,
    }

    curr := dummyNode
    currM := m
    currN := n

    for curr.Next != nil {
        if currM > 0 {
            currM--
            curr = curr.Next
            continue
        }

        if currN > 0 {
            currN--
            curr.Next = curr.Next.Next
            continue
        }

        currM = m
        currN = n
    }

    return dummyNode.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	ordered := []int{}

	for head != nil {
		ordered = append(ordered, head.Val)
		head = head.Next
	}

	i := 0
	j := len(ordered) - 1
	for i <= j {

		if ordered[i] != ordered[j]{
			return false
		}

		i++
		j--
	}

	return true
}

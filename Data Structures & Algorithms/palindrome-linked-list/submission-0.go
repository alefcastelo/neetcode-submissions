import (
	"slices"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	ordered := []int{}
	reversed := []int{}

	for head != nil {
		reversed = append([]int{head.Val}, reversed...)
		ordered = append(ordered, head.Val)
		head = head.Next
	}

	return slices.Equal(ordered, reversed)
}

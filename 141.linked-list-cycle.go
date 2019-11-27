/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return ok
		}
		m[head] = true
		head = head.Next
	}
	return false
}

// @lc code=end


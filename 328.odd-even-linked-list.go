/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddNode := head
	evenNode := head.Next
	evenNodeHead := evenNode

	for evenNode != nil && evenNode.Next != nil {
		oddNode.Next = evenNode.Next
		oddNode = oddNode.Next
		evenNode.Next = oddNode.Next
		evenNode = evenNode.Next
	}

	oddNode.Next = evenNodeHead

	return head
}

// @lc code=end


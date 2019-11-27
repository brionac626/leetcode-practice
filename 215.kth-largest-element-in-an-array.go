/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}
// @lc code=end


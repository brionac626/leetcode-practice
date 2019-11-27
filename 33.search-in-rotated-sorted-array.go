/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	result := -1
	for i, v := range nums {
		if v == target {
			result = i
			break
		}
	}

	return result
}

// @lc code=end


/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	for i, v := range nums {
		if v == target {
			result[0] = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			result[1] = i
			break
		}
	}

	return result
}

// @lc code=end


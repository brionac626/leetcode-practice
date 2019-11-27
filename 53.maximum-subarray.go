/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum, max := nums[0], nums[0]
	for i := 1; i < n; i++ {
		sum += nums[i]
		if sum <= nums[i] {
			sum = nums[i]
		}

		if sum >= max {
			max = sum
		}
	}

	return max
}

// @lc code=end


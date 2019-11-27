/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}

	n := len(nums)
	max := 1
	dp := make([]int, n)
	d[0] = 1

	for {

	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


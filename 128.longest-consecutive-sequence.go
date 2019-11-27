/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sort.Ints(nums)

	longest, current := 1, 1

	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				current++
			} else {
				longest = max(longest, current)
				current = 1
			}
		}
	}

	return max(longest, current)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


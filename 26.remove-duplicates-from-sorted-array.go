/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	count := 1
	currentNumber := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] != currentNumber {
			nums[count] = nums[i]
			currentNumber = nums[i]
			count++
		}
	}

	return count 
}

// @lc code=end


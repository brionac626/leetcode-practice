/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	count := 0
	var isNagitive bool
	if divisor < 0 {
		isNagitive = true
		divisor = -divisor
	}

	if dividend < 0 {
		if isNagitive {
			isNagitive = false
		}
		dividend = -dividend
	}

	for dividend-divisor >= 0 {
		dividend -= divisor
		count++
	}

	if isNagitive {
		count = -count
	}
	return count
}

// @lc code=end


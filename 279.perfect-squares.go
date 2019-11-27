/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {
	result := make([]int, n+1)

	for i := 1; i <= n; i++ {
		result[i] = i
		for j := 1; j*j <= i; j++ {
			result[i] = min(result[i], result[i-j*j]+1)
		}
	}

	return result[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end


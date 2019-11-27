/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
	m := make([]bool, n+1)
	for i := range m {
		m[i] = true
	}

	var temp, result int
	for i := 2; i < n; i++ {
		if !m[i] {
			continue
		}
		result++

		for j := 2; ; j++ {
			temp = i * j
			if temp > n {
				break
			}
			m[temp] = false
		}
	}

	return result
}

// @lc code=end


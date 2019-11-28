/*
 * @lc app=leetcode id=961 lang=golang
 *
 * [961] N-Repeated Element in Size 2N Array
 */

// @lc code=start
func repeatedNTimes(A []int) int {
	m := make(map[int]struct{})
	for _, v := range A {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = struct{}{}
	}
	return 0
}

// @lc code=end


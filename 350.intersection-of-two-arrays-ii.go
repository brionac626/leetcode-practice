/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		if c, ok := m[v]; ok {
			m[v] = c + 1
		} else {
			m[v] = 1
		}
	}

	var result []int
	for _, v := range nums2 {
		if c, ok := m[v]; ok && c > 0 {
			result = append(result, v)
			m[v] = c - 1
		}
	}

	return result
}

// @lc code=end


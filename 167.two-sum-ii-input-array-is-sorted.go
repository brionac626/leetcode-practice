/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	result := []int{0, 0}
	for i, v := range numbers {
		m[v] = i
	}

	for i1, v := range numbers {
		if i2, ok := m[target-v]; ok {
			result = []int{i1 + 1, i2 + 1}
			break
		}
	}

	return result
}

// @lc code=end


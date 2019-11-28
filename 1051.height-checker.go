/*
 * @lc app=leetcode id=1051 lang=golang
 *
 * [1051] Height Checker
 */

// @lc code=start
func heightChecker(heights []int) int {
	sortHeights := make([]int, len(heights))
	copy(sortHeights, heights)
	sort.Ints(sortHeights)
	
	var count int
	for i, v := range heights {
		if v != sortHeights[i] {
			count++
		}
	}
	return count
}

// @lc code=end


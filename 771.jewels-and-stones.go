/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]struct{})
	for _, v := range J {
		jewels[v] = struct{}{}
	}
	var result int
	for _, v := range S {
		_, ok := jewels[v]
		if ok {
			result++
		}
	}
	return result
}

// @lc code=end


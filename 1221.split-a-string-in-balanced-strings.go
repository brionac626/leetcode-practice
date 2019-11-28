/*
 * @lc app=leetcode id=1221 lang=golang
 *
 * [1221] Split a String in Balanced Strings
 */

// @lc code=start
func balancedStringSplit(s string) int {
	var l, r, result int
	for _, v := range s {
		switch v {
		case 'L':
			l++
		case 'R':
			r++
		}
		if l == r {
			l = 0
			r = 0
			result++
		}
	}

	return result
}

// @lc code=end


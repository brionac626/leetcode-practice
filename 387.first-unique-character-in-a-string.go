/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	n := make([]int, 26)
	for _, v := range s {
		n[int(v-'a')]++
	}

	for i, v := range s {
		if n[int(v-'a')] == 1 {
			return i
		}
	}

	return -1
}

// @lc code=end

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}

	return -1
}

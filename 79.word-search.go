/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	letters := make(map[byte]int)
	for i := 0; i < len(board); i++ {
		for _, v := range board[i] {
			if c, ok := letters[v]; ok {
				c++
			} else {
				letters[v]++
			}
		}
	}

	for _, v := range word {
		c, ok := letters[byte(v)]
		if !ok {
			return false
		}

		c--
		if c < 0 {
			return false
		}
	}

	return true
}

// @lc code=end


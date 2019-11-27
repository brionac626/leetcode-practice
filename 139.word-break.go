/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	checkIndex := make([]bool, len(s)+1)
	checkIndex[0] = true

	for end := 1; end <= len(s); end++ {
		for start := 0; start <= end-1; start++ {
			if _, ok := dict[s[start:end]]; checkIndex[start] == true && ok {
				checkIndex[end] = true
				break
			}
		}
	}

	return checkIndex[len(s)]
}

// @lc code=end


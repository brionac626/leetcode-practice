/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {		
		words :=strings.Split(strings.TrimSpace(s)," ")		
		return len(words[len(words)-1])
	
}
// @lc code=end


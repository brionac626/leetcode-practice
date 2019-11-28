/*
 * @lc app=leetcode id=804 lang=golang
 *
 * [804] Unique Morse Code Words
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	transform := make(map[string]struct{})
	for _, word := range words {
		morseCodeTransform := ""
		for _, v := range word {
			morseCodeTransform += morseCode[v-'a']
		}
		transform[morseCodeTransform] = struct{}{}
	}
	return len(transform)
}

// @lc code=end


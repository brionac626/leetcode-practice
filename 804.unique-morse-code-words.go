/*
 * @lc app=leetcode id=804 lang=golang
 *
 * [804] Unique Morse Code Words
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	letters := make(map[rune]string)
	letterStart := 97
	for i, v := range morseCode {
		letters[rune(letterStart+i)] = v
	}

	transform := make(map[string]struct{})
	for _, word := range words {
		morseCodeTransform := ""
		for _, v := range word {
			morseCodeTransform += letters[v]
		}
		transform[morseCodeTransform] = struct{}{}
	}
	return len(transform)
}

// @lc code=end


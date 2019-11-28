/*
 * @lc app=leetcode id=832 lang=golang
 *
 * [832] Flipping an Image
 */

// @lc code=start
func flipAndInvertImage(A [][]int) [][]int {
	limit := len(A)
	for _, image := range A {
		for i := 0; i < (limit+1)/2; i++ {
			image[i], image[limit-1-i] = image[limit-1-i]^1, image[i]^1
		}
	}

	return A
}

// @lc code=end


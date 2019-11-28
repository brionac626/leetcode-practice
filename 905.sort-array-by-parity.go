/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

// @lc code=start
func sortArrayByParity(A []int) []int {
	var odd, even []int
	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return append(even, odd...)
}

// @lc code=end

// in-place
func sortArrayByParity(A []int) []int {
	start, end := 0, len(A)-1
	for start < end {
		if A[start]%2 > A[end]%2 {
			A[start], A[end] = A[end], A[start]
		}

		if A[start]%2 == 0 {
			start++
		}
		if A[end]%2 == 1 {
			end--
		}
	}

	return A
}
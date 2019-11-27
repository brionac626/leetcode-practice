/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 {
		return x
	}

	start, end := 0, x
	var mid, sqr int

	for {
		mid = (start + end) / 2
		if start == mid {
			return mid
		}

		sqr = mid * mid
		if sqr == x {
			return mid
		} else if sqr > x {
			end = mid
		} else if sqr < x {
			start = mid
		}
	}

	return 0
}

// @lc code=end


/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 */

// @lc code=start
type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{origin: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	pickRange := len(this.origin)
	shuffled := make([]int, pickRange)
	copy(shuffled, this.origin)
	for i := 0; i < pickRange; i++ {
		pickIndex := rand.Intn(pickRange)
		shuffled[i], shuffled[pickIndex] = shuffled[pickIndex], shuffled[i]
	}

	return shuffled
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end


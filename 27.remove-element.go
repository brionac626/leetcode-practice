/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	indexLimit := len(nums)-1
    for i,v := range nums {
		if v == val {
			if i + 1 > indexLimit{
				nums = nums[:i-1]
			}else{
				nums = nums[i+1:]
			}
		}
	}
	return len(nums)
}
// @lc code=end


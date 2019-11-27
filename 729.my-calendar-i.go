/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 */

// @lc code=start
type MyCalendar struct {
	calender map[int]struct{}
}

func Constructor() MyCalendar {
	return MyCalendar{calender: make(map[int]struct{})}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if _, notAvailable := this.calender[start]; notAvailable {
		return false
	}

	for i := start; i <= end; i++ {
		this.calender[i] = struct{}{}
	}
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end


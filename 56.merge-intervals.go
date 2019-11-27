/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Sort(sortByFrist(intervals))	
	interLen := len(intervals)

	for i := 0; i < interLen-1; {
		current,next := intervals[i],intervals[i+1]
		if next[0] <= current[1] {
			if next[1] > current[1] {
				intervals[i][1] = intervals[i+1][1]
			}
			intervals = append(intervals[:i+1],intervals[i+2:]...)
			interLen--
			continue
		}
		i++
	}
	return intervals
}

type sortByFrist [][]int

func (bf sortByFrist) Len() int {
	return len(bf)
}

func (bf sortByFrist) Swap(i, j int) {
	bf[i], bf[j] = bf[j], bf[i]
}

func (bf sortByFrist) Less(i, j int) bool {
	return bf[i][0] < bf[j][0]
}

// @lc code=end


/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
type record struct {
	key   int
	count int
}

type numberRecords []record

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]*record)

	for _, v := range nums {
		if r, ok := m[v]; ok {
			r.count++
		} else {
			m[v] = &record{key: v, count: 1}
		}
	}

	records := make([]record, 0)
	for _, v := range m {
		records = append(records, *v)
	}

	sort.Sort(sort.Reverse(numberRecords(records)))
	
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, records[i].key)
	}
	
	return result
}

func (nr numberRecords) Len() int {
	return len(nr)
}

func (nr numberRecords) Swap(i, j int) {
	nr[i], nr[j] = nr[j], nr[i]
}

func (nr numberRecords) Less(i, j int) bool {
	return nr[i].count < nr[j].count
}

// @lc code=end


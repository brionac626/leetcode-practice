/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type LRUCache struct {
	Cap   int
	Len   int
	Start *Node
	End   *Node
	Map   map[int]*Node
}

type Node struct {
	Next *Node
	Prev *Node
	Key  int
	Val  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Cap: capacity, Map: make(map[int]*Node)}
}

func (this *LRUCache) Get(key int) int {
	foundNode, ok := this.Map[key]
	if !ok {
		return -1
	}

	if this.Start == foundNode {
		return foundNode.Val
	}

	if this.End == foundNode {
		this.End = foundNode.Prev
	}

	if foundNode.Next != nil {
		foundNode.Next.Prev = foundNode.Prev
	}

	if foundNode.Prev != nil {
		foundNode.Prev.Next = foundNode.Next
	}

	this.Start.Prev = foundNode
	foundNode.Next = this.Start
	this.Start = foundNode

	return foundNode.Val
}

func (this *LRUCache) Put(key int, value int) {
	if foundNode, ok := this.Map[key]; ok {
		foundNode.Val = value
		this.Get(foundNode.Key)
		return
	}

	n := &Node{Key: key, Val: value}

	if len(this.Map) == 0 {
		this.End = n
	} else {
		this.Start.Prev = n
		n.Next = this.Start
	}

	this.Map[key] = n
	this.Start = n

	if this.Cap == this.Len {
		delete(this.Map, this.End.Key)
		this.Len -= 1
		this.End = this.End.Prev
	}
	this.Len++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end


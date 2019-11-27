/*
 * @lc app=leetcode id=1108 lang=golang
 *
 * [1108] Defanging an IP Address
 */

// @lc code=start

func defangIPaddr(address string) string {
	var newIP []rune
	for _,v:=range address{
		if v == rune('.'){
			newIP = append(newIP,rune('['))
			newIP = append(newIP,rune('.'))
			newIP = append(newIP,rune(']'))
		} else {
			newIP = append(newIP,v)
		}
	}
	return string(newIP)
}

// @lc code=end



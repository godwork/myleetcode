package leetcode

import "fmt"

// 获取字符窜中，没有重复的字符窜最长长度
func LengthOfLongestSubStrings(s string) int {
	// if have no string, return zero
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	fmt.Printf("left:%d,right:%d \n", left, right)
	return result
}

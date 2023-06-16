package main

import (
	"fmt"

	"myleetcode/leetcode"
)

func main() {
	var s string = "abcaddcbfga"
	count := leetcode.LengthOfLongestSubStrings(s)
	fmt.Println("num4:", count)
}

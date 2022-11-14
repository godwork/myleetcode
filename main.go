package main

import (
	"fmt"
	"myleetcode/leetcode"
)

func main() {
	var num1 *leetcode.MyNode = &leetcode.MyNode{
		Val: 2,
		Next: &leetcode.MyNode{
			Val:  4,
			Next: &leetcode.MyNode{Val: 3},
		},
	}
	var num2 *leetcode.MyNode = &leetcode.MyNode{
		Val: 5,
		Next: &leetcode.MyNode{
			Val:  6,
			Next: &leetcode.MyNode{Val: 5},
		},
	}
	resNode := leetcode.AddTowNumberByNodeList(num1, num2)
	for resNode != nil {
		fmt.Println(*resNode)
		resNode = resNode.Next
	}
}

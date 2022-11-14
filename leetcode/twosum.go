package leetcode

// tow numbers add equals target
// param target is the result target
func TwoSum(nums []int, target int) (result []int) {
	var mapResult = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := mapResult[temp]; ok {
			return []int{mapResult[temp], i}
		}
		mapResult[nums[i]] = i
	}
	return nil
}

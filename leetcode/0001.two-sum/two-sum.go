package leetcode

func twoSum(nums []int, target int) []int {
	dataMap := make(map[int]int)

	for i, x := range nums {
		_, val := dataMap[x]
		if val {
			return []int{dataMap[x], i}
		} else {
			dataMap[target-x] = i
		}
	}

	return nil
}

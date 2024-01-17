package leetcode

func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)

	copy(ans[0:len(nums)], nums)
	copy(ans[len(nums):len(nums)*2], nums)

	return ans
}

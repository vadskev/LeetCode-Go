package leetcode

func Shuffle(nums []int, n int) []int {
	var fly, flx int = n, 0
	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			ans[i] = nums[flx]
			flx++
		} else {
			ans[i] = nums[fly]
			fly++
		}
	}
	return ans
}

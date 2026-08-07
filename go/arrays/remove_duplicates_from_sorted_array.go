package arrays

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3258/
func removeDuplicates(nums []int) int {
	i := 0
	for j := range nums {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

package arrays

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3157/
func moveZeroes(nums []int) {
	// [0] -> [0]
	// [1] -> [1]
	// [1, 2] -> [1, 2]
	// [0, 0] -> [0, 0]
	// [1, 0, 0, 5, 3, 0] -> [1, 5, 3, 0, 0, 0]

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if j < i {
				j = i
			}
			for ; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}

			if j == len(nums) {
				return
			}
		}
	}

}

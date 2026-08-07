package arrays

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/
func replaceElements(arr []int) []int {
	m := -1
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], m = m, max(m, arr[i])
	}

	return arr
}

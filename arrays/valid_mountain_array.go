package arrays

// https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3251/
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 1
	for ; i < len(arr) && arr[i] > arr[i-1]; i++ {
	}
	if i == 1 {
		return false
	}

	for ; i < len(arr) && arr[i] < arr[i-1]; i++ {
	}

	return i == len(arr) && arr[i-1] < arr[i-2]
}

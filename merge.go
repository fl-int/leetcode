package main

// https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0

	// [1,3,3,5,6,0,0,0] [2,4,7]
	// [1,2,3,5,6,0,0,0] [3,4,7]
	// [1,2,3,5,6,0,0,0] [3,4,7]
	// [0] [4]
	// [3] []
	// [0] [4]
	if nums1[i] > nums2[j] {
		nums1[i], nums2[j] = nums2[j], nums1[i]
		i++
	} else {
		i++
	}

	for i+j < m+n {

	}
}

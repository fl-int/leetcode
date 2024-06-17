package main

func removeElement(nums []int, val int) int {
	/*
		https://leetcode.com/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150
	*/
	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}

	return k
}

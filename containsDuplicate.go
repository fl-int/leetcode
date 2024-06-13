package main

type IntSet map[int]struct{}

func containsDuplicate(nums []int) bool {
	/*
		https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/

		Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
	*/
	u := make(IntSet)
	for _, v := range nums {
		_, ok := u[v]
		if ok {
			return true
		}
		u[v] = struct{}{}
	}

	return false
}

package main

func plusOne(digits []int) []int {
	/*
		You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
		The digits are ordered from most significant to least significant in left-to-right order.
		The large integer does not contain any leading 0's.

		Increment the large integer by one and return the resulting array of digits
	*/
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			for j := i + 1; j < len(digits); j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	newLen := len(digits) + 1
	oversized := make([]int, newLen)
	for i, _ := range oversized {
		oversized[i] = 0
	}
	oversized[0] = 1
	return oversized
}

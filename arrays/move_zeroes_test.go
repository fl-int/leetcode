package arrays

import (
	"slices"
	"testing"
)

type tmz struct {
	nums     []int
	expected []int
}

func TestMoveZeroes(t *testing.T) {
	tests := []tmz{
		{
			nums:     []int{0},
			expected: []int{0},
		},
		{
			nums:     []int{1},
			expected: []int{1},
		},
		{
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			nums:     []int{3, -1},
			expected: []int{3, -1},
		},
		{
			nums:     []int{1, 0, 0, 5, -1, 0},
			expected: []int{1, 5, -1, 0, 0, 0},
		},
	}

	for _, test := range tests {
		moveZeroes(test.nums)
		if !slices.Equal(test.expected, test.nums) {
			t.Fail()
		}
	}

	// [0] -> [0]
	// [1] -> [1]
	// [1, 2] -> [1, 2]
	// [0, 0] -> [0, 0]
	// [1, 0, 0, 5, 3, 0] -> [1, 5, 3, 0, 0, 0]
}

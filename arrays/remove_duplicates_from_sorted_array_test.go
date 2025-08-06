package arrays

import "testing"

type foo struct {
	nums    []int
	uniques int
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []foo{
		{
			nums:    []int{1},
			uniques: 1,
		},
		{
			nums:    []int{1, 1},
			uniques: 1,
		},
		{
			nums:    []int{1, 1, 2},
			uniques: 2,
		},
		{
			nums:    []int{1, 2, 5},
			uniques: 3,
		},
		{
			nums:    []int{1, 1, 2, 2, 2},
			uniques: 2,
		},
	}

	for _, test := range tests {
		actual := removeDuplicates(test.nums)
		if actual != test.uniques {
			t.Fatalf("array: %v, expected: %d, actual: %d", test.nums, test.uniques, actual)
		}
	}
}

package arrays

import (
	"slices"
	"testing"
)

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/
func TestReplaceElements(t *testing.T) {
	input := []int{17, 18, 5, 4, 6, 1}
	expect := []int{18, 6, 6, 6, 1, -1}
	actual := replaceElements(input)
	if len(actual) != 6 {
		t.Fail()
	}
	if !slices.Equal(input, expect) {
		t.Fail()
	}
}

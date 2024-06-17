package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	m := math.MaxInt

	for i, s := range strs {
		if len(s) < m {
			m = len(s)
		}
		if i == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			if s[j] != strs[0][j] {
				m = j
				continue
			}
		}
	}

	return strs[0][:m]
}

func testLongestCommonPrefix() {
	strs := []string{"flower", "flow", "flight", "fight", ""}
	fmt.Println(longestCommonPrefix(strs))
}

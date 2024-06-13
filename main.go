package main

import "fmt"

func main() {
	a := []int{5, 5, 7, 7, 8, 9}
	fmt.Println(a)
	k := removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(k)
}

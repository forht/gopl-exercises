package main

import (
	"fmt"
)

func rotate(slice []int, shift int) []int {
	if len(slice) == 0 {
		return slice
	}
	shift = shift % len(slice)
	return append(slice[shift:], slice[:shift]...)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	s = rotate(s, 7)
	fmt.Println("%d", s)
}

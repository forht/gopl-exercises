package main

import (
	"fmt"
)

func sort(bs []rune) {
	for i := 1; i < len(bs); i++ {
		j := i - 1
		for j >= 0 && bs[j] > bs[j+1] {
			bs[j], bs[j+1] = bs[j+1], bs[j]
			j--
		}
	}
}

func anagram(s1 string, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	n1 := len(r1)
	n2 := len(r2)
	if n1 != n2 {
		return false
	}
	sort(r1)
	sort(r2)
	for i := 0; i < n1; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "abc"
	s2 := "acb"
	fmt.Println(anagram(s1, s2))
}

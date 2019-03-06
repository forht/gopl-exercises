package main

import (
	"fmt"
)

func deldup(s []string) []string {
	i := 0
	j := 0
	for i < len(s) {
		s[j] = s[i]
		i++
		j++
		for i < len(s) && s[i] == s[i-1] {
			i++
		}
	}
	return s[:j]
}

func main() {
	s := []string{"b", "a", "a", "a"}
	s = deldup(s)
	fmt.Println("%q", s)
}

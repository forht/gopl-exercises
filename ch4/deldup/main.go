package main

import (
	"fmt"
)

func deldup(s []string) []string {
	i := 0
	j := 0
	for {
		flag := false
		for i < len(s)-1 && s[i] == s[i+1] {
			i++
			flag = true
		}
		if flag {
			i++
		}
		if i >= len(s) {
			break
		}
		s[j] = s[i]
		i++
		j++
	}
	return s[:j]
}

func main() {
	s := []string{"a", "b", "b", "c", "d", "d", "d", "e"}
	s = deldup(s)
	fmt.Println("%q", s)
}

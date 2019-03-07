package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseutf8(s []byte) {
	for i := 0; i < len(s); {
		_, l := utf8.DecodeRune(s[i:])
		if l > 1 {
			reverse(s[i : i+l])
		}
		i += l
	}
	reverse(s)
}

func main() {
	s := []byte("ßłĸðđ»¢“·")
	reverseutf8(s)
	fmt.Println(string(s))
}

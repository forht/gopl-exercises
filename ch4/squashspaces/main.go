package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(s []byte) []byte {
	i := 0
	j := 0
	for i < len(s) {
		r, l := utf8.DecodeRune(s[i:])
		if r != utf8.RuneError {
			if !unicode.IsSpace(r) {
				copy(s[j:], s[i:i+l])
				j += l
			} else {
				s[j] = ' '
				j += 1
			}
		}
		i += l
		if unicode.IsSpace(r) {
			for i < len(s) {
				r, l = utf8.DecodeRune(s[i:])
				if !unicode.IsSpace(r) {
					break
				}
				i += l
			}
		}
	}
	return s[:j]
}

func main() {
	s := []byte("nel\r \n   mezzo \ndel cammin   \r    di")
	s = squashSpaces(s)
	fmt.Printf("%q", string(s))
}

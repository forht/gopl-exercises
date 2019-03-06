package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func sha256diff(h1 [32]uint8, h2 [32]uint8) int {
	sum := 0
	for i := 0; i < 32; i++ {
		sum += int(pc[h1[i]^h2[i]])
	}
	return sum
}

func main() {
	a := [32]uint8{15}
	b := [32]uint8{0}
	fmt.Println(sha256diff(a, b))
}

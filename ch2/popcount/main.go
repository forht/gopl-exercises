// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var sum byte
	var i byte
	for i = 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

func PopCountShift(x uint64) int {
	var sum uint64
	var i uint
	for i = 0; i < 64; i++ {
		sum += (x >> i) & 1
	}
	return int(sum)
}

func PopCountClear(x uint64) int {
	sum := 0
	for x != 0 {
		x &= x - 1
		sum++
	}
	return int(sum)
}

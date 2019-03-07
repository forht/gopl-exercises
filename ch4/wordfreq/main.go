// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	count := make(map[string]int)
	for in.Scan() {
		w := in.Text()
		count[w]++
	}
	if err := in.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("word\t\tfreq\n")
	for w, n := range count {
		fmt.Printf("%q\t\t%d\n", w, n)
	}
}

//!-

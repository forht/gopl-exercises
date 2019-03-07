// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type runePredicate func(rune) bool

func main() {
	runePredicates := map[string]runePredicate{
		"Control": unicode.IsControl,
		"Digit":   unicode.IsDigit,
		"Graphic": unicode.IsGraphic,
		"Letter":  unicode.IsLetter,
		"Lower":   unicode.IsLower,
		"Mark":    unicode.IsMark,
		"Number":  unicode.IsNumber,
		"Print":   unicode.IsPrint,
		"Punct":   unicode.IsPunct,
		"Space":   unicode.IsSpace,
		"Symbol":  unicode.IsSymbol,
		"Title":   unicode.IsTitle,
		"Upper":   unicode.IsUpper,
	}
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		for cat, fn := range runePredicates {
			if fn(r) {
				counts[cat]++
			}
		}
	}
	fmt.Printf("category\tcount\n")
	for c := range runePredicates {
		fmt.Printf("%v\t\t%d\n", c, counts[c])
	}
}

//!-

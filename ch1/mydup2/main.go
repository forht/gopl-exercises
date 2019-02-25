// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// Exercise 1.4: Modify dup2 to print the names of all files in which each
// duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	infiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, infiles, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, infiles, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, infiles[line])
		}
	}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int,
	infiles map[string][]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if fs, found := infiles[line]; !found || !contains(fs, filename) {
			infiles[line] = append(infiles[line], filename)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

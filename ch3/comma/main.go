package main

import "fmt"
import "strings"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaf(s string) string {
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		return comma(s[:dot-1]) + "." + s[dot+1:]
	} else {
		return comma(s)
	}
}

func main() {
	fmt.Println(commaf("23893248.93423843233892"))
}

package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	r := n % 3
	if r != 0 {
		buf.WriteString(s[:r])
		buf.WriteByte(',')
	}
	i := r
	for ; i < (n - 3); i += 3 {
		buf.WriteString(s[i : i+3])
		buf.WriteByte(',')
	}
	buf.WriteString(s[i : i+3])
	return buf.String()
}

func main() {
	fmt.Println(comma("2374"))
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Fprintf(os.Stderr, "usage: github get|create user repository [id]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "get":
		if len(os.Args[1:]) != 4 {
			fmt.Fprint(os.Stderr, "usage: github get user repository id")
			os.Exit(1)
		}
		var id int
		var err error
		if id, err = strconv.Atoi(os.Args[4]); err != nil {
			fmt.Fprint(os.Stderr, "usage: github get user repository id")
			os.Exit(1)
		}
		if err := getIssue(os.Args[2], os.Args[3], id); err != nil {
			fmt.Fprintf(os.Stderr, "github: %v", err)
			os.Exit(1)
		}
	case "create":
		if len(os.Args[1:]) != 3 {
			fmt.Fprint(os.Stderr, "usage: github create user repository")
			os.Exit(1)
		}
		if err := createIssue(os.Args[2], os.Args[3]); err != nil {
			fmt.Fprintf(os.Stderr, "github: %v", err)
			os.Exit(1)
		}
	default:
		fmt.Fprint(os.Stderr, "usage: github get|create user repository [id]")
		os.Exit(1)
	}
}

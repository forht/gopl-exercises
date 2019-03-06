package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var hashFunc = flag.String("f", "sha256", "SHA-2 specific function")

func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	s := ""
	for input.Scan() {
		s += input.Text()
	}
	switch *hashFunc {
	case "sha256":
		sum := sha256.Sum256([]byte(s))
		fmt.Printf("%x  -\n", sum)
	case "sha384":
		sum := sha512.Sum384([]byte(s))
		fmt.Printf("%x  -\n", sum)
	case "sha512":
		sum := sha512.Sum512([]byte(s))
		fmt.Printf("%x  -\n", sum)
	default:
		log.Fatal("Usage: sumstdin -f hashfunc")
	}

}

package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	s := flag.Args()
	fmt.Println(s)
}

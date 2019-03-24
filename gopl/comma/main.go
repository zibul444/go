// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for i := 1; i < len(os.Args); i++ {
			fmt.Printf("  %s\n", comma(os.Args[i]))
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			fmt.Println(comma(input.Text()))
		}
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBuf(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	for range s {

	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-

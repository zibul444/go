package main

import "fmt"

func main() {
	for i := 32; i < 127; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

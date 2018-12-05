package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fizzBuzz(i)
	}
}
func fizzBuzz(i int) {
	if i%3 == 0 {
		fmt.Print("Fizz")
	}
	if i%5 == 0 {
		fmt.Print("Buzz")
	} else if i%3 != 0 {
		fmt.Print(i)
	}
	fmt.Print("\n")
}

func fizzBuzz_butiful(i int) {
	switch {
	case i%15 == 0:
		fmt.Println("FizzBuzz")
	case i%3 == 0:
		fmt.Println("Fizz")
	case i%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(i)
	}
}

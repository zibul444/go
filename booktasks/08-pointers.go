package main

import "fmt"

func main() {
	testSwap()

	//x := 1.5
	//square(&x)
}

func square(x *float64) {
	*x = *x * *x
	fmt.Println()
	fmt.Println(*x)
}

func testSwap() {
	x, y := 1, 2
	fmt.Printf("x=%d, y=%d", x, y)
	fmt.Println()
	swap(&x, &y)
	fmt.Printf("x=%d, y=%d", x, y)
}

func swap(x, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
	//элегантно:
	//*x, *y = *y, *x
}

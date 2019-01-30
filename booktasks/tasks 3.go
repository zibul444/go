package main

import "fmt"

func main() {

	mass := []int{
		1,
		2,
		3,
		4,
		5,
	}
	//mass []int := []int(1,2,3,4,5,)

	getElements(mass)

	fmt.Println("\n")

	fmt.Print(make([]int, 3, 9))

	fmt.Println("\n")

	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Print(x[2:5])
	fmt.Println()
	fmt.Print(x[2:])
	fmt.Println()
	fmt.Print(x[:5])

	fmt.Println("\n")

	fmt.Println(getLitel())

	fmt.Println()

	fmt.Println("\n")

}

func getElements(mass []int) {
	fmt.Print(mass[3])
}

func getLitel() (litel int) {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := x[0]
	for _, i := range x {
		if min > i {
			min = i
		}

	}
	return min

}

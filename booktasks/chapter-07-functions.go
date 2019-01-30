package main

import (
	"fmt"
)

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	//
	//fmt.Println(sum([]int{1,2,3,4}))
	//
	//fmt.Println(half(1))
	//
	//fmt.Println(max([]int{-5,1,2,3,4}))
	//
	//
	//nextOdd := makeOddGenerator()
	//fmt.Println(nextOdd()) // 1
	//fmt.Println(nextOdd()) // 3
	//fmt.Println(nextOdd()) // 5

	fmt.Println(fib(1))

	panic("PANIC")
}

func sum(i []int) (sum int) {
	//sum = 0
	for _, x := range i {
		sum += x
	}

	return sum
}

func half(i int) (int, bool) {
	status := false
	result := i

	if (result % 2) == 0.0 {
		status = true
	}

	return result, status
}

func max(i []int) (max int) {

	max = i[0]

	for _, x := range i {
		if max < x {
			max = x
		}
	}

	return max
}

func makeOddGenerator() func() int {
	i := -1
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func fib(n int64) (f int64) {
	f = fib(n+1) + fib(n+2)
	return f
}

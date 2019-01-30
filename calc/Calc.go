package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go calc()
	wg.Wait()
}

func calc() {
	defer wg.Done()

	var f, s, result int
	var sign string
	fmt.Println("Input first value")
	fmt.Scanln(&f)
	fmt.Println("Input sign")
	fmt.Scanln(&sign)
	fmt.Println("Input second value")
	fmt.Scanln(&s)
	fmt.Println("Result is:")

	if sign == "+" {
		result = f + s
	} else if sign == "-" {
		result = f - s
	} else if sign == "*" {
		result = f * s
	} else if sign == "/" {
		result = f / s
	}

	fmt.Println(result)

	//fmt.Println(<-c)
	time.Sleep(time.Duration(time.Second + 5))
}

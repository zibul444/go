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
	fmt.Scanln(&f)
	fmt.Scanln(&sign)
	fmt.Scanln(&s)

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

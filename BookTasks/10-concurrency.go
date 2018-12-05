package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//testF()
	testPingerPrinter()

}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(205))
		time.Sleep(time.Millisecond * amt)
	}
	fmt.Println()
}

func testF() {
	for i := 0; i < 10; i++ {
		go f(0)
	}
	var input string
	fmt.Scanln(&input)
}


func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c <-chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func testPingerPrinter() {
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func ponger(c chan<- string) {
	for{
		c <- "pong"
	}
}
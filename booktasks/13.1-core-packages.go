package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan string)
var cancel = make(chan string)
var m = new(sync.Mutex)
var C = 5
var wg sync.WaitGroup

const Q = 5

func main() {

	go mu()
	go print()

	fmt.Println(<-cancel)

	//var input string
	//fmt.Scanln(&input)
}

func mu() {
	for i := 0; i < Q; i++ {
		wg.Add(1)
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			ch <- fmt.Sprint(i, " Start")
			ch <- fmt.Sprint(i, " end")
			time.Sleep(time.Millisecond * 50)
			wg.Done()
		}(i)
	}

	//time.Sleep(time.Second * 2)
	wg.Wait()
	cancel <- "for - is ended"
}

func print() {
	for {
		fmt.Println(<-ch)
	}
}

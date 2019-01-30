package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu()

	var input string
	fmt.Scanln(&input)
}

func mu() {
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "Start")
			time.Sleep(time.Second)
			fmt.Println(i, "emd")
			m.Unlock()
		}(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("for - is end")
}

package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for {
		time.Sleep(1 * time.Second)
		fmt.Print("Hello world! ", i, "\n")
		i++
	}

	//grab()

}

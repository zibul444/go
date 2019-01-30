package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.NewSource(time.Now().UnixNano()))
}

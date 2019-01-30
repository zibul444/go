package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	g := rand.Intn(90)
	fmt.Println(g)
	s := Generator(g)
	fmt.Println(s)
}

func Generator(l int) string {
	var result string = ""
	if l == 0 {
		return result
	}

	for i := 0; i < l; i++ {
		r := rand.Intn(len(cha))
		c := cha[r]
		result += string(c)
	}

	return result
}

var cha = []string{
	"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "[", "a", "s", "d",
	"f", "g", "h", "j", "k", "l", ";", "'", "z", "x", "c", "v", "b", "n",
	"m", ",", ".", "/", "'", "]", "{", "}", ":", "\"", "<", ">", "?", "~",
	"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "_", "+", "`", "1",
	"2", "3", "4", "5", "6", "7", "8", "9", "0", "-", "=", "`", "|"}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(rand.NewSource(time.Now().UnixNano()))
	//a,s := 1, 2
	//fmt.Println(incr(&a))
	//fmt.Println(incr2(&s))
	//incr2(&s)
	//fmt.Println(a,s)

	//t := Properties{}
	//s := reflect.ValueOf(&t).Elem()
	//typeOfT := s.Type()
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}

	s := "5"

	var i int

	//i = strconv.ParseInt(s, int, 16)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(i)

}

func incr(i *int) int {
	*i++
	return *i
}
func incr2(i *int) {
	*i++
}

type T struct {
	A int
	B string
}

type Properties struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

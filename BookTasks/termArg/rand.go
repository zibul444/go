package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd1()
}


func rnd1() {
	// Определение флагов
	maxp := flag.Int("max", 6, "the max value")
	// Парсинг
	flag.Parse()
	// очень важная деталь! Seed - принимает 
	rand.Seed(time.Now().UnixNano())

	// Генерация числа от 0 до max
	fmt.Println("Int63						",			rand.Int63())
	fmt.Println("Uint32						",			rand.Uint32())
	fmt.Println("Uint64						",			rand.Uint64())
	fmt.Println("Int31						",			rand.Int31())
	fmt.Println("Int							",			rand.Int())
	fmt.Println("Int63n до 	",maxp,"	", 				rand.Int63n(int64(*maxp)))
	fmt.Println("Int31n до 	",maxp,"	",				rand.Int31n(int32(*maxp)))
	fmt.Println("Intn до 	",maxp,"	", 				rand.Intn(*maxp))
	fmt.Println("Float64						",			rand.Float64())
	fmt.Println("Float32						",			rand.Float32())


	fmt.Println("")
	fmt.Println("", rand.Int())

}

func rnd2()  {

}
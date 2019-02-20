package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
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
	// очень важная деталь! Seed - принимает число
	rand.Seed(time.Now().UnixNano())

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 8)
	defer w.Flush()
	show := func(name string, v ...interface{}) {

		//fmt.Fprintf(w, "%s\t", name)
		//for n := range v{
		//	fmt.Fprintf(w, "%v\t", n)
		//}
		//fmt.Fprintln(w)

		fmt.Fprintf(w, "%s\t%v\t\n", name, v)
	}

	show("Int63 ", rand.Int63())
	show("Uint32 ", rand.Uint32())
	show("Uint64 ", rand.Uint64())
	show("Int31 ", rand.Int31())
	show("Int ", rand.Int())
	show("Int63n до ", *maxp, rand.Int63n(int64(*maxp)))
	show("Int31n до ", *maxp, rand.Int31n(int32(*maxp)))
	show("Intn до ", *maxp, rand.Intn(*maxp))
	show("Float64 ", rand.Float64())
	show("Float32 ", rand.Float32())

	show("__________________")
}

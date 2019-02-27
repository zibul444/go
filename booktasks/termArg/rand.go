package main

import (
	"flag"
	"go/logger"
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
	// очень важная деталь! Seed - принимает число
	rand.Seed(time.Now().UnixNano())

	//w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 8)
	//defer w.Flush()
	//show := func(name string, v ...interface{}) {
	//
	//	fmt.Fprintf(w, "%s\t%d\t", name, v)
	//	for n := range v{
	//		fmt.Fprintf(w, "%19d\t", n)
	//		fmt.Fprint(w,  n)
	//	}
	//	fmt.Fprintln(w)
	//
	//	//fmt.Fprintf(w, "%s\t%v\t\n", name, v)
	//}

	log := logger.MyLogger

	log.Debugf("%10s %22v", "Int63 ", rand.Int63())
	log.Debugf("%10s %22v", "Uint32 ", rand.Uint32())
	log.Debugf("%10s %22v", "Uint64 ", rand.Uint64())
	log.Debugf("%10s %22v", "Int31 ", rand.Int31())
	log.Debugf("%10s %22v", "Int ", rand.Int())
	log.Debugf("%10s %2d %19d", "Int63n до ", *maxp, rand.Int63n(int64(*maxp)))
	log.Debugf("%10s %2d %19d", "Int31n до ", *maxp, rand.Int31n(int32(*maxp)))
	log.Debugf("%10s %2d %19d", "Intn до ", *maxp, rand.Intn(*maxp))
	log.Debugf("%10s %22v", "Float64 ", rand.Float64())
	log.Debugf("%10s %22v", "Float32 ", rand.Float32())

	//show("__________________")
}

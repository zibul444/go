package main

import (
	"github.com/op/go-logging"
	"go/logger"
	_ "go/logger"
	"math"
)

func main() {
	log := logger.MyLogger

	//log.Debugf("debug %s  ______________________________________________________", logger.Password("secret") )
	//log.Info("info _______________________________________________________________")
	//log.Notice("notice _____________________________________________________________")
	//log.Warning("warning ____________________________________________________________")
	//log.Error("err ________________________________________________________________")
	//log.Critical("crit _______________________________________________________________")

	var u uint8 = 255
	log.Debug(u, u+1, u, u*u)
	var i int8 = 127
	log.Debug(i, i+1, i, i*i)

	var x, y = 2, 2
	log.Debug(x < y)

	log.Debug(false != true)

	var q, w, e, r = false, false, true, true
	log.Debug(q, w, e, r)

	//bitwiseOperation(log)

	log.Debug(math.MaxInt8)

}

func bitwiseOperation(log *logging.Logger) {
	log.Debug(2 * 2)
	log.Debug("__&_")
	log.Debug(2 & 10)
	log.Debug(4 & 2)
	log.Debug(2 & 4)
	log.Debug("__|_")
	log.Debug(2 | 2)
	log.Debug(4 | 2)
	log.Debug(2 | 4)
	log.Debug("__^_")
	log.Debug(2 ^ 2)
	log.Debug(4 ^ 2)
	log.Debug(2 ^ 4)
	log.Debug("__&^_")
	log.Debug(2 &^ 2)
	log.Debug(4 &^ 2)
	log.Debug(2 &^ 4)
	log.Debug("__<<_")
	log.Debug(2 << 2)
	log.Debug(4 << 2)
	log.Debug(2 << 4)
	log.Debug("__>>_")
	log.Debug(2 >> 2)
	log.Debug(4 >> 2)
	log.Debug(2 >> 4)
}

/*
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
}*/

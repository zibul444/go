package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.Println(runtime.GOMAXPROCS(4))
	compact()
	log.Println(runtime.GOMAXPROCS(0))

}

func compact() {
	for {
		select {
		case <-time.After(time.Minute):
			// do compaction
			//case <- stop:
			return
		}
	}
}

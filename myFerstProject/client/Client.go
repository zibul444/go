package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func Client(name string) {
	c, err := net.Dial("tcp", "127.0.0.2:8800")
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		var msg string
		fmt.Scanln(&msg)
		//<- time.After(time.Second * time.Duration(100))

		log.Println(name, "sending: ", msg)
		err = gob.NewEncoder(c).Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer c.Close()
}

package main

import (
	"encoding/gob"
	"log"
	"net"
)

func Server() {
	ln, err := net.Listen("tcp", "127.0.0.2:8800")
	if err != nil {
		log.Fatal(err)
		return
	}
	//for {
	c, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
		//continue
		return
	}
	go handleServerConnection(c)
	//}
}

func handleServerConnection(c net.Conn) {
	for {
		var msg string
		err := gob.NewDecoder(c).Decode(&msg)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Resived ", msg)
		}
	}

	defer c.Close()
}

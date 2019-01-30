package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	go server()
	go client()
	var input string
	fmt.Scanln(input)
	<-time.After(time.Second * time.Duration(100))
}

func server() {
	ln, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Resived ", msg)
	}
	defer c.Close()
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		log.Fatal(err)
		return
	}

	msg := "Hello World"
	log.Println("Sending ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
}

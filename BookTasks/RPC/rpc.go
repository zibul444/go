package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	log.Println("Negate")
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		log.Fatal(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Fatal("Server.Negate(999) =", result)
	}
}
func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

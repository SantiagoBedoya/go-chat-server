package main

import (
	"log"
	"net"
)

var clients = make(map[string]net.Conn)
var leaving = make(chan Message)
var messages = make(chan Message)

func main() {
	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	go Broadcaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go Handler(conn)
	}
}

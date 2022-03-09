package main

import "net"

func NewMessage(msg string, conn net.Conn) Message {
	addr := conn.RemoteAddr().String()
	return Message{
		Text:    addr + msg,
		Address: addr,
	}
}

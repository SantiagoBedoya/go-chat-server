package main

import (
	"bufio"
	"net"
)

func Handler(conn net.Conn) {
	clients[conn.RemoteAddr().String()] = conn

	messages <- NewMessage(" joined.", conn)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- NewMessage(": "+input.Text(), conn)
	}

	delete(clients, conn.RemoteAddr().String())
	leaving <- NewMessage(" has left.", conn)
	conn.Close()
}

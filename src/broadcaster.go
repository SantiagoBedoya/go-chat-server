package main

import "fmt"

func Broadcaster() {
	for {
		select {
		case msg := <-messages:
			for _, conn := range clients {
				if msg.Address == conn.RemoteAddr().String() {
					continue
				}
				fmt.Fprintln(conn, msg.Text)
			}
		case msg := <-leaving:
			for _, conn := range clients {
				fmt.Println(conn, msg.Text)
			}
		}
	}
}

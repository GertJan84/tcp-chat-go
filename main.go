package main

import (
	"log"
	"net"
)

const (
	IP       = "127.0.0.1"
	PORT     = "8888"
	PROTOCOL = "tcp"
)

func handle_error(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

func handle_connection(con net.Conn) {
	log.Println(con.RemoteAddr().String(), "connected to server")

}

func main() {
	listener, err := net.Listen(PROTOCOL, IP+":"+PORT)

	handle_error(err, "Can't create listener")

	log.Println("Start server on " + IP + ":" + PORT)

	defer func(listener net.Listener) {
		err := listener.Close()
		handle_error(err, "Can't close listener")
	}(listener)

	for {
		con, err := listener.Accept()

		handle_error(err, "Can't accept connection")
		go handle_connection(con)
	}

}

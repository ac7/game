package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ac7/game/server"
)

func main() {
	port := 1030
	id := 0
	log.Printf("starting server on localhost:%d\n", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Warning, could not accept connection: ", err)
			continue
		}
		id = id + 1

		go func(id int) {
			err := server.HandleConn(conn, id)
			if err != nil {
				log.Println("Connection", id, "terminated with error: ", err)
			} else {
				log.Println("Connection", id, "terminated sucessfully.")
			}
		}(id)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/ac7/game/server"
)

func main() {
	testConn := flag.Bool("test_conn_and_quit", false, "Accept only one connection and then terminate.")
	flag.Parse()

	waitGroup := new(sync.WaitGroup)

	port := 1030
	id := 0
	log.Printf("Starting server on localhost:%d\n", port)
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Warning, could not accept connection: ", err)
			continue
		}
		id++

		waitGroup.Add(1)
		go func(id int) {
			err := server.HandleConn(conn, id)
			if err != nil {
				log.Println("Connection", id, "terminated with error: ", err)
			} else {
				log.Println("Connection", id, "terminated sucessfully.")
			}

			err = conn.Close()
			if err != nil {
				log.Printf("Warning, error closing connection %d: %s\n", id, err.Error())
			}
			waitGroup.Done()
		}(id)

		if *testConn {
			break
		}
	}

	log.Println("Waiting on connections to close...")
	waitGroup.Wait()
}

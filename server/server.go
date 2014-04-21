package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type server struct {
	IWorld
	toClient  chan IUnit
	toServer  chan IUnit
	waitGroup *sync.WaitGroup
}

func (s *server) Listen(location string) error {
	ln, err := net.Listen("tcp", location)
	id := 0

	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Warning, could not accept connection: ", err)
			continue
		}

		id++
		s.waitGroup.Add(1)

		go func(id int) {
			err := s.HandleConn(conn, id)

			if err != nil {
				log.Println("Connection", id, "terminated with error: ", err)
			} else {
				log.Println("Connection", id, "terminated sucessfully.")
			}
		}(id)
	}
}

func (s *server) HandleConn(conn io.ReadWriteCloser, id int) error {
	if conn == nil {
		return fmt.Errorf("Invalid connection!")
	}
	s.waitGroup.Add(1)

	// first we check to make sure the client returns the HANDSHAKE_CLIENT string
	var helloMessage []byte = make([]byte, len(HANDSHAKE_CLIENT))
	_, err := conn.Read(helloMessage)
	if err != nil {
		return fmt.Errorf("Could not read from connection: %s", err.Error())
	}
	if !bytes.Equal(helloMessage, []byte(HANDSHAKE_CLIENT)) {
		conn.Close()
		return fmt.Errorf(`Recieved non-protocol handshake "%s" (expected "%s"), abort.`,
			string(helloMessage), HANDSHAKE_CLIENT)
	}
	log.Printf("Connection %d sent correct handshake \"%s\"\n", id, HANDSHAKE_CLIENT)

	// then we send the HANDSHAKE_SERVER string to the client
	n, err := conn.Write([]byte(HANDSHAKE_SERVER))
	if err != nil {
		return fmt.Errorf(`Cannot write to connection %d: "%s"`, id, err.Error())
	} else if n != len(HANDSHAKE_SERVER) {
		return fmt.Errorf("Only wrote %d bytes to connection (expected %d).", n, len(HANDSHAKE_SERVER))
	} else {
		log.Printf(`Sent return handshake "%s" to connection number %d`, HANDSHAKE_SERVER, id)
	}

	s.waitGroup.Done()
	return conn.Close()
}

func NewServer() IServer {
	return &server{
		toClient:  make(chan IUnit, 16),
		toServer:  make(chan IUnit, 16),
		IWorld:    newWorld(),
		waitGroup: new(sync.WaitGroup),
	}
}

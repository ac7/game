package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/vmihailenco/msgpack"
)

const CHAN_BUFFERSIZE = 32
const MAX_PACKET_SIZE = 2048

type server struct {
	IWorld
	toClient  []chan IUnit
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
		return fmt.Errorf("Invalid connection %d!", id)
	}
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()

	// first we check to make sure the client returns the HANDSHAKE_CLIENT string
	helloMessage := make([]byte, len(HANDSHAKE_CLIENT))
	_, err := conn.Read(helloMessage)
	if err != nil {
		return fmt.Errorf("Could not read from connection: %s", err.Error())
	}
	if !bytes.Equal(helloMessage, []byte(HANDSHAKE_CLIENT)) {
		conn.Close()
		return fmt.Errorf(`Recieved non-protocol handshake "%s" (expected "%s"), abort.`,
			string(helloMessage), HANDSHAKE_CLIENT)
	}

	// then we send the HANDSHAKE_SERVER string to the client
	n, err := conn.Write([]byte(HANDSHAKE_SERVER))
	if err != nil {
		return fmt.Errorf(`Cannot write to connection %d: "%s"`, id, err.Error())
	} else if n != len(HANDSHAKE_SERVER) {
		return fmt.Errorf("Only wrote %d bytes to connection (expected %d).", n, len(HANDSHAKE_SERVER))
	}
	log.Printf("Accepted connection %d (handshake verified)", id)

	packet := make([]byte, MAX_PACKET_SIZE)
	table := make(map[string]interface{})
	for {
		_, err := conn.Read(packet)
		if err != nil {
			return fmt.Errorf(`Could not read from connection %d: "%s"`, id, err.Error())
		}

		err = msgpack.Unmarshal(packet, &table)
		if err != nil {
			return fmt.Errorf("Can't unmarshal packet: '%s'", err.Error())
		}
		if typeName, ok := table["type"]; ok {
			switch typeName {
			case "unit":
				if id, ok := table["id"].(int64); ok {
					u := s.Unit(id)
					if u == nil {
						u := new(unit)
						s.AddUnit(u)
					}
					u.Deserialize(table)
					log.Println("Decoded unit: %+v", s.Unit(id))
				}
			case "quit":
				return conn.Close()
			}
		} else {
			log.Printf("Warning: dropping a bad packet on %d", id)
		}
	}
}

func New() IServer {
	return &server{
		toClient:  make([]chan IUnit, 0),
		toServer:  make(chan IUnit, CHAN_BUFFERSIZE),
		IWorld:    newWorld(),
		waitGroup: new(sync.WaitGroup),
	}
}

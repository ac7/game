package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

func HandleConn(conn net.Conn, id int) error {
	if conn == nil {
		return fmt.Errorf("Invalid connection!")
	}

	var helloMessage []byte = make([]byte, 9)
	_, err := conn.Read(helloMessage)
	if err != nil {
		return fmt.Errorf("Could not read from connection: %s", err.Error())
	}
	if !bytes.Equal(helloMessage, []byte("handshake")) {
		return fmt.Errorf(`Recieved non-protocol handshake "%s", abort (expected "%s").`, string(helloMessage), "handshake")
	}

	log.Println("Connection", id, "send correct handshake.")

	return conn.Close()
}

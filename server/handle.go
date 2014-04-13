package server

import (
	"bytes"
	"fmt"
	"net"
)

func HandleConn(conn net.Conn, id int) error {
	if conn == nil {
		return fmt.Errorf("Invalid connection!")
	}

	var helloMessage []byte = make([]byte, 9)
	_, err := conn.Read(helloMessage)
	if err != nil {
		fmt.Println("Could not read from connection:", err)
		return err
	}
	if !bytes.Equal(helloMessage, []byte("handshake")) {
		return fmt.Errorf(`Recieved non-protocol handshake "%s", abort (expected "%s").`, string(helloMessage), "handshake")
	}

	return conn.Close()
}

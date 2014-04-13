package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

const HANDSHAKE_CLIENT = "handshake"
const HANDSHAKE_SERVER = "handshake_part_two"

func HandleConn(conn net.Conn, id int) error {
	if conn == nil {
		return fmt.Errorf("Invalid connection!")
	}

	var helloMessage []byte = make([]byte, 9)
	_, err := conn.Read(helloMessage)
	if err != nil {
		return fmt.Errorf("Could not read from connection: %s", err.Error())
	}
	if !bytes.Equal(helloMessage, []byte(HANDSHAKE_CLIENT)) {
		return fmt.Errorf(`Recieved non-protocol handshake "%s" (expected "%s"), abort.`,
			string(helloMessage), HANDSHAKE_CLIENT)
	}

	log.Println("Connection", id, "sent correct handshake.")

	n, err := conn.Write([]byte(HANDSHAKE_SERVER))
	if err != nil {
		return fmt.Errorf("Cannot write to connection %d: '%s'", id, err.Error())
	} else if n != len(HANDSHAKE_SERVER) {
		return fmt.Errorf("Only wrote %d bytes to connection (expected %d).", n, len(HANDSHAKE_SERVER))
	} else {
		log.Printf(`Sent return handshake "%s" to connection number %d`, HANDSHAKE_SERVER, id)
	}

	return conn.Close()
}

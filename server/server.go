package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

// goroutine to handle the TCP connection
func handleConnection(conn net.Conn, prefix string) {
	// ensures connection closes on termination
	defer conn.Close()

	fmt.Printf("serving %s\n", conn.RemoteAddr().String())
	// log.Info().Str("remote_address", conn.RemoteAddr().String()).Msg("serving")

	// buffered reading client data
	reader := bufio.NewReader(conn)

	for {
		// handle data in bytes
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF { // check end of stream
				log.Error().Err(err).Msg("failed to read data")
			}
			return
		}

		fmt.Printf("request: %s", bytes)

		// prepend prefix and send as response
		line := fmt.Sprintf("%s %s", prefix, bytes)
		fmt.Printf("response: %s", line)
		conn.Write([]byte(line))
	}
}

func main() {
	// check for port and message (prefix)
	if len(os.Args) == 1 {
		log.Error().Msg("please provide a port number and message")
		return
	}

	port := ":" + os.Args[1]
	prefix := os.Args[2]

	// create server and listen for incoming connections at port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Error().Err(err).Msg("failed to create server")
		return
	}

	fmt.Printf("listening on %s, prefix: %s\n", listener.Addr(), prefix)

	// ensures listener closes on termination
	defer listener.Close()

	// accept and handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error().Err(err).Msg("failed to accept connection")
			return
		}
		// handle connections in new goroutine
		go handleConnection(conn, prefix)
	}
}

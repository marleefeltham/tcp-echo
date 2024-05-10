package main

import (
	"net"
	"os"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	prefix := "hello"
	os.Args = []string{"server", "8080", prefix}

	// test server initialization
	go main()

	time.Sleep(1 * time.Second)

	// test connection establishment
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		t.Fatalf("failed to dial")
	}
	defer conn.Close()

	// test data write
	request := "world\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		t.Fatalf("failed to write to server")
	}

	// test response handling
	response := make([]byte, 0)
	buffer := make([]byte, 1024)
	want := prefix + " " + request
	for {
		// test data read
		n, err := conn.Read(buffer)
		if err != nil {
			t.Fatalf("failed to read")
		}
		response = append(response, buffer[:n]...)
		if buffer[n-1] == '\n' {
			break
		}
	}

	// test server response
	if string(response) != want {
		t.Errorf("handleConnection did not respond as expected. got: %s, expected: %s", string(response), want)
	}
}

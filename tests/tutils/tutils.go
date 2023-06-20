package tutils

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	maxRetries        = 15
	secondsForTimeout = 5
)

// StartNewTestServer starts a http server in a random open port at localhost and returns the address used.
// Will panic if server isn't up within 15 retries with different addresses
func StartNewTestServer(server *http.Server) string {
	var addr string
	var err error
	for i := 1; i <= maxRetries; i++ {
		var listener net.Listener
		listener, err = listenOpenPort()
		if err != nil {
			panic(err)
		}
		go startServer(server, listener)

		addr = listener.Addr().String()
		err = dialServer(addr)
		if err != nil {
			log.Printf("Connection refused on address %s, try %d/%d", server.Addr, i, maxRetries)
		}
	}
	if err != nil {
		panic(err)
	}
	return addr
}

func listenOpenPort() (net.Listener, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("failed to find open port - error: %v", err)
	}
	log.Printf("Listening to open port on address: %s", listener.Addr())
	return listener, nil
}

func startServer(server *http.Server, listener net.Listener) {
	err := server.Serve(listener)
	if err != nil && err != http.ErrServerClosed {
		log.Panicf("failed to start test server - error: %v", err)
	}
}

func dialServer(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, secondsForTimeout*time.Second)
	if err != nil {
		return fmt.Errorf("failed to dial test server, error: %v", err)
	}
	conn.Close()
	return nil
}

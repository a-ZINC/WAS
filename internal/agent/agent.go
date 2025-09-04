package agent

import (
	"fmt"
	"net"
	"time"

	"github.com/a-ZINC/WAS/internal/communication/transport"
	"github.com/a-ZINC/WAS/internal/discovery"
)


func StartAgent() {
	fmt.Println("Starting agent...")
	addr, err := startDiscovery()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Discovered broker at %s\n", addr)
	conn, err := connectToBroker(addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		// Keep the connection alive
		fmt.Printf("Connected to broker successfully: %s \n", conn.RemoteAddr().String())
		time.Sleep(10 * time.Second)
	}
}

func startDiscovery() (string, error) {
	discoverer := discovery.NewUDPDiscovery("localhost", 9999)
	addr, err := discoverer.Listen()
	if err != nil {
		return "", err
	}
	return addr, nil
}

func connectToBroker(addr string) (net.Conn, error) {
	tcp := transport.NewTCPTransport(addr)
	conn, err := tcp.Dial(addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
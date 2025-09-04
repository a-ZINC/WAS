package broker

import (
	"fmt"
	"log"
	"net"

	"github.com/a-ZINC/WAS/cmd/was"
	"github.com/a-ZINC/WAS/internal/communication/transport"
	"github.com/a-ZINC/WAS/internal/discovery"
)

func StartBroker() error {

	port, err := startTcpServer()
	if err != nil {
		return err
	}
	log.Printf("Broker started at port %d\n", port)
	err = startDiscovery(port)
	if err != nil {
		return err
	}

	return nil
}

func startTcpServer() (int, error) {
	fmt.Println("Starting TCP server...")
	server := transport.NewTCPTransport(fmt.Sprintf(":%d", was.Port))
	err := server.Start()
	port := server.Listener.Addr().(*net.TCPAddr).Port
	return port, err
}

func startDiscovery(port int) error {
	discoverer := discovery.NewUDPDiscovery("255.255.255.255", 9999)
	err := discoverer.Start(port)
	return err
}

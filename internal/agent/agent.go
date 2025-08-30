package agent

import (
	"fmt"

	"github.com/a-ZINC/WAS/internal/discovery"
)


func StartAgent() {
	fmt.Println("Starting agent...")
	discoverer := discovery.NewUDPDiscovery("localhost", 9999)
	addr, err := discoverer.Listen()
	if err != nil {
		return
	}
	fmt.Printf("Agent listening on %s\n", addr)
}
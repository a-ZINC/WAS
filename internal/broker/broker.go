package broker

import (
	"fmt"

	"github.com/a-ZINC/WAS/internal/discovery"
)

func StartBroker(port int) {
	fmt.Printf("Starting broker on port %d\n", port)
	discoverer := discovery.NewUDPDiscovery("255.255.255.255", 9999)
	discoverer.Start(port)
}

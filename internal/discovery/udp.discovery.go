package discovery

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"
)

type UDPDiscovery struct {
	Addr string
	Port int
}

func NewUDPDiscovery(addr string, port int) *UDPDiscovery {
	return &UDPDiscovery{
		Addr: addr,
		Port: port,
	}
}

// in macos 255.255.255.255 is blocked therfore calculating broadcast ip
func (u *UDPDiscovery) DiscoverBroadcastIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("Error getting network interfaces: %v\n", err)
		return
	}
	fmt.Printf("Found broadcast IP: %+v\n", addrs)
	for _, address := range addrs {
		ipnet, ok := address.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}
		ip := ipnet.IP.To4()
		if ip != nil {
			broadcast := make(net.IP, len(ip))
			broadcastStr := ""
			for i := range ip {
				broadcast[i] = ip[i] | ^ipnet.Mask[i]
				broadcastStr += fmt.Sprintf("%d.", broadcast[i])
			}
			broadcastStr = strings.TrimSuffix(broadcastStr, ".")
			fmt.Printf("Calculated broadcast IP: %s\n", broadcastStr)
			u.Addr = broadcastStr
		}
	}
}

func (u *UDPDiscovery) Start(port int) error {
	addr := net.UDPAddr{
		IP:   net.ParseIP(u.Addr),
		Port: u.Port,
	}
	fmt.Printf("UDPDiscovery Start called with addr: %s, port: %d\n", u.Addr, u.Port)
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return err
	}
	fmt.Printf("UDP discovery started on %s:%d\n", u.Addr, u.Port)
	defer conn.Close()

	buff := make([]byte, 1024)
	ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
	for {
		msg := fmt.Sprintf(`BROADCAST:%d`, port)
		copy(buff, msg)
		fmt.Printf("Broadcasting message: %s\n", msg)
		_, err := conn.WriteToUDP(buff[:len(msg)], &addr)
		if err != nil {
			fmt.Printf("Error sending UDP broadcast: %v\n", err)
			u.DiscoverBroadcastIP()
			u.Start(port)
			continue
		}
		<-ticker.C
	}
}

func (u *UDPDiscovery) Listen() (string, error) {
	addr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: u.Port,
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return "", err
	}
	fmt.Printf("Listening for UDP broadcasts on %s:%d\n", addr.IP.String(), addr.Port)
	defer conn.Close()

	for {
		buff := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buff)
		if err != nil {
			return "", err
		}
		fmt.Printf("Received UDP broadcast from %s: %s\n", addr.IP.String(), string(buff[:n]))
		msgStr := string(buff[:n])
		if bytes.HasPrefix(buff[:n], []byte("BROADCAST:")) {
			port := strings.TrimPrefix(msgStr, "BROADCAST:")
			ip := addr.IP
			return ip.String() + ":" + port, nil
		}
	}
}
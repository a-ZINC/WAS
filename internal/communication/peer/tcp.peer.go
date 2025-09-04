package peer

import "net"

type TCPPeer struct {
	net.Conn
}

func NewTCPPeer(conn net.Conn) *TCPPeer {
	return &TCPPeer{Conn: conn}
}

func (p *TCPPeer) Close() error {
	return p.Conn.Close()
}
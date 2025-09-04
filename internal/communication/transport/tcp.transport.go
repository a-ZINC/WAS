package transport

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"syscall"

	"github.com/a-ZINC/WAS/internal/communication/peer"
)

type TCPTransport struct {
	Address  string
	Listener net.Listener

	peerMu sync.Mutex
	peers  map[string]peer.Peer
}

func NewTCPTransport(address string) *TCPTransport {
	return &TCPTransport{
		Address: address,
		peers:   make(map[string]peer.Peer),
	}
}

func (t *TCPTransport) Start() error {
	listener, err := net.Listen("tcp", t.Address)
	if err != nil {
		switch err.(type) {
		case *net.OpError:
			if syscallErr, ok := err.(*net.OpError).Err.(*os.SyscallError); ok {
				if syscallErr.Err == syscall.EADDRINUSE {
					log.Printf("Port %s in use, trying alternative port\n", t.Address)
					t.retryServerStart()
					return nil
				}
			}
			return err
		default:
			return err
		}
	}
	fmt.Printf("TCP server started on %s\n", listener.Addr().String())
	t.Listener = listener
	go t.accept()
	return nil
}

func (t *TCPTransport) accept() error {
	defer t.Listener.Close()
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			return err
		}
		go t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) error {
	defer conn.Close()
	peer := peer.NewTCPPeer(conn)
	t.peerMu.Lock()
	t.peers[peer.RemoteAddr().String()] = peer
	t.peerMu.Unlock()
	return nil
}

func (t *TCPTransport) Dial(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (t *TCPTransport) retryServerStart() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return
	}
	log.Printf("Port in use, started TCP server on alternative port: %s\n", listener.Addr().String())
	defer listener.Close()
	t.Listener = listener
	t.Address = listener.Addr().String()
	go t.accept()
}

func (t *TCPTransport) Close() error {
	if t.Listener != nil {
		return t.Listener.Close()
	}
	return nil
}

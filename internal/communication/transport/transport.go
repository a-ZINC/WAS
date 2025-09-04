package transport

type Transport interface {
	Start() error
	Dial(addr string) error
	accept() error
	handleConnection() error
	Close() error
}
package discovery

type Discovery interface {
	Start() error
	Listen() (string, error)
}
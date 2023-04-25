package devdb

type EmbeddedDatabase interface {
	Start() error
	Stop() error

	Port() uint32
	Username() string
	Password() string
}

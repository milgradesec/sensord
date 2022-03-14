package serial

type Reader interface {
	Start()
	Close() error
}

package message

type Message interface {
	Get() ([]byte, error)
}

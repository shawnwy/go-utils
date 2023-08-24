package stream

// IMessage is a wrapper for stream message transmitting
type IMessage interface {
	UID() []byte   // UID is a key used for message dispatching
	Bytes() []byte // Bytes return all payload of the message
}

type RawMessage []byte

func (m RawMessage) UID() []byte {
	return nil
}

func (m RawMessage) Bytes() []byte {
	return m
}

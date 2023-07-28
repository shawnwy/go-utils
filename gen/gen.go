package gen

import (
	crand "crypto/rand"
	"encoding/binary"
	"io"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

var (
	rander     = crand.Reader
	characters = []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
)

// RandomBytes - generate random bytes concisely which is costly
func RandomBytes(b []byte) {
	if _, err := io.ReadFull(rander, b); err != nil {
		panic(err.Error()) // rand should never fail
	}
}

// RandomBytesLite - generate random bytes which is performance utilized
func RandomBytesLite(b []byte) {
	for i := range b {
		b[i] = characters[mrand.Intn(len(characters))]
	}
}

type source struct{}

func (s *source) Seed(seed int64) {}

func (s *source) Uint64() (value uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

func (s *source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

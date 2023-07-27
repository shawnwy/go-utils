package gen

import (
	"crypto/rand"
	"io"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

var (
	rander     = rand.Reader
	characters = []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
)

// RandomBytes - generate random bytes strictly which is costly
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

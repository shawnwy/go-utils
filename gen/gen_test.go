package gen

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var (
	x = make([]byte, 20)
)

func TestRandomBytes(t *testing.T) {
	b := make([]byte, 6)
	RandomBytesLite(b)
	fmt.Println(string(b), len(b))
	fmt.Println(hex.EncodeToString(b))
}

func BenchmarkRandomBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomBytes(x)
	}
}

func BenchmarkRandomBytesLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomBytesLite(x)
	}
}

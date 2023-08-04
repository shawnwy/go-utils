package machines

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var cid = ClientID()

func TestClientID(t *testing.T) {
	cid := ClientID()
	fmt.Println(hex.EncodeToString(cid[:]), len(cid))
}

func BenchmarkClientID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClientID()
	}
}

func BenchmarkClientIDHexStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hex.EncodeToString(cid[:])
	}
}

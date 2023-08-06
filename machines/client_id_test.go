package machines

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

var cid = ClientID()

func TestClientID(t *testing.T) {
	cid := ClientID()
	cidStr := hex.EncodeToString(cid[:])
	fmt.Println(cidStr, len(cid), len(cidStr))
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

func TestCopy2RetVal(t *testing.T) {
	x := copy2retVal()
	if !bytes.Equal(x[:], cid[:]) {
		t.FailNow()
	}
}

func copy2retVal() (x [20]byte) {
	copy(x[:], cid[:])
	return
}

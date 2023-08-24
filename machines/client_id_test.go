package machines

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
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

func TestGetHostname(t *testing.T) {
	hostname, err := os.Hostname()
	fmt.Println(hostname, err)
	v4, _ := uuid.NewV4()
	fmt.Println(v4.Bytes(), len(v4.Bytes()))
}

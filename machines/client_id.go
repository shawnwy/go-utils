package machines

import (
	"encoding/binary"
	"os"
	"time"

	"github.com/shawnwy/go-utils/v5/gen"
)

func ClientID() (cid [20]byte) {
	mac, err := MacAddr()
	if err != nil {
		mac = make([]byte, 6)
		gen.RandomBytes(mac)
	}
	binary.BigEndian.PutUint32(cid[:], uint32(os.Getpid()))           // 4 bytes
	binary.BigEndian.PutUint64(cid[:], uint64(time.Now().UnixNano())) // 8 bytes
	copy(cid[12:], mac)
	return
}

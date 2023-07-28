package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/minio/highwayhash"
)

func Secret(secret []byte, t time.Time) []byte {
	payload := make([]byte, len(secret)+8)
	binary.BigEndian.PutUint64(payload, uint64(t.UnixNano()))
	copy(payload[8:], secret)
	return payload
}

func Signature(secret, key []byte) [32]byte {
	return highwayhash.Sum(secret, key)
}

func HmacSignature(secret, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(secret)
	return h.Sum(nil)
}

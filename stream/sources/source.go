package sources

import "github.com/shawnwy/go-utils/v5/stream"

const (
	rawSrcChanSize = 1 << 4
)

type RawSource interface {
	Start()
	RawBytes() chan stream.IMessage
	Close()
}

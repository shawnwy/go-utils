package rotatelogs

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shawnwy/go-utils/v5/errors"
)

var (
	rfc3999Layout = "2006-01-02T15-04-05.999999999"
	compactLayout = "20060102150405.000000000"
	cst, _        = time.LoadLocation("Asia/Shanghai")
)

type FilenameParser func(filename string) (node string, ts time.Time, err error)

// Msg
func MsgBytesFnParser(filename string) (node string, ts time.Time, err error) {
	prefix, t, ok := strings.Cut(filename, "-")
	node = prefix
	if !ok {
		err = ErrInvalidFn
		return
	}
	ts, err = time.Parse(rfc3999Layout, t)
	return
}

// RFC3999FnParser - parse filename by format {date}{millis}_{nodeID}
func RFC3999FnParser(filename string) (node string, ts time.Time, err error) {
	fn := strings.TrimSuffix(filepath.Base(filename), PcapExt)
	fnDetails := strings.Split(fn, "_")
	if len(fnDetails) < 2 {
		err = errors.Wrap(ErrInvalidFn, "irregular filename of pcapfile with less info.")
		return
	}
	node = fnDetails[1]
	if len(fnDetails[0]) < 27 {
		err = errors.Wrapf(ErrInvalidFn, "part of time info is invalid[%s]", fnDetails[0])
		return
	}
	milli, err := strconv.ParseInt(fnDetails[0][14:], 10, 64)
	if err != nil {
		err = errors.Wrapf(ErrInvalidFn, "invalid millis secs[%s]", fnDetails[0][14:])
		return
	}
	ts = time.UnixMilli(milli)
	return
}

// CompactFnParser - parse filename by format {datetime}_{nodeID}
func CompactFnParser(filename string) (node string, ts time.Time, err error) {
	fn := strings.TrimSuffix(filepath.Base(filename), PcapExt)
	fnDetails := strings.Split(fn, "_")
	if len(fnDetails) < 2 {
		err = errors.Wrap(ErrInvalidFn, "irregular filename of pcapfile with less info.")
		return
	}
	node = fnDetails[0]
	if len(fnDetails[1]) < 23 {
		err = errors.Wrapf(ErrInvalidFn, "part of time info is invalid[%s]", fnDetails[0])
		return
	}
	ts, err = time.ParseInLocation(compactLayout, fnDetails[1], cst)
	if err != nil {
		err = errors.Wrapf(ErrInvalidFn, "invalid time string[%s]", fnDetails[1])
		return
	}
	return
}

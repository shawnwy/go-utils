package sizes

import (
	"fmt"
	"strings"
)

type size int64

// constant of data unit base on Bytes
const (
	B size = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func (s size) String() string {
	switch s {
	case B:
		return "B"
	case KB:
		return "KB"
	case MB:
		return "MB"
	case GB:
		return "GB"
	case TB:
		return "TB"
	case PB:
		return "PB"
	default:
		return "Unknown"
	}
}

func (s size) AsInt() int {
	if s > GB {
		panic(fmt.Sprintf("cannot covert to int. %s will overflow", s.String()))
	}
	return int(s)
}

func (s size) AsInt32() int32 {
	if s > GB {
		panic(fmt.Sprintf("cannot covert to int32. %s will overflow", s.String()))
	}
	return int32(s)
}

func (s size) AsInt64() int64 {
	return int64(s)
}

func (s size) AsFloat64() float64 {
	return float64(s)
}

// LiteralStdout literal print as traffic uint
func LiteralStdout(bytes int64) string {
	template := "%.2f"
	if bytes/PB.AsInt64() > 0 {
		return fmt.Sprintf(template, float64(bytes)/PB.AsFloat64())
	}
	if bytes/TB.AsInt64() > 0 {
		return fmt.Sprintf(template, float64(bytes)/TB.AsFloat64())
	}
	if bytes/GB.AsInt64() > 0 {
		return fmt.Sprintf(template, float64(bytes)/GB.AsFloat64())
	}
	if bytes/GB.AsInt64() > 0 {
		return fmt.Sprintf(template, float64(bytes)/GB.AsFloat64())
	}
	if bytes/MB.AsInt64() > 0 {
		return fmt.Sprintf(template, float64(bytes)/MB.AsFloat64())
	}
	if bytes/KB.AsInt64() > 0 {
		return fmt.Sprintf(template, float64(bytes)/KB.AsFloat64())
	}
	return fmt.Sprintf(template, float64(bytes)/B.AsFloat64())
}

// LiteralCompactStdout literal print as traffic uint and trim '.00'
func LiteralCompactStdout(bytes int64) string {
	return strings.TrimSuffix(LiteralStdout(bytes), ".00")
}

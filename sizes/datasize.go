package sizes

// constant of data unit base on Bytes
const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

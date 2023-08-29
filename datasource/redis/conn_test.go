package redis

import "testing"

func TestConn(t *testing.T) {
	Connect("redis://10.10.20.21:6379/0")
}

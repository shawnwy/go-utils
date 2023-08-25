package machines

import (
	"fmt"
	"testing"
)

func TestOutboundNIC(t *testing.T) {
	macs, ips, err := OutboundNICs()
	fmt.Println(macs, ips, err)
}

package machines

import (
	"bytes"
	"net"

	"github.com/shawnwy/go-utils/v5/errors"
)

// MacAddr obtain MAC address as bytes. only accept IEEE MAC-48, EUI-48 and EUI-64 form
// usually the length of MAC address is 6 or 8
func MacAddr() ([]byte, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch ifaces")
	}
	for _, ifs := range ifaces {
		if ifs.Flags&net.FlagUp != 0 &&
			bytes.Compare(ifs.HardwareAddr, nil) != 0 &&
			len(ifs.HardwareAddr) >= 6 && len(ifs.HardwareAddr) <= 8 {
			return ifs.HardwareAddr, nil
		}
	}
	return nil, errors.New("cannot find available interfaces")
}

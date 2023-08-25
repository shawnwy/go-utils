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

func OutboundNICs() (macs []net.HardwareAddr, ips []net.IP, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to fetch ifaces")
	}

	for _, ifs := range ifaces {
		if (ifs.Flags&net.FlagUp) == 0 ||
			bytes.Compare(ifs.HardwareAddr, nil) == 0 {
			continue
		}

		var addrs []net.Addr
		addrs, err = ifs.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() {
				continue
			}

			ip := ipnet.IP.To4()
			if ip == nil {
				continue
			}
			macs = append(macs, ifs.HardwareAddr)
			ips = append(ips, ip)
		}
	}
	if len(macs) == 0 {
		err = errors.New("none available outbound NIC.")
	}
	return
}

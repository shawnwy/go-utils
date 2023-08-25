package machines

import (
	"net"

	"github.com/shawnwy/go-utils/v5/errors"
)

func GetOutboundIP() (ip net.IP, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return nil, errors.Wrap(err, "failed fetch ip from udp dns")
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func GetLocalIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch ip address from iface")
	}
	for _, address := range addrs {
		// check the address type and if it is not a loop-back the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}

	}
	return GetOutboundIP()
}

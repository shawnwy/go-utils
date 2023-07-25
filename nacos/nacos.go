package nacos

import (
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/common/constant"

	"github.com/shawnwy/go-utils/v5/errors"
)

var (
	ErrNacosParseURL    = errors.New("failed to parse nacos urls")
	ErrRegisterFailed   = errors.New("failed to register service")
	ErrDeregisterFailed = errors.New("failed to deregister service")
	ErrReadConfigFailed = errors.New("failed to read config")
)

func ServerCFG(urls string) ([]constant.ServerConfig, error) {
	var serverCFG []constant.ServerConfig
	for _, url := range strings.Split(urls, ";") {
		u := strings.Split(url, ":")
		port, err := strconv.ParseUint(u[1], 10, 64)
		if err != nil {
			return nil, errors.With(err, ErrNacosParseURL)
		}
		serverCFG = append(serverCFG, *constant.NewServerConfig(u[0], port, constant.WithScheme("http")))
	}

	return serverCFG, nil
}

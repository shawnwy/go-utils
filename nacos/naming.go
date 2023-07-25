package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/shawnwy/go-utils/v5/errors"
)

var namingCLI naming_client.INamingClient

// InitNamingCLI create a nacos Naming Client for service registry and discovery
func InitNamingCLI(urls, usr, pwd string) (err error) {
	serverCFG, err := ServerCFG(urls)
	if err != nil {
		return err
	}
	if namingCLI, err = clients.NewNamingClient(vo.NacosClientParam{
		ServerConfigs: serverCFG,
		ClientConfig: constant.NewClientConfig(
			constant.WithNotLoadCacheAtStart(true),
			constant.WithPassword(pwd),
			constant.WithUsername(usr),
		),
	}); err != nil {
		return err
	}
	return nil
}

func RegisterService(name, group, host string, port uint64) error {
	if success, err := namingCLI.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        port,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    nil,
		ServiceName: name,
		GroupName:   group,
		Ephemeral:   true,
	}); !success {
		return errors.With(err, ErrRegisterFailed)
	}
	return nil
}

func DeregisterService(name, group, host string, port uint64) error {
	if success, err := namingCLI.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          host,
		Port:        port,
		ServiceName: name,
		Ephemeral:   true,
		GroupName:   group, // default value is DEFAULT_GROUP
	}); !success {
		return errors.With(err, ErrDeregisterFailed)
	}
	return nil
}

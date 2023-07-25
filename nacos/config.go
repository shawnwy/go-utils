package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/errors"
)

var configCLI config_client.IConfigClient

// InitConfigCLI create a nacos Config Client for manipulate configurations
func InitConfigCLI(urls, usr, pwd string) (err error) {
	serverCFG, err := ServerCFG(urls)
	if err != nil {
		return err
	}
	if configCLI, err = clients.NewConfigClient(vo.NacosClientParam{
		ServerConfigs: serverCFG,
		ClientConfig: constant.NewClientConfig(
			constant.WithNotLoadCacheAtStart(true),
			constant.WithPassword(pwd),
			constant.WithUsername(usr),
			constant.WithLogDir("../.tmp/nacos"),
			constant.WithCacheDir("../.tmp/cache"),
		),
	}); err != nil {
		return err
	}
	return nil
}

func ReadConfig(dataID, group string) (string, error) {
	return configCLI.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
	})
}

func PublishConfig(dataID, group string, content string) error {
	if _, err := configCLI.PublishConfig(vo.ConfigParam{
		DataId:  dataID,
		Group:   group,
		Content: content,
	}); err != nil {
		return err
	}
	return nil
}

func WatchConfig(dataID, group string, cb func(data string) error) error {
	content, err := ReadConfig(dataID, group)
	if err != nil {
		return errors.With(err, ErrReadConfigFailed)
	}
	if err = cb(content); err != nil {
		return errors.Wrap(err, "failed to initial process the data from nacos")
	}
	if err = configCLI.ListenConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			if err := cb(data); err != nil {
				zap.L().Warn("failed to process changed config data",
					zap.String("dataID", dataId),
					zap.String("group", group),
				)
				return
			}
		},
	}); err != nil {
		return errors.Wrap(err, "failed to listen on current config from nacos")
	}
	return nil
}

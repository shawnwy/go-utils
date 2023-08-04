package configs

import (
	"strings"

	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/nacos"
)

func ReadConfigFromNacos(dataID, group, ext string) {
	content, err := nacos.ReadConfig(dataID, group)
	if err != nil {
		zap.L().Panic("Failed to fetch config from nacos!!!", zap.Error(err))
	}
	config.SetConfigType(ext)
	err = config.MergeConfig(strings.NewReader(content))
	if err != nil {
		zap.L().Panic("Failed to load to Viper!!!", zap.Error(err))
	}
}

func WatchOnNacos(dataID, group string, ext string, callback func() error) {
	ReadConfigFromNacos(dataID, group, ext)
	if err := nacos.WatchConfig(dataID, group, func(data string) error {
		if err := config.MergeConfig(strings.NewReader(data)); err != nil {
			zap.L().Warn("failed to load content to Viper! Corrupt format", zap.Error(err))
			return err
		}
		if err := callback(); err != nil {
			zap.L().Warn("failed to do callback for configs watching @ Nacos", zap.Error(err))
			return err
		}
		return nil
	}); err != nil {
		zap.L().Panic("Failed to Listen Config!!!", zap.String("dataID", dataID))
	}
}

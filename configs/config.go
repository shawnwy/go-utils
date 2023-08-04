package configs

import "github.com/spf13/viper"

var config *viper.Viper

func init() {
	config = viper.New()
}

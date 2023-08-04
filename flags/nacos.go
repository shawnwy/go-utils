package flags

import (
	"flag"

	"github.com/shawnwy/go-utils/v5/errors"
)

type NacosCFG struct {
	NacosURI string `json:"nacos-uri"`
	NacosUsr string `json:"nacos-usr"`
	NacosPwd string `json:"nacos-pwd"`
}

func (n NacosCFG) Validate() error {
	if n.NacosURI == "" {
		return nil
	}
	if n.NacosUsr == "" || n.NacosPwd == "" {
		return errors.New("missing nacos parameters ...")
	}
	return nil
}

func DefineNacosFlags(cfg *NacosCFG, cmd string) {
	flag.StringVar(&cfg.NacosURI, "nacos-uri", "", "Set Nacos as Remote Config Center. URI is needed")
	flag.StringVar(&cfg.NacosUsr, "nacos-usr", "", "Set Nacos as Remote Config Center. username")
	flag.StringVar(&cfg.NacosPwd, "nacos-pwd", "", "Set Nacos as Remote Config Center. password")
}

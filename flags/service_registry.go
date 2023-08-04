package flags

import (
	"flag"
	"fmt"

	"github.com/shawnwy/go-utils/v5/errors"
)

type RegistryCFG struct {
	Registry bool   `json:"registry"`
	RegPort  int    `json:"reg-port"`
	RegGroup string `json:"reg-group"`
}

func (r RegistryCFG) Validate() error {
	if r.Registry && (r.RegGroup == "" || r.RegPort == 0) {

		return errors.New("missing service register info ...")
	}
	return nil
}

func DefineRegistryFlags(cfg *RegistryCFG, cmd string) {
	flag.BoolVar(&cfg.Registry, "reg", false, fmt.Sprintf("registry service <%s> or not", cmd))
	flag.IntVar(&cfg.RegPort, "reg-port", 59696, fmt.Sprintf("registry service <%s> with customized port", cmd))
	flag.StringVar(&cfg.RegGroup, "reg-group", "", fmt.Sprintf("registry service <%s> with customized group name", cmd))
}

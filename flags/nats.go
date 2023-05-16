package flags

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type NATsCFG struct {
	Server  string `json:"server"`
	Subject string `json:"subject"`
	Workers int    `json:"workers"`
}

func (n NATsCFG) Validate() error {
	if n.Server == "" {
		return errors.New("server has been missing in NATsCFG")
	}
	if n.Subject == "" {
		return errors.New("'subject' has been missing in NATsCFG")
	}
	if n.Workers <= 0 {
		return errors.New("'workers' has been missing in NATsCFG")
	}
	return nil
}

func DefineNATsFlags(cfg *NATsCFG, cmd, tag string) {
	flag.StringVar(&cfg.Server,
		fmt.Sprintf("%sserver", tag),
		"nats://localhost:4222",
		fmt.Sprintf(strings.ReplaceAll("Set the NATsComm server\n\t  %s --{tag}comm nats --{tag}server 'nats://localhost:4222' --{tag}subject 'sub.1' --{tag}workers 3", "{tag}", tag), cmd))
	flag.StringVar(&cfg.Subject,
		fmt.Sprintf("%ssubject", tag),
		"",
		fmt.Sprintf(strings.ReplaceAll("Set the NATsComm subject\n\t  %s --{tag}comm nats --{tag}server 'nats://localhost:4222' --{tag}subject 'sub.1' --{tag}workers 3", "{tag}", tag), cmd))
	flag.IntVar(&cfg.Workers,
		fmt.Sprintf("%sworkers", tag),
		1,
		fmt.Sprintf(strings.ReplaceAll("Set the NATsComm workers\n\t  %s --{tag}comm nats --{tag}server 'nats://localhost:4222' --{tag}subject 'sub.1' --{tag}workers 3", "{tag}", tag), cmd))
}

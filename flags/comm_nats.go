package flags

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

type NATsCommCFG struct {
	Server  string `json:"server"`
	Subject string `json:"subject"`
	Queue   string `json:"queue"`
	Workers int    `json:"workers"`
}

func (n NATsCommCFG) Validate() error {
	if n.Server == "" {
		return errors.New("server has been missing in NATsCommCFG")
	}
	if n.Subject == "" {
		return errors.New("'subject' has been missing in NATsCommCFG")
	}
	if n.Workers <= 0 {
		return errors.New("'workers' has been missing in NATsCommCFG")
	}
	return nil
}

func DefineNATsFlags(cfg *NATsCommCFG, cmd, tag string) {
	tag += "-"
	flag.StringVar(&cfg.Server,
		fmt.Sprintf("%sserver", tag),
		"nats://localhost:4222",
		fmt.Sprintf(strings.ReplaceAll("Set the NATsComm server\n\t  %s --{tag}comm nats --{tag}server 'nats://localhost:4222' --{tag}subject 'sub.1' --{tag}workers 3", "{tag}", tag), cmd))
	flag.StringVar(&cfg.Subject,
		fmt.Sprintf("%ssubject", tag),
		"",
		fmt.Sprintf(strings.ReplaceAll("Set the NATsComm subject\n\t  %s --{tag}comm nats --{tag}server 'nats://localhost:4222' --{tag}subject 'sub.1' --{tag}workers 3", "{tag}", tag), cmd))
	flag.StringVar(&cfg.Queue,
		fmt.Sprintf("%squeue", tag),
		fmt.Sprintf("%s%d", tag, time.Now().UnixNano()),
		fmt.Sprintf(strings.ReplaceAll("Set the NATsComm queue\n\t  %s --{tag}comm nats --{tag}server 'nats://localhost:4222' --{tag}subject 'sub.1' --{tag}workers 3 --{tag}queue 'g1'", "{tag}", tag), cmd))
	flag.IntVar(&cfg.Workers,
		fmt.Sprintf("%sworkers", tag),
		1,
		fmt.Sprintf(strings.ReplaceAll("Set the NATsComm workers\n\t  %s --{tag}comm nats --{tag}server 'nats://localhost:4222' --{tag}subject 'sub.1' --{tag}workers 3", "{tag}", tag), cmd))
}

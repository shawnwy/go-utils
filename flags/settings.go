package flags

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ISettings interface {
	Validate() error
	AttachPprof(h http.Handler)
}

func ParseFlags(s ISettings, usage func()) {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println(">> !!none arguments detected!! <<")
		usage()
		os.Exit(2)
	}
	flag.Parse()
	if err := s.Validate(); err != nil {
		fmt.Println(">> !!settings is not valid!! <<\n", err)
		os.Exit(2)
	}
	s.AttachPprof(nil)
}

func CreateUsage(authors string, desc ...string) func() {
	return func() {
		descOfCMD := strings.Join(desc, "\r\n")
		descOfAuthors := fmt.Sprintf("Authors: %s", authors)
		fmt.Println(strings.Join([]string{
			"",
			descOfCMD,
			"",
			descOfAuthors,
			"",
		}, "\r\n"))
		flag.PrintDefaults()
		os.Exit(2)
	}
}

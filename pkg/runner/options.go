package runner

import (
	"os"
	"strings"

	"github.com/Henrique-Gomesz/JoeyScan4Me/pkg/logging"

	"github.com/projectdiscovery/goflags"
)

type Options struct {
	Domain  string
	Workdir string
	Server  bool
}

func ParseOptions() *Options {
	opt := &Options{}
	flagSet := goflags.NewFlagSet()

	flagSet.StringVar(&opt.Domain, "d", "", "domain to scan (e.g. example.com)")
	flagSet.StringVar(&opt.Workdir, "w", "./", "working directory for output files, defaults to current directory")
	flagSet.BoolVar(&opt.Server, "server", false, "start gowitness server at the end of scan to view screenshots")

	if err := flagSet.Parse(); err != nil {
		logging.LogError("Error parsing flags:", err)
		os.Exit(1)
	}

	validateDomain(opt)

	return opt
}

func validateDomain(opt *Options) {
	if opt.Domain == "" {
		logging.LogError("Domain is required. Use -d flag to specify a domain (e.g., -d example.com)", nil)
		os.Exit(1)
	}

	opt.Domain = strings.TrimPrefix(opt.Domain, "https://")
	opt.Domain = strings.TrimPrefix(opt.Domain, "http://")

	opt.Domain = strings.TrimSuffix(opt.Domain, "/")
}
